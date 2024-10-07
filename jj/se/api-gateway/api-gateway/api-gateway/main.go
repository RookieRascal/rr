package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Query struct defines the structure of a parsed SQL-like query.
type Query struct {
	Select []string          `json:"select"`
	From   string            `json:"from"`
	Where  map[string]string `json:"where"`
}

// Create a router using Go's built-in `http.ServeMux`.
func router() *http.ServeMux {
	mux := http.NewServeMux()

	// Register the query processing handler at the /query route.
	mux.HandleFunc("/query", queryHandler)

	return mux
}

// Handler function to process the incoming SQL-like query.
func queryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming query.
	var q Query
	err := parseQuery(r, &q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process the query and return the result.
	result := processQuery(q)

	// Send the response as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Function to parse the SQL-like query string from the request.
func parseQuery(r *http.Request, q *Query) error {
	query := r.URL.Query().Get("query")
	if query == "" {
		return fmt.Errorf("query parameter is missing")
	}

	// Split the query into parts to extract SELECT, FROM, and WHERE clauses.
	parts := strings.Split(query, " ")

	if len(parts) < 4 || strings.ToUpper(parts[0]) != "SELECT" || strings.ToUpper(parts[2]) != "FROM" {
		return fmt.Errorf("invalid query format")
	}

	// Parsing the SELECT columns.
	q.Select = strings.Split(parts[1], ",")

	// Parsing the FROM table.
	q.From = parts[3]

	// Parsing the WHERE conditions (if present).
	if len(parts) > 4 && strings.ToUpper(parts[4]) == "WHERE" {
		whereClause := parts[5:]
		q.Where = parseConditions(whereClause)
	}

	return nil
}

// Helper function to parse WHERE conditions in the query.
func parseConditions(conditions []string) map[string]string {
	conditionMap := make(map[string]string)

	for i := 0; i < len(conditions); i += 3 {
		if i+2 < len(conditions) && conditions[i+1] == "=" {
			conditionMap[conditions[i]] = conditions[i+2]
		}
	}

	return conditionMap
}

// Function to process the query and return a result (simulated with dummy data).
func processQuery(q Query) map[string]interface{} {
	// Simulate dummy data that might come from a database.
	dummyData := []map[string]interface{}{
		{"name": "John", "age": 30, "email": "john@example.com"},
		{"name": "Jane", "age": 25, "email": "jane@example.com"},
		{"name": "Doe", "age": 30, "email": "doe@example.com"},
	}

	// Filter the dummy data based on the WHERE conditions.
	filteredData := []map[string]interface{}{}
	for _, record := range dummyData {
		match := true
		for col, val := range q.Where {
			if fmt.Sprint(record[col]) != val {
				match = false
				break
			}
		}
		if match {
			filteredData = append(filteredData, record)
		}
	}

	// Select only the requested columns from the filtered data.
	finalResult := []map[string]interface{}{}
	for _, record := range filteredData {
		selectedRecord := make(map[string]interface{})
		for _, col := range q.Select {
			selectedRecord[col] = record[col]
		}
		finalResult = append(finalResult, selectedRecord)
	}

	// Return the processed result.
	return map[string]interface{}{
		"result": finalResult,
	}
}

func main() {
	// Create a new router.
	mux := router()

	// Start the HTTP server on port 8080.
	fmt.Println("Starting server at :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
