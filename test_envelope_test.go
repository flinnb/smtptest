package smtptest_test

import (
	"testing"

	"github.com/jhillyerd/enmime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flinnb/smtptest"
)

func TestHasRecipient(t *testing.T) {
	msg := enmime.Builder().
		Subject("has to address test").
		From("Sender Name", "sender@localhost.dev").
		ReplyTo("Sender Name", "reply-to@localhost.dev").
		To("Recipient One", "r1@localhost.dev").
		To("Recipient Two", "r2@localhost.dev")

	p, err := msg.Build()
	require.NoError(t, err)
	e, err := enmime.EnvelopeFromPart(p)
	require.NoError(t, err)

	te := smtptest.NewTestEnvelope(e, []string{})

	t.Run("Envelope has correct `to` email addresses", func(t *testing.T) {
		assert.True(t, te.HasToAddr("r1@localhost.dev"))
		assert.True(t, te.HasToAddr("r2@localhost.dev"))
	})
	t.Run("Envelope has correct `to` email name", func(t *testing.T) {
		assert.True(t, te.HasToName("Recipient One"))
		assert.True(t, te.HasToName("Recipient Two"))
	})
}

func TestHasCopiedRecipient(t *testing.T) {
	msg := enmime.Builder().
		Subject("has to address test").
		From("Sender Name", "sender@localhost.dev").
		ReplyTo("Sender Name", "reply-to@localhost.dev").
		CC("Copied Recipient One", "r1@localhost.dev").
		CC("Copied Recipient Two", "r2@localhost.dev")

	p, err := msg.Build()
	require.NoError(t, err)
	e, err := enmime.EnvelopeFromPart(p)
	require.NoError(t, err)

	te := smtptest.NewTestEnvelope(e, []string{})

	t.Run("Envelope has correct `cc` email addresses", func(t *testing.T) {
		assert.True(t, te.HasCcAddr("r1@localhost.dev"))
		assert.True(t, te.HasCcAddr("r2@localhost.dev"))
	})
	t.Run("Envelope has correct `cc` email name", func(t *testing.T) {
		assert.True(t, te.HasCcName("Copied Recipient One"))
		assert.True(t, te.HasCcName("Copied Recipient Two"))
	})
}
