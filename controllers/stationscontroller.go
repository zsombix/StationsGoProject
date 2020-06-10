package controllers

import (
	"github.com/gofiber/fiber"
	stations "github.com/zsombix/StationsGoProject/models"
)

func GetAllStations(ctx *fiber.Ctx) {
	ctx.Status(200).JSON(stations.Stations)
}
