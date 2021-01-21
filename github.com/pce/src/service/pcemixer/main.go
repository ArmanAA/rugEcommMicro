package main

import (
	"context"
	"fmt"
	"net/http"

	pceadderpb "github.com/ArmanAA/pce/src/proto/pceadder"

	pcemixerpb "github.com/ArmanAA/pce/src/proto/pcemixer"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type server struct{}

// AddReq ...
type AddReq struct {
	A int64 `json:"firstNum"`
	B int64 `json:"secondNum"`
}

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:4041", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pceadderpb.NewPceAdderClient(conn)

	router := gin.Default()
	router.POST("/add", func(c *gin.Context) {
		var data AddReq
		if err := c.ShouldBindJSON(&data); err != nil {
			fmt.Println("ERROR: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		req := &pceadderpb.Request{
			A: data.A,
			B: data.B,
		}

		r, err := client.Add(ctx, req)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": r.GetResult(),
		})
	})
	router.Run(":8080")
}

func (s *server) Add(ctx context.Context, req *pcemixerpb.AddRequest) (*pcemixerpb.AddResponse, error) {
	a := req.GetFirstNum()
	b := req.GetSecondNum()
	fmt.Println("Numbers: ", a, b)
	return nil, nil
}
