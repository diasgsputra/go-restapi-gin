package models
import("time")

type Activities struct {
	// ActivityId          int64  `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Email   string `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
}
