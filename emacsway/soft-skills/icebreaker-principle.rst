:canonical-base-url: https://dckms.github.io/system-architecture

.. index:: Icebreaker Principle, Soft Skills
   :name: emacsway-icebreaker-principle

================
Принцип ледокола
================

.. sectionauthor:: Ivan Zakrevsky

В статье ":doc:`./change-making`" говорилось о том, что, изменяя процессы коллектива, важно выбрать такую область воздействия, в которой потребность изменений будет превосходить силу сопротивления.

На этот счет существует интересная аналогия.
Знаете, почему ледокол рубит лед, а обычный корабль нет?

.. contents:: Содержание


Суть принципа ледокола
======================

Ледокол атакует лед в том направлении, в котором способен обеспечить силовое превосходство.
Он атакует его сверху, где сила сопротивления льда наименьшая.
А корабль - с торца, где силовое превосходство остается за льдом, поскольку в горизонтальной плоскости толща льда простирается на сотни километров.

.. figure:: _media/icebreaker-principle/icebreaker-principle.jpg
   :alt: Ледокол атакует лед сверху, где сила его сопротивления наименьшая. Иллюстрация из открытых источников неизвестного автора.
   :align: center
   :width: 90%

   Ледокол атакует лед сверху, где сила его сопротивления наименьшая. Иллюстрация из открытых источников неизвестного автора.

Вот так все просто - иногда достаточно просто изменить направление воздействия, чтобы преодолеть сопротивление.

Лед сильнее ледокола.
Но ледокол способен создать силовое превосходство в нужное время в нужном месте.
Этого достаточно, чтобы шаг за шагом проложить маршрут полностью.

Еще один важный вывод - ледокол колет лед там, где нужно ходить судам.
Т.е. там, где это действительно востребовано остальными участниками зимней навигации.
Проецируя это в профессиональную плоскость - нужно уметь распознавать истинные потребности коллектива.

Принцип создания силового превосходства в нужное время и в нужном месте находит широкое распространение в природе, в технике, в политике, в военном деле, в борьбе, в спорте, в управленческой психологии, в программировании, в торговле на фондовом рынке и т.п.


В программировании
------------------

    "Software design is a constant battle with complexity."

    -- "Domain-Driven Design: Tackling Complexity in the Heart of Software" by Eric Evans

В алгоритмах группы "Divide-and-conquer" действует тот же принцип, который известен по названием "Разделяй и властвуй".
Властвуй - значит обладай превосходством, откуда происходят такие воинские термины как "господствующая высота", "господство в воздухе" и т.д.

Одна из ключевых задач хорошей архитектуры заключается в управлении сложностью, чтобы обеспечить превосходство `краткосрочной памяти <https://ru.wikipedia.org/wiki/%D0%9C%D0%B0%D0%B3%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%BE%D0%B5_%D1%87%D0%B8%D1%81%D0%BB%D0%BE_%D1%81%D0%B5%D0%BC%D1%8C_%D0%BF%D0%BB%D1%8E%D1%81-%D0%BC%D0%B8%D0%BD%D1%83%D1%81_%D0%B4%D0%B2%D0%B0>`__ разработчика над уровнем сложности рассматриваемого изолировано фрагмента кода.

Grady Booch говорил, что архитектура - это многоуровневая система абстракций.
Где назначение абстракций - управление сложностью.

А Len Bass говорил, что

    Architecture Is an Abstraction

    This abstraction is essential to taming the complexity of a system - we simply cannot, and do not want to, deal with all of the complexity all of the time.

    -- "Software Architecture in Practice" 3d edition by Len Bass, Paul Clements, Rick Kazman

И Thomas H. Cormen:

    In software design, separating what operations do from how they do it is known as abstraction.

    -- "Algorithms Unlocked" 3d edition by Thomas H. Cormen

Ну и как же здесь не вспомнить Steve McConnell:

    Managing complexity is the most important technical topic in software development. In my view, it's so important that Software's Primary Technical Imperative has to be managing complexity.

    -- "Code Complete" by Steve McConnell

..

    "Dijkstra pointed out that no one's skull is really big enough to contain a modern
    computer program (Dijkstra 1972), which means that we as software developers
    shouldn't try to cram whole programs into our skulls at once; we should try to organize
    our programs in such a way that we can safely focus on one part of it at a time. The goal
    is to minimize the amount of a program you have to think about at any one time. You
    might think of this as mental juggling—the more mental balls the program requires you
    to keep in the air at once, the more likely you'll drop one of the balls, leading to a design
    or coding error.

    At the software-architecture level, the complexity of a problem is reduced by dividing
    the system into subsystems. Humans have an easier time comprehending several simple
    pieces of information than one complicated piece. The goal of all software-design
    techniques is to break a complicated problem into simple pieces. The more independent
    the subsystems are, the more you make it safe to focus on one bit of complexity at a
    time. Carefully defined objects separate concerns so that you can focus on one thing at a
    time. Packages provide the same benefit at a higher level of aggregation.

    Keeping routines short helps reduce your mental workload. Writing programs in terms
    of the problem domain, rather than in terms of low-level implementation details, and
    working at the highest level of abstraction reduce the load on your brain.

    The bottom line is that programmers who compensate for inherent human limitations
    write code that's easier for themselves and others to understand and that has fewer
    errors."

    -- "Code Complete" by Steve McConnell

..

    Software's Primary Technical Imperative is managing complexity. This is greatly
    aided by a design focus on simplicity.
    Simplicity is achieved in two general ways: minimizing the amount of essential
    complexity that anyone's brain has to deal with at any one time, and keeping
    accidental complexity from proliferating needlessly.

    -- "Code Complete" by Steve McConnell

..

    The number
    "7±2" has been found to be a number of discrete items a person can remember while
    performing other tasks (Miller 1956). If a class contains more than about seven data
    members, consider whether the class should be decomposed into multiple smaller
    classes (Riel 1996).

    -- "Code Complete" by Steve McConnell


В борьбе
--------

Наглядный пример этого принципа - "`висячка <https://youtu.be/svxD8dPGBJw>`__" в Самбо.

.. figure:: _media/icebreaker-principle/hanging.jpg
   :alt: Для одержания победы не нужно быть сильнее противника - достаточно противопоставить свои сильные группы мышц против его слабых групп мышц. Фото из открытых источников неизвестного автора.
   :align: center
   :width: 90%

   Для одержания победы не нужно быть сильнее противника - достаточно противопоставить свои сильные группы мышц против его слабых групп мышц. Фото из открытых источников неизвестного автора.

Противопоставляя свои сильные группы мышц (спины и ног) против слабых групп мышц (бицепс) противника, становится возможным одержать победу даже над превосходящем по силе противником (и именно поэтому "болевые приемы лежа" (с использованием ног) изучаются в силовых ведомствах).

В одном фильме (уже не помню его названия) было красиво сказано:

    Искусство воевать заключается в том, чтобы быть сильным в нужное время в нужном месте.

Но отсюда можно сделать еще один интересный вывод - суть победы в борьбе заключается в умелом использовании потенциальной энергии в поле тяготения Земли.
Тот, кто повален, существенно ограничен в использовании потенциальной энергии.
Это говорит о важности способности видеть действующие силы в окружении, и умело использовать их.


В военном деле
--------------

Как говорил Г.К.Жуков, бой - это сухая математика.
Важно обеспечить перегруппировку сил таким образом, чтобы на заданном участке фронта обеспечить силовое превосходство.
Как вариант, это приводило к взятию группировки противника в "клещи", с последующим ее ослаблением в условиях окружения при отсутствии тылового обеспечения.

При форсированиии водной преграды наступление разворачивается не сразу, а после концентрации сил на плацдарме.


В спорте
--------

Приседая со штангой на плечах, мы держим позвоночник ровно, чтобы вес штанги равномерно распределялся по всей площади контактной поверхности позвонка, минимизируя удельную нагрузку таким образом, чтобы обеспечить превосходство предела прочности позвонка над ней.


В торговле на фондовом рынке
----------------------------

Главный принцип инвестора - это диверсификация, т.е. распределение рисков таким образом, чтобы каждая категория риска не превосходила допустимый предел финансовой устойчивости.
Это обеспечивает психологическое равновесие инвестора.


В технике
---------

Плавучесть судна обеспечивается водонепроницаемыми перегородками, обеспечивающими превосходство гидростатической подъёмной силы над силой тяжести воды на месте пробоины.


В управленческой психологии
---------------------------

Изменяя процессы коллектива важно выбрать такую область воздействия, в которой потребность изменений будет превосходить силу сопротивления, см. ":doc:`./change-making`".


В психологии
------------

Всем известен принцип Дейла Карнеги "Живите в отсеке сегодняшнего дня".
Можно сказать, что основная битва человека за свое счастье - это битва с его собственными мыслями.
См. также ":doc:`./planning-in-psychology`".


В природе
---------

Вода камень точит.
Видели как море режет скалы?
Обязательно посмотрите - вдохновляет.
Стекающие капельки воды прорезают в камне бороздки и углубляют их до тех пор, пока глыба не обрушится.
Капля против скалы!


В быту
------

Разжигая дрова в мангале для шашлыка, мы используем шепки, ветки, бумагу, или горючие жидкости, чтобы энергия пламени спички обладала превосходством над теплоемкостью воспламеняемого материала.
Попытка воспламенить полено спичкой напрямую окажется безуспешной.


.. seealso::

   - :doc:`./change-making`
   - :doc:`./planning-in-psychology`

