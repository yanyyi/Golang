package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(Config)

type Config struct {
	*AppConfig   `mapstructure:"app"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host              string `mapstructure:"host"`
	User              string `mapstructure:"user"`
	Password          string `mapstructure:"password"`
	Db                string `mapstructure:"db"`
	Port              int    `mapstructure:"port"`
	MaxOpenConnection int    `mapstructure:"max_open_connection"`
	MaxIdleConnection int    `mapstructure:"max_idle_connection"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {

	//读取配置文件
	//viper.SetConfigName("config") //配置文件名称
	//viper.SetConfigType("yaml")   //如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath(".")      //查找配置文件所在的路径
	viper.SetConfigFile("./conf/config.yaml") //以后直接用这种啦
	//还可以在工作目录中查找配置
	err = viper.ReadInConfig() //查找并读取配置文件
	if err != nil {
		//读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return

}
