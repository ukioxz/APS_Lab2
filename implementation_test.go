package lab2

import (
	"testing"
	. "gopkg.in/check.v1"
	"fmt"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestPostfixCalculation(c *C) {
  res, err := PostfixCalculation("7 2 +")
	c.Assert(res, Equals, "9")
	c.Assert(err, IsNil)
}

func ExamplePostfixCalculation() {
	res, _ := PostfixCalculation("4 2 - 3 * 5 +")
	fmt.Println(res)

	// Output:
	// 11
}