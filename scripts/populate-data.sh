## Civil War

### Iron Man Team
kubectl create ns ironman
kubectl create ns vision
kubectl create ns spider-man
kubectl create ns adhoc-team-1

### Captain America Team
kubectl create ns captain-america
kubectl create ns ant-man
kubectl create ns hawk-eye
kubectl create ns adhoc-team-2

### Giantswarm infra
kubectl create ns gswarm-critical
kubectl create ns gswarm-lambda

for i in {1..3}; do kubectl run ironman-$RANDOM --image nginx --namespace ironman; done
for i in {1..3}; do kubectl run vision-$RANDOM --image nginx --namespace vision; done
for i in {1..3}; do kubectl run spider-man-$RANDOM --image nginx --namespace spider-man; done
for i in {1..5}; do kubectl run adhoc-team-1-$RANDOM --image nginx --namespace adhoc-team-1; done
for i in {1..3}; do kubectl run captain-america-$RANDOM --image nginx --namespace captain-america; done
for i in {1..5}; do kubectl run adhoc-team-2-$RANDOM --image nginx --namespace adhoc-team-2; done
for i in {1..3}; do kubectl run hawk-eye-$RANDOM --image nginx --namespace hawk-eye; done
for i in {1..3}; do kubectl run ant-man-$RANDOM --image nginx --namespace ant-man; done
for i in {1..5}; do kubectl run gswarm-critical-$RANDOM --image nginx --namespace gswarm-critical; done
for i in {1..5}; do kubectl run gswarm-lambda-$RANDOM --image nginx --namespace gswarm-lambda; done
