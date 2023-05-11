package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    ID       uint   `gorm:"primaryKey;autoIncrement;not null"`
    Username string `gorm:"uniqueIndex;not null"`
    Email    string `gorm:"uniqueIndex;not null"`
    Password string `gorm:"not null"`
    Photos   []Photo
}

type Photo struct {
    gorm.Model
    ID       uint   `gorm:"primaryKey;autoIncrement;not null"`
    Title    string `gorm:"not null"`
    Caption  string `gorm:"not null"`
    PhotoURL string `gorm:"not null"`
    UserID   uint   `gorm:"not null"`
    User     User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
