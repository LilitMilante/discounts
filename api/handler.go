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
	s.r.HandleFunc("/discounts", s.ClientDiscounts).Methods(http.MethodGet)
	s.r.HandleFunc("/discounts/{number}", s.ClientDiscount).Methods(http.MethodGet)
	s.r.HandleFunc("/discounts/{number}", s.ChangeClientDiscount).Methods(http.MethodPatch)

	httpServer := http.Server{
		Addr:    ":" + port,
		Handler: s.r,
	}
	return httpServer.ListenAndServe()
}

func (s Server) ClientDiscounts(w http.ResponseWriter, _ *http.Request) {
	d, err := s.ds.ClientDiscounts()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if d == nil {
		d = make([]entity.ClientDiscount, 0)
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s Server) ClientDiscount(w http.ResponseWriter, r *http.Request) {
	n := mux.Vars(r)["number"]

	d, err := s.ds.ClientDiscountByNumber(n)
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

func (s Server) ChangeClientDiscount(w http.ResponseWriter, r *http.Request) {
	var ud entity.UpdateClientDiscount

	err := json.NewDecoder(r.Body).Decode(&ud)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	n := mux.Vars(r)["number"]

	d, err := s.ds.EditClientDiscountByNumber(ud, n)
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
