package tasks

type Tasks struct {
	ID          uint `gorm:"primaryKey"`
	ISACTIVE    bool
	TITLE       string
	DESCRIPTION string
	CREATETIME  int64 `gorm:"autoCreateTime;index"`
}
