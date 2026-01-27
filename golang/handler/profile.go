package handler

import (
	"html/template"
	"net/http"
)

type User struct {
	Name      string
	Email     string
	AvatarURL string
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	profile := template.Must(template.ParseFiles("./static/src/component/profile/profile-dropdown.html"))

	user := User{
		Name:      "John Doe",
		Email:     "john@example.com",
		AvatarURL: "/avatars/john.jpg",
	}

	err := profile.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleUserIcon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	profile := template.Must(template.ParseFiles("./static/src/component/profile/profile-icon.html"))

	err := profile.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
