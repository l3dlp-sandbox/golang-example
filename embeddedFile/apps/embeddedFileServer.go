package apps

import (
	"log"
"net/http"
)

func EmbeddedFileServerMain() {
	// FS() is created by esc and returns a http.Filesystem.
	http.Handle("/", http.FileServer(FS(false)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
