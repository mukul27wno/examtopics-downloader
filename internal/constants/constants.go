package constants

import "time"

// Request behaviour
const HttpTimeout = 20 * time.Second
const MaxConcurrentRequests = 15
const RequestsPerSecond = 0.5
const MaxRetries = 5

// Backoff configuration
const InitalBackoff = time.Second
const BackoffFactor = 2.0

// HTTP Transport Tuning (in http client)
const MaxIdleConns = 100
const MaxIdleConnsPerHost = 100
const MaxConnsPerHost = 100

// Connection Timeouts (also in http client)
const IdleConnTimeout = 90 * time.Second
const TLSHandshakeTimeout = 10 * time.Second
const ResponseHeaderTimeout = 10 * time.Second
const ExpectContinueTimeout = 1 * time.Second
