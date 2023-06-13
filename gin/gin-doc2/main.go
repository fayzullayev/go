package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyBenchLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("My Custom Logger", context.Request.Method, context.Request.URL)
	}
}

func myLoggerOne() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("logger one")
	}
}

func myLoggerThree() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("logger three")
	}
}

func myLoggerTwo() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("logger two")

		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong: myLoggerTwo",
		})
	}
}

func benchEndpoint(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"Message": "Hello",
	})

}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Users struct {
	Users []User `json:"users"`
	Count int    `json:"count"`
}

func authLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.Header.Get("Authorizations") == "" {

			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "access denied",
			})

			return
		}

		return
	}
}

func main() {

	// Disable log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.New()

	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {

		switch recovered.(type) {
		case string:
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", recovered))
		case int:
			c.String(http.StatusInternalServerError, fmt.Sprintf("error  num: %v", recovered))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
		panic("456")
	})

	router.Run(":8088")

	//router := gin.New()
	//
	//gin.Logger()
	//
	//router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))
	//
	//router.Use(gin.Recovery())
	//
	//router.GET("/ping", func(c *gin.Context) {
	//	c.String(http.StatusOK, "pong")
	//})
	//
	//log.Fatal(router.Run(":8088"))

	//// Disable Console Color, you don't need console color when writing the logs to file.
	//gin.DisableConsoleColor()
	//
	//// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//
	//gin.LoggerWithFormatter()
	//
	//// Use the following code if you need to write the logs to file and console at the same time.
	//// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//
	//router := gin.Default()
	//
	//router.GET("/ping", func(c *gin.Context) {
	//	c.String(http.StatusOK, "pong")
	//})
	//
	//router.Run(":8088")

	//r := gin.New()
	//
	//gin.Default()
	//
	//r.Use(gin.Logger())
	//
	//r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
	//	if err, ok := recovered.(string); ok {
	//		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	//	}
	//
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//}))
	//
	//r.GET("/panic", func(c *gin.Context) {
	//	// panic with a string -- the custom middleware could save this to a database or report it to the user
	//	panic("foo")
	//
	//})
	//
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "ohai")
	//})
	//
	////r.GET("/benchmark", myLoggerOne(), myLoggerTwo(), myLoggerThree(), func(context *gin.Context) {
	////	context.JSON(http.StatusOK, gin.H{
	////		"message": "Success",
	////	})
	////})
	////
	////r.GET("/login", func(context *gin.Context) {
	////	context.JSON(http.StatusOK, gin.H{
	////		"message": "welcome to login page",
	////	})
	////})
	////
	////r.POST("/login", func(context *gin.Context) {
	////	context.JSON(http.StatusOK, gin.H{
	////		"message": "logged successfully",
	////	})
	////})
	////
	////auth := r.Group("/api/v1")
	////
	////auth.Use(myLoggerThree(), myLoggerOne())
	////
	////auth.GET("/users", func(context *gin.Context) {
	////
	////	users := Users{
	////		Users: []User{
	////			{
	////				Name: "Mirodil",
	////				Age:  28,
	////			},
	////			{
	////				Name: "Iroda",
	////				Age:  21,
	////			},
	////		},
	////		Count: 2,
	////	}
	////
	////	context.JSON(http.StatusOK, gin.H{
	////		"users":   users,
	////		"cwecwec": "cwecwec",
	////	})
	////})
	//
	//log.Fatal(r.Run(":9999"))

}
