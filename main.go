package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GeoData struct {
	Country string
	IsoCode string
	Data    []struct {
		Date                            string
		TotalVaccinations               int     `json:"total_vaccinations"`
		PeopleVaccinated                int     `json:"people_vaccinated"`
		PeopleFullyVaccinated           int     `json:"people_fully_vaccinated"`
		DailyVaccinationsRaw            int     `json:"daily_vaccinations_raw"`
		DailyVaccinations               int     `json:"daily_vaccinations"`
		TotalVaccinationsPerHundred     float64 `json:"total_vaccinations_per_hundred"`
		PeopleVaccinatedPerHundred      float64 `json:"people_vaccinated_per_hundred"`
		PeopleFullyVaccinatedPerHundred float64 `json:"people_fully_vaccinated_per_hundred"`
		DailyVaccinationsperMillion     int     `json:"daily_vaccinations_per_million"`
	}
}

type Stats struct {
	Location                        string
	Date                            string
	TotalVaccinations               int
	PeopleVaccinated                int
	PeopleFullyVaccinated           int
	DailyVaccinationsRaw            int
	DailyVaccinations               int
	TotalVaccinationsPerHundred     float64
	PeopleVaccinatedPerHundred      float64
	PeopleFullyVaccinatedPerHundred float64
	DailyVaccinationsperMillion     int
}

type Articles []Stats

var date string
var totalVaccinations, peopleVaccinated, peopleFullyVaccinated, dailyVaccinationsRaw, dailyVaccinations, dailyVaccinationsperMillion int
var totalVaccinationsPerHundred, peopleVaccinatedPerHundred, peopleFullyVaccinatedPerHundred float64

func returnStats(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://raw.githubusercontent.com/owid/covid-19-data/master/public/data/vaccinations/vaccinations.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	data := []GeoData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, val := range data {
		if val.Country == "World" {
			for _, data := range val.Data {
				date = data.Date
				totalVaccinations = data.TotalVaccinations
				peopleVaccinated = data.PeopleVaccinated
				peopleFullyVaccinated = data.PeopleFullyVaccinated
				dailyVaccinationsRaw = data.DailyVaccinationsRaw
				dailyVaccinations = data.DailyVaccinations
				totalVaccinationsPerHundred = data.TotalVaccinationsPerHundred
				peopleVaccinatedPerHundred = data.PeopleVaccinatedPerHundred
				peopleFullyVaccinatedPerHundred = data.PeopleFullyVaccinatedPerHundred
				dailyVaccinationsperMillion = data.DailyVaccinationsperMillion
			}
		}

		// return
	}

	stats := Articles{
		Stats{
			Location:                        "Global",
			Date:                            date,
			TotalVaccinations:               totalVaccinations,
			PeopleVaccinated:                peopleVaccinated,
			PeopleFullyVaccinated:           peopleFullyVaccinated,
			DailyVaccinationsRaw:            dailyVaccinationsRaw,
			DailyVaccinations:               dailyVaccinations,
			TotalVaccinationsPerHundred:     totalVaccinationsPerHundred,
			PeopleVaccinatedPerHundred:      peopleFullyVaccinatedPerHundred,
			PeopleFullyVaccinatedPerHundred: peopleVaccinatedPerHundred,
			DailyVaccinationsperMillion:     dailyVaccinationsperMillion,
		},
	}

	fmt.Println("Endpoint Hit: All Stats Endpoint")
	json.NewEncoder(w).Encode(stats)
}

func homePage(w http.ResponseWriter, r *http.Request) {
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
