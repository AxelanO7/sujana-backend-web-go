package delivery

import (
	"strconv"
	"sujana-be-web-go/domain"

	"github.com/gofiber/fiber/v2"
)

type ReportHandler struct {
	ReportUC domain.ReportUseCase
}

func NewReportHandler(c *fiber.App, das domain.ReportUseCase) {
	handler := &ReportHandler{
		ReportUC: das,
	}
	api := c.Group("/report")

	_ = api.Group("/public")

	private := api.Group("/private")
	private.Get("/package", handler.GetAllReport)
	private.Get("/package/:id", handler.GetReportByID)
	private.Post("/package", handler.AddReport)

	private.Post("/date", handler.GetReportByDate)
}

func (t *ReportHandler) GetAllReport(c *fiber.Ctx) error {
	res, err := t.ReportUC.FetchReports(c.Context())
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
		"message": "Successfully get all user",
	})
}

func (t *ReportHandler) GetReportByID(c *fiber.Ctx) error {
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
	res, err := t.ReportUC.FetchReportByID(c.Context(), uint(strId))
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
		"message": "Successfully get user by id",
	})
}

func (t *ReportHandler) GetReportByDate(c *fiber.Ctx) error {
	req := new(domain.ReqByDate)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  422,
			"success": false,
			"message": "Failed to parse request",
			"error":   err.Error(),
		})
	}
	startDate := req.StartDate
	endDate := req.EndDate
	res, err := t.ReportUC.FetchReportByDate(c.Context(), startDate, endDate)
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
		"message": "Successfully get user by date",
	})
}

func (t *ReportHandler) AddReport(c *fiber.Ctx) error {
	req := new(domain.Report)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  422,
			"success": false,
			"message": "Failed to parse request",
			"error":   err.Error(),
		})
	}
	err := t.ReportUC.AddReport(c.Context(), *req)
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
		"message": "Successfully add user",
	})
}
