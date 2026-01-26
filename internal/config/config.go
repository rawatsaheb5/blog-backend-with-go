type Config struct{
	DBUrl string
	Port string
	JWTKey string
}

func LoadConfig() Config{

	return Config{
		DBUrl: os.Getenv("DB_URL"),
		Port: os.Getenv("PORT"),
		JWTKey: os.Getenv("JWT_KEY"),
	}
}