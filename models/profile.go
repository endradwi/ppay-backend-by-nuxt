package models

import (
	"context"
	"fmt"
	"log"
	"test/lib"

	"github.com/jackc/pgx/v5"
)

type Profile struct {
	Id           int    `json:"id" form:"id"`
	First_Name   string `json:"first_name" form:"first_name" example:"Salah"`
	Last_Name    string `json:"last_name" form:"last_name" example:"Alaudin"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Image        string `json:"image" form:"-" example:"salah.jpg"`
	Email        string `json:"email" form:"email" example:"salah@mail.com"`
	Password     string `json:"password" form:"password" example:"Salah1!"`
}

type PointProfile struct {
	Profile
	Point int
}

type RelationProfile struct {
	Id           int    `json:"id"`
	First_Name   string `json:"first_name" form:"first_name"`
	Last_Name    string `json:"last_name" form:"last_name"`
	Phone_Number string `json:"phone_number" form:"phone_number"`
	Image        string `json:"image" form:"image"`
	User_Id      int    `json:"user_id" form:"user_id"`
}

type TestUsers struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ListProfile []Profile

func FindOneProfile(paramId int) Profile {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var profile Profile

	conn.QueryRow(context.Background(), `
	SELECT users.id, profile.first_name, profile.last_name,
       profile.image, profile.phone_number, users.email, users.password
	FROM users
	JOIN profile ON profile.user_id = users.id 
	WHERE users.id = $1
	`, paramId).Scan(&profile.Id, &profile.First_Name,
		&profile.Last_Name, &profile.Image, &profile.Phone_number, &profile.Email, &profile.Password)

	hash := lib.CreateHash(profile.Password)
	if profile.Password != "" {
		profile.Password = hash
	}
	return profile
}

func FindProfile(paramId int) PointProfile {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var profile PointProfile

	err := conn.QueryRow(context.Background(), `
	SELECT users.id, profile.first_name, profile.last_name,
       profile.image,profile.phone_number, profile.point, users.email, users.password
	FROM users
	JOIN profile ON profile.user_id = users.id 
	WHERE users.id = $1
	`, paramId).Scan(&profile.Id, &profile.First_Name,
		&profile.Last_Name, &profile.Image, &profile.Phone_number, &profile.Point, &profile.Email, &profile.Password)
	log.Println(err)

	hash := lib.CreateHash(profile.Password)
	if profile.Password != "" {
		profile.Password = hash
	}
	return profile
}

func AddProfile(profile RelationProfile) RelationProfile {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var new RelationProfile
	conn.QueryRow(context.Background(),
		`INSERT INTO profile (first_name, last_name, image, user_id, phone_number)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, first_name, last_name, image, user_id, phone_number`,
		profile.First_Name, profile.Last_Name, profile.Image, profile.User_Id, profile.Phone_Number).Scan(
		&new.Id, &new.First_Name, &new.Last_Name, &new.Image, &new.User_Id, &new.Phone_Number,
	)

	return new
}

func UpdatedProfile(profile Profile, userId int) Profile {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var update Profile

	row := conn.QueryRow(context.Background(), `
		UPDATE profile
		SET first_name=$1,
			last_name=$2,
			image=$3,
			phone_number=$4
		FROM users
		WHERE users.id = $5
		AND profile.user_id = users.id
		RETURNING users.id, profile.first_name, profile.last_name, profile.image, profile.phone_number
	`, profile.First_Name, profile.Last_Name, profile.Image, profile.Phone_number, userId)

	if err := row.Scan(&update.Id, &update.First_Name, &update.Last_Name, &update.Image, &update.Phone_number); err != nil {
		fmt.Println("Error updating profile:", err)
	}

	row = conn.QueryRow(context.Background(), `
		UPDATE users
		SET email=$1, password=$2
		WHERE id=$3
		RETURNING users.email, users.password
	`, profile.Email, profile.Password, userId)
	if err := row.Scan(&update.Email, &update.Password); err != nil {
		fmt.Println("Error updating user:", err)
	}

	return update
}

func DeleteProfile(id int) Profile {
	conn := lib.DB()
	defer conn.Close(context.Background())

	fmt.Println("data connection =", conn)

	var profileDelete Profile

	text := conn.QueryRow(context.Background(), `
	SELECT profile.user_id, profile.first_name, profile.last_name,
	profile.image, profile.phone_number,
	users.email, users.password
	FROM profile
	JOIN users ON users.id = profile.user_id
	WHERE profile.user_id =$1
	`, id).Scan(&profileDelete.Id,
		&profileDelete.First_Name,
		&profileDelete.Last_Name,
		&profileDelete.Image,
		&profileDelete.Phone_number,
		&profileDelete.Email,
		&profileDelete.Password,
	)
	fmt.Println("data select =", text)

	conn.QueryRow(context.Background(), `
	DELETE FROM users 
	USING profile
	WHERE users.id = profile.user_id AND profile_id = $1
	`, id)

	conn.QueryRow(context.Background(), `
	DELETE FROM profile WHERE id = $1
	`, id)

	return profileDelete
}

func FindAllProfile(search string) ListProfile {
	conn := lib.DB()
	defer conn.Close(context.Background())
	rows, _ := conn.Query(context.Background(), `
	SElECT profile.id, profile.first_name,
	profile.last_name, profile.image, users.email
	FROM profile
	JOIN users ON users(id) = profile.user_id
	WHERE profile.first_name = $1 
	`, search)
	profile, _ := pgx.CollectRows(rows, pgx.RowToStructByName[Profile])
	return profile
}

func ConsepGetUserJoin() {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var update Profile

	conn.QueryRow(context.Background(), `
	SElECT profile.id, profile.first_name,
	profile.last_name, profile.image, users.email, users.password
	FROM profile
	JOIN users ON users.id = profile.user_id 
	`).Scan(
		&update.Id, &update.First_Name, &update.Last_Name,
		&update.Image, &update.Email, &update.Password,
	)

}
