
Pada code sebelumnya, kita telah membuat struktur router dengan menggunakan framework Gin. Router ini akan berisi endpoint-endpoint API yang dapat diakses oleh pengguna, serta controller dan metodenya yang akan dipanggil saat endpoint diakses.

Berikut adalah contoh konfigurasi router yang dapat digunakan untuk API yang telah kita buat sebelumnya:

go
Copy code
func SetupRouter() *gin.Engine {
    // Initialize Gin router
    router := gin.Default()

    // Initialize middleware for JWT authentication
    authMiddleware := middlewares.NewAuthMiddleware()

    // Define User routes
    userController := controllers.UserController{DB: database.DB}
    userGroup := router.Group("/users")
    userGroup.POST("/register", userController.RegisterUser)
    userGroup.POST("/login", userController.LoginUser)
    userGroup.PUT("/:userId", authMiddleware.Authenticate(), userController.UpdateUser)
    userGroup.DELETE("/:userId", authMiddleware.Authenticate(), userController.DeleteUser)

    // Define Photo routes
    photoController := controllers.PhotoController{DB: database.DB}
    photoGroup := router.Group("/photos")
    photoGroup.POST("", authMiddleware.Authenticate(), photoController.CreatePhoto)
    photoGroup.GET("", photoController.GetPhotos)
    photoGroup.PUT("/:photoId", authMiddleware.Authenticate(), photoController.UpdatePhoto)
    photoGroup.DELETE("/:photoId", authMiddleware.Authenticate(), photoController.DeletePhoto)

    return router
}