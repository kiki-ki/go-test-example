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
		userRepo: repository.NewUserRepository(db),
	}
}

type UserHandler interface {
	Index(http.ResponseWriter, *http.Request)
	Show(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userRepo repository.UserRepository
}

func (h *userHandler) Index(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepo.All()
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
	u, err := h.userRepo.Find(uId)
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
	u, err := h.userRepo.Find(uId)
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
	err = h.userRepo.Update(&u)
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
	err = h.userRepo.Create(&u)
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
	_, err = h.userRepo.Find(uId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	err = h.userRepo.Delete(uId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, fmt.Sprintf("id=%d is deleted", uId))
}
