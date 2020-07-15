package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// find represents the find command
var find = &cobra.Command{
	Use:   "find",
	Short: "find the occurence of any word or phrase throughout the entire project",
	Long: `given a word or phrase find where in the project is located what you're 
lookig for. You'll get a list of files in which were found matches.

For example:

term find --path /home/user/project `,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Running find command :D")

		path, _ := cmd.Flags().GetString("path")

		fmt.Println("path: ", path)

		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			fmt.Println(path, info.Size())
			return nil
		})
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(find)
	find.Flags().String("path", "", "--path [directory]")
	find.Flags().BoolP("case", "c", false, "--case")
}
