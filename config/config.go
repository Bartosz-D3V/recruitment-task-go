package config

import (
	"github.com/Bartosz-D3V/recruitment-task-go/helpers"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
)

type AppConfig struct {
	Logger  *zap.SugaredLogger
	Numbers *[]int
}

func New() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}
	logLevel := getLogLevelIota(os.Getenv("LOG_LEVEL"))

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(logLevel)

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	sugarLogger := logger.Sugar()
	defer logger.Sync()

	nums, err := helpers.ReadFileIntoSlice(sugarLogger, os.Getenv("INPUT_PATH"))
	if err != nil {
		log.Fatal("There was a problem reading a file. Please check settings and re-run the application", err)
	}
	sugarLogger.Info("Successfully parsed input file")

	return AppConfig{
		Logger:  sugarLogger,
		Numbers: nums,
	}
}

func getLogLevelIota(logLevel string) zapcore.Level {
	switch strings.ToLower(logLevel) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}
