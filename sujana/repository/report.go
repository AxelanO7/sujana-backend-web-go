package repository

import (
	"fmt"
	"sujana-be-web-go/domain"

	"gorm.io/gorm"
)

type posgreReportRepository struct {
	DB *gorm.DB
}

func NewPostgreReport(client *gorm.DB) domain.ReportRepository {
	return &posgreReportRepository{
		DB: client,
	}
}

func (a *posgreReportRepository) RetrieveReports() ([]domain.Report, error) {
	var res []domain.Report
	err := a.DB.
		Model(domain.Report{}).
		Find(&res).Error
	if err != nil {
		return []domain.Report{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreReportRepository) CreateReport(report domain.Report) error {
	err := a.DB.
		Model(domain.Report{}).
		Create(&report).Error
	if err != nil {
		return err
	}
	fmt.Println(report)
	return nil
}

func (a *posgreReportRepository) RetrieveReportByID(id uint) (domain.Report, error) {
	var res domain.Report
	err := a.DB.
		Model(domain.Report{}).
		Where("id = ?", id).
		First(&res).Error
	if err != nil {
		return domain.Report{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreReportRepository) UpdateReportByID(report domain.Report) (domain.Report, error) {
	err := a.DB.
		Model(domain.Report{}).
		Where("id = ?", report.ID).
		Updates(&report).Error
	if err != nil {
		return domain.Report{}, err
	}
	fmt.Println(report)
	return report, nil
}

func (a *posgreReportRepository) RemoveReportByID(id uint) error {
	err := a.DB.
		Model(domain.Report{}).
		Where("id = ?", id).
		Delete(&domain.Report{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *posgreReportRepository) RetriveByDate(startDate, endDate string) ([]domain.Order, error) {
	var res []domain.Order
	err := a.DB.
		Model(domain.Order{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Find(&res).Error
	if err != nil {
		return []domain.Order{}, err
	}
	fmt.Println(res)
	return res, nil
}
