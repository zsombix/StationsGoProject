package controllers

import (
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	stations "github.com/zsombix/StationsGoProject/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllStations(ctx *fiber.Ctx) {

	collection := mgm.Coll(&stations.Station{})

	stations := []stations.Station{}

	err := collection.SimpleFind(&stations, bson.D{})
	if err != nil {
		ctx.Status(500)
		return
	}

	ctx.Status(200).JSON(stations)
}
