package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "bookinfo"
)

type ratings struct {
	Id   int `json:"Id"`
	Star int `json:"Star"`
}

func main() {
	reviewer1 := ratings{
		Id:   1,
		Star: 5,
	}
	reviewer2 := ratings{
		Id:   2,
		Star: 4,
	}

	reviewer := []ratings{reviewer1, reviewer2}
	bs, err := json.Marshal(reviewer)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/rating", func(w http.ResponseWriter, r *http.Request) {
		w.Write(bs)
	})
	http.ListenAndServe(":8000", nil)

}
