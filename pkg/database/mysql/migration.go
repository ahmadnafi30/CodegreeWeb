package mysql

import (
	"CodegreeWebbs/entity"
	"log"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Onboarding{},
		&entity.OptionBoarding{},
		&entity.UserProfile{},
		&entity.UserAnswerOnBoarding{},
		&entity.Course{},
		&entity.SubLanguage{},
		&entity.LanguageCode{},
		&entity.Recom{},
		&entity.Material{},
		&entity.Question{},
		&entity.Option{},
		&entity.Payment{},
		&entity.UserAnswerGami{},
		&entity.Mentor{},
	); err != nil {
		log.Fatalf("failed migration db: %v", err)
		return err
	}
	return nil
}
