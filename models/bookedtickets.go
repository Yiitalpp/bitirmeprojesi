package models

import (
	"gorm.io/gorm"
)

type BTicket struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Ticket   Ticket `gorm:"foreignKey:TicketID"`
	TicketID uint
	User     User `gorm:"foreignKey:UserID"`
	UserID   uint
}

// create a Plane
func CreateBTicket(db *gorm.DB, BTicket *BTicket) (err error) {
	err = db.Create(BTicket).Error
	if err != nil {
		return err
	}
	return nil
}

// get Planes
func GetBTickets(db *gorm.DB, BTicket *[]BTicket) (err error) {
	err = db.Find(BTicket).Error
	if err != nil {
		return err
	}
	return nil
}

// get Plane by id
func GetBTicket(db *gorm.DB, BTicket *BTicket, id string) (err error) {
	err = db.Where("id = ?", id).First(BTicket).Error
	if err != nil {
		return err
	}
	return nil
}

// update a Plane
func UpdateBTicket(db *gorm.DB, BTicket *BTicket, id string) (err error) {
	err = db.Model(BTicket).Where("id = ?", id).Updates(map[string]interface{}{"user_id": BTicket.UserID, "ticket_id": BTicket.TicketID}).Error
	if err != nil {
		return err
	}
	return nil
}

// delete Plane
func DeleteBTicket(db *gorm.DB, BTicket *BTicket, id string) (err error) {
	err = db.Where("id = ?", id).Delete(BTicket).Error
	if err != nil {
		return err
	}
	return nil
}
