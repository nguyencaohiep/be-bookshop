package controller

import (
	"encoding/json"
	"net/http"
	"soc1-service/pkg/auth"
	"soc1-service/pkg/log"
	"soc1-service/pkg/router"
	"soc1-service/service/soc1/dao"

	"golang.org/x/crypto/bcrypt"
)

const _defaultPasswordGenerateCost = 12

type AccountForm struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var registerForm = &AccountForm{}
	err := json.NewDecoder(r.Body).Decode(registerForm)
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateAccount: Decode(registerForm)", err)
		router.ResponseInternalError(w, "decode registerForm failed", err)

		return
	}

	account := &dao.AccountDAO{
		Email: registerForm.Email,
	}

	exist, err := account.CheckAccountExist()
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateAccount: account.CheckAccountExist()", err)
		router.ResponseInternalError(w, "Error", err)
		return
	}

	if exist {
		router.ResponseBadRequest(w, "B.400", "Account exist!")
		return
	}

	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerForm.Password), _defaultPasswordGenerateCost)
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateAccount: GenerateFromPassword", err)
		router.ResponseInternalError(w, "Error", err)
		return
	}
	account.Password = string(hashedPassword)

	err = account.InsertAccount()
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateAccount: account.InsertAccount()", err)
		router.ResponseInternalError(w, "Error", err)
		return
	}

	router.ResponseSuccess(w, "B.200", "Signup successfully")
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var signinForm = &AccountForm{}
	err := json.NewDecoder(r.Body).Decode(signinForm)
	if err != nil {
		log.Println(log.LogLevelDebug, "SignIn: Decode(signupForm)", err)
		router.ResponseInternalError(w, "decode signin failed", err)

		return
	}

	var account = &dao.AccountDAO{
		Email: signinForm.Email,
	}

	exist, err := account.CheckAccountExist()
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateAccount: account.CheckAccountExist()", err)
		router.ResponseInternalError(w, "Error", err)
		return
	}

	if !exist {
		router.ResponseBadRequest(w, "B.400", "Account not exist!")
		return
	}

	err = account.GetAccountByEmail()
	if err != nil {
		log.Println(log.LogLevelDebug, "SignIn: GetAccount", err)
		router.ResponseInternalError(w, "get account failed", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(signinForm.Password))
	if err != nil {
		log.Println(log.LogLevelDebug, "SignIn: bcrypt.CompareHashAndPassword", err)
		router.ResponseBadRequest(w, "B.400", "Wrong email or password")
		return
	}

	var payload = &auth.Payload{
		Email: account.Email,
	}

	tokenData, err := payload.GetTokenDataJWT()
	if err != nil {
		log.Println(log.LogLevelDebug, "Signin: GetTokenDataJWT", err)
		router.ResponseInternalError(w, "get token failed", err)
	}

	router.ResponseSuccessWithData(w, "B.200", "Signin successfully", tokenData, account)
}
