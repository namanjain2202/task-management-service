package middleware

import (
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
)

// AuthMiddleware is a middleware function that checks for a valid JWT token in the Authorization header.
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        if tokenString == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        claims := &jwt.StandardClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            // Validate the token signing method here
            return []byte("your-secret-key"), nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Set the claims in the context for use in the next handler
        ctx := r.Context()
        ctx = context.WithValue(ctx, "userID", claims.Subject)
        r = r.WithContext(ctx)

        next.ServeHTTP(w, r)
    })
}