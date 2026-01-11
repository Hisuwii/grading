package main

import "net/http"

type ContactForm struct {
	Initial          bool
	SubmitButtonText string

	Name  string
	Email string
	Error string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Index().Render(r.Context(), w)
		View(ContactForm{}).Render(r.Context(), w)
	})
	http.ListenAndServe(":8080", nil)
}

func (m *ContactForm) InputValidation() {
	if m.Initial {
		return
	}

	if m.Name == "" {
		if m.Error != "" {
			m.Error += "\n"
		}
		m.Error += "*Name is required"
	}

	if m.Email == "" {
		if m.Error != "" {
			m.Error += "\n"
		}
		m.Error += "*Email is required"
	}
}
