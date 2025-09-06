package api

import "net/http"

func (app *Application) healthHandler(w http.ResponseWriter , r * http.Request){
	w . Write([]byte("ok"))
}