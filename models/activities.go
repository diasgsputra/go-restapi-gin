package models
import("time")

type Activities struct {
	ActivityId          int64  `gorm:"primaryKey" json:"id"`
	Title string `gorm:"type:varchar(50)" json:"nama_product"`
	Email   string `gorm:"type:varchar(50)" json:"deskripsi"`
	CreatedAt   time.Time `gorm:"type:text" json:"deskripsi"`
}
