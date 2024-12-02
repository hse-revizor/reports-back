package app

import (
	"database/sql"

	"github.com/hse-revizor/reports-back/internal/common/app"
	"github.com/hse-revizor/reports-back/internal/report/domain"
)

type ReportService struct {
	db *sql.DB
}

func CreateReportService(db *sql.DB) *ReportService {
	return &ReportService{db: db}
}

func (s *ReportService) GetAllReports() ([]domain.Report, error) {
	rows, err := s.db.Query("SELECT id, title, description, created_at FROM reports")
	if err != nil {
			return nil, app.InternalErr
	}
	defer rows.Close()

	reports := make([]domain.Report, 0)
	for rows.Next() {
		var r domain.Report
		err := rows.Scan(&r.ID, &r.Title, &r.Description, &r.CreatedAt)

		if err != nil {
			return nil, app.InternalErr
		}
		reports = append(reports, r)
	}

	return reports, nil
}

func (s *ReportService) GetReportByID(id domain.ReportId) (*domain.Report, error) {
	var report domain.Report
	err := s.db.QueryRow("SELECT id, title, description, created_at FROM reports WHERE id = $1", id).
			Scan(&report.ID, &report.Title, &report.Description, &report.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, app.NotFoundErr
	} else if err != nil {
		return nil, app.InternalErr
	}

	return &report, nil
}
