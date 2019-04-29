package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsd610aaa03feef43450db1d0b71a3467f636005b7 = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>form</title>\n</head>\n<body>\n    <form action=\"/\" method=\"post\">\n        <p>Check some colors</p>\n        <label for=\"red\">Red</label>\n        <input type=\"checkbox\" name=\"colors[]\" value=\"red\" id=\"red\"/>\n        <label for=\"green\">Green</label>\n        <input type=\"checkbox\" name=\"colors[]\" value=\"green\" id=\"green\"/>\n        <label for=\"blue\">Blue</label>\n        <input type=\"checkbox\" name=\"colors[]\" value=\"blue\" id=\"blue\"/>\n        <input type=\"submit\"/>\n    </form>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"views"}, "/views": []string{"form.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1556512015, 1556512015244607436),
		Data:     nil,
	}, "/views": &assets.File{
		Path:     "/views",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1556508635, 1556508635076605994),
		Data:     nil,
	}, "/views/form.html": &assets.File{
		Path:     "/views/form.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1556507119, 1556507119000000000),
		Data:     []byte(_Assetsd610aaa03feef43450db1d0b71a3467f636005b7),
	}}, "")
