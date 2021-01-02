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

func NewZeroLog(toFile io.Writer) (zerolog.Logger, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return zerolog.New(nil), err
	}

	logger := zerolog.New(toFile).
		With().
		Timestamp().
		Str("host", hostname).
		Logger()
	return logger, nil
}

// SetGlobalZeroLogPolicy sets default zerolog settings used by Gregson.
// The following policy are enforced:
//
// 1. Always use RFC3339 format ("2006-01-02T15:04:05Z07:00")
// 2. Timestamp returns UTC.
// 3. Prints only INFO level logs or above.
// 4. Sampling is disabled.
//
// The reason for #1 and #2 is Gregson tries to make service log easy
// to read when distributed across multiple regions.
//
// Special notes for #3: Gregson explicitly disallows developers apply
// different logging levels for dev and production enviornment. This is
// intentionally forbids a practice, that developers expect they can
// use dev environment to re-produce an issue reported in production,
// but with more logging info internally. This approach is proven not
// working for Internet oriented services, because many
// issues are triggerred only under high load which can be pretty
// difficult to simulate in dev environment. A correct practice, is to
// make sure service always generate identical information no matter in
// which environment.
//
// #4 is set by almost same reason with #3. Sampling sacrifaces diagnose
// feasibility to get smaller file size. This is usually not worthy
// in production environment.
func SetGlobalZeroLogPolicy() {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(time.UTC)
	}
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.DisableSampling(true)
}

// InitGlobalZeroLog initialize global zerolog.log instance with default
// configuration. It also sets global zerolog settings.
func InitGlobalZeroLog(toFile io.Writer) error {
	logger, err := NewZeroLog(toFile)
	if err != nil {
		return err
	}
	log.Logger = logger
	SetGlobalZeroLogPolicy()
	return nil
}

// SetOffGlobalZeroLog configures global zerolog to discard messages.
// This is useful when writing tests, but do not use in production.
func SetOffGlobalZeroLog() {
	log.Logger = zerolog.Nop()
}

// HookGlobalZeroLog install zerolog as internal logger of Gin.
// We can use zerolog to output information like requested URLs,
// response code, latency, etc.
func HookGlobalZeroLog(r *gin.Engine) {
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
