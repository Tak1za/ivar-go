package main

import (
	"fmt"
	"ivar-go/src/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type", "Accept"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	//Router Setup
	router := mux.NewRouter().StrictSlash(true)
	router.Use(httpsRedirectMiddleware)
	//router.Use(middleware, httpsRedirectMiddleware, authMiddleware)

	//User related routes
	router.HandleFunc("/users/{username}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	//Post related routes
	router.HandleFunc("/posts", controllers.GetPostsByUserId).Queries("u", "{u}").Methods("GET")
	router.HandleFunc("/posts/{postId}", controllers.GetPostByPostId).Queries("u", "{u}").Methods("GET")
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")

	//Comment related routes
	router.HandleFunc("/comments", controllers.CreateComment).Methods("POST")

	//Like related routes
	router.HandleFunc("/likes", controllers.AddLikeToPost).Methods("POST")
	router.HandleFunc("/likes", controllers.GetLikersForPost).Queries("u", "{u}").Queries("p", "{p}").Methods("GET")

	//Follower related routes
	router.HandleFunc("/followers", controllers.GetFollowers).Queries("u", "{u}").Methods("GET")

	fmt.Println("IVAR-Go listening at port: 8080")
	// log.Fatal(http.ListenAndServe(":8080", router))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func httpsRedirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		proto := req.Header.Get("x-forwarded-proto")
		if proto == "http" || proto == "HTTP" {
			http.Redirect(res, req, fmt.Sprintf("https://%s%s", req.Host, req.URL), http.StatusPermanentRedirect)
			return
		}
		next.ServeHTTP(res, req)
	})
}

//func authMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		accessToken := w.Header().Get("accessToken")
//		_, err := client.VerifyAccessToken(accessToken)
//		if err != nil {
//			w.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//		next.ServeHTTP(w, r)
//	})
//}
