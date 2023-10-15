const showOfficesButton = document.getElementById("showOffices");
const showATMsButton = document.getElementById("showATMs");

showOfficesButton.addEventListener("click", () => {
  // Отправьте запрос к API, чтобы получить отделения
  fetch("/salepoint", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    })
    .then((data) => {
      // Обработка данных, полученных от сервера
      console.log(data);
      // Далее, вы можете обновить интерфейс с полученными данными
      updateInterfaceWithData(data);
    })
    .catch((error) => {
      console.error("Error:", error);
    });
});

showATMsButton.addEventListener("click", () => {
  // Отправьте запрос к API, чтобы получить банкоматы
  fetch("/atm_filters", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    })
    .then((data) => {
      // Обработка данных, полученных от сервера
      console.log(data);
      // Далее, вы можете обновить интерфейс с полученными данными
      updateInterfaceWithData(data);
    })
    .catch((error) => {
      console.error("Error:", error);
    });
});

// Функция для обновления интерфейса с данными
function updateInterfaceWithData(data) {
  // Здесь вы можете создать элементы и обновить страницу
}
