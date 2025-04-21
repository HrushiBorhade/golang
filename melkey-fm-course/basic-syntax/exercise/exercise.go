// Starting code for the exercise
package exercise

// Item struct
type Item struct {
	Name string
	Type string
}

// Player struct
type Player struct {
	Name      string
	Inventory []Item
}

// Pick up an item (modifies the player's inventory)
func (p *Player) PickUpItem(item Item) {
	// TODO: Implement function to add an item to inventory
	p.Inventory = append(p.Inventory, item)

}

// Drop an item (removes from inventory)
func (p *Player) DropItem(itemName string) {
	// TODO: Implement function to remove an item from inventory
	for index, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:index], p.Inventory[index+1:]...)
			return
		}
	}

}

// Use an item (if potion, remove it after use)
func (p *Player) UseItem(itemName string) {
	// TODO: Implement function to use an item
}
