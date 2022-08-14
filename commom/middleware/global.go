package middleware

import (
	"fmt"
	"net/http"
)

// CommonJwtAuthMiddleware : with jwt on the verification, no jwt on the verification
type GlobalMiddleware struct {
}

func NewGlobalMiddleware() *GlobalMiddleware {
	return &GlobalMiddleware{}
}

func (m *GlobalMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("global before")
		next(w, r)
		fmt.Println("global end")
	}
}