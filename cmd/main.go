package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"runtime/debug"

	_ "github.com/lib/pq"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	debug.FreeOSMemory()
}
func bToMb(b uint64) uint64 {
	return b // 1024 / 1024
}

var db *sql.DB

// This function will make a connection to the database only once.
func init() {
	var err error

	connStr := "postgres://postgres:admin@localhost:2000/HELLO?sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("Connected to database")
	fmt.Println(db.Stats().OpenConnections)
	fmt.Println(db.Stats())
}

// type sandbox struct {
// 	a int
// 	b string
// 	c int
// 	d int
// }
type sandbox struct {
	a int64
	b int64
	c int64
	d int64
	e int64
	f int64
	g int64
	h int64
}

func retrieveRecord(w http.ResponseWriter, r *http.Request) {
	//PrintMemUsage()
	// checks if the request is a "GET" request
	if r.Method != "GET" {
		//http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	PrintMemUsage()
	// We assign the result to 'rows'

	rows, err := db.Query("SELECT * FROM newtable LIMIT 10000000")
	                     //SELECT * FROM newtable LIMIT 13042571
	PrintMemUsage()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// creates placeholder of the sandbox

	// we loop through the values of rows

	for rows.Next() {
		snb := sandbox{}
		err := rows.Scan(&snb.a, &snb.b, &snb.c, &snb.d, &snb.e, &snb.f, &snb.g, &snb.h)
		//err := rows.Scan(&snb.a, &snb.b, &snb.c, &snb.d)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		//for {
		fmt.Fprintf(w, "%d %d %d %d %d %d %d %d\n", snb.a, snb.b, snb.c, snb.d, snb.e, snb.f, snb.g, snb.h)
		//fmt.Fprintf(w, "%d %s %d %d \n", snb.a, snb.b, snb.c, snb.d)
		debug.FreeOSMemory()
		//}
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	rows.Scan()
	//rows.NextResultSet()
	PrintMemUsage()
}

// func measure() {
// 	for {
// 		PrintMemUsage()
// 		time.Sleep(2 * time.Second)
// 	}
// }
func main() {
	//go measure()
	// mux := http.NewServeMux()
	// mux.HandleFunc("/custom_debug_path/profile", pprof.Profile)
	// log.Fatal(http.ListenAndServe(":7777", mux))
	http.HandleFunc("/retrieve", retrieveRecord) // (1)
	http.ListenAndServe(":8080", nil)            // (2)
}
