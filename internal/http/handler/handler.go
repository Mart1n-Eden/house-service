package handler

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"house-service/internal/http/handler/tools"
	"house-service/internal/http/model/errors"
	"house-service/internal/http/model/request"
	"house-service/internal/http/model/response"
)

func (h *Handler) createHouse(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "createHouse")
	houseReq, err := tools.Decode[request.HouseCreateRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		errors.ResponseError(err.Error(), w, http.StatusBadRequest)
		return
	}

	res, err := h.houseService.CreateHouse(r.Context(), houseReq.Address, houseReq.Year, houseReq.Developer)
	if err != nil {
		h.log.Error("create house", slog.String("error", err.Error()))
		errors.ResponseError(err.Error(), w, http.StatusInternalServerError)
		return
	}

	response.SendResponse(res, w, http.StatusOK)
}

func (h *Handler) getHouse(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "getHouse")

	id, _ := strings.CutPrefix(r.URL.Path, "/house/")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error("convert id", slog.String("error", err.Error()))
		errors.ResponseError(err.Error(), w, http.StatusBadRequest)
		return
	}

	res, err := h.houseService.GetHouse(r.Context(), idInt)
	if err != nil {
		h.log.Error("get house", slog.String("error", err.Error()))
		errors.ResponseError(err.Error(), w, http.StatusInternalServerError)
		return
	}

	response.SendResponse(res, w, http.StatusOK)
}

func (h *Handler) createFlat(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "createFlat")
	flatReq, err := tools.Decode[request.FlatCreateRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		errors.ResponseError(err.Error(), w, http.StatusBadRequest)
		return
	}

	res, err := h.flatService.CreateFlat(r.Context(), flatReq.HouseId, flatReq.Price, flatReq.Rooms)
	if err != nil {
		h.log.Error("create flat", slog.String("error", err.Error()))
		errors.ResponseError(err.Error(), w, http.StatusInternalServerError)
		return
	}

	response.SendResponse(res, w, http.StatusOK)
}

func (h *Handler) updateFlat(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "updateFlat")

	flatReq, err := tools.Decode[request.FlatUpdateRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		errors.ResponseError(err.Error(), w, http.StatusBadRequest)
		return
	}

	res, err := h.flatService.UpdateFlat(r.Context(), flatReq.Id, flatReq.Status)
	if err != nil {
		h.log.Error("update flat", slog.String("error", err.Error()))
		errors.ResponseError(err.Error(), w, http.StatusInternalServerError)
		return
	}

	response.SendResponse(res, w, http.StatusOK)
}

func (h *Handler) registration(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "CreateUser")
	regReq, err := tools.Decode[request.RegistrationRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		errors.ResponseError(err.Error(), w, http.StatusBadRequest)
		return
	}

	userId, err := h.authService.CreateUser(r.Context(), regReq.Email, regReq.Password, regReq.UserType)
	if err != nil {
		h.log.Error("registration", slog.String("error", err.Error()))
		errors.ResponseError(err.Error(), w, http.StatusInternalServerError)
		return
	}

	res := response.UserIdResponse{UserId: userId}

	response.SendResponse(res, w, http.StatusOK)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "login")
	loginReq, err := tools.Decode[request.LoginRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		errors.ResponseError(err.Error(), w, http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(r.Context(), loginReq.Id, loginReq.Password)
	if err != nil {
		h.log.Error("login", slog.String("error", err.Error()))
		errors.ResponseError(err.Error(), w, http.StatusInternalServerError)
		return
	}

	res := response.TokenResponse{Token: token}

	response.SendResponse(res, w, http.StatusOK)
}
