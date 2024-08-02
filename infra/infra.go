package infra

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gookit/goutil/fsutil"
)

var (
	ServiceTemplate = `[Unit]
Description=Domain-Admin deployment service
Documentation=https://github.com/telanflow/domain-deploy
After=network.target nss-lookup.target

[Service]
Type=simple
User=root
NoNewPrivileges=true
WorkingDirectory={{.ExeDir}}
ExecStart={{.ExePath}} -c {{.ExeDir}}/config.yml
Restart=on-failure
RestartSec=5s
TimeoutSec=60
RestartPreventExitStatus=23
LimitNPROC=10000
LimitNOFILE=1000000

[Install]
WantedBy=multi-user.target`
)

func InstallService(path string) error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	exeDir := filepath.Dir(exePath)

	// 创建模板
	serviceConfig := struct {
		ExePath string
		ExeDir  string
	}{
		ExePath: exePath,
		ExeDir:  exeDir,
	}
	tmpl, err := template.New("service").Parse(ServiceTemplate)
	if err != nil {
		return err
	}

	var output bytes.Buffer
	if err = tmpl.Execute(&output, serviceConfig); err != nil {
		return err
	}

	fs, err := fsutil.OpenFile(path, fsutil.FsCWTFlags, 0666)
	if err != nil {
		return err
	}
	defer fs.Close()

	// 写入服务内容
	_, err = fs.WriteString(output.String())
	if err != nil {
		return err
	}

	return nil
}
