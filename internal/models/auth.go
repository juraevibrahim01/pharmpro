package models

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// -------------------- login -------------------
type Auth_Req_login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Auth_Res_login struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Error       string `json:"error"`
	Step        string `json:"step"`
}

// ----------------------------------------------------

// --------------------- otp -------------------------
type Auth_req_otp struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type Auth_res_otp struct {
	Status       string `json:"status"`
	Description  string `json:"description"`
	Error        string `json:"error"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"resresh_token"`
}

var JWTSecret = []byte("UThdfdlhrrhr353")
var JWTSecretRef = []byte("Fsfdsshsrrhrhrh")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// ---------------------------------------------------

// -------------------------------- ref_token --------------------------
type Auth_req_ref_token struct {
	RefToken string `json:"ref_token"`
}

type Auth_res_ref_token struct {
	Status   string `json:"status"`
	Error    string `json:"error"`
	Token    string `json:"token"`
	RefToken string `json:"ref_token"`
}

// --------------------------------------------------------------------

// ----------------------------- Errors ----------------------------------
var ErrTokenInvalid = errors.New("Токен недействителен")
var ErrTokenExpired = errors.New("Токен истек")

// -----------------------------------------------------------------------
