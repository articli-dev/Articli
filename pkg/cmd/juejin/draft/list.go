package draft

import (
	"fmt"
	"os"

	juejinsdk "github.com/k8scat/articli/pkg/platform/juejin"
	"github.com/k8scat/articli/pkg/table"
	"github.com/spf13/cobra"
)

var (
	keyword string
	limit   int

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List drafts",
		RunE: func(cmd *cobra.Command, args []string) error {
			if limit <= 0 {
				fmt.Println("limit must be greater than 0")
				os.Exit(1)
				return nil
			}

			result := make([]*juejinsdk.Draft, 0, limit)
			page := 1
			for {
				drafts, count, err := client.ListDrafts(keyword, page, juejinsdk.MaxPageSize)
				if err != nil {
					return err
				}
				result = append(result, drafts...)
				if len(result) >= limit || len(result) >= count || len(drafts) < juejinsdk.MaxPageSize {
					break
				}
				page++
			}
			if len(result) > limit {
				result = result[:limit]
			}

			header := []string{"ID", "标题", "创建时间"}
			data := make([][]string, 0, len(result))
			for _, draft := range result {
				if draft.Title == "" {
					draft.Title = "无标题"
				}
				data = append(data, []string{
					draft.ID,
					draft.Title,
					juejinsdk.FormatTime(draft.CreateTime, "2006-01-02 15:04"),
				})
			}
			table.Print(header, data)
			return nil
		},
	}
)

func init() {
	listCmd.Flags().StringVarP(&keyword, "keyword", "k", "", "Keyword")
	listCmd.Flags().IntVarP(&limit, "limit", "l", 10, "Maximum number of drafts to list")
}
