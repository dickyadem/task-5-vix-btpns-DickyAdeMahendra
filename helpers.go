package helpers

import (
    "crypto/rand"
    "encoding/base64"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "strings"
)

// GenerateRandomBytes returns securely generated random bytes.
func GenerateRandomBytes(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    // Note that err == nil only if we read len(b) bytes.
    if err != nil {
        return nil, err
    }

    return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded securely generated random string.
// It takes the number of bytes to generate as input.
func GenerateRandomString(n int) (string, error) {
    b, err := GenerateRandomBytes(n)
    return base64.URLEncoding.EncodeToString(b), err
}

// HashPassword returns the bcrypt hash of the password string.
func HashPassword(password string) ([]byte, error) {
    return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPasswordHash checks whether the provided hash matches the password string.
func CheckPasswordHash(password string, hash []byte) bool {
    err := bcrypt.CompareHashAndPassword(hash, []byte(password))
    return err == nil
}

// GetAuthorizationHeader retrieves the "Authorization" header value from the given context.
// It returns an empty string if the header is not present.
func GetAuthorizationHeader(c *gin.Context) string {
    authHeader := c.GetHeader("Authorization")
    if authHeader != "" {
        return strings.TrimPrefix(authHeader, "Bearer ")
    }
    return ""
}

// GetUserID retrieves the user ID from the JWT token in the "Authorization" header.
// It returns 0 if the user ID is not found or the token is invalid.
func GetUserID(c *gin.Context) uint {
    tokenString := GetAuthorizationHeader(c)
    if tokenString == "" {
        return 0
    }

    token, err := ParseJWT(tokenString)
    if err != nil {
        return 0
    }

    claims := token.Claims.(jwt.MapClaims)
    userID := uint(claims["user_id"].(float64))

    return userID
}

// ParseJWT parses the given JWT token string and returns the token object.
func ParseJWT(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(os.Getenv("JWT_SECRET")), nil
    })
    if err != nil {
        return nil, err
    }
    return token, nil
}

// GenerateJWT generates a new JWT token with the given user ID and expiration time.
func GenerateJWT(userID uint, expiresIn time.Duration) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(expiresIn).Unix(),
    })

    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// AuthenticateUser is a middleware function that checks whether the request is authorized by checking the "Authorization" header.
// It sets the user ID in the context if the request is authorized.
// It returns an error and stops the request chain if the request
