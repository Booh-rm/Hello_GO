package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	items []int
}

func (q *Queue) EnqueueAtStart(item int) {
	q.items = append([]int{item}, q.items...)
}

func (q *Queue) EnqueueAtCenter(item int) error {
	if len(q.items) == 0 {
		return errors.New("queue is empty")
	}
	midIndex := len(q.items) / 2
	q.items = append(q.items[:midIndex], append([]int{item}, q.items[midIndex:]...)...)
	return nil
}

func (q *Queue) EnqueueAtEnd(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	queue := Queue{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Selecciona una opción:")
		fmt.Println("1. Agregar número al inicio")
		fmt.Println("2. Agregar número al centro")
		fmt.Println("3. Agregar número al final")
		fmt.Println("4. Eliminar número de la cola")
		fmt.Println("5. Vaciar la cola")
		fmt.Println("6. Salir")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("Ingrese el número que desea agregar al inicio:")
			startInput, _ := reader.ReadString('\n')
			startInput = strings.TrimSuffix(startInput, "\n")
			startInput = strings.TrimSpace(startInput)
			startInput = strings.ReplaceAll(startInput, " ", "")
			startNumber, err := strconv.Atoi(startInput)
			if err != nil {
				fmt.Println("Error: el valor ingresado no es un número válido")
				continue
			}
			queue.EnqueueAtStart(startNumber)
			fmt.Println("Número agregado al inicio de la cola")
		case "2":
			fmt.Println("Ingrese el número que desea agregar al centro:")
			centerInput, _ := reader.ReadString('\n')
			centerInput = strings.TrimSpace(centerInput)
			centerInput = strings.ReplaceAll(centerInput, " ", "")
			centerNumber, err := strconv.Atoi(centerInput)

			if err != nil {
				fmt.Println("Error: el valor ingresado no es un número válido")
				continue
			}
			err = queue.EnqueueAtCenter(centerNumber)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Número agregado al centro de la cola")
		case "3":
			fmt.Println("Ingrese el número que desea agregar al final:")
			endInput, _ := reader.ReadString('\n')
			endInput = strings.TrimSpace(endInput)
			endInput = strings.ReplaceAll(endInput, " ", "")
			endNumber, err := strconv.Atoi(endInput)
			if err != nil {
				fmt.Println("Error: el valor ingresado no es un número válido")
				continue
			}
			queue.EnqueueAtEnd(endNumber)
			fmt.Println("Número agregado al final de la cola")
		case "4":
			item, err := queue.Dequeue()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Número eliminado de la cola:", item)
		case "5":
			queue.items = []int{}
			fmt.Println("Cola vaciada")
		case "6":
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción inválida, por favor seleccione una opción válida.")
		}

		fmt.Println("Cola actual:", queue.items)
	}
}
