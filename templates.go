// Code generated by go-bindata.
// sources:
// templates/compose-dev-template
// templates/compose-template
// templates/docker-template
// templates/jenkins-template
// templates/kubernetes-template
// DO NOT EDIT!

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

var _templatesComposeDevTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\xb1\x6e\xc3\x20\x10\x86\x77\x4b\x7e\x87\x93\xf7\xd8\x52\x47\xb6\x28\xa1\x6d\xda\xa6\x8e\x92\x0c\xdd\x10\x31\x17\x0b\x85\x40\x74\x60\x9a\xc7\xaf\x6c\xa8\xea\xa8\x4c\xdc\xcf\xf7\xdd\x71\x11\xc9\x6b\x67\x19\x54\x4f\x55\x59\x94\x85\x47\x8a\xba\x43\xcf\xca\x02\x40\x6c\xb6\xcb\x17\x2e\x3e\x97\x5b\x3e\xd5\x00\xfa\x2a\x7b\x64\xd0\x77\x54\x6b\xd7\x88\xdd\xbe\x7d\xe3\xab\xe3\x44\x34\x73\xda\xc8\x80\x3e\x24\xe7\x34\x68\xa3\xb2\x0f\xd0\x39\x1b\xf0\x1e\x18\xd4\xbf\x89\x72\xdd\x05\xe9\xac\x0d\xb2\xd9\xdd\x37\x62\xdd\xae\xde\xf9\x5e\x3c\x6f\x3e\x52\xd7\x24\xa0\x8d\x22\xc1\x35\xda\x98\xb2\xb1\xa9\xd4\x16\x49\x58\x79\x45\x06\xd5\xec\x2f\x55\x42\xbe\x1d\x5d\xb4\xed\x85\xd2\xc4\xa0\x19\x3c\x35\x9e\xba\x46\xde\x6e\xe9\x39\x3a\x33\x5c\xf3\xda\xe3\x59\x40\xcd\xfe\x53\x37\x47\x61\xce\x54\xe2\xb5\x3d\x1c\xc5\xae\xdd\x1f\x99\xe0\x5f\xbb\xf6\xc0\xd7\x53\x95\x87\x5a\x0c\xe3\xdc\xb9\xa2\xf0\x2c\x07\x13\xfe\x82\x13\x69\xd5\xe3\x22\xa3\x65\xf1\xe0\x64\x3a\xfb\x8a\x74\x44\x62\x59\x19\xb3\x47\x39\x63\x78\x0f\x48\x56\x1a\x06\x81\x06\x2c\x8b\x9f\x00\x00\x00\xff\xff\xec\x42\x5d\x5a\xe4\x01\x00\x00")

func templatesComposeDevTemplateBytes() ([]byte, error) {
	return bindataRead(
		_templatesComposeDevTemplate,
		"templates/compose-dev-template",
	)
}

func templatesComposeDevTemplate() (*asset, error) {
	bytes, err := templatesComposeDevTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/compose-dev-template", size: 484, mode: os.FileMode(420), modTime: time.Unix(1563422270, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesComposeTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\x4e\xc3\x30\x10\x44\xef\x91\xf2\x0f\x2b\xdf\x89\x25\x8e\xbe\x55\xad\x41\x05\x4a\xa2\xd0\x03\xb7\x95\x71\x96\x68\xd5\xd4\x8e\xd6\x6e\xe0\xf3\x11\x35\xa0\x70\xdb\x99\x7d\x33\x9a\x85\x24\x71\x0c\x06\xd4\xad\xaa\xab\xba\x4a\x24\x0b\x7b\x4a\xa6\xae\x00\x70\x7f\xd8\xdc\x5b\x7c\xde\x1c\xec\x55\x03\xf0\xd9\x8d\x64\x60\xf4\xd2\x70\xd4\xd8\xf5\xed\x83\xdd\x1e\xaf\x84\x5e\xd3\x93\xcb\x94\x72\xc9\xbc\x5d\x78\x1a\x7e\xf2\x00\x3e\x86\x4c\x9f\xd9\x40\xf3\xeb\x0c\xd1\x9f\x48\xde\x79\x22\xb3\xba\x93\xc6\x5d\xbb\x7d\xb4\x3d\xde\xed\x9f\x4a\x6b\x09\x50\x58\xb0\xc0\x0d\x85\xa5\x78\xdf\xa5\x8e\x03\x09\x06\x77\x26\x03\x6a\xb5\x45\x15\xe4\x23\xca\x89\xc3\x88\x03\x8b\x01\x7d\x49\xa2\x93\x78\xed\xe6\xb9\xbc\xe7\x28\x39\xfd\x8d\xbc\x01\x85\xf6\xb5\x6b\x5f\xec\x0e\xbb\xb6\x3f\x9a\x7f\x4a\xd5\xd5\x57\x00\x00\x00\xff\xff\xbf\x61\xfb\xaa\x37\x01\x00\x00")

func templatesComposeTemplateBytes() ([]byte, error) {
	return bindataRead(
		_templatesComposeTemplate,
		"templates/compose-template",
	)
}

func templatesComposeTemplate() (*asset, error) {
	bytes, err := templatesComposeTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/compose-template", size: 311, mode: os.FileMode(420), modTime: time.Unix(1563422270, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDockerTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0b\xf2\xf7\x55\x88\x77\x72\x0c\x76\x8d\xf7\xf4\x75\x74\x77\xe5\xe5\xe2\xe5\x72\x8d\x08\xf0\x0f\x76\x55\x88\x87\xd0\x2e\xf1\x01\xfe\x41\x21\x80\x00\x00\x00\xff\xff\x6a\xb8\xec\x2a\x28\x00\x00\x00")

func templatesDockerTemplateBytes() ([]byte, error) {
	return bindataRead(
		_templatesDockerTemplate,
		"templates/docker-template",
	)
}

func templatesDockerTemplate() (*asset, error) {
	bytes, err := templatesDockerTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/docker-template", size: 40, mode: os.FileMode(420), modTime: time.Unix(1563422270, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJenkinsTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x95\x51\x6f\xe2\x38\x10\xc7\xdf\x2b\xf5\x3b\xcc\x72\x68\x43\x25\x1c\xb6\xfb\x18\xc4\x43\xd5\x63\x4f\x48\xdb\x2e\x2a\xed\xbd\x5c\xef\x22\x93\x0c\x89\xaf\x8e\x9d\x1b\x3b\x39\x21\xca\x77\x5f\x39\x4d\x9a\x00\xed\x96\x3e\x94\x97\xc4\x33\xe3\xc9\x7f\xc6\xbf\x31\xbf\x7d\x1a\x15\x86\x46\x4b\xa1\x46\xa8\x4a\x48\x48\xeb\x72\x7d\x7a\x72\x7a\x92\x8b\x1c\xa5\x50\x08\x9b\xd3\x13\x00\x9e\xa0\xb2\xb0\x01\xc9\x97\x28\xc1\xcb\xb8\xb1\x48\x1e\x6c\x9d\x0f\x55\x29\x48\xab\xac\x8a\x70\x06\x00\xcb\x93\xb0\x44\x82\x09\x98\x14\x06\x4f\x36\x00\x13\x91\xc8\x6d\x00\xbd\x44\x58\x17\x02\x4c\x02\x63\x46\x93\x9d\x94\x01\xe1\x4a\xf1\x0c\xe1\x11\x2c\x17\x12\xd8\x79\x6f\xd8\xec\x23\xb4\x05\xa9\x85\x8d\x75\x61\x03\xb0\x54\x60\xe3\x39\xf3\x2d\x89\x6c\x70\xd6\xac\x23\x59\x38\x61\x61\x95\x69\x02\xbd\xfe\x26\xe7\xc4\x33\xe3\x77\x1d\xdb\x5e\x13\x9e\x93\xfe\x17\x23\x7b\x18\xde\x75\xb4\xe1\x2b\x21\x31\x14\xf1\x4e\x64\x6d\x6b\x83\x62\x6e\xd2\x26\x61\xb7\xf8\xb6\x7a\xb1\x82\x87\x62\x89\x91\x95\x90\xa0\x85\x52\x90\x2d\xb8\x34\x48\xa5\x88\xd0\xf8\x0a\xed\xff\x9a\x1e\x84\x4a\x7c\x61\xac\xd0\xbe\xd0\xc0\x14\x84\xd7\x17\x57\xd3\xc5\xfc\xe2\x72\x5a\xbd\x01\x63\x75\x4d\x90\x3c\x60\xd8\xdf\xec\x48\x0e\xb9\x11\x9c\x19\x5d\xd8\x14\xb9\xb1\xe7\x8c\x87\xfd\xcd\x4e\x0f\xe0\x11\x12\xc2\x1c\xd8\x7f\x10\xce\xae\x2e\xfe\xa8\xb3\x7e\xfe\xbc\xa3\xed\x03\xbe\x7b\x4c\xb9\x1d\x45\xac\x34\xc0\x34\xac\x79\x26\x3b\x92\x97\xb2\x40\x18\x83\x4d\x51\x01\x46\xa9\x06\x2f\x21\x44\xe5\x8d\x01\xa5\xc1\xda\xe4\x82\xbc\x31\xac\x44\x4b\xd2\xd1\x28\x6d\x9f\x1e\xc6\xf2\x04\x4d\x43\x75\xbd\x1e\x78\x52\xf3\xd8\x61\xef\x9d\xb5\x2e\xe7\xc4\xdc\x74\x0d\x0e\x48\xad\x56\x22\xf9\x26\x24\xce\x49\x97\x22\x46\x1a\xfc\xd5\xda\x06\x8e\x9e\x59\x1c\x38\x9e\x9e\x41\x1a\x42\xc9\x49\xf0\xa5\xc4\x00\xbc\xe9\xf5\x9f\xe1\xb7\xd9\xf7\xa9\x77\xf6\xf7\xd9\x5e\xea\xaa\xc8\x1e\x4c\xda\x1f\xfc\xd3\xf9\xc1\x0d\xf2\x58\xa8\xa4\x96\x00\x2b\xd2\x19\x3c\x0f\xf5\x13\x8c\xd0\xdb\xc9\x68\x52\xe8\x51\x06\x8c\x56\xe0\xa3\x2a\x0f\x9d\x51\x0e\xfd\x0d\xaa\xd2\x6f\x64\x6d\x0f\x03\xb7\xed\xe2\xf9\x75\xbb\xd7\xbf\x65\x21\x64\xfc\x66\xf3\xdc\x17\x0d\xc6\xc0\x04\xdc\x7b\x26\x90\xdc\xa2\xb1\x41\x7f\x53\x5f\x2e\xdb\xe0\xde\x83\xf0\xf2\xc7\xd5\xfc\xc7\x62\x5a\xa9\xa9\x80\xe9\xbd\x37\xc5\x41\x05\x6e\x53\xac\xa3\x07\x24\x16\xe9\x2c\xd7\x06\x81\xad\x5e\xf8\x12\x54\x65\xbc\xb2\x15\x44\xe6\xd0\xe9\x1d\xd1\x8d\x66\x9e\x72\xc2\x9c\x13\x1e\xd5\x97\x66\x48\x23\x42\x6e\xf1\x23\xe6\xd4\x60\x44\x68\x21\x41\x85\x24\xa2\x9d\x91\xac\x5d\x8c\x39\xa6\x18\xaa\x92\x39\x78\x27\xae\x91\xc0\x58\x4c\x6b\x46\x85\xea\x4c\x6d\x23\x96\xe7\xb9\x5c\x7f\x84\x56\xb6\x02\xf6\xfa\xb9\x07\x7c\x7c\x3d\xbe\xef\x7f\x5a\xf2\xb1\x19\xb9\x0d\x41\xb7\x9a\x43\x0b\xeb\x6f\x9e\x2f\xf1\xed\xe8\xab\xa3\xac\x75\xfa\xae\xa6\x5f\x30\x36\x2c\x91\x8c\xd0\x2a\x80\x27\xd8\xda\x75\x37\xeb\x30\x79\x67\xd6\x4e\x6c\x4d\xf1\x8e\xa9\x25\xfa\xad\xcc\xaf\x52\x98\x17\x26\x7d\x13\x3d\x42\x4b\xeb\xc1\xf9\x97\xfd\xbb\xa8\x92\x9b\x44\x52\x17\x31\xd4\xfc\x33\x06\x2e\x25\x24\x11\xf9\x42\x8f\xf6\x0e\x75\xf4\xb2\xfa\xd7\x6e\x12\x00\x2b\x32\xd4\x85\x1d\xb8\x67\x00\x5f\xbf\x0c\xa1\x50\xc2\x06\xe0\x5d\xcd\xae\xef\x6e\xa7\x0b\xef\x25\x49\x5e\x75\x47\x46\x29\x46\xee\xff\x05\x52\xe4\xd2\xa6\xbe\xef\x7b\xef\xbd\xb0\x62\xcc\xa5\x5e\x1f\x35\x99\x51\x41\x12\xc2\xbb\xf9\xef\x17\xb7\xd3\x70\x71\x79\x33\x9b\xdf\x86\x77\x37\xdf\xe1\x11\x96\xdc\xa4\xc0\x8c\xeb\xcc\xfe\x00\x74\xd9\xf8\xc5\x61\x55\x8f\xed\xe9\xc9\xcf\x00\x00\x00\xff\xff\x85\x7d\x61\xe7\xb5\x09\x00\x00")

func templatesJenkinsTemplateBytes() ([]byte, error) {
	return bindataRead(
		_templatesJenkinsTemplate,
		"templates/jenkins-template",
	)
}

func templatesJenkinsTemplate() (*asset, error) {
	bytes, err := templatesJenkinsTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jenkins-template", size: 2485, mode: os.FileMode(420), modTime: time.Unix(1563422270, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesKubernetesTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\xc1\x8e\xda\x30\x10\xbd\xe7\x2b\x46\x5b\xf5\x98\xb2\x5c\x7d\xa3\x2c\xad\xa8\xba\x4b\x04\xb4\xea\x0d\x0d\x66\x36\x6b\xd5\xb1\xad\xf1\x84\x6e\xb4\xda\x7f\xaf\x9c\x10\xc8\x16\xa8\x54\xa9\x73\x82\xf1\x9b\xf7\x5e\x9e\x3d\x79\x9e\x67\xef\x60\xe5\x6b\xd6\xa4\xa0\x24\x47\x8c\x42\x3b\xd8\x36\xe9\x0f\xa3\x78\xce\x30\x98\xef\xc4\xd1\x78\xa7\x60\x3f\xce\x7e\x1a\xb7\x53\xb0\x22\xde\x1b\x4d\x59\x45\x82\x3b\x14\x54\x19\x80\x66\x42\x31\xde\xad\x4d\x45\x51\xb0\x0a\x0a\x5c\x6d\x6d\x06\x60\x71\x4b\x36\x26\x0c\x00\x86\xa0\x60\x33\xbf\x9f\x7c\x9e\x6d\x1e\x26\xf7\xb3\x0c\xc0\x61\x45\x6f\x7b\x31\x90\x4e\x70\x69\x02\x29\x98\xda\x3a\x0a\xf1\xbc\xc8\x00\x82\x67\x39\x30\xe5\x87\xc1\x27\x91\xd0\x36\xba\x53\x05\x9b\xd9\x8f\x62\xb1\x9a\xdd\x6d\x8a\xc5\x72\xdd\x9f\xb0\x17\xaf\xbd\x55\xb0\x9e\x16\x87\x9e\x20\x97\x24\xc5\xc5\x99\x48\x96\xb4\x78\xbe\x62\x3a\x0a\x4a\xdd\xfa\xb0\x1e\x77\x1f\xd1\xa2\xd3\xc4\x0a\x5e\x5e\xb3\x7f\xcd\x94\x9e\x85\x5c\xfa\x19\x47\xfb\xf1\x96\x04\xfb\x8c\xef\x28\x58\xdf\x54\xe4\xe4\xff\xc4\x0c\xb0\xef\x35\x2d\x0a\x45\xf9\x6b\xf4\x4c\xc1\x1a\x8d\x51\xc1\xf8\x2c\x8e\x0a\x45\x3f\x7d\x1d\xa8\x5d\xd1\xbb\xa4\x18\x25\xc5\x51\x36\xdd\x60\x77\xc1\x4b\x6f\xad\x71\xe5\xb7\xb0\x43\xa1\xb6\xcf\xc3\x4e\xaf\x51\xe1\xf3\xaa\xe6\x92\x14\x8c\x6f\x6f\xdf\xa7\xe7\x41\x55\xb0\x47\xc0\x30\xa4\x54\xf6\x8d\xc1\xab\x16\x2f\x99\x04\xe8\x63\x48\xa5\xbd\x13\x34\x8e\x78\x40\x96\x83\xa9\x30\x59\xb9\x29\x35\x7f\x30\x7e\xb4\x29\x96\x8b\x2f\xb3\xe9\xba\xa5\x1e\x0d\x64\x54\xc7\x7a\x73\x1c\x85\x6e\xb4\xa8\xad\x2d\xbc\x35\xba\x51\x30\x7f\x7c\xf0\x52\x30\xc5\x74\xd9\x27\xdc\xf9\xed\x9c\xce\x06\x8b\x70\xf2\x74\x74\x7a\xf9\x51\x0f\xeb\xd2\x42\xa4\x22\xb7\xff\xc4\xbe\x1a\x32\xe7\x10\x49\x33\xc9\x92\x1e\xd5\x1f\x2c\x67\x0e\xf3\x0e\x0a\x03\x1c\x53\x6c\x97\x21\xb6\xeb\x71\xec\x09\xb2\xf4\xdf\x3f\xb1\xbf\xb0\x89\xfd\x56\xc1\xcb\xeb\xef\x00\x00\x00\xff\xff\xb3\x61\x03\x64\x99\x04\x00\x00")

func templatesKubernetesTemplateBytes() ([]byte, error) {
	return bindataRead(
		_templatesKubernetesTemplate,
		"templates/kubernetes-template",
	)
}

func templatesKubernetesTemplate() (*asset, error) {
	bytes, err := templatesKubernetesTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/kubernetes-template", size: 1177, mode: os.FileMode(420), modTime: time.Unix(1563423581, 0)}
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
	"templates/compose-dev-template": templatesComposeDevTemplate,
	"templates/compose-template": templatesComposeTemplate,
	"templates/docker-template": templatesDockerTemplate,
	"templates/jenkins-template": templatesJenkinsTemplate,
	"templates/kubernetes-template": templatesKubernetesTemplate,
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
		"compose-dev-template": &bintree{templatesComposeDevTemplate, map[string]*bintree{}},
		"compose-template": &bintree{templatesComposeTemplate, map[string]*bintree{}},
		"docker-template": &bintree{templatesDockerTemplate, map[string]*bintree{}},
		"jenkins-template": &bintree{templatesJenkinsTemplate, map[string]*bintree{}},
		"kubernetes-template": &bintree{templatesKubernetesTemplate, map[string]*bintree{}},
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

