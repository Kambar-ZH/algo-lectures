package main

import (
	"fmt"
)

type CategoryNode struct {
	Name            string          `json:"name"`
	Prices          []int           `json:"prices"`
	UniquePricesCnt int             `json:"uniquePricesCnt"`
	Children        []*CategoryNode `json:"children"`
}

func (m *CategoryNode) Clone() *CategoryNode {
	if m == nil {
		return nil
	}

	clone := *m

	if len(m.Children) > 0 {
		clone.Children = make([]*CategoryNode, len(m.Children))
		for i, category := range m.Children {
			if category == nil {
				continue
			}
			clone.Children[i] = category.Clone()
		}
	}

	return &clone
}

func mergeSets(largeSet, smallSet map[int]struct{}) {
	for val := range smallSet {
		largeSet[val] = struct{}{}
	}
}

func calculateUniquePrices(node *CategoryNode) map[int]struct{} {
	if node == nil {
		return nil
	}

	if len(node.Children) == 0 {
		uniquePrices := make(map[int]struct{}, len(node.Prices))
		for _, price := range node.Prices {
			uniquePrices[price] = struct{}{}
		}
		fmt.Printf("Категория %s: %d уникальных цен\n", node.Name, len(uniquePrices))
		node.UniquePricesCnt = len(uniquePrices)
		return uniquePrices
	}

	bigChild := map[int]struct{}{}
	bigChildIdx := -1
	childrenUniquePrices := make([]map[int]struct{}, 0, len(node.Children))
	for idx, child := range node.Children {
		childUniquePrices := calculateUniquePrices(child)
		if len(childUniquePrices) > len(bigChild) {
			bigChildIdx = idx
			bigChild = childUniquePrices
		}
		childrenUniquePrices = append(childrenUniquePrices, childUniquePrices)
	}

	uniquePrices := bigChild
	for idx, childUniquePrices := range childrenUniquePrices {
		if idx == bigChildIdx {
			continue
		}
		mergeSets(uniquePrices, childUniquePrices)
	}

	fmt.Printf("Категория %s: %d уникальных цен\n", node.Name, len(bigChild))
	node.UniquePricesCnt = len(uniquePrices)
	return uniquePrices
}

func main() {
	products := &CategoryNode{Name: "Товары", Prices: []int{}}
	electronics := &CategoryNode{Name: "Электроника", Prices: []int{}}
	mobiles := &CategoryNode{Name: "Мобильные", Prices: []int{100, 150, 250, 300}}
	tvs := &CategoryNode{Name: "Телевизоры", Prices: []int{250, 350}}
	laptops := &CategoryNode{Name: "Ноутбуки", Prices: []int{400, 500}}
	sport := &CategoryNode{Name: "Спорт", Prices: []int{}}
	cycling := &CategoryNode{Name: "Велоспорт", Prices: []int{50, 100}}
	tennis := &CategoryNode{Name: "Теннис", Prices: []int{250, 350, 400, 500}}

	electronics.Children = []*CategoryNode{mobiles, tvs, laptops}
	sport.Children = []*CategoryNode{cycling, tennis}
	products.Children = []*CategoryNode{electronics, sport}

	calculateUniquePrices(products)
}
