package formdata

type Pager struct {
	Page     int `json:"page" query:"page"`
	PageSize int `json:"pageSize" query:"pageSize"`
}

type PostAdvInfo struct {
	PostId  int      `json:"id"`
	AdsInfo []string `json:"adv_info"`
}
