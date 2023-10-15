// Предположим, что у вас есть массив координат офисов
var officeCoordinates = getOfficeCoordinates(); // Здесь предполагается, что getOfficeCoordinates возвращает массив координат офисов

// Инициализируем карту
ymaps.ready(function () {
    var myMap = new ymaps.Map('map', {
        center: [55.753994, 37.622093],
        zoom: 9,
    });

    // Создаем метки для каждой пары координат
    for (var i = 0; i < officeCoordinates.length; i++) {
        var coordinate = officeCoordinates[i];
        var officePlacemark = new ymaps.Placemark(coordinate, {}, {
            preset: 'islands#blueIcon' // Стиль метки, можно изменить
        });
        myMap.geoObjects.add(officePlacemark);
    }
});
