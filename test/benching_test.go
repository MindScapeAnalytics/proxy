package benching

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	cl "github.com/MindScapeAnalytics/proxy/app"
	"github.com/MindScapeAnalytics/proxy/config"
)

func BenchmarkGetEventList(b *testing.B) {
	// Создание нового экземпляра приложения Fiber
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx, _ := context.WithCancel(context.Background())
	app, err := cl.NewApp(ctx, &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		err = app.Drivers.HTTPClient.Launch()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	// Создание фэйкового HTTP запроса
	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:6000/api/v1/testing/assessment/account", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiJhZDNhYjdlYi1kZTVjLTRhNDMtOTQ5My01NmZjMDhlYjM5N2UiLCJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTU4NTIxMTJ9.i2_3-kYrGbVeLdxg6gebu25jMS2sQvtTryPoz8oJDf0")
	// Переменная для записи результатов выполнения
	var res *http.Response

	// Запуск бенчмарка
	b.ResetTimer()
	for i := 0; i < 100; i++ {
		// Создание фэйкового HTTP-запроса
		resp, _ := app.Fiber.Test(req)
		res = resp
	}
	// Завершение теста
	b.StopTimer()

	// Используйте res по своему усмотрению, например, чтобы убедиться в корректности ответа
	_ = res
}

func BenchmarkGetEventInfo(b *testing.B) {
	// Создание нового экземпляра приложения Fiber
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx, _ := context.WithCancel(context.Background())
	app, err := cl.NewApp(ctx, &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		err = app.Drivers.HTTPClient.Launch()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	// Создание фэйкового HTTP запроса
	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:6000/api/v1/core/events/854955d0-3c5f-4183-950f-f960f8676c80", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiJhZDNhYjdlYi1kZTVjLTRhNDMtOTQ5My01NmZjMDhlYjM5N2UiLCJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTU4NTIxMTJ9.i2_3-kYrGbVeLdxg6gebu25jMS2sQvtTryPoz8oJDf0")
	// Переменная для записи результатов выполнения
	var res *http.Response

	// Запуск бенчмарка
	b.ResetTimer()
	for i := 0; i < 100; i++ {
		// Создание фэйкового HTTP-запроса
		resp, _ := app.Fiber.Test(req)
		res = resp
	}
	// Завершение теста
	b.StopTimer()

	// Используйте res по своему усмотрению, например, чтобы убедиться в корректности ответа
	_ = res
}

func BenchmarkAddEvent(b *testing.B) {
	// Создание нового экземпляра приложения Fiber
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx, _ := context.WithCancel(context.Background())
	app, err := cl.NewApp(ctx, &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		err = app.Drivers.HTTPClient.Launch()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	// Создание фэйкового HTTP запроса
	req := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:6000/api/v1/core/events", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiJhZDNhYjdlYi1kZTVjLTRhNDMtOTQ5My01NmZjMDhlYjM5N2UiLCJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTU4NTIxMTJ9.i2_3-kYrGbVeLdxg6gebu25jMS2sQvtTryPoz8oJDf0")
	// Переменная для записи результатов выполнения
	var res *http.Response

	// Запуск бенчмарка
	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		// Создание фэйкового HTTP-запроса
		resp, _ := app.Fiber.Test(req)
		res = resp
	}
	// Завершение теста
	b.StopTimer()

	// Используйте res по своему усмотрению, например, чтобы убедиться в корректности ответа
	_ = res
}
