package handle

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//UserProfile  to handle login request
func UserProfile(w http.ResponseWriter, r *http.Request) {
	accessToken := r.FormValue("access_token")
	tokens := strings.SplitAfter(accessToken, ".")
	jwt, _ := base64.StdEncoding.DecodeString(tokens[1])
	fmt.Println(string(jwt))
	// var store datastore.UserStore = datastore.NewUserMgoStore()
	// user, err := store.FindUser("")
	// if err != nil {
	// 	log.Println(err)
	// 	json.NewEncoder(w).Encode(err)
	// }
	j, _ := json.Marshal(jwt)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
