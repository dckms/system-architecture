:canonical-base-url: https://dckms.github.io/system-architecture

.. index:: TDD
   :name: emacsway-tdd

===================================
TDD - Разработка через тестирование
===================================

Вокруг TDD (Test-Driven Development) сложилось немало мифов и заблуждений, и здесь я попытаюсь восстановить изначальный его смысл.
Для этого придется обратиться к первоисточникам.

.. contents:: Содержание


Что такое TDD?
==============

Прежде всего, что такое TDD (Test-Driven Development)?

    Чистый код, который работает (clean code that works), — в этой короткой, но содержательной фразе, придуманной Роном Джеффризом (Ron Jeffries), кроется весь смысл методики Test-Driven Development (TDD).
    Чистый код, который работает, — это цель, к которой стоит стремиться, и этому есть причины.

    Clean code that works - now.
    This is the seeming contradiction that lies behind much of the pain of programming.
    Test-driven development replies to this contradiction with a paradox-test the program before you write it.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан


TDD - это о Software Design
---------------------------

Классическое заблуждение заключается в том, что TDD - это методика тестирования.
На самом же деле, TDD - это, прежде всего, методика разработки и проектирования:

    Ирония TDD состоит в том, что это вовсе не методика тестирования.
    Это методика анализа, методика проектирования, фактически методика структурирования всей деятельности, связанной с разработкой программного кода.

    One of the ironies of TDD is that it isn't a testing technique (the Cunningham Koan).
    It's an analysis technique, a design technique, really a technique for structuring all the activities of development.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан

..

    Если сравнивать со средним уровнем индустрии разработки программного обеспечения, методика TDD позволяет вам писать код, содержащий значительно меньше дефектов и формировать значительно более чистый дизайн. Те, кто стремится к изяществу, могут найти в TDD средство для достижения цели.

    It lets you write code with far fewer defects and a much cleaner design than is common in the industry. However, those whose souls are healed by the balm of elegance can find in TDD a way to do well by doing good.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан

..

    TDD базируется на очаровательно-наивном предположении программиста о том, что чем красивее код, тем вероятнее успех.
    TDD помогает вам обращать внимание на правильные вопросы в подходящие для этого моменты времени. Благодаря этому вы можете делать дизайн чище и модифицировать его по мере того, как перед вами встают новые обстоятельства.

    TDD rests on a charmingly naive geekoid assumption that if you write better code, you'll be more successful.
    TDD helps you to pay attention to the right issues at the right time so you can make your designs cleaner, you can refine your designs as you learn.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан

..

    Мы с Робертом Мартином (Robert Martin) занимались исследованием подобного стиля TDD.
    Проблема состоит в том, что дизайн продолжает вас удивлять.
    Идеи, которые на первый взгляд кажутся вам вполне уместными, позже оказываются неправильными.
    Поэтому я не рекомендую целиком и полностью доверять своим предчувствиям относительно паттернов.
    Лучше думайте о том, что, по-вашему, должна делать система, позвольте дизайну оформиться так, как это необходимо.

    Robert Martin and I did some research into this style of TDD. The problem is that the design keeps surprising you.
    Perfectly sensible design ideas turn out to be wrong.
    Better just to think about what you want the system to do, and let the design sort itself out later.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан


TDD - это способ управления сложностью
--------------------------------------

Согласно закономерности `Магического числа семь плюс-минус два <https://en.wikipedia.org/wiki/The_Magical_Number_Seven,_Plus_or_Minus_Two>`__, обнаруженной американским учёным-психологом Джорджем Миллером, кратковременная человеческая память, как правило, не может запомнить и повторить более 7 ± 2 элементов.
Превышение этого порога замедляет темпы разработки, так как мозг не может преодолеть высокую концентрацию сложности.

Тут можно провести аналогию с народной пословицей: веник сложно поломать пока он связан, но, развязав его на отдельные прутики, их можно легко переломать по отдельности.
TDD именно именно это и делает - декомпозирует сложность в процессе разработки.

.. index::
   single: Refactoring; definition
   single: Refactoring; in TDD

Здесь хорошо прослеживается аналогия с рефакторингом, который, в значительной мере, был основан тем же самым человеком - Кент Беком.

    Мой первый опыт проведения дисциплинированного "поэтапного" рефакторинга связан с программированием на пару с Кентом Беком (Kent Beck) на высоте 30 000 футов.

    My first experience with disciplined, "one step at a time" refactoring was when I was pair-programming at 30,000 feet with Kent Beck.

    -- Martin Fowler, the key author of "Refactoring: Improving the Design of Existing Code" [#fnrefactoring]_, перевод С. Маккавеева

К тому же, рефакторинг является необъемлемой частью цикла TDD:

    Красный—зеленый—рефакторинг — это мантра TDD.

    Red/green/refactor - the TDD mantra.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан

По основной версии, слово "refactoring" происходит от математического термина "factoring", и дословно переводится как "факторизация" или "декомпозиция", о чем говорит на своем сайте ключевой автор известной книги "Refactoring: Improving the Design of Existing Code" [#fnrefactoring]_ (благодаря которой, рефакторинг, собственно, и стал популярным):

    The obvious answer comes from the notion of factoring in mathematics. You can take an expressions such as x^2 + 5x + 6 and factor it into (x+2)(x+3). By factoring it you can make a number of mathematical operations much easier. Obviously this is much the same as representing 18 as 2*3^2. I've certainly often heard of people talking about a program as well factored once it's broken out into similarly logical chunks.

    -- "`Etymology Of Refactoring <https://martinfowler.com/bliki/EtymologyOfRefactoring.html>`__" by Martin Fowler

Такое же мнение можно увидеть и на сайте Ward Cunningham:

    Refactoring is a kind of reorganization. **Technically, it comes from mathematics when you factor an expression into an equivalence - the factors are cleaner ways of expressing the same statement.** Refactoring implies equivalence; the beginning and end products must be functionally identical. You can view refactoring as a special case of reworking (see WhatIsReworking).

    Practically, refactoring means making code clearer and cleaner and simpler and elegant. Or, in other words, clean up after yourself when you code. Examples would run the range from renaming a variable to introducing a method into a third-party class that you don't have source for.

    **Refactoring is not rewriting, although many people think they are the same.** There are many good reasons to distinguish them, such as regression test requirements and knowledge of system functionality. The technical difference between the two is that refactoring, as stated above, doesn't change the functionality (or information content) of the system whereas rewriting does. Rewriting is reworking. See WhatIsReworking.

    Refactoring is a good thing because complex expressions are typically built from simpler, more grokable components. Refactoring either exposes those simpler components or reduces them to the more efficient complex expression (depending on which way you are going).

    For an example of efficiency, count the terms and operators: (x - 1) * (x + 1) = x^2 - 1. Four terms versus three. Three operators versus two. However, the left hand side expression is (arguably) simpler to understand because it uses simpler operations. Also, it provides you more information about the structure of the function f(x) = x^2 - 1, like the roots are +/- 1, that would be difficult to determine just by "looking" at the right hand side.

    -- "`What Is Refactoring <http://wiki.c2.com/?WhatIsRefactoring>`__" on wiki.c2.com

Если кому-то имя Ward Cunningham ни о чем не говорит, то вот как представил его сам Kent Beck в книге "Test-Driven Development By Example" [#fntdd]_:

    Я начал свою жизнь настоящего программиста благодаря наставничеству и в рамках постоянного сотрудничества с Уордом Каннингэмом (Ward Cunningham).
    Иногда я рассматриваю разработку, основанную на тестах, как попытку предоставить каждому программисту, работающему в произвольной среде, ощущение комфорта и тесной дружбы, которое было у нас с Уордом, когда мы вместе разрабатывали программы Smalltalk в среде Smalltalk.
    He существует способа определить первоначальный источник идей, если два человека обладают одним общим мозгом.
    Если вы предположите, что все хорошие идеи на самом деле изначально придумал Уорд, вы не будете далеки от истины.

    My life as a real programmer started with patient mentoring from and continuing collaboration
    with Ward Cunningham. Sometimes I see Test-Driven Development (TDD) as an attempt to
    give any software engineer, working in any environment, the sense of comfort and intimacy
    we had with our Smalltalk environment and our Smalltalk programs. There is no way to sort
    out the source of ideas once two people have shared a brain. If you assume that all of the
    good ideas here are Ward's, then you won't be far wrong.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан

Ну и Википедия о факторизации:

    Factorization (or factoring) may also refer to more general decompositions of a mathematical object into the product of smaller or simpler objects.
    For example, every function may be factored into the composition of a surjective function with an injective function.

    -- "`Factorization <https://en.wikipedia.org/wiki/Factorization>`__", Wikipedia

..

    Decomposition in computer science, also known as factoring, is breaking a complex problem or system into parts that are easier to conceive, understand, program, and maintain.

    -- "`Decomposition <https://en.wikipedia.org/wiki/Decomposition_(computer_science)>`__", Wikipedia

..

    В математике факториза́ция или фа́кторинг — это декомпозиция объекта (например, числа, полинома или матрицы) в произведение других объектов или факторов, которые, будучи перемноженными, дают исходный объект.
    Например, число 15 факторизуется на простые числа 3 и 5, а полином x2 − 4 факторизуется на (x − 2)(x + 2).
    В результате факторизации во всех случаях получается произведение более простых объектов, чем исходный.

    -- "`Факторизация <https://ru.wikipedia.org/wiki/%D0%A4%D0%B0%D0%BA%D1%82%D0%BE%D1%80%D0%B8%D0%B7%D0%B0%D1%86%D0%B8%D1%8F>`__", Wikipedia


Таким образом, рефакторинг - это способ управления сложностью программы, который делает программу более читаемой и понимаемой за счет декомпозиции сложности, что позволяет снизить нагрузку на человеческую память.
Процесс рефакторинга подобен факторизации математического выражения, в результате которого выводится более простое эквивалентное выражение, т.е. сохраняется функциональная идентичность.
Именно поэтому рефакторинг оставляет неизменным внешнее поведение системы:

    Рефакторинг представляет собой процесс такого изменения программной системы, при котором не меняется внешнее поведение кода, но улучшается его внутренняя структура.

    Refactoring is the process of changing a software system in such a way that it does not alter the external behavior of the code yet improves its internal structure.

    -- Martin Fowler in "Refactoring: Improving the Design of Existing Code" [#fnrefactoring]_, перевод С. Маккавеева

TDD, как и рефакторинг, расщепляет сложность таким образом, чтобы минимизировать объем сложности, рассматриваемый разработчиком в единицу времени.
Это как песочные часы - одна песчинка в единицу времени.
Именно этим объясняется повышение темпов разработки при использовании TDD.

    Каким образом можно модифицировать одну часть метода или объекта, состоящего из нескольких частей?
    Вначале изолируйте изменяемую часть.
    Мне приходит в голову аналогия с хирургической операцией: фактически все тело оперируемого пациента покрыто специальной простыней за исключением места, в котором, собственно, осуществляется операция.
    **Благодаря такому покрытию хирург имеет дело с фиксированным набором переменных.**
    Перед выполнением операции врачи сколь угодно долго могут обсуждать, какое влияние на здоровье пациента оказывает тот или иной орган, однако во время операции внимание хирурга должно быть сфокусировано.

    How do you change one part of a multi-part method or object? First, isolate the part that has to change.
    The picture that comes to my mind is surgery: The entire patient except the part to be operated on is draped.
    **The draping leaves the surgeon with only a fixed set of variables.**
    Now, we could have long arguments over whether this abstraction of a person to a lower left quadrant abdomen leads to good health care, but at the moment of surgery, I'm kind of glad the surgeon can focus.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан

..

    Несмотря на множество появившихся в последнее время мощных инструментов, программирование по-прежнему остается сложной работой.
    Я часто ощущаю себя в ситуации, когда мне кажется, что я жонглирую шариками, и мне приходится следить за несколькими шариками в воздухе в одно и то же время: малейшая потеря внимания, и все сыпется на пол.
    Методика TDD позволяет избавиться от этого ощущения.

    **Когда вы работаете в стиле TDD, в воздухе постоянно находится лишь один шарик.**
    **Вы можете сконцентрироваться на нем, а значит, хорошо справиться со своей работой.**
    Когда я добавляю в программу новую функциональность, я не думаю о том, какой дизайн должен быть реализован в данной функции.
    Я просто пытаюсь добиться срабатывания тестов самым простым из доступных мне способов.
    Когда я переключаюсь в режим рефакторинга, я не беспокоюсь о добавлении в программу новых функций, я думаю только о правильном дизайне.
    На каждом из этих этапов я концентрируюсь на единственной задаче, благодаря этому мое внимание не распыляется.

    Despite all the fancy tools that we have, programming is still hard.
    I can remember many programming times when I feel like I was trying to keep several balls in the air at once, any lapse of concentration and everything would come tumbling down.
    Test-driven development helps reduce that feeling, and as a result you get this rapid unhurriedness.

    **I think the reason for this is that working in a test-driven development style gives you this sense of keeping just one ball in the air at once, so you can concentrate on that ball properly and do a really good job with it.**
    When I'm trying to add some new functionality, I'm not worried about what really makes a good design for this piece of function, I'm just trying to get a test to pass as easily as I can.
    When I switch to refactoring mode, I'm not worried about adding some new function, I'm just worried about getting the right design.
    With both of these I'm just focused on one thing at a time, and as a result I can concentrate better on that one
    thing.

    -- Martin Fowler, Afterword, "Test-Driven Development By Example" [#fntdd]_, перевод П. Анджан

..

    Снижение количества дефектов приводит к возникновению множества вторичных психологических и социальных эффектов.
    После того как я начал работать в стиле TDD, программирование стало для меня значительно менее нервным занятием.
    **Когда я работаю в стиле TDD, мне не надо беспокоиться о множестве вещей.**
    **Вначале я могу заставить paботать только один тест, потом — все остальные.**
    Уровень стресса существенно снизился.
    Взаимоотношения с партнерами по команде стали более позитивными.
    Разработанный мною код перестал быть причиной сбоев, люди стали в большей степени рассчитывать на него.
    У заказчиков тоже повысилось настроение.
    Теперь выпуск очередной версии системы означает новую функциональность, а не набор новых дефектов, которые добавляются к уже существующим.

    Part of the effect certainly comes from reducing defects.
    The sooner you find and fix a defect, the cheaper it is, often dramatically so (just ask the Mars Lander).
    There are plenty of secondary psychological and social effects from reduced defects. My own practice of programming became much less stressful when I started with TDD.
    **No longer did I have to worry about everything at once.**
    **I could make this test run, and then all the rest.**
    Relationships with my teammates became more positive.
    I stopped breaking builds, andpeople could rely on my software to work.
    Customers of my systems became more positive, too.
    A new release of the system just meant more functionality, not a host of new defects to identify among all of their old favorite bugs.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан

Иногда мозгу сложно удержать все в голове, и разработчик берется за листочек и ручку.
При TDD, вместо листочка и ручки используется файловый редактор.
TDD позволяет сфокусировать мозг на минимально возможной единице сложности, которую можно рассмотреть изолированно, что приводит к перераспределению умственных ресурсов.
Кстати, именно это является одной из ключевых особенностей, благодаря которой, практикование TDD делает код чище.

Если рефакторинг помогает сосредоточиться на одной обязанности, выполняемой функцией, то TDD идет еще дальше, и помогает сосредоточиться на одном конкретном значении функции, а значит, - на одном из ее внутренних состояний.
Это позволяет выводить алгоритм функции путем обобщения пересекаемых триангуляцией ее внутренних состояний (и поведений, производящих эти состояния).
А это, в свою очередь, позволяет моделировать поведение функции небольшими законченными фрагментами, удовлетворяющими конкретным значениям функции, и визуализировать формирование поведения функции прямо в редакторе.
Наглядно это демонстрируется на примере :download:`выведения функции Фибоначи <_media/tdd/tdd-fibonacci.txt>` в приложении книги, см. Appendix II. Fibonacci [#fntdd]_.

    Это еще один паттерн рефакторинга: **разработать код, который работает с некоторым конкретным экземпляром, и обобщить этот код так, чтобы он мог работать со всеми остальными экземплярами**, для этого константы заменяются переменными.
    В данном случае роль константы играет не некоторое значение, а жестко фиксированный код (имя конкретного метода).
    Однако принцип остается одним и тем же.
    В рамках TDD эта проблема решается очень легко: **методика TDD снабжает вас конкретными работающими примерами, исходя из которых вы можете выполнить обобщение**.
    Это значительно проще, чем выполнять обобщение исходя только из собственных умозаключений.

    Here is another general pattern of refactoring: **take code that works in one instance and generalize it to work in many** by replacing constants with variables.
    Here the constant was hardwired code, not a data value, but the principle is the same.
    **TDD makes this work well by giving you running concrete examples from which to generalize**, instead of having to generalize purely with reasoning.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан

..

    Контроль над объемом работы.
    Программисты привыкли пытаться предвидеть возникновение в будущем самых разнообразных проблем.
    Если вы начинаете с конкретного примера и затем осуществляете **обобщение кода**, это помогает вам избавиться от излишних опасений.
    Вы можете **сконцентрироваться на решении конкретной проблемы** и поэтому выполнить работу лучше.
    При переходе к следующему тесту вы опять же концентрируетесь на нем, так как знаете, что предыдущий тест гарантированно работает.

    Scope control - Programmers are good at imagining all sorts of future problems.
    Starting with one concrete example and **generalizing** from there prevents you from prematurely confusing yourself with extraneous concerns.
    You can do a better job of solving the immediate problem **because you are focused**.
    When you go to implement the next test case, you can focus on that one, too, knowing that the previous test is guaranteed to work.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан

Математическое объяснение этого явления можно найти в главе "1. Recurrent Problems : 1.1. The Tower of Hanoi" книги "Concrete Mathematics: A Foundation for Computer Science" 2nd edition by Ronald L. Graham, Donald E. Knuth, Oren Patashnik.

Кроме того, при TDD хорошо отслеживается ниточка, за которую можно распутать клубок сложности, и вопрос "с какого конца подступиться" решается сам собой.


.. index::
   single: TDD; in velocity increase
   :name: emacsway-why-is-tdd-faster

Почему TDD быстрее
------------------

При TDD разработка осуществляется быстрее, хотя объема кода пишется больше.
Суть в том, что в процессе конструирования кода, :ref:`91% времени занимает чтение кода и борьба со сложностью, и только 9% времени (1:10) занимает ввод символов с клавиатуры <emacsway-who-reads-the-code>`.

TDD является эффективным средством управления сложностью и снижения когнитивной нагрузки.
А поскольку чтение кода и борьба со сложностью (обдумывание) занимает более 91% времени конструирования кода, то время на написание тестов полностью перекрывается повышением темпов разработки, т.е. разработка с тестами получается даже быстрей.
Пальцы работают больше, а голова меньше.
Происходит перераспределение составляющих разработки.

Допустим, что разработчику нужно написать вдвое больше кода без роста когнитивной нагрузки (написание тестов не требует борьбы со сложностью).
Т.е. вместо соотношения 1:10 (где 1 - это часть времени ввода символов с клавиатуры, а 10 - это часть времени чтения кода и борьбы со сложностью) получится соотношение 2:10, что равно 17%:83% вместо 9%:91%.
Совокупное время увеличится на ``100%*(12 - 11)/11 = 9%`` - ровно столько времени потребуется свеху для того, чтобы написать вдвое больше кода без роста когнитивной нагрузки.

А теперь представим, что удалось снизить когнитивную нагрузку вдвое.
Т.е. вместо соотношения 1:10 (где 1 - это часть времени ввода символов с клавиатуры, а 10 - это часть времени чтения кода и борьбы со сложностью) получится соотношение 1:5, что равно 17%:83% вместо 9%:91%.
Совокупное время уменьшится на ``100%*(6 - 11)/11 = -45%`` - ровно столько времени сэкономится, если разработчик будет тратить вдвое меньше времени на борьбу со сложностью.

9% (вдвое больше кода) против 45% (вдвое меньше думать).

Конечно, коэффициенты в этом примере сильно завышены, но они хорошо раскрывают механизм ускорения темпов разработки с использованием TDD.
На практике TDD дает прирост разработки около 10% - Jason Gorman публиковал свою статистику многократного прохождения кат как по TDD, так и без TDD (см. главу "Chapter 1. What Is Design and Architecture? :: What went wrong?" книги "Clean Architecture: A Craftsman's Guide to Software Structure and Design" [#fncarch]_ by Robert C. Martin).

Я перепроверял эту особенность на личном опыте, и убедился в том, что это, действительно, работает.

Кроме того, время на написание тестов можно прогнозировать, в отличии от отладки.


TDD - основной катализатор Clean Code
-------------------------------------

Каким образом тестирование улучшает качество кода?

    "The problem with testing code is that you have to isolate that code.
    It is often difficult to test a function if that function calls other functions.
    To write that test you've got to figure out some way to decouple the function from all the others.
    In other words, the need to test first forces you to think about good design.

    If you don't write your tests first, there is no force preventing you from coupling the functions together into an untestable mass.
    If you write your tests later, you may be able to test the inputs and the outputs of the total mass, but it will probably be quite difficult to test the individual functions."

    -- "Clean Coder" [#fnccoder]_ by Robert Martin

Однако, нужно учитывать:

    Я сказал, что предположение наивное, однако, скорее всего, я преувеличил.
    На самом деле наивно предполагать, что чистый код — это все, что необходимо для успеха.
    Мне кажется, что хорошее проектирование — это лишь 20% успеха.
    Безусловно, если проектирование будет плохим, вы можете быть на 100% уверены в том, что проект провалится.
    Однако приемлемый дизайн сможет обеспечить успех проекта только в случае, если остальные 80% будут там, где им полагается быть.

    I say "naive," but that's perhaps overstating.
    What's naive is assuming that clean code is all there is to success.
    Good engineering is maybe 20 percent of a project's success.
    Bad engineering will certainly sink projects, but modest engineering can enable project success as long as the other 80 percent lines up right.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан


Влияние TDD на темпы разработки
===============================

Я уже перечислял `превосходства TDD для быстрой разработки <https://emacsway.github.io/ru/it/agile/easily-about-agile-way-to-rapid-development/#self-testing-code-for-agile-ru>`__, поэтому повторяться не буду.

Однако, перечислю основные методики, которые используются для быстрой разработки:

- Emergent Design
- Evolutionary (Incremental, Continuous) Design
- YAGNI
- Очевидная Реализация (Obvious Implementation)
- Копирование Паттернов (Pattern Copying)

Первые два хорошо подходят для начинающих специалистов, поскольку они позволяют эффективно обрабатывать случаи неполной информированности.
Последние два - для опытных специалистов.

Несмотря на то, что Martin Fowler (как редактор статьи Jim Shore) объединяет смысл Emergent Design и Continuous Design:

    Continuous design is also known as evolutionary or emergent design.
    I prefer the term continuous design because it emphasizes the core of the process: continuously taking advantage of opportunities to improve your design.

    -- "`Continuous Design <https://www.martinfowler.com/ieeeSoftware/continuousDesign.pdf>`__" by Jim Shore

Существует точка зрения, что они, все-таки, отличаются:

    We distinguish between emergent and evolutionary architecture, and this distinction is an important one.

    -- "`Microservices as an Evolutionary Architecture <https://www.thoughtworks.com/insights/blog/microservices-evolutionary-architecture>`__" by Neal Ford, Rebecca Parsons


.. index::
   single: Black Box; in TDD
   single: White Box; in TDD


Black Box or White Box?
=======================

Тесты по возможности должны быть черным ящиком, т.е. тестируем поведение, а не реализацию.
Это позволяет безболезненно подменять реализацию при рефакторинге.
Опускаться в глубь реализации нужно тогда, когда это требуется для сокращения комбинаций условий тестирования, например, класс использует несколько подключаемых стратегий, и нам проще протестировать стратегии по одной.
Но при этом мы должны минимизировать зависимость от реализации.
Нарушение этого принципа, в сочетании со стремлением к высокому уровню покрытия кода тестами, накладывает на код оковы и ставит крест на дальнейшей эволюции программы.
Эту тему раскрывает Бек в первой и второй серии сериала "`Is TDD dead? <https://martinfowler.com/articles/is-tdd-dead/>`__".

..

    My personal practice - I mock almost nothing.
    If I can't figure out how to test efficiently with the real stuff, I find another way of creating a feedback loop for myself.
    I have to have feedback loop and the feedback loop has to be repeatable, but like I just don't go very far down the mock path.
    I look at a code where you have mocks returning mocks returning mocks and my experience is if I use TDD I can refactor stuff.
    And then I heard these stories people say well I use TDD and now I can't refactor anything and I feel like I couldn't understand that and I started looking at their tests well.
    If you have mocks returning mocks returning mocks your test is completely coupled to the implementation, not the interface, but the exact implementation of some object you know three streets away.
    Of course you can't change anything without breaking the test.
    So that for me is too high a price to pay.
    That's not a trade-off I'm willing to make just to get piecemeal development.

    -- Kent Beck, "`Is TDD Dead? Part 1 at 21:10 <https://youtu.be/z9quxZsLcfo?t=1269>`__

..

    Думать об объектах, как о черных ящиках, достаточно тяжело.
    Представим, что у нас есть объект Contract, состояние которого содержится в поле status, которое может принадлежать либо классу Offered, либо классу Running.
    В этом случае можно написать тест, исходя из предполагаемой реализации:

    .. code-block:: java
       :name: code-1-ru
       :linenos:

       Contract contract = new Contract();
       // по умолчанию состояние Offered
       contract.begin();
       // состояние меняется на Running
       assertEquals(Running.class, contract.status.class);

    Этот тест слишком сильно зависит от текущей реализации объекта status.
    Однако тест должен срабатывать даже в случае, если поле status станет булевским значением.
    Может быть, когда status меняется на Running, можно протестировать дату начала работы над контрактом:

    .. code-block:: java
       :name: code-2-ru
       :linenos:

       assertEquals(..., contract.startDate());
       // генерирует исключение, если status равен Offered

    Я признаю, что пытаюсь плыть против течения, когда настаиваю на том, что все тесты должны быть написаны только с использованием публичного (public) протокола.
    Существует специальный пакет JXUnit, который является расширением JUnit и позволяет тестировать значения переменных, даже тех, которые объявлены как закрытые.

    Желание протестировать объект в рамках концепции белого ящика — это не проблема тестирования, это проблема проектирования.
    Каждый раз, когда у меня возникает желание протестировать значение переменной-члена для того, чтобы убедиться в работоспособности кода, я получаю возможность улучшить дизайн системы.
    Если я забываю о своих опасениях и просто проверяю значение переменной, я теряю такую возможность.
    Иначе говоря, если идея об улучшении дизайна не приходит мне в голову, ничего не поделаешь.
    Я проверяю значение переменной, смахиваю непрошеную слезу, вношу соответствующую отметку в список задач и продолжаю двигаться вперед, надеясь, что наступит день, когда смогу найти подходящее решение.

    Thinking about objects as black boxes is hard. If I have a Contract with a Status that can be an instance of either Offered or Running , I might feel like writing a test based on my expected implementation:

    .. code-block:: java
       :name: code-1-en
       :linenos:

       Contract contract= new Contract();
       // Offered status by default
       contract.begin();
       // Changes status to Running
       assertEquals(Running.class, contract.status.class);

    This test is too dependent on the current implementation of status.
    The test should pass even if the representation of status changed to a boolean.
    Perhaps once the status changes to Running, it is possible to ask for the actual start date.

    .. code-block:: java
       :name: code-2-en
       :linenos:

       assertEquals(..., contract.startDate());
       // Throws an exception if the status is Offered

    I'm aware that I am swimming against the tide in insisting that all tests be written using only public protocol.
    There is even a package that extends JUnit called JXUnit, which allows testing the value of variables, even those declared private.

    Wishing for white box testing is not a testing problem, it is a design problem.
    Anytime I want to use a variable as a way of checking to see whether code ran correctly or not, I have an opportunity to improve the design.
    If I give in to my fear and just check the variable, then I lose that opportunity.
    That said, if the design idea doesn't come, it doesn't come. I'll check the variable, shed a tear, make a note to come back on one of my smarter days, and move on.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан

..

    Взгляд на тестирование в рамках TDD прагматичен.
    В TDD тесты являются средством достижения цели.
    Целью является код, в корректности которого мы в достаточной степени уверены.
    Если знание особенностей реализации без какого-либо теста дает нам уверенность в том, что код работает правильно, мы не будем писать тест.
    Тестирование черного ящика (когда мы намеренно игнорируем реализацию) обладает рядом преимуществ.
    Если мы игнорируем код, мы наблюдаем другую систему ценностей: тесты сами по себе представляют для нас ценность.
    В некоторых ситуациях это вполне оправданный подход, однако он отличается от TDD.

    TDD's view of testing is pragmatic.
    In TDD, the tests are a means to an end—the end being code in which we have great confidence.
    If our knowledge of the implementation gives us confidence even without a test, then we will not write that test.
    Black box testing, where we deliberately choose to ignore the implementation, has some advantages.
    By ignoring the code, it demonstrates a different value system—the tests are valuable alone.
    It's an appropriate attitude to take in some circumstances, but that is different from TDD.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан

..

    Many people make bad trade-offs, especially with heavy mocking.
    Kent thinks it's about trade-offs: is it worth making intermediate results testable?
    He used the example of a compiler where an intermediate parse-tree makes a good test point, and is also a better design.

    -- Kent Beck, "`Is TDD Dead? <https://martinfowler.com/articles/is-tdd-dead/>`__"

..

    Separate interface from implementation thinking.
    I have a tendency to pollute API design decisions with implementation speculation.
    I need to find a new way to separate the two levels of thinking while still providing rapid feedback between them.

    -- Kent Beck, "`RIP TDD <https://www.facebook.com/notes/kent-beck/rip-tdd/750840194948847/>`__"

..

    Структурная зависимость

    Структурная зависимость - одна из самых сильных и наиболее коварных форм зависимости тестов.
    Представьте набор тестов, в котором имеются тестовые классы для всех прикладных классов и тестовые методы для всех прикладных методов.
    Такой набор очень тесно связан со структурой приложения.


    Изменение в одном из прикладных методов или классов может повлечь необходимость изменить большое количество тестов.
    Следовательно, тесты слишком хрупкие и могут сделать прикладной код слишком жестким.

    Роль API тестирования - скрыть структуру приложения от тестов.
    Это позволит развивать прикладной код, не влияя на тесты. Это также позволит развивать тесты, не влияя на прикладной код.

    Такая возможность независимого развития абсолютно необходима, потому что с течением времени тесты становятся все более конкретными, а прикладной код, напротив, — все более абстрактным и обобщенным.
    Тесная структурная зависимость препятствует такому развитию - или, по меньшей мере, затрудняет его - и мешает прикладному коду становиться все более обобщенным и гибким.


    STRUCTURAL COUPLING

    Structural coupling is one of the strongest, and most insidious, forms of test coupling.
    Imagine a test suite that has a test class for every production class, and a set of test methods for every production method.
    Such a test suite is deeply coupled to the structure of the application.

    When one of those production methods or classes changes, a large number of tests must change as well.
    Consequently, the tests are fragile, and they make the production code rigid.

    The role of the testing API is to hide the structure of the application from the tests. 

    This allows the production code to be refactored and evolved in ways that don't affect the tests.
    It also allows the tests to be refactored and evolved in ways that don't affect the production code.

    This separation of evolution is necessary because as time passes, the tests tend to become increasingly more concrete and specific.
    In contrast, the production code tends to become increasingly more abstract and general.
    Strong structural coupling prevents - or at least impedes - this necessary evolution, and prevents the production code from being as general, and flexible, as it could be.

    -- "Clean Architecture: A Craftsman's Guide to Software Structure and Design" [#fncarch]_ by Robert C. Martin

..

    "Mock across architecturally significant boundaries, but not within those boundaries."

    -- "`When to Mock <https://blog.cleancoder.com/uncle-bob/2014/05/10/WhenToMock.html>`__" by Robert C. Martin


.. index::
   single: Sociable Test; in TDD
   single: Solitary Test; in TDD


Sociable or Solitary?
=====================

Наверное, самое часто заблуждение, которое мне приходилось слышать, это то, тесты должны быть полностью изолированы, и должны взаимодействовать только с `дублерами <https://martinfowler.com/bliki/TestDouble.html>`__.
Этот вопрос известен как "Solitary or Sociable?".

    Indeed using sociable unit tests was one of the reasons we were criticized for our use of the term "unit testing". I think that the term "unit testing" is appropriate because these tests are tests of the behavior of a single unit. We write the tests assuming everything other than that unit is working correctly.

    As xunit testing became more popular in the 2000's the notion of solitary tests came back, at least for some people. We saw the rise of Mock Objects and frameworks to support mocking. Two schools of xunit testing developed, which I call the classic and mockist styles. One of the differences between the two styles is that mockists insist upon solitary unit tests, while classicists prefer sociable tests. Today I know and respect xunit testers of both styles **(personally I've stayed with classic style)**.

    -- "`Unit Test <https://martinfowler.com/bliki/UnitTest.html#SolitaryOrSociable>`__" by Martin Fowler

..

    At the end of the day it's not important to decide if you go for solitary or sociable unit tests. Writing automated tests is what's important. Personally, I find myself using both approaches all the time.

    --  "`The Practical Test Pyramid <https://martinfowler.com/articles/practical-test-pyramid.html#SociableAndSolitary>`__" by Ham Vocke with support of Martin Fowler.

..

    TestDrivenDevelopment produces Developer Tests. The failure of a test case implicates only the developer's most recent edit. **This implies that developers don't need to use Mock Objects to split all their code up into testable units**. And it implies a developer may always avoid debugging by reverting that last edit.

    -- "`Unit Test <https://wiki.c2.com/?UnitTest>`__" on c2.com

Недостатки и достоинства обоих подходов описаны в статье "`Mocks Aren't Stubs <https://martinfowler.com/articles/mocksArentStubs.html>`__".

Мнение самого основателя TDD:

    "My personal practice - I mock almost nothing."

    -- Kent Beck, "`Is TDD Dead? Part 1 at 21:10 <https://youtu.be/z9quxZsLcfo?t=1269>`__

Лично я считаю что нужно ограничивать использование современных средства мокирования, активно эксплуатирующих Monkey Patch,  поскольку они позволяют создавать и тестировать низкокачественный код.


.. index::
   single: Design Patterns; in TDD


TDD и Design Patterns
=====================

Почему-то многие начинающие программисты, не знакомые с первоисточниками по TDD, думают, что TDD подразумевает только Evolutionary Design, а Simple Design противопоставляется паттернам программирования.

    Я обратил внимание на один важный эффект, который, я надеюсь, смогут принять во внимание и другие.
    Если на основе постоянно повторяющихся действий формулируются правила, дальнейшее применение этих правил становится неосознанным и автоматическим.
    Естественно, ведь это проще, чем обдумывать все за и все против того или иного действия с самого начала.
    Благодаря этому повышается скорость работы, и если в дальнейшем вы сталкиваетесь с исключением или проблемой, которая не вписывается ни в какие правила, у вас появляется дополнительное время и энергия для того, чтобы в полной мере применить свои творческие способности.

    Именно это произошло со мной, когда я писал книгу Smalltalk Best Practice Patterns (Лучшие паттерны Smalltalk).
    В какой-то момент я решил просто следовать правилам, описываемым в моей книге.
    В начале это несколько замедлило скорость моей работы, — мне требовалось дополнительное время, чтобы вспомнить то или иное правило, или написать новое правило.
    Однако по прошествии недели я заметил, что с моих пальцев почти мгновенно слетает код, над разработкой которого ранее мне приходилось некоторое время размышлять.
    Благодаря этому у меня появилось дополнительное время для анализа и важных размышлений о дизайне.

    Существует еще одна связь между TDD и паттернами: TDD является методом реализации дизайна, основанного на паттернах.
    Предположим, что в определенном месте разрабатываемой системы мы хотим реализовать паттерн Strategy (Стратегия).
    Мы пишем тест для первого варианта и реализуем его, создав метод.
    После этого мы намеренно пишем тест для второго варианта, ожидая, что на стадии рефакторинга мы придем к паттерну Strategy (Стратегия).
    Мы с Робертом Мартином (Robert Martin) занимались исследованием подобного стиля TDD.
    Проблема состоит в том, что дизайн продолжает вас удивлять.
    Идеи, которые на первый взгляд кажутся вам вполне уместными, позже оказываются неправильными.
    Поэтому я не рекомендую целиком и полностью доверять своим предчувствиям относительно паттернов.
    Лучше думайте о том, что, по-вашему, должна делать система, позвольте дизайну оформиться так, как это необходимо.

    The effect that I have noticed, and which I hope others find, is that by reducing repeatable behavior to rules, applying the rules becomes rote and mechanical.
    This is quicker than redebating everything from first principles all the time.
    When along comes an exception, or a problem that just doesn't fit any of the rules, you have more time and energy to generate and apply creativity.

    This happened to me when writing the Smalltalk Best Practice Patterns.
    At some point I decided just to follow the rules I was writing.
    It was much slower at first, to be looking up the rules, or to be stopping to write a new rule.
    After a week, however, I discovered that code was ripping off my fingertips that would have required a pause for thought before.
    This gave me more time and attention for bigger thoughts about design and analysis.
    Another relationship between TDD and patterns is TDD as an implementation method for pattern-driven design.
    Say we decide we want a Strategy for something.
    We write a test for the first variant and implement it as a method.
    Then we consciously write a test for the second variant, expecting the refactoring phase to drive us to a Strategy.
    Robert Martin and I did some research into this style of TDD.
    The problem is that the design keeps surprising you.
    Perfectly sensible design ideas turn out to be wrong.
    Better just to think about what you want the system to do, and let the design sort itself out later.

    -- "Test-Driven Development By Example" [#fntdd]_ by Kent Beck, перевод П. Анджан


..

    Добавление новой функциональности при помощи тестов и рефакторинг — это две монологические разновидности программирования.
    Совсем недавно я открыл еще одну разновидность: копирование паттерна.
    Я занимался разработкой сценария на языке Ruby, выполняющего извлечение информации из базы данных.
    Я начал с создания класса, являющегося оболочкой таблицы базы данных, а затем сказал себе, что раз я только что закончил книгу о паттернах работы с базами данных, я должен использовать паттерн.
    Примеры программ в книге были написаны на Java, поэтому нужный мне код легко можно было перенести на Ruby.
    Когда я программировал, я не думал о решении проблемы, я думал лишь о том, каким образом лучше всего адаптировать паттерн для условий, в рамках которых я работал.

    Копирование паттернов само по себе не является хорошим программированием, — я всегда подчеркиваю этот факт, когда говорю о паттернах.
    Любой паттерн — это полуфабрикат, — вы должны адаптировать его для условий своего проекта.
    Однако чтобы сделать это, лучше всего вначале, особо не задумываясь, скопировать паттерн, а затем, воспользовавшись смесью рефакторинга и TDD, выполнить адаптацию.
    В этом случае в процессе копирования паттерна вы также концентрируетесь только на одной вещи — на паттерне.
    Сообщество ХР интенсивно работает над тем, чтобы добавить в общую картину паттерны.
    Со всей очевидностью можно сказать, что сообщество ХР любит паттерны.
    В конце концов, между множеством приверженцев ХР и множеством приверженцев паттернов существует значительное пересечение: Уорд и Кент являются лидерами обоих направлений.
    Наверное, копирование паттерна — это третий монологический режим программирования наряду с разработкой в стиле "тесты вначале" и рефакторингом.
    Как и первые два режима, копирование паттерна — опасная штука, если ее использовать отдельно от двух других режимов.
    Все три вида программирования проявляют свою мощь только тогда, когда используются совместно друг с другом.

    Adding features test-first and refactoring are two of these monological flavors of programming.
    At a recent stint at the keyboard I experienced another one: pattern copying.
    I was writing a little Ruby script that pulled some data out of a database.
    As I did this I started on a class to wrap the database table and thought to myself that since I'd just finished off a book of database patterns I should use a pattern.
    Although the sample code was Java, it wasn't difficult to adapt it to Ruby.
    While I programmed it I didn't really think about the problem, I just thought about making a fair adaptation of the pattern to the language and specific data I was manipulating.
    Pattern copying on its own isn't good programming—a fact I always stress when talking about patterns.
    Patterns are always half baked, and need to be adapted in the oven of your own project.
    But a good way to do this is to first copy the pattern fairly blindly, and then use some mix of refactoring or test-first, to perform the adaptation.
    That way when you're doing the pattern-copying, you can concentrate on just the pattern—one thing at a time.
    The XP community has struggled with where patterns fit into the picture.
    Clearly the XP community is in favor of patterns, after all there is huge intersection between XP advocates and patterns advocates — Ward and Kent were leaders in both.
    Perhaps pattern copying is a third monological mode to go with test-first and refactoring, and like those two is dangerous on its own but powerful in concert.

    -- Martin Fowler, Afterword, "Test-Driven Development By Example" [#fntdd]_, перевод П. Анджан

..

    Patterns and XP

    The JUnit example leads me inevitably into bringing up patterns. The relationship between patterns and XP is interesting, and it's a common question. Joshua Kerievsky argues that patterns are under-emphasized in XP and he makes the argument eloquently, so I don't want to repeat that. But it's worth bearing in mind that for many people patterns seem in conflict to XP.

    The essence of this argument is that patterns are often over-used. The world is full of the legendary programmer, fresh off his first reading of GOF who includes sixteen patterns in 32 lines of code. I remember one evening, fueled by a very nice single malt, running through with Kent a paper to be called "Not Design Patterns: 23 cheap tricks" We were thinking of such things as use an if statement rather than a strategy. The joke had a point, patterns are often overused, but that doesn't make them a bad idea. The question is how you use them.

    One theory of this is that the forces of simple design will lead you into the patterns. Many refactorings do this explicitly, but even without them by following the rules of simple design you will come up with the patterns even if you don't know them already. This may be true, but is it really the best way of doing it? Surely it's better if you know roughly where you're going and have a book that can help you through the issues instead of having to invent it all yourself. I certainly still reach for GOF whenever I feel a pattern coming on. For me effective design argues that we need to know the price of a pattern is worth paying - that's its own skill. Similarly, as Joshua suggests, we need to be more familiar about how to ease into a pattern gradually. In this regard XP treats the way we use patterns differently to the way some people use them, but certainly doesn't remove their value.

    But reading some of the mailing lists I get the distinct sense that many people see XP as discouraging patterns, despite the irony that most of the proponents of XP were leaders of the patterns movement too. Is this because they have seen beyond patterns, or because patterns are so embedded in their thinking that they no longer realize it? I don't know the answers for others, but for me patterns are still vitally important. XP may be a process for development, but patterns are a backbone of design knowledge, knowledge that is valuable whatever your process may be. Different processes may use patterns in different ways. XP emphasizes both not using a pattern until it's needed and evolving your way into a pattern via a simple implementation. But patterns are still a key piece of knowledge to acquire.

    My advice to XPers using patterns would be

    - Invest time in learning about patterns
    - Concentrate on when to apply the pattern (not too early)
    - Concentrate on how to implement the pattern in its simplest form first, then add complexity later.
    - If you put a pattern in, and later realize that it isn't pulling its weight - don't be afraid to take it out again.

    I think XP should emphasize learning about patterns more. I'm not sure how I would fit that into XP's practices, but I'm sure Kent can come up with a way.

    -- "`Is Design Dead? <https://martinfowler.com/articles/designDead.html#PatternsAndXp>`__" by Martin Fowler

Смотрите так же:

- XP and Patterns Ralph Johnson's View:  http://objectclub.jp/community/XP-jp/xp_relate/xp_patterns
- Joshua Kerievsky, Patterns & XP: http://www.industriallogic.com/xp/PatternsAndXP.pdf


One assertion per test?
=======================

Я часто слышу это распространенное убеждение, что один тестовый метод должен содержать только одно утверждение (assertion) и не больше.
И это интересно, потому что Кент Бек этому правилу не очень-то и следует, что заставило меня найти первоисточник этого убеждения.
Источник я нашел, и он, действительно, авторитетный - это "xUnit Test Patterns. Refactoring Test Code." by Gerard Meszaros, глава "Principle: Verify One Condition per Test", но там есть кое-что еще, о чем это широко распространенное убеждение умалчивает:

    One possibly contentious aspect of "Verify One Condition per Test" is what we mean by "one condition."
    Some test drivers insist on one assertion per test.
    This insistence may be based on using a "Testcase Class per Fixture" organization
    of the "Test Methods" and naming each test based on what the one assertion is verifying.
    (For example, AwaitingApprovalFlight.validApproverRequestShouldBeApproved.)
    Having one assertion per test makes such naming very easy but also leads to many more test methods if we have to assert on many output fi elds.
    Of course, we can often comply with this interpretation by extracting a "Custom Assertion" (page 474)
    or "Verification Method" (see "Custom Assertion") that allows us to reduce the multiple assertion method calls to a single call.
    Sometimes that approach makes the test more readable.
    When it doesn't, I wouldn't be too dogmatic about insisting on a single assertion.

    -- "xUnit Test Patterns. Refactoring Test Code." by Gerard Meszaros

Эту же тему рассматривает и Robert C. Martin в главе "Chapter 9: Unit Tests :: One Assert per Test" книги "Clean Code: A Handbook of Agile Software Craftsmanship":

    Я думаю, что правило "одного assert" является хорошей рекомендацией.
    Обычно я стараюсь создать предметно-ориентированный язык тестирования, который это правило поддерживает, как в листинге 9.5.
    Но при этом я не боюсь включать в свои тесты более одной директивы assert.
    Вероятно, лучше всего сказать, что количество директив assert в тесте должно быть сведено к минимуму.

    I think the single assert rule is a good guideline.
    I usually try to create a domainspeciﬁc testing language that supports it, as in Listing 9-5.
    But I am not afraid to put more than one assert in a test.
    I think the best thing we can say is that the number of asserts in a test ought to be minimized.

    -- "Clean Code: A Handbook of Agile Software Craftsmanship" [#fnccode]_ by Robert C. Martin, перевод Е. Матвеев

Здесь он отсылает к статье "`One Assertion Per Test <https://www.artima.com/weblogs/viewpost.jsp?thread=35578>`__" by Dave Astels в качестве первоисточника.

.. rubric:: Footnotes

.. [#fntdd] "Test-Driven Development By Example" by Kent Beck
.. [#fnccoder] "The Clean Coder: a code of conduct for professional programmers" by Robert C. Martin
.. [#fnccode] "Clean Code: A Handbook of Agile Software Craftsmanship" by Robert C. Martin
.. [#fncarch] "Clean Architecture: A Craftsman's Guide to Software Structure and Design" by Robert C. Martin
.. [#fnrefactoring] "Refactoring: Improving the Design of Existing Code" by Martin Fowler, Kent Beck, John Brant, William Opdyke, Don Roberts

