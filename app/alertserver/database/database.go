package database

import (
	"context"
	"log"

	"github.com/VictoriaMetrics/VictoriaMetrics/app/alertserver/config"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

// соединение с YDB
type DataBase struct {
	conf *config.Config
	conn ydb.Connection
	pool *table.SessionPool
}

func funcNewDataBase(cfg *config.Config) *DataBase {
	// подключение к YDB
	conn, err := ydb.New(context.Background(), ydb.WithConnectionString(cfg.YDB.ConnectionString))
	if err != nil {
		log.Fatalf("Подключение к YDB не успешно: %v", err)
	}

	// инициализация пула сессий (чтобы реальзовать многопоточность)
	pool := table.NewSessionPool(conn, table.WithSizeLimit(cfg.YDB.PoolSize))

	log.Println("Подключение к YDB успешно")
	return &DataBase{
		conf: cfg,
		conn: conn,
		pool: pool,
	}
}

/*
	Метод GetAlertStatus возвращает статус алерта по ID:

1. Подключается к базе через пул.
2. Выполняет YQL-запрос с использованием переданного alertID.
3. Проверяет, есть ли данные в результате выполнения запроса.
4. Если данные есть, извлекает их в переменную status.
5. Пока что: возвращает полученный статус или всегда nil, если данные не найдены. НУЖНО ПЕРЕДЕЛАТЬ: делать проверку.
- если получен несуществующий alertID, то возвращать уведомление "Некорректный alertID";
- если получен удаленный alertID, то выводить уведомление с причину удаления ("Данный alertID был удален администратором",
"Данный alertID был удален по причине: статус Отклонено").
6. Автоматически закрывает сессию после выполнения.
*/
func (db *DataBase) GetAlertStatus(ctx context.Context, alertID string) (string, error) {
	session, err := db.pool.Get(ctx)
	if err != nil {
		return "", err
	}
	defer session.Close(ctx)

	/* Запрос на YQL для YDB.
	   $alertID - переменная для передачи ID события алерта.
	   Берем столбец status из таблицы alert_statuses,
	   где значение в столбце alert_id равно переданныму $alertID. */
	query := `
		DECLARE $alertID AS Utf8;
		SELECT status
		FROM alert_statuses
		WHERE alert_id = $alertID;
	`
	/* В методе Execute выполняется запрос к базе в рамках сессии.
	   Передаются параметры table.NewQueryParameters для замены $alertID на значения alertID.

	   res — результат выполнения запроса.
	   err — ошибка при выполнении запроса.*/

	res, err := session.Execute(ctx, query, table.NewQueryParameters(
		table.ValueParam("$alertID", ydb.UTF8Value(alertID)),
	))
	if err != nil {
		return "", err
	}

	/* Метод NextResultSet позв. проверять наборы результатов ResultSet, если по данному алерту их несколько.
	   Если набор пустой (ничего не найдено, либо ошибка), то возвращает false */
	if !res.NextResultSet(ctx) {
		return "", nil
	}
	/* Метод NextRow позв. итерироваться по набору строк ResultSet, чтобы найти строку с status.
	   Если строк больше не найдено, возвращет false. Если строк нет (самого alert_id нет) -> пустая строка*/
	if !res.NextRow() {
		return "", nil
	}

	// статус события алерта (или инцидента, заведенного по этому событию???)
	var status string
	err = res.Scan(&status)
	return status, err
}

// Метод UpdateAlertStatus обновляет статус события алерта, добавляя отметку времени.
func (db *DataBase) UpdateAlertStatus(ctx context.Context, alertID string, status string) error {
	session, err := db.pool.Get(ctx)
	if err != nil {
		return err
	}
	defer session.Close(ctx)

	/* Запросом UPSERT добавляем новые алерты или обновляем существующие записи.
	   Если запись не существует, создаём новую строку с указанными значениями.
	   С помощью функции CurrentUtcDateTime фиксируем timestamp обновления статуса*/
	query := `
	DECLARE $alertID AS Utf8;
	DECLARE $status AS Utf8;

	UPSERT INTO alert_statuses (alert_id, status, updated_at)
	VALUES ($alertID, $status, CurrentUtcDateTime());
`
	/* С помощью table.ValueParam связываем имя переменной с её значением:
	   $alertID -> alertID;
	   $status -> status
	*/
	_, err = session.Execute(ctx, query, table.NewQueryParameters(
		table.ValueParam("$alertID", ydb.UTF8Value(alertID)),
		table.ValueParam("$status", ydb.UTF8Value(status)),
	))
	return err
}

// Метод Close закрывает соединение и очищает пул сессий, когда сервис завершает работу.
func (db *DataBase) Close() {
	db.pool.Close()
	db.conn.Close(context.Background())
	log.Println("Database connection closed")
}
