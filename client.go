package file

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	proto "github.com/laoqiu/go-file/proto"

	"golang.org/x/net/context"
)

const (
	blockSize = 512 * 1024
)

type fc struct {
	c proto.FileService
}

func (c *fc) Open(ctx context.Context, filename string, new bool) (int64, error) {
	rsp, err := c.c.Open(ctx, &proto.OpenRequest{Filename: filename, New: new})
	if err != nil {
		return 0, err
	}
	return rsp.Id, nil
}

func (c *fc) Stat(ctx context.Context, filename string) (*proto.StatResponse, error) {
	return c.c.Stat(ctx, &proto.StatRequest{Filename: filename})
}

func (c *fc) GetBlock(ctx context.Context, sessionId, blockId int64) ([]byte, error) {
	return c.ReadAt(ctx, sessionId, blockId*blockSize, blockSize)
}

func (c *fc) ReadAt(ctx context.Context, sessionId, offset, size int64) ([]byte, error) {
	rsp, err := c.c.Read(ctx, &proto.ReadRequest{Id: sessionId, Size: size, Offset: offset})
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

func (c *fc) Read(ctx context.Context, sessionId int64, buf []byte) (int, error) {
	b, err := c.ReadAt(ctx, sessionId, 0, int64(cap(buf)))
	if err != nil {
		return 0, err
	}
	copy(buf, b)
	return len(b), nil
}

func (c *fc) Close(ctx context.Context, sessionId int64) error {
	_, err := c.c.Close(ctx, &proto.CloseRequest{Id: sessionId})
	return err
}

func (c *fc) Write(ctx context.Context, sessionId, offset int64, buf []byte) error {
	_, err := c.c.Write(ctx, &proto.WriteRequest{Id: sessionId, Offset: offset, Data: buf})
	return err
}

func (c *fc) Download(filename, saveFile string) error {
	return c.DownloadAt(context.Background(), filename, saveFile, 0)
}

func (c *fc) DownloadAt(ctx context.Context, filename, saveFile string, blockId int) error {
	stat, err := c.Stat(ctx, filename)
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

	sessionId, err := c.Open(ctx, filename, false)
	if err != nil {
		return err
	}
	defer c.Close(ctx, sessionId)

	for i := blockId; i < blocks; i++ {
		buf, rerr := c.GetBlock(ctx, sessionId, int64(i))
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
	return c.UploadAt(context.Background(), localfile, filename, 0)
}

func (c *fc) UploadAt(ctx context.Context, localfile, filename string, blockId int) error {

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

	sessionId, err := c.Open(ctx, filename, true)
	if err != nil {
		return err
	}
	defer c.Close(ctx, sessionId)

	for i := blockId; i < blocks; i++ {
		buf := make([]byte, blockSize)
		n, err := file.ReadAt(buf, int64(i)*blockSize)
		if err != nil && err != io.EOF {
			return err
		}

		buf = buf[:n]

		if err := c.Write(ctx, sessionId, int64(i)*blockSize, buf); err != nil {
			return err
		}

		if err == io.EOF {
			break
		}
	}
	log.Printf("Upload %s completed", filename)

	return nil
}

func (c *fc) UploadStream(ctx context.Context, b []byte, filename string, blockId int) error {
	reader := bytes.NewReader(b)

	blocks := int(reader.Size() / blockSize)
	if reader.Size()%blockSize != 0 {
		blocks += 1
	}

	sessionId, err := c.Open(ctx, filename, true)
	if err != nil {
		return err
	}
	defer c.Close(ctx, sessionId)

	for i := blockId; i < blocks; i++ {
		buf := make([]byte, blockSize)
		n, err := reader.ReadAt(buf, int64(i)*blockSize)

		if err != nil && err != io.EOF {
			return err
		}

		buf = buf[:n]

		if err := c.Write(ctx, sessionId, int64(i)*blockSize, buf); err != nil {
			return err
		}

		if err == io.EOF {
			break
		}
	}

	log.Printf("Upload Stream %s completed", filename)

	return nil
}

func (c *fc) Remove(ctx context.Context, filename string) error {
	_, err := c.c.Remove(ctx, &proto.RemoveRequest{Filename: filename})
	return err
}
