package v1

import (
	"app_jira/internal/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/gorilla/mux"
)

// @Summary Создание задачи для Jira
// @Description Создаёт задачу для Jira и вернёт uuid созданной задачи
// @Tags task
// @Accept json
// @Produce json
// @Param task body entity.TaskJira true "Json задачи для Jira"
// @Success 200 {object} uuid.UUID "uuid созданной в Jira задачи"
// @Failure 400 {string} string "Ошибка в теле запроса"
// @Failure 500 {string} string "Ошибка обработки алерта"
// @Router /api/v1/task [post]
func (s *Server) createTaskJira(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("createTaskJira")

	var task entity.TaskJira
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Ошибка в теле запроса: %v", err)))
		return
	}

	// Обрабатываем полученные алерты
	uuID, err := s.u.CreateTaskJira(r.Context(), task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Ошибка обработки алертов: %v", err)))
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(uuID)
}

// @Summary Получение всех задач
// @Description Вернет массив всех задач, отсортированных по дате
// @Tags task
// @Produce json
// @Success 200 {array} entity.TaskJira
// @Failure 500 {string} string "Ошибка получения данных"
// @Router /api/v1/task [get]
func (s *Server) getAllTashJira(w http.ResponseWriter, r *http.Request) {
	tasks, err := s.u.GetAllTaskJira(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Ошибка получения данных: %v.", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// @Summary Получение списка задач
// @Description Вернет список задач отсортированных по дате с begin в размере count
// @Tags task
// @Produce json
// @Param begin query int true "Начальный индекс"
// @Param count query int true "Колличество получаемых задач"
// @Success 200 {array} entity.TaskJira
// @Failure 400 {string} string "Недопустимые параметры"
// @Failure 500 {string} string "Ошибка получения данных"
// @Router /api/v1/task/list [get]
func (s *Server) getListTasksJira(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	beginStr := vars["begin"]                        // Получаем alert_id из параметров
	begin, err := strconv.ParseInt(beginStr, 10, 64) // Преобразуем в int64
	if err != nil || begin < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Недопустимый параметр `begin` (begin >= 0): %v.", err)))
		return
	}

	countStr := vars["count"]                        // Получаем alert_id из параметров
	count, err := strconv.ParseInt(countStr, 10, 64) // Преобразуем в int64
	if err != nil || count <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Недопустимый параметр `count` (count > 0): %v.", err)))
		return
	}

	tasks, err := s.u.GetListTaskJira(r.Context(), begin, count)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Ошибка получения данных: %v.", err)))
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// @Summary Получение задачи по UUID
// @Description Вернет задачу по указанному UUID
// @Tags task
// @Produce json
// @Param task_uuid path int true "UUID задачи"
// @Success 200 {object} entity.TaskJira
// @Failure 400 {string} string "Недопустимые параметры"
// @Failure 500 {string} string "Ошибка получения данных"
// @Router /api/v1/task/{alert_id} [get]
func (s *Server) getTaskJiraByID(w http.ResponseWriter, r *http.Request) {

	uuId := uuid.UUID{}

	// Вызов метода GetAlert
	task, err := s.u.GetTaskJiraByUUID(r.Context(), uuId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Ошибка получения данных: %v.", err)))
		return
	}

	// Установка заголовка Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Возврат JSON
	if err := json.NewEncoder(w).Encode(task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Failed to encode response: %v.", err)))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
