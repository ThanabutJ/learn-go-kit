echo "curl uppercase"
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase

echo "curl Count"
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
