package controller

import (
	"fmt"
	"strconv"

	"github.com/Darari17/user-management/fiber/v2/model/dto"
	"github.com/Darari17/user-management/fiber/v2/service"
	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	Service service.UserService
	router  fiber.Router
}

func NewUserController(service service.UserService, router fiber.Router) *UserControllerImpl {
	return &UserControllerImpl{Service: service, router: router}
}

func (u *UserControllerImpl) Route() {
	u.router.Post("/user", u.CreateController)
	u.router.Put("/user", u.UpdateController)
	u.router.Delete("/user/:id", u.DeleteController)
	u.router.Get("/user", u.GetController)
	u.router.Get("/user/:id", u.FindByIdController)
}

func (u *UserControllerImpl) CreateController(c *fiber.Ctx) error {
	var payload dto.CreateRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON Format",
		})
	}

	user, err := u.Service.CreateService(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create a new user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   user,
	})
}

func (u *UserControllerImpl) UpdateController(c *fiber.Ctx) error {
	var payload dto.UpdateRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON Format",
		})
	}

	user, err := u.Service.UpdateService(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   user,
	})
}

func (u *UserControllerImpl) DeleteController(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	err = u.Service.DeleteService(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}

func (u *UserControllerImpl) GetController(c *fiber.Ctx) error {
	users, err := u.Service.GetService()
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   users,
	})
}

func (u *UserControllerImpl) FindByIdController(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := u.Service.FindByIdService(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("User with ID %d not found", id),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   user,
	})
}
