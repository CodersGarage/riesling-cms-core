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
			if user.GetByEmailAndPassword(user.Email, user.Password) {
				session := data.Session{
					AccessToken:  utils.GetUUID(),
					RefreshToken: utils.GetUUID(),
					ExpireTime:   utils.GetExpireTime(),
					Hash:         user.Hash,
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
				Code:    http.StatusNotAcceptable,
				Message: "Email & Password mismatch.",
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

}

func ReCreateSession(w http.ResponseWriter, r *http.Request) {

}
