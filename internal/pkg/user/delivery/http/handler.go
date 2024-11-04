package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	"2024_2_ThereWillBeName/internal/pkg/user"
	"context"
	"encoding/json"
	"errors"
	"log"
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
	jwt     *jwt.JWT
}

func NewUserHandler(usecase user.UserUsecase, jwt *jwt.JWT) *Handler {
	return &Handler{
		usecase: usecase,
		jwt:     jwt,
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
	var credentials Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
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
			response := httpresponse.ErrorResponse{
				Message: "user already exists",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusConflict)
			return
		}
		response := httpresponse.ErrorResponse{
			Message: "Registration failed",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		return
	}

	token, err := h.jwt.GenerateToken(user.ID, user.Email, user.Login)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Token generation failed",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})

	response := models.User{
		ID:    user.ID,
		Login: user.Login,
		Email: user.Email,
	}

	httpresponse.SendJSONResponse(w, response, http.StatusOK)
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
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	user, err := h.usecase.Login(context.Background(), credentials.Email, credentials.Password)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid email or password",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized)
		return
	}

	token, err := h.jwt.GenerateToken(user.ID, user.Email, user.Login)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Token generation failed",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})

	response := models.User{
		ID:    user.ID,
		Login: user.Login,
		Email: user.Email,
	}

	httpresponse.SendJSONResponse(w, response, http.StatusOK)

}

// Logout godoc
// @Summary Logout a user
// @Description Log out the user by clearing the authentication token
// @Produce json
// @Success 200 {string} string "Logged out successfully"
// @Router /logout [post]
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   -1,
	})

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
	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized)
		return
	}

	login, ok := r.Context().Value(middleware.LoginKey).(string)
	if !ok {
		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized)
		return
	}

	response := models.User{
		ID:    userID,
		Login: login,
	}

	httpresponse.SendJSONResponse(w, response, http.StatusOK)
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
	userIDStr := mux.Vars(r)["userID"]
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		log.Printf("upload error: %s", err)
		response := httpresponse.ErrorResponse{
			Message: "Invalid user ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return

	}

	authUserID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok || authUserID != uint(userID) {
		response := httpresponse.ErrorResponse{
			Message: "User is not authorized to upload avatar for this ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		response := httpresponse.ErrorResponse{
			Message: "File is too large",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("avatar")
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid file upload",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}
	defer file.Close()

	var avatarPath string
	if avatarPath, err = h.usecase.UploadAvatar(context.Background(), uint(userID), file, header); err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Failed to upload avatar",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
	}

	response := map[string]string{
		"message":    "Avatar uploaded successfully",
		"avatarPath": avatarPath,
	}
	httpresponse.SendJSONResponse(w, response, http.StatusOK)

	w.WriteHeader(http.StatusOK)
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
	userIDStr := mux.Vars(r)["userID"]
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		response := httpresponse.ErrorResponse{
			Message: "Invalid user ID",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	requesterID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		response := httpresponse.ErrorResponse{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusUnauthorized)
		return
	}

	profile, err := h.usecase.GetProfile(context.Background(), uint(userID), requesterID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			response := httpresponse.ErrorResponse{
				Message: "User not found",
			}
			httpresponse.SendJSONResponse(w, response, http.StatusNotFound)
			return
		}
		response := httpresponse.ErrorResponse{
			Message: "Failed to retrieve user profile",
		}
		httpresponse.SendJSONResponse(w, response, http.StatusInternalServerError)
		return
	}

	httpresponse.SendJSONResponse(w, profile, http.StatusOK)
}
