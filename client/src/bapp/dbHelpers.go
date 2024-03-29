package bapp

import (
	"fmt"
	"net/http"

	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
)

func (a *App) CreateUser(user *models.User, w http.ResponseWriter, r *http.Request) error {
	err := a.DB.Create(user).Error
	if err != nil {
		fmt.Printf("Error creating user: %s", err.Error())
		http.Error(w, "Could not insert user into database", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (a *App) CreateClub(club *models.Club, w http.ResponseWriter, r *http.Request) error {
	err := a.Cdb.Create(club).Error
	if err != nil {
		fmt.Printf("Error creating club: %s", err.Error())
		http.Error(w, "Could not insert club into database", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (a *App) CreateEvent(event *models.Event, w http.ResponseWriter, r *http.Request) error {
	err := a.Edb.Create(event).Error
	if err != nil {
		fmt.Printf("Error creating event: %s", err.Error())
		http.Error(w, "Could not insert event into database", http.StatusInternalServerError)
		return err
	}
	return nil
}

//tested
// called to search for user when adding user, returns FALSE if no user found and TRUE if found
func (a *App) UserExists(name string, w http.ResponseWriter, r *http.Request) bool {
	//call is based on User Strcut not credentials struct, may need to change
	user := models.User{}
	if err := a.DB.First(&user, models.User{Name: name}).Error; err != nil {
		fmt.Println("User not located, adding to database...")
		return false
	}
	return true
}

// called to search for user when adding user, returns FALSE if no user found and TRUE if found
func (a *App) ClubExists(name string, w http.ResponseWriter, r *http.Request) bool {
	//call is based on User Strcut not credentials struct
	club := models.Club{}
	if err := a.Cdb.First(&club, models.Club{Name: name}).Error; err != nil {
		fmt.Println("Club not located, adding to database...")
		return false
	}
	return true
}
func (a *App) EventExists(name string) bool {
	event := models.Event{}
	if err := a.Edb.First(&event, models.Event{Event: name}).Error; err != nil {
		return false
	}
	return true
}

// searches DB for user, returns nil if none found
func (a *App) QueryDbByID(id string, w http.ResponseWriter, r *http.Request) *models.User {
	//call is based on User Struct not credentials struct
	user := models.User{}
	if err := a.DB.First(&user, models.User{ID: id}).Error; err != nil {
		//respondError(w, http.StatusNotFound, err.Error())
		http.Error(w, "User not located", http.StatusNotFound)
		return nil
	}
	return &user
}

// searches DB fpr user, returns nil if none found
func (a *App) QueryByName(email string, w http.ResponseWriter, r *http.Request) *models.User {
	//call is based on User Strcut not credentials struct
	user := models.User{}
	if err := a.DB.First(&user, models.User{Email: email}).Error; err != nil {
		fmt.Printf("Error: %s", err.Error())
		return nil
	}
	return &user
}

func (a *App) GetUserByName(name string, w http.ResponseWriter, r *http.Request) (*models.User, error) {
	fmt.Println("Entering GetUserByName")
	user := a.QueryByName(name, w, r)
	if user == nil {
		return nil, fmt.Errorf("user with name %s not found", name)
	}
	return user, nil
}

func (a *App) QueryByNName(name string, w http.ResponseWriter, r *http.Request) *models.User {
	//call is based on User Strcut not credentials struct
	user := models.User{}
	if err := a.DB.First(&user, models.User{Name: name}).Error; err != nil {
		fmt.Printf("Error: %s", err.Error())
		return nil
	}
	return &user
}

func (a *App) GetUserByNName(name string, w http.ResponseWriter, r *http.Request) (*models.User, error) {
	fmt.Println("Entering GetUserByNName")
	user := a.QueryByNName(name, w, r)
	if user == nil {
		return nil, fmt.Errorf("user with name %s not found", name)
	}
	return user, nil
}

//tested
func (a *App) GetUserByID(id string, w http.ResponseWriter, r *http.Request) (*models.User, error) {
	fmt.Println("Entering GetUserByID")
	//TREY: QUERY DB HERE FOR USER ID (Call QueryDbByID)
	user := a.QueryDbByID(id, w, r)
	if user == nil {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}
