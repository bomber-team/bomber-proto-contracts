# Спека по маршрутизации

## Топики на бомберах

1. bombers.tasks.$id - где $id - идентификатор бомбера, выдаваемый системой. На входе контракт: `contracts/rest_contracts/task.proto`

2. bombers.trigger.$id - отправка команды на запуск/остановку задачи. в теле отправляем идентификатор задачи. На входе контракт `contracts/system/task_action.proto`

## Топики на Bomber-Server

1. bomber.status - топик, в который отправляются текущие статусы конкретных бомберов. Контракт на входе: `contracts/system/status_changer.proto`

2. bomber.result - топик, в который отправляются результаты выполнения задач. Контракт на входе: `contracts/rest_contracts/result.proto`
