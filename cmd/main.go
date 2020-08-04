package main

import (
​    "net/http"

    "github.com/gorilla/mux"
​    "github.com/gorilla/handlers" // Пакет с набором готовых хендлеров
)


func main() {

​    rt := mux.NewRouter()

    rt.Handle("/status", StatusHandler).Methods("GET")
    rt.Handle("/get-token", GetTokenHandler).Methods("GET")
    rt.Handle("/test", jwtMiddleware.Handler(TestHandler)).Methods("GET")
    
    http.ListenAndServe( ":3000", handlers.LoggingHandler(os.Stdout, r) )

//--------------------------------------------------------------------------------

	type TempSt struct {
	​    Id    int
	​    Name  string
	​    Slug  string
	}

	var products = []Product{
	​    Product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters"},
	​    Product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer"},
	}


	var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	​    // Конвертируем в json
	​    payload, _ := json.Marshal(products)
	    w.Header().Set("Content-Type", "application/json")
	    w.Write([]byte(payload))
		})

//---------------------------------------------------------------------------------

	var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	​    w.Write([]byte("Hello world API is up and running."))
	})

//---------------------------------------------------------------------------------

    // Глобальный секретный ключ
    var mySigningKey = []byte("secret")

    var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, 
                                            r *http.Request){
        // Создаем новый токен
        token := jwt.New(jwt.SigningMethodHS256)
    
        // Устанавливаем набор параметров для токена
        token.Claims["admin"] = true
        token.Claims["name"] = "First Last"
        token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
    
        // Подписываем токен нашим секретным ключем
        tokenString, _ := token.SignedString(mySigningKey)
    
        // Отдаем токен клиенту
        w.Write([]byte(tokenString))
    })

//-------------------------------------------------------------------------------------
	
	// проверка токена
	var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	    ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	        return mySigningKey, nil
	    },
	    SigningMethod: jwt.SigningMethodHS256,
	})

}
