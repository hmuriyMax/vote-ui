package ui_server

import (
	"github.com/gin-gonic/gin"
	"strconv"
	auth "vote-ui/internal/services/auth-service"
)

func (s *Server) auth(ctx *gin.Context) {
	//_, err := ctx.Cookie(auth.CookieName)
	//if err != nil {
	//	ctx.Set("redirect", "login")
	//	ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	//}
	ctx.SetCookie(auth.CookieName, "test", 3600, "/", "", false, false)
	ctx.SetCookie(auth.CookieUserId, "1", 3600, "/", "", false, false)
}

func (s *Server) getAuthCookies(ctx *gin.Context) (int64, string, error) {
	token, _ := ctx.Cookie(auth.CookieName)
	userId, _ := ctx.Cookie(auth.CookieUserId)
	parseInt, _ := strconv.ParseInt(userId, 10, 64)
	//if err != nil {
	//	return 0, "", fmt.Errorf("error parsing user id cookie: %w", err)
	//}
	return parseInt, token, nil
}
