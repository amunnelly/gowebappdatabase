package routing

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"gowebappdatabase/carpenter"
	"gowebappdatabase/commas"
	"gowebappdatabase/connectdb"
	"strings"
	"time"
	"os"

)

var t *template.Template
var i interface{}

// SeasonGraphDetail is a struct that supplies the title of a particular graph (the
// season itsself, invariably.)
type SeasonGraphDetail struct {
	Title string
	Results []connectdb.PointsGdTable
}

// ServeHome returns the home page template
func ServeHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nHome.")
	var g SeasonGraphDetail

	if len(r.URL.Query()) == 0 {
		fileBytes, err := ioutil.ReadFile("./queries/points_v_goal_difference.sql")
		if err != nil {
			log.Fatal(err)
		}
		season := "2017/2018"
		q := string(fileBytes)
		q = fmt.Sprintf(q, season)

		results := connectdb.PointsGdTableQuery(q)

		t, err := template.ParseGlob("./templates/*")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nSeason %s.", season)
		g.Title = season
		g.Results = results
		commas.CompileCsv(results)
		t.ExecuteTemplate(w, "index.html", g)
	}
}

