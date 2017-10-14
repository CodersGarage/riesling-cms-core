package api

import (
	"net/http"
	"github.com/s4kibs4mi/govalidator"
	"encoding/json"
	"riesling-cms-core/app/data"
	"riesling-cms-core/app/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := data.User{}
	rules := govalidator.MapData{
		"name":     []string{"required", "between:3,30"},
		"email":    []string{"required", "email"},
		"password": []string{"required", "between:8,20"},
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
		if !user.IsEmailExists() {
			user.Hash = utils.GetUUID()
			if user.Hash != "" && user.Save() {
				json.NewEncoder(w).Encode(APIResponse{
					Code:    http.StatusOK,
					Message: "User has been created.",
					Data:    user,
				})
				return
			}
			json.NewEncoder(w).Encode(APIResponse{
				Code:    http.StatusInternalServerError,
				Message: "Something went wrong.",
			})
			return
		}
		json.NewEncoder(w).Encode(APIResponse{
			Code:    http.StatusConflict,
			Message: "Email address exists.",
		})
		return
	}
	apiResp := APIResponse{
		Code:  http.StatusBadRequest,
		Error: err,
	}
	json.NewEncoder(w).Encode(apiResp)
}
