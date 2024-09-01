package tickets

import (
	"errors"
	"os"
	"strconv"
	"strings"
)
// esta estructura tiene un campo de tipo puntero a un slice de Ticket
type GestorTickets struct {
	
	objetoReservas *[]Ticket
}

type Reservaciones interface {
	GetTotalTickets(destination string) (int, error)
	GetCountByPeriod(time string) (int, error)
	PercentageDestination(destination string) (float64, error)
}

type Ticket struct {
	id          int
	nombre      string
	email       string
	paisDestino string
	horaVuelo   string
	precio      int
}

// extraerData lee del disco, y crea un slice de Tickets. Ejecuta unos panics si no puede 
// leer correctamente o el contenido del archivo no es compatible con Tickets
func (gt *GestorTickets) ExtraerData() (*[]Ticket, error) {
	var listaReservas []Ticket
	archivoCsv, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic("Error en la lectura de nuestro archivo.")
	}

	dataCruda := strings.Split(string(archivoCsv), "\n")
	for lineaCruda := range dataCruda {
		registro := strings.Split(string(dataCruda[lineaCruda]), ",")

		if len(registro) < 6 {
			panic("Falla en la estructura de datos de un registro")
		}

		id, _ := strconv.Atoi(registro[0])
		precio, _ := strconv.Atoi(registro[5])

		reserva := Ticket{
			id:          id,
			nombre:      registro[1],
			email:       registro[2],
			paisDestino: registro[3],
			horaVuelo:   registro[4],
			precio:      precio,
		}
		listaReservas = append(listaReservas, reserva)
	}
	return &listaReservas, nil
}

// Esta función es llamada por las demás funciones para cargar los datos. 
// Se ejecuta una sola vez. 
func (gt *GestorTickets) getReservas() error {
	if gt.objetoReservas != nil {
		return nil
	}
	reservas, err := gt.ExtraerData()
	if err != nil {
		return err
	}
	gt.objetoReservas = reservas
	return nil
}

// ejemplo 1
func (gt *GestorTickets) GetTotalTickets(destination string) (int, error) {
	
	err := gt.getReservas()
	
	if err != nil {
		return 0, err
	}

	totalTickets := 0

	for _, reserva := range *gt.objetoReservas {
		if reserva.paisDestino == destination {
			totalTickets++
		}
	}
	return totalTickets, nil
}

// ejemplo 2
func (gt *GestorTickets) GetCountByPeriod(time string) (int, error) {
	err := gt.getReservas()
	if err != nil {
		return 0, err
	}
	switch time {
	case "madrugada":
		return gt.contarReservasPorRango(0, 6)
	case "mañana":
		return gt.contarReservasPorRango(7, 12)
	case "tarde":
		return gt.contarReservasPorRango(13, 19)
	case "noche":
		return gt.contarReservasPorRango(20, 23)
	default:
		return 0, errors.New("La categoría no existe")
	}

}


// esta función realiza la cuenta de reservas de acuerdo al rango de horas definido
func (gt *GestorTickets) contarReservasPorRango(inicial, final int) (int, error) {
	total := 0
	for _, reserva := range *gt.objetoReservas {
		horaString := strings.Split(reserva.horaVuelo, ":")[0]
		horaInt, err := strconv.Atoi(horaString)
		if err != nil {
			return 0, errors.New("Error al convertir las horas a entero: " + err.Error())
		}
		if horaInt >= inicial && horaInt <= final {
			total++
		}
	}
	return total, nil
}

// ejemplo 3
func (gt *GestorTickets) PercentageDestination(destination string) (float64, error) {
	err := gt.getReservas()
	if err != nil {
		return 0, err
	}
	totalReservas := len(*gt.objetoReservas)
	totalTickets := 0.0

	for _, reserva := range *gt.objetoReservas {
		if reserva.paisDestino == destination {
			totalTickets++
		}

	}

	porcentaje := totalTickets / float64(totalReservas) * 100
	return porcentaje, nil
}
