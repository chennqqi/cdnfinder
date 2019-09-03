// Code generated by go-bindata.
// sources:
// assets/cnamechain.json
// assets/resourcefinder.js
// DO NOT EDIT!

package resource

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _assetsCnamechainJson = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x96\xc1\x72\xe3\x36\x0c\x86\xef\x79\x0a\x4d\x0e\x3d\xad\xd5\xd9\xc9\x21\xe7\xc4\xdb\x36\x99\xcd\x66\xda\x26\x9d\x9d\x76\x67\x0f\x10\x05\x49\x1c\x53\xa4\x86\xa4\xac\x28\x4f\x5f\x92\x96\x48\x58\xb6\x9c\x9b\x81\xff\x23\x09\x80\x00\xad\x1f\x57\x59\xf6\xe3\x3a\x67\x82\xa3\xb4\x26\xb7\xbd\x2e\x54\x31\x5a\x34\xb9\x44\x7b\xfd\x29\xbb\x7e\xf5\x9e\x7b\xef\xb9\xfe\xf9\xe9\x00\x27\x68\xc3\x4a\x99\x33\xd5\xae\x80\x50\xbd\x79\x60\xda\x89\x58\x11\xd8\x41\x0b\x7c\x06\xee\x82\xb5\x10\xb1\xac\x71\x1d\x28\xa5\xb9\xb8\xda\x22\x6b\xa4\x12\xaa\xe6\x2e\xa3\x29\xd0\x05\x58\x1b\x51\xe4\xb6\x60\xc0\x1a\x8c\x88\xe0\xad\x83\x22\xc3\x84\xea\xcb\x4a\x2b\x69\xe3\x69\x2d\xbc\x2b\x99\x6d\xa3\x90\x0e\x96\x20\x77\x38\x15\x26\x2f\x74\xa0\x83\x2f\x21\xd6\x6e\x4a\x03\x71\xaf\xd7\x5f\x5e\x93\xf6\xce\x95\x24\x55\xbb\xf3\x76\x54\x0b\x14\x7d\x0d\xa4\xea\xf7\xc1\xb1\xfd\xf2\x9c\x10\xd1\x63\x03\xd6\x2d\x1f\x94\xde\x45\xce\x79\xb3\x07\xb0\xd9\xf3\xc1\x1f\x71\x33\x1a\x8b\x2d\x39\xf0\x37\x57\xf0\x2d\x98\x94\x50\xa8\x4c\x25\xc6\x19\xd8\x4e\x76\x02\x4a\x79\x7b\x1b\xd5\x2f\xcf\xb7\xb7\x0b\x49\xe9\xfa\x54\xea\x40\xda\x06\x35\xc9\xc5\xe9\x87\xe0\x0c\x5d\x5f\x33\xb2\xf5\x52\xaf\x19\x89\x7c\x55\x67\x72\x7d\x7b\x5e\x8d\x39\x57\x93\xee\x8c\xa4\xb1\xd0\x19\x73\x6c\x0d\x97\x10\x32\x5f\x02\xf3\xe1\x67\x80\x9b\x43\x53\x5d\x20\xbc\xeb\x43\xe8\x86\xe6\x78\xaa\x8b\x37\x3f\x05\x34\x4e\x97\xe7\xd6\x8d\x33\xea\x74\xcb\x1d\x62\xc9\x08\xf7\xd7\x3f\x77\xcf\xaf\x8f\x4f\xbf\xae\xf0\xed\x10\x5a\x9e\x5c\xce\x07\x0b\x0e\x13\x22\x40\xc7\x19\xda\x46\xcf\x39\x68\x4e\xe7\x14\xf2\x03\xcf\x5c\xff\x5d\x6a\x49\x37\xf7\xf9\x59\x61\x00\xb6\x22\x74\x2b\x42\xe5\x8c\xd4\xdb\xbf\x07\x6b\x21\xa6\x4b\x5e\xc8\xb5\x52\xb5\xc0\xb0\xef\x1f\xe1\xe7\xac\x1c\x04\x33\xca\x92\x33\xb0\x6e\x80\xcf\x30\xa3\xea\x6d\x5f\x9c\x5b\x3d\xed\xdb\x1b\x37\x1d\xee\x65\x71\x85\x9e\x6b\xba\xe0\x44\x5e\xaa\xbe\x10\xe8\x5e\x6f\xb6\x9b\x63\x3c\x66\x4a\x37\x02\xe1\x12\x83\xb4\x61\x8a\x54\xba\xe1\x05\x6a\xc9\xe9\x73\xf2\x30\xb9\x12\x33\x90\x7b\x78\xe0\x75\x33\x70\x59\xa6\x01\xe2\x92\x41\x47\x1e\xe1\x47\x6f\x9b\x5e\x00\x21\x0c\x93\xd0\x62\x22\x8c\x05\x6d\xfd\x93\xcc\x28\x24\x87\x0f\x08\xd7\x6d\x12\x3a\x12\xcd\xe3\xe4\x8a\xcc\xee\x8d\x64\xf2\x15\x47\xfa\x2a\x0a\x43\x13\x79\x42\x30\xf8\x1d\x8b\x8c\x22\x95\x52\xb6\xd3\x3c\xbd\xf1\x4f\xb8\x47\x71\x93\xb6\x10\x72\x28\xa3\xc6\x5b\x14\xae\x1c\x96\xc8\xa4\x0e\xa7\xb2\x13\x4a\x09\xf4\x0f\xf3\x1b\xbc\xd1\xe3\x27\xc0\x18\x71\x11\x58\x11\x5d\xc9\xd8\xae\x03\xdb\x90\x19\x7f\xf1\xbe\x3f\x9d\x2f\x8d\xb5\xa4\xe7\x63\xc9\x41\xaa\xfd\xd1\x5d\xb9\x25\x53\xcb\x4d\x99\x7c\xe3\x5a\x2b\x9d\x3d\xb6\x50\xa7\xce\x69\x83\x73\xc3\xbd\xf3\x22\xe8\xba\x61\xd3\x72\x7e\x91\xd1\x3e\xa8\xcf\x73\x54\x7f\xbb\xbf\x16\x64\x16\xcb\xec\xe4\xb5\x36\xbc\xed\x04\x92\x6b\x7c\x09\x8e\xa3\x4b\x34\x03\xaf\x2c\xdd\xef\xc5\x3b\x4e\x08\x37\x5a\x7b\x3c\x42\x5e\xbc\x67\xf1\x51\x00\xaa\x00\x15\xbf\x6f\x82\x45\xff\x3a\xf2\x82\xdb\x5a\xc3\x9e\xdb\x31\x41\x16\x32\xf7\xbb\xed\xe5\x34\xf9\x47\x7f\x36\xb9\x45\x81\x95\xf2\x5a\x5c\x11\x3d\x11\xdc\xab\xbc\x35\xc8\x64\x6c\xb6\xef\x6e\xe6\xd4\x60\xb2\xbb\xf7\x9e\x8c\x2f\x8c\x9f\xf3\x22\x1f\xa1\x51\x31\xc6\x7f\xbd\x11\x81\x91\xb7\x75\x7e\xea\x7e\x47\x49\x3f\xaa\xfe\x3b\x98\xe9\xfb\x61\x43\x2a\x7c\xdf\x4b\x79\x34\x47\x3b\x33\xf6\xb4\x89\xbe\x72\x59\x1b\x55\xb9\x3e\xbf\xfa\x79\xf5\x7f\x00\x00\x00\xff\xff\x2a\x19\x5e\xd9\x4d\x0a\x00\x00"

func assetsCnamechainJsonBytes() ([]byte, error) {
	return bindataRead(
		_assetsCnamechainJson,
		"assets/cnamechain.json",
	)
}

func assetsCnamechainJson() (*asset, error) {
	bytes, err := assetsCnamechainJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/cnamechain.json", size: 2637, mode: os.FileMode(436), modTime: time.Unix(1549250601, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsResourcefinderJs = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x57\x5f\x73\xdb\x36\x0c\x7f\xf7\xa7\x40\xbc\xbb\x5a\xba\xd8\x52\xb2\xde\x76\x5b\xbc\xac\xb7\xdb\xba\x6b\x77\x5d\x7b\x97\xec\x2d\x4b\xef\x68\x09\xb6\x99\x4a\xa2\x46\x52\x71\xbd\xcc\xdf\x7d\x00\x49\xfd\x73\x5b\x57\x0f\x89\x04\x82\x00\xf8\x03\xf0\x03\xfd\x28\x34\xd8\x39\x88\x3c\xd7\x68\xcc\x72\xc2\xdf\xf4\xa6\x1a\x9d\xa1\x81\x6b\x78\x3a\x78\x99\xd9\x1b\x8b\x25\x09\x34\xfe\xd3\x48\x8d\xd1\xcc\x4b\x66\xf1\x72\xe2\x14\x36\x68\x5f\x29\x63\x2b\x51\x22\x69\xad\x9b\x2a\xb3\x52\x55\x91\xb1\x3a\x86\xa7\x09\x40\x9a\x82\xb1\xaa\x28\xb0\x82\xb5\x56\x25\x6c\xad\xad\xaf\xd2\x74\x85\x42\xe7\x26\xd3\xc2\x66\x5b\xd4\x26\xc9\x54\x99\x3e\x90\xf7\x4a\x14\x69\x63\x64\xb5\x59\x3c\x88\x47\x41\x0a\xb2\xb6\x0b\xab\x16\xe4\x66\x61\xb7\xb8\xd8\x06\x5f\x0b\xb5\x5e\x88\x45\xa3\x0b\x72\x61\xf5\x9e\x3d\x01\xf8\x43\x50\x18\x15\xee\xe0\x06\x37\x2f\x3f\xd6\xd1\xec\x7d\xf4\xe2\x6a\xfd\xdf\xd6\xc6\xb6\xa6\x37\x13\xbf\xf8\x9b\xdc\x47\x77\xef\xd3\xfb\xf3\x78\x36\x87\x99\x74\x67\xe1\xed\x1a\x2d\x05\x40\xe1\xea\xa4\xe4\xb8\x22\x8d\xf1\xdd\xe5\x7d\x62\xd5\xad\xd5\x14\x52\xe4\xf4\x0e\x90\xf1\x22\x44\xa8\xc3\x11\xf9\x90\xb5\x56\x2b\xb1\x2a\xf6\x90\x0b\x2b\xa0\xd1\x12\x76\x5b\x49\x5a\x3b\x84\x5c\x55\x96\xf6\x50\x60\x62\x65\x83\x7a\xa6\x2a\xa3\x0a\x4c\x0a\xb5\x71\x50\x0d\xfd\x57\x4d\x51\x38\x47\x93\xc3\x64\x92\xa6\x6f\x15\x1d\x07\x73\x58\x2b\x0d\x5b\x14\x39\xa1\xe5\xde\x51\x90\x79\x59\xe5\xf2\x51\xe6\x8d\x28\x40\xad\x1e\x30\xb3\xb4\xe1\x56\x96\x75\x21\xd7\x7b\x50\x8d\xad\x1b\x96\xfc\xe9\x02\xf6\x9f\xbc\x97\x4e\x07\x6a\xed\x52\x91\xbc\xf2\x26\x43\x01\xe4\x4d\x86\xad\x93\x41\x36\x83\x28\xe6\xd3\xb2\x22\xe1\xdb\x6b\x3d\x1d\x82\xf4\x03\xee\xe7\xf4\x52\x34\x38\x07\xb9\x0c\xc2\x2c\xaf\x38\x21\xe1\x4c\xe3\x93\xff\x71\xfb\xee\x6d\x62\x1c\xb6\x14\x6f\xe7\x85\xd1\xa0\x28\x23\x79\x7d\xb1\x94\x3f\x05\x69\x42\x15\xb4\xb1\xdb\xa5\x3c\x3f\x8f\x3d\xe8\xe4\x8e\x2c\x87\xe5\x3b\x79\x9f\x70\x5d\x2c\x43\x21\x50\x10\xe3\x45\x27\x5a\x12\x18\xe0\x9f\x61\x1c\x7d\xe0\x3e\x11\xfd\xf1\xee\x68\xe9\x9e\x0c\xdd\xb9\xd5\x7b\x9f\x96\x3e\x53\x9d\xde\x92\x73\xc5\xc7\x2d\xc5\x07\xd4\x58\x2b\x6d\x87\xf8\xc9\x8a\x90\xef\xd0\x73\xee\xe8\x8f\x99\xc3\x4a\x18\xac\xc5\x06\x73\x55\x0a\x59\xcd\x43\x8e\x1c\x7c\x3e\xd2\x61\x98\x53\x32\x4e\x50\x81\xb7\x3f\xe5\x58\xd9\x0a\x39\x7a\xe7\x72\x9f\xf0\x97\xf7\x95\x74\xbd\x1c\x7f\x15\x76\xde\xd5\x62\x0e\x01\x74\x96\x7d\x09\x71\x5e\x23\x44\x3d\xd2\x47\xde\x1c\x5e\x49\x5f\x1b\xa3\x8a\x3a\x0e\x6d\xa4\x1c\x9f\xb0\x27\x0d\xe3\x44\xe6\xfc\x6a\x0b\x1a\x73\x01\x5c\xbb\x80\x4e\x05\xd3\xb3\x53\xa7\x98\xa6\x39\x16\x68\xf1\x64\xf8\x6d\xb2\xd3\xb4\x52\x39\x3e\x18\x10\x75\x4d\x07\x12\xb9\xf1\x4c\x16\x20\x9d\x8c\x93\x74\x04\xae\xcf\x7c\xec\xea\x63\xc2\x35\xf1\x9b\xb0\x98\x54\x6a\xc7\x64\x12\xd8\x97\xa4\x9e\x54\x13\xa1\x37\x86\x18\x27\x50\x6b\xa1\x44\xce\xe7\x1c\x56\x12\x31\x5e\x57\x47\x61\xad\x63\xe6\x1d\xae\x58\x34\x8b\x93\x8c\xc2\xb4\x18\xc5\x2e\x7a\x6e\x73\x43\x9c\xb7\x91\x76\xdb\xac\x1c\xd3\x0a\x2d\xf7\x22\xad\xb7\xa2\xb2\xaa\x7c\x30\xa9\x34\xa6\x41\x93\x5e\x5e\x3c\xff\xe1\xc7\x6f\xdc\x07\x69\x95\x58\xd9\x05\x89\xbe\xff\xee\xe2\xf2\xdb\xe7\x64\x8a\x8d\x27\xaa\x7a\x2b\x1e\xe5\x46\x70\x34\x37\xe4\x19\x29\xf0\xfc\x28\xc2\x39\xd8\x7d\x4d\x24\xb0\x93\x45\x11\xb4\xe9\x8b\x2b\xbc\x25\x4c\xca\xd5\x1a\x22\x96\xc0\xb3\x67\x44\x97\xc5\xd9\x75\x00\xa3\xd7\xe0\xa7\x47\x88\x74\x96\x83\x85\x71\x45\x4f\xa9\xce\x08\x82\x8c\x79\xb6\xd9\x6c\x7d\x6f\xb4\x8f\x0b\x3b\x2b\x94\xf1\x80\xb4\x8f\x41\xfb\x97\x2c\x91\x3a\x2e\x9a\xb5\x48\x47\x6d\x10\xb3\xf9\x65\xbc\x04\x26\x5f\x2a\x12\x9a\x3b\x03\x6d\x22\x15\x8d\xc1\x0e\x97\x07\x0d\xca\x1e\x9b\x9b\x50\x49\x37\x98\xa1\x7c\x1c\x03\xa3\x3d\x5a\x71\x3f\xab\x1c\x54\x46\xfe\x4b\xe0\xb4\x55\x3a\x6f\x39\x2b\x50\xa8\x57\x1c\x36\x94\x33\x32\xac\xd1\x8e\xd2\x4e\xd5\x61\xeb\x3c\xb4\x19\x79\x1e\x18\xeb\xb0\xe5\xa4\x9c\x45\x1c\x51\x1c\xb7\x69\xe0\x2f\xd2\x6d\x2d\x24\x2b\x95\xef\x6f\x59\xf6\x02\x8e\x45\x57\x70\x11\xec\x1f\x3e\x33\xe7\xb8\x74\xfd\xea\xa0\x25\x07\xd7\x87\x81\x02\xc7\x11\xb5\x5a\x31\x97\x88\x0b\x0a\x7e\x26\x07\xf1\xb0\x82\xce\xa2\xbe\x77\x5b\xfd\xfb\x5e\x05\xe0\x33\xcb\xe1\x76\x73\x42\x83\x7a\xa4\xa9\xb8\x59\x2f\x4e\xab\xad\xf6\xd6\x5d\x96\x3a\xb5\xc3\xe4\x6b\x36\xcf\xaf\xe1\xb2\xd5\x4e\xfb\x0e\x84\x42\xa2\x39\x03\xa3\xf8\xc2\x60\x10\x39\x95\x96\xfb\xcf\x73\x70\x28\x00\x90\x44\x42\x8f\x42\x16\x74\xd3\x68\x4b\x70\x40\xda\x5f\x9c\x94\x2d\x5a\x47\xb3\x92\x2e\x36\x6f\xd4\x0e\xf5\xaf\x82\x5b\x83\x99\x74\x3a\x76\x3b\x8d\x87\xbd\x18\x0a\xa1\x16\xda\xe0\xeb\xca\x46\xc7\xc3\x35\x1e\xf6\xe7\x8a\x18\xe8\x43\x2f\x38\x74\x00\x9d\x40\xc8\xc3\x49\x08\xb1\xa7\x1e\x24\x23\x1e\x7d\x0f\x16\x82\xf8\x9e\x36\xd6\x54\x54\xd8\xf5\x44\x4d\xb8\xb0\x8d\x13\x86\xfb\xf6\x19\xb4\x4d\xe8\xde\x89\xeb\xa0\x5b\x49\x27\xb6\x74\x7d\x93\x9b\x4a\xd1\x4d\x8d\x32\x42\xd7\xbb\xbe\xaf\x5f\x6a\x4d\x38\x0f\x9a\xb9\x34\x1b\x62\x39\x2d\x32\xf4\x84\x75\xe8\x49\xa0\xc6\xaa\xe5\x91\x79\xb7\x81\x0a\xd8\x0a\xdb\x74\xf4\xc6\x5d\xed\x87\x7d\x5f\xf1\x5e\x03\xce\x28\x11\x33\xd3\x64\x74\x06\x33\xeb\xe9\x70\xd8\x4a\xb3\xa7\x29\x72\x44\xd3\x2b\x98\xfe\xfe\xcb\xeb\x37\xd3\x43\x7b\x87\x3d\x00\x16\x04\x4e\xbb\x69\x3c\x71\x60\x01\x76\xcc\x18\x9f\x18\x7e\x43\x4c\xc8\x77\x0c\x4b\x64\x07\x33\x38\x27\x0b\xe7\xf4\xbf\x34\x98\xcd\xba\x04\x87\x9b\xe4\xb0\x8f\xbc\xe8\x68\x34\x7b\x40\x90\xab\x83\xc7\x51\x8f\xc5\x98\xe4\xc3\x6d\x2a\x57\x59\xc3\x33\x87\x02\xc9\xdc\x78\xe9\x46\x77\xd7\x5e\xf1\x72\x32\xf6\x37\xfc\xc9\xd2\xbd\xb7\xea\xfd\x5d\x2c\xf2\xea\x71\x5f\x54\x27\xa8\x32\xe8\x8e\x98\x2c\x34\x6a\x82\x1f\xa5\x0d\x3f\x03\xfc\x4c\xff\x64\x72\x2c\x27\xff\x07\x00\x00\xff\xff\x0e\x85\xee\xee\x5b\x0d\x00\x00"

func assetsResourcefinderJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsResourcefinderJs,
		"assets/resourcefinder.js",
	)
}

func assetsResourcefinderJs() (*asset, error) {
	bytes, err := assetsResourcefinderJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/resourcefinder.js", size: 3419, mode: os.FileMode(436), modTime: time.Unix(1521745412, 0)}
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
	"assets/cnamechain.json": assetsCnamechainJson,
	"assets/resourcefinder.js": assetsResourcefinderJs,
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
	"assets": &bintree{nil, map[string]*bintree{
		"cnamechain.json": &bintree{assetsCnamechainJson, map[string]*bintree{}},
		"resourcefinder.js": &bintree{assetsResourcefinderJs, map[string]*bintree{}},
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

