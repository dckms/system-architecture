:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: Simplicity; in Agile
   :name: emacsway-agile-simplicity


===========================
Role of Simplicity in Agile
===========================

.. sectionauthor:: Ivan Zakrevsky

.. contents:: Содержание

Не очень глубоко осведомленные в архитектуре люди почему-то иногда думают, что архитектурное решение всегда несет перфекционизм и overengineering.
Вероятно, они просто не знакомы с ":ref:`The Primary Technical Imperative <emacsway-primary-technical-imperative>`".

На самом деле, overengineering означает решение не соответствующее контексту, т.е. когда :ref:`уровень привнесенной сложности превышает уровень управляемой сложности <emacsway-agile-patterns>`.
С точки зрения ":ref:`The Primary Technical Imperative <emacsway-primary-technical-imperative>`" это значит, что он ухудшает внутреннее качество программы, а не повышает его.

Более того, хороший архитектор стремится исключить само возникновение сложной проблемы, нежели пытаться героически ее решить.

    📝 "Усложнять - просто, упрощать - сложно".

    -- "Закон Мейера"

Вершиной простоты авиационной инженерной мысли был, пожалуй, самолет Ил-62.

См. видео "`Ил-62 - идеальный вариант <https://youtu.be/VyrN9AJm7sk>`__" / @SkyShips.

Его конструкция получилась настолько удачной и простой, что огромный межконтинентальный лайнер управлялся исключительно мускульной силой пилотов посредством безбустерной (т.е. без усилителей) системы управления, что стало возможным благодаря удачно выбранной балансировке.
Для этого, правда, пришлось привнести в конструкцию заднюю штангу - довольно простое решение, которое на корню исключило возникновение довольно сложной проблемы.
Идеальный пример воплощения `KISS-principle <https://people.apache.org/~fhanik/kiss.html>`__.
Более того, этот лайнер не имел даже предкрылков, благодаря оригинальному аэродинамическому решению в виде "зуба" на передней кромки крыла.

Это яркий пример того, как, вместо того, чтобы создавать сложные решения для сложных проблем, можно просто не допускать самого возникновения этих сложных проблем, благодаря простым и удачным проектным решениям.

    📝 "Вобрав в себя лучшие технологические наработки, лайнер сохранил идеологию простоты конструкции и систем, что сделало его оптимальным для магистральных перевозок в СССР и одним из лучших самолетов своего времени."

    -- "`Ил-62 - идеальный вариант <https://youtu.be/VyrN9AJm7sk>`__" / @SkyShips

..

    📝 Дальнейший этап в творческой деятельности Ильюшина – пассажирский трансконтинентальный лайнер Ил-62, вышедший на воздушные линии в 1967 году и его модификация Ил-62М, ставший флагманом Аэрофлота.
    Примечательно, что даже такой очень большой самолет сохранил простоту и легкость управления, присущую всем "илам".
    В этом – одно из проявлений творческого стиля С.В.Ильюшина, стиля, которому свойственно стремление к оптимальному проектированию, упорство в достижении максимальной надежности и безопасности самолета в сочетании с высокой экономичностью или боевой эффективностью.

    Характерной чертой творческой деятельности Ильюшина являлась простота проектных решений.
    В своих воспоминаниях генеральный конструктор, академик А.С.Яковлев особо отмечает эту черту, называя Ильюшина "мастером простых решений".
    Конечно, эта "простота" требовала огромного творческого напряжения и совершенно четкого и ясного представления эксплуатационной жизни проектируемого самолета.

    В каждом самолете, созданном в конструкторском бюро под руководством С.В.Ильюшина, воплощены творческие особенности Генерального конструктора.
    Умение технически просто решать сложные, а порой противоречивые проблемы – это талант, это стиль С.В.Ильюшина, конструктора и ученого, инженера и творца авиационной техники, что позволяло создавать такие машины, которые сыграли значительную роль в развитии Военно-Воздушных Сил СССР и обеспечили выполнение большой доли работы гражданского воздушного транспорта.
    Они заняли достойное место в истории отечественной авиации.

    Успех С.В.Ильюшиным достигался в результате решения технических задач на основе последних достижений науки, путем смелого внедрения нового и благодаря его исключительной дальновидности.

    От легкого планера с полетным весом 100 кг до межконтинентального лайнера с полетным весом 160 т прошло почти 40 лет.
    Под руководством С.В.Ильюшина спроектировано, построено и испытано в полете десятки машин, многие из которых оказались непревзойденными по летным характеристикам, простоте конструкции, технологии и надежности.

    -- "`Биография Ильюшина <https://www.ilyushin.org/about/history/biography/>`__" / ПАО "Ил" ("Ильюшин" - группа компаний ОАК)

Давайте послушем другого известного "мастера простых решений", оружейного конструктора, создавшего наиболее надежный и простой автомат в истории:

    📝 "хочу сказать, что сделать простое иногда во много раз сложнее, чем сложное."

    -- М.Т. Калашников в интервью журналисту газеты "Metro Москва", 2009 год.

И послушаем выдающегося русского художника Илью Ефимовича Репина:

    📝 "Сначала художник рисует плохо и просто. Потом сложно и плохо. Потом сложно и хорошо. И только потом - просто и хорошо."

    -- И.Е. Репин


Единица измерения
=================

Посмотрим, к примеру, мотивацию Mediator pattern:

    📝 "Mediator promotes loose **coupling** by keeping objects from referring to each other explicitly,
    and it lets you vary their interaction independently."

    -- "Design Patterns: Elements of Reusable Object-Oriented Software" by Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides

Конечно, тут важно найти :ref:`баланс между стоимостью Coupling и стоимостью Decoupling <emacsway-kent-beck-constantine's-law>`.
Но ключевой целью принципа "`Low Coupling & High Cohesion <http://wiki.c2.com/?CouplingAndCohesion>`__" является управление сложностью, т.е. упрощение, а не усложнение!
Именно это позволяет :ref:`рассматривать фрагмент кода изолированно в пределах возможностей краткосрочной памяти человека <emacsway-icebreaker-principle>`.

Поэтому размер программного элемента исчисляется количеством его обязанностей, а не количеством символов.
Если кто-то считает иначе, и думает, что меньше сложности означает "меньше кода", тогда попробуйте понять, что делает этот, весьма лаконичный, код:

    .. code-block:: bash
       :name: emacsway-rm-rf

        echo "test... test... test..." | perl -e '$??s:;s:s;;$?::s;;=]=>%-{<-|}<&|{;;y; -/:-@[-{-};`-{/" -;;s;;$_;see'

    P.S.: Не вздумайте запустить! Он выполняет ``rm -rf /*``.

    -- "`Источник 1 <https://ru.stackoverflow.com/questions/1144804/%D0%A7%D1%82%D0%BE-%D0%B4%D0%B5%D0%BB%D0%B0%D0%B5%D1%82-%D0%B4%D0%B0%D0%BD%D0%BD%D1%8B%D0%B9-%D0%BE%D0%B4%D0%BD%D0%BE%D1%81%D1%82%D1%80%D0%BE%D1%87%D0%BD%D0%B8%D0%BA-%D0%BD%D0%B0-perl>`__", "`Источник 2 <https://lurkmore.to/Rm_-rf>`__"

Как красиво сказал Vladik Khononov: "абзац - это единица измерения мыслей, а не количества слов".
Лаконичность кода определяется уровнем его сложности на горизонте его рассмотрения (т.е. на рассматриваемом уровне абстракции), а не количеством символов.


Качественный код всегда прост!
==============================


Eric Evans
----------

    📝 "Software design is a constant battle with complexity."

    -- "Domain-Driven Design: Tackling Complexity in the Heart of Software" by Eric Evans


Edsger W. Dijkstra
------------------

    📝 "Simplicity and elegance are unpopular because they require hard work and discipline to achieve and education to be appreciated."

    -- Edsger W. Dijkstra

..

    📝 "Simplicity is prerequisite for reliability."

    -- Edsger W. Dijkstra

..

    📝 "Simplicity is a great virtue but it requires hard work to achieve it and education to appreciate it.
    And to make matters worse: complexity sells better."

    -- Edsger W. Dijkstra, 1984 `On the nature of Computing Science <http://www.cs.utexas.edu/users/EWD/transcriptions/EWD08xx/EWD896.html>`__ (EWD896)

..

    📝 "Хороший специалист всегда осознает строго ограниченные размеры своего черепа, поэтому подходит к задачам с максимальной скромностью.

    The competent programmer is fully aware of the strictly limited size of his own skull;
    therefore, he approaches the programming task in full humility"

    -- Edsger W. Dijkstra, 1972


Steve McConnell
---------------

    📝 "Главным Техническим Императивом Разработки ПО является управление сложностью.
    Управлять сложностью будет гораздо легче, если при проектировании вы будете стремиться к простоте.

    Есть два общих способа достижения простоты:
    минимизация объема существенной сложности, с которой приходится иметь дело в любой конкретный момент времени,
    и подавление необязательного роста несущественной сложности.

    Software's Primary Technical Imperative is managing complexity.
    This is greatly aided by a design focus on simplicity.

    Simplicity is achieved in two general ways:
    minimizing the amount of essential complexity that anyone's brain has to deal with at any one time,
    and keeping accidental complexity from proliferating needlessly."

    -- "Code Complete" 2nd edition by Steve McConnell, перевод: Издательско-торговый дом "Русская Редакция"


Kent Beck
---------

    📝 "On the surface, being an XP programmer looks a lot like being a programmer within other software development disciplines.
    You spend your time working with programs, making them bigger, simpler, faster.
    Beneath the surface, however, the focus is quite different.
    Your job isn't over when the computer understands what to do.
    Your first value is communication with other people.
    If the program runs, but there is some vital component of communication left to be done, you aren't done.
    You write tests that demonstrate some vital aspect of the software.
    You break the program into more smaller pieces, or merge pieces that are too small into larger, more coherent pieces.
    You find a system of names that more accurately reflects your intent.

    This may sound like a high-minded pursuit of perfection.
    It is anything but.
    You try to develop the most valuable software for the customer, but not to develop anything that isn't valuable.
    If you can reduce the size of the problem enough, then you can afford to be careful with the work you do on what remains.
    Then, you are careful by habit."

    -- "Extreme Programming Explained" by Kent Beck

..

    📝 "Of course, you can do a better job if you have more tools in your toolbox than if you have fewer, but it is much more important to have a handful of tools that you know when not to use, than to know everything about everything and risk using too much solution."

    -- "Extreme Programming Explained" by Kent Beck

..

    📝 Mastery - The spirit of xUnit is simplicity.
    Martin Fowler said, "Never in the annals of software engineering was so much owed by so many to so few lines of code."
    Some of the implementations have gotten a little complicated for my taste.
    Rolling your own will give you a tool over which you have a feeling of mastery.

    -- "Test-Driven Development By Example" by Kent Beck

..

    📝 Travel light - You can't expect to carry a lot of baggage and move fast.
    The artifacts we maintain should be:

    - Few
    - Simple
    - Valuable

    The XP team becomes intellectual nomads, always prepared to quickly pack up the tents and follow the herd.
    The herd in this case might be a design that wants to go a different direction than anticipated, or a customer that wants to go a different direction than anticipated, or a team member who leaves, or a technology that suddenly gets hot, or a business climate that shifts.

    Like the nomads, the XP team gets used to traveling light.
    They don't carry much in the way of baggage except what they must have to keep producing value for the customer—tests and code.

    <...>

    Travel light - suggests that the manager doesn't impose a lot of overhead - long all-hands meetings, lengthy status reports.
    Whatever the manager requires of the programmers shouldn't take much time to fulfill.

    <...>

    Travel light - The design strategy should produce no "extra" design.
    There should be enough to suit our current purposes (the need to do quality work), but no more.
    If we embrace change, we will be willing to start simple and continually refine.

    -- "Extreme Programming Explained" by Kent Beck

..

    📝 "It's hard to do simple things.
    It seems crazy, but sometimes it is easier to do something more complicated than to do something simple.
    This is particularly true when you have been successful doing the complicated thing in the past.
    Learning to see the world in the simplest possible terms is a skill and a challenge.
    The challenge is that you may have to change your value system.
    Instead of being impressed when someone (like you, for instance) gets something complicated to work, you have to learn to be dissatisfied with complexity, not to rest until you can't imagine anything simpler working."

    -- "Extreme Programming Explained" by Kent Beck

..

    📝 "I'm not a great programmer; I'm just a good programmer with great habits."

    -- Kent Beck at "Refactoring: Improving the Design of Existing Code" 1st edition by Martin Fowler, Kent Beck, John Brant, William Opdyke, Don Roberts

..

    📝 "Solution Complexity

    Sometimes systems grow big and complicated, out of proportion to the problem they solve.
    The challenge is to stop making the problem worse.
    It is difficult for a struggling team to keep going when every defect fixed creates three more.
    XP can help.

    One client began by getting the build process under control.
    The team improved the build so instead of taking 24 hours on a dedicated machine with lots of manual intervention, the build took an hour and could run completely automatically on any machine.
    Then, the team instituted stories and a story board so everyone knew who was working on what and how long they were taking.
    After two years of steady improvement the team reduced costs 60%, going from seventy engineers to twenty; reduced the time to fix defects 66%; and reduced the time to release for major and minor point releases by 75%, from ten weeks to two weeks.
    Once the team had stopped digging itself in deeper, it began to climb out by eliminating excess complexity while also fixing defects.

    The XP strategy for dealing with excess complexity is always the same: chip away at the complexity while continuing to deliver.
    Brighten the corner where you are.
    If you are fixing a defect in an area, clean up while you are there.
    One objection is that this "extra" cleanup takes too long.
    The team is likely wasting time on interruptions to fix defects.
    Cleaning up helps reduce the overhead of work.
    Visible planning can make it easier for every one to see where the time is already going so it is easier to accept the estimates necessary to do the job right."

    -- "Extreme Programming Explained" 2nd edition by Kent Beck, "Chapter 15. Scaling XP :: Solution Complexity"


Martin Fowler
-------------

    📝 "A little time spent refactoring can make the code better communicate its purpose. Programming in this mode is all about saying exactly what you mean."

    -- "Refactoring: Improving the Design of Existing Code", Martin Fowler


Robert C. Martin
----------------

    📝 "Professionals avoid getting so vested in an idea that they can't abandon it and turn around.
    They keep an open mind about other ideas so that when they hit a dead end they still have other options."

    -- "The Clean Coder: a code of conduct for professional programmers" by Robert C. Martin

..

    📝 "A good architecture comes from understanding it more as a journey than as a destination, more as an ongoing process of enquiry than as a frozen artifact."

    -- "Clean Architecture: A Craftsman's Guide to Software Structure and Design" by Robert C. Martin


Bjarne Stroustrup
-----------------

    📝 "I like my code to be elegant and efficient.
    The logic should be straightforward to make it hard for bugs to hide,
    the dependencies minimal to ease maintenance, error handling complete according to an articulated strategy,
    and performance close to optimal so as not to tempt people to make the code messy with unprincipled optimizations.
    Clean code does one thing well."

    -- Bjarne Stroustrup, inventor of C++ and author of The C++ Programming Language.
    "Clean Code: A Handbook of Agile Software Craftsmanship" by Robert C. Martin


Другие
------

    📝 "Simplicity--the art of maximizing the amount of work not done--is essential."

    -- "`Principles behind the Agile Manifesto <http://agilemanifesto.org/principles.html>`__"

..

    📝 "The design goal for Eventlet's API is simplicity and readability.
    You should be able to read its code and understand what it's doing.
    Fewer lines of code are preferred over excessively clever implementations."

    -- "`Eventlet's docs <http://eventlet.net/doc/basic_usage.html>`__"

..

    📝 "Будьте скромны, не считайте себя супергением — это ваша первая ошибка.
    Оставаясь скромным, вы в конечном итоге достигнете уровня супергения, и даже если нет, какая разница.
    Ваш код должен быть прост настолько, что вам не нужно быть гением, чтобы работать с ним.

    Be Humble, don't think of yourself as a super genius, this is your first mistake.
    By being humble, you will eventually achieve super genius status =), and even if you don't, who cares!
    your code is stupid simple, so you don't have to be a genius to work with it."

    -- "`KISS principle <https://people.apache.org/~fhanik/kiss.html>`__"

..

    📝 "Когда кто-либо привязывается к одной какой-нибудь, хотя бы и верной, идее, то он, в сущности, попадает в то же положение, в каком находился бы человек, привязавший себя к столбу, для того чтобы не заблудиться.
    То, что может быть желанной истиной на известной ступени духовного роста, может быть помехой к этому росту и заблуждением на другой, более высокой ступени."

    -- Люси Малори


.. seealso::

   - ":ref:`emacsway-icebreaker-principle`"
   - ":ref:`emacsway-agile-software-design`"
   - ":ref:`emacsway-agile-patterns`"
