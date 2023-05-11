package common

type Module struct {
	Imports    Imports
	Controller IController
}

type Service struct {
	Imports Imports
}

type Controller struct {
	Service Service
	Imports Imports
}
