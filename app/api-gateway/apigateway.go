package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Обработчик для проксирования алертов к AlertService
func proxyAlertsToAlertService(alertServiceURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Создание нового запроса с использованием метода и тела исходного запроса
		req, err := http.NewRequest(c.Request.Method, alertServiceURL, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		// Копирование заголовков из исходного запроса
		copyHeaders(req.Header, c.Request.Header)

		// Установка необходимых заголовков для AlertService 
		req.Header.Set("X-Forwarded-For", c.ClientIP()) // Передача IP клиента
		req.Header.Set("Content-Type", "application/json") // Явное указание типа контента

		// Отправка запроса к AlertService
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to proxy request to AlertService"})
			return
		}
		defer resp.Body.Close()

		// Передача ответа от AlertService в vmalert
		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	}
}

/* Копирование заголовков */
func copyHeaders(dst, src http.Header) {
	for k, v := range src {
		dst[k] = v
	}
}

