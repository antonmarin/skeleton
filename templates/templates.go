// Code generated for package templates by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/files/plugins/readme/README.md
// templates/files/testfile
package templates

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

var _templatesFilesPluginsReadmeReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x53\x4d\x6f\xeb\x38\x0c\xbc\xeb\x57\x10\x2f\xb7\xbe\xd8\x2e\x7a\x0c\x16\xc1\x7b\x49\xbb\x1f\x40\xbb\x28\x9a\x74\x81\xbd\x55\x96\x19\x59\x89\x4c\x0a\xa2\x9c\xd4\xff\x7e\x21\x3b\x6d\xd2\xbd\xbc\x93\x01\x7b\x38\x1c\xce\x8c\x67\xb0\x75\xc9\xa3\x52\xb3\x19\xdc\xa3\x98\xe8\x42\x72\x4c\x4a\xad\x7a\x71\x84\x22\x10\x22\x9b\xfc\xd4\x21\x78\x67\x74\xfe\x0a\x21\xa2\x20\x25\x47\xb6\x84\x87\xf7\x05\xac\x75\xd2\x9e\xed\x1c\x36\x2d\x87\x80\x11\x22\x5a\x27\x29\xea\x89\xeb\x39\x72\xd3\x9b\x71\xf2\xf5\xe5\x71\x01\xbf\xb5\x29\x05\x59\x54\x55\xc3\x9d\x76\x54\x1a\xee\x96\x57\x1b\x1b\x36\x7d\x87\x94\xc6\xe9\x2b\xb4\x61\xda\xf9\x1e\xc9\x60\xf9\x65\x70\x13\xd0\xb8\xdd\x59\xdb\x15\x5e\x4e\xda\x5a\x8c\x57\xe0\x4a\x30\x1e\x9d\xc1\x9f\xe5\xd0\xf9\xe5\xfc\x17\xd4\xd5\xc3\x11\x29\xad\x7a\x59\x8e\xf6\x6c\xf2\x3d\x68\x07\xe0\x5d\x36\x65\x8f\x26\x29\xf5\xf4\xcf\x33\x24\xe7\x3d\xdc\xde\x95\x77\xb7\x77\xb7\x80\x79\x06\x2a\xd8\x24\x5d\x4f\xbe\xce\x60\xa5\xc5\x19\x88\x4e\x0e\xa2\x54\x01\x8f\x7c\x82\x03\xf1\xc9\x63\x63\x31\xb3\x4d\x2b\x55\x01\x4f\xbd\x24\xa8\x11\x5c\xb6\xf8\xd3\xb3\x1a\x77\x1c\x11\xfe\xfe\x77\x62\x7b\x8e\x8e\x8c\x0b\x1e\x47\xb2\x57\x41\x68\x50\x9c\x25\xa8\x87\x33\x13\x54\xf0\x7b\xd6\xd4\xe9\x90\xe9\xeb\x0f\x63\x13\x7f\x49\xd1\x3b\x3a\x08\x54\xf0\x93\xb4\x1f\x04\x3f\xf6\x34\x78\x44\xcf\x41\x15\xb0\x66\x4a\x91\xbd\xc7\x08\x39\xd4\x4f\x1e\x31\x48\x3a\x3a\x86\xef\x20\x1d\x1f\x10\x12\x4a\x52\x05\x6c\xef\xef\xe1\x3b\x10\x4f\x7a\x54\x01\x2d\xbe\x6b\xcb\x94\x71\xc1\xbb\x04\x3d\xb9\x34\x77\x94\xd0\x4e\xd5\x18\x07\x45\x15\xa0\x8d\xc1\x90\x34\x99\x89\x4b\x40\xa7\x4b\x38\xd6\xa5\xb6\xaf\xc7\x44\x96\x65\x59\x8e\x61\x64\x69\xda\xa4\xd1\x82\x2d\x9a\x96\x9c\xd1\x7e\x01\x1b\x8c\x16\x7e\x08\x46\xab\x0a\x38\xd7\x6e\x01\x7f\x1d\x35\xc1\x0f\x77\xd4\x59\xd3\xdb\x2c\xa1\x69\xdf\x4a\x78\xc1\xa3\xc3\xd3\xfc\xc3\x3d\x2d\xe2\x64\x94\x50\xaa\x02\x5e\x30\x70\x4c\x50\xf7\x76\xd4\xb2\x77\x51\x5f\x37\x63\xf5\xfa\xc7\x26\xc3\xd6\x7d\x8c\x39\x6f\x0e\x78\xbe\x48\x92\x4e\xfd\x45\x7f\xae\xe1\xf8\xe6\xba\xb2\xe5\xb9\x4f\xda\x3a\xb2\x4a\xfd\xc9\xa7\x1c\x4d\x83\xc1\xf3\x00\x3b\x8e\xff\xb7\xc3\x91\x05\x4d\x0d\x34\xd8\x31\xb4\x18\xa7\xdf\x75\x1b\xb9\xaf\x3d\x4a\xcb\x9c\x46\x9e\xb5\x4e\xa6\xcd\xd8\xcd\xd0\xed\x98\x06\x78\x78\xcf\x34\x59\xd5\xb7\xed\x10\x10\x30\x46\x8e\x0b\xf8\x96\xd5\xdd\xdc\xdc\xcc\xd5\x79\x25\xe1\x09\x9a\xbe\x0b\xa5\x52\x5b\x06\x61\x7f\x44\xd8\x46\x4d\x92\x3d\xb8\x90\x44\x3c\xe3\x75\x08\x97\x74\xf6\x48\x07\x47\xd7\xf7\x55\x7b\xae\xab\xa5\xfa\x2f\x00\x00\xff\xff\x1a\x2d\xcd\x68\x5b\x04\x00\x00")

func templatesFilesPluginsReadmeReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_templatesFilesPluginsReadmeReadmeMd,
		"templates/files/plugins/readme/README.md",
	)
}

func templatesFilesPluginsReadmeReadmeMd() (*asset, error) {
	bytes, err := templatesFilesPluginsReadmeReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/files/plugins/readme/README.md", size: 1115, mode: os.FileMode(420), modTime: time.Unix(1605639460, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFilesTestfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x49\x2d\x2e\xe1\x02\x04\x00\x00\xff\xff\xc6\x35\xb9\x3b\x05\x00\x00\x00")

func templatesFilesTestfileBytes() ([]byte, error) {
	return bindataRead(
		_templatesFilesTestfile,
		"templates/files/testfile",
	)
}

func templatesFilesTestfile() (*asset, error) {
	bytes, err := templatesFilesTestfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/files/testfile", size: 5, mode: os.FileMode(420), modTime: time.Unix(1605723061, 0)}
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
	"templates/files/plugins/readme/README.md": templatesFilesPluginsReadmeReadmeMd,
	"templates/files/testfile":                 templatesFilesTestfile,
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
	"templates": &bintree{nil, map[string]*bintree{
		"files": &bintree{nil, map[string]*bintree{
			"plugins": &bintree{nil, map[string]*bintree{
				"readme": &bintree{nil, map[string]*bintree{
					"README.md": &bintree{templatesFilesPluginsReadmeReadmeMd, map[string]*bintree{}},
				}},
			}},
			"testfile": &bintree{templatesFilesTestfile, map[string]*bintree{}},
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
