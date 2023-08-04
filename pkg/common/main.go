package common

import ( 
 	"os"
    "log"
    "fmt"
    "io"
    "net/http"
    "encoding/base64"
    "encoding/hex"
    "crypto/md5"
)

func host() string {
    hostname, err := os.Hostname()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return hostname
}

func GetIP() string {
    req, err := http.Get("https://ipinfo.io")
    if err != nil {
        log.Fatal(err)
    }
    content, err := io.ReadAll(req.Body)
    req.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}

func md5(a string) string {
    hash := md5.new()
    hash.Write([byte](a))
    return hex.EncodeToString(hash.Sum(nil))
}

func b64Encode(a string) string {
    return base64.StrEncoding.EncodeToString([]byte(a))
}