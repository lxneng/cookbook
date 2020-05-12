package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"strconv"
	"strings"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type Item struct {
	ItemId      int64    `json:"item_id"`
	ItemType    string   `json:"item_type"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Tags        []string `json:"tags"`
	IsPublish   bool     `json:"is_publish"`
	Weight      int64    `json:"weight"`
	PublishedAt string   `json:"published_at"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
	})
}

func main() {
	kafkaURL := getEnv("kafkaURL", "localhost:9092")
	topic := getEnv("topic", "items")
	groupID := getEnv("groupID", "item-profile-worker")
	esHost := getEnv("esHost", "http://localhost:9200")

	reader := getKafkaReader(kafkaURL, topic, groupID)
	ctx := context.Background()

	defer reader.Close()

	es, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esHost))
	if err != nil {
		panic(err)
	}
	info, code, err := es.Ping(esHost).Do(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	log.Println("start consuming ...")

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("message at topic:%v partition:%v offset:%v    %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		var item Item
		err = json.Unmarshal(m.Value, &item)
		if err != nil {
			log.Fatal(err)
		}

		var rawJSON map[string]json.RawMessage
		err = json.Unmarshal(m.Value, &rawJSON)
		if err != nil {
			log.Fatal(err)
		}

		_, err = UpsertItemToEs(item, es, ctx)
		if err != nil {
			log.Fatal(err)
		}

		err = reader.CommitMessages(ctx, m)
		if err != nil {
			log.Printf("fail to commit msg: ", err)
		}
	}
}

func UpsertItemToEs(item Item, es *elastic.Client, ctx context.Context) (*elastic.UpdateResponse, error) {
	docId := strconv.FormatInt(int64(item.ItemId), 10)
	res, err := es.Update().Index("items").Id(docId).Doc(map[string]interface{}{"metainfo": item}).DocAsUpsert(true).Do(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
