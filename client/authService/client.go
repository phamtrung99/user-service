package authservice

import (
	"fmt"
	"log"
	"os"
	"time"

	authServiceProto "github.com/phamtrung99/auth-service/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var (
	conn   *grpc.ClientConn
	client authServiceProto.AuthServiceClient
)

func init() {

	conn, err := grpc.Dial(
		os.Getenv("ENDPOINT"),
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                5 * time.Minute,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "fail to connect to grpc"))
	}

	client = authServiceProto.NewAuthServiceClient(conn)

	fmt.Println("Connected to auth-service")
}

func Close() {
	conn.Close()
}

func GetClient() authServiceProto.AuthServiceClient {
	return client
}
