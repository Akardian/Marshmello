package handler

import "net/http"

func HandleData(w http.ResponseWriter, r *http.Request) {
	html := `
        <div class="result">
            <h3>Data from Server</h3>
            <p>This was loaded via htmx!</p>
        </div>
    `
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
