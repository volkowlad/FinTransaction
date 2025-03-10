# FinTransaction

## Выполнен учебный REST API проект FinTransaction на Go
### В проекте разобраны следующие концепции:
- Разработка Веб-Приложений на Go, следуя дизайну REST API.
- Работа с фреймворком gin.
- Подход Чистой Архитектуры в построении приложения. Техника внедрения зависимостей.
- Работа с БД PostgreSQL: выполнены миграции с помощью migrate, запуск, используя Docker и Docker Compose.
- Для работы с БД ипользовалась стандартная библиотека database/sql. Выполнение SQL запросов, выполнение транзакций, отслеживание операций пользователя.
- Регистрация и аутентификация. Работа с JWT. 
- Ограничение количества запросов по одному кошельку.
- Работа с переменными окружения
- Работа со swagger
- Сделан Makefile для удобной работы с проектом
- Graceful Shutdown

### Для запуска приложения:
```
make run
```
### Для остановки приложения:
```
make down
```