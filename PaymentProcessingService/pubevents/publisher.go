package pubevents

import (
	"PaymentProcessingService/models"
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
)

const OrderTopic = "payment-events"

func PublishOrderCreated(ctx context.Context, client *pubsub.Client, order models.PaymentEvent) error {

	ctx = context.Background()
	client, err := pubsub.NewClient(ctx, "niharikaprojects")
	if err != nil {
		return err
	}

	defer client.Close()

	topic := client.Topic(OrderTopic)
	data, _ := json.Marshal(order)

	result := topic.Publish(ctx, &pubsub.Message{Data: data})

	print("Publishing 'payment-events' to PubSub...\n")
	print("Payment is success...\n")

	_, err = result.Get(ctx)
	return err
}
