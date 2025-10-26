package logger

import "context"

type ContextLogger interface {
	Infof(format string, args ...any)
	Errorf(format string, args ...any)
	Debugf(format string, args ...any)
	Warnf(format string, args ...any)

	Infow(format string, keysAndValues ...any)
	Errorw(format string, keysAndValues ...any)
	Debugw(format string, keysAndValues ...any)
	Warnw(format string, keysAndValues ...any)

	InfowCtx(ctx context.Context, msg string, keysAndValues ...any)
	ErrorwCtx(ctx context.Context, msg string, keysAndValues ...any)
	WarnwCtx(ctx context.Context, msg string, keysAndValues ...any)
	DebugwCtx(ctx context.Context, msg string, keysAndValues ...any)
}
