package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ReadIDParam(c *fiber.Ctx) (int64, error) {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}
