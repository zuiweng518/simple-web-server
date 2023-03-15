package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "net/http"
  "os"
  "bufio"
  "fmt"
  "strconv"
)

// some change - 2

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  // Start server
  e.Logger.Fatal(e.Start(":1325"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World! -- 17")
}
func  CreatePidFile(){
    filePath := "/opt/simple-web/pid"
    file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("文件打开失败", err)
    }
    //及时关闭file句柄
    defer file.Close()
    //写入文件时，使用带缓存的 *Writer
    write := bufio.NewWriter(file)
  write.WriteString(strconv.Itoa(os.Getppid()))
    //Flush将缓存的文件真正写入到文件中
    write.Flush()
}
