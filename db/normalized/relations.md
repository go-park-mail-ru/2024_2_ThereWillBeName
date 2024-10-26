## user 

Таблица Пользователей и его персональные данные ( логин, почта, пароль)

Связанные таблицы: trip, review

```
user:
{id} -> {login, email, password}
{login} -> {email, password}
{email} -> {login, password}

```


## trip 
Таблица поездок и их данные( название, описание, город, дата начала и конца и тд)

Связанные таблицы: trips_place

```
trip:
{id}-> {name, description, city, start_date, end_date, private, user}

```


## city
Tаблица городов с текстовым названием города

Связанные таблицы: trip, place

```
city:
{id}->{name}

```


## place 
Таблица достопримечательностей и их данные (название, описание, рейтинг, количество отзывов, город, адрес)

```
place:
{id}->{name, description, rating, numberOfReviews, address, city, phonenumber}
{name, city_id}->{name, description, rating, numberOfReviews, address, phonenumber}

```


## trips_place
Таблица-связка между достопримечательностями (place) и поездками (trip)

```
trips_place:
{id}->{trip_id, place_id}

```


## category
Таблица категорий с текстовым названием категории
Связанные таблицы: places_category

```
category:
{id}->{ name}

```


## places_category 
Таблица-связка между категориями (category) и жостопримечательностями (places)

```
places_category:
{id}->{place_id, category_id}

```


## review
Таблица отзывов

```
review:
{id}->{user, place, rating, reviewText}

```



### Cсылка на ER-диаграммы:
https://www.plantuml.com/plantuml/uml/ZLF1QiCm3BtdAt8Uv0zAQIdP21jhIRgDdOmMQwaXSOgjj2jj_tsTiKUgTI4v14lsavxVasUMm53Nr15gKdI8NueqZuzHVFdkOYNz0XjGjL_NRQMqNs_1sdrhmh7I811A0HITer1naQtVUSKayR661eV0gwVv8Xs3LWrKk0BQ-5YYBOtQace3LmEaT1MGNVE1PlcPuqxXtleiI6dGXgYy4CXaF9dSFqwduSARLkEp0_TOvkbhbhxzIxEpQ8JYbMeeaMKPIUF82S8l6XHulauVblcmo5pJGZ0O1ztcLm9XrPf3R-CS_rPUvzbgX5dCdA3rdsPkDeK42ZZKiXzkZTgvY6jpK_Rp5Sr238yMNVYDyeoRuXxGwPGafwEVYyYz09zoNIFtPwU27o7DRBh2yCr-LtUzNa-V-Fdr4uLWLFZ63BNicleR
![Alt text](image.png)
