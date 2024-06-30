package usecase

import (
	"context"
	"sujana-be-web-go/domain"
	"time"
)

type reportUseCase struct {
	reportRepository domain.ReportRepository
	contextTimeout   time.Duration
}

func NewReportUseCase(report domain.ReportRepository, t time.Duration) domain.ReportUseCase {
	return &reportUseCase{
		reportRepository: report,
		contextTimeout:   t,
	}
}

func (c *reportUseCase) AddReport(ctx context.Context, report domain.Report) error {
	err := c.reportRepository.CreateReport(report)
	if err != nil {
		return err
	}
	return nil
}

func (c *reportUseCase) FetchReportByID(ctx context.Context, id uint) (domain.Report, error) {
	res, err := c.reportRepository.RetrieveReportByID(id)
	if err != nil {
		return domain.Report{}, err
	}
	return res, nil
}

func (c *reportUseCase) FetchReports(ctx context.Context) ([]domain.Report, error) {
	res, err := c.reportRepository.RetrieveReports()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *reportUseCase) FetchReportByDate(ctx context.Context, startDate, endDate string) ([]domain.Order, error) {
	order, err := c.reportRepository.RetriveByDate(startDate, endDate)
	if err != nil {
		return nil, err
	}
	return order, nil
}
