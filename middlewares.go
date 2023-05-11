func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := helpers.GetTokenFromHeader(c)
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // validate token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // set user id from token to context
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            userID := claims["userID"].(string)
            c.Set("userID", userID)
        }

        c.Next()
    }
}

func checkPermission() gin.HandlerFunc {
    return func(c *gin.Context) {
        // get user id from context
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // get user from database
        var user User
        if err := db.First(&user, userID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            c.Abort()
            return
        }

        // check if the user making the request is the same user being updated/deleted
        if c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
            requestedUserID := c.Param("userId")
            if user.ID != requestedUserID {
                c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
                c.Abort()
                return
            }
        }

        c.Next()
    }
}
