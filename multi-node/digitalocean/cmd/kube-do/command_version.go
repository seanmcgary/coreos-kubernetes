package main

import (
	"fmt"

	"github.com/coreos/coreos-kubernetes/multi-node/digitalocean/pkg/cluster"
	"github.com/spf13/cobra"
)

var (
	cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Print version information and exit",
		Long:  ``,
		Run:   runCmdVersion,
	}
)

func init() {
	cmdRoot.AddCommand(cmdVersion)
}

func runCmdVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("kube-do version %s\n", cluster.VERSION)
}
