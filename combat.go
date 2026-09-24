package main

import (
	"fmt"
	"strings"
	"time"
)

func poisonEnemy(m *Monster) {
	fmt.Println("Corrupteur bionique envoye sur", m.Name)
	for i := 1; i <= 3 && m.HP > 0; i++ {
		time.Sleep(time.Second)
		m.HP = clamp(m.HP-10, 0, m.MaxHP)
		fmt.Printf("Cycle %d : %s perd 10 PV. PV : %d/%d\n", i, m.Name, m.HP, m.MaxHP)
	}
}

func initGoblin() Monster {
	return Monster{
		Name:       "Bandit blinde d'entrainement",
		HP:         40,
		MaxHP:      40,
		Attack:     5,
		Initiative: 4,
		XP:         20,
		Gold:       15,
	}
}
func goblinPattern(p *Character, m *Monster, turn int) {
	if m.Stunned > 0 {
		m.Stunned--
		fmt.Printf("%s est stun (Blackout) et ne peut pas agir.\n", m.Name)
		return
	}

	damage := m.Attack

	// Pattern spécifique pour Adam Smasher (écrasement tous les 4 tours)
	if m.Name == "Adam Smasher" {
		if turn%4 == 0 {
			damage = 150
			fmt.Println("💥 ÉCRASSEMENT D'ADAM SMASHER !")
		}
	} else { // Pattern classique pour les autres monstres/boss (double dégâts tous les 3 tours)
		if turn%3 == 0 {
			damage *= 2
			fmt.Println("Attaque speciale ennemie !")
		}
	}

	p.HP = clamp(p.HP-damage, 0, p.MaxHP)
	fmt.Printf("%s inflige a %s %d de degats\n", m.Name, p.Name, damage)
	fmt.Printf("%s PV : %d/%d\n", p.Name, p.HP, p.MaxHP)
}
func hit(p *Character, m *Monster, attack string, damage int) {
	m.HP = clamp(m.HP-damage, 0, m.MaxHP)
	fmt.Printf("%s utilise %s : %d degats a %s.\n", p.Name, attack, damage, m.Name)
	fmt.Printf("%s PV : %d/%d\n", m.Name, m.HP, m.MaxHP)
}
func effect(a Attack) string {
	switch a.Name {
	case blackout:
		return "stun 1 tour"
	case polymyeline:
		return "rejouer un tour"
	}
	return fmt.Sprintf("%d degats", a.Damage)
}
func attackMenu(p *Character, m *Monster) bool {
	fmt.Println()
	for i, a := range attacks {
		fmt.Printf("%d. %s : %s, %.1f RAM\n", i+1, a.Name, effect(a), a.Cost)
	}
	fmt.Println("0. Retour")
	fmt.Print("Action : ")
	choice := readInt()
	if choice < 1 || choice > len(attacks) {
		return false
	}
	a := attacks[choice-1]
	if a.Implant && !has(p.Skills, a.Name) {
		fmt.Println("Implant non installe.")
		return false
	}
	if p.RAM < a.Cost {
		fmt.Println("RAM insuffisante.")
		return false
	}
	p.RAM -= a.Cost
	switch a.Name {
	case blackout:
		m.Stunned = 1
		fmt.Printf("%s utilise Blackout : %s est stun 1 tour.\n", p.Name, m.Name)
	case polymyeline:
		p.ExtraTurns++
		fmt.Printf("%s active Polymyeline : un tour supplementaire est accorde.\n", p.Name)
	default:
		hit(p, m, a.Name, a.Damage)
	}
	if a.Cost > 0 {
		fmt.Printf("RAM : %.1f/%.1f\n", p.RAM, p.MaxRAM)
	}
	return true
}
func fightInventory(p *Character, m *Monster) bool {
	showInventory(p)
	fmt.Print("Numero objet, 0 annuler : ")
	choice := readInt()
	if choice < 1 || choice > len(p.Inventory) {
		return false
	}
	item := p.Inventory[choice-1]
	fmt.Println("Vous utilisez", item)
	return useItem(p, item, m)
}
func charTurn(p *Character, m *Monster) bool {
	for {
		fmt.Println("\n1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Fuir")
		fmt.Print("Action : ")
		switch readInt() {
		case 1:
			if attackMenu(p, m) {
				return false
			}
		case 2:
			if fightInventory(p, m) {
				return false
			}
		case 3:
			fmt.Println("\nVous prenez la fuite lâchement...")
			time.Sleep(2 * time.Second)
			return true
		}
	}
}
func playerTurn(p *Character, m *Monster, auto bool) bool {
	fled := false
	if auto {
		autoPlayerTurn(p, m)
	} else {
		fled = charTurn(p, m)
	}

	if fled {
		return true
	}

	for p.ExtraTurns > 0 && m.HP > 0 && p.HP > 0 {
		p.ExtraTurns--
		fmt.Println("-- Tour supplementaire (Polymyeline) --")
		if auto {
			time.Sleep(2 * time.Second)
			autoPlayerTurn(p, m)
		} else {
			fled = charTurn(p, m)
			if fled {
				return true
			}
		}
	}
	return false
}
func fight(p *Character, m *Monster) bool {
	fmt.Println("\n=== COMBAT :", m.Name, "===")
	fmt.Print("Activer le combat automatique ? (o/n) : ")
	auto := strings.ToLower(readText()) == "o"
	for turn := 1; p.HP > 0 && m.HP > 0; turn++ {
		fmt.Printf("\nTour %d | Vous %d/%d PV | %s %d/%d PV\n",
			turn, p.HP, p.MaxHP, m.Name, m.HP, m.MaxHP)
		if p.Initiative >= m.Initiative {
			if playerTurn(p, m, auto) {
				return false
			}
			ClearTerminal()
			if m.HP > 0 {
				goblinPattern(p, m, turn)
			}
		} else {
			goblinPattern(p, m, turn)
			if p.HP > 0 {
				if playerTurn(p, m, auto) {
					return false
				}
			}
		}
	}
	if m.HP <= 0 {
		fmt.Println("VICTOIRE contre", m.Name)
		p.Gold += m.Gold
		gainXP(p, m.XP)
		fmt.Printf("+%d XP et +%d credits\n", m.XP, m.Gold)
		return true
	} else {
		isDead(p)
		fmt.Println("Combat perdu.")
		return false
	}
}
func trainingFight(p *Character) {
	m := initGoblin()
	fight(p, &m)
}
func missions(p *Character) {
	ClearTerminal()
	fmt.Println("\n=== MISSIONS ===")
	fmt.Println("1. Corpodrone")
	fmt.Println("2. Cabot sentinelle")
	fmt.Println("3. Brute chromee")
	fmt.Println("0. Retour")
	fmt.Print("Choix : ")
	var m Monster
	switch readInt() {
	case 1:
		m = Monster{Name: "Corpodrone", HP: 60, MaxHP: 60, Attack: 7, Initiative: 5, XP: 30, Gold: 25}
	case 2:
		m = Monster{Name: "Cabot sentinelle", HP: 80, MaxHP: 80, Attack: 10, Initiative: 6, XP: 50, Gold: 50}
	case 3:
		m = Monster{Name: "Brute chromee", HP: 120, MaxHP: 120, Attack: 14, Initiative: 7, XP: 100, Gold: 100}
	default:
		return
	}
	fight(p, &m)
}
func bosses(p *Character) {
	ClearTerminal()
	fmt.Println("\n=== COMBATS DE BOSS ===")
	fmt.Println("1. Saburo Arasaka")
	fmt.Println("2. Yorinobu Arasaka")
	fmt.Println("3. Dexter DeShawn")
	fmt.Println("0. Retour")
	fmt.Print("Choix : ")

	switch readInt() {
	case 1:
		saburo := Monster{Name: "Saburo", HP: 165, MaxHP: 165, Attack: 25, Initiative: 6, XP: 150, Gold: 200}
		fight(p, &saburo)

	case 2:
		yorinobu := Monster{Name: "Yorinobu", HP: 210, MaxHP: 210, Attack: 40, Initiative: 7, XP: 250, Gold: 300}
		if fight(p, &yorinobu) {
			if !has(p.Skills, heritageRebecca) {
				p.Skills = append(p.Skills, heritageRebecca)
				fmt.Println("\n✨ Recompense : Vous avez obtenu la competence 'Héritage de Rebecca' (75 degats, 4 RAM) !")
			}
		}

	case 3:
		dexter := Monster{Name: "Dexter", HP: 270, MaxHP: 270, Attack: 60, Initiative: 8, XP: 400, Gold: 500}
		if fight(p, &dexter) {
			fmt.Println("\n ATTENTION ! Un nouvel ennemi surgit immédiatement !")
			time.Sleep(2 * time.Second)
			smasher := Monster{Name: "Adam Smasher", HP: 350, MaxHP: 350, Attack: 70, Initiative: 10, XP: 1000, Gold: 1000}
			fight(p, &smasher)
		}
	}
}
func autoAttack(p *Character, m *Monster) {
	best := attacks[0]
	for _, a := range attacks {
		if a.Implant && !has(p.Skills, a.Name) {
			continue
		}
		if p.RAM < a.Cost {
			continue
		}
		if a.Damage > best.Damage {
			best = a
		}
	}
	p.RAM -= best.Cost
	switch best.Name {
	case blackout:
		m.Stunned = 1
		fmt.Printf("%s utilise Blackout : %s est stun 1 tour.\n", p.Name, m.Name)
	case polymyeline:
		p.ExtraTurns++
		fmt.Printf("%s active Polymyeline : un tour supplementaire est accorde.\n", p.Name)
	default:
		hit(p, m, best.Name, best.Damage)
	}

	time.Sleep(1500 * time.Millisecond)
}

func autoPlayerTurn(p *Character, m *Monster) {
	if p.HP <= p.MaxHP/3 && count(p, healPotion) > 0 {
		fmt.Println(p.Name, "utilise automatiquement un Polyconducteur.")
		takePot(p)
		return
	}
	autoAttack(p, m)
}
