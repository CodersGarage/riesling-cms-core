package api

import (
	"net/http"
	"github.com/s4kibs4mi/govalidator"
	"encoding/json"
	"riesling-cms-core/app/data"
	"riesling-cms-core/app/utils"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := data.User{}
	rules := govalidator.MapData{
		"name":  []string{"required", "between:3,30"},
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
		switch user.Count() {
		case 0:
			user.Level = data.USER_LEVEL_ADMIN
		default:
			user.Level = data.USER_LEVEL_MEMBER
		}
		if len(user.Password) >= 8 && len(user.Password) <= 30 {
			if !user.IsEmailExists() {
				user.Hash = utils.GetUUID()
				if user.Hash != "" && user.Save() {
					resp := APIResponse{
						Code:    http.StatusOK,
						Message: "User has been created.",
						Data:    user,
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
				Code:    http.StatusConflict,
				Message: "Email address exists.",
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	if len(pathParams) > 0 {
		hash := pathParams[HASH]
		user := data.User{}
		if user.Get(hash) {
			user.Password = ""
			resp := APIResponse{
				Code: http.StatusOK,
				Data: user,
			}
			ServeAsJSON(resp, w)
			return
		}
		resp := APIResponse{
			Code:    http.StatusNotFound,
			Message: "User not found.",
		}
		ServeAsJSON(resp, w)
		return
	}
	resp := APIResponse{
		Code:    http.StatusBadRequest,
		Message: "User hash required.",
	}
	ServeAsJSON(resp, w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := data.User{}
	json.NewDecoder(r.Body).Decode(&user)
	params := mux.Vars(r)
	if len(params) > 0 {
		hash := params["hash"]
		isUpdateOk, u := user.Update(hash)
		if isUpdateOk {
			u.Password = ""
			resp := APIResponse{
				Code:    http.StatusOK,
				Message: "User has been updated.",
				Data:    u,
			}
			ServeAsJSON(resp, w)
			return
		}
		resp := APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Couldn't update user.",
		}
		ServeAsJSON(resp, w)
		return
	}
	resp := APIResponse{
		Code:    http.StatusBadRequest,
		Message: "User hash required.",
	}
	ServeAsJSON(resp, w)
}
