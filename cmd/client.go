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
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "join as client",
	Long:  `This command assigns you the role of a client in the connection`,
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
		serverIP, _ := cmd.Flags().GetString("server-ip")
		fmt.Println("Launching client...")

		conn, err := net.Dial("tcp", serverIP) // Replace with server's IP address
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		for {
			fmt.Print("Send message: ")
			messageReader := bufio.NewReader(os.Stdin)
			message, err := messageReader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			message = strings.TrimSpace(message)

			_, err = conn.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println(err)
				return
			}

			reply, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			reply = strings.TrimSpace(reply)

			fmt.Println("Received message: " + reply)
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.Flags().StringP("server-ip", "i", "", "specify the ip address of the server")
}
