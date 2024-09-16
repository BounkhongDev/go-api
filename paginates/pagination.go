package paginates

import "math"

// PaginateRequest represents the structure for pagination input parameters
type PaginateRequest struct {
	Limit int `json:"limit" validate:"required"`
	Page  int `json:"page" validate:"required"`
}

// PaginatedResponse represents the structure for pagination output data
type PaginatedResponse struct {
	Rows         interface{} `json:"rows"`
	Count        int64       `json:"count"`
	CountPage    int         `json:"countPage"`
	CurrentPage  int         `json:"currentPage"`
	NextPage     int         `json:"nextPage"`
	PreviousPage int         `json:"previousPage"`
}

// IFindAndCountAll represents the structure for input data to be paginated
type IFindAndCountAll struct {
	Count int64       `json:"count"`
	Rows  interface{} `json:"rows"`
}

// PaginationResult performs pagination calculations and returns the paginated response
func PaginationResult(page int, limit int, result IFindAndCountAll) PaginatedResponse {
	// Calculate total number of pages
	countPage := float64(result.Count) / float64(limit)
	if countPage != math.Floor(countPage) {
		countPage = math.Floor(countPage) + 1
	}

	// Calculate next and previous pages
	var nextPage, previousPage int
	if int(countPage) > page {
		nextPage = page + 1
	} else {
		nextPage = 0
	}

	if page > 1 {
		previousPage = page - 1
	} else {
		previousPage = 0
	}

	// Return the paginated response
	return PaginatedResponse{
		Rows:         result.Rows,
		Count:        result.Count,
		CountPage:    int(countPage),
		CurrentPage:  page,
		NextPage:     nextPage,
		PreviousPage: previousPage,
	}
}
