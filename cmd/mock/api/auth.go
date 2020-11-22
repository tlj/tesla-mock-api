package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
	"github.com/tlj/tesla"
	"net/http"
	"time"
)

type AuthRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    int64  `json:"created_at"`
}

var (
	AccessToken  = "nkJndyWujACm4aeZ5njFcZeen3bObu1a"
	RefreshToken = "h1FTKMFbrUUYAOXxsCWV5APYWb0dt2N0"

	authEmail    = "mock@user"
	authPassword = "mock"
)

func generateTokens() {
	AccessToken = randstr.String(32)
	RefreshToken = randstr.String(32)
}

func AuthRequired(c *gin.Context) {
	bearer := c.Request.Header.Get("Authorization")
	if bearer == "" || bearer != "Bearer " + AccessToken {
		c.AbortWithStatusJSON(http.StatusUnauthorized, nil)
		return
	}

	c.Next()
}

func AuthHandler(c *gin.Context) {
	var input AuthRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.GrantType == "bearer" {
		if input.Email != authEmail || input.Password != authPassword {
			c.JSON(http.StatusUnauthorized, tesla.StringResponse{Response: "https://myt-websvc-vip1.tesla.com:443/teslamobileservices/users/authenticate_=>_authorization_required"})
			return
		}

		generateTokens()
	} else if input.GrantType == "refresh_token" {
		if input.RefreshToken != RefreshToken {
			c.JSON(http.StatusUnauthorized, tesla.StringResponse{Response: "https://myt-websvc-vip1.tesla.com:443/teslamobileservices/users/authenticate_=>_authorization_required"})
			return
		}

		generateTokens()
	} else {
		c.JSON(http.StatusBadRequest, tesla.StringResponse{Response: "invalid grant type"})
		return
	}

	resp := AuthResponse{
		AccessToken:  AccessToken,
		TokenType:    "bearer",
		RefreshToken: RefreshToken,
		CreatedAt:    (time.Now()).Unix(),
		ExpiresIn:    3888000,
	}

	c.JSON(http.StatusOK, resp)
}
