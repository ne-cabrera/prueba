package DB

import(
	"log"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

type Query struct{
	URL string	`json:"url"`
	Time int	`json:"time"`
	Grade string	`json:"grade"`
	Servers []QueryServers `json:"servers"`
	Changed bool `json:"servers_changed"`
	PreviousGrade string `json:"previous_ssl_grade"`
	Logo string `json:"logo"`
	Title string `json:"title"`
	IsDown bool `json:"is_down"`
}

type QueryServers struct{
	Address string `json:"addres"`
	Grade string `json:"ssl_grade"`
	Country string `json:"country"`
	Owner string `json:"owner"`
}

func InsertQuery(query Query){
	var url, time, grade, changed, prevGrade, logo, title, isDown = query.URL, query.Time, query.Grade, query.Changed,
	query.PreviousGrade, query.Logo, query.Title, query.IsDown
	rows, err := db.Query(
		fmt.Sprintf("INSERT INTO queries (url, queryTime, grade, prevGrde, logo, changed, title, isDown) VALUES ('%s', %d, '%s', '%s', '%s', %t, '%s', %t) RETURNING id", url, time, grade, prevGrade, logo, changed, title, isDown))
	if err != nil {
		log.Fatal(err)
	}
	var id int
	if rows.Next(){
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(id)
	for _, server := range query.Servers{
		var addres, sslGrade, country, owner = server.Address, server.Grade, server.Country, server.Owner
		InsertServer(id, addres, sslGrade, country, owner)
	}
}

func GetQuerys()(queries []Query){
	rows, err := db.Query("SELECT * FROM queries")
    if err != nil {
        log.Fatal(err)
    }
	defer rows.Close()
	for rows.Next(){
		var url, grade, prevGrade, logo, title string
		var id, time int
		var changed, isDown bool
		if err := rows.Scan(&id, &url, &time, &grade, &prevGrade, &logo, &changed, &title, &isDown); err != nil {
            log.Fatal(err)
		}
		servers := GetServers(id)
		query := Query{URL: url,Time: time, Grade: grade, Servers: servers, Changed: changed, PreviousGrade: prevGrade, Logo: logo, Title: title, IsDown: isDown}
		queries = append(queries, query)
	}
	return
}

func GetPrevGrade(hostName string)(prevGrade string){
	rows, err := db.Query("SELECT grade, queryTime FROM queries WHERE url ='" + hostName + "' ORDER BY queryTime DESC")
    if err != nil {
        log.Fatal(err)
	}
	if rows.Next(){
		var grade string
		var qTime int
		now := time.Now()
		unixNano := now.UnixNano()                                                                      
        nowMilli := unixNano / 1000000
		if err := rows.Scan(&grade, &qTime); err != nil {
            log.Fatal(err)
		}
		if int(nowMilli) - qTime > 3600000{
			prevGrade = grade
		}else{
			prevGrade = ""
		}
	}
	defer rows.Close()
	return
}