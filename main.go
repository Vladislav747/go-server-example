package main

import "fmt"

const pi = 3.14

func main(){
	var pharse = "Go рулит!"
	fmt.Println(pharse);
	printCircleArea(2);
	printCircleArea(4);


}

func printCircleArea(radius int) {
	fmt.Printf("Радиус круга: %d cm.\n", radius);
	fmt.Printf("Формула для расчета площади круга A=nr2\n");

	circleRadius := calculateCircleArea(radius);

	fmt.Printf("Площадь круга: %f32 см. кв.\n", circleRadius);
}

func calculateCircleArea(radius int) float32 {
	return float32(radius) * float32(radius) * pi
}