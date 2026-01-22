package utils

import "math"

type PageResponse struct {
	List      interface{} `json:"list"`       // 数据列表
	Total     int64       `json:"total"`      // 总条数
	Page      int         `json:"page"`       // 当前页码
	Size      int         `json:"size"`       // 每页条数
	TotalPage int         `json:"total_page"` // 总页数
}

// NewPageResponse 构建分页响应
func NewPageResponse(list interface{}, total int64, page, size int) *PageResponse {
	if size <= 0 {
		size = 10
	}
	return &PageResponse{
		List:      list,
		Total:     total,
		Page:      page,
		Size:      size,
		TotalPage: int(math.Ceil(float64(total) / float64(size))), // 向上取整计算总页数
	}
}
