package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/kvendingoldo/gu-common/pkg/gu-common/model"
	"strconv"
)

//GeneratePaginationFromRequest ...
func GeneratePaginationFromRequest(c *gin.Context) model.Pagination {
	limit := 2
	page := 1
	sort := "created_at asc"

	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}
	return model.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
