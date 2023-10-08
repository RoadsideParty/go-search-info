package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

type Info struct {
	Phone        string `json:"phone"`
	QQ           string `json:"qq"`
	PhoneAddress string `json:"phoneAddress"`
	SinaId       string `json:"sinaId"`
	SinaUrl      string `json:"sinaUrl"`
}

func GetInfo(c *gin.Context) {
	qq := c.Query("qq")
	phone := c.Query("phone")
	sinaId := c.Query("sinaId")
	var info1 Info
	var info2 Info
	var info3 Info
	var info4 Info
	if qq != "" {
		info1 = getPhoneAndAddressByQQ(qq)
	}
	if phone != "" {
		info2 = getQQAndAddressByPhone(phone)
		info3 = getSinaAndAddressByPhone(phone)
	}
	if sinaId != "" {
		info4 = getPhoneAndAddressBySinaID(sinaId)
	}
	c.JSON(http.StatusOK, map[string]any{
		"data": []Info{
			info1, info2, info3, info4,
		},
	})
}

// 封装请求方法
func request(url string) []byte {
	client := &http.Client{Timeout: time.Second * 10}
	res, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return body
}

// https://zy.xywlapi.cc/qqapi?qq=
// QQ号获取手机号和归属地
func getPhoneAndAddressByQQ(qq string) Info {
	body := request("https://zy.xywlapi.cc/qqapi?qq=" + qq)
	V := struct {
		Phone     string `json:"phone"`
		QQ        string `json:"qq"`
		Phonediqu string `json:"phonediqu"`
	}{}
	if err := json.Unmarshal(body, &V); err != nil {
		fmt.Println(err.Error())
		return Info{}
	}
	info := Info{
		Phone:        V.Phone,
		QQ:           V.QQ,
		PhoneAddress: V.Phonediqu,
	}
	return info
}

// https://zy.xywlapi.cc/qqphone?phone=
// 手机号获取qq和归属地
func getQQAndAddressByPhone(phone string) Info {
	body := request("https://zy.xywlapi.cc/qqphone?phone=" + phone)
	V := struct {
		QQ        string `json:"qq"`
		Phonediqu string `json:"phonediqu"`
	}{}
	if err := json.Unmarshal(body, &V); err != nil {
		fmt.Println(err.Error())
		return Info{}
	}
	info := Info{
		Phone:        phone,
		QQ:           V.QQ,
		PhoneAddress: V.Phonediqu,
	}
	return info
}

// https://zy.xywlapi.cc/wbphone?phone=
// 手机号获取微博url和归属地
func getSinaAndAddressByPhone(phone string) Info {
	body := request("https://zy.xywlapi.cc/wbphone?phone=" + phone)
	V := struct {
		SinaId    string `json:"id"`
		Phonediqu string `json:"phonediqu"`
	}{}
	if err := json.Unmarshal(body, &V); err != nil {
		fmt.Println(err.Error())
		return Info{}
	}
	info := Info{
		Phone:        phone,
		PhoneAddress: V.Phonediqu,
		SinaId:       V.SinaId,
		SinaUrl:      "https://weibo.com/u/" + V.SinaId,
	}
	return info
}

// https://zy.xywlapi.cc/wbapi?id=
// 微博ID获取手机号和归属地
func getPhoneAndAddressBySinaID(sinaId string) Info {
	body := request("https://zy.xywlapi.cc/wbapi?id=" + sinaId)
	V := struct {
		Phone     string `json:"phone"`
		Phonediqu string `json:"phonediqu"`
	}{}
	if err := json.Unmarshal(body, &V); err != nil {
		fmt.Println(err.Error())
		return Info{}
	}
	info := Info{
		Phone:        V.Phone,
		PhoneAddress: V.Phonediqu,
		SinaId:       sinaId,
	}
	return info
}
