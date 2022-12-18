package conf

import (
	"fmt"
	"go-mall/dao"
	"gopkg.in/ini.v1"
	"strings"
)

// 读取配置文件

// 定义全局变量
var (
	AppModel string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host        string
	ProductPath string
	AvatarPath  string
)

// Init 初始化配置
func Init() {
	// 本地读取环境变量  注意：此路径是从main.go开始算起的路径
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		// 处理读取配置文件异常
		// panic 直译为 运行时恐慌 当panic被抛出异常后，如果我们没有在程序中添加任何保护措施的话，程序就会打印出panic的详细情况之后，终止运行
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}

	//读取配置
	LoadServer(file)
	LoadMySql(file)
	LoadEmail(file)
	LoadPhotoPath(file)

	//mysql读写分离配置
	// 读 (80%) 主库
	pathRead := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	// 写 (20%) 从库
	pathWrite := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")

	dao.Database(pathRead, pathWrite)
}

// LoadEmail 加载邮箱配置信息
func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
}

// LoadPhotoPath 加载图片路径信息
func LoadPhotoPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}

// LoadMySql 加载数据库配置信息
func LoadMySql(file *ini.File) {
	Db = file.Section("mysql").Key("DB").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

// LoadServer 加载服务配置信息
func LoadServer(file *ini.File) {
	// 将读取进来的配置字符串化
	AppModel = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}
