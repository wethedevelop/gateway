package account

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/wethedevelop/gateway/serializer"
	pb "github.com/wethedevelop/proto/auth"
)

const (
	address = "localhost:50051"
)

// SignupForm 注册表单
type SignForm struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// 用户注册
func Signup(c *gin.Context) {
	var form SignForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAccountAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Signup(ctx, &pb.SignupRequest{Account: form.Account, Password: form.Password})
	if err != nil {
		c.JSON(200, serializer.Err(serializer.CodeSignupErr, "注册失败", err))
	} else {
		c.JSON(200, serializer.Response{
			Data: gin.H{
				"nickName": r.GetNickname(),
			},
		})
	}
}
