package main

import "fmt"

type Item struct {
	ID   string
	Name string
}

type StockItem struct {
	Item
	Quantity int
}

func addItem(stockItems *[]StockItem, newItem StockItem) {
	*stockItems = append(*stockItems, newItem)
	fmt.Println("新しい在庫一覧")
	fmt.Println(*stockItems)
}

func main() {
	initialStock := []StockItem{
		{Item{"p1", "Pen"}, 10},
		{Item{"n1", "Notebook"}, 5},
		{Item{"e1", "Eraser"}, 20},
	}
	newItem := StockItem{Item{"p1", "Pen"}, 10}
	addItem(&initialStock, newItem)
}
