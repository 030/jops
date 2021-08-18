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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("done called")
		j := done.Jira{User: user, Pass: pass, FQDN: fqdn, Project: project, TicketNumber: ticketNumber, Comment: comment}
		if err := j.Done(); err != nil {
			log.Fatal(err)
		}
	},
}

var comment, ticketNumber string

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	doneCmd.PersistentFlags().StringVarP(&ticketNumber, "ticketNumber", "t", "", "The ticketNumber that should be moved to done")
	if err := doneCmd.MarkPersistentFlagRequired("ticketNumber"); err != nil {
		log.Fatal(err)
	}

	doneCmd.PersistentFlags().StringVarP(&comment, "comment", "c", "", "The comment that should be added to the ticket that will be closed")
	if err := doneCmd.MarkPersistentFlagRequired("comment"); err != nil {
		log.Fatal(err)
	}
}
