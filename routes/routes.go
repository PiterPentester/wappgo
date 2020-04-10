package routes

import (
	"net/http"

	"wappgo/middleware"
	"wappgo/models"
	"wappgo/sessions"
	"wappgo/utils"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.AuthRequired(indexGetHandler)).Methods("GET")
	r.HandleFunc("/", middleware.AuthRequired(indexPostHandler)).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/logout", logoutGetHandler).Methods("GET")
	r.HandleFunc("/register", registerGetHandler).Methods("GET")
	r.HandleFunc("/register", registerPostHandler).Methods("POST")

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/{username}", middleware.AuthRequired(userGetHandler)).Methods("GET")
	return r
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	updates, err := models.GetAllUpdates()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	utils.ExecuteTemplate(w, "index.html", struct {
		Title       string
		Updates     []*models.Update
		DisplayForm bool
	}{
		Title:       "All updates",
		Updates:     updates,
		DisplayForm: true,
	})
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	untypedUserID := session.Values["user_id"]
	userID, ok := untypedUserID.(int64)
	if !ok {
		utils.InternalServerError(w)
		return
	}
	r.ParseForm()
	body := r.PostForm.Get("update")
	err := models.PostUpdate(userID, body)
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	http.Redirect(w, r, "/", 302)
}

func userGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	untypedUserID := session.Values["user_id"]
	currentUserID, ok := untypedUserID.(int64)
	if !ok {
		utils.InternalServerError(w)
		return
	}
	vars := mux.Vars(r)
	username := vars["username"]
	user, err := models.GetUserByUsername(username)
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	userID, err := user.GetId()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	updates, err := models.GetUpdates(userID)
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	utils.ExecuteTemplate(w, "index.html", struct {
		Title       string
		Updates     []*models.Update
		DisplayForm bool
	}{
		Title:       username,
		Updates:     updates,
		DisplayForm: currentUserID == userID,
	})

}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	user, err := models.AuthenticateUser(username, password)
	if err != nil {
		switch err {
		case models.ErrUserNotFound:
			utils.ExecuteTemplate(w, "login.html", "Unknown user")
		case models.ErrInvalidLogin:
			utils.ExecuteTemplate(w, "login.html", "Invalid login")
		default:
			utils.InternalServerError(w)
		}
		return
	}
	userID, err := user.GetId()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	session, _ := sessions.Store.Get(r, "session")
	session.Values["user_id"] = userID
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func logoutGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	delete(session.Values, "user_id")
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	err := models.RegisterUser(username, password)
	if err == models.ErrUsernameTaken {
		utils.ExecuteTemplate(w, "register.html", "Username taken")
		return
	} else if err != nil {
		utils.InternalServerError(w)
		return
	}
	http.Redirect(w, r, "/login", 302)
}
