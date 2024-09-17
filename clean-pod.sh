# https://qiita.com/tom_negocia/items/d1302cc3ad4657ad3466
# https://qiita.com/reoring/items/7252556bb216e242cddf
#kubectl delete pods {$pod_name} --grace-period=0 --force
# kubectl get namespace
kubectl -n lemontea get pods

# 作成済みpod表示
#kubectl get pods
kubectl delete namespace sample


# https://qiita.com/oguogura/items/fee90b7b2ce78413a518
# namespace指定して全て削除
kubectl get pod -n lemontea |awk '/cockroachdb-*/{print $1}' |xargs kubectl delete pod -n lemontea --grace-period=0 --force

#https://qiita.com/reoring/items/7252556bb216e242cddf
# 強制削除
# kubectl delete pods {$pod_name} --grace-period=0 --force
# これで全部削除できた
kubectl get pod -n lemontea |awk '/cockroach*/{print $1}' |xargs kubectl delete pod -n lemontea --grace-period=0 --force 

