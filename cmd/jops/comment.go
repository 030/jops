package main

import (
	"fmt"

	"github.com/030/jops/pkg/jira/v2/comment"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// commentCmd represents the comment command
var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Get all comments or add a comment to an existing ticket",
	Long:  `Get all comments or add a comment to an existing ticket`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("comment called")
		if !add && !all {
			log.Fatal("Either 'add' or 'all' subcommand is required")
		}

		if add {
			if msg == "" {
				log.Fatal("message subcommand is mandatory")
			}

			j := comment.Jira{User: user, Pass: pass, FQDN: fqdn, TicketNumber: ticketNumber}
			if err := j.Add(msg); err != nil {
				log.Fatal(err)
			}
		}

		if all {
			j := comment.Jira{User: user, Pass: pass, FQDN: fqdn, TicketNumber: ticketNumber}
			if err := j.GetAll(); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var add, all bool
var msg string

func init() {
	rootCmd.AddCommand(commentCmd)

	commentCmd.Flags().BoolVarP(&add, "add", "c", false, "Add a comment to a ticket")
	commentCmd.Flags().BoolVarP(&all, "all", "a", false, "Get all ticket comments")

	commentCmd.PersistentFlags().StringVarP(&msg, "message", "m", "", "The comment to be added to a ticket")

	commentCmd.PersistentFlags().StringVarP(&ticketNumber, "ticketNumber", "t", "", "The ticketNumber that should be moved to done")
	if err := commentCmd.MarkPersistentFlagRequired("ticketNumber"); err != nil {
		log.Fatal(err)
	}
}
