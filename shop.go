package main

import "fmt"

func market(p *Character) {
	items := []Product{
		{healPotion, 15},
		{manaPotion, 20},
		{poisonPotion, 25},
		{pyrocanon, 50},
		{blackout, 120},
		{polymyeline, 80},
		{antenna, 5},
		{plaqueCyber, 10},
		{gaine, 12},
		{chromofibre, 20},
		{upgrade, 100},
		{etherHealPotion, 30},
		{etherRamPotion, 40},
		{connecteur, 150},
	}
	for {
		ClearTerminal()
		fmt.Printf(cyan+"\n=== SCAVENGER : %d credits ===\n", p.Gold)
		for i, item := range items {
			fmt.Printf("%d. %s : %d credits\n", i+1, item.Name, item.Price)
		}
		fmt.Println("0. Retour")
		fmt.Print("Choix : ")
		choice := readInt()
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(items) {
			continue
		}
		item := items[choice-1]
		if p.Gold < item.Price {
			fmt.Println("Credits insuffisants.")
			continue
		}
		if addInventory(p, item.Name) {
			p.Gold -= item.Price
			fmt.Println("Achat confirme :", item.Name, "-", item.Price, "credits.")
		}
	}
}
func craft(p *Character, result string, materials map[string]int) {
	if p.Gold < 10 {
		fmt.Println("Il faut 10 credits.")
		return
	}
	for item, amount := range materials {
		if have := count(p, item); have < amount {
			fmt.Printf("Composants insuffisants : %s (%d/%d).\n", item, have, amount)
			return
		}
	}
	for item, amount := range materials {
		for i := 0; i < amount; i++ {
			removeInventory(p, item)
		}
	}
	p.Gold -= 5
	addInventory(p, result)
	fmt.Println(result, "fabrique.")
}
func workshop(p *Character) {
	for {
		ClearTerminal()
		fmt.Println(magenta + "\n=== ASSEMBLEUR CYBERNETIQUE ===")
		fmt.Println("1. Coque neurale (+10 pv) : antenne + gaine isolante + 10 credits")
		fmt.Println("2. Neuroblindage (+20 pv) : 2 antennes + chromofibre + 10 credits")
		fmt.Println("3. Tunique cybernetique (+25 pv) : 2 antennes + plaque cyber + 10 credits")
		fmt.Println("4. Blindage chrome (+40 pv) : 2 chromofibres + plaque cyber + 10 credits")
		fmt.Println("5. Bottes de runner (+15 pv) : antenne + gaine isolante + 10 credits")
		fmt.Println("6. Cybersemelles (+25 pv) : 2 gaines isolantes + chromofibre + 10 credits")
		fmt.Println("0. Retour")
		fmt.Print("Choix : ")
		switch readInt() {
		case 1:
			craft(p, coqueNeurale, map[string]int{antenna: 1, gaine: 1})
		case 2:
			craft(p, neuroblindage, map[string]int{antenna: 2, chromofibre: 1})
		case 3:
			craft(p, tuniqueCyber, map[string]int{antenna: 2, plaqueCyber: 1})
		case 4:
			craft(p, blindageChrome, map[string]int{chromofibre: 2, plaqueCyber: 1})
		case 5:
			craft(p, bottesRunner, map[string]int{antenna: 1, gaine: 1})
		case 6:
			craft(p, cybersemelles, map[string]int{gaine: 2, chromofibre: 1})
		case 0:
			return
		}
	}
}
