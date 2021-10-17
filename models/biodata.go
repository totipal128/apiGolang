package models

type Biodata struct{
	Id				int 	`json:"id"`
	Name			string 	`json:"name"`
	Age				int		`json:"age"`
	Address			string	`json:"address"`
	Phone_number	string	`json:"phoneNumber"`
}