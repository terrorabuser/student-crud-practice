package service

import (
	"student-crud/internal/model"
	"student-crud/internal/repository"
)

type StudentService struct {
	repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) CreateStudent(req *model.Student) error {
	err := s.repo.Create(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *StudentService) GetStudentByID(id uint) (*model.Student, error) {
	return s.repo.GetByID(id)
}

func (s *StudentService) GetAllStudents(limit, offset int) ([]model.Student, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *StudentService) UpdateStudent(req *model.UpdateStudent) error {
	return s.repo.Update(req)
}

func (s *StudentService) DeleteStudent(id uint) error {
	return s.repo.Delete(id)
}
