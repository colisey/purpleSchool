package output

import "github.com/fatih/color"

// func PrintError(value interface{}) {
func PrintError(value any) {
	// val, ok := value.(int) Если value int, в переменной val будет значение и ок = true
	// intValue, ok := value.(int)
	// if ok {
	// 	color.Red("Код ошибки: %d", intValue)
	// return
	// }
	// stringValue, ok := value.(string)
	// if ok {
	// 	color.Red(stringValue)
	// return
	// }
	// errorValue, ok := value.(error)
	// if ok {
	// 	color.Red(errorValue)
	// return
	// }
	// color.Red("Неизвестный тип ошибки")
	// return

	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвестный тип ошибки")
	}
}

// func sum[T int | int32 | string](a, b T) T {
// 	// switch d := a.(type) { Не работает проверка
// 	switch d := any(a).(type) { // Преобразуем к any хак
// 	case string:
// 		color.Red(d)
// 	}
// 	return a + b
// }

// type List[T any] struct {
// 	elements []T
// }

// func (l List[T]) addElement() {

// }

// func sumInt(a, b int) int {
// 	return a + b
// }
// func sumFloat32(a, b float32) float32 {
// 	return a + b
// }
