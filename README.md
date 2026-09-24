# Projet-Red

Présentation du Projet
Neon Protocol est un jeu de rôle interactif en ligne de commande proposant un système de combat au tour par tour, de la gestion d'inventaire, de l'artisanat et un arbre de compétences basé sur des implants cybernétiques. Le joueur navigue à travers différents menus pour améliorer son personnage, combattre des monstres et affronter des boss scénarisés issus de l'univers cyberpunk.  


Structure des Fichiers
Le projet est divisé en plusieurs fichiers distincts pour séparer la logique de jeu :
- main.go : Point d'entrée de l'application. Il gère la boucle principale du menu du jeu et la fonction permettant de nettoyer le terminal selon le système d'exploitation.   
- character.go : Gère la création du personnage, la définition des statistiques de base selon la classe choisie, le calcul des points de vie maximum, le gain d'expérience et l'apprentissage des implants.   
- combat.go : Contient toute la logique des affrontements. Il gère les tours du joueur et des ennemis, les attaques utilisant de la mémoire RAM, les altérations d'état (comme le poison ou l'étourdissement) et les rencontres de boss spécifiques (ex: Saburo Arasaka, Adam Smasher).   
- inventory.go : Permet la gestion des objets. Les fonctions incluent l'ajout et la suppression d'objets, l'utilisation de consommables (polyconducteurs, restauram), l'équipement d'armures et l'amélioration de la capacité maximale de l'inventaire.   
- shop.go : Gère l'économie du jeu avec la boutique "Scavenger" pour acheter des objets avec des crédits, ainsi que "l'Assembleur cybernétique" qui permet de fabriquer de l'équipement à partir de matériaux récoltés.  
- types.go : Déclare les structures de données principales (Character, Equipment, Monster, Attack) ainsi que les constantes pour les noms d'objets, de sorts et les bonus d'équipement.   
- utils.go : Fournit des fonctions utilitaires pour la lecture des saisies utilisateur (textes et nombres entiers), le formatage des chaînes de caractères et la vérification des bornes de valeurs.   

Mécaniques de JeuClasses de Personnage : Lors de la création de profil, le joueur peut choisir parmi trois voies : le Mercenaire corpo (humain équilibré), le Puceur ombre (elfe spécialisé en RAM et initiative) ou le Biojuggernaut (nain ultra résistant avec beaucoup de points de vie). Chaque classe débute avec une compétence spécifique. 
Combats et Implants : Les affrontements se déroulent au tour par tour, l'ordre d'action étant défini par la statistique d'initiative. Le joueur utilise des attaques de base ou des compétences d'implants qui consomment de la RAM, avec des effets variés comme infliger de lourds dégâts, étourdir la cible (Blackout) ou octroyer des tours supplémentaires (Polymyeline).
 Un mode de combat automatique est également disponible.   
 Équipement et Artisanat : L'équipement se divise en trois emplacements : Tête, Corps et Pieds. L'Assembleur cybernétique permet de combiner des matériaux (comme des antennes de corpodrone ou de la chromofibre) moyennant des crédits pour créer des pièces d'armure qui augmentent les points de vie maximaux du joueur.   

 Concernant les commit et les push , il se peut qu'ils soient tous enregistrés avec mon pseudo (gabinho_14) car nous avons travaillé avec une extension de vs code appelé Live Share qui permet de travailler a plusieurs sur un dossier.
 
 Installation et ExécutionPour lancer le jeu, assurez-vous d'avoir Go installé sur votre machine. Placez-vous dans le répertoire contenant tous les fichiers du projet et exécutez la commande suivante : go run .
 Et comme on est pas des bêtes on a penser a faire un executable pour que ce soit plus simple.