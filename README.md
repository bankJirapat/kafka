# Link Tutorial kafka confluent

https://developer.confluent.io/get-started/go/

## Command create Topic

```
docker compose exec broker kafka-topics --create --topic purchases --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1
```
