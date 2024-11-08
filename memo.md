

# podたち

`minikube`が作成するpodたち。
それぞれ、このような役目

| Tコンポーネント名 | 役割 |
| --- | ---- |
| kube apiserver | KubernetesのAPIを公開するコンポーネント。kubectlからのリソース操作を受け付ける |
| etcd | 高可用性を備えた分散キーバリューストアでKubernetesクラスタのバッキングストアとして利用される |
| kube scheduler | Nodeを監視し、コンテナを配置する最適なNodeを選択する |
| kube controller manager | リソースを制御するコントローラーを実行する |



```
pod name: etcd-minikube.
pod name: kube-apiserver-minikube.
pod name: kube-controller-manager-minikube.
pod name: kube-proxy-bqdbm.
pod name: kube-scheduler-minikube.
pod name: storage-provisioner.
```
