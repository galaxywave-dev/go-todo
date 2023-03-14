package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"galaxywave.com/go-todo/apiserver/controllers"
	"galaxywave.com/go-todo/apiserver/grpcsvc"
	"galaxywave.com/go-todo/apiserver/initializers"
	"galaxywave.com/go-todo/apiserver/middleware"
	"galaxywave.com/go-todo/apiserver/models"
	"galaxywave.com/go-todo/apiserver/services"
	pb "galaxywave.com/go-todo/todoapi"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	models.InitDBConnection() // new
	services.Init()
	go hostRestApi(8088)
	hostGrpcApi(8089)

}
func hostRestApi(port int) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true

	r := gin.Default()
	r.Use(cors.New(corsConfig))

	books := r.Group("/books")
	{
		books.GET("/", controllers.FindBooks)
		books.POST("/", controllers.CreateBook)
		books.GET("/:id", controllers.FindBook)
		books.PATCH("/:id", controllers.UpdateBook)
		books.DELETE("/:id", controllers.DeleteBook) // new
	}

	todos := r.Group("/todos")
	{
		todos.GET("/", controllers.FindTodos)
		todos.POST("/", controllers.CreateTodo)
		todos.GET("/:id", controllers.FindTodo)
		todos.PATCH("/:id", controllers.UpdateTodo)
		todos.DELETE("/:id", controllers.DeleteTodo) // new
		// Define OPTIONS route to handle preflight request
		todos.OPTIONS("/:id", func(c *gin.Context) {
			// Set CORS headers
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type")

			// Respond with status 204 (no content)
			c.Status(204)
		})

	}
	// auth
	initializers.ConnectDB()
	router := r.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Implement Google OAuth2 in Golang"})
	})

	auth_router := router.Group("/auth")
	auth_router.POST("/register", controllers.SignUpUser)
	auth_router.POST("/login", controllers.SignInUser)
	auth_router.GET("/logout", middleware.DeserializeUser(), controllers.LogoutUser)

	router.GET("/sessions/oauth/google", controllers.GoogleOAuth)
	router.GET("/users/me", middleware.DeserializeUser(), controllers.GetMe)

	router.StaticFS("/images", http.Dir("public"))
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Route Not Found"})
	})

	r.Run(fmt.Sprintf(":%d", port))
}

func hostGrpcApi(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoManagerServer(s, &grpcsvc.TodoManager{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
