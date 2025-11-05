package smtptest

import "github.com/jhillyerd/enmime"

type TestEnvelope struct {
	toAddrs  map[string]bool
	ccAddrs  map[string]bool
	toNames  map[string]bool
	ccNames  map[string]bool
	envelope *enmime.Envelope
}

func NewTestEnvelope(e *enmime.Envelope) (te *TestEnvelope) {
	te = &TestEnvelope{
		toAddrs: map[string]bool{},
		ccAddrs: map[string]bool{},
		toNames: map[string]bool{},
		ccNames: map[string]bool{},
	}
	to, _ := e.AddressList("to")
	for _, email := range to {
		te.toAddrs[email.Address] = true
		te.toNames[email.Name] = true
	}
	cc, _ := e.AddressList("cc")
	for _, email := range cc {
		te.ccAddrs[email.Address] = true
		te.ccNames[email.Name] = true
	}
	te.envelope = e

	return
}

func (te *TestEnvelope) HasToAddr(to string) bool {
	_, ok := te.toAddrs[to]
	return ok
}

func (te *TestEnvelope) HasCcAddr(cc string) bool {
	_, ok := te.ccAddrs[cc]
	return ok
}

func (te *TestEnvelope) HasToName(to string) bool {
	_, ok := te.toNames[to]
	return ok
}

func (te *TestEnvelope) HasCcName(cc string) bool {
	_, ok := te.ccNames[cc]
	return ok
}

func (te *TestEnvelope) HasSubject(subject string) bool {
	return te.envelope.GetHeader("Subject") == subject
}

func (te *TestEnvelope) HasBody(body string) bool {
	return te.envelope.Text == body || te.envelope.HTML == body
}
