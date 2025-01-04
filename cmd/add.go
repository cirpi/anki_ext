package cmd

import (
	"anki/client"
	"anki/cmd/handlers"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

const CARDMAX = 1

type Card int

const (
	_ Card = iota
	Basic
)

var cardType Card = Basic

func init() {
	rootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Use:   "add",
	Short: "Add cards to decks",
	Run: func(c *cobra.Command, args []string) {
		// cardTypePrompt()
		card := handlers.BasicCard{}
		card.Front = "what's ur name?"
		card.Back = "sirpi yugendar"
		payload := handlers.CreateCard("addNote", "dummy", "Basic", card)
		js, er := json.Marshal(payload)
		if er != nil {
			fmt.Println("ErrOR json", er.Error())
			return
		}
		fmt.Println(string(js))
		client.CreateCard(string(js))
		//client.CreateCard()
	},
}

func cardTypePrompt() {
	fmt.Printf("%s\nCard type: [1-%d]", cardTypes(), CARDMAX)
	var input string
	fmt.Scanln(&input)
	if input != "" {
		num, er := strconv.Atoi(input)
		if er != nil {
			log.Printf("PARSE_USER_INPUT: %s", er.Error())
		} else if num > 0 && num <= CARDMAX {
			cardType = Card(num)
		}
	}
	fmt.Println("Selected type: ", cardType)
}

func cardTypes() string {
	return fmt.Sprintf(`1. Basic`)
}
