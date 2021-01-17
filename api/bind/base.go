package bind

type QueryPage struct {
	PageNo   int64 `form:"page_no"`
	PageSize int64 `form:"page_size,default=20"`
}
