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
    "advent-of-code-2022/day8"
    "advent-of-code-2022/day9"
    "advent-of-code-2022/day10"
    "advent-of-code-2022/day11"
    "advent-of-code-2022/day12"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.HandleFunc("/hello", HelloServer)

    m := map[string]func()(string) {
		"day1": day1.Calculate,
		"day2": day2.Calculate,
		"day3": day3.Calculate,
		"day4": day4.Calculate,
		"day5": day5.Calculate,
		"day6": day6.Calculate,
        "day7": day7.Calculate,
        "day8": day8.Calculate,
        "day9": day9.Calculate,
        "day10": day10.Calculate,
        "day11": day11.Calculate,
        "day12": day12.Calculate,
	}

	for day, function := range m {
        path := fmt.Sprintf("/%s", day)
        http.HandleFunc(path, DayServer(function))
	}

    // Azure App Service sets the port as an Environment Variable
	// This can be random, so needs to be loaded at startup
	port := os.Getenv("HTTP_PLATFORM_PORT")

	// default back to 8080 for local dev
	if port == "" {
		port = "8080"
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