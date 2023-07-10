package interfaces

// 分页请求
type Pages struct {
	Page     int `form:"page,default=1" json:"page"`            // 当前页，默认:1
	PageSize int `form:"page_size,default=50" json:"page_size"` //每页显示数量，默认:50，最大:100
}

// 分页数据结构
type Pagination struct {
	*Pages
	Count int64       `json:"count"` // 总数
	Data  interface{} `json:"data"`  // 数据
}

// 分页数据返回
func PaginationRsp(page, pageSize int, count int64, data interface{}) *Pagination {
	return &Pagination{
		Pages: &Pages{
			Page:     page,
			PageSize: pageSize,
		},
		Count: count,
		Data:  data,
	}
}
