package main

import (
	"fmt"
	"net/http"
	"storageModule/shape"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL.Query().Get("message"))
	})

	http.HandleFunc("/shape", func(w http.ResponseWriter, r *http.Request) {
		var currentShape shape.Shape

		shapeType := r.URL.Query().Get("type")

		switch shapeType {
		case "square":
			currentShape = shape.NewSquare(10)
		case "triangle":
			currentShape = shape.NewTriangle(10, 10, 30)
		default:
			currentShape = shape.NewSquare(11)
		}

		square, _ := currentShape.GetSquare()

		fmt.Fprintf(w, "square of %s is %f", shapeType, square)
	})

	http.ListenAndServe(":80", nil)
}
