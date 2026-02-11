package middleware

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetSessionMiddleware_Success(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))

	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(GetSessionMiddleware(log))

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	serv := httptest.NewServer(router)
	defer serv.Close()

	req, err := http.NewRequest(http.MethodGet, serv.URL+"/test", nil)
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{
		Name:  "sessionId",
		Value: "123",
	})

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	var str string

	require.NoError(t, json.NewDecoder(w.Body).Decode(&str))
	require.Equal(t, http.StatusOK, w.Result().StatusCode)
	require.Equal(t, "ok", str)
}

func TestGetSessionMiddleware_WithoutCookie(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))

	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(GetSessionMiddleware(log))

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	serv := httptest.NewServer(router)
	defer serv.Close()

	resp, err := http.Get(serv.URL + "/test")
	require.NoError(t, err)
	defer resp.Body.Close()

	var respBody struct {
		Err string `json:"error"`
	}

	expBody := "needed cookie with sessionId"

	require.NoError(t, json.NewDecoder(resp.Body).Decode(&respBody))
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	require.Equal(t, expBody, respBody.Err)
}
