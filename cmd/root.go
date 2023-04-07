/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chatctl",
	Short: "chatctl is a CLI for chat control -- server and client",
	Long:  `chatctl is a CLI for chat control -- server and client. This tool helps to send message across ips via command line interfaces - terminal, command prompt`,
}

func Execute() {
	var banner string = `
		
		
	     ___           ___           ___                         ___                                 
	    /\__\         /\  \         /\  \                       /\__\                                
	   /:/  /         \:\  \       /::\  \         ___         /:/  /          ___                   
	  /:/  /           \:\  \     /:/\:\  \       /\__\       /:/  /          /\__\                  
	 /:/  /  ___   ___ /::\  \   /:/ /::\  \     /:/  /      /:/  /  ___     /:/  /      ___     ___ 
	/:/__/  /\__\ /\  /:/\:\__\ /:/_/:/\:\__\   /:/__/      /:/__/  /\__\   /:/__/      /\  \   /\__\
	\:\  \ /:/  / \:\/:/  \/__/ \:\/:/  \/__/  /::\  \      \:\  \ /:/  /  /::\  \      \:\  \ /:/  /
	 \:\  /:/  /   \::/__/       \::/__/      /:/\:\  \      \:\  /:/  /  /:/\:\  \      \:\  /:/  / 
	  \:\/:/  /     \:\  \        \:\  \      \/__\:\  \      \:\/:/  /   \/__\:\  \      \:\/:/  /  
	   \::/  /       \:\__\        \:\__\          \:\__\      \::/  /         \:\__\      \::/  /   
	    \/__/         \/__/         \/__/           \/__/       \/__/           \/__/       \/__/    
		
		`
	fmt.Println(banner)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
