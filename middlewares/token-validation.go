package middlewares

import (
	"net/http"

	"github.com/shadowshot-x/micro-product-go/authservice/jwt"
)

// We want all our routes for REST to be authenticated. So, we validate the token
func tokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// check if token is present
		if _, ok := r.Header["Token"]; !ok {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Token Missing"))
			return
		}
		token := r.Header["Token"][0]
		err := jwt.ValidateToken(token, "Secure_Random_String")

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Token Validation Failed"))
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Authorized Token"))

	})
}
