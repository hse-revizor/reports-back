package infra

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	common "github.com/hse-revizor/reports-back/internal/common/app"
	"github.com/hse-revizor/reports-back/internal/report/app"
	"github.com/hse-revizor/reports-back/internal/report/domain"
)

type ReportController struct {
	service *app.ReportService
}

func CreateReportController(r *gin.RouterGroup, s *app.ReportService) *ReportController {
	c := ReportController{s}

	r.GET("/all", c.GetAllReports)
	r.GET("/one/:id", c.GetReportByID)

	return &c
}

// GetAllReports godoc
// @Summary Getting all the existing reports
// @Success 200 {string} string
// @Failure 500 {string} string
// @Router /reports/all [get]
func (c *ReportController) GetAllReports(ctx *gin.Context) {
	reports, err := c.service.GetAllReports()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"reports": reports,
	})
}

// GetReportByID godoc
// @Summary Getting a specific report
// @Param id path string true "id"
// @Success 200 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /reports/one/{id} [get]
func (c *ReportController) GetReportByID(ctx *gin.Context) {
	id := ctx.Param("id")
	report, err := c.service.GetReportByID(domain.ReportId(id))

	if err != nil {
		if errors.Is(err, common.NotFoundErr) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"err": err,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err,
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, report)
}

