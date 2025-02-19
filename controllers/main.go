package controllers

type PageInfo struct {
	CurentPage int `json:"current_page"`
	NextPage   int `json:"next_page"`
	PrevPage   int `json:"prev_page"`
	TotalPage  int `json:"total_page"`
	TotalData  int `json:"total_data"`
}

type Response struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	PageInfo any    `json:"page_info,omitempty"`
	Results  any    `json:"results,omitempty"`
}
