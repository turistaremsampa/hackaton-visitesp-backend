package models

type Feedback struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	UserID     uint   `json:"user_id"`
	LocationID uint   `json:"location_id"`
	Comment    string `json:"comment"`
	Rating     int    `json:"rating"`
	VisitDate  string `json:"visit_date"`
}
