package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang-course/task2/internal/gateway/usecase"
)

type HTTP struct {
	UC *usecase.GetRepository
}

type swaggerRepo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	CreatedAt   string `json:"created_at"`
}

func (h *HTTP) Register(r *gin.Engine) {
	r.GET("/api/v1/repos/:owner/:repo", h.getRepo)
}

// getRepo godoc
// @Summary      Repository metadata
// @Description  Returns name, description, stars, forks, and creation date (via Collector).
// @Tags         repos
// @Param        owner  path  string  true  "GitHub owner (user or org)"
// @Param        repo   path  string  true  "Repository name"
// @Success      200  {object}  swaggerRepo
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/v1/repos/{owner}/{repo} [get]
func (h *HTTP) getRepo(c *gin.Context) {
	owner := c.Param("owner")
	repo := c.Param("repo")
	info, err := h.UC.Run(c.Request.Context(), owner, repo)
	if err != nil {
		writeGRPCError(c, err)
		return
	}
	c.JSON(http.StatusOK, info)
}

func writeGRPCError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	switch st.Code() {
	case codes.NotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": st.Message()})
	case codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, gin.H{"error": st.Message()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
	}
}
