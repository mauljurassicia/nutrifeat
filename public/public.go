package public

import (
	"embed"
	"net/http"
)

//go:embed **
var assetEmbed embed.FS


func GetAssets() http.FileSystem {
	
	return http.FS(assetEmbed)
}
