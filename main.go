package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Stats struct {
	dosesGiven string `json:"dosesGiven"`
	newDosesGiven string `json:"newDosesGiven"`
	fullyVaccinated string `json:"fullyVaccinated"`
	percentFullyVaccinated string `json:"percentFullyVaccinated"`
}

type Articles []Stats

/*func reqData() {

}*/

func returnStats(w http.ResponseWriter, r *http.Request) {
	stats := Articles{
		Stats{dosesGiven: "1024", newDosesGiven: "1024", fullyVaccinated: "1024", percentFullyVaccinated: "1024"},
	}

	fmt.Println("Endpoint Hit: All Stats Endpoint")
	json.NewEncoder(w).Encode(stats)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", returnStats)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
