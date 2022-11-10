package main

import (
	"fmt"
	"go_project_demo/common"
	"go_project_demo/library/log"
	"go_project_demo/router"

	"github.com/urfave/cli"
	"go_project_demo/component"
)

func setupCmd() {
	cmd = cli.NewApp()
	cmd.Name = name
	cmd.Version = build
	cmd.Commands = []cli.Command{
		{
			Name:     "start",
			HideHelp: true,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "c", Value: ".env", Usage: "config file"},
				cli.StringFlag{Name: "p", Value: "3000", Usage: "http listen port"},
			},
			Before: func(ctx *cli.Context) (err error) {
				err = setupComponent(ctx.String("c"), ctx.Int("p"))
				return
			},
			Action: func(ctx *cli.Context) (err error) {
				component.InfLogger.Info(log.F{
					"log_type":   common.LogTypeForAppStart,
					"name":       name,
					"version":    build,
					"build_time": btime,
				})

				engine := router.RegisterRouter()
				if err = component.HttpServer.Run(engine); err != nil {
					err = fmt.Errorf("common.HttpServer.Run: %w", err)
				}
				return
			},
		},
	}
}
