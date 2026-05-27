package email

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"net/smtp"
	"time"
)

type (
	Email struct {
		Auth     smtp.Auth
		Header   map[string]string
		Template *template.Template
		// TLSConfig, when non-nil, is used for both implicit TLS (SMTPS)
		// and STARTTLS. Callers that need a custom root pool or a
		// specific ServerName should set this. The config is cloned per
		// dial, so callers can reuse a single value across sends; they
		// must not mutate it concurrently with an in-flight Send.
		TLSConfig *tls.Config
		// DialTimeout caps the TCP (and, for SMTPS, TLS) connect phase.
		// It does not bound the full SMTP conversation after the client
		// is returned. Zero means no caller-imposed timeout.
		DialTimeout time.Duration
		smtpAddress string
	}

	Message struct {
		ID          string  `json:"id"`
		From        string  `json:"from"`
		To          string  `json:"to"`
		CC          string  `json:"cc"`
		Subject     string  `json:"subject"`
		BodyText    string  `json:"body_text"`
		BodyHTML    string  `json:"body_html"`
		Inlines     []*File `json:"inlines"`
		Attachments []*File `json:"attachments"`
		buffer      *bytes.Buffer
		boundary    string
	}

	File struct {
		Name    string
		Type    string
		Content string
	}
)

func New(smtpAddress string) *Email { _ = "STUB: not implemented"; return nil }

func (m *Message) writeHeader(key, value string) { _ = "STUB: not implemented"; return }

func (m *Message) writeBoundary() { _ = "STUB: not implemented"; return }

func (m *Message) writeText(content string, contentType string) { _ = "STUB: not implemented"; return }

func (m *Message) writeFile(f *File, disposition string) { _ = "STUB: not implemented"; return }

func (e *Email) Send(m *Message) (err error) {
	_ = "STUB: not implemented"
	// Message header
	return nil
}

// Extra

// Message body

// Inlines/attachments

// Dial. Port 465 is SMTPS (implicit TLS) per IANA and always uses
// TLS. Other ports connect plaintext and opportunistically upgrade
// to STARTTLS only if the server advertises it — if the server
// doesn't, the connection stays in the clear. Operators that
// require TLS must use port 465.

// Authenticate

// Send message

func (e *Email) dial() (*smtp.Client, error) { _ = "STUB: not implemented"; return nil, nil }

// Always clone so we never mutate the caller's TLSConfig.

// Drive EHLO explicitly so we can surface its error. (*Client).Extension
// triggers a lazy hello() and swallows its error, which would silently
// treat a failed EHLO as "STARTTLS not advertised" and stay cleartext.
