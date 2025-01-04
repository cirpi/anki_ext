package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"log"
)

type Decks struct {
	Decks []string `json:"result"`
}

const (
	SCHEME        = "http"
	DOMAIN        = "localhost"
	PORT   string = "8765"
)

var client = &http.Client{}

func GetURL() string {
	return SCHEME + "://" + DOMAIN + ":" + PORT
}

func createRequest(payload io.Reader) *http.Request {
	req, er := http.NewRequest(http.MethodPost, GetURL(), payload)
	if er != nil {
		log.Printf("DECKS REQUEST: %s", er.Error())
	}
	return req
}

func createDecksRequest() *http.Request {
	payload := strings.NewReader("{\"action\":\"deckNames\", \"version\":6}")
	return createRequest(payload)
}

func GetDecks() []string {
	resp, er := client.Do(createDecksRequest())
	if er != nil {
		log.Printf("DECKS: %s", er.Error())
		return []string{}
	}
	decks := Decks{}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&decks)
	return decks.Decks
}

func createAddNotesRequest(payload io.Reader) *http.Request {
	// payload := strings.NewReader(`
	//    {
	//        "action": "addNote",
	//        "version": 6,
	//        "params" : {
	//            "note" : {
	//                "deckName": "dummy",
	//                "modelName": "Basic",
	//                "fields" : {
	//                    "Front": "Name?",
	//                    "Back": "Sirpi"
	//                }
	//            }
	//        }
	//    }
	//    `)
	return createRequest(payload)
}

func CreateCard(payload string) {
	reader := strings.NewReader(payload)
	resp, er := client.Do(createAddNotesRequest(reader))
	if er != nil {
		fmt.Println("er", er.Error())
	}
	s := make([]byte, 100)
	defer resp.Body.Close()
	n, _ := resp.Body.Read(s)
	fmt.Println("read", string(s[:n]))
}
