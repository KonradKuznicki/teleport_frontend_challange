package files_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"challenge/files"
	"github.com/stretchr/testify/assert"
)

func TestFilesStaticHandler(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest("GET", "/files", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(files.SataticsHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)

	assert.Contains(t, rr.Body.String(), "files manager")
}
