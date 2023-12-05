package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// import (
// 	"log"
// 	"os"

//	"github.com/joho/godotenv"
//
// )
var (
	DB_URL string
)

// var (
// 	DBNAME     string
// 	DBURI      string
// 	TestDBURI  string
// 	TestDBNAME string
// )

// type Pagination struct {
// 	Limit int64
// 	Page  int64
// }

// type Store struct {
// 	User    UserStore
// 	Hotel   HotelStore
// 	Room    RoomStore
// 	Booking BookingStore
// }

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	DB_URL = os.Getenv("DB_URL")
}
