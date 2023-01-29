package calculator

type Items struct {
	DLC                bool     `json:"dlc"`
	Name               string   `json:"name"`
	BluePrint          bool     `json:"blueprint"`
	BluePrintCostTable int      `json:"blueprint_cost_table"`
	BluePrintCostBench int      `json:"blueprint_cost_bench"`
	Workbench          bool     `json:"workbench"`
	WorkbenchLevel     string   `json:"workbench_level"`
	Stackable          int      `json:"stackable"`
	Craftable          bool     `json:"craftable"`
	AmountCraft        int      `json:"amount_craft"`
	RessourceNeeded    []string `json:"ressource_needed"`
	AmountNeeded       []int    `json:"amount_needed"`
	DropList           []string `json:"droplist"`
	DropRate           []int    `json:"droprate"`
}

type Wall struct {
	Material string
	Hp       int
}

type ExternalWall struct {
	Material string
	Hp       int
}

type Door struct {
	Material string
	Hp       int
}

type ExternalDoor struct {
	Material string
	Hp       int
}

type Fondation struct {
	Material string
	Hp       int
}

type Celling struct {
	Material string
	Hp       int
}
