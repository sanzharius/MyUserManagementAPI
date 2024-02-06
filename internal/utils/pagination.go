package utils

import (
	"math"
	"strconv"
)

const (
	defaultSize = 10
)

type PaginationQuery struct {
	Size    int    `bson:"size,omitempty"`
	Page    int    `bson:"page,omitempty"`
	OrderBy string `bson:"orderBy,omitempty"`
}

func GetPaginationFromCtx(page string, size string, orderBy string) (*PaginationQuery, error) {
	q := &PaginationQuery{}
	if err := q.SetPage(page); err != nil {
		return nil, err
	}

	if err := q.SetSize(size); err != nil {
		return nil, err
	}

	q.SetOrderBy(orderBy)

	return q, nil
}
func GetTotalPages(totalCount int, pageSize int) int {
	d := float64(totalCount) / float64(pageSize)

	return int(math.Ceil(d))
}

func GetHasMore(currentPage int, totalCount int, pageSize int) bool {
	if pageSize == 0 {
		return false
	}

	return currentPage < totalCount/pageSize
}

func (q *PaginationQuery) GetSkip() int {
	if q.Page == 0 {
		return 0
	}

	return (q.Page - 1) * q.Size
}

func (q *PaginationQuery) GetLimit() int {
	return q.Size
}
func (q *PaginationQuery) SetOrderBy(orderByQuery string) {
	q.OrderBy = orderByQuery
}

func (q *PaginationQuery) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		q.Size = defaultSize
		return nil
	}

	if sizeQuery == "0" {
		q.Size = defaultSize
		return nil
	}

	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}

	q.Size = n

	return nil
}

func (q *PaginationQuery) SetPage(pageQuery string) error {
	if pageQuery == "" {
		q.Size = 0
		return nil
	}

	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}

	q.Page = n

	return nil
}
