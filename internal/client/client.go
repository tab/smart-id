package client

import (
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/tab/smart-id/internal/config"
	"github.com/tab/smart-id/internal/errors"
)

const (
	CertificateLevel = "QUALIFIED"
	HashType         = "SHA512"
	InteractionType  = "displayTextAndPIN"
	Timeout          = 60
	URL              = "https://sid.demo.sk.ee/smart-id-rp/v2"
)

// Client holds the client configuration and the HTTP client
type Client struct {
	config     *config.Config
	httpClient *resty.Client
}

// NewClient creates a new Smart-ID client instance
func NewClient(opts ...config.Option) (*Client, error) {
	cfg := &config.Config{
		CertificateLevel: CertificateLevel,
		HashType:         HashType,
		InteractionType:  InteractionType,
		URL:              URL,
		Timeout:          Timeout,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	httpClient := resty.New().
		SetTimeout(time.Duration(cfg.Timeout)*time.Second).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json")

	return &Client{
		config:     cfg,
		httpClient: httpClient,
	}, nil
}

func validateConfig(cfg *config.Config) error {
	if cfg.RelyingPartyName == "" {
		return errors.ErrMissingRelyingPartyName
	}

	if cfg.RelyingPartyUUID == "" {
		return errors.ErrMissingRelyingPartyUUID
	}

	return nil
}
