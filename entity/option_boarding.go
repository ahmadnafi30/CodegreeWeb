package entity

type OptionBoarding struct {
	ID           uint   `json:"id" gorm:"autoIncrement"`
	Description  string `json:"description"`
	OnboardingID uint   `json:"onboarding_id"`
	Pointrange   uint   `json:"pointrange"`
}
