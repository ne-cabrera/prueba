package serversInfoServices

import(
	"prueba/DB"
	"time"
)

type Request struct{
	HostName string `json:"hostName"`
}

type Response struct{
	Servers []ResponseServer `json:"servers"`
	Changed bool `json:"servers_changed"`
	Grade string `json:"ssl_grade"`
	PreviousGrade string `json:"previous_ssl_grade"`
	Logo string `json:"logo"`
	Title string `json:"title"`
	IsDown bool `json:"is_down"`
}

type ResponseServer struct{
	Address string `json:"addres"`
	Grade string `json:"ssl_grade"`
	Country string `json:"country"`
	Owner string `json:"owner"`
}

func MergeInfo(hostName string)(mergedInfo Response){
	res := Response{}
	sslInfo := getSSL(hostName)
	var resServers []ResponseServer
	for _, server := range sslInfo.Servers{
		whoisInfo := getWhois(server.Addres)
		resServer := ResponseServer{Address: server.Addres, Grade: server.Grade, Country: whoisInfo.Country, Owner: whoisInfo.Organization}
		resServers = append(resServers, resServer)
	}
	res.Servers = resServers
	res.Grade = sslInfo.Grade
	headInfo := getHeadInfo("https://www." + hostName)
	res.Logo = headInfo.Icon
	res.Title = headInfo.Title
	prevGrade := DB.GetPrevGrade(hostName)
	res.PreviousGrade = prevGrade
	res.Changed = sslInfo.Grade != prevGrade
	saveQuery(res, hostName)
	mergedInfo = res
	return
}

func saveQuery(res Response, hostName string){
	now := time.Now()
	unixNano := now.UnixNano()                                                                      
	nowMilli := unixNano / 1000000
	var servers []DB.QueryServers
	for _, server := range res.Servers{
		serv := DB.QueryServers{Address: server.Address, Grade: server.Grade, Country: server.Country, Owner: server.Owner}
		servers = append(servers, serv)
	}
	query := DB.Query{URL: hostName, Time: int(nowMilli), Grade: res.Grade, Servers: servers, Changed: res.Changed, PreviousGrade: res.PreviousGrade, Logo: res.Logo, Title: res.Title, IsDown: false}
	DB.InsertQuery(query)
}