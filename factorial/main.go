package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)

	log.Println("Listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	input, ok := r.URL.Query()["number"]

	log.Printf("URL params received: %v", input)

	if !ok || len(input[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("URL parameter 'number' is missing"))
		return
	}

	n, err := strconv.Atoi(input[0])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not parse URL parameter 'number': not an integer"))
		return
	}

	response, err := factorial(n)

	if err != nil {
		switch err.(type) {
		case NegativeFactorialError:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		case IntegerOverflowError:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(response)))
}
