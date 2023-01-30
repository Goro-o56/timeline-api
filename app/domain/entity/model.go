package entity

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var err error
	d.Time, err = time.Parse("\"2006-01-02\"", string(b))
	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format("2006-01-02"))), nil
}

func (d *Date) Scan(value any) error {
	d.Time = value.(time.Time)
	return nil
}

func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}

type Persona struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Content   string    `gorm:"type:varchar(255);not null" json:"persona"`
	CreatedAt time.Time `gorm:"type:datetime;not null" sql: "default:current_timestamp" json:"-"`
}

type Timeline struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Content   []Persona `gorm:"type:varchar(255);not null" json:"persona"`
	CreatedAt time.Time `gorm:"type:datetime;not null" sql: "default:current_timestamp" json:"-"`
}
