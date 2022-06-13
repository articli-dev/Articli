package article

import (
	"strconv"
	"time"

	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	juejinsdk "github.com/k8scat/articli/pkg/platform/juejin"
	sfsdk "github.com/k8scat/articli/pkg/platform/segmentfault"
	"github.com/k8scat/articli/pkg/table"
)

var (
	limit int

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List articles",
		RunE: func(cmd *cobra.Command, args []string) error {
			if limit <= 0 {
				log.Fatal("limit must be greater than 0")
			}

			result := make([]*sfsdk.ArticleRow, 0, limit)
			page := 1
			opts := &sfsdk.ListArticlesOptions{
				Page: page,
				Size: sfsdk.PageSizeMax,
				Sort: sfsdk.ArticleSortNewest,
			}
			for {
				resp, err := client.ListArticles(opts)
				if err != nil {
					return errors.Trace(err)
				}
				result = append(result, resp.Rows...)
				if len(result) >= limit || page >= resp.TotalPage || len(resp.Rows) < juejinsdk.MaxPageSize {
					break
				}
				opts.Page++
			}
			if len(result) > limit {
				result = result[:limit]
			}

			header := []string{"ID", "标题", "URL", "创建时间"}
			data := make([][]string, 0, len(result))
			for _, a := range result {
				data = append(data, []string{
					strconv.FormatInt(a.ID, 10),
					a.Title,
					a.GetURL(),
					time.Unix(a.Created, 0).Format("2006-01-02 15:04"),
				})
			}
			table.Print(header, data)
			return nil
		},
	}
)

func init() {
	listCmd.Flags().IntVarP(&limit, "limit", "l", 10, "Maximum number of articles to list")
}
