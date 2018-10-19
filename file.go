package file

import (
	"os"

	proto "github.com/laoqiu/go-file/proto"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"golang.org/x/net/context"
)

// Client is the client interface to access files
type Client interface {
	Open(ctx context.Context, filename string, new bool) (int64, error)
	Stat(ctx context.Context, filename string) (*proto.StatResponse, error)
	GetBlock(ctx context.Context, sessionId, blockId int64) ([]byte, error)
	Read(ctx context.Context, sessionId int64, buf []byte) (int, error)
	ReadAt(ctx context.Context, sessionId, offset, size int64) ([]byte, error)
	Close(ctx context.Context, sessionId int64) error
	Download(filename, saveFile string) error
	DownloadAt(ctx context.Context, filename, saveFile string, blockId int) error
	Upload(localfile, filename string) error
	UploadAt(ctx context.Context, localfile, filename string, blockId int) error
	UploadStream(ctx context.Context, b []byte, filename string, blockId int) error
	Remove(ctx context.Context, filename string) error
}

// NewClient returns a new Client which uses a micro Client
func NewClient(service string, c client.Client) Client {
	return &fc{proto.NewFileService(service, c)}
}

// NewHandler is a handler that can be registered with a micro Server
func NewHandler(readDir string) proto.FileHandler {
	return &handler{
		readDir: readDir,
		session: &session{
			files: make(map[int64]*os.File),
		},
	}
}

// RegisterHandler is a convenience method for registering a handler
func RegisterHandler(s server.Server, readDir string) {
	proto.RegisterFileHandler(s, NewHandler(readDir))
}
