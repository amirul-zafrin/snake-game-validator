package handlers

import (
	"github.com/amirul-zafrin/snake-validator/models"
	"github.com/amirul-zafrin/snake-validator/utilities"
	"github.com/gofiber/fiber/v2"
)

func GenerateFruit(c *fiber.Ctx) error {
	width := c.QueryInt("w", -1)
	height := c.QueryInt("h", -1)
	newSnake, err := utilities.GenerateNewGame(width, height)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "invalid request", "message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": newSnake})
}

func Validate(c *fiber.Ctx) error {
	var state models.StateWithTick
	if err := c.BodyParser(&state); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	prevX, prevY := state.Snake.X, state.Snake.Y
	prevVelX, prevVelY := 0, 0
	found := false
	for _, tick := range state.Ticks {
		currX, currY := prevX+tick.VelX, prevY+tick.VelY

		// Out of bound
		if currX < 0 || currX > state.Width || currY < 0 || currY > state.Width {
			return c.Status(418).JSON(fiber.Map{
				"status":  "Failed",
				"message": "Game is over, snake went out of bounds",
			})
		}

		// Invalid move
		if (tick.VelX != 0 && -prevVelX == tick.VelX) || (tick.VelY != 0 && -prevVelY == tick.VelY) {
			return c.Status(418).JSON(fiber.Map{
				"status":  "Failed",
				"message": "Game is over, player made an invalid move",
			})
		}

		if currX == state.Fruit.X && currY == state.Fruit.Y {
			found = true
		}

		prevX, prevY = currX, currY
		prevVelX, prevVelY = tick.VelX, tick.VelY
	}

	if found {
		newFruit := utilities.RandomFruitPositon(state.Height, state.Width)
		state.Fruit = newFruit
		state.Score++

		return c.Status(200).JSON(state)
	}

	return c.Status(404).JSON(fiber.Map{
		"status":  "Failed",
		"message": "Fruit not found, the ticks do not lead the snake to the fruit position!",
	})
}
