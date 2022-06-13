package auth

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/k8scat/articli/internal/config"
	sfsdk "github.com/k8scat/articli/pkg/platform/segmentfault"
)

var (
	tokenStdin bool

	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Authenticate with segmentfault.com",
		RunE: func(cmd *cobra.Command, args []string) error {
			bo := color.New(color.Bold)
			wo := color.New(color.FgWhite)

			var token string
			if tokenStdin {
				b, err := ioutil.ReadAll(os.Stdin)
				if err != nil {
					return errors.Trace(err)
				}
				token = strings.TrimSpace(string(b))
			} else {
				s := bufio.NewScanner(os.Stdin)
				if client != nil {
					for {
						bo.Printf("? You're already logged in as '%s'. Do you want to re-login? ", client.User.Name)
						wo.Print("(y/N) ")

						if !s.Scan() {
							return nil
						}

						in := strings.TrimSpace(strings.ToLower(s.Text()))
						if in != "y" && in != "n" && in != "no" && in != "yes" && in != "" {
							color.Red(`X Sorry, your reply was invalid: "%s" is not a valid answer, please try again.`, in)
							continue
						}
						if in == "n" || in == "no" || in == "" {
							return nil
						}
						break
					}
				}

				for {
					bo.Print("? Paste token: ")
					if !s.Scan() {
						return nil
					}

					token = strings.TrimSpace(s.Text())
					if token != "" {
						break
					}
					color.Red("X Sorry, your reply was invalid: Value is required")
				}
			}

			client, err := sfsdk.NewClient(token)
			if err != nil {
				log.Fatalf("error validating token: %v", err)
			}

			gr := color.New(color.FgGreen)
			gr.Print("✓ ")
			fmt.Print("Logged in as ")
			bo.Printf("%s\n", client.User.Name)

			cfg.Platforms.SegmentFault.Token = token
			err = config.SaveConfig(cfgFile, cfg)
			return errors.Trace(err)
		},
	}
)

func init() {
	loginCmd.Flags().BoolVar(&tokenStdin, "with-token", false, "Read token from standard input")
}
