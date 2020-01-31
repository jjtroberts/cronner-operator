# cronjob-notifier
Kubernetes operator to monitor CronJobs for failure and notify stakeholders.

# Build
`operator-sdk build gcr.io/gd-k8s-builder/gd-cronjob-notifier:v0.0.1`

# Generate CRDs
operator-sdk add api --api-version=notify.cronner.com/v1alpha1 --kind=Cronner
operator-sdk add controller --api-version=notify.cronner.com/v1alpha1 --kind=Cronner
operator-sdk generate k8s
operator-sdk generate crds

# Deploy
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml
kubectl create -f deploy/operator.yaml