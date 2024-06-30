package delivery

import (
	"strconv"
	"sujana-be-web-go/domain"
	"sujana-be-web-go/middleware"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	AccountUC domain.AccountUseCase
}

func NewAccountHandler(c *fiber.App, das domain.AccountUseCase) {
	handler := &AccountHandler{
		AccountUC: das,
	}
	api := c.Group("/account")

	public := api.Group("/public")
	public.Post("/login", handler.Login)

	private := api.Group("/private")
	private.Get("/user", handler.GetAllAccount)
	private.Post("/user", handler.CreateAccount)
	private.Put("/user/:id", handler.UpdateAccount)
	private.Delete("/user/:id", handler.DeleteAccount)

	private.Get("/profile", middleware.ValidateToken, handler.GetProfile)
}

func (t *AccountHandler) GetAllAccount(c *fiber.Ctx) error {
	res, err := t.AccountUC.FetchAccounts(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Successfully get all account",
	})
}

func (t *AccountHandler) Login(c *fiber.Ctx) error {
	req := new(domain.LoginPayload)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	valRes, er := govalidator.ValidateStruct(req)
	if !valRes {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   er.Error(),
		})
	}
	res, token, er := t.AccountUC.LoginAccount(c.Context(), req)
	if er != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": er,
			"error":   er.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"token":   token,
		"message": "Successfully login",
	})
}

func (t *AccountHandler) GetProfile(c *fiber.Ctx) error {
	id := middleware.AccountID(c)
	res, err := t.AccountUC.FetchAccountByID(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Successfully get profile",
	})
}

func (t *AccountHandler) CreateAccount(c *fiber.Ctx) error {
	req := new(domain.Account)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	valRes, er := govalidator.ValidateStruct(req)
	if !valRes {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   er.Error(),
		})
	}
	res, err := t.AccountUC.CreateAccount(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"success": true,
		"data":    res,
		"message": "Successfully create account",
	})
}

func (t *AccountHandler) UpdateAccount(c *fiber.Ctx) error {
	req := new(domain.Account)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	valRes, er := govalidator.ValidateStruct(req)
	if !valRes {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   er.Error(),
		})
	}
	res, err := t.AccountUC.UpdateAccount(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Successfully update account",
	})
}

func (t *AccountHandler) DeleteAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	strId, erStr := strconv.Atoi(id)
	if erStr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse id",
			"error":   erStr.Error(),
		})
	}
	err := t.AccountUC.DeleteAccount(c.Context(), uint(strId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"message": "Successfully delete account",
	})
}
