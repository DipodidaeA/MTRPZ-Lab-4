# MTRPZ-Lab-4

1 Запуст програми відбувається виконавши команду go run .\main.go

2 Перейти за посиланням http://localhost:17000/

3 Виконувати http запити http://localhost:17000/?cmd=ввід команд

  Доступні команди: 
  
  Fwhite - змінює фон на білий
  
  Fgreen - змінює фон на зелений
  
  Fred - змінює фон на червоний
  
  Fblack - змінює фон на чорний
  
  Fblue -  змінює фон на синій

  Fyellow - змінює фон на жовтий
  
  **координати вказувати дробовими числами від 0 до 1 (приклад 0; 0.1; 0.5; 0.7; 1)**
  
  Rblack x1 y1 x2 y2 - малює чорний прямокутник за координатами
  
  Rred x1 y1 x2 y2 - малює червоний прямокутник за координатами
  
  Rgreen x1 y1 x2 y2 - малює зелений прямокутник за координатами
  
  Rblue x1 y1 x2 y2 - малює синій прямокутник за координатами
  
  Ryellow x1 y1 x2 y2 - малює жовтий прямокутник за координатами
  
  reset - стирає картинку та малює чорний екран
  
  update - оновлює картинку, **потрібно завжди ставити останьою командою для оновлення екрану**
  
  %0A - роздільник між новими командами
  
  Приклад команд:
  
  http://localhost:17000/?cmd=Fwhite%0ARred 0.1 0.1 0.7 0.2%0Aupdate - малює на білому фоні червоний прямокутник та оновлює картинку
  
  http://localhost:17000/?cmd=reset%0Aupdate - стирає  картинку перемальовує екран у чорний та оновлює вікно
