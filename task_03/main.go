package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.
func predictableFunc(ctx context.Context) (int64, error) {
	resCh := make(chan int64)
	defer close(resCh)

	timeStart := time.Now()
	defer func() {
		timeEnd := time.Now()
		diff := timeEnd.Sub(timeStart)

		fmt.Printf("Duration: %d \n", diff.Milliseconds())
	}()

	go func() {
		resCh <- unpredictableFunc()
	}()

	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("error timeot")

	case res := <-resCh:
		return res, nil

	}
}

func main() {
	fmt.Println("started")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := predictableFunc(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
