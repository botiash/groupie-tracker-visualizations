package handle

import (
	"net/http"
	"text/template"
)

type structerr struct {
	ErrorText string
	Error     int
}

func ErrorHandler(w http.ResponseWriter, status int) {
	tmp, err := template.ParseFiles("./ui/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	errparse := structerr{
		ErrorText: http.StatusText(status),
		Error:     status,
	}
	if err := tmp.Execute(w, errparse); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
