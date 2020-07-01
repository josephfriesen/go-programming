package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	allWords := []string{"medic", "bunk", "snow", "africa", "gymnast", "mirror", "beer", "film", "laser", "press", "wake", "anthem", "tokyo", "soup", "chick", "smuggler", "saw", "tip", "spurs", "troll", "circle", "beijing", "brother", "bay", "egypt"}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(allWords), func(i, j int) { allWords[i], allWords[j] = allWords[j], allWords[i] })

	redWords := allWords[0:8]
	fmt.Println("Red Words")
	for _, v := range redWords {
		fmt.Println("\t", v)
	}

	blueWords := allWords[8:17]
	fmt.Println("Blue Words")
	for _, v := range blueWords {
		fmt.Println("\t", v)
	}

	BOMB := allWords[18]
	fmt.Println("BOMB:\t", BOMB)
}
