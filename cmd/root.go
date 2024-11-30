/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/felipeivanaga/stresstester/cmd/httptester"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stresstester",
	Short: "Call a website",
	Long:  `Call a website for Stress the server`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		if url == "" {
			fmt.Println("Please inform the url")
			os.Exit(1)
		}

		requests, _ := cmd.Flags().GetInt("requests")
		if requests <= 0 {
			fmt.Println("Requests need to be higher than 0")
			os.Exit(1)
		}

		concurrency, _ := cmd.Flags().GetInt("concurrency")
		if concurrency <= 0 {
			fmt.Println("Concurrency need to be higher than 0")
			os.Exit(1)
		}

		fmt.Printf("Calling %s %d times, %d simultaneously \n", url, requests, concurrency)
		httptester.Execute(url, requests, concurrency)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "Website URL")
	rootCmd.Flags().IntP("requests", "r", 10, "Max numbers of requets")
	rootCmd.Flags().IntP("concurrency", "c", 4, "Simultaneously calls")
}
