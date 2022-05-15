package api

import (
	"discounts/domain/entity"
	"discounts/domain/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	r  *mux.Router
	ds *service.DiscountService
}

func NewServer(ds *service.DiscountService) *Server {
	r := mux.NewRouter()
	s := Server{
		r:  r,
		ds: ds,
	}

	return &s
}

func (s Server) Start(port string) error {
	s.r.HandleFunc("/discounts", s.AddClientDiscount).Methods(http.MethodPost)

	return http.ListenAndServe(":"+port, s.r)
}

func (s Server) AddClientDiscount(w http.ResponseWriter, r *http.Request) {
	var d entity.ClientDiscount

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	d, err = s.ds.CreatedClientDiscount(d)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
