package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _favicon_ico = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x96\x39\xe8\x13\x41\x14\x87\xdf\xfa\x37\xc4\x83\x04\x2d\x34\xab\x46\x62\x61\x50\x08\x1e\x68\x13\x41\xa2\x58\xc5\x46\x2c\xb5\x09\x08\x8a\x8d\x88\x78\x80\x4a\xa2\x89\x08\x5a\x04\x2c\xd2\x5b\x89\x96\x5e\x08\x9e\x9d\x09\x5e\x9d\x10\x09\x36\x01\x8f\x46\x03\x2a\xa2\x82\xd1\xf5\xf7\x32\x6f\x65\x18\x66\x37\xbb\x9b\x5d\x3e\x96\xcc\xcc\xfb\x66\xf2\xf6\xcd\x24\x44\x0e\x6e\xd7\xe5\xe7\x1a\xba\xbc\x90\x68\x39\x11\xad\x07\x68\xa2\xcd\xa4\xda\x27\x17\xfa\xde\x64\x14\xfe\xe5\x79\x1e\x3f\xd2\xe0\x12\x78\x0d\xbe\x82\xa7\xe0\x18\x98\xcf\xfd\x61\xe0\x5a\x00\x4e\x80\x47\xe0\x3d\xf8\x0c\xee\x80\xa2\xf4\x63\x56\x7a\xc1\x53\x59\x78\x09\x16\x87\xb8\x4b\xe0\x5d\x40\xec\x08\xe4\x40\x1b\xdc\x03\x9b\xc0\x52\xb0\x0b\xf4\xb4\x71\x57\x03\xdc\x4b\xc0\x5b\x6d\xdc\x27\x70\x00\xac\x00\xeb\xc0\x0d\x70\x5c\x5c\x8b\x8c\xd8\x94\xe4\xca\x8f\xdd\x67\xf4\xcf\x93\x35\xf9\xfd\x63\xb0\xc5\xe2\xe0\x39\xb6\x06\xac\x6f\x1b\xf8\x23\xf1\x5f\x78\xbd\x5a\xdf\x41\x23\x17\x9d\x00\xc7\x06\x79\x3a\xe0\x90\xbc\x97\x81\xac\x3d\x0b\x6e\x69\x8e\x93\x5a\x5c\x57\x6b\xff\x0e\xf2\x12\x33\x10\x07\xbb\x1c\xed\xfd\xf6\x2c\xef\x87\xeb\xe7\x94\xf6\x79\x08\xe6\xc0\x5a\x63\xdc\x43\x70\xd4\x12\xdf\x13\x77\x0b\xfc\x94\xb5\x1e\x01\xa7\xc5\xc5\xf5\x56\x31\x62\x6a\xa4\xea\x58\x6f\x6b\x80\xfb\x52\x2f\x57\xc0\x45\xf0\x4c\x72\xcb\xee\x3e\xcf\x6f\xe4\x6d\x99\xf8\x57\x1a\xae\x0f\xa4\xea\x5b\x6f\xab\x92\xda\x2f\x45\xc3\xd1\x10\xf7\x37\x90\xb3\xbc\x1b\xae\xd3\x8c\xe5\x7b\x9b\xf0\x16\xde\x6b\x89\xcf\x89\xfb\xb6\xe4\x85\xc7\xb5\x2d\x7b\x73\x9a\xbf\x60\xc4\x9c\x05\xab\xc0\x7e\x71\xf3\x1e\x7c\x02\x1e\x83\x8d\x96\x1a\x9d\xe6\xdf\x6d\xc4\xf0\x3e\x7d\x20\xf9\x2d\x85\xec\x7d\x47\x6a\x63\x9a\xff\x15\x85\x9c\x21\x21\xfe\x33\x11\xdc\x3e\xd7\x13\xf8\x87\x52\x1b\xd5\x10\xaf\xdf\xff\x03\xa4\x62\xfa\xfb\x31\xfc\x7c\x7e\xcc\xc5\xf4\x5f\x88\x91\x9f\x6b\x09\xf2\x53\x8a\xe1\xdf\x13\xd7\x2f\x73\x8c\x22\xb8\x7f\x81\x74\x42\xff\xf3\x08\xfe\x6e\x12\xb7\xf8\x3b\x11\xfc\xcd\x19\xfc\xab\xe5\xfb\x07\xb9\xf9\xbf\x80\x9b\xd4\x2f\x73\x1c\x06\xbf\x2d\x6e\x6e\xab\xcd\xe2\xd6\xe6\xd8\x09\x6e\x92\x3a\x9f\x3f\x82\xbb\xa0\x12\x25\xb6\xe9\xe1\x98\x39\x87\xd1\x75\xa2\x72\x39\x4f\x75\xb7\x42\x4d\x97\x8f\x1e\x75\x95\xf1\xd3\xb7\x1d\x77\x16\x3f\x67\xcc\x98\x76\x08\x05\x41\x7d\xfe\x3b\xe1\xfc\x7f\x3c\xbe\x5b\x1e\xfd\x0b\x00\x00\xff\xff\x5a\x96\x60\xc5\x9e\x09\x00\x00")

func favicon_ico() ([]byte, error) {
	return bindata_read(
		_favicon_ico,
		"favicon.ico",
	)
}

var _robots_txt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2c\x8c\x41\x0a\xc2\x30\x10\x45\xf7\x3d\xc5\x07\x77\x82\x99\x7d\x76\x82\x37\x10\x0f\x50\xeb\xb4\x09\xa4\x19\x49\x7e\x19\x8f\x6f\x40\xe1\xaf\x3e\xef\xbd\x13\xee\xaa\x48\xe4\x3b\x8a\xb8\x7b\x68\xf6\x34\x76\x7e\x18\xac\x6d\xe2\x8b\x54\xfb\x5d\x21\x71\x2f\x58\xad\xe1\x65\xcb\xb1\x6b\xe5\xcc\x6c\x15\x63\xc9\x1c\x34\x1c\x5d\xc1\xa4\xf8\xf3\xa3\x81\x35\x17\x9d\x1e\x5d\xdb\xe5\xba\x0d\x23\xe2\x3c\xdd\x72\x9f\x4b\x31\x8f\x90\xe9\x1b\x00\x00\xff\xff\x96\x37\x39\x67\x7e\x00\x00\x00")

func robots_txt() ([]byte, error) {
	return bindata_read(
		_robots_txt,
		"robots.txt",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"favicon.ico": favicon_ico,
	"robots.txt":  robots_txt,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"favicon.ico": &_bintree_t{favicon_ico, map[string]*_bintree_t{}},
	"robots.txt":  &_bintree_t{robots_txt, map[string]*_bintree_t{}},
}}
