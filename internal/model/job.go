package model

type Job string

const (
	// Tanks
	WAR = Job("WAR")
	PLD = Job("PLD")
	DRK = Job("DRK")
	GNB = Job("GNB")

	// Healers
	WHM = Job("WHM")
	SCH = Job("SCH")
	AST = Job("AST")
	SGE = Job("SGE")

	//Melee
	MNK = Job("MNK")
	DRG = Job("DRG")
	NIN = Job("NIN")
	SAM = Job("SAM")
	RPR = Job("RPR")
	VPR = Job("VPR")

	//Caster
	BLM = Job("BLM")
	SMN = Job("SMN")
	RDM = Job("RDM")
	PCT = Job("PCT")

	//PRanged
	BRD = Job("BRD")
	MCH = Job("MCH")
	DNC = Job("DNC")
)
