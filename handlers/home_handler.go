package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// AppInfo represents information about the application name and version.
type AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// HomeHandler handles the root path and returns application information in JSON format.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Read version from the VERSION file
	version, err := getVersionFromFile("VERSION")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create an AppInfo instance with the application name and version
	appInfo := AppInfo{
		Name:    "pnlyzer-api",
		Version: version,
	}

	// Convert the AppInfo struct to JSON
	jsonData, err := json.Marshal(appInfo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}

func getVersionFromFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
