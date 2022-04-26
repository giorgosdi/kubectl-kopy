package cmd

import (
	"fmt"
	"os"

	"github.com/giorgosdi/kubectl-kopy/pkg/kopy"
	"github.com/giorgosdi/kubectl-kopy/pkg/options"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func NewKopyCommand(streams genericclioptions.IOStreams) *cobra.Command {

	o := options.NewKopyOptions(streams)

	// cmd represents the nodeinfo command
	var cmd = &cobra.Command{
		Use:          "kubectl nodeinfo <node> [flags]",
		Short:        "Information about a given node",
		Args:         cobra.MaximumNArgs(3),
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			o.Complete(cmd, args)
			_, err := os.Stat(o.Kubeconfig)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			Kopy(cmd, o, args)
		},
	}
	o.Config.AddFlags(cmd.Flags())

	return cmd
}

func Kopy(cmd *cobra.Command, o *options.KopyOptions, args []string) {
	kopy.KopyObject(o)
}
