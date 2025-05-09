package repository

import (
	"crud-service/internal/model"

	"gorm.io/gorm"
)

type NoteRepository interface {
	Create(note *model.Note) error
	GetAll(userId uint) ([]model.Note, error)
	GetByID(id uint) (*model.Note, error)
	Update(note *model.Note) error
	Delete(note *model.Note) error
}
type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{
		db: db,
	}
}

func (r *noteRepository) Create(note *model.Note) error {
	return r.db.Create(note).Error
}
func (r *noteRepository) GetAll(userId uint) ([]model.Note, error) {
	var notes []model.Note
	err := r.db.Where(&model.Note{UserID: userId}).Find(&notes).Error
	if err != nil {
		return nil, err
	}
	return notes, nil
}
func (r *noteRepository) GetByID(id uint) (*model.Note, error) {
	var note model.Note
	err := r.db.First(&note, id).Error
	if err != nil {
		return nil, err
	}
	return &note, nil
}
func (r *noteRepository) Update(note *model.Note) error {
	return r.db.Save(note).Error
}
func (r *noteRepository) Delete(note *model.Note) error {
	return r.db.Delete(note).Error
}
