package main

import (
	"fmt"

	"github.com/lemon57/share-moment/models"
)

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)

	//db, err := sql.Open("pgx", "postgres://bjj:sharemoment@localhost:5432/sharemoment?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	us := models.UserService{
		DB: db,
	}
	user, err := us.Create("roman02@test.com", "password")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS users (
	// 	id SERIAL PRIMARY KEY,
	// 	name TEXT,
	// 	email TEXT NOT NULL
	//   );

	//   CREATE TABLE IF NOT EXISTS orders (
	// 	id SERIAL PRIMARY KEY,
	// 	user_id INT NOT NULL,
	// 	amount INT,
	// 	description TEXT
	//   );
	// `)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Tables created.")

	// name := "New"
	// email := "new@test.com"

	// row := db.QueryRow(`
	// 	INSERT INTO users(name, email)
	// 	VALUES($1, $2) RETURNING id;`, name, email)

	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("User created. Id =", id)

	// id := 1
	// row := db.QueryRow(`
	// 	SELECT name, email
	// 	FROM users
	// 	WHERE id = $1;`, id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User: name = %s, email = %s", name, email)

	// userID := 1
	// for i := 1; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)
	// 	_, err := db.Exec(`
	// 		INSERT INTO orders(user_id, amount, description)
	// 		VALUES($1, $2, $3);`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Fake orders created!")

	// type Order struct {
	// 	ID          int
	// 	UserID      int
	// 	Amount      int
	// 	Description string
	// }

	// var orders []Order

	// userID := 1
	// rows, err := db.Query(`
	// 	SELECT id, amount, description
	// 	FROM orders
	// 	WHERE user_id = $1;`, userID)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var order Order
	// 	order.UserID = userID
	// 	err := rows.Scan(&order.ID, &order.Amount, &order.Description)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	orders = append(orders, order)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Orders:", orders)
}
