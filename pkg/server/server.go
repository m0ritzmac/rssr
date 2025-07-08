package server

import (
	"fmt"
	"net/http"

	"github.com/m0ritzmac/rssr/pkg/web"
)

func Start() {
	templates := web.NewTemplate()

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.Render(w, "index", struct{ Title string }{Title: "Yo was geht ab"}); err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
