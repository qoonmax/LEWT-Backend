# LEWT Backend macOS

![](https://github.com/qoonmax/LEWT-Client/blob/main/LEWT%20Client/icons/LEWT_icon_128.png)

Приложение которое подменяет содержимое TouchBar для MacBook.

Приложение на GO, отслеживает нажатия с клавиатуры, если сформированный текст удовлетворяет требованиям, он будет переведён через API Яндекс Переводчик.

Получившийся текст на английском языке будет доступен по localhost:3333.
Клиентская часть приложения получает по http с бэкенда строку с английским языком и выводит в TouchBar.

TODO:
- [ ] Избавиться от бэкенда полностью, перенести всю логику на клиент. Или заменить polling на WS.

### Превью

<img src="https://github.com/qoonmax/LEWT-Backend/blob/main/show_2.gif" width="auto">

<img src="https://github.com/qoonmax/LEWT-Backend/blob/main/show.gif" width="auto">
