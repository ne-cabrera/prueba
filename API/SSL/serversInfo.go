package SSL

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"prueba/service/serversInfoServices"
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

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", GetSSLinfo)
	return router
}

func GetSSLinfo(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
    var req Request
    err := decoder.Decode(&req)
    if err != nil {
        panic(err)
    }
    fmt.Println(req.HostName)
	hostName := req.HostName
	res := serversInfoServices.MergeInfo(hostName)
	render.JSON(w,r, res)
}