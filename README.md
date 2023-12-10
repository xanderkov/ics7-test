# Тестирование

## Лабораторная работа 1

# Запуск всех тестов

```bash
go test ./...
```

# Запуск случайного теста

Любой файл можно (и функцию в нем) можно протестировать в любом порядке.

Например, для теста файла необходимо написать:

```bash
go test hospital/internal/modules/domain/patient/repo
```

Также для запуска тестов необходимо добавить все enviroment переменные из .env файла.

# Документация тестов

Сделана на allure


## Лабораторная работа 3

В prometheus были взяты флаги:

- go_memstats_alloc_bytes --- количество выделенных байт в куче.
- go_memstats_alloc_bytes_total --- количество выделенных всех байт в куче
- go_memstats_heap_objects --- текущее количество объектов в куче
- go_memstats_mallocs_total --- суммарное количество совершённых операций выделения памяти
- go_gc_duration_seconds_sum --- суммарная продолжительность GC пауз
- process_cpu_seconds_total --- суммарное kernel и user CPU time

### Результаты бенчмарк тестирования

#### AddDoctor
```bash
hospital/tests/benchmark ---runs 1500 CPU: 720523 ns/op	 RAM:   3383 B/op	90 allocs/op
```

#### ReadDoctor
```bash
hospital/tests/benchmark ---runs 1500 CPU: 2156856 ns/op  RAM:   8716 B/op	219 allocs/op
```
