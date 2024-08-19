[![build](https://github.com/Mart1n-Eden/house-service/actions/workflows/github-ci.yml/badge.svg)](https://github.com/Mart1n-Eden/house-service/actions/workflows/github-ci.yml)
# Тестовое задание для отбора на Avito Backend Bootcamp
## house-service

## Описание
На Авито ежедневно публикуются тысячи объявлений о продаже или аренде недвижимости. Они попадают в каталог домов, в котором пользователь может выбрать жильё по нужным параметрам в понравившемся доме.

Прежде чем попасть в каталог, объявление проходит тщательную модерацию, чтобы в нём не было недопустимого контента.
https://github.com/avito-tech/backend-bootcamp-assignment-2024

## Начало работы
### Установка
Клонирование репозитория
```sh
git clone https://github.com/Mart1n-Eden/house-service
```

### Запуск сервиса
Запускаем контейнер с помощью Makefile
```sh
make run
```

### Выполненные дополнительные задания
1. Реализована пользовательская авторизация по методам /register и /login.
2. Реализован асинхронный механизм уведомления пользователя о появлении новых квартир в доме по почте, метод /house/{id}/subscribe.
3. Настроен CI через github actions:
    1. в README.md корня вашего репозитория отображён бейдж со статусом билда.
4. Настроен логгер

## Дополнения к решению

- Если жильё успешно создано через endpoint `/flat/create`, то объявление получает статус модерации created. 
Из-за того что с квартирой может работать только один модератор и перед началом работы нужно перевести квартиру в статус "on moderation",
то я реализовал такой функционал: Когда обновляется статус квартиры на "on moderation", в таблице квартир в базе данных добавляется ID модератора в поле "updated_by". 
После этого статус этой квартиры может поменять только этот модератор. Также нельзя изменить статус квартиры на "approved" или "declined" если она не находится в статусе "on moderation",
или на статус "on moderation" если квартира уже прошла модерацию, то есть имеет статус "approved" или "declined".
- Подписка на уведомление о новых квартирах по номеру дома. Я реализовал такой вариант решения: Все подписки приходящие с помощью метода /house/{id}/subscribe
сохраняются в отдельную таблицу в базе данных. При запуске программы запускается горутина с тикером, и каждый 12 часов(можно настроить) сервис проверяет по всем имеющемся
подпискам наличие новых квартир в определенном доме. Для этого в таблице с квартирами добавлено поле "updated_at", значение которого изменяется про модерации квартиры. По подписке 
отправляются только квартиры со статусом "approved". Сообщение имеет такой вид:
`New flats in house 1 :
FlatId 2, Rooms 4, Price 10000
FlatId 3, Rooms 2, Price 5000
`
- Для быстрого отклика по endpoint по получения квартир в доме для пользователей я реализовал кэширование списка квартир в доме.
Кэш заполняется только квартирами со статусом "approved", то есть исключительно для пользователей. Из кэша удаляются списки квартир по определенному дому
в случае если в этом доме прошла модерацию какая-либо квартира и получила статус "approved" 
