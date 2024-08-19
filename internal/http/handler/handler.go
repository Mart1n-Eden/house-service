package handler

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"house-service/internal/http/handler/model/request"
	"house-service/internal/http/handler/model/response"
	"house-service/internal/http/handler/tools"
	"house-service/internal/logger"
	"house-service/pkg/utils/dbErrors"
)

func (h *Handler) CreateHouse(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.CreateHouse"

	houseReq, err := tools.Decode[request.HouseCreateRequest](r)
	if err != nil {
		logger.Error("decode request",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid json", http.StatusBadRequest)
		return
	}

	res, err := h.houseService.CreateHouse(r.Context(), houseReq.Address, houseReq.Year, *houseReq.Developer)
	if err != nil {
		if err.Error() == dbErrors.ErrFailedConnection {
			requestId := r.Context().Value("requestId").(string)
			logger.Error("create house",
				slog.String("op", op),
				slog.String("requestId", requestId),
				slog.String("error", err.Error()))
			tools.SendInternalError(w, err.Error(), requestId, http.StatusInternalServerError)
			return
		}
		logger.Error("create house",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, response.CreateHouseResponse(res), http.StatusOK)
}

func (h *Handler) GetHouse(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.GetHouse"

	id, _ := strings.CutPrefix(r.URL.Path, "/house/")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("convert id",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid id", http.StatusBadRequest)
		return
	}

	res, err := h.flatService.GetHouse(r.Context(), idInt)
	if err != nil {
		if err.Error() == dbErrors.ErrFailedConnection {
			requestId := r.Context().Value("requestId").(string)
			logger.Error("get house",
				slog.String("op", op),
				slog.String("requestId", requestId),
				slog.String("error", err.Error()))
			tools.SendInternalError(w, err.Error(), requestId, http.StatusInternalServerError)
			return
		}
		logger.Error("get house",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, response.CreateListFlatsResponse(res), http.StatusOK)
}

func (h *Handler) CreateFlat(w http.ResponseWriter, r *http.Request) {
	const op = "handler.CreateFlat"

	flatReq, err := tools.Decode[request.FlatCreateRequest](r)
	if err != nil {
		logger.Error("decode request",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid json", http.StatusBadRequest)
		return
	}

	res, err := h.flatService.CreateFlat(r.Context(), flatReq.HouseId, flatReq.Price, flatReq.Rooms)
	if err != nil {
		if err.Error() == dbErrors.ErrFailedConnection {
			requestId := r.Context().Value("requestId").(string)
			logger.Error("create flat",
				slog.String("op", op),
				slog.String("requestId", requestId),
				slog.String("error", err.Error()))
			tools.SendInternalError(w, err.Error(), requestId, http.StatusInternalServerError)
			return
		}
		logger.Error("create flat",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, response.CreateFlatResponse(res), http.StatusOK)
}

func (h *Handler) UpdateFlat(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.UpdateFlat"

	flatReq, err := tools.Decode[request.FlatUpdateRequest](r)
	if err != nil {
		logger.Error("decode request",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid json", http.StatusBadRequest)
		return
	}

	res, err := h.flatService.UpdateFlat(r.Context(), flatReq.Id, flatReq.Status)
	if err != nil {
		if err.Error() == dbErrors.ErrFailedConnection {
			requestId := r.Context().Value("requestId").(string)
			logger.Error("update flat",
				slog.String("op", op),
				slog.String("requestId", requestId),
				slog.String("error", err.Error()))
			tools.SendInternalError(w, err.Error(), requestId, http.StatusInternalServerError)
			return
		}
		logger.Error("update flat",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, response.CreateFlatResponse(res), http.StatusOK)
}

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.Registration"

	regReq, err := tools.Decode[request.RegistrationRequest](r)
	if err != nil {
		logger.Error("decode request",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid json", http.StatusBadRequest)
		return
	}

	userId, err := h.authService.CreateUser(r.Context(), regReq.Email, regReq.Password, regReq.UserType)
	if err != nil {
		if err.Error() == dbErrors.ErrFailedConnection {
			requestId := r.Context().Value("requestId").(string)
			logger.Error("registration",
				slog.String("op", op),
				slog.String("requestId", requestId),
				slog.String("error", err.Error()))
			tools.SendInternalError(w, err.Error(), requestId, http.StatusInternalServerError)
			return
		}
		logger.Error("registration",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, response.CreateUserIdResponse(userId), http.StatusOK)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.Login"

	loginReq, err := tools.Decode[request.LoginRequest](r)
	if err != nil {
		logger.Error("decode request",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid json", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(r.Context(), loginReq.Id, loginReq.Password)
	if err != nil {
		if err.Error() == dbErrors.ErrFailedConnection {
			requestId := r.Context().Value("requestId").(string)
			logger.Error("login",
				slog.String("op", op),
				slog.String("requestId", requestId),
				slog.String("error", err.Error()))
			tools.SendInternalError(w, err.Error(), requestId, http.StatusInternalServerError)
			return
		}
		logger.Error("login",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, response.CreateTokenResponse(token), http.StatusOK)
}

func (h *Handler) DummyLogin(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.DummyLogin"

	userType := r.FormValue("user_type")

	if userType == "" {
		logger.Error("login",
			slog.String("op", op),
			slog.String("error", "user type is empty"))
		tools.SendClientError(w, "user type is empty", http.StatusBadRequest)
		return
	}

	if userType != "moderator" && userType != "client" {
		logger.Error("login",
			slog.String("op", op),
			slog.String("error", "user type is invalid"))
		tools.SendClientError(w, "user type is invalid", http.StatusBadRequest)
		return
	}

	token, err := h.authService.DummyLogin(userType)
	if err != nil {
		logger.Error("login",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
	}

	tools.SendResponse(w, response.CreateTokenResponse(token), http.StatusOK)
}

func (h *Handler) NewSubscription(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.NewSubscription"

	id, _ := strings.CutPrefix(r.URL.Path, "/house/")
	id, _ = strings.CutSuffix(id, "/subscribe")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("login",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid id", http.StatusBadRequest)
		return
	}

	subReq, err := tools.Decode[request.SubscriptionRequest](r)
	if err != nil {
		logger.Error("decode request",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, "invalid json", http.StatusBadRequest)
		return
	}

	err = h.subscribeService.NewSubscription(r.Context(), subReq.Email, idInt)
	if err != nil {
		if err.Error() == dbErrors.ErrFailedConnection {
			requestId := r.Context().Value("requestId").(string)
			logger.Error("login",
				slog.String("op", op),
				slog.String("requestId", requestId),
				slog.String("error", err.Error()))
			tools.SendInternalError(w, err.Error(), requestId, http.StatusInternalServerError)
			return
		}
		logger.Error("login",
			slog.String("op", op),
			slog.String("error", err.Error()))
		tools.SendClientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tools.SendResponse(w, "Subscribe successfully", http.StatusOK)
}
