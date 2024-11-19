package models

type Manager struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"size:100"`
	Email string `gorm:"size:100"`
	Age   int
}
