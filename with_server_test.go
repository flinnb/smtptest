package smtptest_test

import (
	"testing"

	"github.com/jhillyerd/enmime"
	"github.com/stretchr/testify/require"

	"github.com/flinnb/smtptest"
	"github.com/flinnb/smtptest/stassert"
)

func TestWithServerBasic(t *testing.T) {

	msg := enmime.Builder().
		Subject("has to address test").
		From("Sender Name", "sender@localhost.dev").
		ReplyTo("Sender Name", "reply-to@localhost.dev").
		To("Recipient One", "r1@localhost.dev").
		To("Recipient Two", "r2@localhost.dev")

	smtptest.WithSmtp(t, func(t *testing.T, srv *smtptest.SmtpServer) {
		sender := enmime.NewSMTP(*srv.Address, nil)

		err := msg.Send(sender)
		require.NoError(t, err)

		te := srv.GetLastEnvelope()
		require.NotNil(t, te)

		stassert.HasToAddr(t, te, "r1@localhost.dev")
		stassert.HasToAddr(t, te, "r2@localhost.dev")

	})
}
