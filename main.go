package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Items struct {
	Item []struct {
		DLC                bool      `json:"dlc"`
		Name               string    `json:"name"`
		BluePrint          bool      `json:"blueprint"`
		BluePrintCostTable int       `json:"blueprint_cost_table"`
		BluePrintCostBench int       `json:"blueprint_cost_bench"`
		Workbench          bool      `json:"workbench"`
		WorkbenchLevel     string    `json:"workbench_level"`
		Stackable          int       `json:"stackable"`
		Craftable          bool      `json:"craftable"`
		AmountCraft        int       `json:"amount_craft"`
		RessourceNeeded    []string  `json:"ressource_needed"`
		AmountNeeded       []int     `json:"amount_needed"`
		DropList           []string  `json:"droplist"`
		DropRate           []float64 `json:"droprate"`
	} `json:"Item"`
}

type Structures struct {
	Structure []struct {
		StructureName          string   `json:"structure_name"`
		Hp                     int      `json:"hp"`
		RecommandedStuffList   []string `json:"recommanded_stuff_list"`
		RecommandedStuffAmount []int    `json:"recommanded_stuff_amount"`
		SulfurCost             int      `json:"sulfur_cost"`
	} `json:"Structure"`
}

type Ressources struct {
	Ressource []struct {
		RessourceName string   `json:"ressource_name"`
		StuffList     []string `json:"stuff_list"`
		Object        []int    `json:"object"`
		Amount        int      `json:"amount"`
	} `json:"Ressource"`
}

type SmeltOres struct {
	SmeltOre []struct {
		OreName string   `json:"ore_name"`
		Furnace []string `json:"furnace"`
		Amount  int      `json:"amount"`
	} `json:"SmeltOre"`
}

type Waters struct {
	Water []struct {
		WaterName string   `json:"water_name"`
		Furnace   []string `json:"water_catcher"`
		AmountMax []int    `json:"amount_max"`
		AmountMin []int    `json:"amount_main"`
	} `json:"Water"`
}

type Composting struct {
	Compost []struct {
		ItemName string   `json:"item_name"`
		Items    []string `json:"item"`
		Amount   []int    `json:"amount"`
	} `json:"Compost"`
}

type AnimalLoots struct {
	AnimalLoot []struct {
		ItemName string   `json:"item_name"`
		Animal   []string `json:"animal"`
		Amount   []int    `json:"amount"`
	} `json:"AnimalLoot"`
}

func main() {
	data, _ := ioutil.ReadFile("structures.json")
	var structures Structures
	json.Unmarshal([]byte(data), &structures)
	fmt.Print(structures)
	data2, _ := ioutil.ReadFile("item.json")
	var items Items
	json.Unmarshal([]byte(data2), &items)
	fmt.Print(items)
}
