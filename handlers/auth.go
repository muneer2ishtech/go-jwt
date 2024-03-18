package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"go-jwt/models"
	"go-jwt/utils"
)

// SignupHandler handles user signup
func SignupHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert the user into the database
	_, err = db.Exec("INSERT INTO users (email, password, firstname, lastname) VALUES (?, ?, ?, ?)", user.Email, hashedPassword, user.Firstname, user.Lastname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success status
	w.WriteHeader(http.StatusCreated)
}

// SigninHandler handles user signin and JWT generation
func SigninHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Query the database to retrieve user information
	row := db.QueryRow("SELECT email, password FROM users WHERE email = ?", user.Email)
	var storedEmail, storedPassword string
	err = row.Scan(&storedEmail, &storedPassword)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Compare hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Assuming user authentication is successful, generate JWT token
	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send token as response
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
