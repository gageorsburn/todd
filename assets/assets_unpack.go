// Code generated by go-bindata.
// sources:
// agent/testing/bashtestlets/http
// agent/testing/bashtestlets/iperf
// agent/facts/collectors/get_addresses
// agent/facts/collectors/get_hostname
// DO NOT EDIT!

package assets

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

var _testingBashtestletsHttp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x97\xdf\x73\xd3\x38\x10\xc7\xdf\xfd\x57\xec\xb9\x39\x4c\x99\x36\x26\x7d\xb8\x87\x32\x65\xae\xa4\xe1\x67\x28\x25\x29\x73\x77\x43\x98\x8c\x63\x2b\x89\x0f\xc7\xf2\x59\x72\x7a\xbd\xd2\xff\xfd\x56\x92\xad\xac\x9c\x12\x32\xf0\x82\xb5\xfe\xee\xc7\xd2\x6a\xf7\x1b\x38\xf8\x25\xac\x44\x19\xce\xd2\x3c\x64\xf9\x1a\x66\x91\x58\x7a\xde\x01\xf4\x79\x71\x5b\xa6\x8b\xa5\x84\x93\xa7\xbd\xdf\xe0\x7d\x24\x25\x7c\x10\x37\x51\x26\xbb\xf0\x49\x30\xe0\x25\xac\x78\x92\xce\xd3\x38\x92\x29\xcf\x81\xcf\x41\x2e\x53\x81\x99\x82\x57\x65\xcc\x20\xe6\x09\x83\x54\xc0\x82\xaf\x59\x99\xb3\x04\x66\xb7\xa8\x60\x90\xa5\x31\xcb\x11\x50\x94\x7c\x9d\x26\x18\x5f\xb2\x92\x9d\x62\xde\x52\xca\x42\x9c\x86\xe1\x22\x95\xcb\x6a\xd6\x8d\xf9\x2a\x94\x3c\x49\x50\xf7\x37\x8b\xa5\x7e\x0e\x67\x19\x9f\x85\xab\x48\x48\x56\x86\xc3\x37\xfd\xc1\xe5\x78\xe0\x79\xf1\x92\xc5\x5f\x1f\x1f\xc2\x9d\x07\xf8\x87\xc5\x4b\x0e\xfe\xa8\xca\xf3\x34\x5f\x40\x9a\x83\x7e\xad\x36\xcb\x7c\x2d\x40\xf0\x2a\xca\x13\x38\x5e\x43\x5c\x95\x19\x3c\x0f\x13\xb6\x0e\xf3\x2a\xcb\xe0\xe4\xf9\xa3\x1e\x7c\xfb\x06\x77\x86\xf2\xfc\xd1\x09\xf8\x5a\x83\xe7\xc8\xb9\x44\x9a\x90\x51\x96\xb1\xa4\x0b\x83\x7f\x53\x89\x1f\xe8\xfa\xcf\x80\xe1\x23\xf4\x9e\xc1\x3d\xf9\x7e\xdf\x7e\x14\xae\xce\xc7\xe3\xc1\x85\xf9\xb6\x96\x3e\xf5\xee\x3d\x2f\x9d\xc3\xe7\xcf\xd0\xe9\xc1\xd9\x19\x7e\x43\xa9\x7d\xf8\xf2\xe5\x19\xaa\xb0\x48\xb9\xd9\xa8\x8a\x7a\xf3\x54\x5d\xc7\x35\x56\x6e\x9e\x96\x42\x42\x54\x2e\xaa\x15\xcb\xf1\x61\x8e\x55\xd0\x25\x95\x4c\xc8\x8c\x49\xc8\xa3\x15\x83\x9b\x14\x0f\x72\x3e\xfc\xe3\xfc\xaf\x31\xcc\x98\x79\x8f\x39\x0c\xef\x6d\xcc\x8f\xd4\x17\xf1\x30\x51\x76\x13\xdd\x0a\xf2\xf2\x08\x54\x49\xf0\x6c\x96\x2f\x60\x5e\xf2\x15\x74\x4e\xf4\x9b\xaa\x00\xb1\xe4\x55\x96\x28\x66\x11\x09\x81\x17\x97\xe6\x92\x6b\x44\x54\x14\x4d\x55\x3d\x7c\x9e\x22\x42\x9c\xf9\x9d\xdf\xfd\xcd\xaa\x73\xd7\x3c\x1e\x74\x9e\xc2\x3d\x1c\xc0\x88\xad\xb0\x33\xe0\x9a\xec\xfd\x61\x79\xcf\x91\xeb\xdd\x62\x41\x22\xa9\xfb\x0d\x0a\x8e\xdb\x38\x82\x46\x4e\x76\xa9\x0e\x23\x97\x91\x0c\x04\x64\x6c\x2e\x3d\x8f\x57\xb2\xa8\xe4\x59\xe7\xb1\xbe\xd2\x63\x01\xc7\x1c\x36\x77\x7f\x7c\x03\x13\x5d\xf6\x60\x62\xca\x3f\xc9\x5f\x5f\x5f\x5f\xe1\x20\x24\xec\x14\x7e\xbd\x53\xed\x39\x55\x4d\x7d\x0f\xd3\xc1\xe5\xc5\xb9\x55\x0d\x39\xff\x8a\xe5\x91\xe9\x4a\xeb\xd4\xdf\x53\x75\x9a\x4c\xc7\x8d\xfa\x85\x55\xf7\x79\x9e\x63\x3b\xbb\xf2\xd8\x04\x8d\xb6\x6f\xb5\x57\x25\xfb\x73\xae\xee\x98\x6a\x8b\x92\xc9\x32\xca\x05\xbe\x30\xfa\x0b\xab\x1f\xe3\x5d\xca\xed\x0c\xa1\xc2\x6e\xce\xc0\xe6\x5c\x73\xec\x68\x57\x2f\x55\xc8\xe8\x5e\xd2\x7d\x4b\xd5\x75\xf2\xb6\xd0\xca\xd8\xac\xa7\x6a\x6d\xb4\xaf\xda\x75\xd3\x67\x22\xa5\x23\x67\x7c\x6d\xb5\x97\xd5\x6a\x86\x1b\x46\xfb\xa8\x13\xd0\x4c\x84\x4a\xca\xab\x55\x93\x23\x4c\xd2\x9b\x07\x92\x46\x2c\x49\xcb\x56\x56\x59\xc7\xea\xb4\xb7\x9b\xfa\xa4\xff\x31\x95\x94\xf0\x9b\x3c\xe3\x51\xa2\x12\x04\xc6\xa6\x4d\xc0\x24\xbc\xdb\x4a\x58\xb2\x28\x61\xa5\xb0\x7a\xb3\x36\xea\xe1\x96\xba\x64\xff\x54\xd8\xd5\x56\x5d\xaf\x8d\xfc\xfd\x96\xbc\x2a\x9c\xbd\x98\xa5\x11\x5f\x5a\xf1\xf9\x7a\x01\x17\xf5\x2e\x61\x5c\x30\x66\x12\xd4\x43\x6b\xf7\x1f\x9c\x9c\x4f\xc5\x43\x19\xf4\x1b\x57\x56\xdf\xd4\xd2\xed\x86\xa6\x9a\x46\xfd\x71\xd3\x38\xf8\x12\xd0\x00\x74\x77\x41\xd3\x5e\xbb\x9a\x6e\xb4\x19\x19\x74\x70\xf8\x34\x1a\xc2\x4b\x26\xd1\xe5\xf4\xce\x70\x26\xa7\x6c\x3e\x57\x77\xb9\xae\x3b\x6a\x6c\x13\x02\x9c\xcd\x4e\xef\xd0\xf3\xec\x14\xe2\x18\x6b\xa7\xed\x98\xa9\x86\x6f\xa0\x2c\xe9\x38\x87\x40\x84\xdd\x27\x64\x72\x27\x8f\xbb\x4f\x26\x87\x66\x66\xbb\x4f\xc2\x49\x2f\x2c\x82\x43\xaf\x35\xa5\xbb\x69\xce\x84\x13\xde\x8b\x16\xaf\x6e\xd7\xdd\x30\xd7\x00\x08\xad\xdf\xa2\x91\x41\xdf\x4d\x74\x6d\x82\x10\x2f\x5a\x44\xe7\x4e\x76\x33\xdb\x56\x42\xa8\x83\x16\x55\xdb\xc5\x6e\x1a\x35\x19\x42\x7a\xb9\x21\x51\x3b\xf9\x61\xfd\x88\x11\x11\xda\xab\x0d\x8d\x1a\xce\x5e\x8d\x52\x5b\x15\xa1\xbd\xde\xd0\xa8\x13\xed\xa6\x7d\xc7\xcc\x08\xf6\x8d\x8b\xb5\x56\xb5\x2f\xd7\xf5\x3b\x02\x7e\xbb\x01\x3b\x96\xf6\x83\x6b\xde\x72\x44\x82\x7c\xd7\x42\x1a\xd7\xdb\x0f\x68\x1d\x93\xf0\x86\x2d\x5e\xed\x8b\xfb\x01\xad\xa9\x12\xe0\xfb\x16\xd0\xd8\xda\x7e\xbc\xc6\x75\x09\xee\x92\xe0\x1c\x63\xdd\x4d\x7c\xc8\x9a\x09\xf5\x43\x9b\xba\xcf\x2e\xb7\xad\x9b\x10\xaf\x5a\x03\xd8\x34\xd1\x6e\x64\xcb\xdd\x09\xef\xe3\xcf\xdb\xc4\xf7\x7e\x04\x08\x7d\xb4\xa1\x3b\x16\xff\x03\xcb\xdd\xfa\x85\x20\xc8\xf1\x06\xe9\xd5\xff\xf6\x9d\x26\x91\x8c\xce\x02\xf3\x5f\x00\xdf\xfe\x4a\xf8\xa7\x7e\xd0\xb1\xab\xc0\x3f\x72\xde\xeb\x11\xa5\x12\x1d\xb0\xaa\xd6\x6f\x84\x7f\x1a\x74\x5a\xa1\x80\x2a\x2d\xaf\x96\x35\x38\xaa\x21\x9e\x6e\x75\x24\xe6\x68\x9d\x6b\xb0\x6a\x27\xea\xe8\xb5\x0f\x5b\x9d\x5e\x35\xef\xa9\xbb\xea\xf3\xd2\x80\x3d\x2f\xf5\x39\xc5\xa1\xeb\x80\x68\xac\x69\x35\x22\x1b\x68\x54\xc4\x2e\x94\x86\x2c\x1d\x45\x3d\xd3\x56\x52\xaf\x1d\x4d\x33\x83\x56\xd4\x04\x1c\x95\x99\x29\xab\x31\x4b\xab\x70\x86\x59\x8b\x9c\x88\xab\x23\x28\xb2\x76\x2a\xdd\x9c\xd7\x16\xbb\x09\xfc\xec\xfd\x39\x83\xa1\x2f\xc8\x89\x04\xbe\x77\x1f\x78\x66\x5a\x68\xbf\x7b\xff\x07\x00\x00\xff\xff\xb8\x97\x93\xa5\xbc\x0f\x00\x00")

func testingBashtestletsHttpBytes() ([]byte, error) {
	return bindataRead(
		_testingBashtestletsHttp,
		"testing/bashtestlets/http",
	)
}

func testingBashtestletsHttp() (*asset, error) {
	bytes, err := testingBashtestletsHttpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testing/bashtestlets/http", size: 4028, mode: os.FileMode(493), modTime: time.Unix(1478845223, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testingBashtestletsIperf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x54\x4d\x6f\xdb\x38\x10\xbd\xeb\x57\xcc\x2a\x46\x95\x04\x8e\x15\xe7\xb0\x87\x04\x0e\x36\x9b\xcd\x21\x40\xdb\x0d\xea\x2c\x8a\x45\x53\x18\x94\x34\xb2\xd8\x48\xa4\x40\x8e\xac\xf5\x3a\xfe\xef\x1d\x92\xfe\x6a\xe1\x4b\x93\x0b\x4d\x3e\x3e\xbe\xf7\x66\x46\x27\xbf\xa5\x9d\x35\x69\x26\x55\x8a\x6a\x01\x99\xb0\x55\x14\x9d\xc0\xbd\x6e\x97\x46\xce\x2b\x82\xab\xcb\xf1\xef\xf0\x41\x10\xc1\xdf\xb6\x17\x35\x8d\xe0\x1f\x8b\xa0\x0d\x34\xba\x90\xa5\xcc\x05\x49\xad\x40\x97\x40\x95\xb4\x7c\xd3\xea\xce\xe4\x08\xb9\x2e\x10\xa4\x85\xb9\x5e\xa0\x51\x58\x40\xb6\x64\x04\x42\x2d\x73\x54\x4c\xd0\x1a\xbd\x90\x05\xef\x57\x68\xf0\x9a\xef\x55\x44\xad\xbd\x4e\xd3\xb9\xa4\xaa\xcb\x46\xb9\x6e\x52\xd2\x45\xc1\xb8\x6f\x98\x93\x5f\xa7\x59\xad\xb3\xb4\x11\x96\xd0\xa4\xef\x1f\xef\x1f\x3e\x4e\x1f\xa2\x28\xaf\x30\x7f\x3d\x3d\x83\x55\x04\xfc\x87\x79\xa5\x21\xfe\xd4\x29\x25\xd5\x1c\xa4\x02\x7f\xec\xc4\x62\xec\x01\x4c\xdc\x08\x55\xc0\xc5\x02\x64\x8b\xa6\x84\xdb\xb4\xc0\x45\xaa\xba\xba\x86\xab\xdb\x77\x63\x78\x7b\x83\x55\xa0\xb9\x7d\x77\x05\x71\x00\xb1\x13\xa5\x89\xf9\x2c\x89\xba\xc6\x62\x04\x0f\xff\x49\xe2\x27\x46\xf1\x0d\x20\x2f\x61\x7c\x03\xeb\x03\x05\xf7\xbb\x67\xe1\xe9\x6e\x3a\x7d\xf8\x2b\xbc\xee\xa1\x97\xd1\x3a\x8a\x64\x09\x5f\xbe\xc0\x60\x0c\x93\x09\xc4\x5e\x64\x0c\x5f\xbf\xde\x30\x8a\x63\x52\x41\xaa\xdb\x8d\x4a\xe9\x0a\xf2\xcc\xd9\x95\xd2\x58\x02\x61\xe6\x5d\x83\x8a\x17\x25\xe7\xe0\x43\x25\xb4\x54\x23\x81\x12\x0d\x42\x2f\xd9\xc9\xdd\xfb\xcf\x77\xff\x4e\x21\xc3\x70\xce\x77\x90\x2b\x37\xd5\x43\xf7\x22\x9b\x11\x75\x2f\x96\xf6\xe0\x70\x08\x2e\x14\xf6\xb6\xe3\xb7\x50\x1a\xdd\xc0\xe0\xca\x9f\x74\x2d\xd8\x4a\x77\x75\xe1\x38\x5b\x61\x2d\x97\x4e\x2a\xd2\x9e\x42\xb4\xed\x36\xd7\x88\xd7\x33\xa6\xb0\x93\x78\xf0\x47\xbc\xff\x35\x58\x6d\x97\x27\x83\x4b\x58\xc3\x09\x7c\xc2\x86\x7b\x03\x9e\x0f\xb4\x1f\x87\x8f\x7f\x80\x7b\xb5\x1c\x88\x20\xdf\x71\xd0\x6a\x96\x31\x84\x2d\xfc\x40\xa5\x33\x43\x95\xa0\xc4\x42\x8d\x25\x85\x14\xf9\xc6\x36\xad\xce\xa2\xdd\xb8\x07\xdb\x65\x96\x24\x75\xae\x99\x87\xdc\xc3\x7c\xc3\xdd\x6b\xc4\x2b\xf2\x99\x41\x4f\x04\x5c\x33\x2e\x1f\xd7\xd0\x72\x3a\xdc\x5b\xde\x3b\x3f\x3a\x84\x1e\xc1\x60\x5b\x0b\xee\x7c\x46\xf4\xdc\xc3\xe1\x30\xa7\x4e\xd4\xdb\x47\x1e\x9f\x46\x07\x0e\x4f\x7d\xa3\x0c\x76\xc2\xdf\xc0\x65\x9a\xd8\x74\xb5\xda\x5e\x58\xaf\xd3\x24\x1e\x8c\xe3\x24\x4d\xce\x9c\xfc\x8f\x9a\x36\x4a\x7a\x4c\x58\x54\xa9\x4d\xee\xda\xbc\xc1\xb9\xc8\x24\xfb\xf2\xd3\xc4\xc6\xdc\x66\x7c\x51\x42\x13\x3b\x2f\x2c\x2e\x17\x0a\x2a\xb1\x70\x73\x87\x85\xcc\x49\x64\xb5\x2b\xa3\x71\xc8\x11\x7c\xe6\x75\xe7\xf2\xc4\x26\xe4\x8a\x80\x5c\x73\xab\xc3\x63\x42\x2d\xb9\xbc\xaa\xe4\xd9\x75\x2d\xef\x82\x33\xfb\x01\xf6\xe2\x7b\xad\x12\xda\x81\xb6\x01\x48\x1b\x0c\x9b\x4e\xcd\x36\xfd\x31\xd9\x0c\xd4\xde\xb7\x93\x09\x71\x14\xe9\x8e\x58\x04\xe7\x32\xf8\xe9\x0a\x5b\xa7\xbc\x9d\xf5\x52\x15\xba\x9f\x59\xf9\x3f\xce\x9a\x6c\x49\xb8\x8b\x30\xdc\xe4\x00\xc9\x40\xf2\xa2\x12\x70\xff\x21\xce\x0b\xe5\x12\x1d\x9d\x3f\xdf\x3f\x41\x20\x00\x47\x70\x0d\x2f\xa7\xa3\xf3\x97\xb3\xd1\x79\xfa\x32\x4e\xdb\x64\x1f\x3e\x7c\xf8\x93\xa9\x79\x9f\x23\x07\x70\x1d\xc3\x01\xd4\xfc\xc9\xf1\x00\x9e\x7e\x42\x3f\x7c\xa1\xcd\x9c\x69\xee\x34\x85\x39\x5a\x2b\xcc\x12\x32\xd6\xf1\xc8\x39\x6c\xce\x2c\x72\xa2\x3c\x24\xae\x98\xb4\x9b\x60\xad\xd0\x6d\xf6\xda\xbc\xc2\xb7\x8e\x37\x5c\x22\xfc\x91\xab\x97\x11\x19\xa1\x6c\x89\x26\x38\xb4\xbf\x62\xd1\x62\x0e\xe7\xc1\x56\x30\x61\xb7\xee\xce\xa2\x8c\x63\xec\x65\x41\x15\xf3\x4a\xb2\x33\xc6\xfe\x0a\x75\xa0\xdb\xb3\x3b\x8e\x3d\x79\xb4\x19\xa8\x59\x21\x48\x4c\x92\x55\x7c\xb4\x5a\xf1\x75\x32\x38\x7a\x90\x0c\x21\xfe\xc9\xb6\xc7\xfe\xb8\xe5\x50\x47\x5c\x38\xe4\x91\xed\x64\x9d\x44\xc1\xdd\xa1\xb6\xe8\x7b\x00\x00\x00\xff\xff\x37\xbd\x4f\xf1\xed\x06\x00\x00")

func testingBashtestletsIperfBytes() ([]byte, error) {
	return bindataRead(
		_testingBashtestletsIperf,
		"testing/bashtestlets/iperf",
	)
}

func testingBashtestletsIperf() (*asset, error) {
	bytes, err := testingBashtestletsIperfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testing/bashtestlets/iperf", size: 1773, mode: os.FileMode(493), modTime: time.Unix(1478831961, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _factsCollectorsGet_addresses = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8f\xcd\x6a\x22\x41\x10\xc7\xef\xf3\x14\xb5\xe3\xc0\xe8\xb2\x6b\x27\x1e\x72\x98\x20\x24\x88\x87\x40\x3e\x0e\x21\x27\x23\xa1\xa7\xbb\x74\x2a\xd1\xee\xa1\xab\x54\xc4\xf8\xee\x29\x05\x35\x7d\xaa\xae\xff\x07\xbf\xea\xfc\x31\x2b\x4e\xa6\xa6\x60\x30\xac\xa1\xb6\xdc\x64\x59\x07\x46\xb1\xdd\x26\x9a\x37\x02\x83\xab\xeb\x1b\x78\xb2\x22\xf0\xc2\x1b\xbb\x90\x3e\xbc\x31\x42\x4c\xb0\x8c\x9e\x66\xe4\xac\x50\x0c\x10\x67\x20\x0d\xb1\x26\x39\xae\x92\x43\x70\xd1\x23\x10\xc3\x3c\xae\x31\x05\xf4\x50\x6f\xd5\x81\xb0\x20\x87\x41\x0b\xda\x14\xd7\xe4\x75\xdf\x60\xc2\x4a\x73\x8d\x48\xcb\x95\x31\x73\x92\x66\x55\xf7\x5d\x5c\x1a\x89\xde\xab\xef\x13\x9d\x1c\x67\x53\x2f\x62\x6d\x96\x96\x05\x93\x79\x7c\x18\x8d\x9f\x5f\xc7\x59\x66\x53\xb2\xdb\x0f\x96\x34\x2c\x77\xf9\xbd\xf7\x09\x99\x91\xf3\x6a\x92\x97\x2a\x9e\xfe\xc3\x6e\xd1\xa5\x16\x0e\x7f\xf8\x06\xb0\x9b\x2f\x28\x0d\x05\x14\xb3\x6b\x13\x05\x81\x62\xb0\x2f\x55\x61\x45\x2a\xd9\xbc\x9b\xfe\x5f\x63\xca\x5e\x2f\xcb\x66\x7a\x2b\x01\x05\xc8\x8b\xdd\xb9\x6e\x72\x37\xdd\xe7\xb7\xe0\x63\x06\xfa\x3a\x87\x73\x9c\x0a\xea\xa1\xfc\xb8\xba\x50\x15\xe7\xb1\xa0\x32\xff\xa7\x54\x3e\x06\xfc\xcd\xad\xbd\xa7\xb9\xaa\xfe\x2b\xc7\x74\xaf\xe8\xe8\x9a\x08\x97\xf0\x4f\x00\x00\x00\xff\xff\xd8\xb5\x56\x13\xa9\x01\x00\x00")

func factsCollectorsGet_addressesBytes() ([]byte, error) {
	return bindataRead(
		_factsCollectorsGet_addresses,
		"facts/collectors/get_addresses",
	)
}

func factsCollectorsGet_addresses() (*asset, error) {
	bytes, err := factsCollectorsGet_addressesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "facts/collectors/get_addresses", size: 425, mode: os.FileMode(493), modTime: time.Unix(1477387277, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _factsCollectorsGet_hostname = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2c\x8c\xbd\x4a\xc4\x50\x10\x85\xfb\x3c\xc5\x98\x15\xa2\xcd\x8e\x5a\x58\xa4\x5d\x16\x14\xfc\x29\xc4\x4a\x2c\xee\xcf\x6c\xee\x48\x72\x27\xdc\x99\x44\x82\xf8\xee\x5e\x64\xbb\x73\x0e\xdf\x77\x76\x17\xb8\x68\x41\xcf\x19\x29\xaf\xe0\x9d\xa6\xa6\xd9\xc1\x41\xe6\xad\xf0\x90\x0c\xee\x6e\x6e\xef\xe1\xd9\x99\xc1\xab\x7e\xbb\xd1\xf6\xf0\xae\x04\x52\x60\x92\xc8\x27\x0e\xce\x58\x32\xc8\x09\x2c\xb1\x56\x53\x65\x29\x81\x20\x48\x24\x60\x85\x41\x56\x2a\x99\x22\xf8\xad\x12\x04\x23\x07\xca\xf5\x60\x2e\xb2\x72\xac\x7b\xa2\x42\x7d\xf5\x92\xd9\xac\x3d\xe2\xc0\x96\x16\xbf\x0f\x32\xa1\x49\x8c\x95\xfb\xa2\x60\xff\x19\xfd\x28\x1e\x27\xa7\x46\x05\x9f\x1e\x0f\xc7\x97\xb7\x63\xd3\x50\x48\x02\xdd\x4f\xfb\x20\x6a\xd9\x4d\xd4\xf6\x1f\x6d\x77\x79\x95\xce\xf5\xba\x6b\x3f\x7f\xbb\xbf\x00\x00\x00\xff\xff\x92\x35\x2d\x49\xe7\x00\x00\x00")

func factsCollectorsGet_hostnameBytes() ([]byte, error) {
	return bindataRead(
		_factsCollectorsGet_hostname,
		"facts/collectors/get_hostname",
	)
}

func factsCollectorsGet_hostname() (*asset, error) {
	bytes, err := factsCollectorsGet_hostnameBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "facts/collectors/get_hostname", size: 231, mode: os.FileMode(493), modTime: time.Unix(1477387277, 0)}
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
	"testing/bashtestlets/http": testingBashtestletsHttp,
	"testing/bashtestlets/iperf": testingBashtestletsIperf,
	"facts/collectors/get_addresses": factsCollectorsGet_addresses,
	"facts/collectors/get_hostname": factsCollectorsGet_hostname,
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
	"facts": &bintree{nil, map[string]*bintree{
		"collectors": &bintree{nil, map[string]*bintree{
			"get_addresses": &bintree{factsCollectorsGet_addresses, map[string]*bintree{}},
			"get_hostname": &bintree{factsCollectorsGet_hostname, map[string]*bintree{}},
		}},
	}},
	"testing": &bintree{nil, map[string]*bintree{
		"bashtestlets": &bintree{nil, map[string]*bintree{
			"http": &bintree{testingBashtestletsHttp, map[string]*bintree{}},
			"iperf": &bintree{testingBashtestletsIperf, map[string]*bintree{}},
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

