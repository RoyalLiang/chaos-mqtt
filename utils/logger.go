package tools

import (
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"fms-awesome-tools/configs"
)

var (
	LogLevel = make(chan string, 1)
)

func setLogLevel(al zap.AtomicLevel, level string) {
	switch level {
	case "debug":
		al.SetLevel(zap.DebugLevel)
	case "info":
		al.SetLevel(zap.InfoLevel)
	case "warn":
		al.SetLevel(zap.WarnLevel)
	case "error":
		al.SetLevel(zap.ErrorLevel)
	default:
		al.SetLevel(zap.InfoLevel)
	}
}

func InitialLogger(root string, cfg configs.LoggerConfig) {

	autoLevel := zap.NewAtomicLevel()
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	var f string
	//file := viper.GetString("product.uuid") + ".log"
	if cfg.Dir != "" {
		f = filepath.Join(cfg.Dir, cfg.Name+".log")
	} else {
		f = filepath.Join(root, "logs", cfg.Name+".log")
	}
	//consoleDebugging := zapcore.Lock(os.Stdout)

	setLogLevel(autoLevel, cfg.Level)
	fileOutput := zapcore.AddSync(&lumberjack.Logger{
		Filename: f,
		MaxSize:  cfg.Size,
		MaxAge:   cfg.Backup,
	})

	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeTime = customTimeEncoder
	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, fileOutput, autoLevel),
		//zapcore.NewCore(encoder, consoleDebugging, autoLevel),
	)

	logger := zap.New(core)
	zap.ReplaceGlobals(logger)

	go func() {
		for {
			select {
			case nv := <-LogLevel:
				setLogLevel(autoLevel, nv)
			}
		}
	}()

	defer logger.Sync()
}
