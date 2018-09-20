package file

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	proto "github.com/laoqiu/go-file/proto"
	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"
)

type handler struct {
	readDir string
	session *session
}

func (h *handler) List(ctx context.Context, req *proto.ListRequest, rsp *proto.ListResponse) error {
	files := []*proto.FileObj{}

	if err := filepath.Walk(h.readDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {

			filename := info.Name()

			if len(req.Prefix) > 0 && !strings.HasPrefix(filename, req.Prefix) {
				return nil
			}

			if len(req.Suffix) > 0 {
				var supported bool
				for _, s := range req.Suffix {
					if strings.HasSuffix(filename, s) {
						supported = true
					}
				}
				if !supported {
					return nil
				}
			}

			files = append(files, &proto.FileObj{
				Filename:     info.Name(),
				Size:         info.Size(),
				LastModified: info.ModTime().Unix(),
			})
		}
		return nil
	}); err != nil {
		return err
	}

	rsp.Files = files

	return nil
}

func (h *handler) Open(ctx context.Context, req *proto.OpenRequest, rsp *proto.OpenResponse) error {
	var (
		file *os.File
		err  error
	)
	path := filepath.Join(h.readDir, req.Filename)

	if req.New {
		file, err = os.Create(path)
	} else {
		file, err = os.Open(path)
	}
	if err != nil {
		return errors.InternalServerError("go.micro.srv.file", err.Error())
	}

	rsp.Id = h.session.Add(file)
	rsp.Result = true

	log.Printf("Open %s, sessionId=%d", req.Filename, rsp.Id)

	return nil
}

func (h *handler) Close(ctx context.Context, req *proto.CloseRequest, rsp *proto.CloseResponse) error {
	h.session.Delete(req.Id)
	log.Printf("Close sessionId=%d", req.Id)
	return nil
}

func (h *handler) Stat(ctx context.Context, req *proto.StatRequest, rsp *proto.StatResponse) error {
	path := filepath.Join(h.readDir, req.Filename)
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.InternalServerError("go.micro.srv.file", err.Error())
	}

	if fi.IsDir() {
		rsp.Type = "Directory"
	} else {
		rsp.Type = "File"
		rsp.Size = fi.Size()
	}

	rsp.LastModified = fi.ModTime().Unix()
	log.Printf("Stat %s, %#v", req.Filename, rsp)

	return nil
}

func (h *handler) Read(ctx context.Context, req *proto.ReadRequest, rsp *proto.ReadResponse) error {
	file := h.session.Get(req.Id)
	if file == nil {
		return errors.InternalServerError("go.micro.srv.file", "You must call open first.")
	}

	rsp.Data = make([]byte, req.Size)
	n, err := file.ReadAt(rsp.Data, req.Offset)
	if err != nil && err != io.EOF {
		return errors.InternalServerError("go.micro.srv.file", err.Error())
	}

	if err == io.EOF {
		rsp.Eof = true
	}

	rsp.Size = int64(n)
	rsp.Data = rsp.Data[:n]

	log.Printf("Read sessionId=%d, Offset=%d, n=%d", req.Id, req.Offset, rsp.Size)

	return nil
}

func (h *handler) Write(ctx context.Context, req *proto.WriteRequest, rsp *proto.WriteResponse) error {
	file := h.session.Get(req.Id)
	if file == nil {
		return errors.InternalServerError("go.micro.srv.file", "You must call open first.")
	}

	n, err := file.WriteAt(req.Data, req.Offset)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.file", err.Error())
	}

	rsp.Size = int64(n)

	return nil
}

func (h *handler) Remove(ctx context.Context, req *proto.RemoveRequest, rsp *proto.RemoveResponse) error {
	path := filepath.Join(h.readDir, req.Filename)
	if err := os.Remove(path); err != nil {
		return errors.InternalServerError("go.micro.srv.file", err.Error())
	}
	return nil
}
