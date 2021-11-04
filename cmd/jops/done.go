package main

import (
	"fmt"
	"log"

	"github.com/030/jops/pkg/jira/v2/done"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Close a ticket",
	Long:  `Close a ticket`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("done called")
		j := done.Jira{User: user, Pass: pass, FQDN: fqdn, TicketNumber: ticketNumber, Comment: doneComment}
		if err := j.Done(); err != nil {
			log.Fatal(err)
		}
	},
}

var doneComment, ticketNumber string

func init() {
	rootCmd.AddCommand(doneCmd)

	doneCmd.PersistentFlags().StringVarP(&ticketNumber, "ticketNumber", "t", "", "The ticketNumber that should be moved to done")
	if err := doneCmd.MarkPersistentFlagRequired("ticketNumber"); err != nil {
		log.Fatal(err)
	}

	doneCmd.PersistentFlags().StringVarP(&doneComment, "comment", "c", "", "The comment that should be added to the ticket that will be closed")
	if err := doneCmd.MarkPersistentFlagRequired("comment"); err != nil {
		log.Fatal(err)
	}
}
