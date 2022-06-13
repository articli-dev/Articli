package auth

import (
	"github.com/spf13/cobra"

	"github.com/k8scat/articli/internal/config"
	sfsdk "github.com/k8scat/articli/pkg/platform/segmentfault"
)

var (
	cfgFile string
	cfg     *config.Config
	client  *sfsdk.Client

	cmd = &cobra.Command{
		Use:   "auth",
		Short: "Manage authentication state of segmentfault.com",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			client, _ = sfsdk.NewClient(cfg.Platforms.SegmentFault.Token)
		},
	}
)

func init() {
	cmd.AddCommand(loginCmd)
	cmd.AddCommand(logoutCmd)
	cmd.AddCommand(statusCmd)
}

func NewCmd(cf string, c *config.Config) *cobra.Command {
	cfgFile = cf
	cfg = c
	return cmd
}
