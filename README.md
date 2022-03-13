# YAF

YAML filter pipe

## Usage

```shell
# Exclude VMServiceScrape while operator is starting
kustomize build victoriametrics | go run github.com/albenik/yaf/cmd/kyaf -x operator.victoriametrics.com/v1beta1:VMServiceScrape | kubectl apply -f -
# Apply only VMServiceScrape
kustomize build victoriametrics | go run github.com/albenik/yaf/cmd/kyaf -i operator.victoriametrics.com/v1beta1:VMServiceScrape | kubectl apply -f -
```
