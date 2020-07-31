package main

import "log"

func main() {

	s := apiserver.New()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

// type User struct {
// 	Name  string
// 	Email string
// }

// func main() {

// 	r := mux.NewRouter()
// 	r.HandleFunc("/test", HomeHandler).Methods("GET")
// 	//r.HandleFunc("/auth/register", RegisterHandler)
// 	//r.HandleFunc("/auth/admin, AdminHandler)
// 	http.Handle("/", r)
// }

// // func RegisterHandler(w http.ResponseWriter, r *http.Request) {

// // }

// func HomeHandler(w http.ResponseWriter, r *http.Request) {

// 	data := []User{
// 		{
// 			Name:  "test-username",
// 			Email: "test@email.com",
// 		},
// 	}

// 	res, _ := josn.Marshal(data)
// 	NewResponse(res).WriteResponse(w)
// }
