package historyServices

import(
	"prueba/DB"
)

func GetHistory()(history []DB.Query){
	history = DB.GetQuerys()
	return
}