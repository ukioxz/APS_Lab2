package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

func (s *TestSuite) TestCompute(c *C) {
	var buffer bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("4 2 - 3 * 5 +"),
		Output: &buffer,
	}
	errCompute := handler.Compute()
	result, err := PostfixCalculation("4 2 - 3 * 5 +")

	c.Assert(buffer.String(), Equals, result)
	c.Assert(err, IsNil)
	c.Assert(errCompute, IsNil)
}

func (s *TestSuite) TestComputeError1(c *C) {
	var buffer bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("a b c"),
		Output: &buffer,
	}
	errCompute := handler.Compute()
	c.Assert(errCompute, ErrorMatches, "invalid input")
}

func (s *TestSuite) TestComputeError2(c *C) {
	var buffer bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader(" "),
		Output: &buffer,
	}
	errCompute := handler.Compute()
	c.Assert(errCompute, ErrorMatches, "syntax error")
}

func (s *TestSuite) TestComputeError3(c *C) {
	var buffer bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("4 8 9 10 +"),
		Output: &buffer,
	}
	errCompute := handler.Compute()
	c.Assert(errCompute, ErrorMatches, "syntax error")
}

func (s *TestSuite) TestComputeError4(c *C) {
	var buffer bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("10 8 12.5 8"),
		Output: &buffer,
	}
	errCompute := handler.Compute()
	c.Assert(errCompute, ErrorMatches, "syntax error")
}

func (s *TestSuite) TestComputeError5(c *C) {
	var buffer bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("abc + 15"),
		Output: &buffer,
	}
	errCompute := handler.Compute()
	c.Assert(errCompute, ErrorMatches, "invalid input")
}
