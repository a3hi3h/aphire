package model

// Product struct
type Product struct {
	BaseModel
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Amount      int    `gorm:"not null" json:"amount"`
}

type question struct {
	BaseModel
	Type     uint   `json:"type"`
	Status   uint   `json:"status"`
	Level    uint   `json:"level"`
	Question string `json:"question"`
}

type exam struct {
	BaseModel
	Name   string `json:"name"`
	Time   uint   `json:"time"`
	Type   uint   `json:"type"`
	Status uint   `json:"status"`
	Level  uint   `json:"level"`
}
