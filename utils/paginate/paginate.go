package paginate

import (
	"net/url"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Pagination struct {
	Page   int
	Limit  int
}

func Paginate(page, length int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (page - 1) * length
		return db.Offset(offset).Limit(length)
	}
}

func GetParams(queryParam url.Values) Pagination {
	params := Pagination{
		Page:   cast.ToInt(queryParam.Get("page")),
		Limit:  cast.ToInt(queryParam.Get("limit")),
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	if params.Limit > 30 {
		params.Limit = 30
	}

	return params
}
