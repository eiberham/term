package cmd

import (
	"fmt"
	// "strconv"

	"github.com/spf13/cobra"
)

// find represents the find command
var find = &cobra.Command{
	Use:   "find",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Running find command :D")
	},
}

func init() {
	rootCmd.AddCommand(find)
	// addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}
