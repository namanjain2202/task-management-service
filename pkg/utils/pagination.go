package utils

import (
	"errors"
	"math"
)

// Pagination represents the pagination parameters.
type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalItems int `json:"total_items"`
}

// NewPagination creates a new Pagination instance.
func NewPagination(page, pageSize, totalItems int) (*Pagination, error) {
	if page < 1 {
		return nil, errors.New("page must be greater than 0")
	}
	if pageSize < 1 {
		return nil, errors.New("page size must be greater than 0")
	}
	return &Pagination{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
	}, nil
}

// TotalPages calculates the total number of pages based on total items and page size.
func (p *Pagination) TotalPages() int {
	if p.PageSize == 0 {
		return 0
	}
	return int(math.Ceil(float64(p.TotalItems) / float64(p.PageSize)))
}

// HasNext checks if there is a next page.
func (p *Pagination) HasNext() bool {
	return p.Page < p.TotalPages()
}

// HasPrevious checks if there is a previous page.
func (p *Pagination) HasPrevious() bool {
	return p.Page > 1
}