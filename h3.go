package main

import (
	"github.com/uber/h3-go"
	"fmt"
	"strings"
)

// ,,
func main() {

	//h3baseCell()
	//h3toGeo()
	//h3kring()

	//Polyfill()
	HexRange()
}

func h3baseCell() {
	result := h3.BaseCell(0x8931aa5745bffff)
	fmt.Println(result)
}
func h3toGeo() {
	result := h3.ToGeo(0x8931aa5745bffff)
	fmt.Println(result)
}

func h3kring() {
	geo := h3.GeoCoord{
		Latitude:  40.021758,
		Longitude: 116.374472,
	}
	resolution := 9
	h3index := h3.FromGeo(geo, resolution)

	fmt.Println(toJs(h3.ToGeoBoundary(h3index)))
	result := h3.KRing(h3index, 9)
	//fmt.Println(result)
	v := toJsFromH3Index(result)
	fmt.Println(v)
}

func h3KRingDistance(){
	geo := h3.GeoCoord{
		Latitude:  40.021758,
		Longitude: 116.374472,
	}
	resolution := 9
	h3index := h3.FromGeo(geo, resolution)

	fmt.Println(toJs(h3.ToGeoBoundary(h3index)))
	result := h3.KRingDistances(h3index, 1000)
	fmt.Println(result)
	//v := toJsFromH3Index(result)
	//fmt.Println(v)
}

func HexRange(){
	geo := h3.GeoCoord{
		Latitude:  40.021758,
		Longitude: 116.374472,
	}
	resolution := 9
	h3index := h3.FromGeo(geo, resolution)
	result,err:=h3.HexRange(h3index,10)
	if err!= nil {

	}
	fmt.Println(toJsFromH3Index(result))
}

func Polyfill() {
	validGeofence := h3.GeoBoundary{
		{Latitude: 39.920255, Longitude: 116.403322},
		{Latitude: 39.897555, Longitude: 116.410703},
		{Latitude: 39.892353, Longitude: 116.402292},
		{Latitude: 39.891365, Longitude: 116.389846},
	}
	obj := h3.GeoPolygon{
		Geofence: validGeofence,
	}
	result := h3.Polyfill(obj, 10)
	fmt.Printf("生成了%v个格子\n", len(result))
	finalstr := "["
	for _, v := range result {
		finalstr = finalstr + toJs(h3.ToGeoBoundary(v)) + ","
	}
	finalstr = strings.TrimRight(finalstr, ",")
	finalstr += "]"
	fmt.Println(finalstr)

}

func toJsFromH3Index(value []h3.H3Index) string {
	finalstr := "["
	for _, v := range value {
		finalstr = finalstr + toJs(h3.ToGeoBoundary(v)) + ","
	}
	finalstr = strings.TrimRight(finalstr, ",")
	finalstr += "]"
	return finalstr
}

// 展示多个对变形
func toJs(geo h3.GeoBoundary) string {

	str := "["
	for _, value := range geo {
		tmp := fmt.Sprintf("[%v,%v],", value.Longitude, value.Latitude)
		str = str + tmp
	}
	str = strings.TrimRight(str, ",")
	str += "]"
	return str
}
