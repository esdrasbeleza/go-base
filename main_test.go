package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
