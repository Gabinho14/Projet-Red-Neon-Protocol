package main

import "fmt"

func baseHP(class string) int {
	switch class {
	case "Puceur ombre":
		return 130
	case "Biojuggernaut":
		return 200
	}
	return 165
}
func recalc(p *Character) {
	p.MaxHP = baseHP(p.Class) + (p.Level-1)*10
	p.MaxHP += hpBonus[p.Equipment.Head] + hpBonus[p.Equipment.Body] + hpBonus[p.Equipment.Feet]
	p.HP = clamp(p.HP, 0, p.MaxHP)
}
func initCharacter(name, class string, level, maxHP, hp int, maxRAM float64, initiative int, inventory []string, skills []string) Character {
	return Character{
		Name:       name,
		Class:      class,
		Level:      level,
		HP:         hp,
		MaxHP:      maxHP,
		Gold:       100,
		RAM:        maxRAM,
		MaxRAM:     maxRAM,
		XPMax:      30,
		Initiative: initiative,
		Inventory:  inventory,
		Capacity:   10,
		Skills:     skills, // <-- MODIFIÉ ICI : on utilise le paramètre
	}
}
func characterCreation() Character {
	ClearTerminal()
	fmt.Println("╔═══════════════════════════════════════╗")
	fmt.Println("║       INITIALISATION DU SYSTEME       ║")
	fmt.Println("╚═══════════════════════════════════════╝")
	fmt.Println()

	var name string
	for !letters(name) {
		fmt.Print("[>] Saisissez votre identifiant (Nom) : ")
		name = readText()
		if !letters(name) {
			fmt.Println("[!] Erreur : Veuillez utiliser uniquement des lettres.")
		}
	}

	ClearTerminal()
	fmt.Println("╔═══════════════════════════════════════╗")
	fmt.Println("║         SELECTION DU PROFIL           ║")
	fmt.Println("╚═══════════════════════════════════════╝")
	fmt.Printf("\nBienvenue, %s. Choisissez votre voie :\n\n", formatName(name))

	fmt.Println("  [1] MERCENAIRE CORPO (Humain)")
	fmt.Println("      -> 165 PV | Initiative : 5 | Profil equilibre")
	fmt.Println("  [2] PUCEUR OMBRE (Elfe)")
	fmt.Println("      -> 130 PV | Initiative : 8 | Specialiste RAM")
	fmt.Println("  [3] BIOJUGGERNAUT (Nain)")
	fmt.Println("      -> 200 PV | Initiative : 3 | Ultra resistant")
	fmt.Println() // Ajoute une ligne vide proprement

	fmt.Print("[>] Confirmation du choix (1/2/3) : ")

	class := "Mercenaire corpo"
	initiative := 5
	classSkill := sortCorpo // <-- Sort par défaut (Corpo)

	switch readInt() {
	case 2:
		class = "Puceur ombre"
		initiative = 8
		classSkill = sortOmbre // <-- Sort de l'Elfe
	case 3:
		class = "Biojuggernaut"
		initiative = 3
		classSkill = sortBio // <-- Sort du Nain
	}

	maxHP := baseHP(class)

	ClearTerminal()
	fmt.Printf("\n[OK] Profil %s enregistre avec succes.\n", class)

	// On met à jour l'appel à initCharacter en lui donnant les compétences : "Coup de poing" + le sort de classe
	return initCharacter(formatName(name), class, 1, maxHP, maxHP/2, 4, initiative,
		[]string{healPotion, healPotion, healPotion}, []string{"Coup de poing", classSkill})
}
func displayInfo(p *Character) {
	fmt.Println("\n=== PROFIL ===")
	fmt.Println("Nom :", p.Name)
	fmt.Println("Classe :", p.Class)
	fmt.Printf("Niveau : %d | XP : %d/%d\n", p.Level, p.XP, p.XPMax)
	fmt.Printf("PV : %d/%d | RAM : %.1f/%.1f\n", p.HP, p.MaxHP, p.RAM, p.MaxRAM)
	fmt.Printf("Credits : %d | Initiative : %d\n", p.Gold, p.Initiative)
	fmt.Println("Casque :", p.Equipment.Head)
	fmt.Println("Veste :", p.Equipment.Body)
	fmt.Println("Bottes :", p.Equipment.Feet)
	fmt.Println("Implants :", p.Skills)
}
func isDead(p *Character) bool {
	if p.HP > 0 {
		return false
	}
	p.HP = p.MaxHP / 2
	fmt.Printf("%s est KO et revient avec %d PV.\n", p.Name, p.HP)
	return true
}
func gainXP(p *Character, value int) {
	p.XP += value
	for p.XP >= p.XPMax {
		p.XP -= p.XPMax
		p.Level++
		p.XPMax += 20
		if p.Level%3 == 0 {
			p.MaxRAM++
		}
		p.RAM = p.MaxRAM
		recalc(p)
		p.HP = p.MaxHP
		fmt.Println("NIVEAU SUPERIEUR ! Niveau :", p.Level)
	}
}
func learnImplant(p *Character, name string) {
	if has(p.Skills, name) {
		fmt.Println("Implant deja installe.")
		return
	}
	p.Skills = append(p.Skills, name)
	fmt.Println("Implant installe :", name)
}
