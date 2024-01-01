package middleware

import (
	"net/http"

	"github.com/GbSouza15/sistemaEstoque/pkg/token"
)

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {

			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error in server"))
			return
		}

		_, err = token.ValidatorToken(c.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		f(w, r)
	}
}
