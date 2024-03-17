package service

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/internal/repository"
)

type Scourse interface {
	CreateCourse(courseData *entity.Course) error
	GetAllCourses() ([]entity.Course, error)
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

func (s *CourseService) GetAllCourses() ([]entity.Course, error) {
	return s.CourseRepo.GetAllCourses()
}
