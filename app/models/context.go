package models

type ContextKey string

const (
	RequestStartedAtKey ContextKey = "request-started-at"
	RequestAuthToken    ContextKey = "request-auth-token"
)
