package pubsubevents

import (
	"OrderProcessingService/models"
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
)

const OrderTopic = "order-updates"

func PublishOrderCreated(ctx context.Context, client *pubsub.Client, order models.OrderEvent) error {

	ctx = context.Background()
	client, err := pubsub.NewClient(ctx, "niharikaprojects")
	if err != nil {
		return err
	}

	defer client.Close()

	topic := client.Topic("order-updates")
	data, _ := json.Marshal(order)

	result := topic.Publish(ctx, &pubsub.Message{Data: data})

	print("Publishing order event to PubSub...\n")

	_, err = result.Get(ctx)
	return err
}
