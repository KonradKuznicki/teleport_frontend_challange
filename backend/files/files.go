package files

import (
	"io"
	"log"
	"net/http"
)

const RESP = `[
    {
        "name": "images",
        "type": "folder",
        "size": 4
    },
    {
        "name": "mountains.jpg",
        "type": "jpg",
        "size": 7340032
    },
    {
        "name": "test.pdf",
        "type": "PDF",
        "size": 12000
    },
    {
        "name": "some_file.pdf",
        "type": "PDF",
        "size": 12000
    },
    {
        "name": "stuff with spaces.pdf",
        "type": "PDF",
        "size": 12000
    },
    {
        "name": "lol.pdf",
        "type": "PDF",
        "size": 12000
    }
]`

func FilesHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("content-type", "application/json")
	_, err := io.WriteString(writer, RESP)
	if err != nil {
		log.Printf("error writeing response for %s error: %v", request.URL.Host, err)
	}
}
