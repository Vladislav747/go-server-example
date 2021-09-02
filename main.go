package main

import ("fmt"; "errors")

const pi = 3.14

func main(){
	var pharse = "Go рулит!"
	fmt.Println(pharse);
	printCircleArea(0);
	printCircleArea(2);
	printCircleArea(4);


}

func printCircleArea(radius int) {

	circleArea, err := calculateCircleArea(radius);
	x := 10;

	if err != nil {
		fmt.Println(err.Error());
		return
	}

	
	fmt.Printf("Радиус круга: %d cm.\n", radius);
	fmt.Printf("Формула для расчета площади круга A=nr2\n");

	fmt.Printf("Площадь круга: %f32 см. кв.\n", circleArea);
	fmt.Println("Адрес значения x");
	fmt.Println(&x);
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		fmt.Println("Радиус круга не может быть меньше 0");
		return float32(0), errors.New("Радиус круга не может быть отрицательным");
	}
	return float32(radius) * float32(radius) * pi, nil
}
