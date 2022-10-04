package basic

import (
	"fmt"
	"strings"
)

// 변수는 var/const 이름 자료형 순으로
// :=를 통하면 타입을 자동으로 지정
// func main() {
// 	// total := superAdd(1, 2, 3, 4, 5, 6)
// 	// fmt.Println(total)
// 	// fmt.Println(testSwitch(70))
// 	// pointerTest()
// 	// testArray()
// 	// testMap()
// 	structTest()
// }

// defer를 통해 함수가 리턴한 이후에 실행을 할 수 있다.
// 리턴값이 두개 이상일 수 있음
// naked function이 있음 (하지만 return을 명확히 하는게 좋을 것 같긴함)
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// 반복문은 for만 존재
// ...를 붙이면 무한적으로 인자를 받을 수 있음 (배열형식)
// 안쓰는 변수는 _를 씀
func superAdd(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println("이렇게도 사용가능합니다", numbers[i])
	}
	// range는 index와 데이터를 반환
	// ignore value
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func canIDrink(age int) bool {
	// if문에 변수 생성 가능
	// if age < 18
	// if 문에서 사용하기 위해서만 사용된다는 것을 의미
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}

func testSwitch(record int) string {
	// 해당 코드도 가능
	// switch record {
	// case 90:
	// 	return "A"
	// case 80:
	// 	return "B"
	// default:
	// 	return "C"
	// }
	switch plusPoint := record + 5; plusPoint {
	case 90:
		return "A"
	default:
		return "F"
	}
}

func pointerTest() {
	a := 2
	b := &a

	// 이렇게도 사용 가능
	// var c int = 5
	// var d *int = &c
	// fmt.Println(*d)

	fmt.Println(&a, b)
	fmt.Println(a, *b)
}

func testArray() {
	// 배열의 크기를 제한 하고 싶을 때
	names := [5]string{"nico", "think", "do"}
	names[3] = "hey"
	fmt.Println(names)

	// Slice라는 자료형, limit 없음
	// apppend 사용, 하지만 append사용법은 아래와 같이 해야함
	names2 := []string{}
	names2 = append(names2, "hi") // 업데이트
	fmt.Println(names2)
}

func testMap() {
	// 만약 value 값이 다른 값도 들어오게 짜려면 struct를 써야함
	dongwon := map[string]int{"hi": 1, "two": 2}
	for key, value := range dongwon {
		fmt.Println(key, value)
	}
}

func structTest() {
	// struct에는 constructor가 없음
	// 메서드는 넣을 수 있음
	type person struct {
		name    string
		age     int
		favFood []string
	}

	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	// nico := person{}
	// nico.age = 18
	// nico.favFood = favFood
	// nico.name = "nico"
	fmt.Println(nico)
}
