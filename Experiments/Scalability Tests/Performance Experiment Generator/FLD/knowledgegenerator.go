package main

import (
	"encoding/xml"
	"io/ioutil"
	"strconv"
	"strings"
)

type Delivery struct {
	XMLName xml.Name `xml:"Delivery"`
	Name string `xml:"name"`
	Patient string `xml:"patient"`
	DeliveryLocation string `xml:"delivery_location"`
}

type Location struct {
	XMLName xml.Name `xml:"Location"`
	Name string `xml:"name"`
	Patient string `xml:"patient"`
}

type Patient struct {
	XMLName xml.Name `xml:"Patient"`
	Name string `xml:"name"`
	CanFetch bool `xml:"can_fetch"`
	CanOpen bool `xml:"can_open"`
}

type WorldDb struct {
	XMLName  xml.Name `xml:"world_db"`
	Deliveries []*Delivery
	Locations    []*Location
	Patients []*Patient
}

func MakeDelivery(name, patient, deliveryLocation string) Delivery {
	return Delivery{
		Name: name,
		Patient: patient,
		DeliveryLocation: deliveryLocation,
	}
}

func MakeLocation(name, patient string) Location {
	return Location{
		Name: name,
		Patient: patient,
	}
}

func MakePatient(name string) Patient {
	return Patient{
		Name: name,
		CanOpen: true,
		CanFetch: true,
	}
}

func main() {
	deliveriesNumber := [13]int{2, 3, 5, 10, 20, 50, 100, 200, 1000, 2000,3000,4000,5000}

	kitchenLocation := MakeLocation("Kitchen", "")

	fileIndex := 0
	for _, deliveryNumber := range deliveriesNumber {
		// Add all elements to a WorldDb type struct
		var worldDb WorldDb

		patientCounter := 1
		roomCounter := 1

		for i := 1; i <= deliveryNumber; i++ {
			var deliveryNameBuilder strings.Builder

			deliveryNameBuilder.WriteString("Delivery")
			deliveryNameBuilder.WriteString(strconv.Itoa(i))

			var deliveryPatientName strings.Builder
			var deliveryLocationName strings.Builder

			deliveryPatientName.WriteString("Patient")
			deliveryPatientName.WriteString(strconv.Itoa(patientCounter))
			patientCounter++

			deliveryLocationName.WriteString("Location")
			deliveryLocationName.WriteString(strconv.Itoa(roomCounter))
			roomCounter++

			patient := MakePatient(deliveryPatientName.String())
			location := MakeLocation(deliveryLocationName.String(), patient.Name)
			delivery := MakeDelivery(deliveryNameBuilder.String(), patient.Name, location.Name)


			worldDb.Patients = append(worldDb.Patients, &patient)
			worldDb.Locations = append(worldDb.Locations, &location)
			worldDb.Deliveries = append(worldDb.Deliveries, &delivery)
		}

		worldDb.Locations = append(worldDb.Locations, &kitchenLocation)

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
