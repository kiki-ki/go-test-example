package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/kiki-ki/go-test-example/internal/app/handler/request"
	"github.com/kiki-ki/go-test-example/internal/app/model"
	"github.com/kiki-ki/go-test-example/internal/app/repository"
	"github.com/kiki-ki/go-test-example/internal/interface/database"
)

func NewUserHandler(db database.DB) UserHandler {
	return &userHandler{
		db: db,
		userRepo: repository.NewUserRepository(),
	}
}

type UserHandler interface {
	Index(http.ResponseWriter, *http.Request)
	Show(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Transaction(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	db database.DB
	userRepo repository.UserRepository
}

func (h *userHandler) Index(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepo.All(h.db.Conn())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, users)
}

func (h *userHandler) Show(w http.ResponseWriter, r *http.Request) {
	uId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	u, err := h.userRepo.Find(uId, h.db.Conn())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, u)
}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	uId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	u, err := h.userRepo.Find(uId, h.db.Conn())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	var req request.UserUpdateReq
	err = render.DecodeJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	u.Name = req.Name
	u.Email = req.Email
	u.Age = req.Age
	err = h.userRepo.Update(&u, h.db.Conn())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, u)
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.UserUpdateReq
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	var u model.User
	u.Name = req.Name
	u.Email = req.Email
	u.Age = req.Age
	err = h.userRepo.Create(&u, h.db.Conn())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, u)
}

func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	uId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	_, err = h.userRepo.Find(uId, h.db.Conn())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	err = h.userRepo.Delete(uId, h.db.Conn())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, fmt.Sprintf("id=%d is deleted", uId))
}

func (h *userHandler) Transaction(w http.ResponseWriter, r *http.Request) {
	tx, err := h.db.Conn().Begin()
	defer tx.Rollback()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	u := new(model.User)
	u.Name = "a"
	u.Email = "a@exa.com"
	u.Age = 20
	err = h.userRepo.Create(u, tx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	u.Id = 0
	u.Email = "b@exa.com"
	err = h.userRepo.Create(u, tx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	tx.Commit()
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, u)
}
