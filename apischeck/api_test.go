package apischeck

import (
	"net/http"
	"net/http/httptest"
	"portfolio-user-service/controller"
	"portfolio-user-service/repository/auth"
	authService "portfolio-user-service/services/auth"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// setupTestRouter creates router + db + repo + service
func setupTestRouter() *gin.Engine {
	logger, _ := zap.NewDevelopment()
	repo := &auth.FakeAuthRepo{}                         // 👈 use fake repo, not real sqlite
	svc := authService.NewAuthService(repo, nil, logger) // pass nil for db if unused
	authController := controller.NewAuthController(svc, logger)

	r := gin.Default()
	r.POST("/auth/register", authController.RegisterUser)
	return r
}

func TestRegisterAPI(t *testing.T) {
	router := setupTestRouter()

	body := `{"firstName":"Amit","lastName":"Kumar","email":"amit@example.com","password":"123456"}`
	req := httptest.NewRequest("POST", "/auth/register", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "success")
}
