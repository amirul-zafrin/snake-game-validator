package utilities

import (
	"errors"
	"math/rand"
	"strings"

	"github.com/amirul-zafrin/snake-validator/models"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandomFruitPositon(width, height int) models.Fruit {
	fruit := models.Fruit{
		X: rand.Intn(width),
		Y: rand.Intn(height),
	}
	return fruit
}

func GenerateNewGame(width, height int) (models.State, error) {
	if width <= 0 || height <= 0 {
		return models.State{}, errors.New("invalid width / height")
	}

	InitState := models.State{
		GameID: randomGameId(10),
		Width:  defaultWidthHeight(width),
		Height: defaultWidthHeight(height),
		Score:  0,
		Snake: models.Snake{
			X:    0,
			Y:    0,
			VelX: 1,
			VelY: 0,
		},
	}
	InitState.Fruit = RandomFruitPositon(InitState.Width, InitState.Height)
	return InitState, nil
}

func randomGameId(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func defaultWidthHeight(provided int) int {
	if provided > 0 {
		return provided
	}

	return randomMinMax(10, 1000)
}

func randomMinMax(min, max int) int {
	return min + rand.Intn(max-min+1)
}
