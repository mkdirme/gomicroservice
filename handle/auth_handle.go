package handle

import (
	"encoding/json"
	"go-microservice/auth"
	"go-microservice/datastore"
	"log"
	"net/http"
)

// OAuth is the JSON Web Token response for OAuth access
type OAuth struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

//Login Handler to handle login request
func Login(w http.ResponseWriter, r *http.Request) {
	request := struct{ Email, Password string }{}
	var store datastore.UserStore = datastore.NewUserMgoStore()
	auth := auth.NewAuth()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	isUser, err := store.IsUser(request.Email, request.Password)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err)
		j, _ := json.Marshal(err)
		w.Write(j)
	} else if isUser == false {
		j, _ := json.Marshal(isUser)
		w.Write(j)
	}
	token := auth.GenToken(request.Email)
	response := new(OAuth)
	response.AccessToken = token
	j, _ := json.Marshal(response)
	w.Write(j)
}

// type Request struct {
// 	Username  string `json:"username"`
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// 	Email     string `json:"email"`
// 	Password  string `json:"password"`
// }

//RegisterNewUser Handler to handle login request
func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	request := struct{ Username, Firstname, Lastname, Email, Password string }{}
	var store datastore.UserStore = datastore.NewUserMgoStore()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	result, err := store.SaveUser(request.Email, request.Password)

	j, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
