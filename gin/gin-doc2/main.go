package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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

		context.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})

		context.Abort()
		return
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

			context.JSON(http.StatusOK, gin.H{
				"message": "access denied",
			})

			return
		}

		return
	}
}

func main() {

	r := gin.Default()

	r.Use(myLoggerTwo())

	r.GET("/benchmark", myLoggerOne(), myLoggerThree(), func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Success",
		})
	})

	r.GET("/login", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "welcome to login page",
		})
	})

	r.POST("/login", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "logged successfully",
		})
	})

	auth := r.Group("/api/v1")

	auth.Use(authLogger())

	auth.GET("/users", func(context *gin.Context) {

		users := Users{
			Users: []User{
				{
					Name: "Mirodil",
					Age:  28,
				},
				{
					Name: "Iroda",
					Age:  21,
				},
			},
			Count: 2,
		}

		context.JSON(http.StatusOK, gin.H{
			"users":   users,
			"cwecwec": "cwecwec",
		})
	})

	//var router *gin.Engine
	//
	//router = gin.Default()
	//
	//gin.Def
	//
	//v1 := router.Group("/v1")
	//
	//v1.GET("/login", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "login",
	//	})
	//})
	//
	//v1.GET("/products", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "products",
	//	})
	//})
	//
	//v2 := router.Group("/etyuio")
	//
	//v2.GET("/products", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "login",
	//	})
	//})
	//
	//router.GET("/qwerty", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "qwerty",
	//	})
	//})

	//
	////var name []string = []string{"Mirodil"}
	////
	////router.GET("/ping", func(context *gin.Context) {
	////	context.JSON(http.StatusOK, gin.H{
	////		"Hello": name,
	////	})
	////})
	//
	//router.GET("/product/:id", func(context *gin.Context) {
	//	var id = context.Param("id")
	//
	//	context.JSON(http.StatusOK, gin.H{
	//		"id": id,
	//	})
	//})
	//
	//router.GET("/users/:name/*age", func(context *gin.Context) {
	//	var name = context.Param("name")
	//	age := context.Param("age")
	//
	//	fmt.Println(context.FullPath())
	//
	//	context.JSON(http.StatusOK, gin.H{
	//		"name": name,
	//		"age":  age,
	//	})
	//})
	//
	//router.POST("/users/:name/*action", func(context *gin.Context) {
	//	context.String(http.StatusOK, "%t", true)
	//})
	//
	//router.GET("/users/:name", func(c *gin.Context) {
	//	c.String(http.StatusOK, c.Param("name"))
	//})
	//
	//router.GET("/users/groups", func(c *gin.Context) {
	//
	//	c.String(http.StatusOK, "The available groups are [...]")
	//
	//})
	//
	//router.GET("/welcome", func(context *gin.Context) {
	//
	//	firstname := context.DefaultQuery("firstname", "Guest")
	//	lastname := context.Query("lastname")
	//
	//	context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	//
	//})
	//
	//router.POST("/form_post", func(context *gin.Context) {
	//	message := context.PostForm("message")
	//
	//	nick := context.DefaultPostForm("nick", "anonymous")
	//
	//	context.JSON(http.StatusOK, gin.H{
	//		"status":  "posted",
	//		"message": message,
	//		"nick":    nick,
	//	})
	//})
	//
	//router.POST("/post", func(context *gin.Context) {
	//	id := context.Query("id")
	//	page := context.DefaultQuery("page", "0")
	//	name := context.PostForm("name")
	//	message := context.DefaultPostForm("message", "Hello world")
	//
	//	context.String(http.StatusOK, "id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	//})
	//
	//router.POST("/post2", func(context *gin.Context) {
	//
	//	ids := context.QueryMap("ids")
	//	names := context.PostFormMap("names")
	//
	//	context.JSON(http.StatusOK, gin.H{
	//		"ids":   ids,
	//		"names": names,
	//	})
	//})
	//
	//router.MaxMultipartMemory = 8 << 20
	//
	//router.POST("/upload", func(c *gin.Context) {
	//
	//	fmt.Println("Hello")
	//	// Single file
	//	file, err := c.FormFile("file")
	//
	//	if err != nil {
	//		fmt.Println("Error1: ", err.Error())
	//	}
	//
	//	log.Println(file.Filename)
	//
	//	// Upload the file to specific dst.
	//	err = c.SaveUploadedFile(file, "./files/"+file.Filename)
	//
	//	if err != nil {
	//		fmt.Println("Error : ", err.Error())
	//		return
	//	}
	//
	//	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	//
	//})
	//
	//router.POST("/upload2", func(c *gin.Context) {
	//	// Multipart form
	//	form, _ := c.MultipartForm()
	//	files := form.File["upload[]"]
	//
	//	for _, file := range files {
	//		// Upload the file to specific dst.
	//		c.SaveUploadedFile(file, "./files/"+file.Filename)
	//	}
	//	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	//})

	log.Fatal(r.Run(":9999"))

}
