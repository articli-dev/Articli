package draft

import (
	"fmt"
	"strconv"

	"github.com/juju/errors"
	"github.com/spf13/cobra"

	"github.com/k8scat/articli/pkg/markdown"
	sfsdk "github.com/k8scat/articli/pkg/platform/segmentfault"
)

var (
	createCmd = &cobra.Command{
		Use:   "create <markdownFile>",
		Short: "Create or update a draft from a markdown file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			markdownFile := args[0]
			mark, err := markdown.Parse(markdownFile)
			if err != nil {
				return errors.Trace(err)
			}

			meta, err := sfsdk.ParseMark(mark)
			if err != nil {
				return errors.Trace(err)
			}
			isCreate := false
			if meta.DraftID == "" {
				isCreate = true
			}

			draft, err := meta.IntoDraft()
			if err != nil {
				return errors.Trace(err)
			}

			if isCreate {
				err = client.createDraft(draft)
			} else {
				err = client.UpdateDraft(draft)
			}
			if err != nil {
				return errors.Trace(err)
			}

			if isCreate {
				meta.DraftID = strconv.FormatInt(draft.ID, 10)
			}

			if err := sfsdk.WriteMarkdownMeta(sfsdk.SaveTypeDraft, mark, meta, isCreate); err != nil {
				return errors.Trace(err)
			}
			fmt.Println(draft.GetURL())
			return nil
		},
	}
)
