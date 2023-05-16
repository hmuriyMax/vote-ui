package ui_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) postVoteForm(ctx *gin.Context) {
	uid, tok, err := s.getAuthCookies(ctx)
	if err != nil {
		s.internalError(ctx, err)
		return
	}

	voteID, err := strconv.ParseInt(ctx.PostForm("voteID"), 10, 64)
	if err != nil {
		s.internalError(ctx, err)
		return
	}

	val, err := strconv.ParseInt(ctx.PostForm("choice"), 10, 64)
	if err != nil {
		s.internalError(ctx, err)
		return
	}

	err = s.voteService.Vote(ctx, val, voteID, uid, tok)
	if err != nil {
		s.internalError(ctx, err)
		return
	}
	ctx.Redirect(http.StatusFound, "/")
}
