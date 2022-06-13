package article

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/k8scat/articli/internal/config"
	sfsdk "github.com/k8scat/articli/pkg/platform/segmentfault"
)

var (
	client *sfsdk.Client
	cfg    *config.Config

	articleCmd = &cobra.Command{
		Use:   "article",
		Short: "Manage articles",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			client, _ = sfsdk.NewClient(cfg.Platforms.Juejin.Cookie)
			if client == nil {
				log.Fatal("please login first")
			}
		},
	}
)

func init() {
	articleCmd.AddCommand(listCmd)
	articleCmd.AddCommand(createCmd)
	articleCmd.AddCommand(viewCmd)
	articleCmd.AddCommand(publishCmd)
	articleCmd.AddCommand(deleteCmd)
}

func NewArticleCmd(c *config.Config) *cobra.Command {
	cfg = c
	return articleCmd
}
