package deliveryhttp

import (
	"golden-server/domain/interfaces"
	"golden-server/interface/presenter"
	"log"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	service interfaces.UserService
}

func NewUserHandler(service interfaces.UserService, router fiber.Router) {
	handler := &userHandler{service}

	user := router.Group("/user")

	user.Get("/", handler.GetAll)
}

func (h *userHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAll(c.Context())

	if err != nil {
		log.Println(err.Error())
		return presenter.JsonResponse(c, fiber.StatusInternalServerError, nil)
	}

	return presenter.JsonResponse(c, fiber.StatusOK, users)
}
