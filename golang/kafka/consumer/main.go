package main

import (
  "context"
  "log"
  "os"
  "strings"
  "time"
  kafka "github.com/segmentio/kafka-go"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
  brokers := strings.Split(kafkaURL, ",")
  return kafka.NewReader(kafka.ReaderConfig{
    Brokers:  brokers,
    GroupID:  groupID,
    Topic:    topic,
    MaxBytes: 10e6, // 10MB
    CommitInterval: 10 * time.Second,
  })
}

func main() {
  kafkaURL := os.Getenv("kafkaURL")
  topic := os.Getenv("topic")
  groupID := os.Getenv("groupID")

  reader := getKafkaReader(kafkaURL, topic, groupID)
  ctx := context.Background()

  defer reader.Close()

  log.Println("start consuming ... !!")

  for {
    m, err := reader.ReadMessage(ctx)
    if err != nil {
      log.Fatalln(err)
    }
    log.Printf("message at topic:%v partition:%v offset:%v    %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
    err = reader.CommitMessages(ctx, m)
    if err != nil {
      log.Printf("fail to commit msg: ", err)
    }
  }

}
