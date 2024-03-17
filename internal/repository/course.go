package repository

import (
	"CodegreeWebbs/entity"

	"gorm.io/gorm"
)

type ICourse interface {
	SaveCourse(course *entity.Course) error
	SaveSubLanguage(sub_language *entity.SubLanguage) error
	SaveMaterial(Material *entity.Material) error
	SaveQuestion(gamification *entity.Question) error
	SaveOptionCourse(option entity.Option) error
	GetAllCourses() ([]entity.Course, error)
	GetSubLanguagesByCourseID(courseID uint) ([]entity.SubLanguage, error)
	GetMaterialsBySubLangID(subLangID uint) ([]entity.Material, error)
	GetQuestionsByMaterialID(materialID uint) ([]entity.Question, error)
	GetOptionsByQuestionID(questionID uint) ([]entity.Option, error)
}

type CourseRepo struct {
	db *gorm.DB
}

func NewCourseRepo(db *gorm.DB) ICourse {
	return &CourseRepo{db: db}
}

func (repo *CourseRepo) SaveCourse(course *entity.Course) error {
	if err := repo.db.Debug().Create(course).Error; err != nil {
		return err
	}
	return nil
}

func (repo *CourseRepo) SaveSubLanguage(sub_language *entity.SubLanguage) error {
	if err := repo.db.Debug().Create(sub_language).Error; err != nil {
		return err
	}
	return nil
}

func (repo *CourseRepo) SaveMaterial(Material *entity.Material) error {
	if err := repo.db.Debug().Create(Material).Error; err != nil {
		return err
	}
	return nil
}

func (repo *CourseRepo) SaveQuestion(gamification *entity.Question) error {
	if err := repo.db.Debug().Create(gamification).Error; err != nil {
		return err
	}
	return nil
}

func (repo *CourseRepo) SaveOptionCourse(option entity.Option) error {
	if err := repo.db.Debug().Create(option).Error; err != nil {
		return err
	}
	return nil
}

func (repo *CourseRepo) GetAllCourses() ([]entity.Course, error) {
	var courses []entity.Course
	if err := repo.db.Preload("SubLanguages").Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (repo *CourseRepo) GetSubLanguagesByCourseID(courseID uint) ([]entity.SubLanguage, error) {
	var subLanguages []entity.SubLanguage
	if err := repo.db.Where("course_id = ?", courseID).Find(&subLanguages).Error; err != nil {
		return nil, err
	}
	return subLanguages, nil
}

func (repo *CourseRepo) GetMaterialsBySubLangID(subLangID uint) ([]entity.Material, error) {
	var materials []entity.Material
	if err := repo.db.Where("sub_lang_id = ?", subLangID).Find(&materials).Error; err != nil {
		return nil, err
	}
	return materials, nil
}

func (repo *CourseRepo) GetQuestionsByMaterialID(materialID uint) ([]entity.Question, error) {
	var questions []entity.Question
	if err := repo.db.Where("material_id = ?", materialID).Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func (repo *CourseRepo) GetOptionsByQuestionID(questionID uint) ([]entity.Option, error) {
	var options []entity.Option
	if err := repo.db.Where("question_id = ?", questionID).Find(&options).Error; err != nil {
		return nil, err
	}
	return options, nil
}
