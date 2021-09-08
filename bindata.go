// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// unit/.DS_Store
// unit/service.sh
package main

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _unitDsStore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xd8\x31\x0a\x02\x31\x10\x85\xe1\x37\x31\x45\xc0\x26\xa5\x65\x1a\x0f\xe0\x0d\xc2\xb2\x9e\xc0\x0b\x58\x78\x05\xfb\x1c\x5d\x96\x79\x60\x60\xd5\x4e\x8c\xcb\xfb\x40\xfe\x05\x37\x2a\x16\x31\x23\x00\x9b\xee\xb7\x13\x90\x01\x24\x78\x71\xc4\x4b\x89\x8f\x95\xd0\x5d\x1b\x5f\x43\x44\x44\x44\xc6\x66\x9e\xb4\xff\xf5\x07\x11\x91\xe1\x2c\xfb\x43\x61\x2b\xdb\xbc\xc6\xe7\x03\x1b\xbb\x35\x99\x2d\x6c\x65\x9b\xd7\x78\x5f\x60\x23\x9b\xd8\xcc\x16\xb6\xb2\xcd\xcb\x4d\xcb\x38\x7c\x18\xdf\xd9\x38\xa1\x18\xa7\x10\x2b\x6c\xfd\xce\x77\x23\xf2\xef\x76\x9e\xbc\xfc\xfe\x9f\xdf\xcf\xff\x22\xb2\x61\x16\xe7\xcb\x3c\x3d\x07\x82\xf5\x0d\x00\xae\xdd\xf5\xa7\x43\x40\xf0\x3f\x0b\x0f\xdd\x5a\x1d\x04\x44\x06\xf3\x08\x00\x00\xff\xff\x6a\x00\x88\x6d\x04\x18\x00\x00")

func unitDsStoreBytes() ([]byte, error) {
	return bindataRead(
		_unitDsStore,
		"unit/.DS_Store",
	)
}

func unitDsStore() (*asset, error) {
	bytes, err := unitDsStoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "unit/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1630927699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _unitServiceSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x96\x6d\x6f\xda\xc8\x17\xc5\xdf\xf3\x29\xce\x7f\xfe\xde\x1a\xba\xb2\xc3\x43\x31\x14\x44\xb5\x94\x78\x53\x24\x1a\x28\xf6\x36\xaa\x36\x2b\xe4\x9a\x81\x8c\x6a\xc6\xae\x67\x9c\x26\xea\xe6\xbb\xaf\xc6\x36\x8e\xc9\x03\x4a\xd2\x28\x79\x41\x3c\xdc\x99\x73\xef\x3d\xf3\xf3\x4d\xfe\xff\xbf\x83\xaf\x8c\x1f\x7c\xf5\xc4\x59\xc5\x91\x5e\x2c\x3f\x5e\x3a\x9f\x26\x6d\xb3\x53\xad\xe1\x67\x05\x00\x96\xa1\xff\x8d\xc6\x88\x92\x20\xc0\xe6\x52\x7c\x0f\x7a\x6d\xb3\x63\x36\xdb\xe5\x60\x9c\x70\x18\x06\xf7\x36\x74\x90\x6e\x69\x9b\x1d\x18\x4b\x18\x11\x5a\xad\xba\xd5\x53\x1f\x30\x28\x3e\x7e\x71\x3e\x4d\x16\xf3\xe9\xd4\x5d\xcc\x86\x8e\x73\x32\x9d\x1f\x0e\x1a\xcd\xd6\x9b\xb6\xb5\xa3\x0c\xc3\x58\xd2\x95\x97\x04\xd2\xf0\x12\x79\x46\xb9\x64\xbe\x27\x59\xc8\x8d\x28\x48\xd6\x8c\x67\x39\x16\xdc\x93\xec\x9c\x2e\x22\x4f\x88\x1f\x61\xbc\x4c\xeb\x39\xf1\x98\x64\x7c\x0d\xb2\x2d\x83\x00\xe4\x33\x8d\x05\x0b\x79\x0f\x7a\xa6\xaf\x03\x42\xd5\x2d\x7b\xd0\x0f\xce\xbd\xf8\x20\x4e\xf8\x41\x7a\x60\x99\xff\x32\x55\x5c\x07\xa2\x30\x96\xbd\xb4\x07\x20\x35\x06\xa3\x70\xb3\x49\x38\x93\x97\x70\x68\x7c\x4e\x63\x54\x8f\x66\x93\x1a\x41\xab\x5e\xf6\x23\x08\xd7\x02\xdb\x12\xca\x01\x7a\x41\xfd\x22\x90\x3d\xc0\x48\xe2\x30\x94\x30\xa2\xdc\x0a\x83\x82\x8c\xe6\xf6\xd0\xb5\x71\x38\x74\x87\xef\x87\x8e\x8d\x8b\x84\x63\xf4\x61\x38\x1f\x8e\x5c\x7b\x0e\xc7\x76\x91\xc8\x55\x17\xa3\xe9\x64\xa2\xf6\xa9\xc5\x62\x4d\x39\x8d\xbd\x60\xe1\x33\xf2\x0b\x29\xff\x72\xec\xb9\x4a\xf7\x87\xfe\x9b\x8e\xf1\xa1\x7d\xec\x8e\xff\x1c\xdb\x87\x78\xff\x05\x7a\xb6\x5b\x7f\x92\xfc\xd1\x7c\x78\xec\xc2\xb1\x27\xf6\xc8\xc5\xf4\x58\xa5\x30\x5f\xc3\x9d\x42\xbf\x48\xb8\xae\xb2\x91\x7e\xe5\xaa\x52\x82\xb0\x6b\xd6\xf7\x40\xd8\x35\xeb\x66\xd3\xda\x07\x61\xd7\xac\x5f\x43\xd8\x7d\x38\x84\x99\xf2\xb3\x41\xd8\x35\xeb\xbb\x10\x66\xfa\xcf\x0a\xa1\x81\xa3\xd9\xe4\x7e\x08\xbb\x66\xfd\xee\x2b\x53\x16\xbd\x30\x84\x0f\x4a\xf9\x74\x08\xf7\xcb\x3f\x14\xc2\x9d\x51\x68\xed\x1d\x85\x96\xd9\x6e\xec\x1f\x85\xd6\x35\x85\x9d\xc7\x8c\x42\xa5\xfc\x8c\xa3\xd0\xba\x39\x0a\x95\xfe\xcb\x8e\x42\xeb\xbe\xc1\x61\xbd\xfc\x28\x7c\x40\xca\x5f\x19\x85\xfb\xe4\x1f\x47\xe1\x2c\x14\x72\x1d\x53\xf1\xf6\x1e\x10\xa3\x3c\xde\x7b\x6b\xde\x37\x0e\xa3\x6b\x89\x9c\xc5\xf6\x9b\x56\xb3\xa7\x3e\x54\x41\xb3\xa9\xe3\x1e\xcd\x6d\xe7\x16\x89\xb7\xa4\x0b\xa6\x4a\x8a\x0a\xab\xbc\x46\x45\x05\xe3\x4c\x22\x8a\x43\x9f\x0a\x01\x3f\xdc\x44\x01\x95\xb4\x8f\x98\x7a\xcb\x4b\xac\xc2\x18\x42\x35\x85\x24\xba\x13\x94\x92\xee\x2d\x67\xcb\x5d\x88\x04\x46\xf1\x05\x0c\x1f\x24\x4a\xdd\xf6\xa1\xdf\xc1\x8b\x4e\xf0\x44\xb5\xd3\x9b\x2c\xe0\x64\xec\x7e\xc0\xd6\xa8\x82\x83\x53\xf2\xe4\x14\x7a\x86\xc3\x70\x32\xc1\x6c\x3e\xfe\x3c\x9e\xd8\x47\xb6\xa3\xb0\x28\x3a\x38\x25\x17\x09\x3f\x25\x90\xa1\xaa\xa0\xaf\xba\xb9\xaa\x54\xc6\x62\xae\x3c\x2d\x98\x48\x6f\x5a\xcb\x66\x91\x7f\x46\xfd\x6f\x42\xc6\x03\xad\x99\xae\x63\x2a\x06\x5a\xb5\x6c\xb4\xd6\x40\xf3\xdd\xab\x06\xfe\xc5\x3a\xa6\x11\x88\xb6\x3d\x42\x6a\xe9\x09\xb6\xc2\xdf\x20\x5a\x4c\x05\xc1\x60\x00\x42\xf0\x4f\x1f\x6a\x00\xa5\x51\xf5\x43\xfd\xb3\x10\xa4\x9e\xbd\x01\x34\x10\xf4\x46\xa4\x91\x45\x56\x4c\x15\x9b\x83\xf3\xa0\x62\x03\x2a\x21\xd9\x86\x86\x89\x1c\x68\xad\x4a\xa1\x68\x70\x90\xf4\x8d\x50\x04\x6a\x4a\x01\x66\x96\x83\x89\x14\xaf\x81\x56\xcd\x4d\x01\x49\xe3\xe4\x76\x5f\x92\x6d\x18\x5f\x0f\x32\xf2\x7e\x9c\xb1\x80\x22\x6d\x34\x97\xc8\x9a\xad\xab\x6e\xf3\xeb\x2c\xba\x12\x01\xa5\x11\x1a\xc5\xfa\x11\x49\x4b\x5d\xa9\xe4\xda\xcf\xec\xe1\xea\xf7\xc6\xae\x67\xaa\xc3\xbc\xa5\xe2\x0e\xb4\x6c\x2f\x0c\xfa\x3d\x7d\x56\xb6\xdc\xba\x8b\x6b\xd7\xb1\xf2\x58\x40\x97\xa6\xca\x15\x26\xb2\xba\x3d\x52\x23\x78\xf7\xaa\xb9\x73\x60\x07\x87\xd4\xce\x9b\x3b\xe8\x05\x93\xa5\x8e\x57\x2c\xf7\x84\xd3\x4a\x29\xa3\x5a\x13\x75\xc9\x7e\xb8\xd9\x78\x7c\xa9\xee\xd5\xf7\x04\x85\x96\x7f\x01\x96\x55\xba\xfd\x03\x5d\xc3\xce\x3f\x79\xfd\xfe\x75\xb4\x6d\x76\xca\xd1\xb6\xd9\xd9\x8d\x5a\xbb\x51\x2b\x8f\x96\xde\xb4\x7c\x43\x69\x6a\xe6\x7b\x5e\xd7\xa0\x55\xb3\x92\xa3\x80\xaa\xfa\x18\x8f\x12\x89\xbc\xc8\xd4\x9f\x1a\xfa\xfd\x0a\x15\x9e\xff\x5f\x00\x00\x00\xff\xff\x78\x34\x52\x42\x17\x0d\x00\x00")

func unitServiceShBytes() ([]byte, error) {
	return bindataRead(
		_unitServiceSh,
		"unit/service.sh",
	)
}

func unitServiceSh() (*asset, error) {
	bytes, err := unitServiceShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "unit/service.sh", size: 3351, mode: os.FileMode(493), modTime: time.Unix(1630927699, 0)}
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
	"unit/.DS_Store":  unitDsStore,
	"unit/service.sh": unitServiceSh,
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

var _bintree = &bintree{nil, map[string]*bintree{}}

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
