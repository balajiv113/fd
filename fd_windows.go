package fd

import (
	"errors"
	"net"
	"os"
)

func Get(via *net.UnixConn, num int, filenames []string) ([]*os.File, error) {
	return nil, errors.New("not supported")
}

func Put(via *net.UnixConn, files ...*os.File) error {
	return errors.New("not supported")
}
