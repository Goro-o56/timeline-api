package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"timeline-api/app/domain/entity"
)

func TestGet(t *testing.T) {
	h := &personaHandler{}
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.GET("/", h.Find)
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestGetByQuery(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.GET("/persona", func(c *gin.Context) {
		element := c.DefaultQuery("element", "Jobs")

		c.JSON(http.StatusOK, entity.Persona{
			ID:        0,
			Content:   element,
			CreatedAt: time.Time{},
		})
	})

	req, _ := http.NewRequest("GET", "/persona?element=Jobs", nil)
	r.ServeHTTP(w, req)

	e := entity.Persona{}
	if err := json.Unmarshal(w.Body.Bytes(), &e); err != nil {
		t.Error(err)
	}
	assert.Equal(t, "Jobs", e.Content)

}
