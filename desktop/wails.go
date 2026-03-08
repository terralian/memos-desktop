package desktop

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func GetMemosBackendPost() int {
	return 34561
}

func GetMemosBackendUrl() string {
	return fmt.Sprintf("http://127.0.0.1:%d", GetMemosBackendPost())
}

// WailsServerMiddleware
// see https://v3alpha.wails.io/guides/gin-routing/
func WailsServerMiddleware() assetserver.Middleware {
	return func(next http.Handler) http.Handler {
		target, err := url.Parse(GetMemosBackendUrl())
		if err != nil {
			panic("failed to parse memos backend url: " + err.Error())
		}
		proxy := httputil.NewSingleHostReverseProxy(target)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/wails") {
				next.ServeHTTP(w, r)
				return
			}
			proxy.ServeHTTP(w, r)
		})
	}
}
