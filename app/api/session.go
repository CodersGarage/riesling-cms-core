package api

import (
	"net/http"
	"riesling-cms-core/app/data"
	"github.com/s4kibs4mi/govalidator"
	"riesling-cms-core/app/utils"
	"gopkg.in/mgo.v2/bson"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	user := data.User{}
	rules := govalidator.MapData{
		"email": []string{"required", "email"},
	}
	opts := govalidator.Options{
		Request: r,
		Rules:   rules,
		Data:    &user,
	}
	vr := govalidator.New(opts)
	err := vr.ValidateJSON()
	if len(err) == 0 {
		if len(user.Password) >= 8 && len(user.Password) <= 30 {
			preUser := data.User{}
			if preUser.GetByEmail(user.Email) {
				if utils.CompareHashedPassword(preUser.Password, user.Password) {
					session := data.Session{
						AccessToken:  utils.GetUUID(),
						RefreshToken: utils.GetUUID(),
						ExpireTime:   utils.GetExpireTime(),
						Hash:         preUser.Hash,
					}
					if session.Save() {
						resp := APIResponse{
							Code: http.StatusOK,
							Data: session,
						}
						ServeAsJSON(resp, w)
						return
					}
					resp := APIResponse{
						Code:    http.StatusInternalServerError,
						Message: "Something went wrong.",
					}
					ServeAsJSON(resp, w)
					return
				}
				resp := APIResponse{
					Code:    http.StatusInternalServerError,
					Message: "Email & Password mismatch.",
				}
				ServeAsJSON(resp, w)
				return
			}
			resp := APIResponse{
				Code:    http.StatusNotAcceptable,
				Message: "User not found",
			}
			ServeAsJSON(resp, w)
			return
		}
		resp := APIResponse{
			Code: http.StatusBadRequest,
			Error: bson.M{
				"password": []string{
					"The password field is required.",
					"The password field must be between 8-30 char.",
				},
			},
		}
		ServeAsJSON(resp, w)
		return
	}
	resp := APIResponse{
		Code:  http.StatusBadRequest,
		Error: err,
	}
	ServeAsJSON(resp, w)
}

func CheckSession(w http.ResponseWriter, r *http.Request) {
	hash := r.Header.Get("hash")
	user := data.User{}
	user.Get(hash)
	resp := APIResponse{
		Code: http.StatusOK,
		Data: ResponseValue{
			"is_valid": true,
			"level":    user.Level,
		},
	}
	ServeAsJSON(resp, w)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get(ACCESS_TOKEN)
	session := data.Session{}
	if session.Get(accessToken) && session.Delete() {
		resp := APIResponse{
			Code: http.StatusOK,
			Data: ResponseValue{
				"success": true,
			},
		}
		ServeAsJSON(resp, w)
		return
	}
	resp := APIResponse{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong.",
	}
	ServeAsJSON(resp, w)
}

func DeleteAllSession(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get(ACCESS_TOKEN)
	session := data.Session{}
	if session.Get(accessToken) && session.DeleteAll() {
		resp := APIResponse{
			Code: http.StatusOK,
			Data: ResponseValue{
				"success": true,
			},
		}
		ServeAsJSON(resp, w)
		return
	}
	resp := APIResponse{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong.",
	}
	ServeAsJSON(resp, w)
}

func ReCreateSession(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.Header.Get(REFRESH_TOKEN)
	session := data.Session{}
	if refreshToken != "" && session.GetByRefreshToken(refreshToken) {
		session.Delete()
		session = data.Session{
			AccessToken:  utils.GetUUID(),
			RefreshToken: utils.GetUUID(),
			ExpireTime:   utils.GetExpireTime(),
			Hash:         session.Hash,
		}
		if session.Save() {
			resp := APIResponse{
				Code: http.StatusOK,
				Data: session,
			}
			ServeAsJSON(resp, w)
			return
		}
	}
	resp := APIResponse{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong.",
	}
	ServeAsJSON(resp, w)
}
