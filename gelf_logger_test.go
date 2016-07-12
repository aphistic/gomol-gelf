package gomolgelf

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type GomolSuite struct{}

var _ = Suite(&GomolSuite{})

func (s *GomolSuite) TestGelfNew(c *C) {
	cfg := NewGelfLoggerConfig()
	cfg.Hostname = "host"
	cfg.Port = 1234

	l, err := NewGelfLogger(cfg)
	c.Assert(err, IsNil)
	c.Check(l, NotNil)
	c.Check(l.config.Hostname, Equals, "host")
	c.Check(l.config.Port, Equals, 1234)
}
