# GolangUnitedSchool

Сервис сбора информации о студентах: 
1. База данных, хранящая всю имеющуюся информацию о студентах 
2. Парсеры, которые могут дополнить информацию в базе данными, полученными из файла 
3. API, на который можно постучаться чтобы получить данные из базы Какая инфа имеется по студентам: 
    - ФИО (русский, английский, вероятно еще какие то языки) 
    - Контакты (email, telegram, discord, …. ) 
    - Набор текстовых полей с описанием чего либо (опыт работы, самопрезентация, заметки ментора/наставника etc) 
    - Данные о том на каком потоке учился студент (группа, даты etc), тут можно прикрутить ссылки на видеозаписи уроков 
    - Текущий статус по выполнению домашних заданий, общий результат 
    - Описание дипломного проекта, команда в которой его выполняет 
    - Данные о выданном сертификате (скан, даты выдачи etc) 
    - Результаты прохождения интервью (интервьюер, оценка, заметки, дата/время) 
    - Сводные данные по интервью по всему потоку (расчет рейтинга по формуле + исходные данные) 

Часть полей могут быть обязательными для заполнения Должна быть реализована ролевая модель доступа к данным с использованием токенов: 
- Студент 
- Админ/ментор курса 
- Админ с доступом ко всем данным 
- Read-only роли по всем доступам перечисленным выше 
  Реализовать логирование всех действий пользователя и доступ к просмотру логов для админов системы 
  Реализовать «мягкое удаление» записей из БД (записи не стираем, помечаем как неактивные и скрываем)

Доработки по модели:
1. у каждого пользователя должен быть приоритетный канал связи
2. у сущностей группа нужно предусмотреть групповой канал связи (например общий канал в телеге, в дискорде и тд)

* data model : https://dbdiagram.io/d/635e65665170fb6441c0caec
* workflow : https://docs.google.com/spreadsheets/d/1L_RyOd7D5ubmhDDY0Gzku0tMPRrMKykhcabfBkgHk-c/edit#gid=443895109
* Documentation Swagger
* install  
  go get -u github.com/swaggo/swag/cmd/swag
  go get -u github.com/swaggo/gin-swagger
  go get -u github.com/swaggo/files
* make swagger docs 
  swag init -g cmd/missionPossible/main.go 
* to see swagger documentations 
  go run cmd/missionPossible/main.go 
  [GET] http://localhost:8080/docs/index.html
* to get offline html analog of swagger:
    - npm install -g redoc-cli
    - in docs dir: redoc-cli bundle -o index.html swagger.json
* to build and upload to AWS 
    - GOOS=linux GOARCH=amd64 go build -o bin/application cmd/missionPossible/main.go
    - 7z a eb.zip bin/ # we deploy just binary 
     AWS 
    - then upload zip archive into AWS elastic beanstalk # 
