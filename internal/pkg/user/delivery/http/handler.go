package http

import (
	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"2024_2_ThereWillBeName/internal/pkg/jwt"
	log "2024_2_ThereWillBeName/internal/pkg/logger"
	"2024_2_ThereWillBeName/internal/pkg/middleware"
	"2024_2_ThereWillBeName/internal/pkg/user/delivery/grpc/gen"
	"2024_2_ThereWillBeName/internal/validator"

	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"mime"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Credentials struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Handler struct {
	client gen.UserServiceClient
	jwt    jwt.JWTInterface
	logger *slog.Logger
}

func NewUserHandler(client gen.UserServiceClient, jwt jwt.JWTInterface, logger *slog.Logger) *Handler {
	return &Handler{
		client: client,
		jwt:    jwt,
		logger: logger,
	}
}

// SignUp godoc
// @Summary Sign up a new user
// @Description Create a new user with login and password
// @Accept json
// @Produce json
// @Param credentials body Credentials true "User credentials"
// @Success 201 {object} models.User "User created successfully"
// @Failure 400 {object} httpresponses.Response "Bad Request"
// @Failure 500 {object} httpresponses.Response "Internal Server Error"
// @Router /signup [post]
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Handling request for sign up")

	var credentials Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		h.logger.WarnContext(logCtx, "Failed to decode credentials", slog.String("error", err.Error()))
		response := httpresponse.Response{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	credentials.Login = template.HTMLEscapeString(credentials.Login)
	credentials.Email = template.HTMLEscapeString(credentials.Email)
	credentials.Password = template.HTMLEscapeString(credentials.Password)

	user := models.User{
		Login:    credentials.Login,
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	logCtx = log.AppendCtx(logCtx, slog.String("login", user.Login))
	logCtx = log.AppendCtx(logCtx, slog.String("email", user.Email))

	v := validator.New()
	if models.ValidateUser(v, &user); !v.Valid() {
		h.logger.WarnContext(logCtx, "User data is not valid")
		httpresponse.SendJSONResponse(logCtx, w, nil, http.StatusUnprocessableEntity, h.logger)
		return
	}

	signUpRequest := &gen.SignUpRequest{
		Login:    user.Login,
		Email:    user.Email,
		Password: user.Password,
	}

	var err error
	signupResponse, err := h.client.SignUp(r.Context(), signUpRequest)
	if err != nil {
		if errors.Is(err, models.ErrAlreadyExists) {
			h.logger.WarnContext(logCtx, "User already exists")
			response := httpresponse.Response{
				Message: "user already exists",
			}
			httpresponse.SendJSONResponse(logCtx, w, response, http.StatusConflict, h.logger)
			return
		}

		h.logger.ErrorContext(logCtx, "Failed to sign up user", slog.String("error", err.Error()))

		response := httpresponse.Response{
			Message: "Registration failed",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusInternalServerError, h.logger)
		return
	}

	user.ID = uint(signupResponse.Id)

	h.logger.DebugContext(logCtx, "User signed up successfully", slog.Int("userID", int(user.ID)))

	token, err := h.jwt.GenerateToken(user.ID, user.Email, user.Login)
	if err != nil {
		h.logger.ErrorContext(logCtx, "Token generation failed", slog.String("userID", strconv.Itoa(int(user.ID))), slog.String("error", err.Error()))
		response := httpresponse.Response{
			Message: "Token generation failed",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusInternalServerError, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Token generated")

	response := models.UserResponseWithToken{
		User: models.User{
			ID:    user.ID,
			Login: user.Login,
			Email: user.Email,
		},
		Token: token,
	}

	h.logger.DebugContext(logCtx, "Sign-up request completed successfully")

	httpresponse.SendJSONResponse(logCtx, w, response, http.StatusOK, h.logger)
}

// Login godoc
// @Summary Login a user
// @Description Authenticate a user and return a token
// @Accept json
// @Produce json
// @Param credentials body Credentials true "User credentials"
// @Success 200 {string} string "Token"
// @Failure 400 {object} httpresponses.Response "Bad Request"
// @Failure 401 {object} httpresponses.Response "Unauthorized"
// @Router /login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Handling request for log in")

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		h.logger.WarnContext(logCtx, "Failed to decode credentials", slog.String("error", err.Error()))
		response := httpresponse.Response{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	logCtx = log.AppendCtx(logCtx, slog.String("email", credentials.Email))

	credentials.Email = template.HTMLEscapeString(credentials.Email)
	credentials.Password = template.HTMLEscapeString(credentials.Password)

	loginRequest := &gen.LoginRequest{
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	loginResponse, err := h.client.Login(r.Context(), loginRequest)
	if err != nil {
		h.logger.ErrorContext(logCtx, "Login failed: invalid email or password")

		response := httpresponse.Response{
			Message: "Invalid email or password",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	user := models.User{
		ID:         uint(loginResponse.Id),
		Login:      loginResponse.Login,
		Email:      loginResponse.Email,
		AvatarPath: loginResponse.AvatarPath,
	}

	logCtx = log.AppendCtx(logCtx, slog.Int("userID", int(user.ID)))

	h.logger.DebugContext(logCtx, "User logged in successfully")

	token, err := h.jwt.GenerateToken(user.ID, user.Email, user.Login)
	if err != nil {
		h.logger.ErrorContext(logCtx, "Token generation failed", slog.String("error", err.Error()))
		response := httpresponse.Response{
			Message: "Token generation failed",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusInternalServerError, h.logger)
		return
	}
	h.logger.DebugContext(logCtx, "Token generated", slog.String("login", user.Login))

	response := models.UserResponseWithToken{
		User:  user,
		Token: token,
	}

	h.logger.DebugContext(logCtx, "Login request completed successfully")

	httpresponse.SendJSONResponse(logCtx, w, response, http.StatusOK, h.logger)

}

// Logout godoc
// @Summary Logout a user
// @Description Log out the user by clearing the authentication token
// @Produce json
// @Success 200 {string} string "Logged out successfully"
// @Router /logout [post]
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Handling logout request")

	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "User logged out successfully")

	w.WriteHeader(http.StatusOK)
}

// CurrentUser godoc
// @Summary Get the current user
// @Description Retrieve the current authenticated user information
// @Produce json
// @Success 200 {object} models.User "Current user"
// @Failure 401 {object} httpresponses.Response "Unauthorized"
// @Failure 500 {object} httpresponses.Response "Internal Server Error"
// @Router /users/me [get]
func (h *Handler) CurrentUser(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Fetching current user information")

	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}
	getProfileRequest := &gen.GetProfileRequest{
		Id:          uint32(userID),
		RequesterId: uint32(userID),
	}

	getProfileResponse, err := h.client.GetProfile(r.Context(), getProfileRequest)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			h.logger.ErrorContext(logCtx, "User not found")

			response := httpresponse.Response{
				Message: "User not found",
			}
			httpresponse.SendJSONResponse(logCtx, w, response, http.StatusNotFound, h.logger)
			return
		}

		h.logger.ErrorContext(logCtx, "Error retrieving profile", "error", err)

		response := httpresponse.Response{
			Message: "Failed to retrieve user current user information",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusInternalServerError, h.logger)
		return
	}

	userProfile := models.UserProfile{
		Login:      getProfileResponse.Login,
		AvatarPath: getProfileResponse.AvatarPath,
		Email:      getProfileResponse.Email,
	}

	userResponse := models.UserResponse{
		ID: uint32(userID), Profile: userProfile,
	}

	h.logger.DebugContext(logCtx, "Successfully retrieved current user information")
	httpresponse.SendJSONResponse(logCtx, w, userResponse, http.StatusOK, h.logger)
}

// UploadAvatar godoc
// @Summary Upload user avatar
// @Description Upload an avatar image for the user
// @Accept multipart/form-data
// @Produce json
// @Param avatar formData file true "Avatar file"
// @Success 200 {string} string "Avatar uploaded successfully"
// @Failure 400 {object} httpresponses.Response "Bad Request"
// @Failure 401 {object} httpresponses.Response "Unauthorized"
// @Failure 500 {object} httpresponses.Response "Internal Server Error"
// @Router /users/{userID}/avatar [put]
func (h *Handler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Starting avatar upload process")

	userIDStr := mux.Vars(r)["userID"]
	userID, err := strconv.ParseUint(userIDStr, 10, 32)

	logCtx = log.AppendCtx(logCtx, slog.Int("user_id", int(userID)))
	if err != nil {
		h.logger.WarnContext(logCtx, "Invalid user ID format", "error", err)
		response := httpresponse.Response{
			Message: "Invalid user ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return

	}

	authUserID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok || authUserID != uint(userID) {
		h.logger.WarnContext(logCtx, "Unauthorized access attempt", "authUserID", authUserID)

		response := httpresponse.Response{
			Message: "User is not authorized to upload avatar for this ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	var requestData struct {
		Avatar string `json:"avatar"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		h.logger.WarnContext(logCtx, "Invalid JSON format", "error", err)
		response := httpresponse.Response{
			Message: "Invalid request format",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	if strings.HasPrefix(requestData.Avatar, "data:image/") {
		index := strings.Index(requestData.Avatar, ",")
		if index != -1 {
			requestData.Avatar = requestData.Avatar[index+1:]
		} else {
			h.logger.Error("Invalid base64 image format", "error", "missing ',' separator")
			response := httpresponse.Response{
				Message: "Invalid base64 image format",
			}
			httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
			return
		}
	}

	avatarData, err := base64.StdEncoding.DecodeString(requestData.Avatar)
	if err != nil {
		h.logger.Error("Failed to decode base64 image", "error", err)
		response := httpresponse.Response{
			Message: "Invalid base64 image data",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	fileType := http.DetectContentType(avatarData)
	h.logger.DebugContext(logCtx, "Detected file type", "fileType", fileType)

	if !strings.HasPrefix(fileType, "image/") {
		h.logger.ErrorContext(logCtx, "Invalid file type", "fileType", fileType)
		response := httpresponse.Response{
			Message: "Only image files are allowed",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	ext, err := mime.ExtensionsByType(fileType)
	if err != nil || len(ext) == 0 {
		h.logger.ErrorContext(logCtx, "Unable to determine file extension", "mimeType", fileType)
		response := httpresponse.Response{
			Message: "Unable to determine file extension",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	avatarFileName := fmt.Sprintf("user_%d_avatar%s", userID, ext[0])

	h.logger.DebugContext(logCtx, "Uploading avatar", "avatarFileName", avatarFileName)

	uploadAvatarRequest := &gen.UploadAvatarRequest{
		Id:             uint32(userID),
		AvatarData:     avatarData,
		AvatarFileName: avatarFileName,
	}

	uploadAvatarResponse, err := h.client.UploadAvatar(r.Context(), uploadAvatarRequest)
	if err != nil {
		h.logger.ErrorContext(logCtx, "Failed to upload avatar", "error", err)
		response := httpresponse.Response{
			Message: "Failed to upload avatar",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusInternalServerError, h.logger)
		return
	}

	photoResponce := models.Photo{
		Path: uploadAvatarResponse.AvatarPath,
	}

	h.logger.DebugContext(logCtx, "Avatar uploaded successfully", "avatarPath", uploadAvatarResponse.AvatarPath)

	httpresponse.SendJSONResponse(logCtx, w, photoResponce, http.StatusOK, h.logger)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Retrieve the user profile information
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {object} models.UserProfile "User profile"
// @Failure 401 {object} httpresponses.Response "Unauthorized"
// @Failure 404 {object} httpresponses.Response "Not Found"
// @Failure 500 {object} httpresponses.Response "Internal Server Error"
// @Router /users/{userID}/profile [get]
func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	userIDStr := mux.Vars(r)["userID"]
	logCtx = log.AppendCtx(logCtx, slog.String("userID", userIDStr))
	h.logger.DebugContext(logCtx, "Handling request for retrieval user's profile")

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		h.logger.WarnContext(logCtx, "Invalid user ID format", "error", err)

		response := httpresponse.Response{
			Message: "Invalid user ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	requesterID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		h.logger.WarnContext(logCtx, "Unauthorized access attempt")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "Fetching profile", "requesterID", requesterID)

	getProfileRequest := &gen.GetProfileRequest{
		Id:          uint32(userID),
		RequesterId: uint32(requesterID),
	}

	getProfileResponse, err := h.client.GetProfile(r.Context(), getProfileRequest)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			h.logger.ErrorContext(logCtx, "User not found")

			response := httpresponse.Response{
				Message: "User not found",
			}
			httpresponse.SendJSONResponse(logCtx, w, response, http.StatusNotFound, h.logger)
			return
		}

		h.logger.ErrorContext(logCtx, "Error retrieving profile", "error", err)

		response := httpresponse.Response{
			Message: "Failed to retrieve user profile",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusInternalServerError, h.logger)
		return
	}

	userProfile := models.UserProfile{
		Login:      getProfileResponse.Login,
		Email:      getProfileResponse.Email,
		AvatarPath: getProfileResponse.AvatarPath,
	}

	h.logger.DebugContext(logCtx, "User profile retrieved successfully")
	httpresponse.SendJSONResponse(logCtx, w, userProfile, http.StatusOK, h.logger)
}

func (h *Handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Starting user password updating")
	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}
	logCtx = log.AppendCtx(logCtx, slog.Int("user_id", int(userID)))

	login, ok := r.Context().Value(middleware.LoginKey).(string)
	if !ok {
		h.logger.WarnContext(logCtx, "Failed to retrieve login from context")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	email, ok := r.Context().Value(middleware.EmailKey).(string)
	if !ok {
		h.logger.WarnContext(logCtx, "Failed to retrieve email from context")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	var credentials struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		response := httpresponse.Response{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "updating password", "userID", userID, "oldPassword", credentials.OldPassword, "newPassword", credentials.NewPassword)

	updatePasswordRequest := &gen.UpdatePasswordRequest{
		Id:          uint32(userID),
		Login:       login,
		Email:       email,
		OldPassword: credentials.OldPassword,
		NewPassword: credentials.NewPassword,
	}

	_, err := h.client.UpdatePassword(r.Context(), updatePasswordRequest)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			h.logger.ErrorContext(logCtx, "User not found")

			response := httpresponse.Response{
				Message: "User not found",
			}
			httpresponse.SendJSONResponse(logCtx, w, response, http.StatusNotFound, h.logger)
			return
		} else if errors.Is(err, models.ErrMismatch) {
			h.logger.ErrorContext(logCtx, "Passwords mismatch")

			response := httpresponse.Response{
				Message: "Invalid old password",
			}
			httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
			return
		}
		h.logger.ErrorContext(logCtx, "Failed to update password", slog.Any("error", err.Error()))

		response := httpresponse.Response{
			Message: "Failed to update password",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusInternalServerError, h.logger)
		return
	}

	response := models.ResponseWithId{
		ID:      userID,
		Message: "User's password updated successfully",
	}

	h.logger.DebugContext(logCtx, "User password updated successfully")

	httpresponse.SendJSONResponse(logCtx, w, response, http.StatusOK, h.logger)
}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()
	h.logger.DebugContext(logCtx, "Starting user profile updating")

	userID, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {
		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}
	logCtx = log.AppendCtx(logCtx, slog.Int("user_id", int(userID)))

	var userData struct {
		Login string `json:"username"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		response := httpresponse.Response{
			Message: "Invalid request",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "updating profile", "username", userData.Login, "email", userData.Email)

	updateProfileRequest := &gen.UpdateProfileRequest{
		UserId:   uint32(userID),
		Username: userData.Login,
		Email:    userData.Email,
	}

	_, err := h.client.UpdateProfile(r.Context(), updateProfileRequest)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			h.logger.ErrorContext(logCtx, "User not found")

			response := httpresponse.Response{
				Message: "User not found",
			}
			httpresponse.SendJSONResponse(logCtx, w, response, http.StatusNotFound, h.logger)
			return
		}
		h.logger.ErrorContext(logCtx, "User not found", slog.Any("error", err.Error()))

		response := httpresponse.Response{
			Message: "Failed to update profile",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusInternalServerError, h.logger)
		return
	}

	h.logger.DebugContext(logCtx, "User profile updated successfully")

	response := models.Response{
		Username: userData.Login,
		Email:    userData.Email,
	}
	httpresponse.SendJSONResponse(logCtx, w, response, http.StatusOK, h.logger)
}

func (h *Handler) GetAchievements(w http.ResponseWriter, r *http.Request) {
	logCtx := r.Context()

	vars := mux.Vars(r)
	userIDStr := vars["userID"]

	logCtx = log.AppendCtx(logCtx, slog.String("user_id", userIDStr))

	h.logger.DebugContext(logCtx, "Handling request for getting achivements by user ID")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		h.logger.WarnContext(logCtx, "Failed to parse user ID", slog.Any("error", err.Error()))
		response := httpresponse.Response{
			Message: "Invalid user ID",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusBadRequest, h.logger)
		return
	}
	_, ok := r.Context().Value(middleware.IdKey).(uint)
	if !ok {

		h.logger.WarnContext(logCtx, "Failed to retrieve user ID from context")

		response := httpresponse.Response{
			Message: "User is not authorized",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusUnauthorized, h.logger)
		return
	}

	achievements, err := h.client.GetAchievements(r.Context(), &gen.GetAchievementsRequest{Id: uint32(userID)})
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			h.logger.ErrorContext(logCtx, "achievements are not found")

			response := httpresponse.Response{
				Message: "Achievements are not found",
			}
			httpresponse.SendJSONResponse(logCtx, w, response, http.StatusNotFound, h.logger)
			return
		}

		h.logger.ErrorContext(logCtx, "Error retrieving achievements", "error", err)

		response := httpresponse.Response{
			Message: "Failed to retrieve user's achievements",
		}
		httpresponse.SendJSONResponse(logCtx, w, response, http.StatusInternalServerError, h.logger)
		return
	}

	achievementResponse := make(models.AchievementList, len(achievements.Achievements))
	for i, achievement := range achievements.Achievements {
		achievementResponse[i] = models.Achievement{
			ID:       uint(achievement.Id),
			Name:     achievement.Name,
			IconPath: achievement.IconPath,
		}
	}

	httpresponse.SendJSONResponse(logCtx, w, achievementResponse, http.StatusOK, h.logger)
}
