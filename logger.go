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

	switch ver {
	case "dev":
		output = zerolog.NewConsoleWriter()
	case "test":
		output = io.Discard
	default:
		output = os.Stdout
	}

	log := zerolog.New(output).With().Str("version", ver).Timestamp().Logger()

	return &log
}
