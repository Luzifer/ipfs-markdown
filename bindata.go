// Code generated by go-bindata.
// sources:
// viewer.html
// DO NOT EDIT!

package main

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

var _viewerHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\x6d\x6f\xdb\xb6\x13\x7f\xdf\x4f\x71\x65\x0b\xd4\xce\xbf\x92\xea\x7f\x56\x2c\x4b\xac\x00\x59\x16\x6c\x01\xba\x76\x28\xb2\x27\x14\xc5\x40\x8b\x27\x89\x09\x25\xaa\x24\xe5\xc4\x08\xf2\xdd\x77\xd4\x83\x2d\x29\x49\x87\x16\xd8\x5e\x18\xe6\x9d\xef\x7e\xbc\xbb\xdf\xdd\xd1\xcb\xa7\x3f\xbc\x3b\xbd\xf8\xf3\x97\x33\xc8\x5d\xa1\x8e\x9f\x2c\xfd\x17\x28\x5e\x66\x31\xc3\x92\x1d\x3f\x01\x58\xe6\xc8\x85\x3f\xd0\xb1\x40\xc7\x21\xc9\xb9\xb1\xe8\x62\x56\xbb\x34\x38\x60\xc3\x9f\x72\xe7\xaa\x00\x3f\xd5\x72\x1d\xb3\x3f\x82\x5f\x4f\x82\x53\x5d\x54\xdc\xc9\x95\x42\x06\x89\x2e\x1d\x96\xe4\x77\x7e\x16\xa3\xc8\x70\xe4\x59\xf2\x02\x63\xb6\x96\x78\x5d\x69\xe3\x06\xc6\xd7\x52\xb8\x3c\x16\xb8\x96\x09\x06\x8d\xf0\x12\x64\x29\x9d\xe4\x2a\xb0\x09\x57\x18\x2f\x7a\xa0\xa7\x41\x00\x17\x39\x02\x5f\xe9\x35\xc2\x3e\x34\xc0\x8e\x67\x16\xf6\x8a\xda\xba\x3d\x02\x2d\x10\x52\x69\xac\x23\x08\x70\x64\xea\x73\x3b\x02\x5e\x6e\x40\x93\x68\x1a\xb9\xbf\x1b\xbc\x53\xeb\xb3\xc7\x53\x87\x66\xcf\xbb\x58\x6c\x21\x83\xa0\xbb\xd5\x49\xa7\xf0\xf8\xf6\x16\x56\x86\x97\x42\x96\x19\xdc\xdd\x2d\xa3\x56\xfb\x64\x17\xd8\xf7\x5a\x3b\xeb\x0c\xaf\x76\x9e\x4a\x96\x57\x60\x50\xc5\xcc\xba\x8d\x42\x9b\x23\x52\xe6\xb9\xc1\x34\x66\x51\x94\x88\x32\xbc\xb4\x02\x95\x5c\x9b\xb0\x44\x17\xad\x7a\x84\x68\x3f\xdc\x0f\x5f\x47\x89\xb5\x3b\x5d\x58\xc8\x32\x24\x0d\xfb\x3a\xec\x5c\x66\xb9\xa2\x8f\x23\x75\x74\x10\x1e\x84\xaf\xa2\xd6\x2f\x12\x98\xf2\x5a\xb9\x7f\x0d\x3f\x93\x2e\xaf\x57\x03\xf8\x5d\xc9\x7e\xba\xf8\xf9\xcd\x6b\xb0\xb9\x2c\x88\x21\x01\xef\xd1\x56\xba\x14\x84\x00\xa9\x36\x70\x7e\x76\x00\xb6\xae\x7c\xbb\x80\x4e\x3b\x63\x54\x58\x10\x75\xb6\x71\x28\x50\x48\x0e\x9f\x6a\x34\x12\x07\x84\x79\xe8\xdf\x4f\xde\xbf\x3d\x7f\xfb\xe3\xe1\x10\x54\x68\xb4\xe5\x0b\x07\xd7\xda\x5c\x81\x4c\x61\xa3\x6b\xf0\x0d\xd9\x34\x4a\xc5\x33\x24\x89\x53\xfb\x28\x3c\x8c\xa2\x11\xdc\x07\xb2\x56\x8e\x22\x82\xef\x3e\xb6\x5a\xd2\xdb\xc4\xc8\xca\x81\x35\x49\xcc\xfc\x5c\x58\xf2\xd2\xd6\x86\x05\xbf\xf1\xe5\xa1\xbe\x8a\xfc\xb0\xbd\xa6\xfc\xd6\xc4\xe8\xb7\xe1\xff\x77\x72\x53\x8e\x4b\xaa\xc6\x32\x6a\x61\xbe\x04\xd5\xb4\x29\x45\x8b\xf0\x1b\xc2\xec\xa4\x47\x10\x97\x4f\x3f\x20\x35\x6d\xfa\xb1\x4d\x67\x19\xf5\xc3\xbe\x5c\x69\xb1\xe9\x6c\x4a\xbe\x86\x44\x71\x6b\x63\x46\xc7\x15\x37\xd0\x7e\x05\x5d\x6f\xb0\x6d\x74\x42\x6e\x2d\xfd\x10\x71\x59\xa2\x09\x52\x55\x4b\xb1\xb5\xe9\xc7\xc1\x0f\x4b\xc3\x92\xd3\x59\xa6\x10\x32\x74\x90\x19\x5d\x57\x28\x1a\x7a\x57\xe8\x68\xe4\xa0\xd0\x2b\x2a\x38\x08\x69\x2b\xc5\x37\xdb\xaa\x4f\x6f\xeb\x02\xf2\xd1\xa3\x19\xdc\xe5\x13\xa9\x9d\xd3\x34\xed\x9b\x8a\x16\x4c\x2b\xb0\x89\x5b\x17\x42\xa2\x95\xe2\x95\x45\xc1\x40\x70\xc7\x3b\xb5\x4f\xa5\xd5\xf7\x6a\x6e\x32\xbf\xfe\x9e\xad\x6c\x80\x37\xbc\xa8\x14\x06\x1d\x50\x6f\x19\x2c\x18\x70\x23\x39\xfd\x5e\x51\x8e\x28\x62\x96\x72\x65\x71\x14\x99\x67\x93\x7e\xed\x63\xb1\x26\xd0\xa5\xda\xb0\xe3\x8b\x36\x1a\x82\x94\x19\xed\x4e\x5d\x12\x65\x64\xf7\x19\x57\x49\xc5\x0e\xe8\xfe\x86\xdd\xff\xc4\x74\x19\xb5\x95\x1c\xe9\xf8\xa4\xac\xcd\x42\xec\x97\xc2\x33\x76\x6f\x47\xf2\x01\x97\x11\x91\xd9\xcd\xfe\xb6\x47\x4e\xa9\x9a\x98\xb8\x66\xfe\x7c\x0f\xfa\x9d\x63\x5f\xfa\xee\x28\xe8\xcb\xf7\x4e\xbb\xb4\xfb\x7d\xed\xdb\xa6\xe1\xcc\x5f\xf0\x58\xa7\xf4\x14\xc1\x84\x32\x06\x92\x58\xfa\x2c\xa5\xa3\x64\xc7\xcd\xd7\xa3\x0d\x8e\xc6\xaf\xbb\x29\xe1\xd5\xb4\xf3\xf0\x66\x6a\x03\xf0\x1b\x2d\x1d\x4a\xe1\xd0\x57\xb4\xdb\xa8\x54\x3a\x59\xa5\xf6\xaf\x9c\xdb\x9c\x6a\xd7\xd4\x72\xa8\x18\x15\xb3\x2d\x68\x35\xe6\xab\xa9\xef\x58\x6c\x8a\x1c\x85\x93\x44\x07\x95\x1b\x99\x4d\x06\x7a\xb7\x00\x23\x02\xe8\xd7\xf6\x43\x0b\xe0\xc1\xf5\x60\xf4\x35\x7b\x94\xa0\xa0\x10\xc1\xe2\x15\x74\x27\x9d\xa6\xf4\x77\xc3\x8f\x94\x67\xa8\x23\x7b\x54\xb4\x37\x9a\xfb\xa6\x0a\xc3\x70\xda\x50\xf7\x84\x61\xa3\x8d\x96\xe9\x03\x8f\xd6\xa5\x7f\x3a\x36\xb4\x47\x17\x8b\x70\xbf\x93\x1e\xdb\xa3\xff\x00\x35\x7d\xbb\x2f\xa7\x4f\xf7\x97\x43\x16\xdc\x5c\xa1\x88\x5e\x35\x78\xad\xf0\x95\x48\x0f\x3c\xce\x3b\xd5\x7d\xc8\x21\x66\x5f\xe1\xe7\xb3\xb4\x2e\x13\xbf\xac\x66\xf3\xdb\x2d\x09\x6b\x7a\x2a\x76\x5d\x1a\x03\x9b\xb6\xf1\xd1\x6e\xe2\xbb\x0c\x88\xe9\x77\x95\xc7\xb1\xb3\xdb\x01\xc3\xdb\x78\x0e\xa1\xbf\x08\x66\x89\x16\x38\x87\xdb\x51\xdf\x1b\x74\xb5\x29\x21\x57\x97\x36\xdc\x3a\x9d\xd4\x4e\xb7\xd6\xe1\x9a\xab\x1a\x8f\x06\x2e\x77\xdb\xf3\xdd\x7c\x10\xce\xf3\x90\x76\xfc\x8c\x45\x42\x27\xb5\xff\x47\x11\x31\xf8\xdf\x2e\xf6\x97\xdb\x28\x66\xfe\x49\x98\x0f\x63\x78\x3e\x7b\xf1\xac\x6b\xd2\x17\xf3\xd0\x3f\xe9\xb3\x36\xb7\xd6\x74\x7e\x34\xba\x70\x7c\x1a\xf2\x46\x2b\xb6\x79\x82\x97\x51\xfb\xd7\xfc\xef\x00\x00\x00\xff\xff\x03\x87\x69\x9a\xab\x0b\x00\x00")

func viewerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewerHtml,
		"viewer.html",
	)
}

func viewerHtml() (*asset, error) {
	bytes, err := viewerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "viewer.html", size: 2987, mode: os.FileMode(420), modTime: time.Unix(1443017078, 0)}
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
	"viewer.html": viewerHtml,
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
	"viewer.html": &bintree{viewerHtml, map[string]*bintree{
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

