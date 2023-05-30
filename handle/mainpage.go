package handle

import (
	"group/model"
	"net/http"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	tmp, err := template.ParseFiles("./ui/index.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	artists, errStatus := model.GetAllArtists()
	if errStatus != http.StatusOK {
		ErrorHandler(w, errStatus)
		return
	}
	if err := tmp.Execute(w, artists); err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
