## Инструкции по установке библиотеки

Следуйте этим шагам, чтобы установить библиотеку:
	

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

	osrmHost := "OSRM HOST"
	profile := "driving|cycling|walking"
    etaService := services.NewEtaService(points, osrmHost, profile, 5, 1.5)
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
