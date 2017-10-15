package api

import (
	"net/http"
	"time"
	"riesling-cms-core/app/data"
	"github.com/gorilla/mux"
)

const (
	ACCESS_TOKEN  = "access_token"
	REFRESH_TOKEN = "refresh_token"
	HASH          = "hash"
)

func isAuthorized(r *http.Request) bool {
	accessToken := r.Header.Get(ACCESS_TOKEN)
	hash := r.Header.Get(HASH)
	if accessToken != "" && hash != "" {
		session := data.Session{}
		if session.Get(accessToken) {
			user := data.User{}
			if user.Get(hash) {
				if session.Hash == user.Hash {
					return session.ExpireTime.After(time.Now())
				}
			}
		}
	}
	return false
}

func isAuthorizedOnPathParam(r *http.Request) bool {
	isAuthorized := isAuthorized(r)
	pathParams := mux.Vars(r)
	if len(pathParams) > 0 {
		pathHash := pathParams[HASH]
		headerHash := r.Header.Get(HASH)
		return isAuthorized && pathHash == headerHash
	}
	return false
}

func isAuthorizedOnQueryParam(r *http.Request) bool {
	isAuthorized := isAuthorized(r)
	queryParams := r.URL.Query()
	if len(queryParams) > 0 {
		queryHash := queryParams.Get(HASH)
		headerHash := r.Header.Get(HASH)
		return isAuthorized && queryHash == headerHash
	}
	return false
}

func isAuthorizedAsAdmin(r *http.Request) bool {
	accessToken := r.Header.Get(ACCESS_TOKEN)
	hash := r.Header.Get(HASH)
	if accessToken != "" && hash != "" {
		session := data.Session{}
		if session.Get(accessToken) {
			user := data.User{}
			if user.Get(hash) {
				if session.Hash == user.Hash {
					return session.ExpireTime.After(time.Now()) && user.Level == data.USER_LEVEL_ADMIN
				}
			}
		}
	}
	return false
}

func SelfAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isAuthorized(r) {
			h.ServeHTTP(w, r)
			return
		}
		resp := APIResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized request",
		}
		ServeAsJSON(resp, w)
	}
}

func SelfAuthWithPathParam(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isAuthorizedOnPathParam(r) {
			h.ServeHTTP(w, r)
			return
		}
		resp := APIResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized request",
		}
		ServeAsJSON(resp, w)
	}
}

func SelfAuthWithQueryParam(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isAuthorizedOnQueryParam(r) {
			h.ServeHTTP(w, r)
			return
		}
		resp := APIResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized request",
		}
		ServeAsJSON(resp, w)
	}
}

func AdminAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isAuthorizedAsAdmin(r) {
			h.ServeHTTP(w, r)
			return
		}
		resp := APIResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized request",
		}
		ServeAsJSON(resp, w)
	}
}
