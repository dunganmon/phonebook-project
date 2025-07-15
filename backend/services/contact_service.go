package services

import (
	"TODO/databases"
	"TODO/dtos"
	"TODO/models"
	"errors"
)

func AddContact(contactReq dtos.ContactRequest) (*models.Contact, error) {
	contact := models.Contact{
		FullName: contactReq.FullName,
		Phone:    contactReq.Phone,
		Email:    contactReq.Email,
		Address:  contactReq.Address,
		UserId:   contactReq.UserID,
		UpdateAt: contactReq.UpdateAt,
	}
	if err := databases.DB.Create(&contact).Error; err != nil {
		return nil, errors.New("Không thể tạo")
	}
	return &contact, nil
}

func GetContactsByUserId(userId int) ([]models.Contact, error) {
	contacts := []models.Contact{}
	if err := databases.DB.Where("user_id = ?", userId).Find(&contacts).Error; err != nil {
		return nil, errors.New("Không thể lấy danh sách contact")
	}

	return contacts, nil
}
func DeleteContact(contactId int) error {
	contact := models.Contact{}
	if err := databases.DB.Where("id = ?", contactId).Delete(&contact).Error; err != nil {
		return errors.New("Khong the tao")
	}
	return nil
}
func EditContact(contactId int, contactReq dtos.ContactRequest) (*models.Contact, error) {
	contact := models.Contact{}
	if err := databases.DB.Where("id = ?", contactId).First(&contact).Error; err != nil {
		return nil, errors.New("Khong the tao")
	}
	contact.FullName = contactReq.FullName
	contact.Phone = contactReq.Phone
	contact.Email = contactReq.Email
	contact.Address = contactReq.Address
	if err := databases.DB.Save(&contact).Error; err != nil {
		return nil, errors.New("Khong the tao")
	}
	return &contact, nil
}
