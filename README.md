# Сервис онлайн библиотеки песен

# Использование

**Запуск**

Из корневой директории проекта:

`$ go run "app\cmd\main.go"`

При добавлении песни сервис делает запрос по адресу, заданному в файле конфигурации: "http://localhost:8081/info".

**Переменные окружения**

Все переменные окружения заданы в файле конфигурации .env в директории "app\cmd\configs". В случае невозможности прочтения данных из фалйа в качестве переменный окружения будут выбраны значения, пропасанные в main.go.

**БД**

Конфигурация БД описана в файле "\db\dbex.sql"