package entity

import "time"

// Результат выполения http запроса
type RequestResult struct {
	StatusCode    int           // статус ответа
	ExecutionTime time.Duration // время выполнения запроса
	Error         error         // Ошибка
}
