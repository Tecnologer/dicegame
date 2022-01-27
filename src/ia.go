package dice

import (
	"fmt"
	"time"

	"github.com/tecnologer/dicegame/src/models"
)

func wanstAIEndTurn(computer *models.Player) bool {
	if computer.Shots < 2 {
		return false
	}

	fmt.Println("la computadora esta decidiendo...")
	time.Sleep(5 * time.Second)
	val := calculateProbability()
	if computer.Shots == 2 && val > 70 {
		return false
	}

	return true
}

func calculateProbability() int {
	return 70
}