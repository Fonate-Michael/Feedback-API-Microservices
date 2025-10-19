package middleware

import (
	"log"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(target string, prefix string) gin.HandlerFunc {

	url, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)

	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	return func(ctx *gin.Context) {
		ctx.Request.URL.Path = strings.TrimPrefix(ctx.Request.URL.Path, prefix)
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}

}
