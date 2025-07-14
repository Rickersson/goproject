package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(
			200,

			gin.H{

				"mensagem": "pong",
			},
		)

	},
	)

	r.Run(":3000")
}
