package article

import (
	"fmt"

	"github.com/juju/errors"
	"github.com/spf13/cobra"
)

var (
	syncToOrg bool

	publishCmd = &cobra.Command{
		Use:   "publish <draftID>",
		Short: "Publish a draft into an article",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			draftID := args[0]
			id, err := client.PublishArticle(draftID, syncToOrg)
			if err != nil {
				return errors.Trace(err)
			}
			fmt.Println(id)
			return nil
		},
	}
)

func init() {
	publishCmd.Flags().BoolVarP(&syncToOrg, "sync-org", "s", false, "Sync to org")
}
