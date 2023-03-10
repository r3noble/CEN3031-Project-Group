package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	//"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/initializers"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// var userMap map[int]User
type App struct {
	db *gorm.DB
	u  map[string]User
	r  *mux.Router
	mu sync.Mutex
}

func WriteOnceMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := &responseWriter{w, false}
		h.ServeHTTP(wrappedWriter, r)
		if !wrappedWriter.wroteHeader {
			wrappedWriter.WriteHeader(http.StatusOK)
		}
	})
}

type responseWriter struct {
	http.ResponseWriter
	wroteHeader bool
}

func (w *responseWriter) WriteHeader(statusCode int) {
	if w.wroteHeader {
		return
	}
	w.ResponseWriter.WriteHeader(statusCode)
	w.wroteHeader = true
}

func (a *App) start() {
	//Initialize and open DB here
	a.db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Error in opening DB")
	}
	//calls AutoMigrate and throws error if cannot migrate
	//formats db to replicate user struct
	err := a.db.AutoMigrate(&User{})
	if err != nil {
		panic("Error in migrating db")
	}

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
	a.r.HandleFunc("/api/getUser/{id}", a.IdHandler).Methods("GET")
	a.r.HandleFunc("/api/addUser", a.AddUserHandler).Methods("POST")
	a.r.HandleFunc("/api/login", a.loginHandler).Methods("POST") // handlers login
	http.ListenAndServe(":8080", a.r)
}
func main() {
	app := App{
		db: db,
		u: make(map[string]User),
		r: mux.NewRouter(),
	}
	app.u["Cole"] = User{ID: "1", Name: "Cole", Email: "cole@rottenberg.org", Password: "pass"}

	app.start()
}

func (a *App) QueryDbByID(id string) {

}

func (a *App) GetUserByID(id string) (*User, error) {
	fmt.Println("Entering GetUserByID")
	//TREY: QUERY DB HERE FOR USER ID (Call QueryDbByID)
	user, ok := a.u[id]
	if !ok {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	return &user, nil
}

func (a *App) loginHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST and the URL path is /user/login
	// Decode the JSON payload from the request body
	fmt.Println("Successfully entered Login Handler")
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Bad Json in Body")
		return
	}

	// Check if the required fields (username and password) are present
	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Authenticate the user using the provided credentials (not shown)
	// ...
	//TREY: QUERY DB here for username
	user, ok := a.u[creds.Username]
	if !ok {
		http.Error(w, "Invalid Username", http.StatusUnauthorized)
		fmt.Println("No found user")
		return
	}
	//now we check the password
	knownPass := user.Password
	if knownPass != creds.Password {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		fmt.Println("No found password")
		return
	}
	/*response := struct {
		Message string `json:"message"`
	}{
		Message: "Login successful",
	}*/

	jsonResponse, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("About to pass back user")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	fmt.Println("Passing back success")
	// Send a success response
	return

	// Send a 404 Not Found response if the URL path doesn't match
}
func (a *App) IdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	id := vars["id"]
	// Look up the user with the given id in the map
	//TREY: Get user by ID must be updated for DB support
	user, err := a.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func (a *App) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the new user data
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newUser.ID = strconv.Itoa(rand.Intn(1000))

	// Check if the user ID already exists in the map
	//TREY: Query DB for username, if EXISTS, print same error
	if _, ok := a.u[newUser.Name]; ok {
		http.Error(w, "User with that ID already exists", http.StatusBadRequest)
		return
	}

	// Add the new user to the map
	//TREY: Call function to add new user to map
	a.mu.Lock()
	defer a.mu.Unlock()
	a.u[newUser.Name] = newUser

	// Return the new user data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//next function writes back to the response
	fmt.Println("Health check accessed")
	fmt.Fprintf(w, "API is running")
}

func (a *App) profileHandler(w http.ResponseWriter, r *http.Request) {
	// Get the username parameter from the URL path
	username := r.URL.Query().Get("username")

	// Retrieve the profile data from the map
	//TREY: QUERY DB for username
	profile, ok := a.u[username]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Convert the profile data to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
