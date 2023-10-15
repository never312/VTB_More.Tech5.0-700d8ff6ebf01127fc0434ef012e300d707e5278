function init() {
    let map = new ymaps.Map('map', {
        center: [55.753994, 37.622093],
        zoom: 17
    });

    let placemark1 = new ymaps.Placemark([55.753994, 37.622093], {
        balloonContent: 

            <div class="balloon">
                <div class="balloon__address">г. Париж</div>
                <div class="balloon__contacts">
                    <a href="tel:+7999999999">+7999999999</a>
                </div>
            </div>

        
    }, {
        iconLayout: 'default#image',
        iconImageHref: '/VTB_logo_ru.png',
        iconImageSize: [40, 40],
        iconImageOffset: [-19, -44]
    });

  map.controls.remove('searchControl'); // удаляем поиск
  map.controls.remove('trafficControl'); // удаляем контроль трафика
  map.controls.remove('typeSelector'); // удаляем тип
  map.controls.remove('fullscreenControl'); // удаляем кнопку перехода в полноэкранный режим
  map.controls.remove('zoomControl'); // удаляем контрол зуммирования
  map.controls.remove('rulerControl'); // удаляем контрол правил
  // map.behaviors.disable(['scrollZoom']); // отключаем скролл карты (опционально)

    map.geoObjects.add(placemark1);
}

ymaps.ready(init);