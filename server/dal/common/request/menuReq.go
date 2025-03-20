package request

type MenuReq struct {
	PageInfo
	KeyWord string `json:"keyWord" form:"keyWord"`
}
