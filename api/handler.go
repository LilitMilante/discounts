package api

import (
	"discounts/domain/entity"
	"discounts/domain/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

const timeLayout = "02-01-2006"

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
	s.r.HandleFunc("/discounts/{number}", s.DeleteClientDiscount).Methods(http.MethodDelete)

	httpServer := http.Server{
		Addr:    ":" + port,
		Handler: s.r,
	}
	return httpServer.ListenAndServe()
}

func (s Server) ClientDiscounts(w http.ResponseWriter, r *http.Request) {
	f := entity.ClientFilter{}

	q := r.URL.Query()
	name := q.Get("name")
	if name != "" {
		f.Name = &name
	}

	sale := q.Get("sale")
	if sale != "" {
		slInt, err := strconv.Atoi(sale)
		if err != nil {
			sendERR(w, http.StatusBadRequest, err)

			return
		}
		p := int8(slInt)
		f.Sale = &p
	}

	startDate := q.Get("start")
	if startDate != "" {
		t, err := time.Parse(timeLayout, startDate)
		if err != nil {
			sendERR(w, http.StatusBadRequest, err)

			return
		}

		f.Start = &t
	}

	endData := q.Get("end")
	if endData != "" {
		t, err := time.Parse(timeLayout, endData)
		if err != nil {
			sendERR(w, http.StatusBadRequest, err)

			return
		}

		t = t.Add(24 * time.Hour)

		f.End = &t
	}

	d, err := s.ds.ClientDiscounts(f)
	if err != nil {
		sendERR(w, http.StatusBadRequest, err)

		return
	}

	if d == nil {
		d = make([]entity.Client, 0)
	}

	sendOK(w, http.StatusOK, d)
}

func (s Server) ClientDiscount(w http.ResponseWriter, r *http.Request) {
	n := mux.Vars(r)["number"]

	d, err := s.ds.ClientDiscountByNumber(n)
	if err != nil {
		sendERR(w, http.StatusInternalServerError, err)

		return
	}

	sendOK(w, http.StatusOK, d)
}

func (s Server) AddClientDiscount(w http.ResponseWriter, r *http.Request) {
	var d entity.Client

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		sendERR(w, http.StatusBadRequest, err)

		return
	}

	err = d.Validate()
	if err != nil {
		sendERR(w, http.StatusBadRequest, err)

		return
	}

	d, err = s.ds.CreatedClientDiscount(d)
	if err != nil {
		sendERR(w, http.StatusInternalServerError, err)

		return
	}

	sendOK(w, http.StatusOK, d)
}

func (s Server) ChangeClientDiscount(w http.ResponseWriter, r *http.Request) {
	var ud entity.UpdateClient

	err := json.NewDecoder(r.Body).Decode(&ud)
	if err != nil {
		sendERR(w, http.StatusBadRequest, err)

		return
	}

	err = ud.Validate()
	if err != nil {
		sendERR(w, http.StatusBadRequest, err)

		return
	}

	n := mux.Vars(r)["number"]

	d, err := s.ds.EditClientDiscountByNumber(ud, n)
	if err != nil {
		sendERR(w, http.StatusInternalServerError, err)

		return
	}

	sendOK(w, http.StatusOK, d)

}

func (s Server) DeleteClientDiscount(w http.ResponseWriter, r *http.Request) {
	n := mux.Vars(r)["number"]

	err := s.ds.DeleteClientDiscount(n)
	if err != nil {
		sendERR(w, http.StatusNotFound, err)
	}

	sendOK(w, http.StatusOK, n)
}
