// Module untuk task model
package models

type Task struct {
	IdTask      string
	Judul       string
	Deskripsi   string
	PIC         string
	TglDeadline string
	Status      string
}

var Tugas = []Task{}
