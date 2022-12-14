package ginx

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
	"io"
)

type signer struct {
	w      io.Writer
	h      hash.Hash
	secret string
}

func (s *signer) Write(b []byte) (int, error) {
	n, err := s.w.Write(b)
	s.h.Write(b)
	return n, err
}

func (s *signer) Signature() string {
	s.h.Write([]byte(s.secret))
	return hex.EncodeToString(s.h.Sum(nil))
}

func NewSigner(w io.Writer, secret string) *signer {
	return &signer{w: w, h: md5.New(), secret: secret}
}
