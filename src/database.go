package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"syscall"
	"time"

	"golang.org/x/term"

	_ "github.com/go-sql-driver/mysql"
)

// struct to hold values that will be used for the database connection
type databaseParams struct {
	username string
	password string
	hostname string
	dbName   string
}

func getDBparams() databaseParams {
	/*function to get all the parameters for the database such as username, password,
	hostname and databaseName

	Returns
	_______
	databaseParams
		a struct of type databaseParams with the provided values filled in
	*/

	var databaseName string

	// getting the user input for the username, password, and hostname.
	// These use the term module so that local echo is disabled and doesn't display
	// sensitive information
	fmt.Println("Please enter the username to log into the mysql server: ")
	byteUsername, err := term.ReadPassword(int(syscall.Stdin))

	if err != nil {
		log.Fatal(err)
	}

	username := string(byteUsername)

	fmt.Println("Please enter the password assorted with the username: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))

	if err != nil {
		log.Fatal(err)
	}

	password := string(bytePassword)

	fmt.Println("Please enter ip address of the database: ")
	byteHostname, err := term.ReadPassword(int(syscall.Stdin))

	if err != nil {
		log.Fatal(err)
	}

	hostname := string(byteHostname)

	// Asking for the user to provide a name for the database to be made. This
	// can be locally echoed
	fmt.Println("Please enter a name for the database to be created: ")
	fmt.Scanln(&databaseName)

	return databaseParams{username, password, hostname, databaseName}
}

func createDatabaseString(parameter_struct databaseParams, dbName string) string {
	/*function to create the string that will be passed to the sql.Open function
	Parameters
	__________
	parameter_struct databaseParams
		structure that has fields for the username, the password, the hostname, and a database name.
		The database name will not be used in this function

	dbName string
		string that has the name for database. It will be empty when we create a new database

	Returns
	_______
		string
		returns the string with all the necessary parameters
	*/
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", parameter_struct.username,
		parameter_struct.password,
		parameter_struct.hostname,
		dbName)
}
func create_new_db(database *sql.DB, databaseName string) {
	/*function to create the new database
	Parameters
	__________
	database *sql.DB
		pointer to the database object

	databaseName	*/
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := database.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+databaseName)

	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
	}

	// no, err := res.RowsAffected()
	// if err != nil {
	// 	log.Printf("Error %s when fetching rows", err)
	// }
	// log.Printf("rows affected %d\n", no)
}

func initialize_db() (*sql.DB, string, error) {
	/* function to create the database */

	parameters := getDBparams()

	// making sure that the connection is correct
	database, err := sql.Open("mysql", createDatabaseString(parameters, ""))

	if err != nil {
		log.Fatal(err)
	}

	create_new_db(database, parameters.dbName)

	database.Close()

	database, err = sql.Open("mysql", createDatabaseString(parameters, parameters.dbName))

	if err != nil {
		log.Fatal(err)
	}
	// defer database.Close()
	// setting the max number of connections
	database.SetMaxOpenConns(20)
	// setting the max number of idle connections
	database.SetMaxIdleConns(20)
	// setting the max time that a connection can have
	database.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// Checking to make sure that there is a connection to the database
	err = database.PingContext(ctx)

	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, parameters.dbName, err
	}
	log.Printf("Connected to DB %s successfully\n", parameters.dbName)
	return database, parameters.dbName, nil
}

func make_table(database *sql.DB, dbName string) {
	query := `CREATE TABLE IF NOT EXISTS ` + dbName + `(gene_id varchar(10), gene_start int,  
        gene_stop int, OMIM_id varchar(10), gene_name varchar(10), chromosome_number int)`
	fmt.Println(query)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	res, err := database.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rows affected when creating table: %d", rows)

	// continue following this https://golangbot.com/mysql-create-table-insert-row/

}
