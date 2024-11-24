package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 定义全局变量, 用来保存配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port int `mapstructure:"port"`

	*LogConfig `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
	MaxSize int `mapstructure:"maxsize"`
	MaxAge int `mapstructure:"maxage"`
	MaxBackups int `mapstructure:"maxbackups"`
}

type MySQLConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	User string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName string `mapstructure:"db_name"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db int `mapstructure:"db"`
	PoolSize int `mapstructure:"pool_size"`
}

func Init() error {
	viper.SetConfigName("../config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.readinconfig() error:%v\n", err)
		return err
	}
	// 将读取到的配置信息保存(反序列化)到全局变量Conf中
	err = viper.Unmarshal(Conf)
	if err != nil {
		fmt.Printf("viper.unmarshal() error:%v\n", err)
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已经修改")
		err = viper.Unmarshal(Conf)
		if err != nil {
			fmt.Printf("viper.unmarshal() error:%v\n", err)
		}
	})

	return err
}
