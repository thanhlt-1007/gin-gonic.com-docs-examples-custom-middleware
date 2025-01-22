package main

import (
    "time"
    "github.com/gin-gonic/gin"
    "log"
)

func middlewareLogger() gin.HandlerFunc {
    return func(context *gin.Context)  {
        now := time.Now()
        context.Set("example", "12345")

        // before request
        context.Next()

        // after request
        log.Println(time.Since(now))
        log.Println(context.Writer.Status())
    }
}

func getExample(context *gin.Context) {
    example := context.MustGet("example")
    log.Println(example)
}

func main() {
    engine := gin.New()
    engine.Use(middlewareLogger())
    engine.GET("/example", getExample)
    engine.Run()
}
