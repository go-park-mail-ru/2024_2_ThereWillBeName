package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	"2024_2_ThereWillBeName/internal/pkg/user"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Credentials struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Handler struct {
	usecase user.UserUsecase
	jwt     jwt.JWTInterface
	logger  *slog.Logger
}

func NewUserHandler(usecase user.UserUsecase, jwt jwt.JWTInterface, logger *slog.Logger) *Handler {
	return &Handler{
		usecase: usecase,
		jwt:     jwt,
		logger:  logger,
	}
}

// SignUp godoc
// @Summary Sign up a new user
// @Description Create a new user with login and password
// @Accept json
// @Produce json
// @Param credentials body Credentials true "User credentials"
// @Success 201 {object} models.User "User created successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Bad Request"
// @Failure 500 {object} httpresponses.ErrorResponse "Internal Server Error"
// @Router /signup [post]
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for sign up")

	var credentials Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		h.logger.Warn("Failed to decode credentials", slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	user := models.User{
		Login:    credentials.Login,
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	var err error
	user.ID, err = h.usecase.SignUp(context.Background(), user)
	if err != nil {
		if errors.Is(err, models.ErrAlreadyExists) {
			h.logger.Warn("User already exists", slog.String("login", user.Login), slog.String("email", user.Email))
			response := httpresponse.ErrorResponse{
				Message: "user already exists",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusConflict, h.logger)
			return
		}

		h.logger.Error("Failed to sign up user", slog.String("error", err.Error()))

		response := httpresponse.ErrorResponse{
			Message: "Registration failed",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError, h.logger)
		return
	}

	h.logger.Debug("User signed up successfully", slog.Int("userID", int(user.ID)), slog.String("login", user.Login))

	token, err := h.jwt.GenerateToken(user.ID, user.Email, user.Login)
	if err != nil {
		h.logger.Error("Token generation failed", slog.String("userID", strconv.Itoa(int(user.ID))), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Token generation failed",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError, h.logger)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})

	h.logger.Debug("Token generated and set as cookie", slog.String("userID", strconv.Itoa(int(user.ID))), slog.String("login", user.Login), slog.String("email", user.Email))

	response := models.User{
		ID:    user.ID,
		Login: user.Login,
		Email: user.Email,
	}
	h.logger.DebugContext(logCtx, "Sign-up request completed successfully")

	httpresponse.SendJSONResponse(w, response, http.StatusOK, h.logger)
}

// Login godoc
// @Summary Login a user
// @Description Authenticate a user and return a token
// @Accept json
// @Produce json
// @Param credentials body Credentials true "User credentials"
// @Success 200 {string} string "Token"
// @Failure 400 {object} httpresponses.ErrorResponse "Bad Request"
// @Failure 401 {object} httpresponses.ErrorResponse "Unauthorized"
// @Router /login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling request for log in")

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		h.logger.Warn("Failed to decode credentials", slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	user, err := h.usecase.Login(context.Background(), credentials.Email, credentials.Password)
	if err != nil {
		h.logger.Warn("Login failed: invalid email or password", slog.String("email", credentials.Email))

		response := httpresponse.ErrorResponse{
			Message: "Invalid email or password",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	h.logger.Debug("User logged in successfully", slog.Int("userID", int(user.ID)), slog.String("email", user.Email))

	token, err := h.jwt.GenerateToken(user.ID, user.Email, user.Login)
	if err != nil {
		h.logger.Error("Token generation failed", slog.String("userID", strconv.Itoa(int(user.ID))), slog.String("error", err.Error()))
		response := httpresponse.ErrorResponse{
			Message: "Token generation failed",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError, h.logger)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})

	h.logger.Debug("Token generated and set as cookie", slog.String("userID", strconv.Itoa(int(user.ID))), slog.String("login", user.Login), slog.String("email", user.Email))

	response := models.User{
		ID:    user.ID,
		Login: user.Login,
		Email: user.Email,
	}
	h.logger.DebugContext(logCtx, "Sign-up request completed successfully")

	httpresponse.SendJSONResponse(w, response, http.StatusOK, h.logger)

}

// Logout godoc
// @Summary Logout a user
// @Description Log out the user by clearing the authentication token
// @Produce json
// @Success 200 {string} string "Logged out successfully"
// @Router /logout [post]
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Handling logout request")

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   -1,
	})

	h.logger.DebugContext(logCtx, "User logged out successfully")

	w.WriteHeader(http.StatusOK)
}

// CurrentUser godoc
// @Summary Get the current user
// @Description Retrieve the current authenticated user information
// @Produce json
// @Success 200 {object} models.User "Current user"
// @Failure 401 {object} httpresponses.ErrorResponse "Unauthorized"
// @Failure 500 {object} httpresponses.ErrorResponse "Internal Server Error"
// @Router /users/me [get]
func (h *Handler) CurrentUser(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Fetching current user information")

	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.Warn("Failed to retrieve user ID from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	login, ok := r.Context().Value(middleware.LoginKey).(string)
	if !ok {
		h.logger.Warn("Failed to retrieve user login from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	email, ok := r.Context().Value(middleware.EmailKey).(string)
	if !ok {
		h.logger.Warn("Failed to retrieve user email from context")

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	response := models.User{
		ID:    userID,
		Login: login,
		Email: email,
	}
	h.logger.DebugContext(logCtx, "Successfully retrieved current user information", "userID", userID)

	httpresponse.SendJSONResponse(w, response, http.StatusOK, h.logger)
}

// UploadAvatar godoc
// @Summary Upload user avatar
// @Description Upload an avatar image for the user
// @Accept multipart/form-data
// @Produce json
// @Param avatar formData file true "Avatar file"
// @Success 200 {string} string "Avatar uploaded successfully"
// @Failure 400 {object} httpresponses.ErrorResponse "Bad Request"
// @Failure 401 {object} httpresponses.ErrorResponse "Unauthorized"
// @Failure 500 {object} httpresponses.ErrorResponse "Internal Server Error"
// @Router /users/{userID}/avatar [put]
func (h *Handler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Starting avatar upload process")

	userIDStr := mux.Vars(r)["userID"]
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		h.logger.Error("Invalid user ID format", "userID", userIDStr, "error", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid user ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return

	}

	authUserID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok || authUserID != uint(userID) {
		h.logger.Warn("Unauthorized access attempt", "authUserID", authUserID, "targetUserID", userID)

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized to upload avatar for this ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.logger.Error("File size exceeds limit", "error", err)

		response := httpresponse.ErrorResponse{
			Message: "File is too large",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	file, header, err := r.FormFile("avatar")
	if err != nil {
		h.logger.Error("Error reading avatar file", "error", err)

		response := httpresponse.ErrorResponse{
			Message: "Invalid file upload",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}
	defer file.Close()

	h.logger.Debug("Uploading avatar", "userID", userID, "fileName", header.Filename)

	var avatarPath string
	if avatarPath, err = h.usecase.UploadAvatar(context.Background(), uint(userID), file, header); err != nil {
		h.logger.Error("Failed to upload avatar", "userID", userID, "error", err)

		response := httpresponse.ErrorResponse{
			Message: "Failed to upload avatar",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError, h.logger)
	}

	response := map[string]string{
		"message":    "Avatar uploaded successfully",
		"avatarPath": avatarPath,
	}

	h.logger.DebugContext(logCtx, "Avatar uploaded successfully", "userID", userID, "avatarPath", avatarPath)

	httpresponse.SendJSONResponse(w, response, http.StatusOK, h.logger)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Retrieve the user profile information
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {object} models.UserProfile "User profile"
// @Failure 401 {object} httpresponses.ErrorResponse "Unauthorized"
// @Failure 404 {object} httpresponses.ErrorResponse "Not Found"
// @Failure 500 {object} httpresponses.ErrorResponse "Internal Server Error"
// @Router /users/{userID}/profile [get]
func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	logCtx := log.LogRequestStart(r.Context(), r.Method, r.RequestURI)
	h.logger.DebugContext(logCtx, "Starting user profile retrieval")

	userIDStr := mux.Vars(r)["userID"]
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		h.logger.Error("Invalid user ID format", "userID", userIDStr, "error", err)

		response := httpresponse.ErrorResponse{
			Message: "Invalid user ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	requesterID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		h.logger.Warn("Unauthorized access attempt", "userID", userID)

		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	h.logger.Debug("Fetching profile", "userID", userID, "requesterID", requesterID)

	profile, err := h.usecase.GetProfile(context.Background(), uint(userID), requesterID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			h.logger.Warn("User not found", "userID", userID)

			response := httpresponse.ErrorResponse{
				Message: "User not found",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusNotFound, h.logger)
			return
		}

		h.logger.Error("Error retrieving profile", "userID", userID, "error", err)

		response := httpresponse.ErrorResponse{
			Message: "Failed to retrieve user profile",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError, h.logger)
		return
	}
	h.logger.Debug("User profile retrieved successfully", "userID", userID)

	httpresponse.SendJSONResponse(w, profile, http.StatusOK, h.logger)
}

func (h *Handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	login, ok := r.Context().Value(middleware.LoginKey).(string)
	if !ok {
		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	email, ok := r.Context().Value(middleware.EmailKey).(string)
	if !ok {
		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized, h.logger)
		return
	}

	var credentials struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
		return
	}

	user := models.User{
		ID:       userID,
		Login:    login,
		Email:    email,
		Password: credentials.OldPassword,
	}

	err := h.usecase.UpdatePassword(r.Context(), user, credentials.NewPassword)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			response := httpresponse.ErrorResponse{
				Message: "User not found",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusNotFound, h.logger)
			return
		} else if errors.Is(err, models.ErrMismatch) {
			response := httpresponse.ErrorResponse{
				Message: "Invalid old password",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusBadRequest, h.logger)
			return
		}

		response := httpresponse.ErrorResponse{
			Message: "Failed to update password",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError, h.logger)
		return
	}

	response := struct {
		ID       uint   `json:"id"`
		Password string `json:"password"`
	}{
		ID:       user.ID,
		Password: credentials.NewPassword,
	}

	httpresponse.SendJSONResponse(w, response, http.StatusOK, h.logger)
}
