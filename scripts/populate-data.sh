## Civil War

### Iron Man Team
kubectl create ns ironman
kubectl create ns vision
kubectl create ns spiderman
kubectl create ns adhoc-team-1

### Captain America Team
kubectl create ns captainamerica
kubectl create ns antman
kubectl create ns hawkeye
kubectl create ns adhoc-team-2

for i in {1..3}; do kubectl run ironman-$RANDOM --image nginx --namespace ironman; done
for i in {1..3}; do kubectl run vision-$RANDOM --image nginx --namespace vision; done
for i in {1..3}; do kubectl run spiderman-$RANDOM --image nginx --namespace spiderman; done
for i in {1..3}; do kubectl run adhoc-team-1-$RANDOM --image nginx --namespace adhoc-team-1; done
for i in {1..3}; do kubectl run cap-$RANDOM --image nginx --namespace captainamerica; done
for i in {1..3}; do kubectl run adhoc-team-2-$RANDOM --image nginx --namespace adhoc-team-2; done
for i in {1..3}; do kubectl run hawkeye-$RANDOM --image nginx --namespace hawkeye; done
for i in {1..3}; do kubectl run antman-$RANDOM --image nginx --namespace antman; done
