package tag

import (
	"strconv"
	"strings"

	"github.com/juju/errors"

	"github.com/k8scat/articli/pkg/table"

	"github.com/spf13/cobra"
)

var (
	keyword string

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List tags",
		RunE: func(cmd *cobra.Command, args []string) error {
			keyword = strings.TrimSpace(keyword)
			resp, err := client.SearchTags(keyword)
			if err != nil {
				return errors.Errorf("list tags failed: %+v", errors.Trace(err))
			}

			header := []string{"ID", "名称", "链接"}
			data := make([][]string, 0, resp.Size)
			for _, t := range resp.Rows {
				data = append(data, []string{
					strconv.FormatInt(t.ID, 10),
					t.Name,
					t.GetURL(),
				})
			}
			table.Print(header, data)
			return nil
		},
	}
)

func init() {
	listCmd.Flags().StringVarP(&keyword, "keyword", "k", "", "Filter keyword")
}
