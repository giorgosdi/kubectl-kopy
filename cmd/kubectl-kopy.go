package cmd

import (
	"os"

	"github.com/giorgosdi/kubectl-kopy/pkg/cmd"

	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func KubectlKopy() {
	flags := pflag.NewFlagSet("kubectl-kopy", pflag.ExitOnError)
	pflag.CommandLine = flags
	kopyCmd := cmd.NewKopyCommand(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := kopyCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
