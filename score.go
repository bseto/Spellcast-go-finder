package main

import (
	"sort"
	"strings"
)

var PointMap = map[string]int{
	"A": 1, "B": 4, "C": 5, "D": 3, "E": 1,
	"F": 5, "G": 3, "H": 4, "I": 1, "J": 7,
	"K": 3, "L": 3, "M": 4, "N": 2, "O": 1,
	"P": 4, "Q": 8, "R": 2, "S": 2, "T": 2,
	"U": 4, "V": 5, "W": 5, "X": 7, "Y": 4,
	"Z": 8,
}

type Score struct {
	Word  string
	Score int
}

func CalculateAndSortByScore(words []string) []Score {
	scores := make([]Score, len(words))
	for i, word := range words {
		var currentScore int
		for _, letter := range word {
			currentScore += PointMap[strings.ToUpper(string(letter))]
		}
		scores[i] = Score{
			Word:  word,
			Score: currentScore,
		}
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score < scores[j].Score
	})
	return scores
}