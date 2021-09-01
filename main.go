package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/MagicTheGathering/mtg-sdk-go"
	"github.com/briandowns/spinner"
)

var Red = "\033[31m"
var Reset = "\033[0m"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetCard(cardName string) (mtg.Card, error) {
	query := "https://api.magicthegathering.io/v1/cards?name=" + cardName
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
	return errorCard, errors.New("{{Carta não encontrada}}: ")
}

func GetCards(cardsInput []string, cards *[]mtg.Card, wg *sync.WaitGroup) {
	for i := 0; i < len(cardsInput); i++ {
		card, err := GetCard(cardsInput[i])
		if err != nil {
			fmt.Println()
			fmt.Println(Red + err.Error() + ": " + cardsInput[i] + Reset)
			fmt.Println()
			continue
		}
		*cards = append(*cards, card)
	}
	wg.Done()
}

func GetChosenFormat() string {
	mtgFormats, err := mtg.GetFormats()
	Check(err)

	fmt.Println("Qual formato você quer jogar?")
	for i := 0; i < len(mtgFormats); i++ {
		fmt.Println(strconv.Itoa(i) + " - " + mtgFormats[i])
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	chosenFormat, err := strconv.Atoi(text)
	Check(err)

	return mtgFormats[chosenFormat]
}

func main() {
	s := spinner.New(spinner.CharSets[26], 1000*time.Millisecond)

	chosenFormat := GetChosenFormat()
	s.Prefix = "Checando legalidade das cartas para " + chosenFormat
	s.Start()

	dat, err := os.ReadFile("cards.csv")
	Check(err)

	cardsInput := strings.Split(string(dat), "\n")

	cards := make([]mtg.Card, 0)

	var wg sync.WaitGroup
	wg.Add(1)
	go GetCards(cardsInput, &cards, &wg)
	wg.Wait()
	s.Stop()

	validCards := make([]string, 0)
	invalidCards := make([]string, 0)
	for i := 0; i < len(cards); i++ {
		isLegal := false
		for j := 0; j < len(cards[i].Legalities); j++ {
			if cards[i].Legalities[j].Format == chosenFormat {
				isLegal = true
				break
			}
		}

		if isLegal {
			validCards = append(validCards, cards[i].Name)
		} else {
			invalidCards = append(invalidCards, cards[i].Name)
		}
	}

	fmt.Println()

	if len(validCards) > 0 && len(invalidCards) == 0 {
		fmt.Println("Seu deck é TOTALMENTE jogável no " + chosenFormat + "!!! :D")
	} else if len(invalidCards) > 0 && len(validCards) == 0 {
		fmt.Println("NENHUMA carta do seu deck é jogável no " + chosenFormat + " :(")
	} else {
		fmt.Println("[[Cartas LEGAIS no " + chosenFormat + "]]:")
		for i := 0; i < len(validCards); i++ {
			fmt.Println(validCards[i])
		}
		fmt.Println()
		fmt.Println("------------------------------------------")
		fmt.Println()
		fmt.Println("[[Cartas ILEGAIS no " + chosenFormat + "]]:")
		for i := 0; i < len(invalidCards); i++ {
			fmt.Println(invalidCards[i])
		}
	}
}
