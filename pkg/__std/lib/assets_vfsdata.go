// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package lib

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 11, 12, 0, 10, 3, 567165603, time.UTC),
		},
		"/std.js": &vfsgen۰CompressedFileInfo{
			name:             "std.js",
			modTime:          time.Date(2018, 11, 12, 22, 32, 15, 982381054, time.UTC),
			uncompressedSize: 13357,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5b\x5d\x6f\xe3\x36\xb3\xbe\x3f\xbf\x42\xc9\x45\x20\x56\xb2\x20\x29\x8e\x37\x6b\x99\x5e\xec\xb6\x49\x9b\x62\x9b\x14\x9b\xb4\x45\x1b\x04\x86\x62\x53\x31\x4f\xbd\x92\x4b\xd2\x9b\x93\xda\x3e\xbf\xfd\x05\x3f\x24\x91\x12\x65\x3b\xc5\xbe\x40\x6f\x9a\x5a\xe4\x3c\x33\x9c\x0f\x72\x66\xc8\xfd\x92\x12\x87\x41\xb6\xd9\xac\xb7\x09\x0b\x2e\x0b\xf2\x39\x65\x70\xfd\x7e\xc5\x8a\x61\xe8\xff\x78\x7b\x73\x3d\x8c\xfc\xdf\xdf\xff\xf4\x71\x18\xf3\xf1\xdf\x08\x66\xe8\x3d\x79\xa2\x30\x5b\xe5\x53\x86\x8b\xdc\x05\x6b\x36\xc7\x34\x78\x7c\x84\xf9\x6a\xb1\x48\xd4\x8f\xc9\xb2\xa0\x30\x34\x48\x82\x25\x29\x58\xc1\x5e\x96\x28\x98\x4c\x70\x8e\x59\x8d\xc1\x7c\x54\xc1\x08\x4a\x56\xe2\x40\x94\x10\xc4\x56\x24\x77\xf8\x07\x13\xef\x09\xb1\x4f\x45\xc1\xde\x53\x8b\x54\xc8\xc7\x60\x2d\x29\x5d\xbc\xd9\xe4\xe8\xd9\xd1\x48\x81\x92\xc0\x45\x01\x41\xe9\xec\x2a\x67\xa7\xb1\x8b\x82\x65\x41\xb1\x5c\x13\xf0\xf4\x5f\x3e\x02\x5d\x2b\x59\xa6\x6c\xae\xad\x03\xac\xb9\x3e\x11\x54\xd2\x07\x93\x49\x91\x65\x14\x31\x57\x5b\x9c\xdf\x07\xe5\x9a\xd0\xbb\x7a\x22\x65\x04\xe7\x4f\xfa\x44\x0f\xf9\x0c\x0c\xb9\x56\xbb\xb8\x7f\x49\x17\x2b\xf4\x5a\xf6\x83\xaf\xc6\x9e\xff\x47\x77\x84\x03\x98\x9f\x5b\x98\x2b\x13\x9c\x9b\xcc\xc1\xb0\x74\xc7\x80\x7b\xa3\x29\x03\x65\x29\x61\x16\xb3\x33\xb0\x66\x72\xf0\xe6\xf1\x7f\xd1\x94\xb9\xa7\x0d\xcb\xa5\xb3\xd9\xcf\xa6\xc9\x84\xeb\xf1\xef\x97\x18\x2d\x66\x37\x52\xe4\xd0\x47\x7e\xd8\x26\xfd\xb5\xa1\x6f\x2b\x6d\x64\xa7\xbd\x33\x94\x25\xfc\x13\x55\xa4\x62\xf9\xb1\x8f\x7d\x63\xcd\x0d\x10\x94\xcf\xec\x4b\x56\x7a\xe7\x13\xd4\xb2\x6b\x2d\x6f\x13\x39\x8a\x44\x84\xa3\x40\x10\xaf\xaf\x6f\xae\x2f\x86\xa1\x5f\xc1\x0d\x23\x3e\xf6\x13\xa2\x34\x7d\x42\x87\xc6\x76\x45\xf0\x55\x22\xbb\x46\xab\xe2\xba\x25\x0f\x6b\x47\x75\x45\x56\xc5\x34\xd3\x62\x9a\x19\x31\xad\xff\xf2\x19\xb0\xaf\x20\x25\x4f\xf4\xce\xe2\xd8\xec\xd0\xa0\x66\x86\x63\xff\x82\x9b\x9e\xcd\xc0\x50\x9a\x21\xe0\x56\xe8\x16\xe2\xeb\xc4\xf5\x2a\x97\x8a\x6b\xc4\x96\x0c\xeb\x9a\xb5\x88\x99\xb6\xbe\x9b\xe1\x14\x1b\x4a\x4b\x67\xb3\xf7\x2d\x6d\x09\x1b\x31\xd3\xb3\x43\x1f\xfb\xda\x9a\x6d\x20\x87\x46\x55\x4d\x88\xf2\x99\x4d\xe0\xfd\xc1\x80\xe1\x7a\x9b\xe0\xe0\xf6\xea\x8f\x8b\x9b\xcb\xc9\xed\x0f\x37\x9f\xee\x60\x5c\x7f\xb8\xba\xbe\x83\xfd\x04\x07\x97\x57\x1f\x2f\x26\x57\xdf\x5d\x5c\xdf\x5d\x5d\x5e\x5d\x7c\x9a\x7c\xbc\xb8\xfe\xfe\xee\x07\x31\x74\x91\x4f\x8b\x19\xce\x9f\xe0\xfa\x97\xbb\xcb\xf3\xc9\x87\xdf\xef\x2e\x6e\x87\x91\xff\xcb\xdd\x65\x34\x98\xdc\xde\x7d\xba\xba\xfe\x9e\x9f\x97\x38\xc0\xdc\x0b\x21\xf7\x53\xe1\x8f\xef\x09\x49\x5f\xdc\x18\x24\x38\xc8\x16\x45\x5a\x8e\x5d\xca\xff\x97\xa3\x8a\x28\x78\x5c\x65\x19\x22\xd5\xd4\x41\xbf\x9e\x3a\xe8\x77\x4d\xc5\xf4\x23\x66\x6c\x81\x2e\xf2\x19\x4e\x73\x41\xc1\x1d\x30\x1a\x48\x82\xf2\xf7\xb9\xfc\x79\x1f\xf9\xe1\x03\x28\xc9\xef\xc3\x07\x08\x61\x94\xe0\xe0\x63\x91\x3f\xd9\x42\x78\x51\x3c\x43\xb6\x09\x65\x04\xcf\xf1\xd3\x1c\xa2\x4d\xb8\x55\x04\xc1\x94\xa0\x94\x35\x77\xc7\x32\x26\x20\x0c\x4f\x4e\x10\x84\xe1\x3b\x35\xfb\x8f\x8b\x4f\x37\x43\x2e\x8f\xfc\x2d\x66\x57\x50\xda\x11\x53\xa8\x15\xeb\xe1\xa8\x76\x80\x52\xa6\xf1\x78\x1c\x02\xaf\x12\xea\x9b\x7e\xfc\xb6\xff\x76\xf0\x26\x7e\x3b\xb0\x00\xa2\xbf\x56\xe9\xc2\x8c\x2d\x6d\x1b\x12\x4b\x84\x8c\xff\x39\x39\xa9\x97\x09\x99\xf8\x5b\xc1\x71\xe1\xa1\x26\x7c\xe8\x87\x5c\xfb\x1f\x56\x78\x31\x43\xc4\x00\xc7\x99\x7b\x54\x79\x65\x14\xc6\xfd\x2d\x5a\x50\x54\x7a\xe9\xb6\xdc\x0c\x71\xf0\xe1\x85\xa1\x0f\xc2\x12\x41\xba\x58\x14\xd3\x94\x21\x17\x01\xa9\x6b\xba\x4c\xa7\x08\x22\xf9\xe3\x33\xce\xd3\x05\x7e\xca\x61\x24\x7f\x7f\x61\xe9\xe3\x02\x69\x3b\xb4\xfc\x30\xc1\xf9\x64\x45\x11\x54\xe6\xc2\xf4\x1a\x51\x86\x66\x30\x4b\x17\x54\x21\x15\x22\x3e\x26\x22\xc0\xcb\x79\x92\x98\xc2\xfb\x07\xf5\x1b\x4d\x59\x41\x26\xf9\xea\xf3\x04\x2d\xd0\x67\x5a\xce\xcb\x0a\x32\x45\x93\x19\xca\xd2\xd5\x82\x51\x89\xba\xad\x75\xa0\x69\x5c\xcc\xfc\xae\x9a\xa8\xef\x2d\x16\x20\x66\x07\x99\xa5\x2c\x95\xda\x69\x3b\x82\xa3\x94\x68\xa7\x4c\x69\xed\xf1\xdd\xb4\xc1\xe3\x0b\x43\xd4\x05\x01\x5d\x3d\xa6\x22\x38\xca\x01\xfd\xc8\x68\x7d\x92\x5e\xa7\x36\x63\x00\xba\x24\x10\xcc\xf7\x89\x5f\x89\x20\xe3\x31\xa0\x0b\x3c\x45\x5f\x51\x8c\x25\x41\xcb\x46\x74\xe2\xcc\x65\x63\xc3\xa9\x94\x51\x2a\x1f\x63\x5b\xee\xa9\x04\xfe\x7f\x25\xc8\x34\x5d\xa6\x53\xcc\x5e\x5c\xd0\xab\x7d\xd3\x43\xc0\x8b\x4e\x58\x2f\x4a\x9e\xe7\x78\xa1\xa4\x16\x23\x23\xe2\x31\x0f\xc9\x00\xa0\xb0\x0d\x92\x68\x11\xa0\x44\x7e\x22\xc5\x73\x1d\x0d\x25\x63\x3d\x14\x3c\x0b\x50\x8f\xca\x60\x5a\xa6\x33\x97\x74\xa9\x20\x9d\x19\x0e\x98\x15\xc4\x95\x91\x18\x26\x68\xc4\x12\xe4\x79\x55\xb6\x12\x3c\xf3\xe4\x48\x1c\x5f\x3d\x6d\xa5\xfc\x10\xb2\x83\x57\xf3\xdb\x3e\x6e\xa0\xd5\x58\x3d\x18\x89\x34\x64\x17\x5a\x34\xd8\x09\x17\x0d\x0c\xbc\x78\x2f\xde\x69\xbc\x13\x8f\xe7\x4c\x1a\x5e\x7f\x2f\x9e\xbe\x33\x5b\xf0\x06\x7d\x03\xef\x7c\x37\x9e\x3a\x08\xbb\x11\xd5\x84\xd7\xc8\xd8\x3a\x3f\xac\x98\x07\xcb\x99\xce\x66\x76\x1b\xf3\xf0\x72\x23\x7e\x10\x88\x9f\x9a\xbd\x77\x22\xd9\xec\x2b\xa0\xe2\x16\x14\xb7\xf5\x4e\x2c\x9b\xe6\x04\x56\xbf\x85\xc5\x75\xb8\x13\xcb\xa6\x31\x81\x75\xde\xc2\xe2\xba\xeb\xc6\xea\xb4\xa9\x4d\xb2\xca\xbe\x7b\xf0\x0e\x95\xae\xb2\xed\x0e\xbc\x32\x4b\x35\xf6\x46\x9e\xc4\xf2\xdd\xb1\x7d\x42\x6d\x36\xe8\x08\x62\xc5\x53\x79\x43\x7d\x50\x2f\x0a\xc6\x99\xed\xe5\x66\x98\xfd\x95\xec\xa2\xc1\xeb\xf9\x19\x06\x78\x25\xbf\xd3\xf8\xf5\xfc\x0c\x03\xed\xe6\x77\x54\x26\x65\x2e\x06\x06\xdf\x41\xff\x95\x7c\xdb\xae\xf6\x9a\x95\x96\xbe\xf7\x0f\x78\xbe\x62\xb5\x16\x9e\xaf\x5e\xa7\x2c\x89\xfe\x21\x4b\x55\x4f\xbd\x8e\xe3\x2d\x23\xab\xa9\x95\xa3\x06\x9e\x8b\x04\xf3\x40\xe4\x5c\x65\xa3\x66\xaa\xcc\x8e\xa0\x99\xca\xac\xd9\x9c\x14\xcf\x0e\xcf\xb5\x2f\x08\x29\x88\x7b\x7c\xb9\x48\x99\xcc\x0c\xe8\xd0\xa1\x42\x2e\xe7\xf3\x8a\x32\xe7\x11\x39\x14\x11\x9c\x2e\xf0\xdf\x68\xe6\xe0\x7c\x81\x73\x14\x1c\x77\xf2\x2f\xd8\x75\x43\x84\x5a\x83\x65\xb2\xbc\x8f\xbd\x4c\xa0\x2b\xb6\x29\x87\x91\xc2\xe4\x85\x10\x48\xae\xb2\x5b\x0a\xae\xa3\xf6\x4e\x26\x73\xf0\x7b\xf6\x60\x2a\xc3\x0e\x51\x34\x9c\xa1\x95\x56\x5a\x13\x36\x1d\xcb\xcc\xb7\xac\xc5\xb4\x96\xaf\x71\x9b\x9f\x9c\xc6\x71\x14\xc7\x67\xfd\x37\xf1\x3e\x15\x4d\xd3\x9c\xeb\x82\xf3\x70\x64\x62\xeb\x3c\xa2\x97\x22\x9f\x39\xb1\xf3\x84\x9f\x52\x91\xf4\x72\x05\xc9\x3c\x13\x8d\x46\x51\x22\x73\x45\x7b\x49\x44\x40\x42\x03\x8a\xd8\xcf\x65\xf2\x4b\x7a\x88\x7f\xaa\xf2\x77\xc4\x5c\x56\xfe\xf2\xc5\xa0\xd2\x07\xed\x74\xf1\x56\x3c\xe9\x27\x8a\xde\x17\xb0\x1e\xa3\xba\x89\x7a\xcc\xd3\xe7\x77\x98\x4c\xeb\xa8\xb4\x99\x56\x8e\x29\x75\xad\x97\x78\xa2\xc6\x33\x7c\x04\xde\x3f\x6c\x2d\x35\x1f\x4b\xba\x33\x5b\xe5\x5c\xe8\x01\x86\x5b\xb3\x32\x64\x64\x65\x2b\x0c\x0f\xf0\xc1\xaa\xd7\x62\x0b\x26\x5d\xf8\xcd\xe6\xe8\x55\x01\x56\x01\x3b\xd3\x74\xb1\x40\x33\xe7\x19\xb3\x79\xb1\x62\x8e\xa6\xc2\x63\xb0\x35\x0f\xab\x10\x24\x5a\xc3\xae\x94\x3c\xd1\x5a\x68\x86\xb6\x7a\x91\xd0\x56\x82\xc6\x30\x54\x45\x7f\xad\x23\xae\xbe\x5e\x0f\xac\x4b\xf7\xf4\xb4\xc9\x72\xc4\x3c\x98\x4d\xf2\x23\x18\xbe\x63\x3d\xf3\xdb\x30\x94\xbe\x4e\x61\x9c\x34\x68\x7b\x2d\xe5\x4b\xa9\x73\xe8\x12\x8f\x82\x6f\xcc\x9e\x55\x83\x3a\x97\x73\x0b\x18\x8a\xbf\x19\xac\x83\x3d\x61\x43\x2e\xb4\xf2\x05\xad\xca\x0f\x16\x28\x7f\x62\x73\xe9\x1d\x9c\x2a\xb5\x15\x56\x3a\xc5\x3d\x7a\xe0\x5e\x99\x43\xa8\x37\x39\xa5\x04\x29\xa8\x2b\xaa\x39\x6c\x08\x3b\x1f\xe5\xc9\xdc\x6b\x7c\xad\x9d\xc4\x00\xca\xbc\x39\x38\xb2\x31\xf0\xe6\x00\xac\xa7\x45\xce\x70\xbe\x42\x0e\xdb\x6e\x0b\xd8\x14\xee\x91\xa0\xf4\xcf\xed\x16\x67\x6e\xa1\x6c\x23\x9b\x27\xb6\x75\x25\x3b\xab\x1f\xbf\xe8\x31\x20\xbb\x35\x86\xce\x96\x2b\x3a\x37\xa3\x1e\x74\x02\x35\x18\xfa\x8d\xcd\x02\x6c\x6d\xbd\x99\x72\xff\xee\x68\xa7\xe0\x1c\xd3\xe6\x05\x06\xdf\x94\xa5\x05\x49\xd9\x26\x12\x9b\x97\x51\xcc\xfb\xfa\xd6\xe4\x75\x35\x38\xc5\xae\x43\x94\x67\x1c\xc1\xce\x69\x7b\x22\x37\xc3\x0b\xe4\xe0\x19\xca\x19\xce\x30\x22\xd5\x11\x2d\x71\x9d\xe3\x6e\xfe\xdb\xd2\x87\x68\x27\xf3\x5e\x94\x50\x1e\x80\xb4\x0a\xc0\xba\xdc\x22\xc1\x74\x9e\x92\x6f\x8b\x19\x7a\xcf\x5c\x0a\xc0\x76\x7b\x88\x3a\x40\xd2\xc8\x8f\x58\x6d\x54\xfd\xb4\xa9\xdd\xa3\x63\x2b\x24\xe8\xaf\x15\x26\x48\xa6\x4d\x0d\x2b\xc9\xe6\xb3\xd5\x13\xa5\xe9\x70\xaf\xe1\xf4\xa7\xb1\x8b\x41\x62\xf6\x4e\xea\x78\x20\x1e\x02\x47\x30\xe4\x06\x3b\xa2\xfb\x0d\x82\x16\x33\xe7\xd8\x43\xde\xb1\x96\x2f\xb1\xee\xd4\x84\xef\x3f\xbf\x8a\xee\x9f\x99\xf7\x11\xcb\x41\x25\x03\xa4\xd9\x2a\xd4\x3d\xd1\x38\x46\xd9\x37\x65\x92\x28\xc6\x88\xf8\xd0\x79\xb4\x34\x85\x68\x98\xbc\x8c\xb5\x26\x7b\xa0\xdf\x2f\xed\x39\xc0\x64\xef\xfa\x56\xdc\x81\xb6\x72\x52\x07\xe7\x94\xa5\xf9\x14\x15\x99\xd6\x3d\xaf\x92\x23\xbd\xa1\x7b\xff\xa0\x6e\x18\x42\xd5\x08\xc3\x23\xa6\x42\x49\x85\x67\x69\x4c\xdd\x4b\xb1\xe7\x89\xa8\xa3\xa3\xb3\xb3\xf8\xed\x60\xb3\xa1\x63\x78\x36\x38\x8d\x43\xb0\x26\x90\xd6\xf8\xb9\x85\x8c\x40\x97\x8e\x46\x51\x08\xbc\xdc\x73\x07\x67\x67\xa7\x83\x9e\x2b\x50\xc4\xc7\x9e\x84\xe1\xdb\x21\x19\x45\xf1\x39\x58\x23\xb9\x7f\x11\xb5\xb3\x89\x81\x38\xec\x6b\x23\xe3\xf1\xe0\xe4\x34\xda\x44\x6f\x63\x7d\x8e\x80\xd6\x27\x45\xf1\x49\x74\xb6\x89\xe3\xbe\x9a\xa5\x8d\x9c\x9f\xbc\xd9\xc4\xfd\xd0\x97\xb3\x06\xa7\x1b\xce\x79\xab\xe3\x37\xbf\x55\x1f\xb6\x5b\xa3\xae\x2e\xf3\x2d\xcd\x19\xdd\xc8\x47\x4a\xa5\x7e\xb4\x2f\x48\x7b\xb0\x9c\x0b\xaa\x94\x08\xc3\xd0\x2f\xb4\xd3\xd1\xcf\xaa\xd0\x52\x69\x63\x82\x47\x25\x59\xc2\x95\xbc\xce\xee\x0b\xcf\x7b\x80\xe8\x1e\x3f\x6c\x75\xa7\xaa\x9c\x73\xb7\x5f\x59\xee\x53\x14\x8a\x71\x75\x52\xdd\x81\x74\x24\xe3\x52\x4a\x2e\xe2\xa4\xbc\x46\x2d\x5b\xc0\x13\x18\x9a\x84\x55\xc2\x6c\xbb\xe9\x90\x37\x17\x5a\x8b\xb5\x71\x2f\xc4\x40\x43\x0c\x6d\x45\x82\x7f\x67\xb5\x21\x84\xeb\xa4\x2d\xa5\xed\x22\xaf\x56\xd3\x89\xa0\x59\xd9\x92\xae\x57\xca\x60\x9d\x00\xe5\xa6\xbb\x7b\x05\xca\xf8\x9d\x28\xe5\x8b\x85\xce\x6b\x24\xed\xe6\x17\x8c\x46\x71\x7f\x3c\x8e\xfb\x3b\xd1\xc4\xe4\x4e\x38\x29\xd5\x3d\x7b\xd8\x27\x51\xa3\x9b\x68\x13\x49\x74\x11\x47\xa3\x68\x30\x1e\x47\x83\xbd\x32\xed\x00\xac\x84\xda\x18\x3f\xbd\xe8\x61\x34\x3a\xdf\x27\x68\xa3\x21\x78\x30\x6e\xe3\x5b\xfc\xc0\x57\xd2\xf8\x78\xfa\xc0\x55\xbe\x77\x69\x3b\x44\xd0\x1e\x10\x80\xf1\x78\xdc\x8c\xad\xd6\x6a\x1a\xed\x48\x23\xcc\xe4\xed\x66\x13\xd5\x6f\x7e\xf1\xfa\x3b\x82\xae\x14\xf9\x70\x3e\x72\x81\x06\xa3\xf2\xd3\x5e\x4e\xb6\x8e\xad\xba\x6a\xbe\x0f\x55\x4f\x42\x5f\x4b\x52\x6d\x66\xea\x52\xfb\x3e\xdc\xed\xa8\xb6\x16\x6e\xc9\xa0\x79\x87\xfd\x2e\x1c\x46\x36\x9e\x9d\xf3\xa3\xa1\x45\x46\xaf\xdf\x94\x72\xd0\xdf\x25\xa5\xed\x2e\xa7\x7e\xc9\x52\x3a\x28\x44\xbb\x01\x9a\x41\xfd\x7a\x84\x76\xa7\xb8\x8d\x90\x34\xc3\x04\xa2\xf1\xb8\x3b\x00\x2b\xc9\xfe\x2b\xc0\xed\x56\xf3\xa1\xb8\x49\x33\xb2\xf9\xd7\x68\x90\x34\x63\x9b\x7f\xde\x11\xdd\xd5\xf2\xfe\x05\x62\xb4\xfb\xe0\x96\xd4\x95\x67\x33\xc5\xb3\xa5\xaf\xe4\xf5\x7d\x24\x9e\x1d\x74\x47\x6b\xb5\xd8\x1d\x6c\xca\xb0\x6f\xf3\xa9\x37\x84\xc3\x18\xd9\xda\xeb\x3c\x70\xeb\xa8\x2f\x95\x6a\xac\xaf\xde\x3a\x0e\xc1\x6f\xad\x44\x8f\xd7\x9d\xf8\xd6\x9d\xc3\xae\xd7\x9d\x9b\x47\xb7\x94\x4f\x48\x95\x53\x57\x55\x5d\x6b\x6b\x78\x19\x39\xc4\xc8\x4c\x4d\xbc\xc3\x6a\xf0\x3d\xb5\x5c\x2d\x9d\x83\xa9\xc3\x8a\xc2\xa1\xf3\x82\x30\x87\x15\xce\xb4\xc8\x59\x8a\x73\x27\xcd\xb5\xea\xbb\xec\xaf\x32\x78\x7c\xdc\x68\x0d\x76\x89\xa0\x3a\x86\x1e\x94\x45\x51\x90\x91\xe2\xf3\xb7\xaa\xf6\x30\x8e\x33\x75\xa1\xdd\xb1\x42\x04\xc0\xd6\xe8\x67\xd8\xf4\x5a\x3e\x67\xb3\x97\xcb\xbd\xce\x03\x47\x35\xb3\xea\x82\x18\x83\x77\xcd\x2f\x1e\x02\xc3\xee\xf3\x5b\xbd\x8f\x6b\x3d\x3e\x53\x8f\x15\x91\xd7\xe0\xcd\xcb\xd6\xe0\xf1\x51\x9c\x2f\xc9\x01\xeb\xa2\xcd\x8a\x52\xc0\x7b\x96\xf3\x4c\x36\x00\xec\xdf\x29\x37\x9b\x2c\x00\xc3\x84\x69\x4d\xb4\xab\xeb\x3b\xd1\x8f\x87\x10\xd6\x8f\xd2\x82\xfa\x4d\x9a\x35\xb7\xad\x5f\xb8\xf8\xcc\x23\x60\x2b\x0b\xd5\x7c\x44\xa4\xc6\x0b\xbd\x7d\xa8\xe5\xb1\x5e\xae\xaa\xd4\x6c\xc4\x4b\xc3\x75\x01\xb3\xba\x32\x4d\x77\x4d\xe7\x35\xe2\xba\x80\x6e\x76\x72\x1a\x81\xd1\x68\xb0\x49\x4f\x06\xa7\x35\xed\x7c\x27\x6d\x3f\x54\xb4\xd1\x19\xcf\x5a\xe3\x8d\xcb\xa9\x05\xcc\xdc\x80\x59\x75\xc0\x08\xe2\x37\x9c\xf6\xbc\xa2\xe5\x30\xf3\x12\x66\xc5\x61\x64\xc7\xb0\xac\x72\xa9\xdd\xed\x0b\x55\xea\x16\x3d\x28\x26\x26\x1d\xf3\xdc\x62\x3c\xe6\x15\xb9\xa8\xc3\x7d\xb7\x38\x71\x23\x59\x8d\x47\xc0\x53\x05\xf9\xd6\xb8\x94\xb0\x3b\x0f\xce\x67\x98\x34\xef\x08\x4a\x83\x36\x3d\x93\x75\x6f\x5b\x93\xc9\x97\x66\x1b\x67\x17\x8e\x11\xc0\x7b\x41\x27\x0b\x94\x1f\x92\x47\xb7\xf9\xec\x12\x78\x9e\xd2\x09\xb6\x6c\xb2\xaa\x23\xf3\x6f\xe8\x4e\x1e\xb2\x7d\x0a\x59\xb5\x8e\x0d\x2a\xdb\xda\x87\xee\x9d\xa5\x36\xe5\xf3\xba\x6a\x27\x25\xab\xee\xbc\xf1\x9f\xb4\x1c\xa6\x45\x4e\x99\x43\xa0\xbc\x50\x09\x32\x82\xd0\xdf\xc8\x2d\xdf\xa2\x83\xa4\x04\x72\xa8\x4b\xfd\x1c\x1e\x1f\xfb\x05\x24\xf2\x89\xfa\x5a\xd2\x66\xea\x39\xa4\x6a\x81\xb8\x51\x18\xf7\x81\xc2\x4d\xe1\x8f\xb7\x37\xd7\x81\xdc\x0b\x71\xf6\xe2\xd2\x72\x64\x0e\x33\xa3\xf5\xe6\xa6\xe5\xc8\xaa\x39\x92\x83\x1d\xff\x00\xc0\xcd\x80\xf5\xa9\xbe\x9b\xf9\xf3\xd6\xc8\xcf\x29\x9b\xbb\x99\xbf\x6a\x0d\xdc\xbd\x2c\x39\x45\x51\xca\xb0\x84\x5d\xaf\xef\x39\x3f\xfb\xdb\x69\x73\x44\x7b\x20\xed\x66\xe5\x0b\xe8\xfa\x5f\xa1\xb4\x67\xba\x99\xbf\x2c\xd9\xcf\xa0\xed\xad\x33\x67\x90\xa9\x4b\x00\x77\x06\x92\x5f\xcf\x7f\x2b\xc8\x9f\x88\xc4\x01\x45\xf9\xcc\xcd\xcc\xd7\x86\x2e\x00\xdb\xca\x78\xca\x0f\xa8\xcb\xfc\xe3\x63\x6e\x77\x79\x5f\xb4\x5e\x14\x4f\xc3\xdc\x97\xb6\x1e\x12\x5f\xe4\x4a\x43\xba\x4d\xd0\xff\x2d\x79\x52\xa1\x1e\x12\x38\x45\xf2\x3f\xff\x09\x00\x00\xff\xff\xb0\xac\xfb\x79\x2d\x34\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/std.js"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}