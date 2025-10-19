package middlewares

import "github.com/gin-gonic/gin"

func AdminMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		role, exist := context.Get("role")

		if !exist {
			context.IndentedJSON(401, gin.H{"error": "Unauthorized Role Not Found!"})
			context.Abort()
			return
		}

		if role != "admin" {
			context.IndentedJSON(401, gin.H{"error": "Unauthorized you must be admin to acces this page hehehe"})
			context.Abort()
			return
		}

		context.Next()
	}
}
