# YAF

YAML filter pipe

## Usage

```shell
# Exclude VMServiceScrape while operator is starting
kustomize build victoriametrics | kubectl apply -f - | go run github.com/albenik/yaf/cmd/kyaf -x operator.victoriametrics.com/v1beta1:VMServiceScrape
# Apply only VMServiceScrape
kustomize build victoriametrics | kubectl apply -f - | go run github.com/albenik/yaf/cmd/kyaf -i operator.victoriametrics.com/v1beta1:VMServiceScrape
```
