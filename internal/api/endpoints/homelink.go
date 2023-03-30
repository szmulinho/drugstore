package endpoints

import (
	"net/http"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "C:\\Program Files\\Go\\src\\github.com\\szmulinho\\drugstore\\cmd\\server\\static\\index.html")
}
