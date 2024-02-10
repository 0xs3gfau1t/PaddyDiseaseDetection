package client

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"segFault/PaddyDiseaseDetection/ent"
	"segFault/PaddyDiseaseDetection/ent/disease"
	"segFault/PaddyDiseaseDetection/ent/diseaseidentified"
	"segFault/PaddyDiseaseDetection/types"
	"strings"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/semaphore"
)

type MlConsumerClient interface {
	Run()
	UpdateStatus(context.Context, *types.ProcessedMessage) error
	FindDiseaseIdFromName(string, context.Context) (*uuid.UUID, error)
}

type MlConsumer struct {
	dbDisease           *ent.DiseaseClient
	dbDiseaseIdentified *ent.DiseaseIdentifiedClient
	channel             *amqp091.Channel
}

func (m MlConsumer) Run() {
	semapho := semaphore.NewWeighted(10)

	autoAck, exclusive, noLocal, noWait := false, false, false, false

	if resultChan, err := m.channel.Consume(os.Getenv("RABBIT_QUEUE_CONSUMER"), "", autoAck, exclusive, noLocal, noWait, nil); err == nil {
		for msg := range resultChan {
			ctx := context.Background()
			semapho.Acquire(ctx, 1)
			go func(cotx context.Context, release func(int64), msg amqp091.Delivery) {
				defer release(1)
				message := types.ProcessedMessage{}
				if err := json.Unmarshal(msg.Body, &message); err != nil {
					log.Println("Counldn't read response from queue", err, string(msg.Body))
					return
				}
				if err := m.UpdateStatus(cotx, &message); err == nil {
					msg.Ack(false)
				} else {
					//TODO: Solve infinite loop problem
					msg.Ack(false) // Temporary workaround
					//msg.Nack(false, true)
					log.Println("Error updating message: ", err)
				}
			}(ctx, semapho.Release, msg)
		}
	}
}

func (m MlConsumer) UpdateStatus(ctx context.Context, msg *types.ProcessedMessage) error {
	// TODO: Avaid setting status from processed to processing
	// THis could happen due to network level race conditions
	if id, err := m.FindDiseaseIdFromName(m.PrepareName(msg.Disease), ctx); err != nil {
		return err
	} else {
		return m.dbDiseaseIdentified.Update().Where(diseaseidentified.ID(msg.Id)).SetDisease(&ent.Disease{
			ID: *id,
		}).SetStatus(diseaseidentified.Status(msg.Status)).Exec(ctx)
	}
}

func (m MlConsumer) PrepareName(name string) string {
	finalString := ""
	for _, partName := range strings.Split(name, "_") {
		capitalized := strings.ToUpper(string(partName[0])) + partName[1:]
		finalString += capitalized + " "
	}
	return strings.Trim(finalString, " ")
}

func (m MlConsumer) FindDiseaseIdFromName(name string, ctx context.Context) (*uuid.UUID, error) {
	if d, err := m.dbDisease.Query().Where(disease.Name(name)).First(ctx); err == nil {
		return &d.ID, nil
	} else {
		return nil, err
	}
}
