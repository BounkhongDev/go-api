package paginates

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaginateRequest struct {
	Limit int `json:"limit" validate:"required"`
	Page  int `json:"page" validate:"required"`
}

type PaginatedResponse struct {
	Rows         interface{} `json:"rows"`
	Count        int         `json:"count"`
	CountPage    int         `json:"countPage"`
	CurrentPage  int         `json:"currentPage"`
	NextPage     int         `json:"nextPage"`
	PreviousPage int         `json:"previousPage"`
}

func Paginate(db *gorm.DB, model interface{}, paginate PaginateRequest, results interface{}) (*PaginatedResponse, error) {

	if paginate.Limit <= 0 {
		return nil, errors.New("limit must be greater than 0")
	}
	var total int64

	db.Model(model).Count(&total)
	countPage := (int(total) + paginate.Limit - 1) / paginate.Limit
	offset := (paginate.Page - 1) * paginate.Limit

	// Fetch paginated results
	result := db.Preload(clause.Associations).Limit(paginate.Limit).Offset(offset).Find(results)
	if result.Error != nil {
		return nil, result.Error
	}

	nextPage := paginate.Page + 1
	if nextPage > countPage {
		nextPage = 0
	}

	previousPage := paginate.Page - 1
	if previousPage < 1 {
		previousPage = 0
	}

	pagination := &PaginatedResponse{
		Count:        int(total),
		CountPage:    countPage,
		CurrentPage:  paginate.Page,
		NextPage:     nextPage,
		PreviousPage: previousPage,
		Rows:         results,
	}
	return pagination, nil
}
