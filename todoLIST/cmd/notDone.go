/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
)

// notDoneCmd represents the notDone command
var notDoneCmd = &cobra.Command{
	Use:   "notDone",
	Short: "This would mark a task as undone",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		i, _ := strconv.Atoi(args[0])
		TaskNotDone(i)
	},
}

func init() {
	rootCmd.AddCommand(notDoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notDoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notDoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
