package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	var user User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = '"+username+"' AND password = '"+password+"'").Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate token (simplified)
	token := generateToken(user.ID)

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check for empty username or password
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Store the user in the database with plain text password (security issue)
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err = db.Exec(query, user.Username, hashPassword(user.Password))
	if err != nil {
		http.Error(w, "Error storing user in the database", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User signed up successfully")
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	rows, err := db.Query("SELECT id, user_id, text FROM todos WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.UserID, &todo.Text)
		todos = append(todos, todo)
	}

	json.NewEncoder(w).Encode(todos)
}
