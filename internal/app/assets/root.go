package assets

import "embed"

//go:embed favicon.ico
var FaviconFS embed.FS

//go:embed index.html
var IndexFS embed.FS
