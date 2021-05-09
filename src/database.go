package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/cheggaaa/pb/v3"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/term"
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
func create_new_db(database *sql.DB, databaseName string) error {
	/*function to create the new database
	Parameters
	__________
	database *sql.DB
		pointer to the database object

	databaseName
	*/
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := database.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+databaseName)

	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
	}

	return err
}

func initialize_db() (*sql.DB, string, error) {
	/* function to create the database */

	parameters := getDBparams()

	// making sure that the connection is correct
	database, err := sql.Open("mysql", createDatabaseString(parameters, ""))

	if err != nil {
		log.Fatal(err)
	}

	_ = create_new_db(database, parameters.dbName)

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

func insert(db *sql.DB, data_struct Data, dbName string) error {
	/*function to insert a new row into the mysql table
	Parameters
	__________
	db *sql.DB
		pointer to the database object

	data_struct Data
		The data struct that has fileds for things like the gene id,
		start and stop positions, omim id, full gene name, and chromosome id
	*/
	// creating the query string
	query := "INSERT INTO " + dbName + "(gene_id, gene_start, gene_stop, OMIM_id, gene_name, chromosome_number) VALUES (?, ?, ?, ?, ?, ?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	// preparing the query statement to be executed
	query_stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer query_stmt.Close()

	// executing the query statement to add values to a row
	response, err := query_stmt.ExecContext(ctx, data_struct.Gene_name, data_struct.Data.Gene.Start, data_struct.Data.Gene.Stop, data_struct.Data.Gene.Omim_id, data_struct.Data.Gene.Name, data_struct.Data.Gene.Chrom)

	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = response.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func make_table(database *sql.DB, dbName string, gene_info_list []Data) error {
	fmt.Println("creating a mysql table...")

	query := `CREATE TABLE IF NOT EXISTS ` + dbName + `(gene_id varchar(10), gene_start int,  
        gene_stop int, OMIM_id varchar(10), gene_name varchar(50), chromosome_number int)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	// executing the above query
	_, err := database.ExecContext(ctx, query)

	if err != nil {
		log.Fatal(err)
		return err
	}
	// creating a progress bar
	bar := pb.StartNew(len(gene_info_list))

	// iterating through each struct to place the info into a mysql
	// database
	fmt.Println("writing values to the database...")

	for i := 0; i < len(gene_info_list); i++ {
		err = insert(database, gene_info_list[i], dbName)

		if err != nil {
			log.Fatal(err)
			return err
		}
		bar.Increment()
	}
	bar.Finish()

	return nil
}
