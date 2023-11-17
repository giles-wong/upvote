package logger

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"path"
	"runtime/debug"
	"time"

	"giles.wang/upvote/config"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

type Options struct {
	Path  string //日志记录路径
	Level string //日志级别
}

var (
	maxSize    int  = 200
	maxBackups int  = 20
	maxAge     int  = 60
	compress   bool = true
)

func InitLogger(cfg *config.LogConfig) {
	if cfg.Path == "" {
		cfg.Path = "./runtime/logs"
	}
	if cfg.Level == "" {
		cfg.Path = "debug"
	}
	maxSize = cfg.MaxSize
	maxBackups = cfg.MaxBackups
	maxAge = cfg.MaxAge
	compress = cfg.Compress

	logFile := path.Join(cfg.Path, "runtime.log") //正常日志记录信息
	errFile := path.Join(cfg.Path, "error.log")   //错误日志记录信息
	encoder := getEncoder()

	teeCore := zapcore.NewTee(
		zapcore.NewCore(encoder, getLogWriter(logFile), level(cfg.Level)),
		zapcore.NewCore(encoder, getLogWriter(errFile), zapcore.ErrorLevel),
	)

	logger = zap.New(teeCore, zap.AddCaller())
	zap.ReplaceGlobals(logger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	defer logger.Sync()
}

func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05"))
	}
	encodeConfig.TimeKey = "time"
	return zapcore.NewJSONEncoder(encodeConfig)
}

func getLogWriter(file string) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   file,       //日志记录的文件路径
		MaxSize:    maxSize,    //当个文件最大空间  单位M
		MaxBackups: maxBackups, //最大保存数量
		MaxAge:     maxAge,     //最大备份天数
		Compress:   compress,   //是否启用压缩
	}

	return zapcore.AddSync(lumberjackLogger)
}

func level(levelStr string) zapcore.Level {

	switch levelStr {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func WithConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				logger.Error("[Recovery from panic]",
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)

				logger.Error(c.Request.URL.Path,
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
				)
				// If the connection is dead, we can't write a status to it.
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  fmt.Sprintf("%v", err),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
