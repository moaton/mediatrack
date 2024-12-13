package application

import (
	"context"
	"mediatrack/config"
	"mediatrack/pkg/logger/zap"
	"path/filepath"

	"github.com/go-logr/logr"
	"go.uber.org/zap/zapcore"
)

type Application struct {
	ctx    context.Context
	cfg    *config.Config
	logger logr.Logger
}

func New(ctx context.Context, cfg *config.Config) *Application {
	return &Application{
		ctx: ctx,
		cfg: cfg,
	}
}

func (a *Application) InitLogger() {
	a.logger = zap.New(
		zap.Level(a.cfg.Level),
		zap.UseDevMode(true),
		zap.TimeEncoder(zapcore.ISO8601TimeEncoder),
		zap.ConsoleEncoder(
			func(ec *zapcore.EncoderConfig) { ec.EncodeLevel = zapcore.CapitalColorLevelEncoder },
			func(ec *zapcore.EncoderConfig) {
				ec.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
			},
			func(ec *zapcore.EncoderConfig) {
				ec.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
					encoder.AppendString(filepath.Base(caller.FullPath()))
				}
			},
		),
	)
}

func (a *Application) Run() {
	a.logger.Info("Application started")
}
