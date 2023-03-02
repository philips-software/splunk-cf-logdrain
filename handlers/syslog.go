package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	syslog "github.com/RackSec/srslog"

	v2syslog "github.com/influxdata/go-syslog/v2"
	"github.com/influxdata/go-syslog/v2/rfc5424"
	"github.com/labstack/echo/v4"
)

type SyslogHandler struct {
	debug  bool
	token  string
	writer *syslog.Writer
	parser v2syslog.Machine
}

func NewSyslogHandler(token, syslogAddr string) (*SyslogHandler, error) {
	if token == "" {
		return nil, fmt.Errorf("missing TOKEN value")
	}
	handler := &SyslogHandler{}
	handler.token = token

	parser := rfc5424.NewParser()

	if os.Getenv("DEBUG") == "true" {
		handler.debug = true
	}
	writer, err := syslog.Dial("tcp", syslogAddr,
		syslog.LOG_WARNING|syslog.LOG_DAEMON, "splunk-cf-logdrain")
	if err != nil {
		return nil, fmt.Errorf("syslog: %w", err)
	}
	//writer.SetFramer(syslog.RFC5425MessageLengthFramer)
	writer.SetFormatter(RFC5424PassThroughFormatter)
	handler.writer = writer
	handler.parser = parser
	return handler, nil
}

func RFC5424PassThroughFormatter(_ syslog.Priority, _, _, content string) string {
	return content
}

func (h *SyslogHandler) Handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		t := c.Param("token")
		if h.token != t {
			return c.String(http.StatusUnauthorized, "")
		}
		b, _ := io.ReadAll(c.Request().Body)
		_, err := h.parser.Parse(b)
		if err != nil {
			return err
		}
		_, _ = h.writer.Write(b)
		return c.String(http.StatusOK, "")
	}
}
