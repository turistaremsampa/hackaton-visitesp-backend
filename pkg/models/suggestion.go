package models

type Suggestion struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id"`
	RouteID     uint   `json:"route_id"`
	LocationID  uint   `json:"location_id"`
	SuggestedOn string `json:"suggested_on"`
	Details     string `json:"details"`
}
