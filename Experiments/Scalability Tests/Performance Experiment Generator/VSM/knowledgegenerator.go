package main

import (
	"encoding/xml"
	"io/ioutil"
	"strconv"
	"strings"
)

type Patient struct {
	XMLName           xml.Name `xml:"Patient"`
	Name              string   `xml:"name"`
	InfectiousDisease bool     `xml:"infectious_disease"`
	Diabetes          bool     `xml:"diabetes"`
	PostSurgical      bool     `xml:"post_surgical"`
	Available         bool     `xml:"available"`
	Checked           bool     `xml:"checked"`
}

type Room struct {
	XMLName    xml.Name `xml:"Room"`
	Name       string   `xml:"name"`
	IsOccupied bool     `xml:"is_occupied"`
	Patients   string   `xml:"patients"`
}

type WorldDb struct {
	XMLName  xml.Name `xml:"world_db"`
	Rooms    []*Room
	Patients []*Patient
}

func MakePatient(name string) Patient {
	return Patient{
		Name:              name,
		InfectiousDisease: false,
		Diabetes:          true,
		PostSurgical:      false,
		Available:         true,
		Checked:           false,
	}
}

func MakeRoom(name string, patients string) Room {
	return Room{
		Name:       name,
		IsOccupied: false,
		Patients:   patients,
	}
}

func main() {
	roomsNumber := [13]int{2, 3, 5, 10, 20, 50, 100, 200, 1000, 2000, 3000, 4000, 5000}
	patientNumber := 2

	sanitizationRoom := MakeRoom("SanitizationRoom", "")

	fileIndex := 0
	for _, roomNumber := range roomsNumber {
		// Add all elements to a WorldDb type struct
		var worldDb WorldDb

		patientCounter := 1

		for i := 0; i < roomNumber; i++ {
			var roomNameBuilder strings.Builder

			roomNameBuilder.WriteString("Room")
			roomNameBuilder.WriteString(strconv.Itoa(i))

			var roomPatients strings.Builder

			for j := 0; j < patientNumber; j++ {
				var patientNameBuilder strings.Builder

				patientNameBuilder.WriteString("Patient")
				patientNameBuilder.WriteString(strconv.Itoa(patientCounter))
				patientCounter++

				patient := MakePatient(patientNameBuilder.String())
				worldDb.Patients = append(worldDb.Patients, &patient)

				roomPatients.WriteString(patient.Name)
				if j < patientNumber-1 {
					roomPatients.WriteString(" ")
				}
			}

			room := MakeRoom(roomNameBuilder.String(), roomPatients.String())
			worldDb.Rooms = append(worldDb.Rooms, &room)
		}

		worldDb.Rooms = append(worldDb.Rooms, &sanitizationRoom)

		var nameBuilder strings.Builder

		nameBuilder.WriteString("world_db")
		nameBuilder.WriteString(strconv.Itoa(fileIndex))
		nameBuilder.WriteString(".xml")

		data, err := xml.MarshalIndent(worldDb, "\t", "\t")
		if err != nil {
			panic(err)
		}

		_ = ioutil.WriteFile(nameBuilder.String(), data, 0644)

		fileIndex++
	}
}
