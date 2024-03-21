package entity

type Mentor struct {
	ID             uint `json:"id" gorm:"primary_key"`
	LanguageCodeID uint
	LanguageCode   LanguageCode `gorm:"foreignKey:LanguageCodeID"`
	Language       string
	Name           string `json:"name" gorm:"varchar(255) not null;"`
	Description    string `json:"description" gorm:"text"`
	Company        string `json:"company" gorm:"varchar(255)"`
	Linkwhatsapp   string `json:"link_whatsapp,omitempty" gorm:"varchar(100);"`
}
