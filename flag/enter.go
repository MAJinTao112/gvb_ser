package flag

import (
	sys_flag "flag"
	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string // -u admin -u user
	ES   string // -es creat -es creat
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	es := sys_flag.String("es", "", "创建索引")
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
		ES:   *es,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}

func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}

	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	if option.ES == "create" {
		EsCreateIndex()
	}
	//sys_flag.Usage()

}
