package main

const (
	healPotion      = "Polyconducteur"
	manaPotion      = "Restauram"
	poisonPotion    = "Corrupteur bionique"
	pyrocanon       = "Pyrocanon"
	blackout        = "Blackout"
	polymyeline     = "Polymyeline"
	antenna         = "Antenne de corpodrone"
	plaqueCyber     = "Plaque cybernetique"
	gaine           = "Gaine isolante"
	chromofibre     = "Chromofibre"
	upgrade         = "Extension inventaire"
	coqueNeurale    = "Coque neurale"
	neuroblindage   = "Neuroblindage"
	tuniqueCyber    = "Tunique cybernetique"
	blindageChrome  = "Blindage chrome"
	bottesRunner    = "Bottes de runner"
	cybersemelles   = "Cybersemelles"
	heritageRebecca = "Héritage de Rebecca"
	etherHealPotion = "Polyconducteur éthéré"
	etherRamPotion  = "Restauram éthéré"
	connecteur      = "Connecteur de données"
	sortCorpo       = "Tir suppressif"
	sortOmbre       = "Surcharge neurale"
	sortBio         = "Onde de choc"
)

var hpBonus = map[string]int{
	coqueNeurale:   10,
	neuroblindage:  20,
	tuniqueCyber:   25,
	blindageChrome: 40,
	bottesRunner:   15,
	cybersemelles:  25,
}

type Equipment struct {
	Head, Body, Feet string
}
type Character struct {
	Name, Class        string
	Level, HP, MaxHP   int
	Gold               int
	RAM, MaxRAM        float64
	XP, XPMax          int
	Initiative         int
	Inventory          []string
	Capacity, Upgrades int
	Skills             []string
	Equipment          Equipment
	ExtraTurns         int
}
type Monster struct {
	Name                                    string
	HP, MaxHP, Attack, Initiative, XP, Gold int
	Stunned                                 int
}
type Product struct {
	Name  string
	Price int
}
type Attack struct {
	Name    string
	Cost    float64
	Damage  int
	Implant bool
}

var attacks = []Attack{
	{"Attaque basique", 0, 5, false},
	{"Coup de poing", 0.5, 8, false},
	{pyrocanon, 1, 18, true},
	{blackout, 1.5, 8, true},
	{polymyeline, 1, 10, true},
	{sortCorpo, 1, 12, true},
	{sortOmbre, 2, 22, true},
	{sortBio, 1.5, 16, true},
}
