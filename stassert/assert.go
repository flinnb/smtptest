package stassert

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flinnb/smtptest"
)

func HasToAddr(t *testing.T, te *smtptest.TestEnvelope, to string) bool {
	return assert.True(t, te.HasToAddr(to))
}

func HasCcAddr(t *testing.T, te *smtptest.TestEnvelope, cc string) bool {
	return assert.True(t, te.HasCcAddr(cc))
}

func HasToName(t *testing.T, te *smtptest.TestEnvelope, to string) bool {
	return assert.True(t, te.HasToName(to))
}

func HasCcName(t *testing.T, te *smtptest.TestEnvelope, cc string) bool {
	return assert.True(t, te.HasCcName(cc))
}

func HasSubject(t *testing.T, te *smtptest.TestEnvelope, subject string) bool {
	return assert.True(t, te.HasSubject(subject))
}

func HasBody(t *testing.T, te *smtptest.TestEnvelope, body string) bool {
	return assert.True(t, te.HasBody(body))
}
