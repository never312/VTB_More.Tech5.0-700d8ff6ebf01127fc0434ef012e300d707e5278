// Предположим, что у вас есть координаты пользователя и координаты офисов
// Координаты пользователя (замените на реальные данные)
var userCoordinates = [userLatitude, userLongitude];

// Получение координат офисов
var officeCoordinates = getOfficeCoordinates(); // Здесь предполагается, что getOfficeCoordinates возвращает массив координат офисов

// Функция для вычисления расстояния между двумя точками
function getDistance(point1, point2) {
    return ymaps.coordSystem.geo.getDistance(point1, point2);
}

// Вычисление расстояния между пользователем и каждым офисом
for (var i = 0; i < officeCoordinates.length; i++) {
    var distance = getDistance(userCoordinates, officeCoordinates[i]);
    console.log("Расстояние между пользователем и офисом " + i + ": " + distance + " метров");
}
