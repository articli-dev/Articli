package draft

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/k8scat/articli/pkg/table"
)

var (
	deleteCmd = &cobra.Command{
		Use:   "delete <draftIDs...>",
		Short: "Delete drafts",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			failedList := make([][]string, 0)
			for _, id := range args {
				idInt, err := strconv.ParseInt(id, 10, 64)
				if err != nil {
					failedList = append(failedList, []string{id, err.Error()})
					continue
				}
				if err := client.DeleteDraft(idInt); err != nil {
					failedList = append(failedList, []string{id, err.Error()})
				}
			}
			if len(failedList) > 0 {
				header := []string{"ID", "Error"}
				table.Print(header, failedList)
			}
			return nil
		},
	}
)
