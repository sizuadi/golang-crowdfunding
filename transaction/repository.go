package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignID(CampaignID int) ([]Transaction, error)
	GetByUserID(UserID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignID(CampaignID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", CampaignID).Order("id DESC, created_at DESC").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByUserID(UserID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", UserID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
