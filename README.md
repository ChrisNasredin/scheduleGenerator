# Schedule Generator

## Описание

Микросервис генерирует расписание для официантов.

## Формат

Принимает на вход JSON в формате

```JSON
{
  "dateFrom": "2026.12.01",
  "dateTo": "2026.12.31",
  "employees": [
    {
      "name": "Some Waiter 1",
      "preferedDates": ["2026.12.05", "2026.12.13"],
      "rejectedDates": ["2026.12.02", "2026.12.16"]
    }
  ]
}
```

На выходе выдает JSON:

```JSON
{
  "month": "2026-12",
  "shifts": [
    {
      "date": "2026-12-01",
      "employees": ["Some Waiter 1", "Some Waiter2"],
      "status": "dayoff"
    }
    {
      "date": "2026-12-02",
      "employees": ["Some Waiter 1", "Some Waiter2"],
      "status": "workday"
    }
  ]
}
```


## ToDo:

* Добавить роли
* добавить параметр "strongHandsRequiered" и сотрудникам добавить параметр strongHand
* Переход на ID
* Написать транспорт GRPC