package options

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type KopyOptions struct {
	Namespace  string
	Kubeconfig string
	Config     *genericclioptions.ConfigFlags
	Name       string
	Kind       string
	Args       []string

	genericclioptions.IOStreams
}

func NewKopyOptions(streams genericclioptions.IOStreams) *KopyOptions {
	return &KopyOptions{
		Config: genericclioptions.NewConfigFlags(true),

		IOStreams: streams,
	}
}

func (o *KopyOptions) Complete(c *cobra.Command, args []string) error {
	var err error
	homeDir, _ := os.UserHomeDir()
	o.Namespace, err = c.Flags().GetString("namespace")
	if err != nil {
		return err
	}
	o.Kubeconfig = fmt.Sprintf("%s/.kube/config", homeDir)

	if value, _ := c.Flags().GetString("kubeconfig"); value != "" {
		o.Kubeconfig, err = c.Flags().GetString("kubeconfig")
		if err != nil {
			return err
		}
	}
	if os.Getenv("KUBECONFIG") != "" {
		o.Kubeconfig = os.Getenv("KUBECONFIG")
	}
	o.Kind = args[0]
	o.Name = args[1]
	o.Args = args
	return nil
}
