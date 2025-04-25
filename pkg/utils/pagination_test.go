package utils

import (
    "testing"
)

func TestNewPagination(t *testing.T) {
    tests := []struct {
        name          string
        page          int
        pageSize      int
        totalItems    int
        expectedError bool
    }{
        {
            name:          "Valid pagination parameters",
            page:          1,
            pageSize:      10,
            totalItems:    100,
            expectedError: false,
        },
        {
            name:          "Invalid page number",
            page:          0,
            pageSize:      10,
            totalItems:    100,
            expectedError: true,
        },
        {
            name:          "Invalid page size",
            page:          1,
            pageSize:      0,
            totalItems:    100,
            expectedError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            pagination, err := NewPagination(tt.page, tt.pageSize, tt.totalItems)
            if tt.expectedError {
                if err == nil {
                    t.Error("expected error, got nil")
                }
            } else {
                if err != nil {
                    t.Errorf("unexpected error: %v", err)
                }
                if pagination.Page != tt.page {
                    t.Errorf("expected page %d, got %d", tt.page, pagination.Page)
                }
                if pagination.PageSize != tt.pageSize {
                    t.Errorf("expected pageSize %d, got %d", tt.pageSize, pagination.PageSize)
                }
                if pagination.TotalItems != tt.totalItems {
                    t.Errorf("expected totalItems %d, got %d", tt.totalItems, pagination.TotalItems)
                }
            }
        })
    }
}

func TestPagination_TotalPages(t *testing.T) {
    tests := []struct {
        name       string
        pagination *Pagination
        want       int
    }{
        {
            name: "Even division",
            pagination: &Pagination{
                PageSize:   10,
                TotalItems: 100,
            },
            want: 10,
        },
        {
            name: "Uneven division",
            pagination: &Pagination{
                PageSize:   10,
                TotalItems: 95,
            },
            want: 10,
        },
        {
            name: "Zero page size",
            pagination: &Pagination{
                PageSize:   0,
                TotalItems: 100,
            },
            want: 0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.pagination.TotalPages(); got != tt.want {
                t.Errorf("Pagination.TotalPages() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPagination_HasNext(t *testing.T) {
    tests := []struct {
        name       string
        pagination *Pagination
        want       bool
    }{
        {
            name: "Has next page",
            pagination: &Pagination{
                Page:       1,
                PageSize:   10,
                TotalItems: 100,
            },
            want: true,
        },
        {
            name: "No next page",
            pagination: &Pagination{
                Page:       10,
                PageSize:   10,
                TotalItems: 100,
            },
            want: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.pagination.HasNext(); got != tt.want {
                t.Errorf("Pagination.HasNext() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPagination_HasPrevious(t *testing.T) {
    tests := []struct {
        name       string
        pagination *Pagination
        want       bool
    }{
        {
            name: "Has previous page",
            pagination: &Pagination{
                Page:       2,
                PageSize:   10,
                TotalItems: 100,
            },
            want: true,
        },
        {
            name: "No previous page",
            pagination: &Pagination{
                Page:       1,
                PageSize:   10,
                TotalItems: 100,
            },
            want: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.pagination.HasPrevious(); got != tt.want {
                t.Errorf("Pagination.HasPrevious() = %v, want %v", got, tt.want)
            }
        })
    }
}
