package main

import "time"

// Request describes a login request.
type Request struct {
	Token     string
	Email     string
	IPAddress string
	CreatedAt time.Time
	Expiry    time.Time
}

// CreateRequest creates a new login request.
// TODO: Add bruteforce protection, to disallow spamming e-mails to a client.
// TODO: Add a mechanism to prevent phishing (e.g. only allow requests from the originating IP).
// TODO: Add interface to allow external storage, i.e. Redis or a RMDBS.
// func CreateRequest(email, ipAddress string) {
// 	token, err := GenerateRandomHash(64)
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	now := time.Now()
// 	expiry := now.Add(time.Minute * time.Duration(configuration.Config.Expiry))
//
// 	r := Request{Token: token, Email: email, IPAddress: ipAddress, CreatedAt: now, Expiry: expiry}
//
// 	// Save the login request to the internal database with login requests (for validation/expiry purposes)
//
// 	// Add the login request to the mailer queue.
// }

// ValidateRequest validates that the login request exists.
func ValidateRequest() {
	// Original IP must match the current IP. If not, log an error.
}

// RevokeRequest revokes the login request.
func RevokeRequest() {
	// Revokation can happen when:
	// a) the login request has met its expiry timer
	// b) all login requests need to be invalidated
}
