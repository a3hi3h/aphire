package model

type Account struct {
	BaseModel
	Status  uint   `gorm:"not null" json:"status"`
	Type    uint   `json:"type"`
	Address string `json:"address"`
}

type Org struct {
	BaseModel
	Names   string  `json:"names"`
	Account Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// User struct
type User struct {
	BaseModel
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Names    string `json:"names"`
	Type     uint   `json:"type"`
	Status   uint   `json:"status"`
}
