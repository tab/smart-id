package utils

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/tab/smartid/internal/errors"
)

var (
	identityRegex        = regexp.MustCompile(`^(PAS|IDC|PNO)([A-Z]{2})-([A-Za-z0-9]+)$`)
	latvianIdentityRegex = regexp.MustCompile(`^(PAS|IDC|PNO)([A-Z]{2})-([A-Za-z0-9]+)-([A-Za-z0-9]+)$`)
)

type Person struct {
	IdentityNumber string
	PersonalCode   string
	FirstName      string
	LastName       string
}

func Extract(encodedCert string) (*Person, error) {
	certBytes, err := base64.StdEncoding.DecodeString(encodedCert)
	if err != nil {
		return nil, errors.ErrFailedToDecodeCertificate
	}

	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, errors.ErrFailedToParseCertificate
	}

	parts := strings.Split(cert.Subject.CommonName, ",")
	if len(parts) < 2 {
		return nil, errors.ErrInvalidCertificate
	}

	identity, err := parse(cert.Subject.SerialNumber)
	if err != nil {
		return nil, err
	}
	firstName := strings.TrimSpace(parts[0])
	lastName := strings.TrimSpace(parts[1])

	return &Person{
		IdentityNumber: cert.Subject.SerialNumber,
		PersonalCode:   identity.ID,
		FirstName:      firstName,
		LastName:       lastName,
	}, nil
}

type Identity struct {
	Country string
	Type    string
	ID      string
}

func parse(value string) (*Identity, error) {
	if value == "" {
		return nil, errors.ErrInvalidIdentityNumber
	}

	matches := identityRegex.FindStringSubmatch(value)
	latvian := latvianIdentityRegex.FindStringSubmatch(value)

	if len(matches) != 4 && len(latvian) != 5 {
		return nil, errors.ErrInvalidIdentityNumber
	}

	if len(latvian) != 0 {
		return &Identity{
			Type:    latvian[1],
			Country: latvian[2],
			ID:      fmt.Sprintf("%s-%s", latvian[3], latvian[4]),
		}, nil
	} else {
		return &Identity{
			Type:    matches[1],
			Country: matches[2],
			ID:      matches[3],
		}, nil
	}
}
