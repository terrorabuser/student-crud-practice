package repository

import (
	"gorm.io/gorm"
	"student-crud/internal/model"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) Create(student *model.Student) error {
	return r.db.Create(student).Error
}

func (r *StudentRepository) GetByID(id uint) (*model.Student, error) {
	var student model.Student
	if err := r.db.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) GetAll(limit, offset int) ([]model.Student, error) {
	var students []model.Student
	if err := r.db.Limit(limit).Offset(offset).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *StudentRepository) Update(student *model.UpdateStudent) error {
	return r.db.Table("students").
		Where("id = ?", student.ID).
		Updates(student).Error
}

func (r *StudentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Student{}, id).Error
}
