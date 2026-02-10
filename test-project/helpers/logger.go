package helpers

import (
	"os"
	"github.com/rs/zerolog"
    "test-project/configs"
)

func InitializeZeroLogs() zerolog.Logger {
	output := os.Stdout
	// Add file based logging logic if needed (from configs)
	
	logger := zerolog.New(output).With().Timestamp().Str("service", configs.Cfg.General.ServiceName).Logger()
	return logger
}
