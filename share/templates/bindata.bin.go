// Code generated by go-bindata.
// sources:
// index.html
// DO NOT EDIT!

// +build bin

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5c\xeb\x73\xdb\x36\x12\xff\xee\xbf\x02\xc7\x6b\x6f\xda\x4e\x29\xd5\x79\xdc\xf5\x52\xcb\x37\x8e\xe5\x24\x9a\xc6\x8e\xc6\x8f\xcc\xf5\xbe\x74\x20\x12\x94\x10\x53\x24\x0b\x40\x72\x7c\x1a\xfd\xef\xb7\x78\xf1\xa9\x17\x45\xda\xf1\x4d\xfa\x21\x31\x89\xc7\xee\x62\xb1\xf8\xed\x02\x58\xf1\xe8\x2f\xfd\x0f\xa7\xd7\xbf\x0d\xcf\xd0\x44\x4c\xc3\x83\xe3\x23\xf5\x07\x21\x78\x20\xd8\x87\x07\xf9\x38\x25\x02\x23\x6f\x82\x19\x27\xa2\xe7\xcc\x44\xe0\xfe\xec\xe4\xab\x26\x42\x24\x2e\xf9\x63\x46\xe7\x3d\xe7\xdf\xee\xcd\x89\x7b\x1a\x4f\x13\x2c\xe8\x28\x24\x0e\xf2\xe2\x48\x90\x08\xfa\x0d\xce\x7a\xc4\x1f\x93\x42\xcf\x08\x4f\x49\xcf\x99\x53\x72\x97\xc4\x4c\xe4\x1a\xdf\x51\x5f\x4c\x7a\x3e\x99\x53\x8f\xb8\xea\xe5\x47\x44\x23\x2a\x28\x0e\x5d\xee\xe1\x90\xf4\x0e\x2d\x21\x41\x45\x48\xd4\x33\xbc\x2d\x16\x9d\x3e\x16\xb8\xf3\x2e\xe6\x42\x12\x5f\x2e\x11\x3c\x01\xc9\xa3\x6e\xd6\xee\xf8\x28\xa4\xd1\x2d\x62\x24\xec\x39\x14\x58\x3a\x48\xdc\x27\x20\x07\x9d\xe2\x31\xe9\x26\xd1\xd8\x41\x13\x46\x82\x9e\xd3\x0d\xf0\x5c\x36\xe8\xc8\xb2\x4a\x57\x2e\xee\x43\xc2\x27\x84\x08\x4b\x40\x90\xcf\xa2\xeb\x71\x9e\xf6\x87\xe7\x2e\x8d\x7c\xf2\xb9\x23\x4b\x0d\x05\xee\x31\x9a\x88\x7c\x97\x4f\x78\x8e\x75\xa9\x53\xd6\x33\xe2\xcc\x03\x42\x9f\x78\x97\x49\x0d\x33\x02\x4f\xcf\x3a\x87\x9d\xc3\x9f\x6d\x41\x67\x4a\xa3\xce\x27\xe0\xe9\xc3\xc0\xdd\x29\xa6\x91\x6e\xbf\x58\xd0\x00\x75\xae\x4f\xde\xbe\x3d\xeb\x8f\x68\xb4\x5c\x42\x3b\x23\x8c\xee\xb1\x58\x90\x90\x83\x86\x80\x43\x77\x4a\xc3\x5b\x53\xa9\x2a\x22\x7f\xb9\x74\xac\x52\x8f\xba\x5a\x38\x23\x7f\xd7\x98\xc6\xf1\xd1\x28\xf6\xef\x4d\x61\x84\xe7\xc8\x0b\x31\xe7\x3d\x07\x1e\x47\x98\x21\xfd\xc7\xf5\x49\x80\x67\xa1\xb0\xaf\x5c\x80\x61\x78\xae\x88\x13\x07\xb1\x18\xe6\x51\x36\xa7\x63\x28\x84\x79\x48\x19\xfa\x34\xa5\x26\x4d\x02\xc6\x44\x98\x1b\x84\x33\xea\xdb\x36\xa5\x56\x86\xba\x94\x8c\x30\x30\xa3\xd0\x9d\xfa\xee\x21\x4a\xb0\xef\xd3\x68\xec\x86\x24\x80\xc9\xb7\x53\x60\xfb\x8f\x66\x42\xc4\x51\x89\x84\x88\xc7\xe3\x90\x48\x12\x21\x4e\x38\xf1\xed\xdc\xea\xc6\x46\xcb\xba\x91\x14\x4e\xb7\xb2\xc5\x98\x8d\xe5\xd4\xfd\xd5\xd0\x4a\xab\x73\x6c\x95\x05\x24\x38\x65\xcb\x99\x1b\x47\xe1\x7d\xb1\x09\x34\xba\xd6\x72\x64\xda\x81\x59\x80\x6e\x1b\x28\x49\x4b\x75\x81\x6d\x85\xd4\x93\xe9\xd9\xd5\x4a\x2c\x4c\x02\x46\xd4\xef\x39\x13\xd0\x6c\x71\x1e\x46\x0c\x47\x3e\x94\x51\x30\x14\xb5\xd2\x7a\xce\x14\x7f\xd6\x58\xf0\x0a\x1d\x3e\xf3\x26\xe9\x2a\x83\x29\x92\x6b\x1b\xa8\x98\x55\x8f\xaa\x30\x50\x9a\x82\x6a\x83\xa3\x2e\x2e\xc8\xd5\x05\xeb\x2a\xd9\x9a\x14\xb4\x3c\xb3\x99\x99\xea\x02\xb4\x61\xea\x8f\x8f\x66\x61\x6e\x94\xb6\x29\xfc\x29\x1b\x48\x48\x6d\x3b\xec\x09\x3a\x27\x55\xf5\x62\x33\x78\x89\xbd\xfc\x55\xb7\x7b\x77\x77\xd7\x81\xb1\x30\xf8\xd7\xf1\xe2\x69\x57\xa3\x1e\x60\x44\x48\x30\x27\xbc\x1b\x62\x41\xb8\xf8\x97\x37\x4d\x7a\x76\xec\x1f\xcf\x2e\xaf\x06\x1f\x2e\xca\xba\x51\xf4\x2d\x68\xe2\xea\xbc\x86\x74\xad\xb0\x76\x09\x01\xa6\xa1\xdc\x33\x80\x12\x1b\xd3\x48\x2e\x40\x14\x50\xc6\x85\x2a\xad\x8e\xc9\x8b\x7d\x52\x22\x25\x8b\x00\x71\x2b\x0a\x50\x22\xa6\xd3\xdc\x1f\x5c\x5d\x5f\x0e\x5e\xcb\x49\x94\x1d\x5a\x96\xb9\xa9\xa0\x34\xf9\x5b\x34\xe2\xc9\x2f\x7a\xe1\x48\x23\x02\x9b\xae\x34\xcb\x19\xe5\x60\x28\x47\x52\x59\x3f\x66\x2c\x4f\x70\x84\xb3\xca\x08\x67\x9b\x47\x78\x93\x08\xaa\x17\xdd\xff\xd1\x28\x43\x5c\x1e\x65\x88\x37\x8e\xf2\xfd\x49\x0b\x23\xec\xce\xc2\x4d\xb8\x94\x7b\x85\x17\x18\x84\xf1\xc7\xdb\x3d\x68\xa1\x0d\x8b\xef\xca\x9e\x55\x0e\x70\x4a\xa6\xe0\xd4\x74\x1c\x97\xe2\x9c\x74\xad\x2f\xac\x8f\x8d\x83\x00\xa2\x15\xf7\xb0\x80\x73\x8b\x85\x20\xd3\x44\x42\x0e\x72\xc0\xff\x03\x6f\xfe\xea\x95\x7e\xf8\x1d\xb4\x41\x42\xa0\xeb\xa0\x8e\xd4\xce\x6a\x94\xf5\x83\xd5\x6c\x5f\x36\x63\xeb\x07\x2b\xb9\x16\x74\xb8\x5d\x29\x5e\x32\x7b\x10\xa5\x00\xdd\x8d\x4a\xa1\x0f\xa3\x14\xda\x8a\x52\x12\xbe\x5a\xb8\xc3\x9f\x9a\x49\x97\xf0\x36\xa4\x9b\x8f\x1f\x46\xba\xf9\x78\xbb\x74\xe9\xcb\x96\xc8\xdf\x76\x98\x43\xec\x2c\xe1\x03\xf5\xac\x7f\x5b\x2e\x7f\x29\x87\xe0\x26\xf2\x96\xc1\xb8\xda\xb7\x2d\x16\xdd\x1f\x0e\x7e\xe8\x2e\x97\x8b\x85\x96\xaf\x2a\x30\x98\x17\x68\x87\x3b\xb2\xcd\x1d\x15\x13\x24\xdf\x60\xcc\x72\x00\xc0\x48\xfe\xa7\xf7\x03\x8b\x05\x84\x5e\x63\x82\xbe\xa1\x3f\xa2\x6f\xbc\x98\x11\xf4\xaa\x87\x34\xa2\x9d\x0e\x6f\x3a\xef\x29\x17\x30\x64\xc1\x40\x3c\x55\xdd\xb9\xe8\x9c\x08\xc1\x7e\x25\xf7\x48\xad\x0d\xa0\x3b\xba\x77\x2f\x80\x91\x92\x50\xf8\x56\xe9\x0a\x8f\x19\x1d\x4f\x00\xa1\xe3\x3b\x86\x8d\x8f\x38\x4e\xe9\x48\x4d\x0a\x7f\x6d\xaf\x74\x03\x25\xf1\x77\xb1\x50\xf5\xcb\x25\x38\x1b\x4e\x12\xc2\x3c\x08\x58\x74\x60\xa1\x83\xf1\xac\xb4\x67\x19\xdc\x70\xc2\x94\x54\x05\xb6\xba\x34\x07\xd4\xc7\x0f\x2b\xc5\xd5\x3d\xaf\x0a\xa1\x0a\xdb\x96\xc1\xa5\xd1\x9c\xb0\x74\x7b\xb2\x42\x96\x81\x1f\x92\xaa\x30\xba\x74\xa5\x34\x5d\xc1\xa4\xb5\xa9\x0d\x62\xf9\xcf\x2e\x26\xa8\x97\xa1\xb4\x1f\xf9\x90\x8e\x4c\xbd\xa8\xff\xdd\x49\x0c\x32\x3b\x7a\xe8\xb9\x43\x08\xe0\x6b\x57\x94\x98\xe4\xd7\x5a\xfa\x26\x2b\xb6\x98\x9a\x6a\x26\xe7\xbb\xb0\x89\x99\x45\xb4\x10\x8d\x1e\x7f\x5b\xf2\xda\xfb\x70\x81\x09\x7d\x78\x26\x72\xa2\x9a\x70\xe9\x1a\xad\xaa\x52\xbb\xab\x17\xd9\xb6\x7e\x13\x00\x5a\x3c\xd1\x00\x28\x32\x44\x52\xd3\x98\x33\x92\xad\x66\xe1\x07\xa3\x7b\xd8\x9d\xfc\x5e\x13\x9e\x7c\xca\x6f\x33\x68\xea\xbf\x51\x44\x8a\xf0\x24\x9b\x74\xfa\x94\x5d\xc0\x16\x2f\x03\x29\xc3\xcf\x00\x95\x4f\x99\xdc\x01\x66\x70\x95\x82\x92\xee\x4d\xe6\x17\x66\x83\x98\x2e\xc8\x52\x0b\x4d\x7f\x37\xf0\xb2\x9d\x4e\xe6\x98\x86\xbb\x76\x59\x3f\x09\xd9\x92\x76\xb4\x42\x24\x9a\x0d\x75\xc9\x72\xa9\xe3\x55\xcb\x12\x6a\xfc\x7a\x42\x5e\xc7\x02\xe7\x85\x6c\xb4\xf8\xed\x2c\x5b\x04\x00\x26\x49\x36\x7d\x43\xcc\xf0\x94\x37\xc2\x85\x95\x7a\xd2\x27\x42\x24\x9a\x4d\xe5\xc1\x1d\x28\x29\xe9\x7c\xc4\x8c\xca\x46\xf0\xd8\x0f\x6e\xd1\x21\x72\xfa\xea\x9c\xd1\x41\x8e\x12\x6b\x2f\x2a\xcf\x90\x73\x1e\xcf\x22\x21\x0f\x8c\x1a\x90\x79\x8e\x1c\x65\x19\x40\x24\x37\x2b\x7b\x93\x7b\x81\x9c\x1b\x75\x86\xd5\x0a\xb5\x97\xc8\x51\x16\x51\x26\xd7\x08\x46\x0a\x8b\xbf\x0d\x2c\xa1\x11\x6c\xac\x1a\x83\x89\xa6\xb2\x1b\x9a\xe8\xb6\x5f\x14\x4e\x06\x01\x23\x3b\x77\xa9\x03\x27\x83\xd9\x5a\x3c\x91\x55\x35\x01\x65\xa0\x34\xd5\x1e\xa2\x98\xb9\x6e\x27\xa8\xb0\xfd\xf4\x42\xc8\x79\x59\x0d\x0f\x6b\xdc\x73\xa5\xb5\x81\x81\x8d\xcd\x51\x59\x43\xaa\x99\x5a\xf9\x7b\xf4\x93\x4b\x7c\x8f\x6e\x6a\x2d\xb7\x16\x0e\x14\x17\x5e\xf3\x85\x4c\xf7\x0b\x0a\x68\x90\xad\xe2\xc1\xaa\x90\x80\x06\x9d\xe2\xfa\xa5\x85\x68\x60\xcd\xda\x35\xbd\x76\xb6\x76\x68\xdf\x27\xa1\xc0\x83\xa8\x76\x97\x0f\x33\x51\xa7\x4f\x3d\x0e\x45\xe2\x8d\x56\x1f\x2d\xf9\xf3\xb6\x22\xfa\x01\xac\x1e\x16\xe0\xb5\x0b\xae\x1a\x0f\xdb\xe3\xff\xd7\x83\xeb\x2b\x04\xf8\x85\x38\xf1\xe2\x28\x7f\x55\x34\x88\x36\x47\xc9\x47\xa5\x53\xb7\xd1\x51\x37\x5f\x72\x9c\xf0\xbd\x83\xf5\xed\xc2\xc1\x9c\x7c\x71\xe9\x84\xc4\x02\xf4\xfa\xb7\xeb\xb3\x2b\x34\x8d\xfd\x59\x18\xa3\x17\x6f\x1b\x28\xf0\x75\x49\xc4\x6f\x5f\xbc\x7d\x70\x19\x6b\xeb\xb1\xae\x90\x4d\xe0\x91\xb6\x1b\xe6\xd0\x80\x30\x16\xb3\x86\xf0\xa8\x69\x6c\xc5\x47\xdd\xec\x4f\x80\xac\x0d\x90\x66\x8e\x9e\x0e\x42\xae\x01\xc7\xfc\xbd\xc6\xba\xc5\xd3\x02\xc8\xac\x43\xbf\x47\x62\xbf\x1a\xd7\x76\x61\xde\x06\x7a\xad\x41\xac\x36\xd8\x37\xc3\xa5\x3c\x90\xb4\x01\x4c\x09\xf6\x6e\x89\x68\x88\x4c\x86\xc8\x56\x68\x32\xed\xfe\xc4\xa6\xda\xd8\x64\xa7\xe9\x4f\x70\x7a\x12\xe0\xa4\x83\x9b\x2f\x0b\x51\xeb\x65\x78\x12\x40\x55\x00\x96\xe6\x48\x35\x25\xd3\x3d\xee\xc3\xa0\x57\x86\x53\xe7\x67\xe7\x45\x84\x82\xda\xce\xaf\x34\xf2\x33\x88\x92\x37\xe8\x1a\x9c\x6e\xa1\x7c\x05\x38\xd9\x2e\x3b\x03\x81\xec\xf0\xa6\xfd\xd3\x1e\x49\x76\xe5\xd9\xb1\xa9\xa8\x27\x60\xab\x27\xc7\x40\xf0\xd1\xaf\x8d\x72\xc6\x2f\x95\x5d\xa7\xfd\x86\xe3\x98\x47\x38\x87\xb1\x66\xdd\x7c\x81\xa4\x49\x09\x6b\xcf\xea\xf5\xcd\xf6\xf1\x51\x76\xc3\x8d\xf5\x10\x64\x7b\xb5\x06\xde\x31\x12\x98\x4c\x46\x28\x39\x8d\xa3\x80\x8e\x81\xe4\x72\x99\xbf\xb5\x1c\x89\xc8\x1d\x85\xb1\x77\x9b\xde\x81\x23\x64\xaf\x36\x0d\x9d\x53\xd9\xf6\xa2\x40\x03\x39\x93\x17\x68\x34\x76\x69\x14\xc4\x8e\x7c\x71\xd2\x8b\x4c\xa0\x70\x3a\xbc\x29\x40\xd1\xb1\xcd\x65\x4b\x2f\xe4\xf3\x62\xcb\x3c\x81\xcd\xbc\x3c\xf5\x68\xb2\x89\x54\xf6\x6a\xa5\x28\xcd\x19\x75\x27\xd4\xf7\x49\x04\xf2\xd8\x7c\x11\xdd\x34\xcb\xbc\x09\x62\x36\xcd\x18\x9e\x78\x32\xbf\xf3\x0d\x94\x15\xf5\x22\x5b\xc1\xf0\x8a\xe9\x48\xb0\x8b\x8d\x92\x99\x48\x0f\xf8\x14\x2b\x97\xcf\x46\x53\x9a\xe6\x21\x9b\xb7\x5c\x9f\x5c\x8a\x84\xd4\xb6\x88\xe3\xb0\x94\xb5\x59\x69\x33\x66\xf1\x2c\x41\xe9\x93\xcb\xa7\x36\x63\x57\xbd\x97\x53\x16\xf1\x2a\x05\xbe\xa3\x3e\x51\xea\x03\x32\x8a\x94\xcd\x06\x36\x69\x8d\x95\x0a\x33\x87\x6b\xec\xc7\x90\xcb\xcd\xb3\xe1\xfe\x4e\x69\xa1\x9c\xae\xb8\x4a\xa8\x2c\x8f\xe1\xec\x33\x58\x87\x2f\x17\x05\xbc\x55\x05\xa9\xca\xec\x53\x2e\x5b\xfb\x9b\x65\xd4\x64\x57\x4a\x69\xf3\xc2\x32\xee\x12\x0f\xe0\xad\x9c\x7d\x9a\xb3\xd2\x96\xa6\x67\x13\x62\x80\xf8\x8c\xf0\x89\x83\xbe\x3b\x9d\x60\x1a\x69\x15\x21\xe7\x52\x17\x4b\x00\xf8\xbe\x9c\xe2\x52\x11\x10\x5e\xa5\xb5\x1e\x54\x2a\xeb\x2c\xb7\xd4\x5a\xca\x0b\x49\xdd\xa8\x1d\xec\x30\x94\x34\xc1\xa0\x98\x95\x93\x13\x42\x3d\xd4\x05\x41\x3f\xa8\x8d\x81\x1b\x11\xd0\x0f\xb6\x01\xe0\x16\xf8\xf3\x83\xcd\xe8\xd7\x97\xb7\x59\x33\x8e\xc7\xa4\x2d\x10\x94\x1c\x1b\x60\xa0\x1f\x7c\xcd\x10\x28\x95\xd7\x1e\x02\x4a\xeb\x69\x0a\x80\x19\xf8\xf5\xdf\xb4\x8d\x7d\xab\xe4\xb3\xd0\x97\xc1\x5e\xff\xcd\x53\x46\x3d\x58\xf1\xf5\x41\xaf\x92\xdd\x2f\x93\x0e\x39\xaa\x2c\x9b\x1c\x81\x90\x96\x26\xe8\x5a\x5f\xb6\x0b\x99\x96\x90\x9a\x88\x53\x50\xe8\xea\x61\xad\xbf\xb8\x57\xb4\xf4\xcd\xab\xa6\x55\xc8\x32\xde\x24\xc3\xb3\x16\x65\x00\x5a\xaf\xe5\x09\xf8\x4a\x11\xb2\x24\xe7\x06\xae\x43\x43\xd4\x06\xcf\xb1\xa2\x73\x5e\xdd\x4e\xb5\x7b\x1e\x53\x77\xba\xfc\x5c\xe5\x7d\xb6\x70\x7e\xd6\x90\x73\xfe\x16\xae\xca\xb8\x35\x1f\x48\xeb\xfb\xc0\x2d\x5e\x90\x6e\xf5\x82\x5b\xb7\x01\x74\x8b\x1f\x4c\x0f\xa3\x78\x5b\x7e\x90\x36\xf3\x83\xf4\xab\xf6\x83\xb4\x55\x3f\x48\x5b\xf5\x83\x83\xd6\xfd\xe0\x2a\xf9\xaa\x7e\x70\xf0\xa4\xfd\x20\xfd\xa2\x7e\x70\xd0\xa2\x1f\x34\xb4\x86\xfa\x10\xb1\x86\x23\x1c\xb4\xe8\x08\x0d\xad\x33\x75\xe5\x52\x53\x86\xe7\x2d\xca\xf0\xfc\xc1\x9d\x31\xdd\xd7\x19\x0f\x1a\x3b\xe3\xf2\xdd\xc6\xce\xde\x78\xd0\xd8\x1b\x97\xae\x7c\x6b\x71\x7e\xde\x90\xf3\x23\xc5\x01\xf2\xa7\x5b\x2d\x07\x02\x40\xb2\x71\x24\x20\xcf\xe5\x37\x86\x02\xe7\x64\x1a\xb3\xfb\xb6\xc2\x00\xc5\xae\x41\x1c\x00\xfd\xbf\xe6\x40\x40\xa9\xaf\xbd\x48\x40\x1b\xd0\xfe\xb1\x80\x21\xc3\xef\x70\xd2\xa6\x58\x92\xde\xca\x10\x85\x20\x59\xf5\x74\xbd\xbe\x5c\xe3\x8f\x78\xe6\x67\x16\xd3\x9e\x67\x7e\xe9\xed\x50\xbb\x67\x7e\x09\x6f\xe3\xcc\x4f\x8d\x71\xc8\xa3\x7d\xce\xfb\x86\xa6\xaf\x82\xb2\x02\xb8\xe5\x71\x6d\xc8\x62\xd8\xde\xf0\x86\x3b\x9c\x8c\xd7\x56\x04\x5b\x81\x7b\x06\xd4\x12\xfe\x35\x63\x5a\xc2\x9b\x61\xc7\x7f\x08\x8b\x53\x6b\x69\xba\xaf\x19\x5e\xa9\x35\x01\x7f\xa3\x58\xf4\x89\xc7\x08\x56\x1b\x97\xa6\x7b\x9c\xf7\xf2\x1b\x0e\xeb\x84\x74\xf7\x94\x2f\xdb\x84\xed\x21\x5e\x51\xbe\x73\xf9\x73\xd5\x75\xf2\xd9\xcd\x57\xc6\x3c\x09\x67\xfc\x1a\x36\x61\x4f\x79\x0b\x06\x40\xf4\x88\x58\x2c\xad\x78\x6f\x28\x4e\xf8\x43\x20\xf1\x7c\xdc\x76\xc0\x39\x1f\x37\x8e\x37\xe7\xe3\xcd\xe1\xe6\x47\x3c\x66\x58\x7e\x41\xa5\x9d\x78\x53\xb2\x6b\x10\x6e\xce\xc7\x5f\x33\x32\x4b\xe5\xb5\x17\xd5\x49\xe3\xd9\x15\x9e\x9f\x0a\x86\xc0\x12\x7a\x44\x0c\xd1\xc6\xba\x27\x86\xcc\xc7\x6d\x63\x08\xaf\x99\xe6\x95\x40\x44\x95\x83\x1a\xe3\x2c\x0a\x79\x5e\xb2\x49\x67\x38\xe8\x67\x79\x5e\x89\xcd\x41\x4d\xa8\xbf\xe9\xcb\x07\x26\x4d\x2a\x4f\x64\xb7\xc4\xaa\xb4\xc7\x4d\xa9\xc7\x81\x51\xa9\xa9\x35\x1f\x33\xa8\x41\x70\xc8\x68\xcc\xa8\xb8\xaf\xd9\xed\x82\x7a\x3b\xe6\xa0\xa5\x5d\xae\xe8\x7f\xeb\x76\xb9\x24\x1c\x2c\x2a\xda\x90\x55\x2b\xd3\xd5\xec\x0f\xf0\x52\x35\x5c\xd3\x75\xbf\xb1\xd4\xb2\x17\xd3\x81\x1b\x65\xa7\x59\xa7\xf7\x34\x7e\xd1\x3c\xe4\xea\x17\xcd\x60\x56\x2d\xfc\xe8\x57\x12\xfb\x19\x39\x37\xad\x11\xfb\x27\x10\xbb\x3a\xbb\xdc\xfb\x27\xd2\x92\xc6\x33\x18\xdd\x65\x4b\xf2\x3c\x47\xce\xc5\xa0\x25\x5a\x2f\x90\xf3\x71\x70\x79\xdd\x12\xb5\x97\x80\xdd\x67\x57\x2d\x11\xfb\x3b\x72\xae\x07\xe7\x67\x96\x9a\x59\x31\x7b\x93\xfb\x07\x72\x4e\x3f\x9c\x9f\x9f\x5c\xf4\xed\xf9\x71\x83\x7c\x45\x03\xcf\xcd\xd3\x15\xad\x07\x54\xeb\x50\x7e\xef\x44\x2e\xc5\xef\xd4\x37\x2d\x51\x07\xfd\xf4\xbd\x2a\x8f\x0a\x85\x87\xdf\x67\x8b\x56\x75\x29\x04\x96\xe5\x2d\xb2\xaa\x52\x47\xf7\x85\xc6\x1d\x7d\x9a\x8f\x80\xb6\xa3\x02\x23\xe3\xd5\x73\xcf\xe0\xd7\x91\x11\xcf\x54\x4e\x30\x77\xef\x30\x8b\xa8\xfc\xd8\xe8\xce\xdd\xac\x63\x29\x7c\x81\x31\xd7\x1e\xfb\xbe\xfd\x8a\xe6\xb1\xf1\xfc\x69\x00\x2a\x43\x5a\x1d\xb7\x95\xe3\x3b\xf9\x7d\x30\x08\x3a\x52\x56\xba\x95\xfa\xd0\xa2\x1b\xc4\x33\x26\xbf\xd5\x95\x7d\xb2\xc8\x41\x30\x91\x1e\x99\xc4\x21\x58\x49\xcf\xa8\xba\x73\x4e\x23\xc3\x50\x05\x90\x46\x6f\x12\x66\x4d\x31\xa8\x27\x5f\xf3\x11\x87\xb3\x42\xd5\x81\x32\x93\x4f\xfc\xf3\x69\x18\x73\x62\x54\xe2\xd8\x10\xa0\x86\x19\xe4\x92\x9a\xa1\x77\xf9\x03\x39\x21\x1e\x91\x70\xd5\xe7\x70\x3a\x56\x04\x78\xb0\xdf\x4c\xa9\xc1\x75\x3e\x56\x07\xff\x8e\x8e\x10\x6e\xc9\xbd\x8c\xb9\x75\x91\xf5\x59\x20\x45\x1c\x4a\xba\x52\x8e\x97\xa9\xbf\x32\xdf\x62\xd4\xdb\x05\x65\x4a\x6b\x1d\xd3\x0e\x52\xa4\xa1\x8e\x89\x65\xa6\xd8\x9b\x64\x2e\xc9\x70\x39\x87\x42\xd9\xaf\x94\xbb\x0e\xa5\x9d\x9b\x9b\x7c\x50\x33\xd7\xed\x4d\x64\x33\x9b\xd1\x95\x09\xec\xb6\xdf\x6a\x7f\xab\xaa\xd7\x7f\xf2\x40\x55\x0f\x59\x3c\x07\x2f\xcf\x36\x34\x81\x28\x4c\x6c\x22\x61\x86\x16\xd0\x50\xee\x1a\xc5\x64\x9b\x77\xdf\x41\x97\x1c\xc2\x45\x6f\xb2\x4b\xe0\x48\x03\x44\xfe\x40\xd5\x99\xf4\x91\x60\x33\xb2\x06\x67\x4b\x96\x83\x0a\x14\x37\xb5\x4f\xf1\x72\xdf\x88\xc5\x86\xd8\x2d\xfe\xe8\xa7\x5f\x4a\x78\xcf\xaa\xe4\xd4\xaf\xad\xb4\x13\xbf\xb6\x81\x9a\xf6\xb5\xb5\x7d\xca\x88\x27\xd4\xb5\x4e\x4b\xd9\xf3\xe9\xac\x37\xf0\x47\xa9\xeb\x04\x0a\xe6\x33\xad\x60\x2f\xef\xa1\x44\x6e\x26\xe5\x17\xa9\xa3\xd3\x90\x7a\xb7\xbd\x85\x98\x50\xde\x99\xe0\xc8\x0f\x89\x2a\xb1\x10\x24\xdb\x66\xc7\x52\xbb\xf0\x2c\x39\x6d\x39\xb3\x93\x3c\xee\x99\xef\x40\x00\xed\x93\x90\x8e\xa3\x53\x5d\xae\xe7\x77\xa5\x8c\x59\x5f\x53\x6e\xba\xa4\xb0\x95\xc9\x78\x90\x1d\x96\xe4\xfa\x9c\x62\x46\x44\xbe\x93\xd2\x62\xfe\x8b\x62\xd8\x28\x76\x92\x8e\xf0\x7f\x01\x00\x00\xff\xff\x87\xf1\xab\xe5\x8e\x5c\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 23694, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"index.html": indexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"index.html": &bintree{indexHtml, map[string]*bintree{
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, filepath.Join(name, child))
                if err != nil {
                        return err
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

