/*
Copyright © 2024 NAME HERE <danubio.bwm@gmail.com>
*/
package cmd

import (
	"github.com/danubiobwm/goExpert/CLI/internal/database"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
// var createCmd = &cobra.Command{
// 	Use:   "create",
// 	Short: "A brief description of your command",
// 	Long:  `A longer description that spans multiple lines and likely contains examples`,
// 	RunE:  runCreate(GetCategoryDB(GetDb())),
// Run: func(cmd *cobra.Command, args []string) {
// 	db := GetDb()
// 	category := GetCategoryDB(db)

// 	name, _ := cmd.Flags().GetString("name")
// 	description, _ := cmd.Flags().GetString("description")
// 	category.Create(name, description)
// },
// }

func newCreateCmd(categoryDB database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		RunE:  runCreate(categoryDB),
	}
}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "Y", "Name of category")
	createCmd.Flags().StringP("description", "d", "Y", "Description of category")
	createCmd.MarkFlagsRequiredTogether("name", "description")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
