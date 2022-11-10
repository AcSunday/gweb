package main

import (
	"os"
	"runtime"

	"go_project_demo/library/clean"

	"github.com/urfave/cli"
)

var (
	cmd   *cli.App
	name  string
	btime string
	build string
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	setupCmd()
}

func main() {
	if err := cmd.Run(os.Args); err != nil {
		clean.ExitErr(err)
	}
	clean.Exit()
}
