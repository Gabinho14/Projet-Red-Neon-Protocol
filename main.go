package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	player := characterCreation()
	for {
		ClearTerminal()
		fmt.Println("\n=== NEON PROTOCOL ===")
		fmt.Println("1. Profil")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Scavenger")
		fmt.Println("4. Assembleur cybernetique")
		fmt.Println("5. Entrainement")
		fmt.Println("6. Missions de combat")
		fmt.Println("7. Boss")
		fmt.Println("8. Qui sont-ils ?")
		fmt.Println("9. Quitter")
		fmt.Print("Choix : ")

		switch readInt() { // readInt est définie dans utils.go
		case 1:
			displayInfo(&player) // Dans character.go
		case 2:
			accessInventory(&player) // Dans inventory.go
		case 3:
			market(&player) // Dans shop.go
		case 4:
			workshop(&player) // Dans shop.go
		case 5:
			trainingFight(&player) // Dans combat.go
		case 6:
			missions(&player) // Dans combat.go
		case 7:
			bosses(&player) // Dans combat.go
		case 8:
			fmt.Println("Artistes caches : ABBA et Daft Punk.")
		case 9:
			fmt.Println("Deconnexion.")
			return
		}
	}
}

func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
