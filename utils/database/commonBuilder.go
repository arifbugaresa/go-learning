package database

import (
	"github.com/doug-martin/goqu/v9"
	"go-learning/utils/common"
	"strings"
)

func BuildDatasetPaginationWithTotalData(dataset *goqu.SelectDataset, page *int64, limit *int64, sortField string, sortOrder string) (ds *goqu.SelectDataset, count int64, err error) {
	count, err = dataset.GroupBy("id").Count()
	if err != nil {
		return dataset, 0, err
	}

	if !common.IsEmptyField(sortOrder) {
		if !common.IsEmptyField(sortField) {
			if strings.ToUpper(sortOrder) == "DESC" {
				dataset = dataset.Order(
					goqu.I(sortField).Desc(),
				)
			} else {
				dataset = dataset.Order(
					goqu.I(sortField).Asc(),
				)
			}
		}
	}

	if !common.IsEmptyField(page) {
		if !common.IsEmptyField(limit) {
			offset := (*page - 1) * *limit
			dataset = dataset.Limit(uint(*limit)).Offset(uint(offset))
		}
	}

	return dataset, count, nil
}
