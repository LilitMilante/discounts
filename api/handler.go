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
	s.r.HandleFunc("/users", s.AddUser).Methods(http.MethodPost)
	s.r.HandleFunc("/users", s.Users).Methods(http.MethodGet)
	s.r.HandleFunc("/users/{number}", s.User).Methods(http.MethodGet)
	s.r.HandleFunc("/users/{number}", s.ChangeUser).Methods(http.MethodPatch)
	s.r.HandleFunc("/users/{number}", s.DeleteUser).Methods(http.MethodDelete)

	httpServer := http.Server{
		Addr:    ":" + port,
		Handler: s.r,
	}
	return httpServer.ListenAndServe()
}

func (s Server) Users(w http.ResponseWriter, r *http.Request) {
	f := entity.UserFilter{}

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

	u, err := s.ds.Users(f)
	if err != nil {
		sendERR(w, http.StatusBadRequest, err)

		return
	}

	if u == nil {
		u = make([]entity.User, 0)
	}

	sendOK(w, http.StatusOK, u)
}

func (s Server) User(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)["phone"]

	ph, err := s.ds.UserByPhone(p)
	if err != nil {
		sendERR(w, http.StatusInternalServerError, err)

		return
	}

	sendOK(w, http.StatusOK, ph)
}

func (s Server) AddUser(w http.ResponseWriter, r *http.Request) {
	var u entity.Client

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		sendERR(w, http.StatusBadRequest, err)

		return
	}

	err = u.Validate()
	if err != nil {
		sendERR(w, http.StatusBadRequest, err)

		return
	}

	u, err = s.ds.CreateUser(u)
	if err != nil {
		sendERR(w, http.StatusInternalServerError, err)

		return
	}

	sendOK(w, http.StatusOK, u)
}

func (s Server) ChangeUser(w http.ResponseWriter, r *http.Request) {
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

func (s Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	n := mux.Vars(r)["number"]

	err := s.ds.DeleteClientDiscount(n)
	if err != nil {
		sendERR(w, http.StatusNotFound, err)
	}

	sendOK(w, http.StatusOK, n)
}
