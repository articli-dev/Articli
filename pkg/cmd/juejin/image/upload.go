package image

import (
	"fmt"

	"github.com/juju/errors"
	"github.com/spf13/cobra"

	juejinsdk "github.com/k8scat/articli/pkg/platform/juejin"
)

var (
	region string

	uploadImageCmd = &cobra.Command{
		Use:   "upload <imagePath>",
		Short: "UploadImage image to juejin.cn",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			imagePath := args[0]
			imageURL, err := client.UploadImage(region, imagePath)
			if err != nil {
				return errors.Errorf("upload image failed: %s", errors.Trace(err))
			}
			fmt.Println(imageURL)
			return nil
		},
	}
)

func init() {
	uploadImageCmd.Flags().StringVarP(&region, "region", "r", juejinsdk.RegionCNNorth, "region")
}
