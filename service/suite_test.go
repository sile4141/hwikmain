package service

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	// policy
	suite.Run(t, new(PolicyTestSuite))
}
