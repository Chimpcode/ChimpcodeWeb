package models

// User struct
type User struct {
	ID int `storm:"id,increment"`

	Name        string //`json:"Name"`
	Email       string //`json:"Email"`
	Phonenumber string //`json:"Phonenumber"`
	// Dni           string   //`json:"Dni"`
	DeveloperType []string //`json:"DeveloperType"`
	// TechStacks     []string //`json:"TechStacks"`

	LastJob  string //`json:"LastJob"`
	Job      string //`json:"TechStacks"`
	SinceJob string //`json:"TechStacks"`
}
