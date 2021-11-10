package main

import (
	"fmt"

	"github.com/030/jops/pkg/jira/v2/changelog"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// changelogCmd represents the changelog command
var changelogCmd = &cobra.Command{
	Use:   "changelog",
	Short: "changelog",
	Long:  `changelog`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("changelog called")
		j := changelog.Jira{User: user, Pass: pass, FQDN: fqdn, TicketNumber: ticketNumber}
		s, err := j.Get()
		if err != nil {
			log.Fatal(err)
		}
		log.Info(s)
	},
}

func init() {
	rootCmd.AddCommand(changelogCmd)

	changelogCmd.PersistentFlags().StringVarP(&ticketNumber, "ticketNumber", "t", "", "The ticketNumber that should be moved to done")
	if err := changelogCmd.MarkPersistentFlagRequired("ticketNumber"); err != nil {
		log.Fatal(err)
	}
}
