package logger

import (
	"io"
	"os"

	"cufoon.litkeep.service/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() func() {
	cg, err := conf.NewConf("../../dev.yaml")
	if err != nil {
		panic(err)
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.Level(cg.LogZap.Level))
	var writer zapcore.WriteSyncer
	if cg.LogZap.Output == 1 {
		writer = zapcore.AddSync(os.Stdout)
	} else if cg.LogZap.Output == 2 {
		writer = zapcore.AddSync(getLogWriter())
	} else {
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(getLogWriter()), os.Stdout)
	}

	core := zapcore.NewCore(
		getEncoder(),
		writer,
		atomicLevel,
	)
	var tempLogger *zap.Logger
	if cg.Mode == "Debug" {
		tempLogger = zap.New(core, zap.AddCaller(), zap.Development())
	} else {
		tempLogger = zap.New(core)
	}

	perLog := zap.ReplaceGlobals(tempLogger)
	return perLog
}

func getLogWriter() io.Writer {
	cg, err := conf.NewConf("../../dev.yaml")
	if err != nil {
		panic(err)
	}
	c := cg.LogFile

	hook := &lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
	}
	return hook
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}
