package main

import (
	"fmt"

	"github.com/030/jops/pkg/jira/v2/create"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Jira issue",
	Long:  `Create a Jira issue`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")

		j := create.Jira{User: user, Pass: pass, FQDN: fqdn, Priority: priority, Project: project, Summary: summary, Description: description, Labels: labels}
		_, err := j.Create()
		if err != nil {
			log.Fatal(err)
		}
	},
}

var description, priority, summary string
var labels []string

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringVarP(&description, "description", "d", "", "The description of the Jira issue")
	if err := createCmd.MarkPersistentFlagRequired("description"); err != nil {
		log.Fatal(err)
	}

	createCmd.PersistentFlags().StringSliceVarP(&labels, "labels", "l", []string{""}, "The labels of the Jira issue")

	createCmd.PersistentFlags().StringVarP(&priority, "priority", "p", "", "The priority of the Jira issue")
	if err := createCmd.MarkPersistentFlagRequired("priority"); err != nil {
		log.Fatal(err)
	}

	createCmd.PersistentFlags().StringVarP(&summary, "summary", "s", "", "The summary of the Jira issue")
	if err := createCmd.MarkPersistentFlagRequired("summary"); err != nil {
		log.Fatal(err)
	}
}
