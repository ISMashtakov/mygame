package components

import "github.com/yohamta/donburi"

type ObstacleData struct{}

var Obstacle = donburi.NewComponentType[ObstacleData]()

type GardenData struct{}

var Garden = donburi.NewComponentType[GardenData]()
