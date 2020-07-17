package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "term",
	Short: "Term is a tool that lets you find occurences of words in files within your project",
	Long: `The advantages of this tool is that it was designed leveraging the goods of golang's
	concurrency model to perform at good speed`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("inside init")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	fmt.Println("Init config")
}
