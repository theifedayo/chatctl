/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/theifedayo/chatctl/utils"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "get your encrypted ip address",
	Long:  `This command generates and encrypts your ip address`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.IPAndEncrypt()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

}
