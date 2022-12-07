package main

import (
    "fmt"
    "net/http"
	"advent-of-code-2022/day1"
    "advent-of-code-2022/day2"
    "advent-of-code-2022/day3"
    "advent-of-code-2022/day4"
    "advent-of-code-2022/day5"
    "advent-of-code-2022/day6"
    "advent-of-code-2022/day7"
)

func main() {
    http.HandleFunc("/hello", HelloServer)

    m := map[string]func()(string) {
		"day1": day1.Calculate,
		"day2": day2.Calculate,
		"day3": day3.Calculate,
		"day4": day4.Calculate,
		"day5": day5.Calculate,
		"day6": day6.Calculate,
        "day7": day7.Calculate,
	}

	for day, function := range m {
        path := fmt.Sprintf("/%s", day)
        http.HandleFunc(path, DayServer(function))
	}

    http.ListenAndServe(":8080", nil)
}

func DayServer(f func()(string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, f())
    }
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}