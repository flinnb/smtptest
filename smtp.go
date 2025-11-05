package smtptest

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/fzerorubigd/smtpd"
	"github.com/jhillyerd/enmime"
)

type SmtpServer struct {
	Address      *string
	lastEnvelope *enmime.Envelope
	listener     net.Listener
}

func getFreePort() (port int, err error) {
	var a *net.TCPAddr
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			defer l.Close()
			return l.Addr().(*net.TCPAddr).Port, nil
		}
	}
	return
}

func (s *SmtpServer) handler(remoteAddr net.Addr, from string, to []string, data []byte) {
	e, _ := enmime.ReadEnvelope(bytes.NewReader(data))
	s.lastEnvelope = e
}

func (s *SmtpServer) GetLastEnvelope() *TestEnvelope {
	return NewTestEnvelope(s.lastEnvelope)
}

func (s *SmtpServer) ListenAndServe(ctx context.Context) (err error) {

	hn, _ := os.Hostname()
	port, _ := getFreePort()
	address := fmt.Sprintf(":%d", port)
	s.Address = &address
	appName := "smtptest"
	hostName := hn

	lc := &net.ListenConfig{}
	s.listener, err = lc.Listen(ctx, "tcp", *s.Address)
	if err != nil {
		return err
	}

	opts := []smtpd.OptionSetter{
		smtpd.WithAddress(*s.Address),
		smtpd.WithAppName(appName),
		smtpd.WithHostname(hostName),
		smtpd.AllowAuthMechanisms("LOGIN", true),
		smtpd.WithAuthHandler(
			func(
				remoteAddr net.Addr,
				mechanism string,
				username []byte,
				password []byte,
				shared []byte,
			) (bool, error) {
				// Require auth, but don't actually validate it.
				return true, nil
			},
			true),
	}

	srv, err := smtpd.NewServer(s.handler, opts...)
	if err != nil {
		return err
	}
	go func() {
		if err := srv.ServeContext(ctx, s.listener); err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

func (s *SmtpServer) Close() {
	s.listener.Close()
}
