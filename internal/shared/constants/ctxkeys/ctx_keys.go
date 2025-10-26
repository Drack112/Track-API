package ctxkeys

type contextKey string

const (
	UserID    contextKey = "user_id"
	Token     contextKey = "token"
	RequestID contextKey = "request_id"
	TraceID   contextKey = "trace_id"
	SpanID    contextKey = "span_id"
	Claims    contextKey = "claims"
)
