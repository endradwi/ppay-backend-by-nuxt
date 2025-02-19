package models

import (
	"context"
	"fmt"
	"log"
	"test/lib"
	"time"

	"github.com/jackc/pgx/v5"
)

type order struct {
	Id          int
	Profile_id  int      `json:"profile_id" form:"profile_id" db:"profile_id"`
	Movie_id    int      `json:"movie_id" form:"movie_id" db:"movie_id"`
	Cinema_id   int      `json:"cinema_id" form:"cinema_id" db:"cinema_id"`
	Seat        []string `json:"seat" form:"seat[]"`
	Quantity    int      `json:"quantity" form:"quantity" db:"quantity"`
	TotalPrice  int      `json:"total_price" form:"total_price" db:"total_price"`
	Cinema_name string   `json:"cinema_name" form:"cinema_name" db:"cinema_name"`
	Location    string   `json:"location" form:"location" db:"location"`
}

type OrderBody struct {
	order
	Date string `json:"date" form:"date" `
	Time string `json:"time" form:"time" `
}

type OrderData struct {
	order
	Date time.Time `db:"date"`
	Time time.Time `db:"time"`
}

type Orders struct {
	Id         int    `json:"id"`
	Profile_Id int    `json:"profile_id" form:"profile_id"`
	Movie_Id   int    `json:"movie_id" form:"movie_id"`
	Tittle     string `json:"tittle" form:"tittle"`
	Genre      string `json:"genre" form:"genre"`
	Images     string `json:"image" form:"image"`
	Qty        int    `json:"qty" form:"qty"`
	Seat       string `json:"seat" form:"seat"`
	Cinema     string `json:"cinema" form:"cinema"`
	TotalPrice int    `json:"total_price" form:"total_price"`
}

type MoviesCinema struct {
	Movie       MoviesNoTag
	Cinema      string    `json:"cinema" form:"cinema"`
	Cinema_date time.Time `json:"cinema_date" form:"cinema_date"`
	Cinema_time time.Time `json:"cinema_time" form:"cinema_time"`
	Location    string    `json:"location" form:"location"`
	// Tag          string    `json:"tag" form:"tag"`
}
type ListAllCinema []MoviesCinema

type seat string

type SeatCinema struct {
	Id              int    `json:"id"`
	Tittle          string `json:"tittle" form:"tittle"`
	Genre           string `json:"genre" form:"genre"`
	Images          string `json:"image" form:"image"`
	Cinema          string `json:"cinema" form:"cinema"`
	Cinema_date     string `json:"cinema_date" form:"cinema_date"`
	Cinema_time     string `json:"cinema_time" form:"cinema_time"`
	Cinema_location string `json:"cinema_location" form:"cinema_location"`
	Price           int    `json:"price" from:"price"`
	Seats           []seat `json:"seat" from:"seat[]"`
}

type GetCinemaDTO struct {
	MovieId  int    `form:"movie_id"`
	Location string `form:"location"`
	Date     string `form:"date"`
	Time     string `form:"time"`
}
type GetCinema struct {
	Movie_Id int       `json:"movie_id" form:"movie_id"`
	Cinema   string    `json:"cinema" form:"cinema"`
	Location string    `json:"location" form:"location"`
	Date     time.Time `json:"date" form:"date"`
	Time     time.Time `json:"time" form:"time"`
}

type TotalSeatCinema struct {
	SeatCinema
	Total_Seat int
}

type ListCinema []MoviesCinema

type ListOrders []Orders
type GetCinemas []GetCinema

func OrderTicket(data OrderBody) OrderData {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var order OrderData
	cinema := struct {
		price int `db:"price"`
	}{}

	conn.QueryRow(context.Background(),
		`SELECT price FROM cinema WHERE id = $1`, data.Cinema_id).Scan(&cinema.price)

	totalPrice := cinema.price * len(data.Seat) * data.Quantity

	conn.QueryRow(context.Background(), `
	SELECT cinema.name, cinema_location.name_location
	FROM cinema
	JOIN cinema_location ON cinema_location.cinema_id = cinema.id
	WHERE cinema.id = $1`,
		data.Cinema_id).Scan(&order.Cinema_name, &order.Location)

	conn.QueryRow(context.Background(), `
	INSERT INTO orders (profile_id, movie_id, cinema_id, seat, date_order, qty, total_price) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING id, profile_id, movie_id, cinema_id, seat, date_order, qty, total_price`,
		data.Profile_id, data.Movie_id, data.Cinema_id, data.Seat, data.Date, data.Quantity, totalPrice).Scan(
		&order.Id, &order.Profile_id, &order.Movie_id, &order.Cinema_id, &order.Seat, &order.Date, &order.Quantity, &order.TotalPrice,
	)

	return order
}

func BookingCinema(paramId int, Location string) ListAllCinema {
	conn := lib.DB()
	defer conn.Close(context.Background())
	// var movie MoviesNoTag

	searchLoc := fmt.Sprintf("%%%s%%", Location)

	rows, err := conn.Query(context.Background(), `
	SELECT movie_schedules.movie_id, movies.tittle, movies.genre, 
movies.images, movies.synopsis, movies.author, 
movies.actors, movies.release_date, movies.duration, cinema.name,
cinema_date.name_date, cinema_time.name_time, cinema_location.name_location
	FROM movies
    JOIN movie_schedules On movie_id = movies.id
    JOIN cinema ON cinema.id = cinema_id
    JOIN cinema_date ON cinema_date.id = date_id
    JOIN cinema_time ON cinema_time.id = time_id
    JOIN cinema_location ON cinema_location.id = location_id 
    WHERE movie_schedules.movie_id = $1 AND cinema_location.name_location ILIKE $2
	`, paramId, searchLoc)
	if err != nil {
		fmt.Println(err)
	}
	cinema, err := pgx.CollectRows(rows, pgx.RowToStructByName[MoviesCinema])
	if err != nil {
		fmt.Println(err)
	}
	// .Scan(&movie.Id, &movie.Tittle, &movie.Genre, &movie.Image,
	// &movie.Synopsis, &movie.Author, &movie.Actors,
	// &movie.Release_date, &movie.Duration)
	return cinema
}

func FindCinema(input GetCinemaDTO) (GetCinema, error) {
	conn := lib.DB()
	fmt.Println("data conn=", conn)
	defer conn.Close(context.Background())

	var movie GetCinema
	// parsedDate, err := time.Parse("2006-01-02", input.Date)
	// if err != nil {
	// fmt.Println("Error parsing date:", err)
	// return
	// }
	// parsedTime, err := time.Parse("15:04:05", input.Time)
	// if err != nil {
	// fmt.Println("Error parsing time:", err)
	// return
	// }

	log.Println("data time=", input.Time)
	log.Println("data date=", input.Date)
	log.Println("data location=", input.Location)
	// log.Println("data movie_id=", input.MovieId)

	// log.Println(parsedDate)
	// log.Println(parsedTime)

	err := conn.QueryRow(context.Background(), `
	SELECT movies.id, cinema.name, cinema_location.name_location, cinema_date.name_date, 
		cinema_time.name_time
	FROM movie_schedules
	JOIn movies ON movie_id = movies.id
	JOIN cinema ON cinema_id = cinema.id
	JOIN cinema_location ON location_id = cinema_location.id
	JOIN cinema_date ON date_id = cinema_date.id
	JOIN cinema_time ON time_id = cinema_time.id
	WHERE movie_id = $1 
    AND cinema_location.name_location = $2 
	AND cinema_date.name_date = $3 
    AND cinema_time.name_time = $4
	`, input.MovieId, input.Location, input.Date, input.Time).Scan(&movie.Movie_Id, &movie.Cinema, &movie.Location,
		&movie.Date, &movie.Time)
	fmt.Println(err)
	log.Println("data time=", movie.Time)
	log.Println("data date=", movie.Date)
	log.Println("data location=", movie.Location)
	log.Println("data movie_id=", movie.Movie_Id)
	// fmt.Println("Query executed successfully, movie:", movie)

	// fmt.Println("data = )
	return movie, nil
}
