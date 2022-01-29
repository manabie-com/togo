package tasks

type Tasks struct {
	Id         uint   `json:"id" gorm:"primaryKey"`
	IsActive   bool   `json:"isActive" gorm:"default:true"`
	Title      string `json:"title"`
	Desciption string `json:"description"`
	CreatedAt  int64  `json:"createdAt" gorm:"autoCreateTime;index"`
	UpdatedAt  int64  `json:"updatedAt" gorm:"autoUpdateTime"`
}
