package pubsubevents

import (
	"PaymentProcessingService/models"
	"PaymentProcessingService/repository"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func SubscribeToOrderEvents(repo *repository.PaymentRepoImpl) {
	fmt.Println(" Subscribing to 'payment-service-sub' pubsub topic...")

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "niharikaprojects")
	if err != nil {
		log.Fatalf("PubSub Client error: %v", err)
	}

	sub := client.Subscription("payment-service-sub")

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("✅ Payment Service received: %s\n", string(msg.Data))

		var event models.OrderCreatedEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("❌ Failed to parse message: %v", err)
			msg.Nack()
			return
		}

		newPayment := &models.PaymentCreate{
			OrderID: event.OrderID,
			Amount:  event.TotalAmount,
			Method:  models.DebitCard,
		}

		repo.MakePayment(client, ctx, newPayment)
		msg.Ack()
	})

	if err != nil {
		log.Fatalf("❌ Failed to receive messages: %v", err)
	}
}
