package usecase

import "errors"

// Определение ошибок
var (
	ErrNotFound     = errors.New("not found")      // Ошибка, когда элемент не найден
	ErrInvalidInput = errors.New("invalid input")  // Ошибка, когда входные данные недействительны
	ErrInternal     = errors.New("internal error") // Ошибка, связанная с внутренними проблемами
	ErrUnauthorized = errors.New("unauthorized")   // Ошибка, когда доступ запрещен
	ErrConflict     = errors.New("conflict")       // Ошибка, когда возникает конфликт
)
