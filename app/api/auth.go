package api

import "net/http"

const (
	ACCESS_TOKEN  = "access_token"
	REFRESH_TOKEN = "refresh_token"
	HASH          = "hash"
)

func SelfAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get(ACCESS_TOKEN)
		hash := r.Header.Get(HASH)
		if hash != "" && accessToken != "" {
			h.ServeHTTP(w, r)
			return
		}
	}
}

func AdminAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get(ACCESS_TOKEN)
		hash := r.Header.Get(HASH)
		if hash != "" && accessToken != "" {
			h.ServeHTTP(w, r)
			return
		}
	}
}

func MemberAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get(ACCESS_TOKEN)
		hash := r.Header.Get(HASH)
		if hash != "" && accessToken != "" {
			h.ServeHTTP(w, r)
			return
		}
	}
}
