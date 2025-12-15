package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lin-snow/ech0/internal/i18n"
)

// I18n 国际化中间件
func I18n() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		// 如果语言不支持，则使用默认语言
		if !i18n.IsLanguageSupported(lang) {
			lang = i18n.DefaultLanguage.String()
		}
		c.Set("lang", lang)
		c.Next()
	}
}
