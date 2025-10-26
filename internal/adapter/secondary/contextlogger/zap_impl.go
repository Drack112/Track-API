package contextlogger

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/Drack112/Track-API/internal/shared/constants/ctxkeys"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	failedToFlushLogger = "Failed to flush context logger: %v"
)

type ZapLoggerContextual struct {
	base *zap.SugaredLogger
}

func New() (*ZapLoggerContextual, func()) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	encoder := zapcore.NewJSONEncoder(encoderCfg)

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel && lvl < zapcore.ErrorLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	infoWriter := zapcore.Lock(os.Stdout)
	errorWriter := zapcore.Lock(os.Stderr)

	infoCore := zapcore.NewCore(encoder, infoWriter, infoLevel)
	errorCore := zapcore.NewCore(encoder, errorWriter, errorLevel)

	tee := zapcore.NewTee(infoCore, errorCore)

	logger := zap.New(tee, zap.AddCaller(), zap.AddCallerSkip(1))
	sugar := logger.Sugar()

	cleanup := func() {
		if err := sugar.Sync(); err != nil {
			log.Printf(failedToFlushLogger, err)
		}
	}

	return &ZapLoggerContextual{base: sugar}, cleanup
}

func (l *ZapLoggerContextual) Infof(format string, args ...any) {
	l.base.Infof(format, args...)
}

func (l *ZapLoggerContextual) Errorf(format string, args ...any) {
	l.base.Errorf(format, args...)
}

func (l *ZapLoggerContextual) Debugf(format string, args ...any) {
	l.base.Debugf(format, args...)
}

func (l *ZapLoggerContextual) Warnf(format string, args ...any) {
	l.base.Warnf(format, args...)
}

func (l *ZapLoggerContextual) Infow(format string, keysAndValues ...any) {
	l.base.Infow(format, keysAndValues...)
}

func (l *ZapLoggerContextual) Errorw(format string, keysAndValues ...any) {
	l.base.Errorw(format, keysAndValues...)
}

func (l *ZapLoggerContextual) Debugw(format string, keysAndValues ...any) {
	l.base.Debugw(format, keysAndValues...)
}

func (l *ZapLoggerContextual) Warnw(format string, keysAndValues ...any) {
	l.base.Warnw(format, keysAndValues...)
}

func (l *ZapLoggerContextual) InfowCtx(ctx context.Context, msg string, keysAndValues ...any) {
	fields := EnrichFieldsFromContext(ctx)
	l.base.Infow(msg, append(fields, keysAndValues...)...)
}

func (l *ZapLoggerContextual) ErrorwCtx(ctx context.Context, msg string, keysAndValues ...any) {
	fields := EnrichFieldsFromContext(ctx)
	l.base.Errorw(msg, append(fields, keysAndValues...)...)
}

func (l *ZapLoggerContextual) DebugwCtx(ctx context.Context, msg string, keysAndValues ...any) {
	fields := EnrichFieldsFromContext(ctx)
	l.base.Debugw(msg, append(fields, keysAndValues...)...)
}

func (l *ZapLoggerContextual) WarnwCtx(ctx context.Context, msg string, keysAndValues ...any) {
	fields := EnrichFieldsFromContext(ctx)
	l.base.Warnw(msg, append(fields, keysAndValues...)...)
}

func EnrichFieldsFromContext(ctx context.Context) []any {
	var fields []any
	if reqID := GetRequestID(ctx); reqID != "" {
		fields = append(fields, string(ctxkeys.RequestID), reqID)
	}

	span := trace.SpanFromContext(ctx)
	sc := span.SpanContext()
	if sc.IsValid() {
		fields = append(fields, string(ctxkeys.TraceID), sc.TraceID().String())
		fields = append(fields, string(ctxkeys.SpanID), sc.SpanID().String())
	} else {
		if traceID := GetTraceID(ctx); traceID != "" {
			fields = append(fields, string(ctxkeys.TraceID), traceID)
		}
	}

	if userID := GetUserID(ctx); userID != "" {
		fields = append(fields, string(ctxkeys.UserID), userID)
	}
	return fields
}

func GetRequestID(ctx context.Context) string {
	if v := ctx.Value(ctxkeys.RequestID); v != nil {
		if s, ok := v.(string); ok {
			return s
		}

		// Fallback
		switch vv := v.(type) {
		case string:
			return vv
		case []byte:
			return string(vv)
		case int:
			return strconv.Itoa(vv)
		case int64:
			return strconv.FormatInt(vv, 10)
		case uint64:
			return strconv.FormatUint(vv, 10)
		}
	}

	return ""
}

func GetTraceID(ctx context.Context) string {
	if v := ctx.Value(ctxkeys.TraceID); v != nil {
		if s, ok := v.(string); ok {
			return s
		}

		// Fallback
		switch vv := v.(type) {
		case string:
			return vv
		case []byte:
			return string(vv)
		case int:
			return strconv.Itoa(vv)
		case int64:
			return strconv.FormatInt(vv, 10)
		case uint64:
			return strconv.FormatUint(vv, 10)
		}
	}

	return ""
}

func GetUserID(ctx context.Context) string {
	if v := ctx.Value(ctxkeys.UserID); v != nil {
		if s, ok := v.(string); ok {
			return s
		}

		// Fallback
		switch vv := v.(type) {
		case string:
			return vv
		case []byte:
			return string(vv)
		case int:
			return strconv.Itoa(vv)
		case int64:
			return strconv.FormatInt(vv, 10)
		case uint64:
			return strconv.FormatUint(vv, 10)
		}
	}

	return ""
}
