package models

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	ID            uint
	Plane         Plane `gorm:"foreignKey:PlaneID"` // Relationship with Plane model
	PlaneID       uint  // Foreign key referencing Plane ID
	From          string
	To            string
	DepartureDate string
	ReturnDate    string
	DHour         string
	RHour         string
	NofSeats      string
	Price         string
}

// create a Plane
func CreateTicket(db *gorm.DB, Ticket *Ticket) (err error) {
	err = db.Create(Ticket).Error
	if err != nil {
		return err
	}
	return nil
}

// get Planes
func GetTickets(db *gorm.DB, Ticket *[]Ticket) (err error) {
	err = db.Find(Ticket).Error
	if err != nil {
		return err
	}
	return nil
}

// get Planes
func FilterTickets(db *gorm.DB, Ticket *[]Ticket) (err error) {
	err = db.Find(Ticket).Error
	if err != nil {
		return err
	}
	return nil
}

// get Plane by id
func GetTicket(db *gorm.DB, Ticket *Ticket, id string) (err error) {
	err = db.Where("id = ?", id).First(Ticket).Error
	if err != nil {
		return err
	}
	return nil
}

// update a Plane
func UpdateTicket(db *gorm.DB, Ticket *Ticket, id string) (err error) {
	err = db.Model(Ticket).Where("id = ?", id).Updates(map[string]interface{}{"plane_ID": Ticket.PlaneID, "From": Ticket.From, "To": Ticket.To, "Departure_date": Ticket.DepartureDate, "Return_date": Ticket.ReturnDate, "DHour": Ticket.DHour, "RHour": Ticket.RHour, "NofSeats": Ticket.NofSeats, "price": Ticket.Price}).Error
	if err != nil {
		return err
	}
	return nil
}

// delete Plane
func DeleteTicket(db *gorm.DB, Ticket *Ticket, id string) (err error) {
	err = db.Where("id = ?", id).Delete(Ticket).Error
	if err != nil {
		return err
	}
	return nil
}
