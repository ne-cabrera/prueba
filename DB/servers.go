package DB

import(
	"fmt"
	"log"
)

func InsertServer(queryID int, addres string, sslGrade string, country string, owner string){
	if _, err3 := db.Exec(
		fmt.Sprintf("INSERT INTO servers (query_id, address, ssl_grade, country, owner) VALUES (%d, '%s', '%s', '%s', '%s')", queryID, addres, sslGrade, country, owner)); err3 != nil {
		log.Fatal(err3)
	}
}

func GetServers(queryID int)(servers []QueryServers){
	rows, err := db.Query(fmt.Sprintf("SELECT address, ssl_grade, country, owner FROM servers WHERE query_id = %d", queryID))
    	if err != nil {
        	log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next(){
			var addres, sslGrade, country, owner string
			if err := rows.Scan(&addres, &sslGrade, &country, &owner); err != nil {
				log.Fatal(err)
			}
			server := QueryServers{Address: addres, Grade: sslGrade, Country: country, Owner: owner}
			servers = append(servers, server)
		}
	return
}