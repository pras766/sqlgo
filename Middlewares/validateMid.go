package validate

import "net/http"

func CheckUsername(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Adding middleWare \n"))
		next.ServeHTTP(w, r)
	})
}
