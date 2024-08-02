package web

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/telanflow/domain-deploy/infra/config"
	"github.com/telanflow/domain-deploy/infra/logger"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type CertificateRequest struct {
	Domains           []string `json:"domains"`
	SslCertificate    string   `json:"ssl_certificate"`
	SslCertificateKey string   `json:"ssl_certificate_key"`
	StartTime         string   `json:"start_time"`
	ExpireTime        string   `json:"expire_time"`
}

// IssueCertificateHandler 证书部署
func IssueCertificateHandler(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != config.GetDeploy().Token {
		_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	keySavePath := ctx.GetHeader("key-save-path")
	deployCmd := ctx.GetHeader("deploy-cmd")

	certReq := &CertificateRequest{}
	err := ctx.ShouldBindBodyWithJSON(certReq)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("Error decoding request body: %v", err))
		return
	}

	for _, domain := range certReq.Domains {
		certPath := fmt.Sprintf("%s/%s.pem", keySavePath, domain)
		keyPath := fmt.Sprintf("%s/%s.key", keySavePath, domain)

		err = os.WriteFile(certPath, []byte(certReq.SslCertificate), 0644)
		if err != nil {
			logger.Errorf("writing certificate for domain %s: %v", domain, err)
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		err = os.WriteFile(keyPath, []byte(certReq.SslCertificateKey), 0644)
		if err != nil {
			logger.Errorf("writing key for domain %s: %v", domain, err)
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	cmd, ok := config.GetDeploy().Cmds[deployCmd]
	if !ok {
		logger.Errorf("invalid deploy-cmd: %s", deployCmd)
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid deploy-cmd: %s", deployCmd))
		return
	}

	// parts := strings.Fields(cmd)
	// out, err := exec.Command(parts[0], parts[1:]...).Output()
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		logger.Errorf("command execution failed: %s\nError: %v", cmd, err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("command execution failed: %s", deployCmd))
		return
	}

	logger.Infof("Certificate issued and command executed successfully for domains: %s. Command: %s", strings.Join(certReq.Domains, ", "), cmd)
	logger.Infof("cmd run result %s", out)
}
