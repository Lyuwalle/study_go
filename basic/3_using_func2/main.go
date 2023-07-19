package main

import "fmt"

// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
	kingdom string // 界。
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

type Animal struct {
	scientificName string // 学名。
	AnimalCategory        // 动物基本分类。字段声明AnimalCategory代表了Animal类型的一个嵌入字段
}

// 声明一个方法，方法隶属于一个类型。从它的接收者声明可以看出它隶属于AnimalCategory类型。
func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s", ac.kingdom, ac.phylum, ac.class, ac.order, ac.family, ac.genus, ac.species)
}

func (a Animal) Category() string {
	return a.AnimalCategory.String()
}

func (a Animal) String() string {
	return fmt.Sprintf("%s (category: %s)", a.scientificName, a.AnimalCategory)
}

func main() {
	category := AnimalCategory{species: "cat"}
	fmt.Printf("The animal category: %s\n", category)

	animal := Animal{
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	fmt.Printf("The animal: %s\n", animal)

}
