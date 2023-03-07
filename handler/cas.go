package handler

import (
	"errors"
	"gin-cas/config"
	"github.com/gin-gonic/gin"
	"net/url"
)

func Logout(c *gin.Context) {
	c.Redirect(302, config.C.Cas.ServerLogoutUrl)
}

func CallBack(c *gin.Context) {
	to, ok := c.GetQuery("url")
	if ok == false {
		c.AbortWithError(403, errors.New("没有url参数"))
	}
	u, err := url.Parse(config.C.Cas.ClientHostUrl)
	if err != nil {
		panic(err)
	}
	u.Path = to
	c.Redirect(302, u.String())
}

func Login(c *gin.Context) {
	to, ok := c.GetQuery("url")
	if ok == false {
		c.AbortWithError(403, errors.New("没有url参数"))
	}
	callBackUrl, err := url.Parse(config.C.Cas.ClientHostUrl)
	if err != nil {
		panic(err)
	}
	callBackUrl.Path = "/api/cas/call-back"
	callBackValues := url.Values{}
	callBackValues.Add("url", to)
	callBackUrl.RawQuery = callBackValues.Encode()

	redirUrl, err := url.Parse(config.C.Cas.ServerLoginUrl)
	if err != nil {
		panic(err)
	}
	redirValues := url.Values{}
	callBackValues.Add("service", callBackUrl.String())
	redirUrl.RawQuery = redirValues.Encode()
	c.Redirect(302, redirUrl.String())
}

func CasInfo(c *gin.Context) {
	c.Redirect(302, config.C.Cas.ServerLogoutUrl)
}
