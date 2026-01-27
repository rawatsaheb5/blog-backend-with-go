# Logger Package

This package provides structured logging using the [zap](https://github.com/uber-go/zap) library with support for multiple environments.

## Features

- **Environment-based configuration**: Different log formats and levels for development, production, and test environments
- **Structured logging**: JSON output in production, human-readable console output in development
- **Gin integration**: Middleware for HTTP request logging and panic recovery
- **Global logger**: Easy access to logger instance throughout the application

## Environments

### Development
- **Level**: Debug
- **Format**: Console (colored, human-readable)
- **Features**: Includes caller information, stack traces, and detailed timestamps

### Production
- **Level**: Info
- **Format**: JSON (structured, machine-readable)
- **Features**: Optimized for log aggregation tools, minimal overhead

### Test
- **Level**: Warn
- **Format**: Console (simplified)
- **Features**: Reduced verbosity for cleaner test output

## Usage

### Initialization

Initialize the logger in your `main.go`:

```go
import "github.com/rawatsaheb5/blog-backend-with-go/pkg/logger"

func main() {
    cfg := config.LoadConfig()
    
    // Initialize logger based on environment
    if err := logger.Init(cfg.Environment); err != nil {
        log.Fatalf("Failed to initialize logger: %v", err)
    }
    defer logger.Sync()
    
    // Your application code...
}
```

### Basic Logging

```go
import "github.com/rawatsaheb5/blog-backend-with-go/pkg/logger"

// Using structured logger
logger.GetLogger().Info("User logged in",
    zap.String("user_id", "123"),
    zap.String("email", "user@example.com"),
)

// Using sugared logger (easier syntax)
logger.GetSugar().Infof("User %s logged in", username)
logger.GetSugar().Errorf("Failed to process request: %v", err)
```

### Log Levels

```go
// Debug - detailed information for debugging
logger.GetSugar().Debug("Debug message")

// Info - general informational messages
logger.GetSugar().Info("Info message")

// Warn - warning messages
logger.GetSugar().Warn("Warning message")

// Error - error messages
logger.GetSugar().Error("Error message")

// Fatal - logs and then calls os.Exit(1)
logger.GetSugar().Fatal("Fatal error")
```

### With Fields

```go
// Structured logging with fields
logger.GetLogger().Info("Processing request",
    zap.String("method", "POST"),
    zap.String("path", "/api/users"),
    zap.Int("status", 200),
    zap.Duration("latency", time.Since(start)),
)
```

### Gin Middleware

The logger package includes Gin middleware for HTTP request logging:

```go
import "github.com/rawatsaheb5/blog-backend-with-go/pkg/logger"

r := gin.New()
r.Use(logger.GinLogger())    // Logs all HTTP requests
r.Use(logger.GinRecovery())  // Recovers from panics and logs them
```

## Environment Variables

Set the `ENV` or `ENVIRONMENT` variable to control the logger behavior:

```bash
# Development (default)
ENV=development

# Production
ENV=production

# Test
ENV=test
```

## Examples

### Logging in Handlers

```go
func Register(svc *Service) gin.HandlerFunc {
    return func(c *gin.Context) {
        logger.GetSugar().Info("Registration attempt", "email", req.Email)
        
        if err := svc.Register(req.Email, req.Password); err != nil {
            logger.GetSugar().Errorf("Registration failed: %v", err)
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        
        logger.GetSugar().Infof("User registered successfully: %s", req.Email)
        c.JSON(201, gin.H{"message": "user created"})
    }
}
```

### Logging in Services

```go
func (s *Service) Login(email, password string) (*User, string, error) {
    logger.GetSugar().Debugf("Attempting login for email: %s", email)
    
    user, err := s.repo.FindByEmail(email)
    if err != nil {
        logger.GetSugar().Errorf("Login failed - user not found: %s", email)
        return nil, "", err
    }
    
    logger.GetSugar().Infof("Login successful for user: %d", user.ID)
    return user, token, nil
}
```

## Best Practices

1. **Use appropriate log levels**: Debug for development details, Info for important events, Error for errors
2. **Include context**: Always include relevant fields (user ID, request ID, etc.)
3. **Don't log sensitive data**: Avoid logging passwords, tokens, or personal information
4. **Use structured logging**: Prefer `zap.Field` over string formatting for better log parsing
5. **Sync before exit**: Always call `logger.Sync()` before application exit (use defer)
