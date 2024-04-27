package pagination

import (
	"fmt"
	"gin-api/domain/types/apiErros"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

type PaginatedReturn struct {
	Page     int64 `json:"page"`
	Limit    int64 `json:"limit"`
	PageSize int64 `json:"pageSize"`
	Data     any   `json:"data"`
}

type Pagination struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

func defaultPagination() Pagination {
	return Pagination{
		Page:  1,
		Limit: 1000,
	}
}

func Validate(c *gin.Context) Pagination {
	apiErros.NewApiError()

	// sets a default pagination
	pagination := defaultPagination()

	pageString := c.Query("page")
	page, err := strconv.Atoi(pageString)
	// sets the value of page to default if less then 1
	if err == nil && page > 0 {
		fmt.Println("here")

		pagination.Page = int64(page)
	}

	limitString := c.Query("limit")
	limit, err := strconv.Atoi(limitString)
	// sets the value of limit to default if less then 1
	if err == nil && limit > 0 {
		pagination.Limit = int64(limit)
	}

	return pagination
}

func (p *Pagination) PaginatedInfo(data any) PaginatedReturn {
	var dataValue interface{}
	var pageSize int64

	// If data is nil, set it as an empty slice
	if data == nil {
		dataValue = make([]interface{}, 0)
		pageSize = 0
	} else {
		// Otherwise, use the original data
		dataValue = data

		// Get the length of the data if it's a slice
		val := reflect.ValueOf(data)
		if val.Kind() == reflect.Slice {
			pageSize = int64(val.Len())
		} else {
			// If data is not a slice, set page size to 1
			pageSize = 1
		}
	}

	return PaginatedReturn{
		Page:     p.Page,
		Limit:    p.Limit,
		Data:     dataValue,
		PageSize: pageSize,
	}
}
