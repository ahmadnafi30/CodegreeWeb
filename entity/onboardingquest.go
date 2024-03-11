package entity

// import "github.com/google/uuid"

type Onboarding struct {
	ID       uint             `json:"id" gorm:"autoIncrement"`
	Question string           `json:"question" gorm:"type:text"`
	Options  []OptionBoarding `json:"options" gorm:"foreignKey:OnboardingID"`
}
