package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// PageInfo Paging common input parameter structure
type ResponsePageInfo struct {
	CurrentPage int64 `json:"currentPage"`                // 页码
	Total       int64 `json:"total" form:"total"`         // 页码
	TotalPage   int64 `json:"totalPage" form:"totalPage"` // 页码

}

type ReturnPage struct {
	List             interface{} `json:"list"`
	ResponsePageInfo `json:"pageInfo"`
}
