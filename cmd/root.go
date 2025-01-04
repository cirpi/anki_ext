package cmd

import (
	"anki/client"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Anki cli",
	Short: "An command-line extension for anki",
	Run: func(c *cobra.Command, args []string) {
		res := client.GetDecks()
		fmt.Println(res)
	},
}

func Execute() {
	er := rootCmd.Execute()
	if er != nil {
		log.Fatalf("INIT_ERROR: %s", er.Error())
	}
}
