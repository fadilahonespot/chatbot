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

func SetLoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxLogger := logres.Context{
			ServiceName:    "chatbot-service",
			ServiceVersion: "1.0.0",
			ServicePort:    cast.ToInt(os.Getenv("APP_PORT")),
			ThreadID:       uuid.New().String(),
			ReqMethod:      r.Method,
			ReqURI:         r.URL.Path,
			Header:         r.Header,
		}

		reqBody, _ := io.ReadAll(&LoggingRequestReader{Request: r})
		lrw := &LoggingResponseWriter{ResponseWriter: w,}

		ctx := logres.SetCtxLogger(r.Context(), ctxLogger)
		ctx = request.SetRequestInContext(ctx, reqBody)
		r = r.WithContext(ctx)
		logger.Info(ctx, "Incoming Request")

		next.ServeHTTP(lrw, r)
		logger.TDR(r.Context(), reqBody, lrw.body)
	}
}

func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			err := errors.SetError(http.StatusUnauthorized, "missing Authorization header")
			response.ResponseError(w, err)
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.SetError(http.StatusUnauthorized, "signature not valid")
			}
			return []byte(constrans.JwtSecret), nil
		})

		if err != nil || !token.Valid {
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
