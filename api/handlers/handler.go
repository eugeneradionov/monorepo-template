package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"monorepo-template/pkg/units"

	"github.com/go-chi/chi/v5"
)

func New() http.Handler {
	r := chi.NewRouter()

	r.Get("/", helloWorld)
	r.Get("/units/{unit}", unitHandler)

	return r
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Hello World!`))
}

func unitHandler(w http.ResponseWriter, r *http.Request) {
	var unit = strings.ToLower(chi.URLParam(r, "unit"))

	switch unit {
	case "kb", "kilobyte":
		w.Write([]byte(fmt.Sprintf("Kilobyte = %d bytes", units.KiB)))
	case "mb", "megabyte":
		w.Write([]byte(fmt.Sprintf("Megabyte = %d bytes", units.MiB)))
	case "gb", "gigabyte":
		w.Write([]byte(fmt.Sprintf("Gigabyte = %d bytes", units.GiB)))
	}
}
