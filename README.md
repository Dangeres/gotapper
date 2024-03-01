<h1 align="center">
  gotapper
</h1>


<h4 align="center">Оглавление README:</h4>
<div align="center">
    <a href="#сторитейлинг"> • Сторитейлинг проекта • </a><br>
    <a href="#про-скрипт"> • Для чего нужен этот скрипт • </a><br>
    <a href="#установка"> • Установка • </a><br>
    <a href="#примечание"> • Примечание • </a>
</div>

## Сторитейлинг
<i>
Жил был Дмитрий, он был разработчиком до мозга костей.<br>
И вот однажды Дмитрий пришел в ресторан, заказал позиции для восполнения жизненной энергии, в ожидании приготовления блюд стал рассматривать окружение - приглушенный свет, вечернее время суток, красивые фонари, подсвечивающие элементы дизайна...<br>
Взгляд перешел на более близкие предметы - стол и предметы, внезапно взгляд зацепился за плашку на столе с QR кодом, отсканировав код попадаем на сайт tapper, который имеет функции - создание и оплата заказа, позвать официанта.<br>
Насытившись едой, Дмитрий был приятно удивлен от вкусных блюд и красивой атмосферы ресторана, было решено прийти в этот ресторан еще раз.<br>
Так как Дмитрий по сути своей любопытный, он заказал другие позиции, поэтому второй поход ему не особо понравился.<br>
И тут Дмитрий, будучи программистом, задался вопросом - как заказать позиции, которые будут вкусными, но так что бы это было интересно и в игровой форме...<br>
Решение пришло сразу же - у нас имеется на каждом столике табличка с оплатой заказа, можно вытащить из этой таблички данные и сделать подсчет позиций.<br>
Так как Дмитрий интересовался всяким, у него в окружении появился новый друг с именем гоша, с помощью которого он смог получить интересные позиции путем перекладывания json'ов.
</i>


## Про скрипт
По сути своей <a href="https://tapper.cloud">tapper</a> представляет собой сервис, который занимается подключением бизнеса(рестораны и кафе) к системе оплаты и предоставляет дополнительные возможности в виде раздельного счета, вызова оффицианта и прочего.

Концепция данного скрипта в том, что зная domain (<i>shop_token внутри программы</i>) заведения (его можно получить, отсканировав любой попавшийся qr код со стола с одноименной подписью), можно пройтись по столам(они вроде как числовые) и получить позиции, которые были заказаны этим столиком. Потом положить это в файлы и в дальнейшем сделать аналитику нужных данных <i>(популярные столики, популярные блюда, средний чек и прочее).</i>

## Установка
1. Скачиваете ветку;
2. Устанавливаете golang;
3. Скачиваем сторонние зависимости `go get main`;
4. Запускаете `main.go` и наслаждаетесь результатом.


## Примечание
README будет изменять и дополняться, запросы собранные публиковать не буду так там какая-то странная шляпа с email официантами имеется, не очень хочется распространять данную информацию.