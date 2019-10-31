package protobuff

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

func GrpcClient() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := NewAddServiceClient(conn)

	g := gin.Default()
	g.GET("/add/:a/:b", func(context *gin.Context) {
		a, err := strconv.ParseUint(context.Param("a"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Parameter",
			})
		}
		b, err := strconv.ParseUint(context.Param("b"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Parameter",
			})
		}
		req := &Request{
			A: int64(a),
			B: int64(b),
		}
		if resp, err := client.Add(context, req); err == nil {
			context.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(resp.Result),
			})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "",
			})
		}
	})
	g.GET("/mult/:a/:b", func(context *gin.Context) {
		a, err := strconv.ParseUint(context.Param("a"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Parameter",
			})
		}
		b, err := strconv.ParseUint(context.Param("b"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Parameter",
			})
		}
		req := &Request{
			A: int64(a),
			B: int64(b),
		}
		if resp, err := client.Multiply(context, req); err == nil {
			context.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(resp.Result),
			})
		}
	})
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
