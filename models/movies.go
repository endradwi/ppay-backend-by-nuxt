package models

import (
	"context"
	"fmt"
	"log"
	"test/lib"
	"time"

	"github.com/jackc/pgx/v5"
)

type MoviesbyTag struct {
	Id       int    `json:"id"`
	Tittle   string `json:"tittle" form:"tittle" example:"Spiderman"`
	Genre    string `json:"genre" form:"genre" example:"Action"`
	Synopsis string `json:"synopsis" form:"synopsis" example:"film action universal"`
	Author   string `json:"author" form:"author"`
	Actors   string `json:"actors" form:"actors"`
	// Release_date time.Time `json:"release_date" form:"release_date"`
	// Duration     time.Time `json:"duration" form:"duration"`
	Tag   string `json:"tag" form:"tag"`
	Image string `json:"image" form:"image" example:"Spiderman.jpg"`
}
type MoviesNoTag struct {
	Id           int       `json:"id"`
	Tittle       string    `json:"tittle" form:"tittle" example:"Spiderman"`
	Genre        string    `json:"genre" form:"genre" example:"Action"`
	Image        string    `json:"image" form:"image" example:"Spiderman.jpg"`
	Synopsis     string    `json:"synopsis" form:"synopsis" example:"film action universal"`
	Author       string    `json:"author" form:"author"`
	Actors       string    `json:"actors" form:"actors"`
	Release_date time.Time `json:"release_date" form:"release_date"`
	Duration     time.Time `json:"duration" form:"duration"`
	// Tag          string    `json:"tag" form:"tag"`
}

type GetAllMovie struct {
	Id           int       `json:"id"`
	Tittle       string    `json:"tittle" form:"tittle" example:"Spiderman"`
	Genre        string    `json:"genre" form:"genre" example:"Action"`
	Images       string    `json:"image" form:"image" example:"Spiderman.jpg"`
	Release_date time.Time `json:"release_date" form:"release_date"`
}

type Movie_body struct {
	MoviesbyTag
	Release_date string `json:"release_date" form:"release_date"`
	Duration     string `json:"duration" form:"duration" `
}

type Movie_Data struct {
	MoviesbyTag
	Release_date time.Time `db:"release_date"`
	Duration     time.Time `db:"duration"`
}

type ListMovie []MoviesbyTag
type ListAllMovie []GetAllMovie

func FindAllMovie(page int, limit int, search string, sort string) ListAllMovie {
	conn := lib.DB()
	defer conn.Close(context.Background())

	offset := (page - 1) * limit

	searching := fmt.Sprintf("%%%s%%", search)

	query := fmt.Sprintf(`
	SELECT id, tittle, genre ,images, release_date
	FROM movies
	WHERE tittle ILIKE $1
 	ORDER BY id %s
 	LIMIT $2 OFFSET $3
 `, sort)

	rows, _ := conn.Query(context.Background(), query, searching, limit, offset)

	movie, _ := pgx.CollectRows(rows, pgx.RowToStructByName[GetAllMovie])

	return movie
}
func CountData(search string) int {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var count int
	search = fmt.Sprintf("%%%s%%", search)

	conn.QueryRow(context.Background(), `
	SELECT COUNT(id)
	FROM movies
	WHERE tittle ILIKE $1
	`, search).Scan((&count))
	return count
}

func FindOneMovie(paramId int) MoviesNoTag {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var movie MoviesNoTag

	conn.QueryRow(context.Background(), `
	SELECT id, tittle, genre, images, synopsis, 
	author, actors, release_date, duration
	FROM movies 
    WHERE id = $1
	`, paramId).Scan(&movie.Id, &movie.Tittle, &movie.Genre, &movie.Image,
		&movie.Synopsis, &movie.Author, &movie.Actors,
		&movie.Release_date, &movie.Duration)
	return movie
}

func CountMovie(search string) int {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var total int

	search = fmt.Sprintf("%%%s%%", search)

	conn.QueryRow(context.Background(), `
	SELECT COUNT(id) FROM movie WHERE title ILIKE $1
	`, search).Scan(&total)

	return total
}

func InsertMovie(data Movie_body) (Movie_Data, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movieInsert Movie_Data

	if data.Release_date == "" {
		return Movie_Data{}, fmt.Errorf("release date is required")
	}

	movieDate, _ := time.Parse(time.DateOnly, data.Release_date)

	log.Println("movieDate =", movieDate)
	movieDuration, _ := time.Parse(time.TimeOnly, data.Duration)

	log.Println("data duration =", movieDuration)

	text := conn.QueryRow(context.Background(), `
	INSERT INTO movies (tittle, genre, images, synopsis,
	author, actors, release_date, duration, tag) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id, tittle, genre, images, synopsis,
	author, actors, release_date, duration, tag
	`, data.Tittle, data.Genre, data.Image, data.Synopsis,
		data.Author, data.Actors, movieDate, movieDuration, data.Tag).
		Scan(
			&movieInsert.Id,
			&movieInsert.Tittle,
			&movieInsert.Genre,
			&movieInsert.Image,
			&movieInsert.Synopsis,
			&movieInsert.Author,
			&movieInsert.Actors,
			&movieInsert.Release_date,
			&movieInsert.Duration,
			&movieInsert.Tag,
		)
	log.Println("data text =", text)
	return movieInsert, nil
}

func UpdateMovie(movie Movie_body) Movie_Data {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movieUpdate Movie_Data

	// Check if Release_date is empty
	if movie.Release_date == "" {
		return Movie_Data{}
	}

	movieDate, err := time.Parse(time.DateOnly, movie.Release_date)
	if err != nil {
		return Movie_Data{}
	}

	log.Println(movie)

	movieDuration, err := time.Parse(time.TimeOnly, movie.Duration)
	if err != nil {
		return Movie_Data{}
	}

	conn.QueryRow(context.Background(), `
		UPDATE movie
		SET title = $1, genre = $2, images = $3, synopsis = $5,
		author = $6, actors = $7, release_date = $8, duration = $9,
		tag = $10  
		WHERE id = $11
		RETURNING id, title, genre, images, synopsis, author, actors, release_date, duration, tag  
	`, movie.Tittle, movie.Genre, movie.Image, movie.Synopsis,
		movie.Author, movie.Actors, movieDate, movieDuration, movie.Tag, movie.Id).Scan(
		&movieUpdate.Id,
		&movieUpdate.Tittle,
		&movieUpdate.Genre,
		&movieUpdate.Image,
		&movieUpdate.Synopsis,
		&movieUpdate.Author,
		&movieUpdate.Actors,
		&movieUpdate.Release_date,
		&movieUpdate.Duration,
		&movieUpdate.Tag,
	)

	return movieUpdate
}

func DeleteMovie(iddb int) MoviesbyTag {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movieDelete MoviesbyTag

	conn.QueryRow(context.Background(), `
	DELETE FROM movie WHERE id = $1
	RETURNING  id, title, genre, images, synopsis,
	author, actors, release_date, duration, tag
	`, iddb).Scan(
		&movieDelete.Id,
		&movieDelete.Tittle,
		&movieDelete.Genre,
		&movieDelete.Image,
		&movieDelete.Synopsis,
		&movieDelete.Author,
		&movieDelete.Actors,
		// &movieDelete.Release_date,
		// &movieDelete.Duration,
		&movieDelete.Tag,
	)

	return movieDelete
}
