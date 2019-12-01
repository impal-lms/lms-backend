package domain

type Alumni struct {
	ID           int64  `json:"id" gorm:"primary_key"`
	Name         string `json:"name" binding:"required"`
	GraduateYear int    `json:"graduate_year" binding:"required"`
}
