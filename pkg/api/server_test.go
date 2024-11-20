package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/pazams/go-create-api/pkg/api/models"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/ping", nil)
	testExpectedStatusAndBody(t, req, []int{200}, &models.PingResponse{Message: "pong"})
}

// helpers
func testExpectedStatus(t *testing.T, req *http.Request, expectedStatuses []int) {

	s, err := InitializeServer()
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	s.srv.Handler.ServeHTTP(w, req)
	res := w.Result()
	assert.Contains(t, expectedStatuses, res.StatusCode)
}

func testExpectedStatusAndBody(t *testing.T, req *http.Request, expectedStatuses []int, expectedBody any) {

	s, err := InitializeServer()
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	s.srv.Handler.ServeHTTP(w, req)
	res := w.Result()
	assert.Contains(t, expectedStatuses, res.StatusCode)

	defer res.Body.Close()
	m := reflect.New(reflect.TypeOf(expectedBody).Elem()).Interface()
	err = json.NewDecoder(res.Body).Decode(m)
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, m)

}
