:canonical-base-url: https://dckms.github.io/system-architecture

.. index:: Agile Development
   :name: emacsway-agile-development


===========================
Что такое Agile Development
===========================

.. sectionauthor:: Ivan Zakrevsky

..

    📝 "Agile development - software development approach based on :ref:`iterative development <emacsway-iterative-development>`, frequent inspection and adaptation, and incremental deliveries, in which requirements and solutions evolve through collaboration in cross‐functional teams and through continual stakeholder feedback."

    -- "ISO/IEC/IEEE 12207:2017 Systems and software engineering - Software life cycle processes"

..

    📝 "\"Agile\" methods actually can be applied within a variety of models.
    While Agile methods are common in executing an evolutionary lifecycle model, they can be used in other lifecycle models at various stages.
    What the methods have in common is an emphasis on continuous inspection and collaboration in the rapid production of working software in an environment where changes, including changes to requirements, are expected."

    -- "ISO/IEC/IEEE 12207:2017 Systems and software engineering - Software life cycle processes"

..

    📝 "As discussed in 5.4.2, the life cycle models used in agile projects are often highly :ref:`incremental <emacsway-incremental-development>` and :ref:`evolutionary <emacsway-evolutionary-development>`."

    -- "ISO/IEC/IEEE 12207:2017 Systems and software engineering - Software life cycle processes"

Agile является естественным следствием эволюции итеративной разработки, краткий обзор которой можно посмотреть в превосходной статье Craig Larman "`Iterative and Incremental Development: A Brief History <https://www.craiglarman.com/wiki/downloads/misc/history-of-iterative-larman-and-basili-ieee-computer.pdf>`__".

Как уже говорилось ранее, итеративная модель разработки открывает широкие возможности для :ref:`удешевления обработки неопределенности <emacsway-adaptation>`.
Однако долгое время эти возможности оставались экономически нецелесообразными по причине быстрорастущего характера роста стоимости :ref:`Adaptation <emacsway-adaptation>`, приближющегося к экспоненциальному.
При таком характере роста возникает экономическая целесообразность принимать решения в момент наименьшей стоимости их реализации, вплоть до заблаговременного проектирования (BDUF).

.. figure:: _media/agile/exponential-cost-of-change.png
   :alt: Figure 1. The cost of change rising exponentially over time. The image source is "Extreme Programming Explained" 1st edition by Kent Beck, "Chapter 5. Cost of Change".
   :align: left
   :width: 90%

   Figure 1. The cost of change rising exponentially over time. The image source is "Extreme Programming Explained" 1st edition by Kent Beck, "Chapter 5. Cost of Change".

Однако, в конце 1990-х - начале 2000-х, в архитектурном мире произошли существенные изменения - обрели массовую популярность высокоуровневые объектно-ориентированные языки, появились шаблоны и принципы проектирования, методики управления сложностью (ROM, POSA, GOF, OOAD, SOLID, Use Case Driven Approach, Object-Oriented Software Construction etc.), появились TDD, Refactoring и т.п.

Унификация знаний в области архитектуры, переход ментального оперирования на элементы унифицированных шаблонных конструкций более высокого уровня абстракции, позволили сократить когнитивную и коммуникативную нагрузку на разработчика, уменьшить порог вхождения в новый проект, смягчить негативное воздействие Закона Брукса.

Рост количественных изменений привел к изменениям качественным - ведущим умам архитектуры своего времени удалось снизить характер роста стоимости адаптации вплоть до пологого графика, максимально приближенного к горизонтальной асимптоте.
Это означало, что стоимость реализации решения больше не зависело от момента его принятия, что позволило отказаться от заблаговременного проектирования и откладывать принятие решения до момента наибольшей полноты информированности, даже после частичной реализации продукта.

    📝 "What would we do if all that investment paid off?
    What if all that work on languages and databases and whatnot actually got somewhere?
    What if the cost of change didn't rise exponentially overtime, but rose much more slowly, **eventually reaching an asymptote**?
    What if tomorrow's software engineering professor draws Figure 3 on the board?"

    -- "Extreme Programming Explained" 1st edition by Kent Beck, "Chapter 5. Cost of Change"

.. figure:: _media/agile/flatten-cost-of-change.png
   :alt: Figure 3. The cost of change may not rise dramatically over time. The image source is "Extreme Programming Explained" 1st edition by Kent Beck, "Chapter 5. Cost of Change".
   :align: left
   :width: 90%

   Figure 3. The cost of change may not rise dramatically over time. The image source is "Extreme Programming Explained" 1st edition by Kent Beck, "Chapter 5. Cost of Change".

Что такое асимтота, можно посмотреть в "§284 Асимтоты" Справочника по высшей математике / М.Я. Выгодский:

    📝 "Прямая АВ называется асимптотой линии L, если расстояние МК (черт. 297) от точки М линии L до прямой АВ стремится к нулю при удалении точки М в бесконечность."

    -- "Справочник по высшей математике" / М.Я. Выгодский

В нашем случае, нас интересует Асимптоты, параллельная оси абсцисс (там же):

    📝 "Для разыскания горизонтальных асимптот линии y = f(х) ищем пределы f(х) при х -> +∞ и при х -> -∞. Если lim х->∞ f(x) = b, то прямая у = b - асимптота (при бесконечном удалении вправо; черт. 299)."

    -- "Справочник по высшей математике" / М.Я. Выгодский

Вся суть Agile (итеративной) модели разработки была лаконично и метко выражена Кент Беком всего одним предложением:

    📝 "Сделайте изменение легким, а потом делай легко изменение.

    **Make the change easy then make the easy change.**"

    -- Kent Beck, dddeu 20

Невероятный талант Kent Beck объяснять сложные вещи простым языком. Именно об этом я говорил здесь (https://t.me/emacsway_log/810). И это при необычайной эрудированности Kent Beck. Cписок использованной литературы в его книгах просто ошеломляет.

Thanks to Vladik Khononov for https://youtu.be/ybYtgII151g?t=9808

Более развернутый вариант его фразы:

    📝 "At the core of understanding this argument is the software change curve.
    The change curve says that as the project runs, it becomes exponentially more expensive to make changes.
    The change curve is usually expressed in terms of phases "a change made in analysis for $1 would cost thousands to fix in production".
    This is ironic as most projects still work in an ad-hoc process that doesn't have an analysis phase, but the exponentiation is still there.
    **The exponential change curve means that evolutionary design cannot possibly work.**
    It also conveys why planned design must be done carefully because any mistakes in planned design face the same exponentiation.

    **The fundamental assumption underlying XP is that it is possible to flatten the change curve enough to make evolutionary design work.**
    This flattening is both enabled by XP and exploited by XP.
    This is part of the coupling of the XP practices: specifically **you can't do those parts of XP that exploit the flattened curve without doing those things that enable the flattening.**
    This is a common source of the controversy over XP.
    Many people criticize the exploitation without understanding the enabling.
    Often the criticisms stem from critics' own experience where they didn't do the enabling practices that allow the exploiting practices to work.
    As a result they got burned and when they see XP they remember the fire."

    -- "`Is Design Dead? <https://martinfowler.com/articles/designDead.html>`__" by M.Fowler

..

    📝 "**This is one of the premises of XP. It is the technical premise of XP.**
    If the cost of change rose slowly over time, you would act completely differently from how you do under the assumption that costs rise exponentially.
    You would make big decisions as late in the process as possible, to defer the cost of making the decisions and to have the greatest possible chance that they would be right.
    You would only implement what you had to, in hopes that the needs you anticipate for tomorrow wouldn’t come true.
    You would introduce elements to the design only as they simplified existing code or made writing the next bit of code simpler.

    **If a flattened change cost curve makes XP possible, a steep change cost curve makes XP impossible.**
    If change is ruinously expensive, you would be crazy to charge ahead without careful forethought.
    But if change stays cheap, the additional value and reduced risk of early concrete feedback outweighs the additional cost of early change."

    -- "Extreme Programming Explained" 1st edition by Kent Beck

Поскольку это было произнесено еще до встречи 2001 года и принятия Agile Manifesto, то под XP следует понимать Agile (или даже любую итератиную модель разработки) в принципе, поскольку XP - это частный случай Agile.

Иными словами, внутреннее качество программы является первичным условием в Agile и в любой другой итеративной разработке.

    📝 "Engineers who don't understand exponential growth and the cost curve as economies of scale kick in come to wildly incorrect conclusions."

    -- Kent Beck, https://twitter.com/KentBeck/status/1402276528910704655?s=19

..

    📝 "The incremental and iterative nature of Agile development can facilitate **efficient technical and management processes and practices to reduce the cost associated with change**.
    In comparison, projects managed at the waterfall end of the continuum seek to reduce total rework cost by minimizing the number of changes, limiting the number of control points, and baselining detailed specifications which are reviewed and traced throughout the project."

    -- "ISO/IEC/IEEE 12207:2017 Systems and software engineering - Software life cycle processes"



.. seealso::

   - ":ref:`emacsway-adaptation`"
   - ":ref:`emacsway-prediction`"
   - ":ref:`emacsway-balancing-prediction-adaptation`"
   - ":doc:`../../uncertainty-management/adaptation/crash-course-in-software-development-economics`"
