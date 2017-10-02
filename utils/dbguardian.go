package utils

// import "github.com/fzzy/radix/redis"
import "github.com/asdine/storm"
import m "../models"
import "fmt"

var (
	// client      *redis.Client
	db          *storm.DB
	initialized bool
)

// Initialize VOLTDB
func Initialize(dbpath string) {
	if dbpath != "" {
		var err error
		db, err = storm.Open(dbpath)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Print("no db path!")
	}
}

// ReadData VOLTDB
func ReadData() []m.User {

	defer db.Close()

	var user []m.User
	err := db.All(&user)
	if err != nil {
		panic(err)
	}
	return user
}

// WriteData VOLTDB
func WriteData(user *m.User) {
	defer db.Close()

	err := db.Save(user)
	if err != nil {
		panic(err)
	}
}

// Initialize Redis Client
// func Initialize() {
// 	initialized = false
// 	var err error
// 	client, err = redis.Dial("tcp", "localhost:6379")
// 	if err != nil {
// 		// handle err
// 		initialized = false
// 		panic(err)
// 	}
// 	initialized = true
// }

// Query testing radix query
// func Query() (string, error) {
//
// 	if initialized == false {
// 		Initialize()
// 	}
//
// 	users, err := client.Cmd("GET", "users").Str()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return users, nil
// }
