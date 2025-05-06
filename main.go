package main
import "fmt"
func main() {
	Ex1()
	Ex3()
	Ex4()
	Ex5()
}
func Ex1() {
	students := []Student{
		{Name: "John", Age: 20, Group: "A", GPA: 3.6},
		{Name: "Alice", Age: 22, Group: "B", GPA: 3.8},
		{Name: "Bob", Age: 21, Group: "A", GPA: 3.2},
	}
	for _, student := range students {
		fmt.Printf("Name: %s, Age: %d, Group: %s, GPA: %.2f\n", student.Name, student.Age, student.Group, student.GPA)
	}
	printExcellentStudents(students)
}
type Student struct {
	Name  string
	Age   int
	Group string
	GPA   float64
}
func printExcellentStudents(students []Student) {
	fmt.Println("\nОтличные студенты (GPA > 3.5):")
	for _, student := range students {
		if student.GPA > 3.5 {
			fmt.Printf("Name: %s, Age: %d, Group: %s, GPA: %.2f\n",
				student.Name, student.Age, student.Group, student.GPA)
		}
	}
}
type City struct {
	Name       string
	Population int
	IsCapital  bool
}

func Ex3() {
	citys := []City{
		{Name: "Moscow", Population: 11920000, IsCapital: true},
		{Name: "Saint Petersburg", Population: 5384000, IsCapital: false},
		{Name: "Novosibirsk", Population: 1612833, IsCapital: false},
		{Name: "Yekaterinburg", Population: 1493748, IsCapital: false},
	}
	var capitalCount int
	var totalPopulation int

	for _, city := range citys {
		if city.IsCapital {
			capitalCount++
		}
		totalPopulation += city.Population
		fmt.Printf("City: %s, Population: %d, Is Capital: %t\n", city.Name, city.Population, city.IsCapital)
	}

	fmt.Printf("\nNumber of capitals: %d\n", capitalCount)
	fmt.Printf("Total population: %d\n", totalPopulation)
}
func Ex4(){
games := []Game{
		{Title: "Game1", Platform: "PC", HoursPlayed: 10},
		{Title: "Game2", Platform: "Xbox", HoursPlayed: 5},
		{Title: "Game3", Platform: "PlayStation", HoursPlayed: 20},
		{Title: "Game4", Platform: "PC", HoursPlayed: 15},
		{Title: "Game5", Platform: "Xbox", HoursPlayed: 8},
		{Title: "Game6", Platform: "PlayStation", HoursPlayed: 12},
		{Title: "Game7", Platform: "PC", HoursPlayed: 25},
		{Title: "Game8", Platform: "Xbox", HoursPlayed: 30},

	}
	total:=totalPlaytime(games)
	fmt.Println("Total playtime:", total)
}
type Game struct {
	Title string 
	Platform string
	HoursPlayed int
}
func totalPlaytime(games []Game) int {
	sum := 0
	for _, game := range games {
		sum += game.HoursPlayed
	}
	return sum
}
func Ex5(){
	
	orders:= []Order{
		{Item: "Apple", Quantity: 10, PricePerUnit: 0.5},
		{Item: "Banana", Quantity: 5, PricePerUnit: 0.2},
		{Item: "Orange", Quantity: 8, PricePerUnit: 0.3},
		{Item: "Grapes", Quantity: 12, PricePerUnit: 0.4},
	}

	total := 0.0
	for _, order := range orders {
		orderTotal := float64(order.Quantity) * order.PricePerUnit
		fmt.Printf("Товар: %s, сумма: %.2f\n", order.Item, orderTotal)
		total += orderTotal
	}

	fmt.Printf("Итоговая сумма всех заказов: %.2f\n", total)
}
type Order struct {
	Item string
	Quantity int
	PricePerUnit float64
}
