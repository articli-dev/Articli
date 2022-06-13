package note

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/juju/errors"
	"github.com/spf13/cobra"

	sfsdk "github.com/k8scat/articli/pkg/platform/segmentfault"
	"github.com/k8scat/articli/pkg/table"
	"github.com/k8scat/articli/pkg/utils"
)

var (
	limit int

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List drafts",
		RunE: func(cmd *cobra.Command, args []string) error {
			if limit <= 0 {
				fmt.Println("limit must be greater than 0")
				os.Exit(1)
				return nil
			}

			result := make([]*sfsdk.DraftRow, 0)
			page := 1
			for {
				opts := &sfsdk.ListNotesOptions{
					Page: page,
					Size: sfsdk.PageSizeMaxDraft,
				}
				resp, err := client.ListDrafts(opts)
				if err != nil {
					return errors.Trace(err)
				}
				result = append(result, resp.Rows...)
				if len(result) >= limit || page >= resp.TotalPage || len(resp.Rows) < sfsdk.PageSizeMaxDraft {
					break
				}
				page++
			}
			if len(result) > limit {
				result = result[:limit]
			}

			header := []string{"ID", "标题", "类型", "更新时间", "链接"}
			data := make([][]string, 0, len(result))
			for _, draft := range result {
				if draft.Title == "" {
					draft.Title = "无标题"
				}
				data = append(data, []string{
					strconv.FormatInt(draft.ID, 10),
					draft.Title,
					draft.TypeName,
					time.Unix(draft.Modified, 0).Format(utils.DefaultTimeLayout),
					draft.GetURL(),
				})
			}
			table.Print(header, data)
			return nil
		},
	}
)

func init() {
	listCmd.Flags().IntVarP(&limit, "limit", "l", 10, "Maximum number of drafts to list")
}
