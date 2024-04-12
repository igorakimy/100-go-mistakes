package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// detach пользовательская структура, действующая
// как обертка исходного контекста.
type detach struct {
	ctx context.Context
}

func (d detach) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (d detach) Done() <-chan struct{} {
	return nil
}

func (d detach) Err() error {
	return nil
}

func (d detach) Value(key any) any {
	// Делегируем вызов на получение значения
	// родительскому контексту.
	return d.ctx.Value(key)
}

func handler(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		// Использование detach поверх контекста HTTP.
		// Теперь переданный для публикации контекст никогда
		// не окажется устаревшим или отмененным - он будет
		// нести в себе все значения из родительского контекста.
		_ = publish(detach{r.Context()}, response)
	}()

	writeResponse(response)
}

func doSomeTask(ctx context.Context, r *http.Request) (http.ResponseWriter, error) {
	return nil, nil
}

func publish(ctx context.Context, w http.ResponseWriter) error {
	return nil
}

func writeResponse(w http.ResponseWriter) {
	_, _ = fmt.Fprint(w, "Some response")
}

func main() {

}
