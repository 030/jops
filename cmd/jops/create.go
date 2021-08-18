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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.PersistentFlags().StringVarP(&summary, "summary", "s", "", "The summary of the Jira issue")
	if err := createCmd.MarkPersistentFlagRequired("summary"); err != nil {
		log.Fatal(err)
	}

	createCmd.PersistentFlags().StringVarP(&description, "description", "d", "", "The description of the Jira issue")
	if err := createCmd.MarkPersistentFlagRequired("description"); err != nil {
		log.Fatal(err)
	}
}
