// Code generated by go-bindata.
// sources:
// cmd/streepjes/test/files/testdata.sql
// DO NOT EDIT!

package test

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

var _cmdStreepjesTestFilesTestdataSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x92\x4d\x8f\xda\x30\x10\x86\xef\xfc\x8a\xd1\x5e\xc2\x4a\x39\x24\xf6\xd8\x09\xe2\xd4\x6a\x69\x45\xbb\xb0\xd5\x2e\xfd\x3a\xad\x26\xb1\x03\x2e\xf9\x92\x1d\x5a\xf1\xef\xab\x90\x84\x85\x48\xad\xb6\xa7\xce\x01\x8d\x87\xc9\xe3\x79\x5f\xcf\xdd\xe2\x7e\xb1\x59\xc0\xbb\xc7\x87\x15\xdc\x54\x56\x69\x7b\x33\x9f\x5c\x16\x0f\x4e\xdb\xeb\x4a\xa1\x8b\x64\x5c\xab\x6d\xa5\x0e\x69\x73\x5d\x4c\xa9\xd1\xdb\xca\x1e\xe7\x93\xc9\x72\xfd\xb4\x78\xdc\xc0\x72\xbd\x79\x38\x11\xa7\x25\x15\xda\x3f\xa5\x5d\x56\x93\x73\xbf\x2a\xab\x7c\xb0\x55\xae\x6f\xe1\xcb\x9b\xfb\xcf\x8b\xa7\x09\x00\xc0\xd4\x5b\x91\x6d\x8e\x9e\x0f\x5e\xd1\x27\xdf\x3c\x86\x9c\xc9\x90\x21\x0f\x38\x32\x14\x24\xb8\xe0\x1c\x65\x8c\xb1\x4c\x22\x92\x31\x0f\x65\x84\x1c\x13\x44\x8c\x90\x4b\xc1\x43\x11\xc8\x19\x8f\xa5\x90\x4a\x66\x51\x28\x13\xc1\x91\x30\x12\x28\x33\xde\x7e\xa5\x38\x4a\x25\x13\x96\xa1\x8a\x28\x12\x32\x8e\x66\x72\x26\x23\x29\x71\x86\x89\xcc\x04\x8b\xe2\x28\xc0\xcc\xf3\x21\xbc\xf5\xfb\xb9\xee\x74\x41\x9e\x0f\xe0\xa9\x3e\xf9\xaf\x73\xb1\xdb\x91\xd1\xdd\x43\x4d\xd3\xfc\x90\xf8\xd0\xd9\x9c\x50\x4e\x65\x3a\xf2\x37\xf4\xc1\x5b\x1f\xcb\xbd\x86\xb7\xda\x6e\x5d\x27\xa5\x8d\x60\x50\xda\x76\x7c\x30\x79\x03\xdf\xeb\x97\x7f\xdb\x08\x83\xe0\xdc\xc4\x7c\xf0\x56\xd4\xec\xcc\x0f\x07\x9f\xac\x29\x9d\x2e\xbb\xd6\x90\x71\xbc\x6c\x5a\xe6\x4e\xc3\x4f\x2a\x41\x69\x0b\x5f\xb5\x51\xba\x75\x95\x71\x14\xd7\x28\xdb\x98\x12\x3e\x92\xdd\x57\xce\xf5\xb7\x06\x63\x8d\xc3\x8e\x9d\x16\x6a\xb4\x36\xca\x9a\x72\xef\xbc\xf3\x73\xb9\x92\xd2\xf6\x3c\x42\xf4\xbb\x3b\x1d\x50\xcf\x46\x0d\x76\xd5\xd6\xa4\xfa\xb9\x26\x4b\x49\x55\xe5\xc3\x79\x9b\x93\x32\xd4\x54\xd6\x5d\x5d\x38\x8a\xd6\xb4\x97\xf0\x32\x6b\xdc\xa5\x73\x27\x63\xc2\xc0\x6f\x7f\xfa\x09\xff\x0e\x78\x6f\xab\xdc\xa5\xbb\x2b\x46\xc8\x5a\x80\x78\x1d\xc0\xd5\x3a\x35\x44\x79\x62\xb4\x1d\x28\x2c\x08\x7c\x60\x7f\x04\xb0\xb1\x04\xdd\x8c\x34\x74\x80\xe0\x75\x80\xa6\x72\x8d\x81\x3d\xd1\x85\x13\x9d\x04\xf6\x2f\x80\x1d\x15\x50\x9e\x31\xa1\xe8\x3d\x98\xff\x0e\x00\x00\xff\xff\xfb\x74\x95\x1c\xcc\x04\x00\x00")

func cmdStreepjesTestFilesTestdataSqlBytes() ([]byte, error) {
	return bindataRead(
		_cmdStreepjesTestFilesTestdataSql,
		"cmd/streepjes/test/files/testdata.sql",
	)
}

func cmdStreepjesTestFilesTestdataSql() (*asset, error) {
	bytes, err := cmdStreepjesTestFilesTestdataSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/streepjes/test/files/testdata.sql", size: 1228, mode: os.FileMode(420), modTime: time.Unix(1616165499, 0)}
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
	"cmd/streepjes/test/files/testdata.sql": cmdStreepjesTestFilesTestdataSql,
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
	"cmd": &bintree{nil, map[string]*bintree{
		"streepjes": &bintree{nil, map[string]*bintree{
			"test": &bintree{nil, map[string]*bintree{
				"files": &bintree{nil, map[string]*bintree{
					"testdata.sql": &bintree{cmdStreepjesTestFilesTestdataSql, map[string]*bintree{}},
				}},
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