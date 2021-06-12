package dbUtils

import (
	"strconv"

	"github.com/DdZ-Fred/fiber-server-1/models"
	"github.com/gofiber/fiber/v2"
)

func GeneratePaginationFromRequest(c *fiber.Ctx) models.Pagination {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("perPage", "10"))

	return models.Pagination{
		Page:       page,
		PerPage:    perPage,
		TotalPages: 1,
	}
}

func SetPaginationTotalPages(pagination *models.Pagination, totalRows int64) {
	if int(totalRows)%pagination.PerPage == 0 {
		pagination.TotalPages = int(totalRows) / pagination.PerPage
	} else {
		pagination.TotalPages = (int(totalRows) / pagination.PerPage) + 1
	}
}
