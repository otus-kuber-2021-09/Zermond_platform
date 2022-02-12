# Zermond_platform
Zermond Platform repository


# Выполнено ДЗ № 2 (kubernetes-controllers)

>Руководствуясь материалами лекции опишите произошедшую ситуацию, почему обновление ReplicaSet не повлекло обновление 
запущенных pod?

ReplicaSet не поддерживает обновление. И пока состояние реплика сет удовлетворяет настройкам
(запущено указанное кол-во реплик, то ничего не происходит). Для обновления нужно удалить старые поды вручную,
тогда состояние репликисет будет не желаемым и репликасет заново поднимет поды с указанными контейнерами (уже с новым образом)
Для разворачивания приложений нужно использовать Deployment.

(Выписка из документации:
Once the original is deleted, you can create a new ReplicaSet to replace it. 
As long as the old and new .spec.selector are the same, then the new one will adopt the old Pods. 
However, it will not make any effort to make existing Pods match a new, different pod template. 
To update Pods to a new spec in a controlled way, use a Deployment, 
as ReplicaSets do not support a rolling update directly.)

> Основное ДЗ:
> 
Описаны манифесты (deployment, replicaset) для приложений hipster-frontend, hipster-paymentservice


# ДЗ со звездочкой

Описаны манифесты для разворачивания paymentservice по стратегиям: blue-green, reverse rolling update

# ДЗ со звездочкой

Описан манифест для node-exporter, который поднимается как DaemonSet (на каждой ноде, включая мастер-ноду)