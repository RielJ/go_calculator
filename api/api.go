package api

import (
	"fmt"
)

const (
	_ = iota
	RED
	GREEN
	BLUE
	YELLOW
	PINK
	PURPLE
	ORANGE
)

type ItemSet int

type Item struct {
	Name   string
	Price  float64
	Set    ItemSet
	Orders int
}

type Store map[ItemSet]Item

func (il *Store) Init() {
	*il = Store{
		RED:    Item{"Red", 50, RED, 0},
		GREEN:  Item{"Green", 40, GREEN, 0},
		BLUE:   Item{"Blue", 30, BLUE, 0},
		YELLOW: Item{"Yellow", 50, YELLOW, 0},
		PINK:   Item{"Pink", 80, PINK, 0},
		PURPLE: Item{"Purple", 90, PURPLE, 0},
		ORANGE: Item{"Orange", 120, ORANGE, 0},
	}
}

func (il *Store) AddOrder(name ItemSet) {
	if item, ok := (*il)[name]; ok {
		(*il)[name] = Item{item.Name, item.Price, item.Set, item.Orders + 1}
	} else {
		fmt.Println("Item not found")
	}
}

func (il *Store) CalculateTotal() float64 {
	var total float64
	for _, item := range *il {
		switch item.Set {
		case ORANGE, PINK, GREEN:
			total += calculateBundle(item.Orders, item.Price)
		default:
			total += item.Price * float64(item.Orders)
		}
	}
	return total
}

func (il *Store) CalculateTotalCustomerMember() float64 {
	total := il.CalculateTotal()
	return total * 0.9
}

func (il *Store) PrintItems() {
	fmt.Println()
	for _, item := range *il {
		fmt.Printf("[%s]%s: $%.2f\n", item.Name[0:2], item.Name[2:], item.Price)
	}
	fmt.Println()
}

func (il *Store) PrintOrders() {
	for _, item := range *il {
		if item.Orders > 0 {
			fmt.Printf("%s: %d\n", item.Name, item.Orders)
		}
	}
}

func (il *Store) GetItemByName(name string) (Item, bool) {
	switch name {
	case "Re":
		return (*il)[RED], true
	case "Gr":
		return (*il)[GREEN], true
	case "Bl":
		return (*il)[BLUE], true
	case "Ye":
		return (*il)[YELLOW], true
	case "Pi":
		return (*il)[PINK], true
	case "Pu":
		return (*il)[PURPLE], true
	case "Or":
		return (*il)[ORANGE], true
	default:
		return Item{}, false
	}
}

func calculateBundle(orders int, price float64) float64 {
	discounted := float64(orders / 2)
	regular := float64(orders % 2)
	return (discounted*2)*(price*0.95) + regular*price
}
