package cmd

import (
	"fmt"

	cli "github.com/spf13/cobra"

	"github.com/zenoss/boilr/pkg/boilr"
	"github.com/zenoss/boilr/pkg/util/tlog"
	"github.com/zenoss/boilr/pkg/util/validate"
)

// Version contains the cli-command for printing the current version of the tool.
var Version = &cli.Command{
	Use:   "version",
	Short: "Show the boilr version information",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{})

		shouldntPrettify := GetBoolFlag(c, "dont-prettify")
		if shouldntPrettify {
			fmt.Println(boilr.Version)
		} else {
			tlog.Info(fmt.Sprint("Current version is ", boilr.Version))
		}
	},
}
