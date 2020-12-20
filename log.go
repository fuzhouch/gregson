package gregson

import (
	"io"
	"os"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitGlobalZeroLog initialize global zerolog.log instance with default
// configuration.
func InitGlobalZeroLog(toFile io.Writer) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	log.Logger = zerolog.New(toFile).
		With().
		Timestamp().
		Str("host", hostname).
		Logger()

	// Always set to UTC timestamp for readability reason.
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(time.UTC)
	}
	zerolog.TimeFieldFormat = time.RFC3339

	// Always ignore debug log and force all debug & prod
	// environments use same log-level settings. We should not
	// expect an issue reported in production can always be
	// reproduced "with more info" by rerunning with debug level
	// enabled. A correct practice is we should ALWAYS ensure
	// production log contain all information for diagnose.
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Never allow sampling. It will make diagnose
	// of online issue pretty difficult.
	zerolog.DisableSampling(true)
	return nil
}

// SetOffGlobalZeroLog configures global zerolog to discard messages.
// This is useful when writing tests, but do not use in production.
func SetOffGlobalZeroLog() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

// SetGlobalZeroLogToGin adds global logger to Gin engine.
func SetGlobalZeroLogToGin(r *gin.Engine) {
	r.Use(logger.SetLogger(logger.Config{
		Logger: &log.Logger,
		UTC:    true,
		// I used to think we should skip /metrics for
		// Prometheus, but then I change my mind because
		// I don't want anything hiding from logs. If
		// something wrong from internal that causes continuous
		// call to /metrics, we should be able to spot it.
	}))
}
