package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var profiles = []Profile{}

func addProfile(w http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)
	w.Header().Set("Content-Type", "application/json")
	profiles = append(profiles, newProfile)
	json.NewEncoder(w).Encode(profiles)
}

func getAllProfiles(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)

}

func getProfileById(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Id Could not be converted into integer"))
		return
	}

	if id >= len(profiles) {
		w.WriteHeader(404)
		w.Write([]byte("NO profile found"))
		return
	}
	profile := profiles[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Id Could not be converted into integer"))
		return
	}

	if id >= len(profiles) {
		w.WriteHeader(404)
		w.Write([]byte("NO profile found"))
		return
	}
	var updatedProfile Profile
	json.NewDecoder(r.Body).Decode(&updatedProfile)
	profiles[id] = updatedProfile
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProfile)

}

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Id Could not be converted into integer"))
		return
	}

	if id >= len(profiles) {
		w.WriteHeader(404)
		w.Write([]byte("NO profile found"))
		return
	}
	profiles = append(profiles[:id], profiles[id+1:]...)
	w.WriteHeader(200)

}
