package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/a1comms/gcp-iap-auth/jwt"
)

type userIdentity struct {
	Subject string `json:"sub,omitempty"`
	Email   string `json:"email,omitempty"`
}

func authHandler(res http.ResponseWriter, req *http.Request) {
	claims, err := jwt.RequestClaims(req, cfg)
	if err != nil {
		if claims == nil || len(claims.Email) == 0 {
			log.Printf("Failed to authenticate (%v)\n", err)
		} else {
			log.Printf("Failed to authenticate %q (%v)\n", claims.Email, err)
		}
		res.WriteHeader(http.StatusUnauthorized)
		return
	}
	user := &userIdentity{
		Subject: claims.Subject,
		Email:   claims.Email,
	}
	expiresAt := time.Unix(claims.ExpiresAt, 0).UTC()
	log.Printf("Authenticated %q (token expires at %v)\n", user.Email, expiresAt)
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(user)
}

func authAudHandler(res http.ResponseWriter, req *http.Request) {
	aud := req.Header.Get("X-GCP-IAP-AUD")
	if aud == "" {
		log.Printf("Audience header not found")
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	reqCfg, err := initRequestCfg(aud)
	if err != nil {
		log.Printf("Request configuration error: %s", err)
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims, err := jwt.RequestClaims(req, reqCfg)
	if err != nil {
		if claims == nil || len(claims.Email) == 0 {
			log.Printf("Failed to authenticate (%v)\n", err)
		} else {
			log.Printf("Failed to authenticate %q (%v)\n", claims.Email, err)
		}
		res.WriteHeader(http.StatusUnauthorized)
		return
	}
	user := &userIdentity{
		Subject: claims.Subject,
		Email:   claims.Email,
	}
	expiresAt := time.Unix(claims.ExpiresAt, 0).UTC()
	log.Printf("Authenticated %q (token expires at %v)\n", user.Email, expiresAt)
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(user)
}
