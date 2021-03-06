package auth

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "View authentication status",
		Run: func(cmd *cobra.Command, args []string) {
			bo := color.New(color.Bold)
			gr := color.New(color.FgGreen)

			if client == nil {
				fmt.Print("You are not logged into gitlab. Run ")
				bo.Print("acli gitlab auth login")
				fmt.Println(" to authenticate.")
				os.Exit(1)
			} else {
				gr.Print("✓ ")
				gr.Printf("Logged in to %s as %s (%s)\n", cfg.Platforms.Gitlab.BaseURL, client.User.Name, cfgFile)
			}
		},
	}
)
