package database
import (
    "database/sql"
    "fmt"
    "strconv"
)

// Database instance
var DB *sql.DB

// Connect function
func Connect() error {
    var err error
    p := config.Config("5050")
    // because our config function returns a string, we are parsing our str to int here 
    port,err := strconv.ParseUint(p, 10, 32) 
    if err != nil {
        fmt.Println("Error parsing str to int")
    }
	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("localhost"), port,
	 config.Config("root"), config.Config("fbr01994"), config.Config("gomysql")))

    if err != nil {
        return err
    }
    if err = DB.Ping(); err != nil {
        return err
    }
    CreateProductTable()
    fmt.Println("Connection Opened to Database")
    return nil
}
It exports a single m