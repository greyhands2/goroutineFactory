package goroutines

import (
	"log"
	"time"
)

func WorkGoroutine[T any](workFunction func(T) (result string, message string, error error), data T) {
 resultChannel := make(chan Result)

 go acceptWork[T](callJob[T], workFunction, data, resultChannel, time.Now())

 result := <-resultChannel
 if result.Error != nil {
  log.Println(result.Error, result.Message)
 } else {
  log.Println(result.Message, result.Data)
 }
}

func callJob[T any](workFunction func(T) (result string, message string, error error), data T) WorkFunctionReturn {
 result, message, err := workFunction(data)

 return WorkFunctionReturn{Error: err, Data: result, Message: message}

}
