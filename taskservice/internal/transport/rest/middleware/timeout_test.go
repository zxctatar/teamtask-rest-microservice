package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestTimeout(t *testing.T) {
	tests := []struct {
		name    string
		timeMs  time.Duration
		status  int
		expBody string
	}{
		{
			name:    "good test",
			timeMs:  1 * time.Millisecond,
			status:  http.StatusOK,
			expBody: "ok",
		},
		{
			name:    "timeout test",
			timeMs:  3 * time.Millisecond,
			status:  http.StatusRequestTimeout,
			expBody: "timeout",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.Use(TimeoutMiddleware(2 * time.Millisecond))

			r.GET("/test", func(ctx *gin.Context) {
				time.Sleep(tt.timeMs)
				ctx.JSON(http.StatusOK, "ok")
			})

			serv := httptest.NewServer(r)
			defer serv.Close()

			resp, err := http.Get(serv.URL + "/test")
			require.NoError(t, err)
			defer resp.Body.Close()

			var respBody string

			require.NoError(t, json.NewDecoder(resp.Body).Decode(&respBody))
			require.Equal(t, tt.expBody, respBody)
			require.Equal(t, tt.status, resp.StatusCode)
		})
	}
}
