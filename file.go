package file

import (
	"os"

	proto "github.com/laoqiu/go-file/proto"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
)

// Client is the client interface to access files
type Client interface {
	Open(filename string, new bool) (int64, error)
	Stat(filename string) (*proto.StatResponse, error)
	GetBlock(sessionId, blockId int64) ([]byte, error)
	Read(sessionId int64, buf []byte) (int, error)
	ReadAt(sessionId, offset, size int64) ([]byte, error)
	Close(sessionId int64) error
	Download(filename, saveFile string) error
	DownloadAt(filename, saveFile string, blockId int) error
	Upload(localfile, filename string) error
	UploadAt(localfile, filename string, blockId int) error
	UploadStream(b []byte, filename string, blockId int) error
	Remove(filename string) error
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
