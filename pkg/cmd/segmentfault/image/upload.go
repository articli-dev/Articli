package image

import (
	"fmt"

	"github.com/juju/errors"
	"github.com/spf13/cobra"
)

var (
	uploadImageCmd = &cobra.Command{
		Use:   "upload <imagePath>",
		Short: "Upload image on segmentfault.com",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			imagePath := args[0]
			imageURL, err := client.UploadImage(imagePath)
			if err != nil {
				return errors.Errorf("upload image failed: %s", errors.Trace(err))
			}
			fmt.Println(imageURL)
			return nil
		},
	}
)
