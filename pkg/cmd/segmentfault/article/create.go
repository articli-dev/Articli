package article

import (
	"fmt"

	"github.com/juju/errors"
	"github.com/spf13/cobra"

	"github.com/k8scat/articli/pkg/markdown"
	juejinsdk "github.com/k8scat/articli/pkg/platform/juejin"
)

var (
	createCmd = &cobra.Command{
		Use:   "create <markdownFile>",
		Short: "Create or update an article from a markdown file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			markdownFile := args[0]
			mark, err := markdown.Parse(markdownFile)
			if err != nil {
				return errors.Trace(err)
			}

			params, err := client.ParseMark(mark)
			if err != nil {
				return errors.Trace(err)
			}
			isCreate := false
			if params.ArticleID == "" {
				isCreate = true
			}

			if err := client.SaveArticle(params); err != nil {
				return errors.Trace(err)
			}

			if err := juejinsdk.WriteBack(juejinsdk.SaveTypeArticle, mark, params, isCreate); err != nil {
				return errors.Trace(err)
			}
			fmt.Println(juejinsdk.BuildArticleURL(params.ArticleID))
			return nil
		},
	}
)

func init() {
	createCmd.Flags().BoolVarP(&syncToOrg, "sync", "s", false, "Sync to org")
}
