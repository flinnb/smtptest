package smtptest_test

import (
	"testing"

	"github.com/jhillyerd/enmime"
	"github.com/stretchr/testify/require"

	"github.com/flinnb/smtptest"
	"github.com/flinnb/smtptest/stassert"
	"github.com/flinnb/smtptest/testhelpers"
)

func TestWithServerBasic(t *testing.T) {
	b1, t1 := testhelpers.GetAttachment("puppy-asleep.jpg")
	b2, t2 := testhelpers.GetAttachment("theodolite.jpg")
	b3, t3 := testhelpers.GetAttachment("test.txt")

	msg := enmime.Builder().
		Subject("has to address test").
		From("Sender Name", "sender@localhost.dev").
		ReplyTo("Sender Name", "reply-to@localhost.dev").
		To("Recipient One", "r1@localhost.dev").
		To("Recipient Two", "r2@localhost.dev").
		CC("Copy Recipient One", "r1@localhost.dev").
		CC("Copy Recipient Two", "r2@localhost.dev").
		AddAttachment(b1, t1, "puppy-asleep.jpg").
		AddAttachment(b2, t2, "theodolite.jpg").
		AddAttachment(b3, t3, "test.txt")

	smtptest.WithSmtp(t, func(t *testing.T, srv *smtptest.SmtpServer) {
		sender := enmime.NewSMTP(*srv.Address, nil)

		err := msg.Send(sender)

		require.NoError(t, err)

		te := <-srv.Messages
		require.NotNil(t, te)

		stassert.HasToAddr(t, te, "r1@localhost.dev")
		stassert.HasToAddr(t, te, "r2@localhost.dev")

		stassert.HasToName(t, te, "Recipient One")
		stassert.HasToName(t, te, "Recipient Two")

		stassert.HasCcAddr(t, te, "r1@localhost.dev")
		stassert.HasCcAddr(t, te, "r2@localhost.dev")

		stassert.HasCcName(t, te, "Copy Recipient One")
		stassert.HasCcName(t, te, "Copy Recipient Two")

	})
}
