package handlers

import (
	"log"

	"github.com/danik-tro/usd-rate/pkg/core"
	"github.com/danik-tro/usd-rate/pkg/models"
	storage "github.com/danik-tro/usd-rate/pkg/storages"
	"github.com/danik-tro/usd-rate/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

const BATCH_SIZE = 500

type rateHandler struct {
	s  *storage.Storage
	rc *storage.Cache
	c  *core.Config
}

func NewRateHandler(c *core.Config, s *storage.Storage, rc *storage.Cache) rateHandler {
	return rateHandler{s, rc, c}
}

func (rt *rateHandler) Rate(c *fiber.Ctx) error {
	rate, err := rt.rc.GetRate()

	if err != nil {
		rate, err = utils.UsdRate()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal error",
			})
		}
		rt.rc.SetRate(rate)
	}

	return c.Status(fiber.StatusOK).JSON(rate)
}

func (rt *rateHandler) Subscribe(c *fiber.Ctx) error {
	s := new(models.Subscriber)

	if err := c.BodyParser(s); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse email",
		})
	}

	is_exists, err := rt.s.IsEmailSubscribed(s.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal error",
		})
	}

	if is_exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "subscriber already exists",
		})
	}

	if err := rt.s.SubscribeEmail(s); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func (rt *rateHandler) SendEmails(c *fiber.Ctx) error {
	go func(batchSize int, rt *rateHandler) {
		rate, err := utils.UsdRate()

		if err != nil {
			log.Printf("Failed to get current rate %s", err)
			return
		}

		totalSubscribers := rt.s.TotalSubscribers()

		numBatches := int(totalSubscribers) / batchSize
		if int(totalSubscribers)%batchSize != 0 {
			numBatches++
		}

		for batch := 0; batch < numBatches; batch++ {
			offset := batch * batchSize

			subscribers := rt.s.FetchBatchSubscribers(int64(batchSize), int64(offset))

			for _, subscriber := range subscribers {
				err := utils.SendMessage(rt.c, subscriber.Email, rate)

				if err != nil {
					log.Printf("Error sending an email %s: %v", subscriber.Email, err)
				}
			}
		}
	}(BATCH_SIZE, rt)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}
