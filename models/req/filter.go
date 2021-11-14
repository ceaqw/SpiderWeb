package req

type Filter struct {
	DateRangeType uint8  `json:"dateRangeType"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	PlatForm      string `json:"platForm"`
	Project       string `json:"project"`
	ShowType      uint8  `json:"showType"`
}
