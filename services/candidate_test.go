package services

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
)

func TestSum(t *testing.T) {
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	CandidateList(context)
}
