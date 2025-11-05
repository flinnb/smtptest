package tsrequire

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/flinnb/smtptest"
)

func HasToAddr(t *testing.T, te *smtptest.TestEnvelope, to string) {
	require.True(t, te.HasToAddr(to))
}

func HasCcAddr(t *testing.T, te *smtptest.TestEnvelope, cc string) {
	require.True(t, te.HasCcAddr(cc))
}

func HasToName(t *testing.T, te *smtptest.TestEnvelope, to string) {
	require.True(t, te.HasToName(to))
}

func HasCcName(t *testing.T, te *smtptest.TestEnvelope, cc string) {
	require.True(t, te.HasCcName(cc))
}

func HasSubject(t *testing.T, te *smtptest.TestEnvelope, subject string) {
	require.True(t, te.HasSubject(subject))
}

func HasBody(t *testing.T, te *smtptest.TestEnvelope, body string) {
	require.True(t, te.HasBody(body))
}
