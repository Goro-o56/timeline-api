package server

import (
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	h := &timelineHandler{}
	w := httptest.NewRecorder()

}
