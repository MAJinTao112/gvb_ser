package image_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/service/image_ser"

	//qiniiu "gvb_server/plugins/qiniu"

	"io/fs"
	"os"
)

// 白名单
var (
	WhiteImageList = []string{
		// 常见位图格式
		"jpg", "jpeg", "png", "gif", "bmp",
		// 新一代高效图片格式
		"webp", "heic", "heif", "avif",
		// 矢量图
		"svg",
		// 专业格式
		"tiff", "tif", "ico", "raw", "psd",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  //文件名
	Msg       string `json:"msg"`        //消息
	IsSuccess bool   `json:"is_success"` //是否上传成功
}

// 上传单个图片，返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}
	//判断路径是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 递归创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	var resList []image_ser.FileUploadResponse

	for _, file := range fileList {
		serviceRes := service.ServiceGroupApp.ImageService.ImageUploadService(file)
		//serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		// 成功的
		if !global.Config.QiNiu.Enable {
			// 本地还得保存一下
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
	}

	res.OkWithData(resList, c)
}
