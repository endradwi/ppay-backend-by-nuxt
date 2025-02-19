package models

import (
	"context"
	"fmt"
	"test/lib"
)

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" binding:"email"`
	Password string `json:"password" form:"password" binding:"min=6,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"`
}
type UserAdmin struct {
	Id           int    `json:"id" form:"id"`
	First_Name   string `json:"first_name" form:"first_name" example:"Salah"`
	Last_Name    string `json:"last_name" form:"last_name" example:"Alaudin"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Image        string `jsoon:"image" form:"image" example:"salah.jpg"`
	Email        string `json:"email" form:"email" example:"salah@mail.com"`
	Password     string `json:"password" form:"password" example:"Salah1!"`
}

type ListUser []Users

func FindOneUserByEmail(email string) Users {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var user Users
	conn.QueryRow(context.Background(), `
	SELECT id, email, password 
	FROM users WHERE email = $1
	`, email).Scan(&user.Id, &user.Email, &user.Password)
	return user
}

func InsertUser(user Users) Users {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var new Users

	conn.QueryRow(context.Background(), `
	INSERT INTO users(email, password) VALUES
	($1, $2)
	RETURNING id, email, password
	`, user.Email, user.Password).Scan(&new.Id,
		&new.Email, &new.Password)
	return new
}

func AddUsers(profile UserAdmin) UserAdmin {
	conn := lib.DB()
	defer conn.Close(context.Background())

	fmt.Println("data connection =", conn)

	var profileAdd UserAdmin
	var user_id int

	// Query pertama: Insert data ke tabel users dan dapatkan user_id
	err := conn.QueryRow(context.Background(), `
		INSERT INTO users (email, password) 
		VALUES ($1, $2) RETURNING id
	`, profile.Email, profile.Password).Scan(&user_id)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return profileAdd
	}

	fmt.Println("user_id =", user_id)

	// Query kedua: Insert data ke tabel profile dengan user_id yang didapat dari query pertama
	err = conn.QueryRow(context.Background(), `
		INSERT INTO profile (first_name, last_name, image, phone_number, user_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, first_name, last_name, phone_number, image
	`, profile.First_Name, profile.Last_Name, profile.Image, profile.Phone_number, user_id).Scan(
		&profileAdd.Id, &profileAdd.First_Name, &profileAdd.Last_Name, &profileAdd.Phone_number, &profileAdd.Image,
	)
	if err != nil {
		fmt.Println("Error inserting profile:", err)
		return profileAdd
	}

	return profileAdd
}
