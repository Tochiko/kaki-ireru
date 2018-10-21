package controllers_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"kaki-ireru/web/controllers"
	"kaki-ireru/web/models"
	"kaki-ireru/web/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnv_AllItems(t *testing.T) {
	router := gin.Default()
	db := &tests.TestDb{}
	env := controllers.Env{Db: db}

	router.GET("/items", env.AllItems)
	req, _ := http.NewRequest("GET", "/items", nil)
	httpRespTest(t, router, req, func (w *httptest.ResponseRecorder) bool {
		statusOk := w.Code == http.StatusOK
		bodyOk := false
		var items []*models.Item
		err := json.NewDecoder(w.Body).Decode(&items)
		if err == nil {
			bodyOk = true
		}
		return statusOk && bodyOk
	})

}

func httpRespTest(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}