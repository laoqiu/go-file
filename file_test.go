package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/micro/go-micro"
	b "github.com/micro/go-micro/broker/mock"
	r "github.com/micro/go-micro/registry/mock"

	proto "github.com/laoqiu/go-file/proto"
	"golang.org/x/net/context"
)

func TestFileServer(t *testing.T) {
	// service cancellation context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// wait chan
	wait := make(chan bool)

	// make service
	s := micro.NewService(
		micro.Name("go.micro.srv.file"),
		micro.Registry(r.NewRegistry()),
		micro.Broker(b.NewBroker()),
		micro.Context(ctx),
		micro.AfterStart(func() error {
			close(wait)
			return nil
		}),
	)

	td := os.TempDir()
	f := filepath.Join(td, "/server_test.file")

	// write a file
	err := ioutil.WriteFile(f, []byte(`hello world`), 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f)

	// register file handler
	proto.RegisterFileHandler(s.Server(), NewHandler(td))

	// start service
	go s.Run()

	// wait for start
	<-wait

	// new file client
	cl := NewClient("go.micro.srv.file", s.Client())

	if err := cl.Download("server_test.file", "client_test.file"); err != nil {
		// no fatal as we need cleanup
		t.Error(err)
		return
	}
	defer os.Remove("client_test.file")

	// got file!
	b, err := ioutil.ReadFile("client_test.file")
	if err != nil {
		t.Error(err)
		return
	}

	if string(b) != "hello world" {
		t.Errorf("got %s, expected 'hello world'", string(b))
		return
	}

	if err := cl.Upload("client_test.file", "upload_test.file"); err != nil {
		t.Error(err)
		return
	}

	sessionId, err := cl.Open("upload_test.file", false)
	if err != nil {
		t.Error(err)
		return
	}

	buf := make([]byte, len(b))
	if _, err := cl.Read(sessionId, buf); err != nil {
		t.Error(err)
		return
	}

	if string(buf) != "hello world" {
		t.Errorf("got %s, expected 'hello world'", string(buf))
		return
	}

	if err := cl.Remove("upload_test.file"); err != nil {
		t.Error(err)
		return
	}

	if err := cl.UploadStream(buf, "upload_test.file", 0); err != nil {
		t.Error(err)
		return
	}

	if err := cl.Remove("upload_test.file"); err != nil {
		t.Error(err)
		return
	}
}
