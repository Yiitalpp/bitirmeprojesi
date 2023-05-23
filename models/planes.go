package models

import (
	"gorm.io/gorm"
)

type Plane struct {
	gorm.Model
	ID         uint
	FirmName   string
	SeatNumber string
}

// create a Plane
func CreatePlane(db *gorm.DB, Team *Plane) (err error) {
	err = db.Create(Team).Error
	if err != nil {
		return err
	}
	return nil
}

// get Planes
func GetPlanes(db *gorm.DB, Plane *[]Plane) (err error) {
	err = db.Find(Plane).Error
	if err != nil {
		return err
	}
	return nil
}

// get Plane by id
func GetPlane(db *gorm.DB, Plane *Plane, id string) (err error) {
	err = db.Where("id = ?", id).First(Plane).Error
	if err != nil {
		return err
	}
	return nil
}

// update a Plane
func UpdatePlane(db *gorm.DB, Plane *Plane, id string) (err error) {
	err = db.Model(Plane).Where("id = ?", id).Updates(map[string]interface{}{"plane_name": Plane.FirmName, "seat_number": Plane.SeatNumber}).Error
	if err != nil {
		return err
	}
	return nil
}

// delete Plane
func DeletePlane(db *gorm.DB, Plane *Plane, id string) (err error) {
	err = db.Where("id = ?", id).Delete(Plane).Error
	if err != nil {
		return err
	}
	return nil
}
