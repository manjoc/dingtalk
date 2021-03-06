package dingtalk

import (
	"github.com/manjoc/dingtalk"
	"github.com/manjoc/gutils/log"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "send link message with DingTalk robot",
	Long:  `send link message with DingTalk robot`,
	Args:  cobra.MinimumNArgs(0),
	Run:   runLinkCmd,
}

func runLinkCmd(_ *cobra.Command, args []string) {
	if len(linkVars.title) < 1 {
		log.L(log.Red, "title can not be empty")
		return
	}

	if len(linkVars.text) < 1 {
		log.L(log.Red, "text can not be empty")
		return
	}

	if len(linkVars.messageURL) < 1 {
		log.L(log.Red, "messageURL can not be empty")
		return
	}

	client, err := newClient()
	if err != nil {
		log.L(log.Red, err.Error())
		return
	}

	msg := dingtalk.NewLinkMessage().
		SetLink(linkVars.title, linkVars.text, linkVars.picURL, linkVars.messageURL)
	if _, err := client.Send(msg); err != nil {
		log.L(log.Red, err.Error())
	}
}

// LinkVars struct
type LinkVars struct {
	title      string
	text       string
	picURL     string
	messageURL string
}

var linkVars LinkVars

func init() {
	rootCmd.AddCommand(linkCmd)
	linkCmd.Flags().StringVarP(&linkVars.title, "title", "i", "", "title")
	linkCmd.Flags().StringVarP(&linkVars.text, "text", "e", "", "text")
	linkCmd.Flags().StringVarP(&linkVars.picURL, "picURL", "p", "", "picURL")
	linkCmd.Flags().StringVarP(&linkVars.messageURL, "messageURL", "u", "", "messageURL")
}
