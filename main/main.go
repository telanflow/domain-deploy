package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/judwhite/go-svc"
	"github.com/telanflow/domain-deploy/infra"
	"github.com/telanflow/domain-deploy/infra/config"
	"github.com/telanflow/domain-deploy/infra/logger"
	"github.com/telanflow/domain-deploy/main/distro"
	"github.com/urfave/cli/v2"
)

var (
	appName    string
	appVersion string
	gitHash    string
	buildTime  string
)

func main() {
	logger.Init()
	app := &cli.App{
		Name:        appName,
		Version:     appVersion,
		Usage:       appName + " for " + runtime.GOOS + "/" + runtime.GOARCH,
		Copyright:   fmt.Sprintf("(c) 2024-%s %s.\n", time.Now().Format("2006"), appName),
		Description: fmt.Sprintf("%s is the deployment service for domain-admin", appName),
		Commands: cli.Commands{
			&cli.Command{
				Name:  "version",
				Usage: "print the version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					fmt.Printf("Copyright (c) 2024-%s\n", time.Now().Format("2006"))
					return nil
				},
			},
			&cli.Command{
				Name:  "install",
				Usage: "install to service",
				Action: func(ctx *cli.Context) error {
					servicePath := fmt.Sprintf("/etc/systemd/system/%s.service", appName)
					if err := infra.InstallService(servicePath); err != nil {
						return err
					}

					_ = exec.Command("/bin/sh", "-c", "systemctl daemon-reload").Run()

					logger.Infof("install success to %s", servicePath)
					logger.Info("")
					logger.Infof("systemctl start %s", appName)
					logger.Infof("systemctl stop %s", appName)
					logger.Infof("systemctl restart %s", appName)
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "配置文件",
				DefaultText: "config.yml",
				Value:       "config.yml",
			},
		},
		Action: func(ctx *cli.Context) error {
			configFile := ctx.String("config")
			if err := initConfig(configFile); err != nil {
				return err
			}

			// 启动服务
			program := &distro.Program{}
			return svc.Run(program)
		},
	}
	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}

func initConfig(configFile string) error {
	if configFile == "" {
		return errors.New("配置文件不能为空")
	}
	// 配置初始化
	if err := config.InitForFile(configFile); err != nil {
		return err
	}
	// 日志初始化
	logger.InitForConfig(config.GetLog())
	return nil
}
