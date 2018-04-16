package main

import (
    "os"
    "github.com/zserge/webview"
    "strconv"
    "fmt"
)

func main() {
    var webServerUrl = NewUrl()
    var apiServerUrl = NewUrl()

    teachApiUrltoJS(apiServerUrl)
    go startHttpServer(webServerUrl)
    go startApiHandler(apiServerUrl)


    if os.Args[1]=="server" {
        fmt.Printf("WebServerPort:%d APIServerPort:%d",webServerUrl.port,apiServerUrl.port)
    }else {
        webview.Open("Idol Produce Framework DevVer",
            "http://localhost:" + strconv.Itoa(webServerUrl.port) , 1280, 720, true)
    }
}

func teachApiUrltoJS(apiUrl *url){
    _, err := os.Stat("page/apiUrl.js")
    var fp *os.File

    if err != nil {
        fp, err = os.Create("page/apiUrl.js")
    } else {
        fp, err = os.Open("page/apiUrl.js")
    }

    b := []byte("var apiUrl = http://localhost:" + strconv.Itoa(apiUrl.port) + ";")
    //書き込み
    fp.Write(b)
    fp.Close()
}
