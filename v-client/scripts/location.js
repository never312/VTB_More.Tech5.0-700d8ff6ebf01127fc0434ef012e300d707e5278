ymaps.geolocation.get({
    // Выставляем опцию для определения положения по ip    
    provider: 'yandex',
    // Карта автоматически отцентрируется по положению пользователя.
    mapStateAutoApply: true
}).then(function (result) {
    var geoObjects = result.geoObjects;
    if (geoObjects.getLength() > 0) {
        var coordinates = geoObjects.get(0).geometry.getCoordinates();
        console.log("Координаты пользователя: " + coordinates);
    } else {
        console.log("Местоположение пользователя не найдено.");
    }
});
