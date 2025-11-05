package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = iota + 1 //本地
	QiNiu                      //七牛云
)

func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	var str string
	switch s {
	case Local:
		str = "local"
	case QiNiu:
		str = "qiniu"
	default:
		str = "其他"
	}
	return str

}
