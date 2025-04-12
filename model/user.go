package model

import "time"

type User struct {
    ID                uint      `gorm:"primaryKey"`
    Name              string
    Email             string    `gorm:"unique"`
    CountryCode       string
    Phone             string
    Password          string
    CreatedAt         time.Time
    UpdatedAt         time.Time
    DeletedAt         *time.Time
}
