package calculator

type Items struct {
	Name               string
	BluePrint          bool
	BluePrintCostTable int
	BluePrintCostBench int
	Workbench          bool
	WorkbenchLevel     string
	Stackable          int
	Craftable          bool
	AmountCraft        int
	RessourceNeeded    []string
	AmountNeeded       []int
	DropList           []string
	DropRate           []int
}
