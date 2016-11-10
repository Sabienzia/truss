// Code generated by go-bindata.
// sources:
// NAME-service/NAME-cli-client/client_main.gotemplate
// NAME-service/NAME-server/server_main.gotemplate
// NAME-service/generated/cli/handlers/client_handler.gotemplate
// NAME-service/generated/client/grpc/client.gotemplate
// NAME-service/generated/client/http/client.gotemplate
// NAME-service/generated/endpoints.gotemplate
// NAME-service/generated/transport_grpc.gotemplate
// NAME-service/generated/transport_http.gotemplate
// NAME-service/handlers/server/server_handler.gotemplate
// DO NOT EDIT!

package template

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	name    string
	size    int64
	mode    os.FileMode
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

var _nameServiceNameCliClientClient_mainGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xef\x6f\x9b\x48\x10\xfd\x0c\x7f\xc5\x1c\x6a\x24\x5c\x39\xd0\x93\xee\xbe\x58\xca\x87\x34\xd7\x5e\x23\x5d\x2b\x2b\x89\xd4\x2f\x27\x9d\x08\xac\x61\xaf\xb0\x4b\x77\x17\x27\x15\xf2\xff\x7e\x6f\x16\x8c\xed\xc6\x96\xef\x57\xa4\x04\x58\x66\xde\xcc\xbe\x79\x6f\x49\x9b\xe5\x5f\xb2\x52\x50\x93\x49\x15\x86\xb2\x69\xb5\x71\x14\x87\x41\x24\x54\xae\x0b\xa9\xca\xf4\x4f\xab\x55\x84\x85\x55\x9d\x95\xfe\xda\x38\xbe\x68\xcb\x7f\xad\x33\xb9\x56\xeb\xf1\x16\xe1\x7e\xd5\xc9\x46\x44\x21\x6e\x4a\x5d\x67\xaa\x4c\xb4\x29\xd3\xe7\x54\x09\x97\x22\xd8\x89\x67\x0f\x50\x6a\x5d\xd6\x22\xd9\x0b\x29\x4d\x9b\x0f\x69\xd2\x55\xdd\x63\x92\xeb\x26\x6d\xbf\x94\xa9\x30\x46\x1b\xcb\x6f\xd2\x94\x1e\x2a\x69\xe9\x5e\x98\xb5\xcc\x45\x18\xe4\xb5\x14\xca\x7d\xc8\x54\x51\x0b\x43\x51\xdf\x27\xb7\x7e\x0b\xcb\xcc\x55\x74\xb9\xd9\x50\x5a\x0a\x25\x4c\xe6\x44\x91\x22\x36\xad\x86\x48\xee\x92\xab\x0d\xe9\xe7\xf3\x10\x34\x76\x17\x54\xce\xb5\xff\x24\x8d\xe3\x91\xd6\x3e\xfa\xf0\xe5\xdb\xc3\x84\x28\x9c\x85\xe1\x3a\x33\x4c\xf9\x1f\x74\x45\x23\x9f\xc9\x32\x33\x56\xdc\x2a\x37\xad\x32\xb5\xc9\x7d\x5b\xcb\x71\x89\xa7\x92\xdc\xe8\x06\xf3\x1b\x57\x06\x96\x92\xcf\x26\x6b\x57\xc3\x4a\xfb\x98\xdc\x89\x52\x5a\x27\x0c\x4a\x8f\x9c\x25\x9f\xb2\x46\x6c\x36\xfc\x24\x0c\x57\x0f\x57\x9d\xca\xfd\xfc\xe3\x19\xf5\x23\xc7\x82\xb2\xa2\xc0\x0e\xa8\x35\xc2\x76\x8d\xb0\xa4\x34\xd9\x01\x81\x0a\x69\x73\x8d\xec\x6f\x64\xbf\x01\xbc\x99\x13\x58\x25\xf1\xdc\x8a\xdc\x59\xea\x10\x66\xc9\x69\x8f\xd4\x1a\xbd\x96\x85\x20\x57\x71\x9a\x41\x00\x03\x03\xd3\x92\x5e\x21\x6d\x8b\x99\x0c\x73\x1d\xaa\xb5\x4e\x6a\x45\x78\x34\x62\x55\x23\x45\x14\x04\x71\x32\x1c\xc3\x70\x57\x8f\x52\x65\x28\xcf\x65\x79\x69\x5c\xe6\x99\x8c\x7a\xb6\x0b\xbf\x78\xe9\x4c\xa6\x2c\x13\x9e\x70\x59\x62\x0d\x5b\x8f\xc4\xa9\xe0\x5d\xea\xce\x6e\x53\x41\x3c\x88\xee\x72\x07\x1a\xe9\x51\x63\x40\xc3\x96\xa8\xd2\xd6\x2d\xbc\x31\xb6\x83\x80\x14\xc7\xa1\x79\x3d\x5c\x33\xf6\xf0\x73\xe5\x6b\x24\xf7\x3e\x30\x8e\xf8\xad\x2f\x1d\xcd\x29\xe2\xdf\x0f\x0f\x0f\xcb\x03\x0a\x8a\xc2\xae\xf3\x68\x06\x24\x16\xd8\x69\x24\x7e\x3b\x21\x2d\x7e\x7e\xf3\xd3\x1b\xbe\x29\xef\x96\x37\x14\x33\xe8\xec\x04\x6a\x23\x5c\xa5\x0b\xa2\xe3\xa8\xc3\x5b\x46\xea\x7b\x50\x85\x83\xe0\x95\x54\x85\x78\x9e\xe3\x4a\x8b\x2b\x9a\x54\xf3\xd1\x07\xda\xcd\xa6\xef\xe5\x6a\x0c\xe2\x07\x51\x5b\xc1\xd7\x07\xfd\x9b\x7e\x82\x05\x5f\xc9\x51\x60\x78\xa5\x8a\xe9\x12\xcd\x83\x7f\x57\x61\x3e\xe1\x9c\xaa\xc0\xbb\x9c\xed\x0d\x04\x52\xbf\xf1\x03\xbd\x36\x98\xd4\x75\x5d\xbf\xe7\xa1\x6f\x36\x1c\x15\xf8\xcd\x7b\x73\xc5\x2c\xfd\x29\x69\x2b\x6d\x58\xa6\xef\x7f\xd5\x5c\x80\x8e\x3b\x26\x08\x60\x34\x4f\xa5\x37\x9c\x47\x45\xbf\xaf\x27\x1d\xfc\x70\x85\x49\xb3\x91\xb6\xa0\x73\x8e\x04\xf1\xbb\x93\x23\xf9\x24\x9e\xe2\x29\x03\x00\x1b\x62\x1e\x89\x71\x26\x15\xec\x70\x20\x4c\x35\x80\x80\x2f\xaf\x83\x5f\x64\x56\xc7\x53\xe8\x7c\x58\xfc\x8c\x53\xf3\x56\x59\x91\x77\x06\xbb\xdb\x5b\x7c\xc0\x71\xac\x3b\x17\xf3\xb1\x8c\x4d\x01\xae\x98\xb1\x34\x50\x8e\x41\x51\x48\xc9\xda\x57\x0a\x70\xb8\x27\xef\x5b\x68\xc3\xad\x62\x8d\x03\xc7\x15\x88\x80\x3a\xde\xf1\x5e\xe9\xa9\x92\x35\xfb\x38\xab\x21\x1e\x8f\xcf\xa6\x51\xf0\x08\x0c\xbb\xa0\x8b\x75\xe4\xdb\x64\xec\x00\xd9\xef\x9e\xa5\x8b\x7f\xe4\x27\x90\x1f\x14\x62\x85\xe9\x71\x3c\xe6\xa3\xfd\x00\x5e\x50\xb4\x3b\x93\x3d\x45\x1c\xbc\x63\x87\x1b\x3c\xd5\x9f\x9f\xc5\x82\x0f\x29\x23\x1a\xed\xc4\x64\x06\x0b\x07\xcb\x95\x14\xc5\xef\xca\xdb\x61\xbf\xad\x4d\x78\x84\x82\x33\x15\x2e\xd6\x00\x9a\x76\x79\x88\x16\x06\xf6\x49\xba\xbc\xa2\xd7\xa3\xe9\xfa\xb0\xef\x09\x4b\x15\xbd\x42\x4b\x2c\x76\x96\xe1\xce\x07\xbc\x84\x37\x2f\x2d\xc0\x43\xcf\xb0\xe5\xe8\x88\xec\xa3\x05\xd3\xdb\xf7\x97\x03\xb2\xf7\x89\x47\xd9\x53\x3d\x03\xf1\xcd\x36\x8b\x3f\x35\x9c\xe5\xf5\xce\x1b\x1e\xb5\xeb\x81\x92\x8f\xf0\x43\x95\xed\x7c\x32\xa0\x7b\x73\xf1\x83\x11\x5f\x3b\x61\xdd\xa4\xc0\x83\x6f\x2e\xec\x32\x75\x16\xf7\xfd\xdf\x6f\x89\xfd\x9b\xdc\x64\x75\xcd\x8b\x93\x99\xbd\x74\x8e\xe8\xf2\x8c\x30\x73\xe0\xb0\x24\x4f\xb6\xf6\xfd\xe0\xbe\xd3\xa7\x1f\x1e\xd8\x99\xf6\xb8\xfd\x2c\xed\xef\x6e\xfc\xdf\x25\x79\x8b\x4f\x4c\x69\x74\xa7\x0a\x76\xd9\xc8\xce\x7f\x6c\xfc\x48\xbd\xf3\x2d\x8f\xe8\x4b\x06\xaf\x55\x1c\x0d\x6c\xd3\xdd\xd0\x11\xbe\x9a\x3c\x8d\x85\x97\xfd\x41\xe0\xff\x34\xa6\x83\xda\xc3\xe9\x88\xda\xf8\xda\x02\xf7\x64\xed\xf5\x2c\xdc\x53\xd7\x9e\xce\x70\x3c\x64\x5d\xed\x16\xe7\x2d\x28\xd5\x1a\x07\x50\x41\xa3\xc9\x2e\xbe\x7a\x96\x86\xa7\x97\xa6\xc4\x64\xff\x0a\x00\x00\xff\xff\x6c\xf0\x67\x0a\xe4\x0a\x00\x00")

func nameServiceNameCliClientClient_mainGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceNameCliClientClient_mainGotemplate,
		"NAME-service/NAME-cli-client/client_main.gotemplate",
	)
}

func nameServiceNameCliClientClient_mainGotemplate() (*asset, error) {
	bytes, err := nameServiceNameCliClientClient_mainGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/NAME-cli-client/client_main.gotemplate", size: 2788, mode: os.FileMode(436), modTime: time.Unix(1478723139, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceNameServerServer_mainGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x4f\x6f\xdb\x36\x14\x3f\x4b\x9f\x82\x15\x3a\x40\x1a\x1c\xaa\xc3\xd6\x1d\x82\xe6\xd0\xa6\x59\x1a\x2c\x09\x8c\xd8\xc3\xce\xb2\x44\xd3\x44\x24\x52\x20\xa9\xd4\x86\xe0\xef\xbe\xf7\x48\xca\x96\x53\xdb\xd9\xa1\x45\x63\x51\x7c\x3f\xfe\xf8\xfe\x3f\xb5\x45\xf9\x5c\x70\x46\x9a\x42\xc8\x38\x16\x4d\xab\xb4\x25\x69\x1c\x25\xcb\xba\xe0\x09\x3e\x1b\x8b\x0f\xc9\x86\x47\xbe\xb2\xb6\x1d\xaf\xf3\xb6\xd5\x6a\x89\x3b\xca\xf8\xdf\xdc\x08\x2e\x8b\x1a\x5f\xcc\xc6\x94\x45\x0d\xcb\x38\xca\x73\xf2\x7b\x45\xa6\x85\xb6\x1b\x10\x70\x55\x17\x92\x53\xa5\x79\xbe\xce\x91\xaa\x54\xd2\xb2\xb5\xbb\x85\x2b\xc5\x6b\x46\x47\x10\xae\xdb\x32\x70\xdc\x2a\xf2\xb7\xb0\x88\x12\x76\xd5\x2d\x68\xa9\x9a\x9c\xab\x8b\x67\x61\x73\xfc\x63\xb2\x6a\x95\x90\x9e\xe7\x28\xa2\x56\x3c\x50\xcd\x57\xc2\x90\x19\xd3\x2f\xa2\x64\x71\xb4\x2a\x64\x55\x33\x4d\x92\xbe\xa7\x77\xce\x11\xd3\xc2\xae\xc8\xc5\x76\x4b\xf2\x20\x03\xcb\x00\xcd\x34\x90\x9b\x97\xf2\x28\x92\x33\xc9\x74\x61\x59\x05\x98\x76\xe1\x20\xd3\x2f\x87\x20\xb8\x3d\x8b\xe3\x65\x27\x4b\xe7\xf6\x34\x23\x7d\x1c\xbd\x14\x1a\xfd\x1e\x55\x6c\xd1\xf1\xcf\x55\xa5\x89\xfb\x77\x45\x30\x10\x74\x66\xb5\x90\x3c\x4d\x9c\x94\x16\x20\x4e\x26\x24\xb9\xfc\xf8\xe1\xcf\x0f\xb8\xf8\x8a\xdb\x04\x74\x24\x0d\x03\x64\x69\x48\x2d\x8c\x65\x92\x20\x92\x19\x93\x64\xc0\x8c\xb1\xda\x13\xbf\x66\x46\xe9\x98\xf8\xa3\x23\xfe\x36\x9f\x4f\x8f\x71\x61\x3c\x4e\x73\xa1\x74\xcc\xf5\x87\xe3\xe2\x4f\xd3\x6b\x92\x22\x63\x76\x84\x12\xfe\x3b\x0e\xc8\x0f\xc3\xd2\xcc\x07\xe8\x5e\x71\x0e\x8c\xa4\x52\xe8\x28\xea\xbd\x04\xf1\xe3\xcc\x3d\xe8\xbd\x5b\xc6\x11\xf8\x2f\x0a\xdb\x57\x4e\xf0\xc8\xbe\x83\x0c\x72\xd7\x23\x52\x65\x40\xb9\x4a\x75\x36\x3b\x82\xbc\xf6\x99\x97\xfa\xfd\x8c\xfe\x0b\x69\x93\x26\xd6\x80\xd2\x88\xf8\xca\x96\x45\x57\xdb\xb9\x68\x98\xb1\x45\xd3\xfe\x33\xbf\xfe\xff\x2c\x98\xfc\x4c\x1f\x32\x5d\xbb\x3d\xe0\xd8\xc6\x81\x05\x0d\x49\x93\xc6\x70\xf4\xd3\x8a\xd5\xb5\x42\x97\x54\x6c\xc9\x06\x73\x0f\x10\x50\x1f\xd5\x62\xc3\x92\xe0\xa5\x2f\x9d\x11\x12\xdc\x78\xe8\x26\xe3\xf3\x9a\xb4\x0b\xda\xf7\xb7\xea\xb1\x68\x18\xa1\x21\xd9\x29\xbe\x6d\xb7\x33\x97\xcc\xde\x7d\x03\xfc\x8a\x84\x5c\x47\x9b\x02\x3c\x45\x7b\xe1\x22\x08\xd7\x8e\xb6\x0e\xa1\x19\x27\xdd\x8a\x69\x86\x56\x39\xad\x6e\x42\x21\xee\xb5\xea\x7b\x0d\xf5\xcc\xc8\x7b\x41\x2e\xaf\xf6\xba\x3c\x30\xbb\x52\x95\xd9\x6e\xbd\xde\x7d\x3f\x57\xf7\xea\x3b\x58\xfe\x5e\x04\x3d\x77\x54\x43\x71\xd3\x61\xc7\xeb\x7e\xee\xc8\x15\x81\x4a\xa5\x0f\xc5\x33\xeb\xfb\x1f\xa4\x69\xb0\x26\xd8\x07\x19\xbd\xbb\x82\x58\x5d\x94\x60\xe0\x84\x08\x69\xac\xee\x1a\x26\x6d\x61\x85\x92\xce\xe2\xc1\xfa\xc1\x62\x50\x01\x0e\xa2\x09\xc3\x79\x83\x36\xe2\xd5\xc3\x5d\xa6\x7f\xcb\x03\xd8\x1b\x9c\x35\x3f\xe8\x79\x89\x45\x76\xc6\xca\xc9\x48\x81\xe0\xfe\x07\x56\x42\x20\x05\xe4\xdf\x3e\x00\x4c\xeb\x12\x2f\x6e\xc0\x1b\x29\x8a\x09\xec\x28\xcc\xc4\xd2\xae\x51\x10\x7a\x30\xfd\x02\x53\x81\x6b\xd5\xc9\x6a\xa8\xc4\x3b\x10\x68\xdd\xb5\x76\x97\x1e\x71\xc4\x15\xc1\x26\xe6\xfb\x57\xf4\x8a\x19\x2b\xce\x4d\x81\x09\xf9\x0d\xdd\xeb\x47\x02\x7d\x54\x56\x2c\x37\x69\x39\x21\x61\x32\xd0\xd9\xdd\xed\xdd\xe3\xfc\xe0\x7d\x7e\xf3\xf4\x80\x67\x9c\xbe\x9f\x2e\x08\xd4\x31\xbd\x41\x4d\x97\x69\xf2\x0b\x96\xe5\xa7\x8b\x12\xcb\x67\x50\xce\xb7\x3f\xdf\x53\x8e\x68\x16\x2a\xf5\xf2\xad\x82\x87\xd8\x18\xec\xd3\x58\x62\xae\xd1\xba\x02\x8b\x1a\x3c\xe9\xda\x63\xa8\x08\xf6\xd0\xad\x5d\x49\x34\xf4\x9b\x73\x46\x9a\xe4\x0e\xef\x27\x61\x0e\xe7\x1d\xdc\x0b\xf5\x5f\xa8\x89\x93\xd0\x3b\x59\xb1\x75\x76\xe6\x68\xd9\x54\x35\xd4\xf2\x69\x86\x6b\x0f\x38\xc7\x81\x3f\xa2\x3e\xc3\x31\xf5\x80\x73\x1c\x66\xd3\x2c\x54\x7d\x9a\x62\xe6\xe4\xe7\x18\xb0\x7c\xce\xe8\x30\x47\x71\xe6\xfc\x3b\x6e\x70\x61\x68\xfc\xba\x9b\x82\xe3\x34\x70\x54\xf7\x2e\xca\x9f\x65\xe5\x22\x91\xee\x91\x13\xd2\x8c\x73\xc2\x4d\xae\x5d\x48\x7f\x4a\x4e\x20\xa5\x9f\xa4\x43\x6d\x63\x5b\xc1\xdd\x60\x5f\x0a\x65\x34\xd9\x75\x10\x33\x09\xbd\x3b\x3b\x61\xe4\x30\x90\xdf\xb4\x71\x00\x82\x37\xc7\x26\xba\x81\xfa\x73\x4d\x44\x4a\x9f\xf5\xb5\x9c\x60\x77\xc0\xe3\xf0\x75\x16\x54\x02\x74\xd9\xa2\xea\xc3\xfc\x47\xd5\xc5\xd2\x01\xdf\x01\x50\xd4\xee\xe6\x9d\x35\xf0\xc4\x57\xcd\x6c\xa7\x25\xac\xb0\x37\x45\x46\xbf\x8c\xfd\x77\x0b\x57\xfa\x41\xf4\xca\x7d\xae\x6b\x20\xd2\x7d\x4e\x0c\xb5\xa7\x5d\xe5\xc1\x50\x7b\x62\x1c\x75\xd2\xf0\x81\x75\x6c\xaa\xa5\xe0\x7e\xb8\xe9\x64\x86\x8d\x2d\x18\xd4\x35\xd4\x3b\xbc\x96\x63\x37\x3f\x75\xf2\xdd\xe1\xa0\x66\x6b\x61\x5d\x17\xc2\x83\x59\xbc\x8d\xff\x0b\x00\x00\xff\xff\xf4\x3e\x62\x2e\x46\x0b\x00\x00")

func nameServiceNameServerServer_mainGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceNameServerServer_mainGotemplate,
		"NAME-service/NAME-server/server_main.gotemplate",
	)
}

func nameServiceNameServerServer_mainGotemplate() (*asset, error) {
	bytes, err := nameServiceNameServerServer_mainGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/NAME-server/server_main.gotemplate", size: 2886, mode: os.FileMode(436), modTime: time.Unix(1478128977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedCliHandlersClient_handlerGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x52\x4d\x6f\xea\x30\x10\x3c\xc7\xbf\x62\x85\x10\x4a\x9e\xc0\xdc\x91\xde\xe1\xf1\xd4\x22\x0e\xad\x10\xed\x1f\x30\xc9\x02\x56\x13\x27\xb5\x9d\x7e\xc8\xf2\x7f\xef\xda\x98\x50\x7a\xa1\x17\xb0\x77\x66\xc7\xb3\xb3\xe9\x44\xf9\x22\x0e\x08\x65\x2d\x51\xd9\xa3\x50\x55\x8d\x9a\x31\xd9\x74\xad\xb6\x90\xb3\xac\xdb\xc1\xc8\xb9\x31\xdf\x2c\xd7\xb1\xb6\x11\xf6\x08\x33\xef\x47\xac\x60\xcc\x39\x78\x97\x74\x1f\x5b\x84\xc5\x5f\xe0\xde\xb3\xcc\x39\x2d\x14\x29\x8e\x65\x28\x11\xc2\x9f\x50\xbf\xc9\x12\xf9\x03\xda\x63\x5b\x99\x40\x9a\xcf\x81\x44\x25\x7f\x14\x0d\x7a\x0f\xf4\x5c\x8d\x0d\x19\x30\x70\x26\xb3\x6c\xdf\xab\xf2\x3b\x2b\x77\x2e\x3e\x26\x55\x85\x1f\x51\xf8\x7f\x34\xfd\x4f\x1f\x4c\xd4\x0e\x07\x18\xe8\xce\xad\xda\x70\x02\x7e\x4f\x42\x56\xb6\x2a\xe0\xa1\x8e\xaa\xf2\xbe\x80\xfc\x4f\xb7\xe3\x03\x8b\xfa\xb6\xf8\xda\xa3\xb1\xcf\x9f\x1d\x26\x8d\x29\xa0\xd6\xad\x2e\x1c\xcb\x32\x7d\x42\xc3\x50\x37\xfb\x02\xdf\xb9\x59\x0a\xa7\x21\x6f\xa1\xed\x57\xc6\x43\xb6\xd4\x7d\x89\xb1\x13\x5a\x34\x31\xca\xa0\xc3\x23\x37\x71\xe2\x1b\x72\x0f\xb4\x35\xc8\xe3\x8f\x6a\x6d\xea\xe0\x6b\xb3\x14\x06\x83\xa9\xe2\xaa\xbe\xc5\x0e\x85\xc5\xaa\xb8\x2e\xdf\xa9\xbe\x29\x92\x6c\x76\x99\xee\x04\xa6\x35\x2d\x60\xf2\x13\x59\xb5\x64\x88\x82\x1a\xec\x60\x6d\x10\x6e\xea\xdc\x96\xa1\x69\x86\x24\x2e\x97\x04\xc5\xb3\x8f\x4b\xb1\xbd\x56\x30\x49\xcb\x99\x82\x92\x35\xcb\xe2\x57\x78\xa2\x9d\xff\xbf\x02\x00\x00\xff\xff\x50\xcc\x31\x90\xe8\x02\x00\x00")

func nameServiceGeneratedCliHandlersClient_handlerGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedCliHandlersClient_handlerGotemplate,
		"NAME-service/generated/cli/handlers/client_handler.gotemplate",
	)
}

func nameServiceGeneratedCliHandlersClient_handlerGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedCliHandlersClient_handlerGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/cli/handlers/client_handler.gotemplate", size: 744, mode: os.FileMode(436), modTime: time.Unix(1478723139, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedClientGrpcClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x55\xdf\x6f\xdb\x36\x10\x7e\x16\xff\x8a\x9b\x51\x0c\x52\xa0\x50\xef\x1d\xfc\x52\x77\x2b\x3a\xac\x9d\xd1\x06\xdb\xc3\x30\x0c\x34\x45\xcb\x84\x6d\x52\xa3\x68\x27\x81\xa0\xff\x7d\x77\xfc\x91\xc8\x5e\x92\xa2\x41\xc4\xe3\x77\x77\x1f\xbf\xfb\x91\xa6\x81\xb5\x90\x7b\xd1\x29\xe8\x5c\x2f\xa1\x77\xf6\xac\x5b\x35\x80\x80\xee\xdb\x7a\x05\xf2\xa0\x95\xf1\xb0\xb5\x0e\xfc\x4e\xc1\x38\xf2\xef\xca\x9d\xb5\x54\xfc\xab\x38\xaa\x69\x82\x21\x1d\x59\x3f\x0b\xc3\x98\x3e\xf6\xd6\x79\x28\x59\xb1\xe8\xec\x41\x98\x8e\x5b\xd7\x35\x0f\x8d\x51\xbe\x91\xd6\x78\xf5\xe0\x17\xe1\xce\x76\x07\xc5\x67\x10\x72\x7f\xfd\xa6\x39\x2a\x2f\x5a\xe1\x45\x80\x68\xbf\x3b\x6d\xb8\xb4\xc7\xa6\xdf\x77\x8d\x72\xce\xba\x61\xc1\x2e\x6f\x3a\x7b\xbb\xd7\xbe\xa1\x1f\x65\xda\xde\x6a\x43\x89\x29\x96\x77\xc2\x0c\x81\xe5\x2b\xf8\x27\x40\x22\xc5\x0a\x54\xeb\x6e\xa7\x07\x48\x1a\xb0\x62\x38\x4b\x58\xa0\x28\x9f\xc3\x73\xd7\xc2\xef\xe0\x16\x45\x69\x3a\x65\x94\x13\x5e\xb5\x98\xab\xdf\x04\xc8\xfa\xc3\x25\x68\xc1\x2a\xc6\x30\xe0\x57\x75\x0f\x4e\xf9\x93\x33\x28\xba\xc9\x7a\xc2\x06\xe5\x54\x2d\x6c\x1e\xaf\x2a\x81\xe2\x19\x25\xbd\xb6\x86\xc3\x67\x0f\x48\x06\xeb\xc2\xb6\x27\x23\x29\x52\x49\xd7\x70\x43\x7c\xf9\x2a\x38\xac\xd0\x50\x83\xed\xc9\x63\x00\xce\x93\xf9\xf7\x60\xa8\xa0\xec\x37\x7c\x1c\x3f\x59\xaa\x26\x5c\xd5\x96\x4e\xca\xd5\x10\x84\xad\x60\x64\xc5\x59\x38\x90\x32\x51\xc1\xc8\x5b\xdd\xa1\x2a\xd4\x1c\xff\xd4\xb0\x85\xf7\x4b\x40\xc9\xb0\x07\x72\x3a\x74\x29\xd0\x9b\x2e\xb6\xe5\x8f\x52\x56\x78\xd6\x5b\x0a\x08\x3f\x2c\xc1\xe8\x43\x40\x14\xf1\xf9\x74\x4e\xc9\x06\xfe\xa7\x13\x7d\x89\xdf\x35\x2c\xa4\x30\xc6\x7a\x10\x7d\x7f\x78\x4c\x91\x17\x14\x68\x62\xf8\x9f\x15\x72\xf6\x9e\x81\x32\xfd\xf5\xf7\x45\x75\x2f\x1e\x4c\xe9\x5e\xba\xfd\xa0\xf0\x11\xaa\x24\x32\xa9\x3b\xff\x10\x87\x93\x1a\xee\xec\x27\xd4\xfe\x4b\x6a\xba\x52\x4a\xbe\x53\xa2\x55\x6e\xa8\xaa\x9a\xd2\x17\xe3\x78\x0b\xf7\xd8\x3d\xf0\xce\x2b\x4a\xce\x27\x34\xce\xac\xd8\x98\x41\x5a\xbc\x42\x04\x4f\xb3\x16\xf5\xa5\x6c\x84\x8c\x9a\xbd\xd3\x19\x94\xab\x80\x69\x77\xb6\x1d\x22\x30\x68\x3f\x8e\x77\xf6\x37\x7b\xaf\x1c\xa2\x53\x91\x7e\x4e\x4d\x0d\xb9\xbb\x79\xb6\x04\xaf\xa0\x2f\xa5\x79\xdd\x71\x09\x97\x8a\x60\x1b\x45\x51\xca\xe8\x4b\x8a\x98\x3a\x7d\x63\x27\xe7\x37\x4d\x13\x76\xce\x9c\x6f\x34\x2e\xe6\x50\x7d\x6d\xc4\x81\x41\x82\xd2\xb6\x8a\x84\x9d\x21\xbe\xa9\x7f\x51\x6f\x3f\xc7\x7d\x54\x2f\xe2\x90\xa5\x19\x54\x06\xce\xfb\x17\x41\xf9\xfa\xee\xb1\xcf\x84\xc6\x29\x63\x2f\x5a\x05\x47\x21\xd9\xab\x27\xc9\xca\x2a\x58\x52\x65\x50\xd1\x54\xcd\xf4\x95\x3f\x58\xee\xd8\xf8\x9a\xe8\x3b\x8c\x04\x98\xd7\xf2\xba\x90\x34\xf5\x21\xdc\xff\x6a\xf0\x1e\xf0\xdf\x1b\x35\xaa\x9f\x73\x17\x53\x4d\x83\xc2\x90\x85\xc7\x37\x5e\xcc\x22\x0c\xde\x9d\xa4\xa7\xa1\x4a\x6d\x8a\xc3\x80\x36\x6d\x3a\xc2\xe3\xb2\x99\xcf\x02\xed\x0e\x01\xb4\x39\xc2\xc9\xef\x84\x87\xa3\x6d\xf5\x56\xab\xb0\x54\x66\x1b\x87\xe6\x3c\x64\xbb\xf0\x27\xd7\xf2\x66\x4e\xa0\x8a\xe3\xcb\xe2\x3e\x5a\xf9\x87\x3c\x45\xdf\x91\x7c\xb9\x57\x8f\x61\x03\x45\x46\xd5\x65\xb0\xf1\x49\xd4\x10\xd6\xc2\x4b\x81\xc3\xba\xb0\x79\x06\xb1\x73\x29\x24\x9b\x2f\x10\x1a\xca\x29\xe5\x7f\x6b\x92\x03\x97\x2c\x4e\x75\x35\x01\xa9\x17\x7f\xa1\x20\x57\xbc\xa4\x7f\xc8\x71\xf9\x2a\xfe\xae\xe1\xd8\xc2\x4d\xfe\xc3\xc4\xbf\x7c\xac\xae\x11\x81\x36\xcd\x6f\x2f\xf4\xbc\x26\x45\x5e\x9d\xfb\xe7\xd5\x19\x88\x85\xa9\xc5\x45\x79\xc6\xcd\x1d\xee\x30\x2d\x0f\xef\x28\xf7\x15\x2f\x13\xeb\x9f\xe8\x32\x0e\x78\x0c\xbc\xa4\x25\x49\x4a\x87\x23\x86\xad\xe1\x1c\x3a\x7a\x0a\xeb\x32\x2e\xdf\x08\x9d\xaf\xdf\x1b\xe4\xbf\x84\xa7\x07\xfc\x8a\x1d\x57\xa2\xad\x7e\x36\xad\xc9\x27\x46\xc5\xfa\x55\x55\x0e\x97\x94\x41\x76\x51\xf7\xff\x02\x00\x00\xff\xff\x3a\x03\xa3\xb4\x4e\x08\x00\x00")

func nameServiceGeneratedClientGrpcClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedClientGrpcClientGotemplate,
		"NAME-service/generated/client/grpc/client.gotemplate",
	)
}

func nameServiceGeneratedClientGrpcClientGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedClientGrpcClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/client/grpc/client.gotemplate", size: 2126, mode: os.FileMode(436), modTime: time.Unix(1478128977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedClientHttpClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2c\xcb\x31\x0e\xc2\x30\x0c\x05\xd0\x9d\x53\x78\xae\x44\x7c\x08\x96\x8e\x48\xe4\x02\x51\xf9\x84\x0a\x13\x47\xce\x67\x8a\x7a\x77\x18\xd8\xdf\x9b\x53\x17\xb9\x01\x52\xfd\xcc\xf8\x8c\xa1\x15\xad\xfa\x6b\xa7\x3e\xc9\xce\x28\x6d\x74\x0f\x2a\xf1\xee\x56\x88\x54\x5d\x1e\x1e\xb2\xf9\x1d\xb2\xe8\x71\x9c\xe6\xdc\x8a\x99\xa4\x35\xe7\xeb\x0a\xeb\x88\x74\xb1\x1d\x8d\xf9\x5f\x24\xfd\xd4\x37\x00\x00\xff\xff\x0b\x3c\x4c\x9e\x69\x00\x00\x00")

func nameServiceGeneratedClientHttpClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedClientHttpClientGotemplate,
		"NAME-service/generated/client/http/client.gotemplate",
	)
}

func nameServiceGeneratedClientHttpClientGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedClientHttpClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/client/http/client.gotemplate", size: 105, mode: os.FileMode(436), modTime: time.Unix(1477930115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedEndpointsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x55\x5f\x6f\xdc\x36\x0c\x7f\x3e\x7f\x0a\x2e\x08\xb0\xbb\xc1\xd1\xbd\x0f\xe8\xcb\x8a\xfd\xc9\x43\x8b\x62\xed\x17\xd0\xc9\xb4\x2d\x44\x96\x5c\x49\xbe\x24\x33\xfc\xdd\x47\x4a\xb6\xe3\xec\x8c\xa1\x05\x1a\xdb\x14\xf9\x23\xf9\xfb\x91\xba\x5e\xaa\x27\xd9\x20\x84\xab\x2a\x8a\xf3\x19\xbe\xb5\x3a\x40\xad\x0d\x82\x72\x36\x4a\x6d\x03\x74\x18\x5b\x57\x05\x88\x0e\x3a\xf9\x84\xa0\x6d\xa5\xaf\xba\x1a\xa4\x01\xb4\x55\xef\xb4\x8d\x14\xe2\x5d\x07\x01\xfd\x55\x2b\x0c\x25\x23\x79\xfc\x3e\x60\x88\x20\x6d\x45\xef\xa1\x77\x36\x20\xc4\xd7\x1e\x13\x12\xbb\xd2\x67\xeb\xc8\xb8\xa2\x94\x20\x03\x3c\xa3\x31\xfc\x44\xab\x5c\x85\x3e\x30\x00\xe3\x55\x38\x7f\xd7\xce\xcf\x81\x09\xad\x4c\x06\x49\x41\xae\x06\x37\x78\x08\x43\xdf\x3b\x1f\xb1\x82\xe8\xa5\x0d\xfc\xce\xe9\xb4\x34\xfa\x1f\x19\xb5\xb3\x8c\x46\x31\x9d\x8c\x41\xc0\x23\x55\x68\x82\xa3\xae\x94\x19\x2a\x0c\x6b\x35\xd0\xe9\xaa\x32\xf8\x2c\xa9\x78\x51\x14\xba\x4b\x40\xc7\xe2\x70\xd7\x38\x23\x6d\x23\x9c\x6f\xce\x2f\x67\x8b\xf1\xcc\x54\xe1\x4b\xbc\x2b\xf8\x50\xc7\x76\xb8\x08\xe5\xba\x73\xe3\x1e\x9e\x74\x3c\xf3\xff\x05\x94\x5d\xfa\x0b\xdc\x8d\xa3\xf8\xf2\xdb\x63\x82\xfc\x22\x63\x0b\x0f\xd3\x74\x57\x9c\x92\x02\xbf\xaf\x9c\x2a\x67\x0c\x2a\x7a\x99\x9b\x8b\xed\x86\x2b\xfa\x92\x91\x5c\x08\x83\x98\x90\x16\x64\x55\x2d\x02\x70\x57\x3f\x07\x06\xeb\x50\x52\x27\xc4\xf7\x05\x61\x08\x44\x09\x11\x2b\xa1\x45\xd3\x23\x11\x15\xfd\xa0\x62\xc9\xc7\x73\xaa\xfd\x4c\xf4\xc7\x81\x64\xb8\xa0\x6d\x43\xa3\xd1\x4b\x2f\x69\x2a\xd0\x0b\x32\xb2\xfd\x91\xd2\x67\x49\x7d\x09\x9a\x72\x73\xb2\x7a\x30\x49\x9a\x7a\xb0\x8a\x69\x9f\x4b\xb6\xc8\xca\x38\x70\x54\x82\x8c\x08\x8e\x63\xe9\xfd\x61\x49\xc8\x80\x17\x19\x34\x89\xf3\x07\x85\xe3\x8b\xec\x7a\x83\x25\xbc\xba\x81\x34\x69\xda\x48\xf9\x03\x8f\xc5\x86\x2a\x2e\x70\x4d\x94\xf3\xf4\xde\x55\x03\x0d\x23\xc3\x91\x6f\x1b\x63\x2f\xfe\xa2\x59\x32\x5c\xe3\x33\xa9\x04\x28\x55\x3b\x4f\x37\x1c\x97\xec\x27\x3a\xf3\x54\xe1\xd0\x67\xd0\xd0\xa3\xd2\xb5\x56\x94\x34\xb6\x02\x8e\x8f\xa9\x3e\x5a\x12\xc2\xbf\xc8\x8b\x79\x25\x9f\x4e\x87\x98\x37\x83\xa6\x34\xe8\xc6\x72\xa8\xb6\x57\xf7\x84\x89\xca\xaf\x59\x96\x75\x93\x52\x89\xf8\x5e\xec\x2c\x06\x43\x2c\x4c\x8a\xd3\x96\x5d\x65\x34\xda\xf8\x9e\xdd\x8d\x70\x6f\x4b\x49\x15\xd1\x3c\x66\x38\xea\xe3\x7f\x64\xe4\xf5\xc9\x5c\x69\x66\xb8\xc3\x3c\x56\x6f\xf5\x52\x04\xfa\x5a\xf2\x40\xed\x2b\xc1\x60\x6b\xb2\xfd\x8b\x61\xe0\x64\x6f\x9b\x78\x4e\x3a\x7c\xc6\xe7\x8f\x73\x3f\x34\xc1\x17\x6d\x13\x4f\x5d\x62\x36\x55\xb9\xd1\xb6\x9c\x6f\x90\x38\x78\x4b\xdd\x73\xd3\x5c\xa3\xa2\x4e\x69\x86\xd3\x3c\xcf\xf5\x8a\x22\x75\x74\xc3\xe9\x58\x8c\x23\xe5\xa7\x5b\xee\x5e\xc3\xaf\x1f\x40\x2c\xfe\x9f\xb2\x1e\xd3\x54\x1c\xc6\xf1\x5e\x8b\xcf\x34\xd5\xd3\xb4\xc4\x03\xfd\x5b\xfa\x10\x8b\x91\xa0\x1e\xd8\x4a\x31\xd3\xfb\x75\xfd\x81\x24\x3c\xa0\x34\x69\x6f\x31\x27\xd8\xe4\x3d\xaa\xf8\x02\xf3\x55\x22\x3e\xe6\x67\xc9\x03\xf1\x4b\x7f\x11\xe3\xf8\xa7\x63\x37\x42\x17\x7f\xe7\x9b\xf5\x1b\x35\x3b\x87\x9e\xe0\x78\xeb\x94\xaf\xdc\x8d\x57\x09\xe8\xbd\xf3\x94\xb4\x38\x1c\x96\x2b\x39\x19\xb9\x60\x14\x3b\x1c\x70\x4d\x5c\xc3\x89\x22\x74\x9d\x5c\x7f\xfa\x00\x56\x9b\x84\x71\x98\x55\xa1\xef\x04\x43\xa6\xa9\x58\xad\x4b\x06\xf1\x23\xb5\x9d\x4a\x46\x29\x28\x7e\x1c\x33\xbd\x4c\xee\x27\xde\xaa\x2d\xc3\x69\x6f\xef\xe9\xd6\x60\x86\xb3\x6e\x5b\xd2\xe9\x64\x8f\xf7\x4c\x3c\x83\xed\xb5\x48\xab\xcc\xe5\x6d\x63\xb3\xc7\xd7\xb4\x86\xa7\xdb\x21\x78\xd7\x3c\x63\xef\x4b\xb7\xfc\x02\xae\x6b\x34\xb2\x50\xeb\x6f\xe1\xc6\x9c\x45\xd8\xa8\xc3\xe8\xdf\xb9\xa3\x19\x63\x8f\xc3\x9b\x21\x48\x71\xd7\x55\xd0\x20\xfe\x33\x5c\xa9\xa2\xec\xb5\xa3\xe5\x9e\x9a\x59\xcf\xf5\xe4\x3a\x8b\x94\xcd\x89\xfd\xac\xd5\xf2\xfc\x37\x00\x00\xff\xff\x5a\xa4\x38\xcd\x4e\x08\x00\x00")

func nameServiceGeneratedEndpointsGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedEndpointsGotemplate,
		"NAME-service/generated/endpoints.gotemplate",
	)
}

func nameServiceGeneratedEndpointsGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedEndpointsGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/endpoints.gotemplate", size: 2126, mode: os.FileMode(436), modTime: time.Unix(1478128977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedTransport_grpcGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x57\xdf\x6f\xdb\x36\x10\x7e\xb6\xfe\x8a\x9b\x51\x0c\x56\x10\xcb\x7b\x2e\x90\x97\xa6\x5d\x53\x6c\xe9\x82\x2e\xd8\x1e\x8a\xa2\xa0\xa5\xb3\x4c\x58\x12\x15\x92\x76\x92\x09\xfe\xdf\x77\x47\x4a\xb2\xfc\x23\xb2\xda\x3c\x04\x0e\xc5\xfb\xf1\xdd\xf7\xdd\x51\x54\x29\xe2\x95\x48\x11\xcc\x26\x0e\x82\xd9\x0c\xee\x97\xd2\xc0\x42\x66\x08\xa5\x56\x1b\x99\xa0\x01\x83\x7a\x83\x7a\x6a\x68\x01\x73\x59\x24\xb2\x48\xc9\x44\x69\xb0\x4b\x84\xf4\xcb\xdd\x35\x58\x2d\x0a\x53\x2a\x6d\x23\x0e\xf1\xc9\xc2\xda\xca\x4c\xfe\x47\xbe\x6c\xd2\xee\xce\x52\x5d\xc6\xd1\xdf\x2e\x5c\x14\x04\x32\xe7\x87\x30\x09\x46\xe3\x02\xed\x6c\x69\x6d\x39\x0e\x68\x91\xaa\x4c\x14\x69\xa4\x74\x3a\x7b\x9a\xf1\x4e\xac\x0a\x8b\x4f\x76\xec\xf6\x54\x9a\x61\xd4\x31\xe1\x98\xb3\x1c\xad\x48\x84\x15\xec\xcf\x0f\xda\x94\x30\x4e\xa5\x5d\xae\xe7\x51\xac\xf2\x59\xaa\xa6\x2b\x69\x67\xfc\xb7\x8f\x89\xdd\x9a\xda\x19\x9e\x8c\x31\x18\x95\x73\x18\x57\x55\x74\xf7\xee\x93\xc3\x79\x27\xec\x12\xa6\xdb\xed\x38\x08\x03\xc7\xd4\xad\x58\xe1\x47\xaa\xde\xd7\x03\x39\x2d\x0d\x08\x62\xcb\x82\x5a\x00\x16\x49\xa9\x64\x61\xe9\xd1\x46\xc8\x4c\xcc\x89\x51\xc1\xfb\x8e\xb0\xaa\xfa\xa8\x3e\x8b\x1c\x21\xaa\xd3\x45\xbc\xda\x6e\x1b\x6e\x16\xeb\x22\x3e\x48\x30\x89\xed\x13\xd4\x4c\x44\xd7\xfe\xf7\xb2\x93\xe6\x43\xf3\x5f\x08\xe5\x3c\xea\x4f\x00\x55\x30\xf2\xaa\xfe\x55\x5a\xa9\x0a\x03\x6f\xaf\xe0\xeb\xb7\x3d\xe6\x6a\x9d\xbc\x01\xd9\x8f\xe0\xd4\xf6\x3b\xa4\x46\xc0\x49\xc3\xff\xbd\xaa\x91\x85\x97\xc1\x68\x1b\x8c\x34\xda\xb5\x2e\xe0\x57\x76\xf5\x0e\x95\x63\xba\xaa\xe0\x5e\xfd\xa9\x1e\x09\xca\x1e\x40\xd8\x92\x53\x55\x51\x12\x6a\xc9\x37\x92\x61\xb5\xfb\xb7\x68\x97\x2a\x31\x6c\x41\x26\x8d\xfb\x1b\x59\x57\xf6\xf6\x00\xdf\x67\x7c\xac\x89\x23\xfb\x11\x91\x77\xc9\xbf\x2d\x5f\xc4\x50\xeb\xda\x50\xe7\x2c\xde\x63\xac\x12\xc7\x7b\xc7\xe2\x0b\x3e\xac\xd1\x78\x83\x0f\xc5\x49\x03\xca\x5a\x18\x74\x16\x7b\xd4\x46\x51\xc4\x0f\x99\x90\xaa\x9a\xb2\x60\x5c\xc1\x36\xd8\xba\x26\xda\x11\x03\x34\x0e\x19\xe6\xc8\x5a\xf2\xd4\x9c\x51\x90\xe0\xa2\x5e\x08\x6a\x54\xfb\x5c\x62\x37\x8e\xb1\x7a\x1d\x5b\x92\xf8\x3c\x8f\x27\x68\x04\x38\xe0\xf1\x46\x14\x49\x86\x3a\xd8\x81\xf7\xc8\xeb\x30\xee\x20\xe8\x64\xb7\x6a\x57\xc8\xf0\x1a\xce\x42\x75\x03\x31\x31\x70\xb1\x4b\x15\xee\xc2\xb7\xe8\x4f\x0f\x89\xc6\x07\xb8\xe8\x0e\x05\xd9\xd7\x8a\xde\x13\x79\xb5\x6f\x08\x93\x63\x23\xaf\x6a\xc7\x8a\x46\x4e\x6b\xc5\xc9\x83\xd1\x77\x0e\x5d\xba\x27\x0c\x9b\x7b\xea\x88\x4f\x3f\x27\xdc\x2d\x8c\xcd\x61\x09\x83\x91\x5c\x38\xa7\x5f\xae\xa0\x90\x19\x87\x6a\x26\x85\x96\x2e\x5e\x77\x7a\x28\x47\x34\x04\x5a\x78\xc9\xee\x24\x4f\x55\x79\xa1\x58\xa6\x9a\x6a\xdf\xd5\xe7\x79\x26\x8f\xbe\x01\x00\xc9\x47\xd8\xc1\x81\xee\x1d\x6a\x8b\xdf\x59\x28\xbb\x14\x96\x65\xa0\xcc\x7c\x00\xba\x46\xf7\xc7\xde\x71\xbf\xe9\x3a\x32\x35\x8e\x80\x35\x8d\xce\x34\x51\xb9\x90\x45\x9f\x71\x04\x77\x5a\xe6\x42\xcb\xec\x99\x5d\x16\xeb\x8c\x7a\xc9\x9d\xbd\x9d\xe3\xb3\xaf\x8e\xc9\xf7\xe3\x2e\xe1\x5a\x68\x7b\xd7\x95\x15\xb7\x44\x67\xd5\x95\x9e\x5b\x8a\x18\xac\x7d\x4e\xc9\x73\xd4\x5e\x1d\x3d\x1f\x8e\x94\x62\x8a\xae\x33\xc9\x43\xf3\x7a\xa9\x7c\x67\xf4\x6a\xe5\x4d\x7e\x42\xac\x92\x28\x1f\x28\x95\xcf\xf1\x92\x56\xb1\xab\xb6\x5f\x2b\x1f\xe1\x65\xb1\x18\xcc\x40\xb9\xd8\xb4\x15\x8c\x16\xc3\x26\xaa\x3b\x83\xd9\x73\xcf\x7c\xf9\x97\xc2\x20\xd1\x7a\xdf\x1f\x27\x45\xf3\x1e\xe7\x44\x1b\x2a\x88\x97\xaf\x5f\xe2\x41\x03\xd6\x5b\xc8\x29\xd1\x5a\x04\x03\x35\x33\x25\xb3\xd8\x36\xd2\x8f\x2a\x66\xca\xbe\x31\x7b\xbd\x62\x2f\x9f\x88\x8d\x60\xbd\x27\xe2\xc0\xb3\xee\xac\x5c\xbd\x27\xe2\xde\x94\xf5\xd5\x71\x5a\xaf\xba\xc4\x1f\x38\x11\x1b\x3c\xaf\x3e\x11\x89\xa1\x1b\xcc\x4a\xd4\x26\xf0\xe8\x8f\xee\x98\xa7\x5f\xf6\x79\x02\x17\x8d\x69\x74\xfb\x3e\x3c\xb4\x60\xac\x7c\x67\x59\x5d\xc2\xc6\x01\x76\xf2\x5f\x90\x1b\xbf\x86\xe9\xc5\xbc\xe9\xbe\x96\xfd\x67\x01\xc2\x0a\x9f\x9d\xd2\x49\x82\x09\xcc\x15\x7d\x0a\x10\xbd\x4d\x1a\xbe\x03\xe5\xa4\xef\x64\x15\xc2\xe3\x52\xc6\x4b\x67\x9a\x65\x90\xb1\x58\x75\x14\xba\x46\xb9\x7b\x1d\x7f\xe6\x44\xd7\xa2\x50\x85\x8c\x45\x76\x83\x22\x41\xfd\x07\x45\xa7\x6f\x06\x5b\x27\x32\xca\xf7\x8b\xa4\x96\x11\x05\xcc\xb1\x09\x11\xc7\x68\x0c\x01\xa0\xdc\x48\x9f\x35\xd4\x08\x3e\x73\x7d\xc3\x85\xab\xb6\xd8\x7f\x69\xfb\x1f\x91\xad\xd1\xdf\x3a\xb8\xd8\xaf\xbf\x7d\x0b\xcf\x1a\xbe\x80\x8e\x2a\xdb\x45\x70\xd7\xd7\x56\x3b\x72\x23\xd9\xfe\x0f\x00\x00\xff\xff\x09\x00\x51\xc2\x47\x0e\x00\x00")

func nameServiceGeneratedTransport_grpcGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedTransport_grpcGotemplate,
		"NAME-service/generated/transport_grpc.gotemplate",
	)
}

func nameServiceGeneratedTransport_grpcGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedTransport_grpcGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/transport_grpc.gotemplate", size: 3655, mode: os.FileMode(436), modTime: time.Unix(1478128977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedTransport_httpGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2c\xcb\x31\x0e\xc2\x30\x0c\x05\xd0\x9d\x53\x78\xae\x44\x7c\x8d\x8e\x48\xcd\x05\xa2\xf2\x09\x88\x50\x47\xce\x87\xc5\xca\xdd\x61\x60\x7f\x2f\x42\x17\xd9\x00\xa9\x76\xa6\xbf\xc7\xd0\x8a\xa3\xda\xf3\x41\xbd\x93\x9d\x5e\x8e\xd1\xcd\xa9\xc4\xab\xb7\x42\xa4\x6a\x72\x33\x97\xdd\xae\x90\x45\xe7\x3c\x45\xec\xa5\x35\x49\x6b\xce\x97\x15\xad\xc3\xd3\x06\xff\xc0\xf3\xbf\x48\xfa\xa9\x6f\x00\x00\x00\xff\xff\xaf\xf5\x0f\x3e\x69\x00\x00\x00")

func nameServiceGeneratedTransport_httpGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedTransport_httpGotemplate,
		"NAME-service/generated/transport_http.gotemplate",
	)
}

func nameServiceGeneratedTransport_httpGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedTransport_httpGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/transport_http.gotemplate", size: 105, mode: os.FileMode(436), modTime: time.Unix(1477930115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceHandlersServerServer_handlerGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x5f\x6e\xd4\x30\x10\xc6\x9f\x3d\xa7\x18\xad\x56\x28\x41\x5b\xe7\xbd\x52\x5f\x40\x02\xf1\xd0\xaa\x02\x2e\xe0\x66\xa7\x59\x97\xac\x1d\x6c\x67\xbb\xc8\xf2\x99\x38\x04\x17\x63\xec\x64\x49\xc4\x1f\xd1\xa7\xc4\xf6\x6f\x3e\x7f\xf3\x79\x06\xd5\x7e\x51\x1d\xe1\x41\x99\x7d\x4f\x0e\x40\x1f\x07\xeb\x02\x56\x20\x36\x9d\xed\x95\xe9\xa4\x75\x5d\x73\x6e\x0c\x85\xa6\xb5\x26\xd0\x39\x6c\x00\xc4\xf0\x80\x9b\x18\xe5\xfd\x9b\x0f\x85\xbf\x57\xe1\x80\x57\x29\x6d\xa0\x06\x68\x1a\xbc\xa3\xe7\x4f\xe4\x4e\xba\x25\x74\x14\x46\x67\x3c\x2a\x34\xea\xc7\xf7\x13\xed\xd0\x07\x15\xa8\x27\xef\x91\x2f\xeb\xe9\x48\x86\x37\xb4\x35\x68\x1f\x71\xae\x92\xf0\x38\x9a\x76\x25\x53\xd5\x38\x3c\xc8\x18\xdf\xdb\x3b\x75\x24\x94\x17\x2e\xaf\x52\xca\x2b\x72\x18\x41\x4c\xb7\x61\xb6\x36\x75\xb6\x00\x8c\xc7\x04\x09\x20\x7c\x1b\xe8\x5f\x04\x9b\x73\x63\x1b\x18\x84\x18\x9f\x35\x77\xb5\x0d\x84\xd7\x37\x28\x31\x25\x10\x31\x3a\x8e\x84\x70\xab\xf3\x1e\x1f\xfd\x32\x72\x4b\xe1\x60\xf7\x3e\x43\x82\x03\x88\x71\xab\x67\x73\x4b\x97\x7e\x69\x4f\x88\xd2\x60\xe5\x33\xc9\x32\x7f\xf3\x52\xaf\x55\xaa\x36\x9c\x71\x7e\x00\xf9\x76\xfa\xee\x50\x1b\x7c\xbd\xce\x85\xf1\x8f\xf4\x75\x24\x1f\x3e\x73\x93\x73\x69\x8d\xd5\x9f\x90\x1f\xac\xf1\xb4\xa2\x76\x48\xce\x59\x57\x73\x88\x42\x9c\x94\xe3\x87\xf3\x03\xfe\xbf\x2e\xe3\x05\xbd\x79\x01\x5c\xc4\x97\x14\x9f\x4a\x8a\xbf\x91\xb7\x3c\x19\x1c\x85\x7c\xa7\xa9\xdf\xfb\x3c\x55\xa5\x68\x4a\xf5\xa2\xfe\x34\x0b\x5e\xe3\xac\x48\x66\x7f\x41\x67\x4b\x65\x10\x5e\x65\x6b\x3b\x34\xba\x87\x72\x50\x40\xc6\x62\xbc\xc2\xe9\xef\x67\x00\x00\x00\xff\xff\x14\x16\x8e\x3b\x02\x03\x00\x00")

func nameServiceHandlersServerServer_handlerGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceHandlersServerServer_handlerGotemplate,
		"NAME-service/handlers/server/server_handler.gotemplate",
	)
}

func nameServiceHandlersServerServer_handlerGotemplate() (*asset, error) {
	bytes, err := nameServiceHandlersServerServer_handlerGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/handlers/server/server_handler.gotemplate", size: 770, mode: os.FileMode(436), modTime: time.Unix(1478128977, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	"NAME-service/NAME-cli-client/client_main.gotemplate": nameServiceNameCliClientClient_mainGotemplate,
	"NAME-service/NAME-server/server_main.gotemplate": nameServiceNameServerServer_mainGotemplate,
	"NAME-service/generated/cli/handlers/client_handler.gotemplate": nameServiceGeneratedCliHandlersClient_handlerGotemplate,
	"NAME-service/generated/client/grpc/client.gotemplate": nameServiceGeneratedClientGrpcClientGotemplate,
	"NAME-service/generated/client/http/client.gotemplate": nameServiceGeneratedClientHttpClientGotemplate,
	"NAME-service/generated/endpoints.gotemplate": nameServiceGeneratedEndpointsGotemplate,
	"NAME-service/generated/transport_grpc.gotemplate": nameServiceGeneratedTransport_grpcGotemplate,
	"NAME-service/generated/transport_http.gotemplate": nameServiceGeneratedTransport_httpGotemplate,
	"NAME-service/handlers/server/server_handler.gotemplate": nameServiceHandlersServerServer_handlerGotemplate,
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
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"NAME-service": &bintree{nil, map[string]*bintree{
		"NAME-cli-client": &bintree{nil, map[string]*bintree{
			"client_main.gotemplate": &bintree{nameServiceNameCliClientClient_mainGotemplate, map[string]*bintree{}},
		}},
		"NAME-server": &bintree{nil, map[string]*bintree{
			"server_main.gotemplate": &bintree{nameServiceNameServerServer_mainGotemplate, map[string]*bintree{}},
		}},
		"generated": &bintree{nil, map[string]*bintree{
			"cli": &bintree{nil, map[string]*bintree{
				"handlers": &bintree{nil, map[string]*bintree{
					"client_handler.gotemplate": &bintree{nameServiceGeneratedCliHandlersClient_handlerGotemplate, map[string]*bintree{}},
				}},
			}},
			"client": &bintree{nil, map[string]*bintree{
				"grpc": &bintree{nil, map[string]*bintree{
					"client.gotemplate": &bintree{nameServiceGeneratedClientGrpcClientGotemplate, map[string]*bintree{}},
				}},
				"http": &bintree{nil, map[string]*bintree{
					"client.gotemplate": &bintree{nameServiceGeneratedClientHttpClientGotemplate, map[string]*bintree{}},
				}},
			}},
			"endpoints.gotemplate": &bintree{nameServiceGeneratedEndpointsGotemplate, map[string]*bintree{}},
			"transport_grpc.gotemplate": &bintree{nameServiceGeneratedTransport_grpcGotemplate, map[string]*bintree{}},
			"transport_http.gotemplate": &bintree{nameServiceGeneratedTransport_httpGotemplate, map[string]*bintree{}},
		}},
		"handlers": &bintree{nil, map[string]*bintree{
			"server": &bintree{nil, map[string]*bintree{
				"server_handler.gotemplate": &bintree{nameServiceHandlersServerServer_handlerGotemplate, map[string]*bintree{}},
			}},
		}},
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

