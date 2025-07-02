# push subgraph to apollo studio
```
APOLLO_KEY=service:My-Graph-1h158s:yAkc89p1eAD5RaCuJjve6g \
  rover subgraph publish My-Graph-1h158s@current \
  --schema ./graph/schema.graphqls \
  --name locations \
  --routing-url http://products.prod.svc.cluster.local:4001/graphql
```

