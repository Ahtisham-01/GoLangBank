package api // replace with a seperate package

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"GoLangBank/utils"

	"github.com/gorilla/mux" // fiber
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APiError struct {
    Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := f(w, r); err != nil {
            WriteJSON(w, http.StatusBadRequest, APiError{Error: err.Error()})
        }
    }
}

type APIServer struct {
    listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
    return &APIServer{
        listenAddr: listenAddr,
    }
}

func (s *APIServer) Run() {
    router := mux.NewRouter()
    router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

    log.Println("JSON API server running on port", s.listenAddr)
    http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
    if r.Method == "GET" {
        return s.handleGetAccount(w, r)
    }
    if r.Method == "POST" {
        return s.handleCreateAccount(w, r)
    }
    if r.Method == "DELETE" {
        return s.handleDeleteAccount(w, r)
    }
    return fmt.Errorf("Method not allowed")
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
    
	account:=utils.NewAccount("Ahtisham","kiani")
    return WriteJSON(w, http.StatusOK, account)
}
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
    return nil
}
