package models
import("time")

type Todo struct {
	// ActivityId          int64  `gorm:"primaryKey" json:"id"`
	ActivityGroupId int64 `json:"activity_group_id"`
	Title   string `json:"title"`
	Priority   string `json:"priority"`
	IsActive   string `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}
