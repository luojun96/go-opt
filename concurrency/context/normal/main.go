package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	helloWorldHandler := http.HandlerFunc(HelloWorld)
	http.Handle("/welcome", injectMsgID(helloWorldHandler))
	http.ListenAndServe(":8080", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msgID := ""
	if m := r.Context().Value("msgId"); m != nil {
		if value, ok := m.(string); ok {
			msgID = value
		}
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("msgId", msgID)
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello, world"})
}

func injectMsgID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msgID := uuid.New().String()
		ctx := context.WithValue(r.Context(), "msgId", msgID)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
