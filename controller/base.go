package controller

import (
	"strconv"

	"github.com/atharv-bhadange/log-ingestor/model"
	"github.com/atharv-bhadange/log-ingestor/service"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(model.Response{
		Status:  200,
		Message: "Server is up and running",
		Data:    nil,
	})
}

func AddLog(c *fiber.Ctx) error {
	var log model.Log

	if err := c.BodyParser(&log); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:  400,
			Message: "Bad request",
			Data:    nil,
		})
	}

	log, err := service.AddLog(log)

	if err != nil {
		return c.Status(500).JSON(model.Response{
			Status:  500,
			Message: "Internal server error",
			Data:    nil,
		})
	}

	return c.Status(200).JSON(model.Response{
		Status:  200,
		Message: "Log added successfully",
		Data:    log.Message,
	})
}

func GetLog(c *fiber.Ctx) error {

	logs, err := service.GetLog()

	if err != nil {
		return c.Status(500).JSON(model.Response{
			Status:  500,
			Message: "Internal server error",
			Data:    nil,
		})
	}

	return c.Status(200).JSON(model.Response{
		Status:  200,
		Message: "Log fetched successfully",
		Data:    logs,
	})
}

func GetLogById(c *fiber.Ctx) error {

	id_str := c.Params("id")

	id, err := strconv.Atoi(id_str)

	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:  400,
			Message: "Bad request",
			Data:    nil,
		})
	}

	log, err := service.GetLogById(id)

	if err != nil {
		return c.Status(500).JSON(model.Response{
			Status:  500,
			Message: "Internal server error",
			Data:    nil,
		})
	}

	return c.Status(200).JSON(model.Response{
		Status:  200,
		Message: "Log fetched successfully",
		Data:    log,
	})
}

func GetLogsByLevel(c *fiber.Ctx) error {
	level := c.Params("level")

	logs, err := service.GetLogByLevel(level)

	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:  400,
			Message: "Bad request.",
			Data:    nil,
		})
	}

	return c.Status(200).JSON(model.Response{
		Status:  200,
		Message: "Logs fetched successfully",
		Data:    logs,
	})
}

