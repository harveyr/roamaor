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
	"labix.org/v2/mgo/bson"
)

var webappPath string
var indexPage string
var currentUser *domain.User  // Use sessions for this

type PageContent struct {
    MenuHtml    template.HTML
    ContentHtml template.HTML
}

type AppPageContent struct {
	MenuHtml	template.HTML
	AdminHtml	template.HTML
}

func Jsonify(m interface{}) (s string) {
    b, err := json.Marshal(m)
    if err != nil {
            s = ""
            return
    }
    s = string(b)
    return
}

func WriteFailureResponse(w http.ResponseWriter, reason string) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]interface{})
	response["success"] = false
	response["reason"] = reason
	fmt.Fprintf(w, Jsonify(response))
	return
}

func WriteSuccessResponse(w http.ResponseWriter, response map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response["success"] = true
	fmt.Fprintf(w, Jsonify(response))
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

func bootstrapBundleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := domain.FetchUser("harveyr@gmail.com")
	if user == nil {
		WriteFailureResponse(w, "Failed to fetch user")
		return
	}
	currentUser = user
	data := make(map[string]interface{})
	data["success"] = true
	data["worldWidth"] = domain.WORLD_WIDTH
	data["worldHeight"] = domain.WORLD_HEIGHT
	data["user"] = user.Publicize()
	if len(user.ToonId) > 0 {
		toon := domain.FetchToonById(user.ToonId)
		data["toon"] = toon
		data["visited"] = domain.FetchLocationsVisited(toon)
	} else {
		data["toon"] = nil
	}
	fmt.Fprintf(w, Jsonify(data))
}

func setDestinationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var args map[string]float64
	if err := decoder.Decode(&args); err != nil {
		log.Fatalf("Failed to decode request body\n\tErr: %s\n\tBody: %s", err, r.Body)
	}
	toon := currentUser.Toon()
	if toon == nil {
		log.Print("No toon for user ", currentUser)
		WriteFailureResponse(w, "No toon for user")
	} else {
		toon.DestX = args["x"]
		toon.DestY = args["y"]
		toon.Save()
		response := make(map[string]interface{})
		response["toon"] = toon
		WriteSuccessResponse(w, response)
	}
	return
}

func setActiveToonHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var args map[string]bson.ObjectId
	if err := decoder.Decode(&args); err != nil {
		log.Print("[setActiveToonHandler] Failed to decode body: ", err)
	}
	log.Print("args: ", args)
	if currentUser == nil {
		log.Print("No current user!")
		WriteFailureResponse(w, "No active user")
		return
	}
	currentUser.ToonId = args["toonId"]
	currentUser.Save()
	WriteSuccessResponse(w, make(map[string]interface{}))
}

func adminNewToonHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	args := make(map[string]string)
	err := decoder.Decode(&args)
	if err != nil {
		log.Fatalf("Failed to decode request body\n\tErr: %s\n\tBody: %s", err, r.Body)
	}

	w.Header().Set("Content-Type", "application/json")
	if !domain.CanCreateToon(args["name"]) {
		w.WriteHeader(http.StatusNotModified)
	} else {
		toon := domain.NewToon(args["name"])
		fmt.Fprintf(w, Jsonify(toon))
	}
	return
}

func adminAllToonsHandler(w http.ResponseWriter, r *http.Request) {
	c := domain.GetCollection(domain.BEING_COLLECTION)
	query := make(map[string]interface{})
	var result []interface{}
	c.Find(query).All(&result)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, Jsonify(result))
	return
}

func adminAllLocsHandler(w http.ResponseWriter, r *http.Request) {
	c := domain.GetCollection(domain.LOCATION_COLLECTION)
	var result []domain.Location
	c.Find(make(map[string]interface{})).All(&result)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, Jsonify(result))
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

	if domain.CanCreateUser("harveyr@gmail.com") {
		domain.NewUser("harveyr@gmail.com")
	}

	http.HandleFunc("/app", appHandler)
	http.HandleFunc("/static/", assetsHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/api/destination", setDestinationHandler)
	http.HandleFunc("/api/activetoon", setActiveToonHandler)
	http.HandleFunc("/api/bootstrap", bootstrapBundleHandler)
	http.HandleFunc("/api/admin/newtoon", adminNewToonHandler)
	http.HandleFunc("/api/admin/alltoons", adminAllToonsHandler)
	http.HandleFunc("/api/admin/alllocations", adminAllLocsHandler)
	http.ListenAndServe(":8080", nil)
}
