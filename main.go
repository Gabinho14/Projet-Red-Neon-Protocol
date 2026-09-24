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
		fmt.Println(cyan + "╔═══════════════════════════════════════╗" + reset)
		fmt.Println(cyan + "║             NEON PROTOCOL             ║" + reset)
		fmt.Println(cyan + "╚═══════════════════════════════════════╝\n" + reset)
		fmt.Println(vert + "1. Profil" + reset)
		fmt.Println(vert + "2. Inventaire" + reset)
		fmt.Println(jaune + "3. Scavenger" + reset)
		fmt.Println(jaune + "4. Assembleur cybernetique" + reset)
		fmt.Println(rouge + "5. Entrainement" + reset)
		fmt.Println(rouge + "6. Missions de combat" + reset)
		fmt.Println(rouge + "7. Boss" + reset)
		fmt.Println(magenta + "8. Qui sont-ils ?" + reset)
		fmt.Println("9. Quitter")
		fmt.Print("Choix : ")

		switch readInt() {
		case 1:
			displayInfo(&player)
			fmt.Println("\nAppuyez sur Entrée et saisissez un caractère pour continuer...")
			readText()
		case 2:
			accessInventory(&player)
		case 3:
			market(&player)
		case 4:
			workshop(&player)
		case 5:
			trainingFight(&player)
		case 6:
			missions(&player)
		case 7:
			bosses(&player)
		case 8:
			fmt.Println("Artistes caches : ABBA et Steven Spielberg.")
			fmt.Println("\nAppuyez sur Entrée et saisissez un caractère pour continuer...")
			readText()
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
