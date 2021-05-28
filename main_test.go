package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSample(t *testing.T) {
	suite.Run(t, new(SampleSuite))
}

type SampleSuite struct {
	suite.Suite
}

func (s *SampleSuite) SetupSuite() {
}

func (s *SampleSuite) TearDownSuite() {
}

func (s *SampleSuite) SetupTest() {
}

func (s *SampleSuite) TearDownTest() {
}

func (s *SampleSuite) TestSomething() {
	sum := 1 + 1
	s.Equal(2, sum)
}
