package lab2

import (
	"testing"
	. "gopkg.in/check.v1"
	"fmt"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestPostfixCalculation1(c *C) {
  res, err := PostfixCalculation("7 2 + 2 *")
	c.Assert(res, Equals, "18")
	c.Assert(err, IsNil)
}

func (s *TestSuite) TestPostfixCalculation2(c *C) {
  res, err := PostfixCalculation("2 3 ^ 8 +")
	c.Assert(res, Equals, "16")
	c.Assert(err, IsNil)
}

func (s *TestSuite) TestPostfixCalculation3(c *C) {
  res, err := PostfixCalculation("4 2 - 2 ^")
	c.Assert(res, Equals, "4")
	c.Assert(err, IsNil)
}

func (s *TestSuite) TestPostfixCalculation4(c *C) {
  res, err := PostfixCalculation("2 3 * 1 + 7 3 + - 3 / 4 3 * + 8 -")
	c.Assert(res, Equals, "3")
	c.Assert(err, IsNil)
}

func (s *TestSuite) TestPostfixCalculation5(c *C) {
  res, err := PostfixCalculation("12 3 * 6 / 18 3 * + 2 2 ^ /")
	c.Assert(res, Equals, "15")
	c.Assert(err, IsNil)
}

func (s *TestSuite) TestPostfixCalculation6(c *C) {
  res, err := PostfixCalculation("20 3 * 6 / 12 3 * + 7 3 * 2 + /")
	c.Assert(res, Equals, "2")
	c.Assert(err, IsNil)
}

func (s *TestSuite) TestPostfixCalculation7(c *C) {
	_, err := PostfixCalculation(" ")
	c.Assert(err, ErrorMatches, "empty string")
}

func (s *TestSuite) TestPostfixCalculation8(c *C) {
	_, err := PostfixCalculation("h + 2")
	c.Assert(err, ErrorMatches, "invalid input")
}

func (s *TestSuite) TestPostfixCalculation9(c *C) {
	_, err := PostfixCalculation("g o")
	c.Assert(err, ErrorMatches, "invalid input")
}

func ExamplePostfixCalculation() {
	res, _ := PostfixCalculation("4 2 - 3 * 5 +")
	fmt.Println(res)

	// Output:
	// 11
}
