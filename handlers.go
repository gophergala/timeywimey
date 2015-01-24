package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	body, err := localFile("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.FormValue("profile")
	passphrase := r.FormValue("p")

	if profile == "" || passphrase == "" {
		// Look for a cookie?
	} else {
		// Set cookies
		expire := time.Now().AddDate(0, 0, 1)
		cookieProfile := &http.Cookie{
			Name:    "profile",
			Value:   profile,
			Expires: expire,
		}
		cookiePassphrase := &http.Cookie{
			Name:    "passphrase",
			Value:   passphrase,
			Expires: expire,
		}
		http.SetCookie(w, cookieProfile)
		http.SetCookie(w, cookiePassphrase)
	}

	if profile == "" || passphrase == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/"+profile, http.StatusSeeOther)
	return
}

func UserIndexHandler(w http.ResponseWriter, r *http.Request) {
	//username := mux.Vars(r)["username"]

	//projects, err := GetByUsername(username)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	body, err := localFile("user_index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func ProjectIndexHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := localFile("project_index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func ProjectNewHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := localFile("project_new.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func MeetingEditHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := localFile("meeting_edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func MeetingShowHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := localFile("meeting_index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func MeetingUpdateHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var m Moment
	if err := m.FromJson(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Retrieve project
	// TODO: Add moment to project
}

func MeetingDeleteHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IssueEditHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := localFile("issue_edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func IssueShowHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := localFile("issue_index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func IssueUpdateHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var t Task
	if err := t.FromJson(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Retrieve project
	// TODO: Add moment to project
}

func IssueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TimeEditHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := localFile("time_edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func TimeShowHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := localFile("time_index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func TimeUpdateHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var t TimeEntry
	if err := t.FromJson(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Retrieve project
	// TODO: Add moment to project
}

func TimeDeleteHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func InvoiceEditHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := localFile("invoice_edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func InvoiceShowHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := localFile("invoice_index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func InvoiceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var invoice Invoice
	if err := invoice.FromJson(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Retrieve project
	// TODO: Add moment to project
}

func InvoiceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	projectName := mux.Vars(r)["project"]

	var p Project
	err := p.Get(projectName, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func localFile(filename string) ([]byte, error) {
	page, err := ioutil.ReadFile("html/" + filename)
	if err != nil {
		return nil, err
	}
	return page, nil
}
