package testApi

import (
	"github.com/go-playground/assert/v2"
	"go-mall/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {

	tests := []struct {
		name string
	}{
		{"success"},
		{"failed"},
	}

	r := routes.NewRouter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock  http
			request := httptest.NewRequest(
				"GET",
				"/ping",
				nil,
			)

			// mock 响应记录器
			w := httptest.NewRecorder()

			// 处理mock请求并响应
			r.ServeHTTP(w, request)

			// 检验状态码
			assert.Equal(t, http.StatusOK, w.Code)

			// 也可以检验响应内容
		})
	}
}
