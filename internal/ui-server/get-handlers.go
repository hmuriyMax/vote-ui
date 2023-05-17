package ui_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (s *Server) vote(ctx *gin.Context) {
	uid, tok, err := s.getAuthCookies(ctx)
	if err != nil {
		s.internalError(ctx, err)
		return
	}

	voteID, err := strconv.ParseInt(ctx.Query("vote_id"), 10, 64)
	if err != nil {
		s.internalError(ctx, err)
		return
	}

	vote, err := s.voteService.GetVoteById(ctx, voteID, uid, tok)
	if err != nil {
		s.internalError(ctx, err)
		return
	}

	ctx.HTML(http.StatusOK, "vote.html", gin.H{
		"title":    vote.Name,
		"voteID":   vote.ID,
		"finishes": vote.FinishesAt,
		"variants": vote.Variants,
	})
}

func (s *Server) internalError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, struct {
		Err string `json:"error"`
	}{err.Error()})
}
