package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/fadilahonespot/chatbot/utils/constrans"
	"github.com/fadilahonespot/chatbot/utils/logger"
	"github.com/fadilahonespot/chatbot/utils/request"
	"github.com/fadilahonespot/chatbot/utils/response"
	"github.com/fadilahonespot/library/errors"
	"github.com/fadilahonespot/library/logres"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	body []byte
}

type LoggingRequestReader struct {
	*http.Request
	body []byte
}

func (lrw *LoggingResponseWriter) Write(b []byte) (int, error) {
	lrw.body = append(lrw.body, b...)
	return lrw.ResponseWriter.Write(b)
}

func (lrr *LoggingRequestReader) Read(b []byte) (int, error) {
	n, err := lrr.Body.Read(b)
	lrr.body = append(lrr.body, b[:n]...)
	return n, err
}

// SetLoggerMiddleware sets the logger middleware for the given http handler
func SetLoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctxLogger := logres.Context{
            // ServiceName is the name of the service
            ServiceName: "chatbot-service",
            // ServiceVersion is the version of the service
            ServiceVersion: "1.0.0",
            // ServicePort is the port on which the service is running
            ServicePort: cast.ToInt(os.Getenv("APP_PORT")),
            // ThreadID is a unique ID for the current request
            ThreadID: uuid.New().String(),
            // ReqMethod is the HTTP request method
            ReqMethod: r.Method,
            // ReqURI is the requested URI
            ReqURI: r.URL.Path,
            // Header is the request headers
            Header: r.Header,
        }

        // Read the request body
        reqBody, _ := io.ReadAll(&LoggingRequestReader{Request: r})
        // Create a new logging response writer
        lrw := &LoggingResponseWriter{ResponseWriter: w,}

        // Set the context logger
        ctx := logres.SetCtxLogger(r.Context(), ctxLogger)
        // Set the request body in the context
        ctx = request.SetRequestInContext(ctx, reqBody)
        // Set the updated request with the context
        r = r.WithContext(ctx)
        // Log the incoming request
        logger.Info(ctx, "Incoming Request")

        // Serve the next handler in the chain
        next.ServeHTTP(lrw, r)
        // Log the request and response bodies
        logger.TDR(r.Context(), reqBody, lrw.body)
    }
}

// JwtMiddleware is a middleware function that verifies the JWT token in the request's Authorization header.
// If the token is valid, it sets the userId in the context.
func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            // If the Authorization header is missing, return an unauthorized error.
            err := errors.SetError(http.StatusUnauthorized, "missing Authorization header")
            response.ResponseError(w, err)
            return
        }
        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                // If the signature is not valid, return an unauthorized error.
                return nil, errors.SetError(http.StatusUnauthorized, "signature not valid")
            }
            return []byte(constrans.JwtSecret), nil
        })

        if err != nil || !token.Valid {
            // If the token is not valid, return an unauthorized error.
            response.ResponseError(w, err)
            return
        }

        claims := token.Claims.(jwt.MapClaims)
        ctx := r.Context()

        // Set userId in context
        ctx = context.WithValue(ctx, "userId", fmt.Sprintf("%v", claims["userId"]))
        r = r.WithContext(ctx)

        next.ServeHTTP(w, r)
    }
}
