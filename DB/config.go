package DB

import(
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB(){
	conn, err := sql.Open("postgres",
		"postgresql://maxroach@localhost:26257/exercise_db?sslmode=disable")
	
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	db = conn
	if _, err := db.Exec(
        "CREATE TABLE IF NOT EXISTS queries (id SERIAL PRIMARY KEY, url TEXT, queryTime INT, grade TEXT, prevGrde TEXT, logo TEXT, changed BOOLEAN, title TEXT, isDown BOOLEAN)"); err != nil {
        log.Fatal(err)
	}
	if _, err := db.Exec(
        "CREATE TABLE IF NOT EXISTS servers (serv_id SERIAL PRIMARY KEY, query_id INTEGER REFERENCES queries(id), address TEXT, ssl_grade TEXT, country TEXT, owner TEXT)"); err != nil {
        log.Fatal(err)
    }
}

func getDB()(dataB *sql.DB){
	return db
}