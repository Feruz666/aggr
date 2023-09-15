# Сервис сбора аналитики (аудита)

---
## Для запуска:
### 1 - Создать сеть:
`
make network
`
### 2 - Запустить:
`
make run
`

---
## Пример запроса:
```
{
    "data": 
        {
            "headers": "Test",
            "body": "data" 
        }
    
}
```

### Запрос должен включать в себя заголовок ``X-Tantum-Authorization``


---

## Тестирование RPS в apache benchmark:
![benchmark](https://github.com/Feruz666/aggr/blob/main/static/benchmark.png)