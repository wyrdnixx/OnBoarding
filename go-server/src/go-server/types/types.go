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
	Id int `json:"id"`
	Enabled int `json:"enabled"`
}

///////// Firmen
type Firma struct {
	Id   int `json:"id"`
	Name string `json:"name"`
	Enabled int `json:"enabled"`
}
type Firmen struct {
	Firmen   []Firma `json:"Firmen"`	
}

type NewCompany struct {
	Name string `json:"name"`
}



////////// Abteilungen
type Department struct {
	Id   int `json:"id"`
	Name string `json:"name"`
	Firma   string `json:"firma"`
	NotifyMail string `json:"notifyMail"`
	Enabled int `json:"enabled"`
}
type Departments struct {
	Departments   []Department `json:"Department"`	
}

type NewDepartment struct {
	NewDepName string `json:"newDepName"`
	NewDepCompany   int `json:"newDepCompany"`
	NewDepMail string `json:"newDepMail"`
}

type NewDeps struct {
	NewDeps   []NewDepartment `json:"NewDepartment"`
	//NewDeps   []string `json:"NewDepartment"`
}