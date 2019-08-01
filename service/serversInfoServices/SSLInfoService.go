package serversInfoServices

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"time"
)

type SSLServers struct{
	Host string `json:"host"`
	Status string `json:"status"`
	Grade string `json:"ssl_grade"`
	Servers []Server `json:"endpoints"`
}

type Server struct{
	Addres string `json:"ipAddress"`
	Grade string `json:"grade"`
	Status string `json:"statusMessage"`
}

func getSSL(endpoint string)(SSLinformation SSLServers){
	var gotServers = getServers(endpoint)
	var servers = gotServers.Servers
	var SSLstatus = gotServers.Status
	for servers == nil {
		time.Sleep(500 * time.Millisecond)
		servers = getServers(endpoint).Servers
	}
	var grade = "A+"
	if SSLstatus != "READY"{
		for index, server := range servers {
			for server.Grade == ""{
				time.Sleep(1000 * time.Millisecond)
				gotServers = getServers(endpoint)
				servers = gotServers.Servers
				server = servers[index]
			}
			if compareGrades(server.Grade, grade) == false{
				grade = server.Grade
			}
		}
		gotServers.Grade = grade
		SSLinformation = gotServers
		return
	}else{
		fmt.Println(gotServers)
		grade = getGrade(gotServers.Servers)
		gotServers.Grade = grade
		SSLinformation = gotServers
		return
	}
}

func getServers(endpoint string)(servers SSLServers){
	resp, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + endpoint)
	if err != nil {
        log.Fatal(err)
    }
	defer resp.Body.Close()
	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		log.Fatal(errRead)
	}

	var respStatusObj SSLServers
	json.Unmarshal(body, &respStatusObj)
	//var status = respStatusObj.Status
	servers = respStatusObj
	return
}

func getGrade(servers []Server)(grade string){
	grade = "A+"
	for _, server := range servers{
		if compareGrades(server.Grade, grade) == false{
			grade = server.Grade
		}
	}
	return
}

func compareGrades( grade1, grade2 string) (isGrater bool){
	if grade1 == grade2{
		isGrater = false
		return
	} else if grade1 == "A+"{
		isGrater = true
		return
	}else if grade2 == "A+"{
		isGrater = false
		return
	}else if grade1 == "A-" && grade2 == "A"{
		isGrater = true
		return
	}else if grade1 == "A" && grade2 == "A-"{
		isGrater = false
		return
	}else{
		isGrater = !(grade1 > grade2)
		return
	}
}