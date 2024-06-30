package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Report struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	ReportID  string         `json:"report_id"`
	Name      string         `json:"name"`
	StartDate string         `json:"start_date"`
	EndDate   string         `json:"end_date"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ReqByDate struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type ReportRepository interface {
	RetrieveReports() ([]Report, error)
	RetrieveReportByID(id uint) (Report, error)
	CreateReport(report Report) error
	UpdateReportByID(report Report) (Report, error)
	RemoveReportByID(id uint) error
	RetriveByDate(startDate, endDate string) ([]Order, error)
}

type ReportUseCase interface {
	AddReport(ctx context.Context, report Report) error
	FetchReports(ctx context.Context) ([]Report, error)
	FetchReportByID(ctx context.Context, id uint) (Report, error)
	FetchReportByDate(ctx context.Context, startDate, endDate string) ([]Order, error)
}
