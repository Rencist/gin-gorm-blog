package dto

import "gin-gorm-blog/entity"

type Meta struct {
	Page      int   `json:"page"`
	MaxPage   int   `json:"max_page"`
	TotalData int64 `json:"total_data"`
}

type PaginationResponse struct {
	DataPerPage any  `json:"data_per_page"`
	Meta        Meta `json:"meta"`
}

type BlogPaginationResponse struct {
	Blog  			BlogResponseDto		`json:"blog"`
	Tags			[]entity.Tag		`json:"tag"`
	Comments    	[]entity.Comment	`json:"comment_per_page"`
	Meta        	Meta 				`json:"meta"`
}