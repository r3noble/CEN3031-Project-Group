package backend_app

import(
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"	

	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
)


// var userMap map[int]User
type App struct {
	db *gorm.DB
	u  map[string]User
	r  *mux.Router
	mu sync.Mutex
}

func (a *App) start() {

	a.r.Use(func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})
	a.r.HandleFunc("/api/health", HealthCheck).Methods("GET")
	//query-based matching using id
	//user CRUD APIs
	a.r.HandleFunc("/api/getUser/{id}", a.IdHandler).Methods("GET")
	a.r.HandleFunc("/api/addUser", a.AddUserHandler).Methods("POST")
	a.r.HandleFunc("/api/login", a.loginHandler).Methods("POST") // handlers login
	//club CRUD APIs
	a.r.HandleFunc("/api/addClub", a.AddClubHandler).Methods("POST")
	//a.r.HandleFunc("/api/getClub", a.GetClubHandler).Methods("GET")
	//a.r.HandleFunc("/api/delClub", a.DeleteClubHandler).Methods("DELETE")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	handler := c.Handler(a.r)

	http.ListenAndServe(":8080", handler)
}