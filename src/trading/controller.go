package trading

import (
	"fmt"
	"net/http"
	"strings"
)

// GetStringInBetween Returns empty string if no start string found
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str, end)
	return str[s:e]
}

func PostTrading(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		var result int
		input := r.FormValue("input")
		result = GetValue(GetSymbol(input))

		//if no result, show default message
		if result == 0 {
			fmt.Fprintf(w, "I have no idea what you are talking about")
		} else {

			if strings.Contains(input, "Silver") {
				result = result * 17
			}

			if strings.Contains(input, "Gold") {
				result = result * 14450
			}

			if strings.Contains(input, "Iron") {
				result = result * 391 / 2
			}

			fmt.Fprintf(w, fmt.Sprintf("%sis %d Credits", GetStringInBetween(input, "is", "?"), result))
		}
	}
}
