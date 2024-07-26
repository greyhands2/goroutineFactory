package goroutines

import (
 "time"
)

type WorkFunctionReturn struct {
 Error   error
 Data    interface{}
 Message string
}

type Result struct {
 Error   error
 Message string
 Data    interface{}
 Latency time.Duration
}

func acceptWork[T any](callJob func(workFunction func(T) (result string, message string, error error), data T) WorkFunctionReturn, workFunction func(T) (result string, message string, error error), data T, ch chan<- Result, start time.Time) {

 result := callJob(workFunction, data)
 if result.Error != nil {
  sendError(ch,
   result.Error, result.Message, start)
 }

 sendSuccess(ch, result.Message, result.Data, start)
}

func sendError(ch chan<- Result, err error, message string, start time.Time) {
 timeTaken := time.Since(start).Round(time.Millisecond)
 ch <- Result{Error: err, Message: message, Latency: timeTaken}

}

func sendSuccess(ch chan<- Result, message string, data interface{}, start time.Time) {
 timeTaken := time.Since(start).Round(time.Millisecond)
 ch <- Result{Message: message, Latency: timeTaken, Data: data}

}
