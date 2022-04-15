package tools

import (
	"os"
	"time"

	"github.com/BurntSushi/toml"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//总配置
type tomlConfig struct {
	Title     string
	LogConfig logConfig `toml:"Log"`
}

type logConfig struct {
	Debug       bool   `toml:"debug"`       //是否开启调试
	MaxSize     int    `toml:"maxSize"`     //日志文件最大多少兆
	MaxDays     int    `toml:"maxDays"`     //日志文件保留天数
	MaxBackups  int    `toml:"maxBackups"`  //保留文件数
	FileName    string `toml:"fileName"`    //日志名字
	Compress    bool   `toml:"compress"`    //日志生成压缩包,大幅降低磁盘空间,必要时使用
	RotateByDay bool   `toml:"rotateByDay"` //每天轮转一次,如果开启,maxBackups的值需要>=maxDays
}

var (
	Log        *zap.Logger
	TomlConfig *tomlConfig
)

func init() {
	cfgfile := "config/log.toml"
	TomlConfig = new(tomlConfig)
	_, err := toml.DecodeFile(cfgfile, TomlConfig)
	if err != nil {
		TomlConfig.LogConfig = *defaultConfig()
	}

	//加入日期轮转
	l := &lumberjack.Logger{
		Filename:   TomlConfig.LogConfig.FileName,
		MaxSize:    TomlConfig.LogConfig.MaxSize,    // megabytes 兆字节
		MaxBackups: TomlConfig.LogConfig.MaxBackups, //保留文件数
		MaxAge:     TomlConfig.LogConfig.MaxDays,    // days
		LocalTime:  true,
		Compress:   TomlConfig.LogConfig.Compress, //日志生成压缩包,大幅降低磁盘空间,必要时使用
	}

	//开启24小时轮转一次
	if TomlConfig.LogConfig.RotateByDay {
		l.Rotate()
		go func() {
			rotateStartTime := time.Now().Format("2006-01-02") + " 23:59:59"
			rotateLeftTime, _ := time.ParseInLocation("2006-01-02 15:04:05", rotateStartTime, time.Local)
			remainSecond := time.Duration(rotateLeftTime.Unix() - time.Now().Unix() + 1)
			for {
				<-time.After(time.Second * remainSecond)
				l.Rotate()
				remainSecond = 24 * 60 * 60
			}
		}()
	}

	//写入磁盘
	w := zapcore.AddSync(l)

	//初始化encoder配置
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeCaller = shortCallerEncoder //日志调用方法
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	var core zapcore.Core
	//调试模式
	if TomlConfig.LogConfig.Debug {
		consoleErrors := zapcore.Lock(os.Stderr)
		core = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), consoleErrors, zap.DebugLevel),
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), w, zap.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderCfg), // zapcore.NewJSONEncoder(encoderCfg) //json格式
			w,
			zap.DebugLevel,
		)
	}

	Log = zap.New(core, zap.AddCaller())
	// defer Log.Sync() //main函数退出时使用
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
}

func shortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// TODO: consider using a byte-oriented API to save an allocation.
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

func defaultConfig() *logConfig {
	return &logConfig{
		Debug:       true,              //是否开启调试
		MaxSize:     1024,              //日志文件最大多少兆
		MaxDays:     29,                //日志文件保留天数
		MaxBackups:  30,                //保留文件数
		FileName:    "log/default.log", //日志名字
		Compress:    true,              //日志生成压缩包,大幅降低磁盘空间,必要时使用
		RotateByDay: false,             //每天轮转一次,如果开启,maxBackups的值需要>=maxDays
	}
}
