# Zermond_platform
Zermond Platform repository


## Выполнено ДЗ № 1 (kubernetes-prepare)

Разберитесь почему все pod в namespace kube-system восстановились после удаления. Укажите причину в описании PR

## В процессе сделано:

```
docker rm -f $(docker ps -a -q) (внутри ВМ minikube)
kubectl delete pod --all -n kube-system
```

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