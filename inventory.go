package main

import (
	"fmt"
	"time"
)

func inventoryFull(p *Character) bool {
	return len(p.Inventory) >= p.Capacity
}
func addInventory(p *Character, item string) bool {
	if inventoryFull(p) {
		fmt.Println("Inventaire plein.")
		return false
	}
	p.Inventory = append(p.Inventory, item)
	return true
}
func removeInventory(p *Character, item string) bool {
	for i, x := range p.Inventory {
		if x == item {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			return true
		}
	}
	return false
}
func count(p *Character, item string) int {
	total := 0
	for _, x := range p.Inventory {
		if x == item {
			total++
		}
	}
	return total
}
func showInventory(p *Character) {
	fmt.Printf("\n=== INVENTAIRE %d/%d ===\n", len(p.Inventory), p.Capacity)
	if len(p.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
	}
	for i, item := range p.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}
func upgradeInventorySlot(p *Character) bool {
	if p.Upgrades >= 3 {
		fmt.Println("Maximum de 3 extensions atteint.")
		return false
	}
	p.Upgrades++
	p.Capacity += 10
	fmt.Println("Capacite :", p.Capacity)
	return true
}
func takePot(p *Character) bool {
	if !removeInventory(p, healPotion) {
		fmt.Println("Aucun polyconducteur.")
		return false
	}
	p.HP = clamp(p.HP+25, 0, p.MaxHP)
	fmt.Printf("PV : %d/%d\n", p.HP, p.MaxHP)
	return true
}
func poisonPot(p *Character) {
	fmt.Println("Corrupteur bionique injecte dans votre systeme.")
	for i := 1; i <= 3; i++ {
		time.Sleep(time.Second)
		p.HP = clamp(p.HP-10, 0, p.MaxHP)
		fmt.Printf("Cycle %d : vous perdez 10 PV. PV : %d/%d\n", i, p.HP, p.MaxHP)
	}
	isDead(p)
}
func equipItem(p *Character, item string) {
	var slot *string
	switch item {
	case coqueNeurale, neuroblindage:
		slot = &p.Equipment.Head
	case tuniqueCyber, blindageChrome:
		slot = &p.Equipment.Body
	case bottesRunner, cybersemelles:
		slot = &p.Equipment.Feet
	default:
		return
	}
	removeInventory(p, item)
	old := *slot
	*slot = item
	if old != "" {
		addInventory(p, old)
	}
	recalc(p)
	fmt.Println(item, "equipe.")
}
func useItem(p *Character, item string, m *Monster) bool {
	switch item {
	case healPotion:
		return takePot(p)
	case manaPotion:
		removeInventory(p, item)
		p.RAM += 1
		if p.RAM > p.MaxRAM {
			p.RAM = p.MaxRAM
		}
		fmt.Printf("RAM : %.1f/%.1f\n", p.RAM, p.MaxRAM)
		return true
	case poisonPotion:
		removeInventory(p, item)
		if m != nil {
			poisonEnemy(m)
		} else {
			poisonPot(p)
		}
		return true
	case etherHealPotion:
		removeInventory(p, item)
		p.HP = clamp(p.HP+50, 0, p.MaxHP)
		fmt.Printf("PV : %d/%d\n", p.HP, p.MaxHP)
		return true
	case etherRamPotion:
		removeInventory(p, item)
		p.RAM += 2
		if p.RAM > p.MaxRAM {
			p.RAM = p.MaxRAM
		}
		fmt.Printf("RAM : %.1f/%.1f\n", p.RAM, p.MaxRAM)
		return true
	case connecteur:
		removeInventory(p, item)
		p.MaxRAM += 1
		p.RAM += 1
		fmt.Printf("RAM : %.1f/%.1f\n", p.RAM, p.MaxRAM)
		return true
	case pyrocanon, blackout, polymyeline:
		removeInventory(p, item)
		learnImplant(p, item)
		return true
	case upgrade:
		if !upgradeInventorySlot(p) {
			return false
		}
		removeInventory(p, item)
		return true
	case coqueNeurale, neuroblindage, tuniqueCyber, blindageChrome, bottesRunner, cybersemelles:
		equipItem(p, item)
		return true
	}
	fmt.Println("Objet non utilisable.")
	return false
}
func accessInventory(p *Character) {
	for {
		ClearTerminal()
		showInventory(p)
		fmt.Println("0. Retour")
		fmt.Print("Objet a utiliser : ")
		choice := readInt()
		if choice == 0 {
			return
		}
		if choice >= 1 && choice <= len(p.Inventory) {
			useItem(p, p.Inventory[choice-1], nil)
		}
	}
}
