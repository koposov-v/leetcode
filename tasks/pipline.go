package tasks

import (
	"fmt"
	"sync"
)

// Определяем структуры данных для каждой стадии
type Stage1Data struct {
	Value int
}

type Stage2Data struct {
	Value int
}

// Функция для первой стадии обработки
func Stage1(input <-chan Stage1Data, output chan<- Stage2Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range input {
		// Обработка данных
		processedData := Stage2Data{Value: data.Value * 2}
		output <- processedData
	}
	close(output) // Закрываем выходной канал, когда обработка завершена
}

// Функция для второй стадии обработки
func Stage2(input <-chan Stage2Data, output chan<- Stage2Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range input {
		// Дополнительная обработка данных
		processedData := Stage2Data{Value: data.Value + 1}
		output <- processedData
	}
	close(output) // Закрываем выходной канал, когда обработка завершена
}

func main() {
	// Создание каналов для передачи данных между стадиями
	stage1Output := make(chan Stage2Data)
	stage2Output := make(chan Stage2Data)

	// Инициализация WaitGroup
	var wg sync.WaitGroup

	// Запуск горутин для каждой стадии конвейера
	wg.Add(2)
	go Stage1(createStage1Data(), stage1Output, &wg)
	go Stage2(stage1Output, stage2Output, &wg)

	// Чтение окончательных результатов из последнего канала
	for result := range stage2Output {
		fmt.Printf("Результат: %d\n", result.Value)
	}

	// Ожидание завершения всех стадий
	wg.Wait()
}

// Функция для генерации и отправки начальных данных в первую стадию
func createStage1Data() chan Stage1Data {
	input := make(chan Stage1Data)

	go func() {
		defer close(input)
		for i := 1; i <= 5; i++ {
			input <- Stage1Data{Value: i}
		}
	}()

	return input
}
