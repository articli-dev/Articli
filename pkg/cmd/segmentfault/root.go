package segmentfault

import (
	"github.com/spf13/cobra"

	"github.com/k8scat/articli/internal/config"
	"github.com/k8scat/articli/pkg/cmd/segmentfault/article"
	"github.com/k8scat/articli/pkg/cmd/segmentfault/auth"
	"github.com/k8scat/articli/pkg/cmd/segmentfault/draft"
	"github.com/k8scat/articli/pkg/cmd/segmentfault/image"
	"github.com/k8scat/articli/pkg/cmd/segmentfault/tag"
)

var (
	cfgFile string
	cfg     *config.Config

	cmd = &cobra.Command{
		Use:   "sf",
		Short: "Manage content on segmentfault.com",
	}
)

func NewCmd(cf string, c *config.Config) *cobra.Command {
	cfgFile = cf
	cfg = c

	cmd.AddCommand(auth.NewCmd(cfgFile, cfg))
	cmd.AddCommand(article.NewArticleCmd(cfg))
	cmd.AddCommand(image.NewCmd(cfg))
	cmd.AddCommand(tag.NewCmd(cfg))
	cmd.AddCommand(draft.NewCmd(cfg))
	return cmd
}
