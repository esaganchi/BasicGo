package main // исполняемая программа; в пакете main должна быть func main()

import "fmt" // пакет для вывода текста в консоль

// В статье сначала бывает так (именованное поле Person):
//
//	type Android struct {
//		Person Person
//		Model  string
//	}
//
// Тогда метод зовут только так: a.Person.Talk()
// Ниже — вариант из той же статьи со встроенным (анонимным) полем Person.

type Person struct { // тип «личность» из примера
	Name string // имя
}

func (p *Person) Talk() { // метод у *Person
	fmt.Println("Hi, my name is", p.Name) // как в учебнике
}

type Android struct { // андроид: Person встроен без имени поля (embedded)
	Person        // анонимное поле типа Person — методы и поля «продвигаются»
	Model  string // модель
}

func main() { // точка входа
	// Литерал как в статье: Model и вложенный Person с ключом Person (имя типа).
	a := Android{
		Model: "model",
		Person: Person{
			Name: "name",
		},
	}

	fmt.Print("a.Person.Talk(): ")
	a.Person.Talk() // явный путь через имя типа встроенного поля

	fmt.Print("a.Talk():        ")
	a.Talk() // продвижение метода Person на Android

	fmt.Println("a.Name:", a.Name) // поле Name тоже доступно как a.Name
}
