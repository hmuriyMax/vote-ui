package ui_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) index(ctx *gin.Context) {
	uid, tok, err := s.getAuthCookies(ctx)
	if err != nil {
		s.internalError(ctx, err)
		return
	}

	votes, err := s.voteService.GetVotesByUserId(ctx, uid, tok)
	if err != nil {
		s.internalError(ctx, err)
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title":    "Главная",
		"username": "test",
		"votes":    votes,
	})
}

func (s *Server) voteForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "vote.html", gin.H{
		"title": "Vote",
	})
}

func (s *Server) internalError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, struct {
		Err string `json:"error"`
	}{err.Error()})
}
