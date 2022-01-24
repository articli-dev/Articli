package draft

import (
	"fmt"
	juejinsdk "github.com/k8scat/articli/pkg/platform/juejin"
	"github.com/spf13/cobra"
	"os"
)

var (
	client *juejinsdk.Client

	draftCmd = &cobra.Command{
		Use:   "draft",
		Short: "Manage drafts",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if client == nil {
				fmt.Println("please login first")
				os.Exit(1)
			}
		},
	}
)

func NewDraftCmd(c *juejinsdk.Client) *cobra.Command {
	client = c
	draftCmd.AddCommand(newListCmd())
	draftCmd.AddCommand(editCmd)
	draftCmd.AddCommand(createCmd)
	draftCmd.AddCommand(deleteCmd)
	return draftCmd
}
