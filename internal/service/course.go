package service

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/internal/repository"
	"CodegreeWebbs/model"

	"github.com/google/uuid"
)

type Scourse interface {
	CreateCourse(courseData *entity.Course) error
	GetAllCourses() ([]model.GetCourse, error)
	SelectSubLang(id uint) (model.GetSublang, error)
	Selectacourse(id uint) (model.GetCoursedetail, error)
	GetGamification(sublangId uint, id uint) ([]model.Gamification, error)
	CheckAnswer(userID uuid.UUID, quest uint, option uint) (bool, error)
	GetnameCerification(name string, courseid uint) (model.Certification, error)
}

type CourseService struct {
	CourseRepo repository.ICourse
}

func NewCourseService(courseRepo repository.ICourse) Scourse {
	return &CourseService{
		CourseRepo: courseRepo,
	}
}

func (s *CourseService) CreateCourse(courseData *entity.Course) error {
	err := s.CourseRepo.SaveCourse(courseData)
	if err != nil {
		return err
	}

	for i := range courseData.SubLanguages {
		subLang := &courseData.SubLanguages[i]
		err := s.CourseRepo.SaveSubLanguage(subLang)
		if err != nil {
			return err
		}

		material := subLang.Material
		err = s.CourseRepo.SaveMaterial(&material)
		if err != nil {
			return err
		}
		for j := range subLang.Questions {
			question := &subLang.Questions[j]
			err := s.CourseRepo.SaveQuestion(question)
			if err != nil {
				return err
			}
			for k := range question.Options {
				option := &question.Options[k]
				err := s.CourseRepo.SaveOptionCourse(*option)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *CourseService) GetAllCourses() ([]model.GetCourse, error) {
	return s.CourseRepo.GetAllCourses()
}

func (s *CourseService) Selectacourse(id uint) (model.GetCoursedetail, error) {
	courses, err := s.CourseRepo.GetCourse(id)
	if err != nil {
		return model.GetCoursedetail{}, err
	}
	return courses, nil
}

func (s *CourseService) SelectSubLang(id uint) (model.GetSublang, error) {
	Sublang, err := s.CourseRepo.GetMaterialsBySubLangID(id)
	if err != nil {
		return model.GetSublang{}, err
	}
	return Sublang, nil
}

func (s *CourseService) GetGamification(sublangId uint, id uint) ([]model.Gamification, error) {
	gamification, err := s.CourseRepo.GetQuestionsBySublangID(sublangId, id)
	if err != nil {
		return []model.Gamification{}, err
	}
	return gamification, nil
}

func (s *CourseService) CheckAnswer(userID uuid.UUID, quest uint, option uint) (bool, error) {
	correctAnswer, err := s.CourseRepo.CheckCorrectAnswer(quest, option)
	if err != nil {
		return false, err
	}

	if correctAnswer {
		userAnswer := &entity.UserAnswerGami{
			UserID:     userID,
			QuestionID: quest,
			Answer:     option,
			Value:      true,
		}
		if err := s.CourseRepo.SaveUserAnswer(userAnswer); err != nil {
			return false, err
		}
	}

	return correctAnswer, nil
}

func (s *CourseService) GetnameCerification(name string, courseid uint) (model.Certification, error) {
	coursename, err := s.CourseRepo.GetnameCerification(courseid)
	if err != nil {
		return model.Certification{}, err
	}
	certification := model.Certification{
		Name:     name,
		Language: coursename,
	}
	return certification, nil
}
