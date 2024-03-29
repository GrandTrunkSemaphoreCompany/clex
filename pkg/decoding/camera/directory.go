package camera

import (
	"fmt"
	"github.com/GrandTrunkSemaphoreCompany/clex/pkg/encoding"
	"github.com/GrandTrunkSemaphoreCompany/clex/pkg/message"
)

// A directory is a kind of sink that creates a series of images in a folder
type Directory struct {
	BasePath string
	Id       int
}

func NewDirectory(basePath string, id int) *Directory {
	if id < 0 {
		return nil
	}

	return &Directory{basePath, id}
}

type sink interface {
	Write(m message.Message) (err error)
}

// Write takes a message and passes it through the Image writer
func (d *Directory) Write(m message.Message) (err error) {
	path := fmt.Sprintf("%s/%d/%d", d.BasePath, d.Id, m.Created.UnixNano())
	im := new(encoding.Image)
	im.BasePath = path

	_, err = im.Write([]byte(m.String()))

	return err
}
