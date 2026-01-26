func ErrorResponse(msg string) gin.H {
    return gin.H{"error": msg}
}
