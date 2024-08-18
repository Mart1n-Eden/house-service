package handler

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"house-service/internal/http/handler/model/request"
	"house-service/internal/http/handler/model/response"
	"house-service/internal/http/handler/tools"
	"house-service/pkg/utils/dbErrors"
)

func (h *Handler) CreateHouse(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "createHouse")

	houseReq, err := tools.Decode[request.HouseCreateRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid json", http.StatusBadRequest)
		return
	}

	res, err := h.houseService.CreateHouse(r.Context(), houseReq.Address, houseReq.Year, *houseReq.Developer)
	if err != nil {
		h.log.Error("create house", slog.String("error", err.Error()))
		if err.Error() == dbErrors.ErrFailedConnection {
			tools.SendInternalError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, response.CreateHouseResponse(res), http.StatusOK)
}

func (h *Handler) GetHouse(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "getHouse")

	id, _ := strings.CutPrefix(r.URL.Path, "/house/")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error("convert id", slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid id", http.StatusBadRequest)
		return
	}

	res, err := h.flatService.GetHouse(r.Context(), idInt)
	if err != nil {
		h.log.Error("get house", slog.String("error", err.Error()))
		if err.Error() == dbErrors.ErrFailedConnection {
			tools.SendInternalError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, response.CreateListFlatsResponse(res), http.StatusOK)
}

func (h *Handler) CreateFlat(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "createFlat")

	flatReq, err := tools.Decode[request.FlatCreateRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid json", http.StatusBadRequest)
		return
	}

	res, err := h.flatService.CreateFlat(r.Context(), flatReq.HouseId, flatReq.Price, flatReq.Rooms)
	if err != nil {
		h.log.Error("create flat", slog.String("error", err.Error()))
		if err.Error() == dbErrors.ErrFailedConnection {
			tools.SendInternalError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, response.CreateFlatResponse(res), http.StatusOK)
}

func (h *Handler) UpdateFlat(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "updateFlat")

	flatReq, err := tools.Decode[request.FlatUpdateRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid json", http.StatusBadRequest)
		return
	}

	res, err := h.flatService.UpdateFlat(r.Context(), flatReq.Id, flatReq.Status)
	if err != nil {
		h.log.Error("update flat", slog.String("error", err.Error()))
		if err.Error() == dbErrors.ErrFailedConnection {
			tools.SendInternalError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, response.CreateFlatResponse(res), http.StatusOK)
}

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "CreateUser")

	regReq, err := tools.Decode[request.RegistrationRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid json", http.StatusBadRequest)
		return
	}

	userId, err := h.authService.CreateUser(r.Context(), regReq.Email, regReq.Password, regReq.UserType)
	if err != nil {
		h.log.Error("registration", slog.String("error", err.Error()))
		if err.Error() == dbErrors.ErrFailedConnection {
			tools.SendInternalError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, response.CreateUserIdResponse(userId), http.StatusOK)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	//h.log = h.log.With("method", "login")

	loginReq, err := tools.Decode[request.LoginRequest](r)
	if err != nil {
		h.log.Error("decode request", slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid json", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(r.Context(), loginReq.Id, loginReq.Password)
	if err != nil {
		h.log.Error("login", slog.String("error", err.Error()))
		if err.Error() == dbErrors.ErrFailedConnection {
			tools.SendInternalError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, response.CreateTokenResponse(token), http.StatusOK)
}

func (h *Handler) DummyLogin(w http.ResponseWriter, r *http.Request) {
	userType := r.FormValue("user_type")

	if userType == "" {
		h.log.Error("login", slog.String("error", "user type is empty"))
		tools.SendClientError(w, "user type is empty", http.StatusBadRequest)
		return
	}

	// TODO: move to tools
	if userType != "moderator" && userType != "client" {
		h.log.Error("login", slog.String("error", "user type is invalid"))
		tools.SendClientError(w, "user type is invalid", http.StatusBadRequest)
		return
	}

	token, err := h.authService.DummyLogin(userType)
	if err != nil {
		h.log.Error("login", slog.String("error", err.Error()))
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
	}

	tools.SendResponse(w, response.CreateTokenResponse(token), http.StatusOK)
}
