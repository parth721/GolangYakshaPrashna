package utils

import (
	"math/rand"
	//"time"
)

var Indices [2]int
var RandomNumber int
var yaksha1 string
var yaksha2 string

func GetRandomQuestions(n int) (string, string) {
	// Initialize the random number generator
	//var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Loop to get n random indices
	for i := 0; i < n; i++ {
		RandomNumber := rand.Intn(4-1) + 1
		Indices[i] = (RandomNumber - 1)
	}
	yaksha1 = Questions[Indices[0]]
	yaksha2 = Questions[Indices[1]]

	return yaksha1, yaksha2
}

func GetSolution(m int) string {
	var index int = m - 1
	return Solutions[Indices[index]]
}
