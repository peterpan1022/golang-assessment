package main

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/anagram"
	"example.com/dictionary"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ErrHandler(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	result := map[string]interface{}{
		"message": err,
	}

	resp, err_ := json.Marshal(result)
	check(err_)

	w.Write(resp)
}

func FindHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.FormValue("lang")
	word := r.FormValue("word")

	if lang == "" {
		lang = "en"
	}

	if word == "" {
		http.Error(w, "word is required", http.StatusBadRequest)
		return
	}

	anagrams, err := anagram.FindAnagrams(lang, word)

	if err != "" {
		ErrHandler(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	result := map[string]interface{}{
		"count": len(anagrams),
		"list":  anagrams,
	}

	resp, err_ := json.Marshal(result)
	check(err_)

	w.Write(resp)
}

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	word1 := r.FormValue("word1")
	word2 := r.FormValue("word2")

	if word1 == "" {
		http.Error(w, "word1 is required", http.StatusBadRequest)
		return
	}

	if word2 == "" {
		http.Error(w, "word1 is required", http.StatusBadRequest)
		return
	}

	result := anagram.CheckAnagram(word1, word2)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(result)
	check(err)

	w.Write(resp)
}

func FindLongestHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.FormValue("lang")
	if lang == "" {
		lang = "en"
	}

	word := r.FormValue("word")

	if word == "" {
		http.Error(w, "word is required", http.StatusBadRequest)
		return
	}

	longestAnagram, distance, err := anagram.FindLongest(lang, word)

	if err != "" {
		ErrHandler(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	result := map[string]interface{}{
		"longestAnagram": longestAnagram,
		"distance":       distance,
	}

	resp, err_ := json.Marshal(result)
	check(err_)

	w.Write(resp)
}

type DictionaryBody struct {
	Lang  string   `json:lang`
	Words []string `json:words`
}

func DictionaryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var body DictionaryBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		if body.Lang == "" {
			http.Error(w, "lang is required", http.StatusBadRequest)
			return
		}

		dictionary.AddDictionary(body.Lang, body.Words)
	case "DELETE":
		var body DictionaryBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		if body.Lang == "" {
			http.Error(w, "lang is required", http.StatusBadRequest)
			return
		}

		err := dictionary.RemoveDictionary(body.Lang, body.Words)
		if err != "" {
			ErrHandler(w, err)
		}
	default:
		ErrHandler(w, "Sorry, only POST method is supported.")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	http.HandleFunc("/find", FindHandler)
	http.HandleFunc("/compare", CompareHandler)
	http.HandleFunc("/find-longest", FindLongestHandler)
	http.HandleFunc("/dictionary", DictionaryHandler)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
