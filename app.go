type User struct {
    ID        int    `json:"id"`
    Username  string `json:"username"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

type Photo struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Caption   string `json:"caption"`
    PhotoUrl  string `json:"photo_url"`
    UserID    int    `json:"user_id"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}