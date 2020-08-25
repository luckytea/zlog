package zlog

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// New - return new logger.
// Output format depend on given arg.
func New(ver string) *zerolog.Logger {
	var output io.Writer

	if ver == "dev" {
		output = zerolog.NewConsoleWriter()
	} else {
		output = os.Stdout
	}

	log := zerolog.New(output).With().Str("version", ver).Timestamp().Logger()

	return &log
}
