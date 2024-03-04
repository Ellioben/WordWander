package handle

import (
	"math/rand"
	"time"
)

const (
	randomPatch = 20
)

func RandomWord(wordList []string) []string {
	randomWords := make([]string, randomPatch)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < randomPatch; i++ {
		randomIndex := rand.Intn(len(wordList))
		randomWords[i] = wordList[randomIndex]
	}

	var endWordList []string
	for _, word := range randomWords {
		endWordList = append(endWordList, word)
	}
	return endWordList
}
