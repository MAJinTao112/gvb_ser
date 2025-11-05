package new_api

import (
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/utils/request"
	"io"
	"time"
)

type params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type header struct {
	SignatureKet string `form:"signature_ket" structs:"signature_ket"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"user_agent" structs:"user_agent"`
}

type NewData struct {
	Index    string `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hot_value"`
	Link     string `json:"link"`
}

type NewResponse struct {
	Code int       `json:"code"`
	Data []NewData `json:"data"`
	Msg  string    `json:"msg"`
}

const NewAPI = "https://api.codelife.cc//top/list"
const timeout = 2 * time.Second

func (NewApi) NewListView(c *gin.Context) {
	var cr params
	var headers header
	err := c.ShouldBindJSON(&cr)
	err = c.ShouldBindHeader(&headers)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	if cr.Size == 0 {
		cr.Size = 1
	}
	httpResponse, err := request.Post(NewAPI, cr, structs.Map(headers), timeout)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	var response NewResponse
	byteData, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	if response.Code != 200 {
		res.FailWithMessage(response.Msg, c)
		return
	}
	res.OkWithData(response.Data, c)
	return
}
