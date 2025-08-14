// Package repositories provides implementations for data persistence and retrieval
// related to contact entities in the API Contact Form application.
//
// It defines the ContactRepository interface and its GORM-based implementation
// for performing CRUD operations on contact records in the database.
package repositories

import (
	"api-contact-form/models"
	"time"

	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(contact *models.Contact) error
	FindAll() ([]models.Contact, error)
	FindByID(id uint) (*models.Contact, error)
	Update(contact *models.Contact) error
	Delete(contact *models.Contact) error
}

// contactRepository is the GORM-based implementation of ContactRepository.
type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db}
}

func (r *contactRepository) Create(contact *models.Contact) error {
	return r.db.Create(contact).Error
}

func (r *contactRepository) FindAll() ([]models.Contact, error) {
	var contacts []models.Contact
	err := r.db.Where("deleted_at IS NULL").Find(&contacts).Error
	return contacts, err
}

func (r *contactRepository) FindByID(id uint) (*models.Contact, error) {
	var contact models.Contact
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&contact).Error
	return &contact, err
}

func (r *contactRepository) Update(contact *models.Contact) error {
	return r.db.Save(contact).Error
}

func (r *contactRepository) Delete(contact *models.Contact) error {
	now := time.Now()
	return r.db.Model(contact).Select("deleted_at").Updates(models.Contact{
		DeletedAt: &now,
	}).Error
}