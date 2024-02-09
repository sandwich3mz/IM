package configs

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const DefaultConfigPath = "./config.yaml"

var Conf = new(Config)

type Config struct {
	API         `mapstructure:"api"`
	DBConfig    `mapstructure:"db"`
	RedisConfig `mapstructure:"redis"`
	EmailConfig `mapstructure:"email"`
	JwtConfig   `mapstructure:"jwt"`
}

type API struct {
	Port int `mapstructure:"port"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type EmailConfig struct {
	Email string `mapstructure:"email"`
	Code  string `mapstructure:"auth_code"`
}

type JwtConfig struct {
	Secret string `mapstructure:"secret"`
}

func init() {
	// 默认配置文件路径
	var configPath string
	flag.StringVar(&configPath, "configs", DefaultConfigPath, "配置文件绝对路径或相对路径")
	flag.Parse()

	logrus.Printf("===> configs path is: %s", configPath)
	// 初始化配置文件
	viper.SetConfigFile(configPath)
	viper.WatchConfig()
	// 观察配置文件变动
	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Printf("configs file has changed")
		if err := viper.Unmarshal(&Conf); err != nil {
			logrus.Errorf("failed at unmarshal configs file after change, err: %v", err)
			panic(fmt.Sprintf("failed at init configs: %v", err))
		}
	})
	// 将配置文件读入 viper
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("failed at ReadInConfig, err: %v", err)
		panic(fmt.Sprintf("failed at init configs: %v", err))
	}
	// 解析到变量中
	if err := viper.Unmarshal(&Conf); err != nil {
		logrus.Errorf("failed at Unmarshal configs file, err: %v", err)
		panic(fmt.Sprintf("failed at init configs: %v", err))
	}

	// 返回 nil 或错误
	logrus.Infoln("global configs init success...")
	logrus.Infof("%+v", Conf)
}
