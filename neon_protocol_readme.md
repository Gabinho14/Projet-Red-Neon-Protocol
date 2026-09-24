# NEON PROTOCOL

## Le contexte dans lequel il est fait
Ce projet a été réalisé dans le cadre d'un cours d'apprentissage du langage de programmation Go (Golang). L'objectif étant de concevoir un jeu en tour par tour avec les bases du language go , en apprenant. Le jeu plonge le joueur dans un univers cyberpunk.

## Les fonctionnalités principales
- **Création de personnage :** Saisie du nom et choix de la classe parmi trois profils uniques (Mercenaire corpo, Puceur ombre, Biojuggernaut), chacun ayant ses propres caractéristiques (PV, Initiative, sorts spécifiques).
- **Système de progression :** Gain d'expérience (XP) et de crédits via les combats, montée en niveau augmentant les statistiques (PV maximums, RAM maximum).
- **Système de combat au tour par tour :** Affrontement contre différents ennemis (Corpodrone, Brute chromée...) ou boss (Saburo, Yorinobu, Adam Smasher) avec gestion de l'initiative, des points de vie (PV) et de l'énergie (RAM) pour lancer des attaques spéciales/implants. Un mode "combat automatique" est également disponible.
- **Gestion de l'inventaire et de l'équipement :** Utilisation de consommables (potions de soin, de RAM, poisons), équipement d'armures (tête, corps, pieds) pour améliorer les statistiques, et possibilité d'améliorer la taille de l'inventaire.
- **Boutique et Artisanat (Crafting) :** Un marché (Scavenger) pour acheter des objets avec des crédits, et un assembleur cybernétique (Workshop) pour fabriquer des équipements avancés en combinant des ressources (antennes, chromofibres, etc.).

## Les prérequis d'installation
- **Version de Go :** Go 1.16 ou supérieur est recommandé. Le projet n'utilise que la bibliothèque standard de Go.
- **Logiciel tiers :** Aucun logiciel supplémentaire n'est requis. Un simple terminal (Invite de commandes, PowerShell, bash, etc.) suffit pour jouer.
- **Base de données :** **Aucune base de données** n'est nécessaire. L'intégralité des données du jeu (personnage, inventaire, progression) est gérée en mémoire (RAM) le temps de l'exécution du programme.

## Le guide d'installation/lancement
1. Téléchargez ou clonez le code source du projet dans un dossier sur votre machine.
2. Ouvrez votre terminal et naviguez jusqu'au dossier contenant les fichiers du projet.
3. Pour lancer le jeu directement sans générer de fichier exécutable, tapez la commande suivante :
   ```bash
   go run .
   ```
4. *Alternative :* 
Un éxecutable est présent directement dans le dossier , ne nécessitant aucune ligne de commande pour le lancer !

## L'organisation du code source
Le projet est découpé en plusieurs fichiers `.go` pour séparer la logique et faciliter la maintenance :

- `main.go` : Point d'entrée du programme. Gère le menu principal, la boucle de jeu et l'effacement de l'écran du terminal.
- `character.go` : Gère la création du personnage, l'affichage du profil, le gain d'XP (montée en niveau) et le calcul des statistiques (PV max).
- `combat.go` : Contient toute la logique des affrontements (tours de jeu, attaques, patterns des ennemis, combats de boss et mode automatique).
- `inventory.go` : Gère l'inventaire du joueur (ajout/retrait d'objets, utilisation des consommables, équipement des armures).
- `shop.go` : Gère les interactions marchandes (le Scavenger pour l'achat) et l'assembleur cybernétique (fabrication/crafting d'équipement).
- `types.go` : Regroupe les constantes (couleurs du terminal, noms des objets) et la définition de toutes les structures de données (`Character`, `Monster`, `Attack`, `Equipment`).
- `utils.go` : Contient des fonctions utilitaires réutilisables (lecture des saisies clavier, formatage du texte, vérifications dans les listes, mathématiques simples).