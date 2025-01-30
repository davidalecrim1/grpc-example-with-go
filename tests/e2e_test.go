package e2e_test

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"grpc-example-with-go/internal/app"
	handler "grpc-example-with-go/internal/handler/grpc"
	gen "grpc-example-with-go/internal/handler/grpc/generated"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var conn *grpc.ClientConn

func TestMain(m *testing.M) {
	c, cancel, err := createTestServer()
	if err != nil {
		log.Fatalf("Failed to create test server: %v", err)
	}
	conn = c
	defer cancel()

	code := m.Run()
	os.Exit(code)
}

func createTestServer() (conn *grpc.ClientConn, cancel func(), err error) {
	// Initialize an In-Memory gRPC Server
	// Using an in-memory connection helps avoid network overhead
	// and makes tests faster and more reliable.
	lis := bufconn.Listen(bufSize)

	svc := app.NewProductService()
	handler := handler.NewProductGrpcHandler(svc)

	s := grpc.NewServer()
	gen.RegisterProductHandlerServer(s, handler)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	conn, err = grpc.NewClient(
		"passthrough://bufnet",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}

	cancel = func() {
		s.Stop()
		lis.Close()
	}

	return conn, cancel, nil
}

func TestE2E(t *testing.T) {
	t.Run("valid product creation", func(t *testing.T) {
		client := gen.NewProductHandlerClient(conn)

		req := &gen.CreateProductRequest{
			Name: "Test-Product-001",
		}

		res, err := client.Create(context.Background(), req)
		require.NoError(t, err)

		assert.Equal(t, res.Name, req.Name)
		assert.NotEmpty(t, res.Id)
	})
}
