package main

import "fmt"

type City struct {
	Name string
	longitude int
	latitude int
	School School
}
type School struct {
	Name string
	//Cla Cla
}
func NewSchool() School {
	return School{Name: "你好"}
}
func NewCity(sch School,CityName string) City{
	return City{School: sch,Name: CityName}
}
func (c City) splicing() string{
	return fmt.Sprintf("所在城市 %s", c.Name)
}
func main()  {
	result := Initialize("杭州").splicing()
	fmt.Println(result)
}