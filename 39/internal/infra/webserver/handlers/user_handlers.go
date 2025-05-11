package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Bran00/go.expert/39/internal/dto"
	"github.com/Bran00/go.expert/39/internal/entity"
	"github.com/Bran00/go.expert/39/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB       database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB:       userDB,
	}
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body     dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500  {object}  Error
// @Router       /users [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwt_expires_in").(int)

	fmt.Printf("jwt: %v, jwtExpiresIn: %d\n", jwt, jwtExpiresIn)
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	acessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acessToken)
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request  body   dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure			 500  {object}  Error
// @Router			 /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return


	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
