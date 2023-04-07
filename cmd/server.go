/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"bufio"
	"net"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theifedayo/chatctl/utils"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "join as server",
	Long:  `This command assigns you the role of a server in the connection`,
	Run: func(cmd *cobra.Command, args []string) {
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
		fmt.Println("Chat with anyone via terminal")
		fmt.Println("Launching server...")
		fmt.Println("Share " + utils.GetIPAddress() + ":8080 to anyone you want to join you")

		ln, err := net.Listen("tcp", ":8080")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer ln.Close()

		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			message = strings.TrimSpace(message)

			fmt.Println("Received message: " + message)

			fmt.Print("Send message: ")
			replyReader := bufio.NewReader(os.Stdin)
			reply, err := replyReader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			reply = strings.TrimSpace(reply)

			_, err = conn.Write([]byte(reply + "\n"))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
