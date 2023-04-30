package ui_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Index",
	})
}

func (s *Server) voteForm(ctx *gin.Context) {}

func (s *Server) auth(ctx *gin.Context) {}
