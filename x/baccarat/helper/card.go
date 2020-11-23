package helper

import (
	"math/rand"
	"time"
	"strings"
)

func DrawCard(deck string) (string, string) {
	deckArray := strings.Split(deck, ",")
	hand := deckArray[0] + "," + deckArray[1] + ";" + deckArray[2] + "," + deckArray[3]
	return hand, strings.Join(deckArray[4:], ",")
}

func GenerateDeck() string {
	suites := []string{"S", "H", "D", "C"}
	values := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	deck :=  make([]string, 0, 52)
	for _, suite := range suites {
		for _, value := range values {
			deck = append(deck, value + suite)
		}
	}
	return strings.Join(Shuffle(deck), ",")
}

func Shuffle(deck []string) []string {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := 0; i < len(deck); i++ {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}


