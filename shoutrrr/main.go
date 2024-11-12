package main

import (
	"os"

	"github.com/devops-golang/shoutrrr/cmd/docs"
	"github.com/devops-golang/shoutrrr/cmd/generate"
	"github.com/devops-golang/shoutrrr/cmd/send"
	"github.com/devops-golang/shoutrrr/cmd/verify"
	"github.com/devops-golang/shoutrrr/internal/meta"
	cli "github.com/devops-golang/shoutrrr/shoutrrr/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmd = &cobra.Command{
	Use:     "shoutrrr",
	Version: meta.Version,
	Short:   "Shoutrrr CLI",
}

func init() {
	viper.AutomaticEnv()
	cmd.AddCommand(verify.Cmd)
	cmd.AddCommand(generate.Cmd)
	cmd.AddCommand(send.Cmd)
	cmd.AddCommand(docs.Cmd)
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(cli.ExUsage)
	}
}
