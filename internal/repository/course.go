package repository

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/model"
	"errors"

	"gorm.io/gorm"
)

type ICourse interface {
	SaveCourse(course *entity.Course) error
	SaveSubLanguage(sub_language *entity.SubLanguage) error
	SaveMaterial(Material *entity.Material) error
	SaveQuestion(gamification *entity.Question) error
	SaveOptionCourse(option entity.Option) error
	GetAllCourses() ([]model.GetCourse, error)
	GetCourse(id uint) (model.GetCoursedetail, error)
	GetMaterialsBySubLangID(subLangID uint) (model.GetSublang, error)
	GetQuestionsBySublangID(sublangId uint, id uint) ([]model.Gamification, error)
	GetOptionsByQuestionID(questionID uint) ([]entity.Option, error)
	SaveUserAnswer(answer *entity.UserAnswerGami) error
	CheckCorrectAnswer(quest uint, option uint) (bool, error)
	GetnameCerification(courseid uint) (string, error)
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

func (repo *CourseRepo) GetAllCourses() ([]model.GetCourse, error) {
	var courses []entity.Course
	if err := repo.db.Debug().Find(&courses).Error; err != nil {
		return nil, err
	}

	var result []model.GetCourse
	for _, course := range courses {
		result = append(result, model.GetCourse{
			Title:       course.Title,
			Description: course.Description,
		})
	}
	return result, nil
}
func (repo *CourseRepo) GetCourse(id uint) (model.GetCoursedetail, error) {
	var course *entity.Course
	if err := repo.db.Preload("SubLanguages").First(&course, id).Error; err != nil {
		return model.GetCoursedetail{}, err
	}

	var subLangs []string
	for _, lang := range course.SubLanguages {
		subLangs = append(subLangs, lang.Title)
	}

	result := model.GetCoursedetail{
		Progress: course.Progres,
		Sublang:  subLangs,
	}

	return result, nil
}

// func (repo *CourseRepo) GetSubLanguagesByCourseID(courseID uint) (model.get, error) {
// 	var subLanguages []entity.SubLanguage
// 	if err := repo.db.Where("course_id = ?", courseID).Find(&subLanguages).Error; err != nil {
// 		return nil, err
// 	}
// 	return subLanguages, nil
// }

func (repo *CourseRepo) GetMaterialsBySubLangID(subLangID uint) (model.GetSublang, error) {
	var sublang *entity.SubLanguage
	if err := repo.db.Where("id = ?", subLangID).Preload("Material").First(&sublang).Error; err != nil {
		return model.GetSublang{}, err
	}

	result := model.GetSublang{
		Title:       sublang.Title,
		Description: sublang.Description,
		Material:    sublang.Material.Material,
	}
	return result, nil
}

func (repo *CourseRepo) GetQuestionsBySublangID(sublangID uint, id uint) ([]model.Gamification, error) {
	var q entity.Question
	if err := repo.db.Where("id = ? AND sub_language_id = ?", id, sublangID).Preload("Options").First(&q).Error; err != nil {
		return nil, errors.New("this subab no longer has any questions")
	}

	var options []model.Option
	for _, opt := range q.Options {
		options = append(options, model.Option{
			ID:     opt.ID,
			Option: opt.Option,
		})
	}

	gamification := model.Gamification{
		QuestionID: q.ID,
		Question:   q.Question,
		Options:    options,
	}

	return []model.Gamification{gamification}, nil
}

func (repo *CourseRepo) GetOptionsByQuestionID(questionID uint) ([]entity.Option, error) {
	var options []entity.Option
	if err := repo.db.Where("id = ?", questionID).Find(&options).Error; err != nil {
		return nil, err
	}
	return options, nil
}

func (repo *CourseRepo) SaveUserAnswer(answer *entity.UserAnswerGami) error {
	if err := repo.db.Create(answer).Error; err != nil {
		return err
	}
	return nil
}

func (repo *CourseRepo) CheckCorrectAnswer(quest uint, option uint) (bool, error) {
	var correctOptID uint
	if err := repo.db.Model(&entity.Option{}).Where("question_id = ? AND value = true", quest).Pluck("id", &correctOptID).Error; err != nil {
		return false, errors.New("failed to find data")
	}

	return option == correctOptID, nil
}

func (repo *CourseRepo) GetnameCerification(courseid uint) (string, error) {
	var course entity.Course
	if err := repo.db.Where("id = ?", courseid).Find(&course).Error; err != nil {
		return "", err
	}
	result := course.Title

	return result, nil
}
