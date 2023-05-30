package handle

import (
	"group/model"
	"net/http"
	"strconv"
	"text/template"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil || id > 52 || id < 1 {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("./ui/artist.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	Artist, statusCode := model.GetOneArtist(idstr)
	if statusCode != http.StatusOK {
		ErrorHandler(w, statusCode)
		return
	}
	Relation, statusCode := model.GetRelation(idstr)
	if statusCode != http.StatusOK {
		ErrorHandler(w, statusCode)
		return
	}

	Location, statusCode := model.GetLocation(idstr)
	if statusCode != http.StatusOK {
		ErrorHandler(w, statusCode)
		return
	}
	Date, statusCode := model.GetDate(idstr)
	if statusCode != http.StatusOK {
		ErrorHandler(w, statusCode)
		return
	}
	asd := struct {
		Arti model.Artist
		Rela model.Relation
		Loca model.Location
		Date model.Date
	}{
		Arti: Artist,
		Rela: Relation,
		Loca: Location,
		Date: Date,
	}
	if err := tmp.Execute(w, asd); err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
