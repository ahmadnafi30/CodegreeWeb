package entity

type OptionBoarding struct {
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Description  string `json:"description" gorm:"type:text"`
	OnboardingID uint   `json:"onboarding_id"`
	Pointrange   uint   `json:"pointrange"`
}
