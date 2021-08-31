package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/MagicTheGathering/mtg-sdk-go"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetCard(cardName string) (mtg.Card, error) {
	query := "https://api.magicthegathering.io/v1/cards?name=" + cardName
	fmt.Println(query)
	resp, err := http.Get(query)
	Check(err)

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)

	defer resp.Body.Close() // for garbage collection

	bodyString := string(body)

	type CardsResponse struct {
		Cards []mtg.Card `json:"cards"`
	}

	cardsResponse := new(CardsResponse)
	json.Unmarshal([]byte(bodyString), &cardsResponse)

	if len(cardsResponse.Cards) > 0 {
		return cardsResponse.Cards[0], nil
	}

	errorCard := mtg.Card{}
	return errorCard, errors.New("card not found")
}

func main() {

	dat, err := os.ReadFile("cards.csv")
	Check(err)

	cardsInput := strings.Split(string(dat), "\n")

	cards := make([]mtg.Card, 3)

	for i := 0; i < len(cardsInput); i++ {
		card, err := GetCard(cardsInput[i])
		if err != nil {
			fmt.Println(err.Error() + ": " + cardsInput[i])
			continue
		}
		cards = append(cards, card)
	}

	// fmt.Println(cards)

}
