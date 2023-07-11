package models

import "time"

type Blogpost struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// CREATE DATABASE blogpost;

// USE blogpost;

// CREATE TABLE blogpost (
//     ID int NOT NULL AUTO_INCREMENT,
//     Title varchar(255),
//     Content varchar(255),
//     Created_at DATETIME,
// 	PRIMARY KEY (ID)
// );
