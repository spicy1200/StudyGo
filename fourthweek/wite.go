// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import "github.com/google/wire"

func Initialize(CityName string) City {
	wire.Build(NewCity,NewSchool)
	return City{}
}