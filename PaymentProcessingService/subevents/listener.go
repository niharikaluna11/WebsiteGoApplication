package subevents

import (
	"PaymentProcessingService/models"
	"PaymentProcessingService/repository"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

const Subscriber = "payment-service-sub"

func SubscribeToOrderEvents(repo *repository.PaymentRepoImpl) {

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "niharikaprojects")
	if err != nil {
		log.Fatalf("PubSub Client error: %v", err)
	}

	sub := client.Subscription(Subscriber)

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Println("✅ Received message from PubSub 'order-updates'...")
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
