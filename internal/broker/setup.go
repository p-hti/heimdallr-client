package broker

import (
	"context"
	"encoding/json"
	"log"

	"github.com/p-hti/heimdallr-client/internal/manage"
	"github.com/segmentio/kafka-go"
)

type Broker struct {
	KfWriter *kafka.Writer
	Machine  *manage.Machine
}

// type BrokerReader struct {
// 	kfReader *kafka.Writer
// }

type MachineUseCases interface {
	// GetFeatures() error
	GetResourceUsage() error
}

func NewBrokerWriter(addresses string, topic string, machine *manage.Machine) *Broker {
	broker := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{addresses},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return &Broker{
		KfWriter: broker,
		Machine:  machine,
	}
}

func (b *Broker) SendResourceUsage(ctx context.Context, errChan chan<- error) {
	for {
		select {
		case <-ctx.Done():
			log.Println("context is done")
			return
		default:
			err := b.Machine.GetResourceUsage()
			if err != nil {
				errChan <- err
				continue
			}
			data, err := json.Marshal(b.Machine)
			if err != nil {
				errChan <- err
				continue
			}

			err = b.KfWriter.WriteMessages(
				context.Background(),
				kafka.Message{
					Value: data,
				})
			if err != nil {
				errChan <- err
				continue
			}
		}
	}
}
