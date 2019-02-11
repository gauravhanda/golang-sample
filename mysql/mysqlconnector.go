package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bavsales")
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(10)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	/*createNewProduct(db, Product{"3", "Spiritual Gems", 1.50, 1, "Charan Singh"})
	upateProductPrice(db, "1", 2.40)
	reduceInventory(db, "1", 1)*/
	addSale(db, "1", 1)
	fmt.Println("Connected successfully")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	_, _ = reader.ReadString('\n')

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}

/**
 *
 */
func createNewProduct(db *sql.DB, product Product) {
	insertSQL := fmt.Sprintf("INSERT INTO %s.%s(ID, NAME, PRICE, INVENTORY, DESCRIPTION) values ('%s', '%s', %f, '%d', '%s')",
		"bavsales", "PRODUCT", product.Id, product.Name, product.Price,
		product.Inventory, product.Description)

	fmt.Println(insertSQL)
	insert, err := db.Query(insertSQL)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

/**
 *
 */
func upateProductPrice(db *sql.DB, productID string, price float32) {
	updateSQL := fmt.Sprintf("UPDATE %s.%s SET price=%f where ID='%s'",
		"bavsales", "PRODUCT", price, productID)

	fmt.Println(updateSQL)

	update, err := db.Query(updateSQL)
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()
}

func getCurrentInventory(db *sql.DB, productID string) int {
	currentInventorySQL := fmt.Sprintf("select inventory from %s.%s where ID=?", "bavsales", "PRODUCT")
	fmt.Println(currentInventorySQL)
	currentQueryCount := 0
	// Open the connection to query the current value of inventory
	err := db.QueryRow(currentInventorySQL, productID).Scan(&currentQueryCount)

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Inventory count - %d", currentQueryCount)

	return currentQueryCount
}

/**
 *
 */
func reduceInventory(db *sql.DB, productID string, reduction int) {
	var inventory = getCurrentInventory(db, productID)
	if inventory == 0 {
		log.Fatal("Inventory is zero. Can't sale item")
		panic("Inventory is zero. Can't sale item")
	} else if reduction < 0 {
		log.Fatal("Sale cannot be negative")
		panic("Sale cannot be negative.")
	} else if inventory-reduction < 0 {
		log.Fatal("Insufficient Stock")
		panic("Insufficient Stock")
	}
	updatedStock := inventory - reduction
	updateInventory(db, productID, updatedStock)

}

func updateInventory(db *sql.DB, productID string, updatedInventory int) error {
	updateSQL := fmt.Sprintf("UPDATE %s.%s SET inventory=%d where ID='%s'",
		"bavsales", "PRODUCT", updatedInventory, productID)
	log.Println(updateSQL)
	updateResult, err := db.Query(updateSQL)
	defer updateResult.Close()
	if err != nil {
		log.Fatal("Error updating the inventory")
	} else {
		log.Println("Inventory updatedate successfully")
	}
	return err
}

func addSale(db *sql.DB, productID string, saleQuantity int) {
	reduceInventory(db, productID, saleQuantity)
	updateSaleSQL := fmt.Sprintf("INSERT INTO %s.%s (PRODUCT_ID, QUANTITY)  VALUES('%s', %d);", "bavsales", "SALES", productID, saleQuantity)
	log.Println("addSale: " + updateSaleSQL)
	updateSales, err := db.Query(updateSaleSQL)
	defer updateSales.Close()
	if err != nil {
		log.Fatal("Failed to update the Sales")
		panic(err.Error())
	}
}

func findDailySales() {
	dailySales := "SELECT C.SALE_DATE, SUM(C.SALE_AMOUNT) FROM (select DATE(SALE_DATE) AS SALE_DATE , A.price * B.quantity AS SALE_AMOUNT from PRODUCT A, SALES B where A.id = B.product_id) C GROUP BY C.SALE_DATE"

}
