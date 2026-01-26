func Register(svc *Service) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Email    string
            Password string
        }
        c.BindJSON(&req)

        if err := svc.Register(req.Email, req.Password); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }

        c.JSON(201, gin.H{"message": "user created"})
    }
}
