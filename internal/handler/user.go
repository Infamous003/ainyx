package handler

import (
	"time"

	"github.com/Infamous003/ainyx/internal/models"
	"github.com/Infamous003/ainyx/internal/repository"
	"github.com/Infamous003/ainyx/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type User struct {
	service *service.User
	logger  *zap.Logger
}

func NewUser(svc *service.User, logger *zap.Logger) *User {
	return &User{
		service: svc,
		logger:  logger,
	}
}

func (h *User) CreateUser(c *fiber.Ctx) error {
	var payload models.UserCreateRequest

	if err := c.BodyParser(&payload); err != nil {
		h.logger.Error("failed to parse payload", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Required-field check (because missing fields don't call UnmarshalJSON)
	if payload.Dob == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "dob is required",
		})
	}

	user, err := h.service.CreateUser(
		c.Context(),
		payload.Name,
		time.Time(*payload.Dob), // ← conversion happens here
	)
	if err != nil {
		h.logger.Error("failed to create user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *User) UpdateUser(c *fiber.Ctx) error {
	id, err := ReadIDParam(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	var payload models.UserUpdateRequest

	if err := c.BodyParser(&payload); err != nil {
		h.logger.Error("failed to parse payload", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if payload.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "name is required",
		})
	}

	if payload.Dob == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "dob is required",
		})
	}

	user, err := h.service.UpdateUser(
		c.Context(),
		int32(id),
		payload.Name,
		time.Time(*payload.Dob),
	)
	if err != nil {
		if err == repository.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}

		h.logger.Error("failed to update user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not update user",
		})
	}

	return c.JSON(user)
}

func (h *User) GetUser(c *fiber.Ctx) error {
	id, err := ReadIDParam(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	user, err := h.service.GetUser(c.Context(), int32(id))
	if err != nil {
		if err == repository.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}
		h.logger.Error("failed to get user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not get user"})
	}

	return c.JSON(user)
}

func (h *User) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers(c.Context())
	if err != nil {
		h.logger.Error("failed to list users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not list users"})
	}

	return c.JSON(users)
}

func (h *User) DeleteUser(c *fiber.Ctx) error {
	id, err := ReadIDParam(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	err = h.service.DeleteUser(c.Context(), int32(id))
	if err != nil {
		if err == repository.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}
		h.logger.Error("failed to delete user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not delete user"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
