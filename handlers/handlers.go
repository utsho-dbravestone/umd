package handlers



import (

	"database/sql"
	"encoding/json"

	"fmt"

	"log"

	"net/http"



	"github.com/gorilla/mux"

)



func Home(resposeWriter http.ResponseWriter, request *http.Request) {

	resposeWriter.Header().Set("Content Type", "application/json")

	json.NewEncoder(resposeWriter).Encode("{msg : welcome to utsho's movie downloader}")

	fmt.Printf("[%v: %v]\n", request.Method, request.RequestURI)

}



func Search(resposeWriter http.ResponseWriter, request *http.Request) {

	db, err := sql.Open("sqlite3", "./movies.db")

	if err != nil {

		log.Fatal(err)

	}

	defer db.Close()



	resposeWriter.Header().Set("Content Type", "application/json")

	params := mux.Vars(request)



	query := fmt.Sprintf(`select name, link from movies where like('%%%v%%', name) order by name asc`, params["name"])

	data, err := db.Query(query)

	var (

		name, link string

	)

	var movies []string

	for data.Next() {

		data.Scan(&name, &link)

		movies = append(movies, fmt.Sprintf("{%s, %s}", name, link))

	}



	json.NewEncoder(resposeWriter).Encode(movies)

	fmt.Printf("[%v: %v]\n", request.Method, request.RequestURI)

}
