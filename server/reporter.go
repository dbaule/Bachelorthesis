package main

import (
	"html/template"
	"log"
	"net/http"
)

type Error struct {
	InitMessage string
	NewMessage  string
}

type ViewControllerResult struct {
	ViewControllerName string
	InitBV             string
	NewBV              string
	InitScreenshot     string
	NewScreenshot      string
	Errors             []*Error
	NumErrors          string
}

type Results struct {
	ViewControllerResults []*ViewControllerResult
}

func generateHTML(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"safeURL": func(s string) template.URL {
			return template.URL(s)
		},
	}

	e1 := Error{InitMessage: "Size changed! Was: 1", NewMessage: "Size changed! Now: -1"}
	e2 := Error{InitMessage: "Position changed! Was: 10", NewMessage: "Position changed! Now: 100"}
	vc1 := ViewControllerResult{
		ViewControllerName: "VC1",
		InitBV:             "3",
		NewBV:              "4",
		InitScreenshot:     "img/abc.png",
		NewScreenshot:      "file:///Users/Danny/Documents/Studium/Bachelorarbeit/bachelor-thesis-3baule/App/GoServer/src/github.com/3baule/server/img/def.png",
		Errors:             []*Error{&e1, &e2},
		NumErrors:          "2"}
	vc2 := ViewControllerResult{
		ViewControllerName: "VC2",
		InitBV:             "3",
		NewBV:              "4",
		InitScreenshot:     "file:///Users/Danny/Documents/Studium/Bachelorarbeit/bachelor-thesis-3baule/App/GoServer/src/github.com/3baule/server/img/abc.png",
		NewScreenshot:      "file:///Users/Danny/Documents/Studium/Bachelorarbeit/bachelor-thesis-3baule/App/GoServer/src/github.com/3baule/server/img/def.png",
		Errors:             []*Error{&e1, &e2},
		NumErrors:          "2"}
	res := Results{ViewControllerResults: []*ViewControllerResult{&vc1, &vc2}}
	t := template.Must(template.New("a").Funcs(funcMap).ParseFiles("/Users/Danny/Documents/Studium/Bachelorarbeit/bachelor-thesis-3baule/App/GoServer/src/github.com/3baule/server/templates/results.html"))
	//t1 := template.Must(template.ParseFiles("/Users/Danny/Documents/Studium/Bachelorarbeit/bachelor-thesis-3baule/App/GoServer/src/github.com/3baule/server/templates/results.html"))
	err = t.ExecuteTemplate(w, "results.html", res)
	if err != nil {
		log.Println(err)
	}
}
