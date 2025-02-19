package models

import (
	"context"
	"log"
	"test/lib"
	"time"
)

type Payment struct {
	No_Rekening   string    `json:"no_rekening" form:"no_rekening"`
	Total_Payment int       `json:"total_payment" form:"total_payment"`
	Limit_Payment time.Time `json:"limit_payment" form:"limit_paymnet"`
}
type PaymentData struct {
	Payment
	Order_id   int       `json:"order_id" form:"order_id"`
	Payment_id int       `json:"payment_id" form:"payment_id"`
	Date       time.Time `json:"date"`
}

type PaymentPaidStatus struct {
	Order_Id         int    `json:"order_id" form:"order_id"`
	Payment_Method   string `json:"payment_method" form:"payment_method"`
	Payment_Rekening string `json:"payment_rekening" form:"payment_rekening"`
	Total_Price      int    `json:"total_price" form:"total_price"`
}

type PaymentPaidOrders struct {
	Order_Id   int    `json:"order_id" form:"order_id"`
	Movie_Name string `json:"movie_name" form:"movie_name"`
	Seat       string `json:"seat" form:"seat"`
}

type PaymentInfo struct {
	PaymentPaidOrders
	Payment_Method   string `json:"payment_method" form:"payment_method"`
	Payment_Rekening string `json:"payment_rekening" form:"payment_rekening"`
	Total_Price      int    `json:"total_price" form:"total_price"`
}

func PaymentMethod(data PaymentData) Payment {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var order Payment

	conn.QueryRow(context.Background(), `
		SELECT total_price FROM orders
		WHERE id = $1
	`, data.Order_id).Scan(&order.Total_Payment)

	conn.QueryRow(context.Background(), `
		SELECT rekening FROM payment_method
		WHERE id = $1
	`, data.Payment_id).Scan(&order.No_Rekening)

	limit_payment := time.Now().Add(24 * time.Hour)

	conn.QueryRow(context.Background(), `
		UPDATE orders
		SET payment_id = $1, expired_payment = $2, status_id = 1
		WHERE id = $3
		RETURNING total_price, expired_payment
	`, data.Payment_id, limit_payment, data.Order_id).Scan(
		&order.Total_Payment, &order.Limit_Payment,
	)

	return order
}

func PaymentPaid(data PaymentInfo) PaymentPaidOrders {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var order PaymentPaidOrders

	conn.QueryRow(context.Background(), `
		SELECT orders.id, payment_method.name, payment_method.rekening, orders.total_price FROM orders
		JOIN payment_method ON payment_method.id = orders.payment_id
		WHERE orders.id = $1 AND payment_method.name = $2 AND payment_method.rekening = $3 AND orders.total_price = $4
	`, data.Order_Id).Scan(
		&order.Order_Id)

	log.Println("data id =", order.Order_Id)
	log.Println("data id =", data.Movie_Name)

	conn.QueryRow(context.Background(), `
		UPDATE orders
		status_id = 2
		WHERE id = $1
	`, data.Order_Id)

	return order
}
