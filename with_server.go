package smtptest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestFunc func(t *testing.T, srv *SmtpServer)

func WithSmtp(t *testing.T, tf TestFunc) {

	srv := &SmtpServer{}
	err := srv.ListenAndServe(context.Background())
	require.NoError(t, err)

	defer srv.Close()

	tf(t, srv)

}
