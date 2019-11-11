package discordbot

import (
	"log"
	"math/rand"
	"strconv"
)

func diceRoll(n []string) string {
	dieSides, err := strconv.Atoi(n[1]) // Converts die sides to int from ASCII
	if err != nil {
		log.Println(err)
		return "String conversion of die sides failed!"
	}

	return strconv.Itoa(rand.Intn(dieSides-1) + 1) // Rolls die and returns result as string
}
