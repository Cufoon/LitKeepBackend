package dto

type QueryStatisticsDayData struct {
	Day   string  `gorm:"column:day" json:"day"`
	Money float64 `gorm:"column:total" json:"money"`
}
