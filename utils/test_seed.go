package main

import (
	"context"
	"log"

	"github.com/Dmitriy-M1319/crystal-golang/config"
	"github.com/Dmitriy-M1319/crystal-golang/internal"
)

func main() {
	err := config.LoadSettings(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	settings := config.GetSettings()
	baseDb, err := internal.NewConnection(settings.BaseDB["ip"], settings.BaseDB["port"], settings.BaseDB["user"],
		settings.BaseDB["password"], settings.BaseDB["database"])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer internal.Close(baseDb)

	// Products
	createProdQuery := `CREATE TABLE IF NOT EXISTS products (
	id serial primary key,
	product_name varchar(255),
	company varchar(255),
	client_price decimal(10, 2),
	purchase_price decimal(10, 2),
	count integer
	)`

	ctx := context.Background()

	tx, err := baseDb.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("failed to begin transaction for products")
	}

	_, err = tx.Exec(createProdQuery)
	if err != nil {
		log.Fatalf("failed to create products table: %s", err.Error())
		tx.Rollback()
	}

	insertQuery := `INSERT INTO products (product_name, company, client_price, purchase_price, count) VALUES
		('AquaPure Filter', 'AquaPure', 1500, 1000, 50),
		('Brita Filter', 'Brita', 1200, 800, 30),
		('Pur Filter', 'Pur', 1300, 900, 40),
		('Whirlpool Filter', 'Whirlpool', 1600, 1100, 25),
		('GE Filter', 'GE', 1400, 950, 35),
		('Culligan Filter', 'Culligan', 1700, 1200, 20),
		('APEC Filter', 'APEC', 1800, 1300, 15),
		('iSpring Filter', 'iSpring', 1900, 1400, 10),
		('Home Master Filter', 'Home Master', 2000, 1500, 5),
		('3M Filter', '3M', 1100, 700, 45)`

	_, err = tx.Exec(insertQuery)
	if err != nil {
		log.Fatalf("failed to insert products into table: %s", err.Error())
		tx.Rollback()
	}

	tx.Commit()

	// Users
	createUserQuery := `CREATE TABLE IF NOT EXISTS users (
		id serial primary key,
		name varchar(255),
		surname varchar(255),
		email varchar(255),
		phone_number varchar(15),
		password varchar(255),
		is_admin boolean
		)`

	tx, err = baseDb.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("failed to begin transaction for users")
	}

	_, err = tx.Exec(createUserQuery)
	if err != nil {
		log.Fatalf("failed to create users table: %s", err.Error())
		tx.Rollback()
	}

	insertQuery = `INSERT INTO users (name, surname, email, phone_number, password, is_admin) VALUES
		('John', 'Doe', 'john.doe@example.com', '1234567890', 'password1', FALSE),
		('Jane', 'Smith', 'jane.smith@example.com', '0987654321', 'password2', FALSE),
		('Alice', 'Johnson', 'alice.johnson@example.com', '1231231230', 'password3', TRUE),
		('Bob', 'Brown', 'bob.brown@example.com', '0980980981', 'password4', FALSE),
		('Charlie', 'Davis', 'charlie.davis@example.com', '1230981230', 'password5', FALSE),
		('Eva', 'Wilson', 'eva.wilson@example.com', '0981230981', 'password6', TRUE),
		('Frank', 'Taylor', 'frank.taylor@example.com', '1231230980', 'password7', FALSE),
		('Grace', 'Anderson', 'grace.anderson@example.com', '0980981231', 'password8', FALSE),
		('Hank', 'Thomas', 'hank.thomas@example.com', '1230980981', 'password9', TRUE),
		('Ivy', 'Harris', 'ivy.harris@example.com', '0981231230', 'password10', FALSE)`

	_, err = tx.Exec(insertQuery)
	if err != nil {
		log.Fatalf("failed to insert users into table: %s", err.Error())
		tx.Rollback()
	}
	tx.Commit()

	//Orders
	createOrderQuery := `CREATE TABLE IF NOT EXISTS orders (
		id serial primary key,
		client_id integer REFERENCES users(id),
		total_price decimal(10, 2) DEFAULT 0,
		address varchar(255),
		is_delivery boolean,
		payment_type varchar(255),
		order_status boolean,
		created_at timestamp,
		updated_at timestamp
		)`

	tx, err = baseDb.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("failed to begin transaction for orders")
	}

	_, err = tx.Exec(createOrderQuery)
	if err != nil {
		log.Fatalf("failed to create orders table: %s", err.Error())
		tx.Rollback()
	}

	insertQuery = `INSERT INTO orders (client_id, total_price, address, is_delivery, payment_type, order_status, created_at, updated_at) VALUES
		(1, 150.75, '123 Main St, Springfield, IL', TRUE, 'Credit Card', TRUE, '2023-10-01 10:00:00', '2023-10-01 10:00:00'),
		(1, 230.50, '456 Elm St, Shelbyville, IL', FALSE, 'PayPal', FALSE, '2023-10-02 11:00:00', '2023-10-02 11:00:00'),
		(3, 99.99, '789 Oak St, Ogdenville, IL', TRUE, 'Cash', TRUE, '2023-10-03 12:00:00', '2023-10-03 12:00:00'),
		(4, 450.00, '321 Pine St, North Haverbrook, IL', FALSE, 'Credit Card', FALSE, '2023-10-04 13:00:00', '2023-10-04 13:00:00'),
		(5, 320.25, '654 Maple St, Brockway, IL', TRUE, 'PayPal', TRUE, '2023-10-05 14:00:00', '2023-10-05 14:00:00')`

	_, err = tx.Exec(insertQuery)
	if err != nil {
		log.Fatalf("failed to insert orders into table: %s", err.Error())
		tx.Rollback()
	}
	tx.Commit()

	// Orders-Products
	createOrdProdQuery := `CREATE TABLE IF NOT EXISTS orders_products (
		id serial primary key,
		product_id integer REFERENCES products(id),
		order_id integer REFERENCES orders(id),
		product_count integer
		)`

	tx, err = baseDb.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("failed to begin transaction for orders and products")
	}

	_, err = tx.Exec(createOrdProdQuery)
	if err != nil {
		log.Fatalf("failed to create orders_products table: %s", err.Error())
		tx.Rollback()
	}

	insertQuery = `INSERT INTO orders_products(product_id, order_id, product_count) VALUES
		(1, 1, 2),
		(2, 1, 2),
		(4, 2, 2),
		(2, 2, 2),
		(6, 3, 2),
		(8, 3, 2),
		(9, 4, 2),
		(1, 4, 2),
		(3, 5, 2),
		(7, 5, 2)`

	_, err = tx.Exec(insertQuery)
	if err != nil {
		log.Fatalf("failed to insert orders and products into table: %s", err.Error())
		tx.Rollback()
	}
	tx.Commit()
}
