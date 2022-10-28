package req

import (
	"time"
)

type Filter struct {
	DateRangeType uint8  `json:"dateRangeType"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	PlatForm      string `json:"platForm"`
	Project       string `json:"project"`
	ShowType      uint8  `json:"showType"`
}

func FilterVerify(filter *Filter) {
	// 时间特殊处理
	baseLayout := "2006-01-02"
	dateRangeCount := []int{-7, -3}
	// 昨天数据
	if filter.DateRangeType == 2 {
		filter.StartDate = time.Now().AddDate(0, 0, -1).Format(baseLayout)
		filter.EndDate = filter.StartDate
	}
	// 没有开始时间
	if filter.StartDate == "" {
		endDate := time.Now()
		if filter.EndDate != "" {
			endDate, _ = time.Parse(baseLayout, filter.EndDate)
		}
		if filter.DateRangeType < 2 {
			filter.StartDate = endDate.AddDate(0, 0, dateRangeCount[filter.DateRangeType]).Format(baseLayout)
		} else {
			filter.StartDate = time.Now().Format(baseLayout)
		}
	}
	// 没有结束时间
	if filter.EndDate == "" {
		startDate, _ := time.Parse(baseLayout, filter.StartDate)
		if filter.DateRangeType < 2 {
			filter.EndDate = startDate.AddDate(0, 0, dateRangeCount[filter.DateRangeType]).Format(baseLayout)
		} else {
			filter.EndDate = startDate.Format(baseLayout)
		}
	}
}
