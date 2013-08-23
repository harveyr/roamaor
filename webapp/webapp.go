package main

import (
	"fmt"
	"bytes"
	"strings"
	"os"
	"log"
	"path/filepath"
	"net/http"
	"html/template"
	"encoding/json"
	"bitbucket.org/harveyr/roamaor/domain"
)

var webappPath string
var indexPage string

type PageContent struct {
    MenuHtml    template.HTML
    ContentHtml template.HTML
}

type AppPageContent struct {
	MenuHtml	template.HTML
	AdminHtml	template.HTML
}

func Jsonify(m map[string]interface{}) (s string) {
    b, err := json.Marshal(m)
    if err != nil {
            s = ""
            return
    }
    s = string(b)
    return
}

func parseTemplate(partial string, data interface{}) (out []byte, error error) {
	// See https://bitbucket.org/jzs/sketchground/src/4defb0a2ea64?at=default
    var buf bytes.Buffer
    file := filepath.Join(webappPath, "/static/templates/", partial)
    t, err := template.ParseFiles(file)
    if err != nil {
    	log.Printf("Error fetching partial %s: %s", partial, err)
        return nil, err
    }
    err = t.Execute(&buf, data)
    if err != nil {
            return nil, err
    }
    return buf.Bytes(), nil
}

func getPage(file string, data interface{}) []byte {
    var active string
    if strings.Contains(file, "project") {
            active = "Projects"
    } else if strings.Contains(file, "about") {
            active = "About"
    } else if strings.Contains(file, "post") {
            active = "Archive"
    } else if strings.Contains(file, "blog") || strings.Contains(file, "home") {
            active = "Blog"
    } else {
            active = ""
    }
    menu, error := parseTemplate("menu.html", map[string]string{active: "active"})
    if error != nil {
            print(error.Error())
    }
    page, error := parseTemplate(file, data)
    if error != nil {
            print(error.Error())
    }
    base, error := parseTemplate(
    	"base.html",
    	PageContent{
    		MenuHtml: template.HTML(menu),
    		ContentHtml: template.HTML(page)})
    if error != nil {
    	print(error.Error())
        return []byte("Internal server error...")
    }
    return base
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	menu, _ := parseTemplate("menu.html", map[string]string{})
	admin, _ := parseTemplate("adminBar.html", map[string]string{})
	page, _ := parseTemplate(
		"app.html",
		AppPageContent{
			MenuHtml: template.HTML(menu),
			AdminHtml: template.HTML(admin)}) 
	w.Write(page)
	// http.ServeFile(w, r, indexPage)
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(webappPath, r.URL.Path))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	if len(email) == 0 {
		page := getPage("login.html", nil)
		w.Write(page)
	}
}

func adminNewToonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	args := make(map[string]string)
	err := decoder.Decode(&args)
	if err != nil {
		log.Fatal("Failed to decode request body: ", r.Body)
	}
	toon := domain.NewToon(args["name"])
	if toon == nil {
		w.WriteHeader(http.StatusNotModified)
	} else {
		fmt.Fprintf(w, Jsonify(toon.Publicize()))
	}
	return
}

func main() {
	domain.InitDb("localhost", "roamaor")
	defer domain.CloseSession()
	
	envPath := os.Getenv("ROAMAOR_WEBAPP")
	p, err := filepath.Abs(envPath)
	if err != nil {
		log.Fatal("Could not resolve webapp path: ", envPath)
	}

	webappPath = p
	indexPage = filepath.Join(webappPath, "static/index.html")

	_, err = os.Stat(indexPage)
	if err != nil {
		log.Fatal("Could not find index.html at ", indexPage)
	}
	http.HandleFunc("/app", appHandler)
	http.HandleFunc("/static/", assetsHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/api/admin/newtoon", adminNewToonHandler)
	http.ListenAndServe(":8080", nil)
}
