package distro

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/judwhite/go-svc"
	"github.com/telanflow/domain-deploy/infra/config"
	"github.com/telanflow/domain-deploy/infra/logger"
)

type Program struct {
	srv *http.Server
}

func (p *Program) Init(env svc.Environment) error {
	ip, port := config.GetAddress()
	p.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", ip, port),
		Handler: LoadCore(),
	}
	return nil
}

func (p *Program) Start() error {
	logger.Info("Domain-deploy starting")
	go func() {
		if err := p.srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Fatal(err)
			}
		}
	}()

	ip, port := config.GetAddress()
	logger.Infof("listen: %s:%d", ip, port)
	return nil
}

func (p *Program) Stop() error {
	_ = p.srv.Close()
	logger.Info("Domain-deploy stopped")
	return nil
}
