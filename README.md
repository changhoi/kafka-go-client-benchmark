# Go Kafka Client Benchmark test

Go 클라이언트 중 다음 두 개의 밴치마크 테스트를 수행

- [kafka-go](https://github.com/segmentio/kafka-go)
- [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go)

## 테스트
- 테스트는 Async로 메시지를 Produce 하는 상황에 대해서만 진행
- confluent-kafka-go가 C 패키지를 같이 사용하고, ARM 환경에서는 해당 패키지가 동작하지 않아서 도커를 사용해 테스트 환경 구성

```shell
make set-kafka # kafka docker compose 실행

make build # 테스트 컨테이너 빌드
make test # 테스트 실행

make clean-kafka # kafka 내리기 
```
---

```plaintext
goos: linux
goarch: amd64
pkg: kafka-client-bench/bench
cpu: QEMU TCG CPU version 2.5+
BenchmarkConfluentKafkaGoAsync
BenchmarkConfluentKafkaGoAsync                     53566             18713 ns/op             552 B/op         14 allocs/op
BenchmarkConfluentKafkaGoAsync-3                  125451              8987 ns/op             552 B/op         14 allocs/op
BenchmarkConfluentKafkaGoAsync-6                  142630              8909 ns/op             552 B/op         14 allocs/op
BenchmarkConfluentKafkaGoAsyncInParalell
BenchmarkConfluentKafkaGoAsyncInParalell           63799             18373 ns/op             521 B/op         13 allocs/op
BenchmarkConfluentKafkaGoAsyncInParalell-3        135698              9914 ns/op             521 B/op         13 allocs/op
BenchmarkConfluentKafkaGoAsyncInParalell-6        120361              8727 ns/op             521 B/op         13 allocs/op
BenchmarkKafkaGo
BenchmarkKafkaGo                                  120118             19040 ns/op             418 B/op          8 allocs/op
BenchmarkKafkaGo-3                                151347              7419 ns/op             418 B/op          8 allocs/op
BenchmarkKafkaGo-6                                159974              7550 ns/op             418 B/op          8 allocs/op
BenchmarkKafkaGoAsync
BenchmarkKafkaGoAsync                             107686             14074 ns/op             410 B/op          7 allocs/op
BenchmarkKafkaGoAsync-3                           239545              5726 ns/op             410 B/op          7 allocs/op
BenchmarkKafkaGoAsync-6                           215001              5349 ns/op             410 B/op          7 allocs/op
PASS
ok      kafka-client-bench/bench        43.256s
```

[source code](bench)
