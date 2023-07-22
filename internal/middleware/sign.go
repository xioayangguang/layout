package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"io/ioutil"
	"layout/global"
	response2 "layout/internal/response"
	"layout/pkg/helper/md5"
	"sort"
)

func Sign(signSalt string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sign := c.Request.Header.Get("sign")
		if global.Config.Debug && sign == "debug-mode-ignore" {
			c.Next()
			return
		}
		if sign != "" {
			paramsMap := map[string]interface{}{}
			data, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			_ = json.Unmarshal(data, &paramsMap)
			queryMap := c.Request.URL.Query()
			for k := range queryMap {
				paramsMap[k] = queryMap[k][0]
			}
			//fmt.Println(paramsMap)
			//fmt.Println(paramsBuild(paramsMap))
			//fmt.Println(signSalt)
			//fmt.Println(c.Request.Header.Get("ApiAuth"))
			//fmt.Println(paramsBuild(paramsMap) + signSalt + c.Request.Header.Get("ApiAuth"))
			//fmt.Println(calculate(paramsBuild(paramsMap), signSalt+c.Request.Header.Get("ApiAuth")))
			//fmt.Println(sign)
			if sign == calculate(paramsBuild(paramsMap), signSalt+c.Request.Header.Get("ApiAuth")) {
				c.Next()
				return
			}
		}
		response2.FailWithCode(c, response2.SignError)
		c.Abort()
		return
	}
}

func paramsBuild(params map[string]interface{}) string {
	var dataParams string
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if data, ok := params[k].([]interface{}); ok {
			dataParams = dataParams + k + "=" + fmt.Sprintf("%v", data) + "&"
		} else {
			dataParams = dataParams + k + "=" + cast.ToString(params[k]) + "&"
		}
	}
	if dataParams == "" {
		return ""
	} else {
		return dataParams[0 : len(dataParams)-1]
	}
}

func calculate(params, signSalt string) string {
	//h := sha1.New()
	//h.Write([]byte(params + signSalt))
	//bs := h.Sum(nil)
	//return hex.EncodeToString(bs)
	return md5.Md5(params + signSalt)
}
