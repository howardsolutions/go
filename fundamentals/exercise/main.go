package main

import "fmt"

type Player struct {
	Name      string
	Inventory []Item
}

type Item struct {
	Name string
	Type string
}

func main() {

}

func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s \n", p.Name, item.Name)
}

func (p *Player) DropItem(itemName string) {
	for i, invItem := range p.Inventory {
		if invItem.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s dropped %s. \n", p.Name, itemName)
			return
		}
	}
	fmt.Printf("%s does not have %s in inventory \n", p.Name, itemName)
}

func (p *Player) UseItem(itemName string) {
	for i, item := range p.Inventory {
		if itemName == item.Name {
			if item.Type == "potion" {
				fmt.Printf("%s used %s and feels rejuvenated! \n", p.Name, itemName)
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			} else {
				fmt.Printf("%s used %s \n", p.Name, itemName)
			}
		}
		return
	}

	fmt.Printf("%s does not have %s in inventory", p.Name, itemName)
}
