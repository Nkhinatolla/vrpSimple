## Инструкции по установке библиотеки

Следуйте этим шагам, чтобы установить библиотеку:

1. Внесите изменения в файл `Dockerfile` и добавьте следующую строку для установки переменной среды `GOPRIVATE`:

   ```
   ENV GOPRIVATE=git.chocofood.kz
   ```

2. Создайте токен доступа в GitLab. Для этого перейдите в "Настройки" и выберите "Токены доступа". Сохраните токен доступа в безопасном месте.

3. Откройте файл `~/.netrc` и добавьте следующую строку, заменив `YOUR_LOGIN` вашим логином в GitLab и `ACCESS_TOKEN` на токен доступа, сгенерированный на предыдущем шаге:

   ```
   machine git.chocofood.kz login YOUR_LOGIN password ACCESS_TOKEN
   ```

4. В файле `go.mod` добавьте следующую строку, чтобы добавить библиотеку в зависимости:

   ```
   require git.chocofood.kz/chocodelivery/assignments/vrp-simple v1.0.4
   ```

## Использование

В этом разделе приведите примеры или инструкции по использованию библиотеки.

```go
    point_1 := domain.EtaPoint{
        ID:                "courier_1",
        Dependencies:      []string{},
        ShouldBeArrivedAt: time.Now().Add(3 * time.Hour),
        Latitude:          43.189297,
        Longitude:         76.871927,
    }
    point_2 := domain.EtaPoint{
        ID:                "point_1",
        Dependencies:      []string{"courier_1"},
        ShouldBeArrivedAt: time.Now().Add(3 * time.Hour),
        Latitude:          43.269373,
        Longitude:         76.936449,
	}
    point_3 := domain.EtaPoint{
        ID:                "point_2",
        Dependencies:      []string{"courier_1"},
        ShouldBeArrivedAt: time.Now().Add(3 * time.Hour),
        Latitude:          43.199297,
        Longitude:         76.871927,
    }

    points := make([]domain.EtaPoint, 3)
    points[0] = point_1
    points[1] = point_2
    points[2] = point_3

    etaService := services.NewEtaService(points, "https://osrm02.chocodelivery.kz", "driving", 5, 1.5)
    points, err := etaService.FindOptimalEta(false)

    if err != nil {
        fmt.Println(err)
    }

    if err != nil {
        fmt.Println(err)
    }
    for _, point := range points {
        fmt.Println(point.Priority, point.ID, point.EstimateAt)
    }
```

Дополнительную информацию и подробную документацию можно найти на веб-сайте библиотеки или в исходном коде.

Пожалуйста, настройте инструкции согласно вашей конкретной установке и использованию библиотеки.