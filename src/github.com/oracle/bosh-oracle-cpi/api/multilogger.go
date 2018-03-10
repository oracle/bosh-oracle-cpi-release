package api

import (
	"bytes"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

// MultiLogger is a Logger that logs messages to two locations. Messages logged to it
// are sent to an underlying byte buffer and to a BOSH logger
type MultiLogger struct {
	boshlog.Logger
	LogBuff *bytes.Buffer
}
