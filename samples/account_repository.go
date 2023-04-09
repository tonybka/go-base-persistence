package account

import (
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db}
}

// Create creates new account
func (repo *AccountRepository) Create(dataModel AccountModel) error {
	if result := repo.db.Create(&dataModel); result.Error != nil {
		return result.Error
	}
	return nil
}

// FindById query account by it's identity
func (repo *AccountRepository) FindById(id uint) (AccountModel, error) {
	var dataModel AccountModel

	if result := repo.db.Where("id = ?", id).First(&dataModel); result.Error != nil {
		return AccountModel{}, result.Error
	}

	return dataModel, nil
}

// GetAll returns all accounts in the table
func (repo *AccountRepository) GetAll() ([]AccountModel, error) {
	var dataModels []AccountModel

	if result := repo.db.Find(&dataModels); result.Error != nil {
		return nil, result.Error
	}

	return dataModels, nil
}
