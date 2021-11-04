package main

import (
	"fmt"
	"log"

	"github.com/030/jops/internal/jira/v2/issue/create"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Jira issue",
	Long:  `Create a Jira issue`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		j := create.Jira{User: user, Pass: pass, FQDN: fqdn, Project: project, Summary: summary, Description: description}
		if err := j.Create(); err != nil {
			log.Fatal(err)
		}
	},
}

var description, summary string

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringVarP(&summary, "summary", "s", "", "The summary of the Jira issue")
	if err := createCmd.MarkPersistentFlagRequired("summary"); err != nil {
		log.Fatal(err)
	}

	createCmd.PersistentFlags().StringVarP(&description, "description", "d", "", "The description of the Jira issue")
	if err := createCmd.MarkPersistentFlagRequired("description"); err != nil {
		log.Fatal(err)
	}
}
