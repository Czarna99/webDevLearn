package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")

}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is an about page and 2 + 2 is  %d", sum))
}

func addValues(x, y int) int {

	sum := x + y
	return sum
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%.2f divided by %.2f is %.2f", 100.0, 0.0, f))

}

func divideValues(x, y float64) (float64, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	fmt.Println("Application is running...")
	http.ListenAndServe(portNumber, nil)
}
