package file

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	proto "github.com/laoqiu/go-file/proto"

	"golang.org/x/net/context"
)

const (
	blockSize   = 512 * 1024
	maxReadSize = 1 * 1024 * 1024
)

type fc struct {
	c proto.FileService
}

func (c *fc) Open(filename string, new bool) (int64, error) {
	rsp, err := c.c.Open(context.TODO(), &proto.OpenRequest{Filename: filename, New: new})
	if err != nil {
		return 0, err
	}
	return rsp.Id, nil
}

func (c *fc) Stat(filename string) (*proto.StatResponse, error) {
	return c.c.Stat(context.TODO(), &proto.StatRequest{Filename: filename})
}

func (c *fc) GetBlock(sessionId, blockId int64) ([]byte, error) {
	return c.ReadAt(sessionId, blockId*blockSize, blockSize)
}

func (c *fc) ReadAt(sessionId, offset, size int64) ([]byte, error) {
	rsp, err := c.c.Read(context.TODO(), &proto.ReadRequest{Id: sessionId, Size: size, Offset: offset})
	if err != nil {
		return nil, err
	}

	if rsp.Eof {
		err = io.EOF
	}

	if rsp.Data == nil {
		rsp.Data = make([]byte, size)
	}

	if size != rsp.Size {
		return rsp.Data[:rsp.Size], err
	}

	return rsp.Data, nil
}

func (c *fc) Read(sessionId int64, buf []byte) (int, error) {
	b, err := c.ReadAt(sessionId, 0, int64(cap(buf)))
	if err != nil {
		return 0, err
	}
	copy(buf, b)
	return len(b), nil
}

func (c *fc) Close(sessionId int64) error {
	_, err := c.c.Close(context.TODO(), &proto.CloseRequest{Id: sessionId})
	return err
}

func (c *fc) Write(sessionId, offset int64, buf []byte) error {
	_, err := c.c.Write(context.TODO(), &proto.WriteRequest{Id: sessionId, Offset: offset, Data: buf})
	return err
}

func (c *fc) Download(filename, saveFile string) error {
	return c.DownloadAt(filename, saveFile, 0)
}

func (c *fc) DownloadAt(filename, saveFile string, blockId int) error {
	stat, err := c.Stat(filename)
	if err != nil {
		return err
	}
	if stat.Type == "Directory" {
		return errors.New(fmt.Sprintf("%s is directory.", filename))
	}

	blocks := int(stat.Size / blockSize)
	if stat.Size%blockSize != 0 {
		blocks += 1
	}

	log.Printf("Download %s in %d blocks\n", filename, blocks-blockId)

	file, err := os.OpenFile(saveFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	sessionId, err := c.Open(filename, false)
	if err != nil {
		return err
	}
	defer c.Close(sessionId)

	for i := blockId; i < blocks; i++ {
		buf, rerr := c.GetBlock(sessionId, int64(i))
		if rerr != nil && rerr != io.EOF {
			return rerr
		}
		if _, werr := file.WriteAt(buf, int64(i)*blockSize); werr != nil {
			return werr
		}

		if i%((blocks-blockId)/100+1) == 0 {
			log.Printf("Downloading %s [%d/%d] blocks", filename, i-blockId+1, blocks-blockId)
		}

		if rerr == io.EOF {
			break
		}
	}
	log.Printf("Download %s completed", filename)

	return nil
}

func (c *fc) Upload(localfile, filename string) error {
	return c.UploadAt(localfile, filename, 0)
}

func (c *fc) UploadAt(localfile, filename string, blockId int) error {

	file, err := os.OpenFile(localfile, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := os.Stat(localfile)
	if err != nil {
		return err
	}
	blocks := int(stat.Size() / blockSize)
	if stat.Size()%blockSize != 0 {
		blocks += 1
	}

	sessionId, err := c.Open(filename, true)
	if err != nil {
		return err
	}
	defer c.Close(sessionId)

	for i := blockId; i < blocks; i++ {
		buf := make([]byte, blockSize)
		n, err := file.ReadAt(buf, int64(i)*blockSize)
		if err != nil && err != io.EOF {
			return err
		}

		buf = buf[:n]

		if err := c.Write(sessionId, int64(i)*blockSize, buf); err != nil {
			return err
		}

		if err == io.EOF {
			break
		}
	}
	log.Printf("Upload %s completed", filename)

	return nil
}

func (c *fc) Remove(filename string) error {
	_, err := c.c.Remove(context.TODO(), &proto.RemoveRequest{Filename: filename})
	return err
}
