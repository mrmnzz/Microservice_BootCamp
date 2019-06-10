 package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type response struct {
	Token string `json:"token"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/stg/tokens/{size}", tokenGen)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func tokenGen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	size := vars["size"]
	ln, err := strconv.Atoi(size)
	if err == nil {
		fmt.Println(ln)
	}

	bytes := make([]byte, ln)
	for i := 0; i < ln; i++ {
		var s string = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890"
        bytes[i] = byte(s[rand.Intn(len(s)-1)])
	}

	encodedToken := hex.EncodeToString(bytes)

	res := response{
		encodedToken,
	}

	fmt.Println(string(bytes))
	fmt.Println(res)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}