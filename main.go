package main

import "fmt"

type Operation struct {
	ID            int
	OperationType string // income, outcome
	Category      string // "Зарплата", "Продукты", "Транспорт"
	Amount        float64
	CreatedAt     string
	Description   string
}

var (
	operations        []Operation
	operationSequence int
)

func AddOperation(o Operation) {
	operationSequence++
	o.ID = operationSequence
	operations = append(operations, o)
}

func GetOperationFromConsole() Operation {
	var o Operation
	fmt.Println("Введите тип операции(income, outcome): ")
	fmt.Scan(&o.OperationType)

	fmt.Println("Введите категорию операции(Зарплата, Продукты, Транспорт): ")
	fmt.Scan(&o.Category)

	fmt.Println("Введите сумму операции: ")
	fmt.Scan(&o.Amount)

	return o
}

func PrintAllOperations() {
	fmt.Printf("ID | OperationType | Category | Amount | CreatedAt\n")
	for _, o := range operations {
		fmt.Printf("%d | %s | %s | %.2f | %s\n",
			o.ID,
			o.OperationType,
			o.Category,
			o.Amount,
			o.CreatedAt)
	}
}

func GetTotalByOperationType(operationType string) (total float64) {
	for _, o := range operations {
		if o.OperationType == operationType {
			total += o.Amount
		}
	}
	return total
}

func EditOperation(id int) {
	found := false
	for i, o := range operations {
		if o.ID == id {
			fmt.Println("Введите новый тип операции(income, outcome): ")
			fmt.Scan(&o.OperationType)

			fmt.Println("Введите новую категорию операции(Зарплата, Продукты, Транспорт): ")
			fmt.Scan(&o.Category)

			fmt.Println("Введите новую сумму операции: ")
			fmt.Scan(&o.Amount)

			operations[i] = o
			found = true
			fmt.Println("Операция успешно изменена!")
			break
		}
	}
	if !found {
		fmt.Println("Операция с таким ID не найдена.")
	}
}

func DeleteOperation(id int) {
	found := false
	for i, o := range operations {
		if o.ID == id {
			operations = append(operations[:i], operations[i+1:]...)
			found = true
			fmt.Println("Операция успешно удалена!")
			break
		}
	}
	if !found {
		fmt.Println("Операция с таким ID не найдена.")
	}
}

func Run() {
	fmt.Println("Добро пожаловать в Coinkeeper")

	for {
		fmt.Println("Выберите команду:")
		fmt.Println("1. Добавить операцию")
		fmt.Println("2. Получить список моих операций")
		fmt.Println("3. Получить итоги доходов и расходов")
		fmt.Println("4. Редактировать операцию")
		fmt.Println("5. Удалить операцию")
		fmt.Println("0. Завершить работу")
		var cmd int
		fmt.Scan(&cmd)
		switch cmd {
		case 1:
			o := GetOperationFromConsole()
			AddOperation(o)
		case 2:
			PrintAllOperations()
		case 3:
			tIncome := GetTotalByOperationType("income")
			tOutcome := GetTotalByOperationType("outcome")
			fmt.Println("Всего заработано:", tIncome)
			fmt.Println("Всего потрачено:", tOutcome)
			fmt.Println("Остаток:", tIncome-tOutcome)
		case 4:
			var id int
			fmt.Println("Введите ID операции, которую хотите отредактировать:")
			fmt.Scan(&id)
			EditOperation(id)
		case 5:
			var id int
			fmt.Println("Введите ID операции, которую хотите удалить:")
			fmt.Scan(&id)
			DeleteOperation(id)
		case 0:
			fmt.Println("Рады были помочь!")
			return
		default:
			fmt.Println("Вы ввели несуществующую команду.")
		}
	}
}

func main() {
	Run()
}
