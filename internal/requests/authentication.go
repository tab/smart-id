package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/tab/smartid/internal/config"
	"github.com/tab/smartid/internal/errors"
	"github.com/tab/smartid/internal/models"
	"github.com/tab/smartid/internal/utils"
)

const (
	MaxIdleConnections        = 10000
	MaxIdleConnectionsPerHost = 10000
	IdleConnTimeout           = 90 * time.Second
	TLSHandshakeTimeout       = 10 * time.Second
	Timeout                   = 60 * time.Second
)

func CreateAuthenticationSession(
	ctx context.Context,
	cfg *config.Config,
	identity string,
) (*models.Session, error) {
	hash, err := utils.GenerateHash(cfg.HashType)
	if err != nil {
		return nil, err
	}

	body := models.AuthenticationSessionRequest{
		RelyingPartyName:       cfg.RelyingPartyName,
		RelyingPartyUUID:       cfg.RelyingPartyUUID,
		NationalIdentityNumber: identity,
		CertificateLevel:       cfg.CertificateLevel,
		Hash:                   hash,
		HashType:               cfg.HashType,
		AllowedInteractionsOrder: []models.AllowedInteraction{
			{
				Type:          cfg.InteractionType,
				DisplayText60: cfg.Text,
			},
		},
	}

	endpoint := fmt.Sprintf("%s/authentication/etsi/%s", cfg.URL, identity)
	response, err := httpClient().R().SetContext(ctx).SetBody(body).Post(endpoint)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, errors.ErrSmartIdProviderError
	}

	var result models.Session
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}

	code, err := utils.GenerateVerificationCode(hash)
	if err != nil {
		return nil, err
	}

	return &models.Session{
		Id:   result.Id,
		Code: code,
	}, nil
}

func FetchAuthenticationSession(
	ctx context.Context,
	cfg *config.Config,
	sessionId string,
) (*models.AuthenticationResponse, error) {
	endpoint := fmt.Sprintf("%s/session/%s", cfg.URL, sessionId)

	response, err := httpClient().R().SetContext(ctx).Get(endpoint)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, errors.ErrSmartIdProviderError
	}

	var result models.AuthenticationResponse
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func httpClient() *resty.Client {
	transport := &http.Transport{
		MaxIdleConns:        MaxIdleConnections,
		MaxIdleConnsPerHost: MaxIdleConnectionsPerHost,
		IdleConnTimeout:     IdleConnTimeout,
		TLSHandshakeTimeout: TLSHandshakeTimeout,
	}

	client := resty.NewWithClient(&http.Client{
		Transport: transport,
		Timeout:   Timeout,
	})

	client.
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json")

	return client
}
