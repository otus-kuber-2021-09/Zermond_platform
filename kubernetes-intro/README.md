# Zermond_platform
Zermond Platform repository


# Выполнено ДЗ № 1 (kubernetes-prepare)

Разберитесь почему все pod в namespace kube-system восстановились после удаления. Укажите причину в описании PR


## Решение
В ВМ (minikube) запущен демон kubelet.service, в конфиге `/var/lib/kubelet/config.yaml` которого указана директива `staticPodPath: /etc/kubernetes/manifests`
по адресу `/etc/kubernetes/manifests` лежат манифесты следующих объектов:

```
root@minikube:/var/lib/kubelet# ls -l /etc/kubernetes/manifests/ | awk '{print $9}'

etcd.yaml
kube-apiserver.yaml
kube-controller-manager.yaml
kube-scheduler.yaml
```
Их, собственно, говоря и запускает kubelet
А за `core-dns` следит уже контроллер `Controlled By: ReplicaSet/coredns-78fcd69978`

# Выполнено ДЗ № 1 (kubernetes-intro) со звездочкой

Выясните причину, по которой pod frontend находится в статусе Error Создайте новый манифест frontend-pod-healthy.yaml . При его
применении ошибка должна исчезнуть. Подсказки можно найти: В логах - kubectl logs frontend
В манифесте по ссылке
В результате, после применения исправленного манифеста pod frontend должен находиться в статусе Running (опустим вопрос,
действительно ли микросервис работает)
Поместите исправленный манифест frontend-pod-healthy.yaml в директорию kubernetes-intro

## Решение

У гошного приложения не хватало заданных переменных окружения
```shell
panic: environment variable "PRODUCT_CATALOG_SERVICE_ADDR" not set
```

Решение - задать недостающие переменные в манифест в блоке env
```yaml
      env:
        - name: PRODUCT_CATALOG_SERVICE_ADDR
          value: "PRODUCTSERVICE:4444"
        - name: CURRENCY_SERVICE_ADDR
          value: "CURRSERVICE:4444"
        - name: CART_SERVICE_ADDR
          value: "CARTSERVICE:4444"
        - name: RECOMMENDATION_SERVICE_ADDR
          value: "recommendationservice:8080"
        - name: SHIPPING_SERVICE_ADDR
          value: "shippingservice:50051"
        - name: CHECKOUT_SERVICE_ADDR
          value: "checkoutservice:5050"
        - name: AD_SERVICE_ADDR
          value: "adservice:9555"
```