package helper

import (
	"math/rand"
	"time"
	"strings"
	"github.com/blockchain/baccarat/x/baccarat/types"
	"strconv"
	"math"
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

func Winner(hand string) (types.BetSide, int) {
	pb := strings.Split(hand, ";")
	player := strings.Split(pb[0], ",")
	banker := strings.Split(pb[1], ",")
	p1, _ := strconv.Atoi(player[0][:len(player[0]) - 1])
	p2, _ := strconv.Atoi(player[1][:len(player[1]) - 1])
	b1, _ := strconv.Atoi(banker[0][:len(banker[0]) - 1])
	b2, _ := strconv.Atoi(banker[1][:len(banker[1]) - 1])
	distPlayer := math.Abs(9 - float64(p1 + p2))
	distBanker :=  math.Abs(9 - float64(b1 + b2))
	if distPlayer == distBanker {
		return types.Tie, 5
	} else if distPlayer < distBanker {
		return types.Player, 2
	} else {
		return types.Banker, 2
	}
}


