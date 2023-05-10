package user

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c *gin.Context) {
	
	var user LoginUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	

	u, err := h.Service.Login(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expiration := 60 * 60 * 24

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    u.AccessToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(expiration)),
		MaxAge:   expiration,
		HttpOnly: true, 
		Secure:   true, 
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, u)
}

func (h *Handler) GetUserInfo(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	cookie, err := c.Request.Cookie("jwt")

     if err != nil {
		fmt.Println("cookie",cookie)
	 } else {
		fmt.Println("error", err)
	 }

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized1"})
		return
	}

	bearerToken := strings.Split(authHeader, "Bearer ")
	if len(bearerToken) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized2"})
		return
	}

	jwtToken := bearerToken[1]

	// Validate and parse the JWT token.
	user, err := h.Service.ParseToken(c.Request.Context(), jwtToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the user information.
	c.JSON(http.StatusOK, user)
}

func (h *Handler) Logout(c *gin.Context) {

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true, 
		Secure:   true, 
		SameSite: http.SameSiteLaxMode, 
	}


	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
