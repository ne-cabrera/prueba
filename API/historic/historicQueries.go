package historic

import(
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"prueba/service/historyServices"
)

type HistoryResponse struct{
	Items []HistoryItem `json:"items"`
}

type HistoryItem struct{
	HostName string `json:"hostName"`
	Servers []HistoryServers `json:"servers"`
	Changed bool `json:"servers_changed"`
	Grade string `json:"ssl_grade"`
	PreviousGrade string `json:"previous_ssl_grade"`
	Logo string `json:"logo"`
	Title string `json:"title"`
	IsDown bool `json:"is_down"`
}

type HistoryServers struct{
	Address string `json:"addres"`
	Grade string `json:"ssl_grade"`
	Country string `json:"country"`
	Owner string `json:"owner"`
}

func RoutesHistoric() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetAllItems)
	return router
}

func GetAllItems(w http.ResponseWriter, r *http.Request){
	history := historyServices.GetHistory()
	res := HistoryResponse{}
	var items []HistoryItem
	for _, item := range history{
		hostItem := HistoryItem{HostName: item.URL, Changed: item.Changed, Grade: item.Grade, PreviousGrade: item.PreviousGrade, Logo: item.Logo, Title: item.Title, IsDown: item.IsDown}
		var servers []HistoryServers
		for _, serverItem := range item.Servers{
			server := HistoryServers{Address: serverItem.Address, Grade: serverItem.Grade, Country: serverItem.Country, Owner: serverItem.Owner}
			servers = append(servers, server)
		}
		hostItem.Servers = servers
		items = append(items, hostItem)
	}
	res.Items = items
	render.JSON(w,r, res)
}