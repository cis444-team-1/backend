package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nedpals/supabase-go"
)

type Config struct {
	DBConnectionUrl     string
	JWTAccessSecret     string
	JWTRefreshSecret    string
	FrontendOrigin      string
	S3BucketName        string
	AWSRegion           string
	AWSAccessKey        string
	AWSSecretAccessKey  string
	AWSCloudfrontDomain string
	SupabaseAnonKey     string
	SupabaseUrl         string
	SupabaseClient      *supabase.Client
}

var (
	once   sync.Once
	config *Config
)

// Loads the environment file using a relative path from the config package.
// Since the config package location is consistent within the project structure,
// we can reliably determine the path to environment files.
//
// Project structure:
// cis444-team-1/
// ├── .env
// ├── .env.test
// ├── config/
// │   └── config.go (this file)
// └── ...
func loadEnvFile(envFile string) error {
	// Get the directory of the current file (config.go)
	_, filename, _, _ := runtime.Caller(0)
	configDir := filepath.Dir(filename)

	// Move up one directory to reach project root
	projectRoot := filepath.Dir(configDir)

	// Construct path to env file
	envPath := filepath.Join(projectRoot, envFile)
	return godotenv.Load(envPath)
}

func LoadConfig() *Config {
	once.Do(func() {
		// Determine which env file to load
		envFile := ".env"
		if os.Getenv("GO_ENV") == "test" {
			envFile = ".env.test"
		}

		err := loadEnvFile(envFile)
		if err != nil {
			log.Printf("Warning: could not load %s file: %v", envFile, err)
		}

		config = &Config{
			DBConnectionUrl:     getEnv("DATABASE_URL"),
			JWTAccessSecret:     getEnv("JWT_ACCESS_SECRET"),
			JWTRefreshSecret:    getEnv("JWT_REFRESH_SECRET"),
			FrontendOrigin:      getEnv("FRONTEND_ORIGIN", "http://localhost:5173"),
			S3BucketName:        getEnv("S3_BUCKET_NAME"),
			AWSRegion:           getEnv("AWS_REGION"),
			AWSAccessKey:        getEnv("AWS_ACCESS_KEY"),
			AWSSecretAccessKey:  getEnv("AWS_SECRET_ACCESS_KEY"),
			AWSCloudfrontDomain: getEnv("AWS_CLOUDFRONT_DOMAIN"),
			SupabaseUrl:         getEnv("SUPABASE_URL"),
			SupabaseAnonKey:     getEnv("SUPABASE_ANON_KEY"),
		}

		// Initialize Supabase client
		config.SupabaseClient = supabase.CreateClient(config.SupabaseUrl, config.SupabaseAnonKey)
	})
	return config
}

func getEnv(key string, defaultValue ...string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue[0]
}

// Added CORS for frontend, need to allow frontend origin otherwise requests get rejected
func InitCORS(e *echo.Echo) {
	cfg := LoadConfig()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{cfg.FrontendOrigin}, // Explicitly allow frontend origin
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Authorization"},
		AllowCredentials: true, // Important: Allow credentials to send over the cookie
	}))
}
