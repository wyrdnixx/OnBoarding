package types

// Configuration -  Allgemeine Config
type Configuration struct {
	ApiPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Info       string
}

type ToggleEntry struct {
	Table   string `json:"table"`
	Id      int    `json:"id"`
	Enabled int    `json:"enabled"`
}

///////// Firmen
type Firma struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Enabled int    `json:"enabled"`
}
type Firmen struct {
	Firmen []Firma `json:"Firmen"`
}

type NewCompany struct {
	Name string `json:"name"`
}

////////// Abteilungen
type Department struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Firma      string `json:"firma"`
	NotifyMail string `json:"notifyMail"`
	Enabled    int    `json:"enabled"`
}
type Departments struct {
	Departments []Department `json:"Department"`
}

type NewDepartment struct {
	NewDepName    string `json:"newDepName"`
	NewDepCompany int    `json:"newDepCompany"`
	NewDepMail    string `json:"newDepMail"`
}

type NewDep struct {
	Dep NewDepartment `json:"NewDepartment"`
}

type Processors struct {
Processors []Processor ` json:"Processor"`
}

//// Processors

type Processor struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	NotifyMail string `json:"notifyMail"`
	Enabled    int    `json:"enabled"`
}

type NewProcessor struct {
	NewProcessorName    string `json:"NewProcessorName"`	
	NewProcessorMail    string `json:"NewProcessorMail"`
}

type NewProc struct {
	Item NewProcessor `json:"NewProcessor"`
}

type Items struct {
	Items []Item ` json:"Item"`
}

type NewItem struct {
	Id	int	`json:"id"`
	ProcessorId	int	`json:"selectedProcessor"`
	DepId	int	`json:"selectedDepartment"`
	Text	string	`json:"newItemText"`
	ItemType	string	`json:"selectedType"`	
	Enabled	int	`json:"enabled"`
}

type NewITM struct {
	Itm NewItem `json:"NewItem"`
}

type Item struct {
	Id	int	`json:"id"`
	ProcessorId	string	`json:"ProcessorId"`
	DepId	string	`json:"DepId"`
	Text	string	`json:"Text"`
	ItemType	string	`json:"ItemType"`	
	Enabled	int	`json:"enabled"`
}