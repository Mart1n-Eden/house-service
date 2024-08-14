package handler

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"house-service/internal/http/handler/tools"
	"house-service/internal/http/model/request"
	"house-service/internal/http/model/response"
	"house-service/pkg/utils/dbErrors"
)

func (h *Handler) createHouse(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "createHouse")

	houseReq, err := tools.Decode[request.HouseCreateRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.houseService.CreateHouse(r.Context(), houseReq.Address, houseReq.Year, houseReq.Developer)
	if err != nil {
		h.log.Error("create house", slog.String("error", err.Error()))
		if err.Error() == dbErrors.ErrFailedConnection {
			tools.SendInternalError(w, err.Error(), http.StatusInternalServerError)
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, res, http.StatusOK)
}

func (h *Handler) getHouse(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "getHouse")

	id, _ := strings.CutPrefix(r.URL.Path, "/house/")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error("convert id", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.houseService.GetHouse(r.Context(), idInt)
	if err != nil {
		h.log.Error("get house", slog.String("error", err.Error()))
		if err.Error() == dbErrors.ErrFailedConnection {
			tools.SendInternalError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, res, http.StatusOK)
}

func (h *Handler) createFlat(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "createFlat")

	flatReq, err := tools.Decode[request.FlatCreateRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.flatService.CreateFlat(r.Context(), flatReq.HouseId, flatReq.Price, flatReq.Rooms)
	if err != nil {
		h.log.Error("create flat", slog.String("error", err.Error()))
		if err.Error() == dbErrors.ErrFailedConnection {
			tools.SendInternalError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, res, http.StatusOK)
}

func (h *Handler) updateFlat(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "updateFlat")

	flatReq, err := tools.Decode[request.FlatUpdateRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.flatService.UpdateFlat(r.Context(), flatReq.Id, flatReq.Status)
	if err != nil {
		h.log.Error("update flat", slog.String("error", err.Error()))
		if err.Error() == dbErrors.ErrFailedConnection {
			tools.SendInternalError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, res, http.StatusOK)
}

func (h *Handler) registration(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "CreateUser")

	regReq, err := tools.Decode[request.RegistrationRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := h.authService.CreateUser(r.Context(), regReq.Email, regReq.Password, regReq.UserType)
	if err != nil {
		h.log.Error("registration", slog.String("error", err.Error()))
		if err.Error() == dbErrors.ErrFailedConnection {
			tools.SendInternalError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := response.UserIdResponse{UserId: userId}

	tools.SendResponse(w, res, http.StatusOK)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "login")

	loginReq, err := tools.Decode[request.LoginRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(r.Context(), loginReq.Id, loginReq.Password)
	if err != nil {
		h.log.Error("login", slog.String("error", err.Error()))
		if err.Error() == dbErrors.ErrFailedConnection {
			tools.SendInternalError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := response.TokenResponse{Token: token}

	tools.SendResponse(w, res, http.StatusOK)
}

func (h *Handler) dummyLogin(w http.ResponseWriter, r *http.Request) {
	userType := r.FormValue("user_type")

	if userType == "" {
		h.log.Error("login", slog.String("error", "user type is empty"))
		http.Error(w, "user type is empty", http.StatusBadRequest)
		return
	}

	// TODO: move to tools
	if userType != "moderator" && userType != "client" {
		h.log.Error("login", slog.String("error", "user type is invalid"))
		http.Error(w, "user type is invalid", http.StatusBadRequest)
		return
	}

	token, err := h.authService.DummyLogin(r.Context(), userType)
	if err != nil {
		h.log.Error("login", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// TODO: refactor to constructor
	res := response.TokenResponse{Token: token}

	tools.SendResponse(w, res, http.StatusOK)
}
