:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: Software Design; in Agile
   :name: emacsway-agile-software-design


================================
Role of Software Design in Agile
================================

.. sectionauthor:: Ivan Zakrevsky

.. figure:: _media/software-design/ouroboros.jpg
   :alt: Уроборос. Иллюстрация из открытых источников неизвестного автора.
   :align: center
   :width: 40%

   Уроборос. Иллюстрация из открытых источников неизвестного автора.

.. contents:: Содержание


.. index::
   single: Software Design; for those who read the code
   :name: emacsway-who-reads-the-code

Кто читает код?
===============

Среди малоопытных разработчиков иногда можно услышать, что им некогда писать качественный код, так как у них мало времени, а этот код все равно читать никто не будет.

Истина в том, что в процессе конструирования кода, 91% времени занимает именно чтение кода и борьба со сложностью, и только 9% времени (1:10) занимает ввод символов с клавиатуры.
И это соотношение дано еще для качественного кода.

Причем, разработчик пишет код в одиночку и лишь единожды.
Зато читают код все разработчики команды и по много раз.

Таким образом, плохо написанный код более чем на 91% (т.е. более чем на порядок!) влияет на снижение темпов разработки всей команды.

Кстати, этот момент хорошо объясняет то, почему :ref:`при TDD разработка осуществляется быстрее <emacsway-why-is-tdd-faster>`, хотя объема кода пишется больше.

    📝 "Кто-то спросит: так ли уж часто читается наш код?
    Разве большая часть времени не уходит на его написание?

    Вам когда-нибудь доводилось воспроизводить запись сеанса редактирования?
    В 80-х и 90-х годах существовали редакторы, записывавшие все нажатия клавиш (например, Emacs). Вы могли проработать целый час, а потом воспроизвести весь сеанс, словно ускоренное кино.
    Когда я это делал, результаты оказывались просто потрясающими.

    Большинство операций относилось к прокрутке и переходу к другим модулям!

    - Боб открывает модуль.
    - Он находит функцию, которую необходимо изменить.
    - Задумывается о последствиях.
    - Ой, теперь он переходит в начало модуля, чтобы проверить инициализацию переменной.
    - Снова возвращается вниз и начинает вводить код.
    - Стирает то, что только что ввел.
    - Вводит заново.
    - Еще раз стирает!
    - Вводит половину чего-то другого, но стирает и это!
    - Прокручивает модуль к другой функции, которая вызывает изменяемую функцию, чтобы посмотреть, как она вызывается.
    - Возвращается обратно и восстанавливает только что стертый код.
    - Задумывается.
    - Снова стирает!
    - Открывает другое окно и просматривает код субкласса. Переопределяется ли в нем эта функция?

    <...>

    В общем, вы поняли.
    На самом деле соотношение времени чтения и написания кода превышает 10:1.
    Мы постоянно читаем свой старый код, поскольку это необходимо для написания нового кода.

    Из-за столь высокого соотношения наш код должен легко читаться, даже если это затрудняет его написание.
    Конечно, написать код, не прочитав его, невозможно, так что упрощение чтения в действительности упрощает и написание кода.
    Уйти от этой логики невозможно.
    Невозможно написать код без предварительного чтения окружающего кода.
    Код, который вы собираетесь написать сегодня, будет легко или тяжело писаться в зависимости от того, насколько легко или тяжело читается окружающий код.
    Если вы хотите быстро справиться со своей задачей, если вы хотите, чтобы ваш код было легко писать — позаботьтесь о том, чтобы он легко читался.

    You might ask: How much is code really read? Doesn't most of the effort go into writing it?

    Have you ever played back an edit session? In the 80s and 90s we had editors like Emacs that kept track of every keystroke.
    You could work for an hour and then play back your whole edit session like a high-speed movie.
    When I did this, the results were fascinating.

    The vast majority of the playback was scrolling and navigating to other modules!

    - Bob enters the module.
    - He scrolls down to the function needing change.
    - He pauses, considering his options.
    - Oh, he's scrolling up to the top of the module to check the initialization of a variable.
    - Now he scrolls back down and begins to type.
    - Ooops, he's erasing what he typed!
    - He types it again.
    - He erases it again!
    - He types half of something else but then erases that!
    - He scrolls down to another function that calls the function he's changing to see how it is called.
    - He scrolls back up and types the same code he just erased.
    - He pauses.
    - He erases that code again!
    - He pops up another window and looks at a subclass. Is that function overridden?

    <...>

    You get the drift. Indeed, the ratio of time spent reading vs. writing is well over 10:1.
    We are constantly reading old code as part of the effort to write new code.

    Because this ratio is so high, we want the reading of code to be easy, even if it makes the writing harder.
    Of course there's no way to write code without reading it, so making it easy to read actually makes it easier to write.

    There is no escape from this logic.
    You cannot write code if you cannot read the surrounding code.
    The code you are trying to write today will be hard or easy to write depending on how hard or easy the surrounding code is to read.
    So if you want to go fast, if you want to get done quickly, if you want your code to be easy to write, make it easy to read."

    -- "Clean Code: A Handbook of Agile Software Craftsmanship" by Robert C. Martin, перевод: Е.Матвеев, ООО Издательство "Питер"


.. index:: Primary Technical Imperative
   :name: emacsway-primary-technical-imperative

Primary Technical Imperative
============================


    📝 "There are two ways of constructing a software design: one way is to make it so simple that there are obviously no deficiencies, and the other is to make it so complicated that there are no obvious deficiencies."

    -- C. A. R. Hoare

..

    📝 "Управление сложностью — самый важный технический аспект разработки ПО.
    По-моему, управление сложностью настолько важно, что оно должно быть Главным Техническим Императивом Разработки ПО.

    Managing complexity is the most important technical topic in software development.
    In my view, it's so important that Software's Primary Technical Imperative has to be managing complexity."

    -- "Code Complete" 2nd edition by Steve McConnell, перевод: Издательско-торговый дом "Русская Редакция"

..

    📝 "Дейкстра пишет, что ни один человек не обладает интеллектом, способным вместить все детали современной компьютерной программы (Dijkstra, 1972), поэтому нам - разработчикам ПО — не следует пытаться охватить всю программу сразу.
    Вместо этого мы должны попытаться организовать программы так, чтобы можно было безопасно работать с их отдельными фрагментами по очереди.
    Целью этого является минимизация объема программы, о котором нужно думать в конкретный момент времени.
    Можете считать это своеобразным умственным жонглированием: чем больше умственных шаров программа заставляет поддерживать в воздухе,
    тем выше вероятность того, что вы уроните один из них и допустите ошибку при проектировании или кодировании.

    На уровне архитектуры ПО сложность проблемы можно снизить, разделив систему на подсистемы.
    Несколько несложных фрагментов информации понять проще, чем один сложный.
    В разбиении сложной проблемы на простые фрагменты и заключается цель всех методик проектирования ПО.
    Чем более независимы подсистемы, тем безопаснее сосредоточиться на одном аспекте сложности в конкретный момент времени.
    Грамотно определенные объекты разделяют аспекты проблемы так, чтобы вы могли решать их по очереди.
    Пакеты обеспечивают такое же преимущество на более высоком уровне агрегации.

    Стремление к краткости методов программы помогает снизить нагрузку на интеллект.
    Этому же способствует написание программы в терминах проблемной области, а не низкоуровневых деталей реализации,
    а также работа на самом высоком уровне абстракции.

    Суть сказанного в том, что программисты, компенсирующие изначальные ограничения человеческого ума,
    пишут более понятный и содержащий меньшее число ошибок код.

    Dijkstra pointed out that no one's skull is really big enough to contain a modern computer program (Dijkstra 1972),
    which means that we as software developers shouldn't try to cram whole programs into our skulls at once;
    we should try to organize our programs in such a way that we can safely focus on one part of it at a time.
    The goal is to minimize the amount of a program you have to think about at any one time.
    You might think of this as mental juggling—the more mental balls the program requires you
    to keep in the air at once, the more likely you'll drop one of the balls, leading to a design or coding error.

    At the software-architecture level, the complexity of a problem is reduced by dividing the system into subsystems.
    Humans have an easier time comprehending several simple pieces of information than one complicated piece.
    The goal of all software-design techniques is to break a complicated problem into simple pieces.
    The more independent the subsystems are, the more you make it safe to focus on one bit of complexity at a time.
    Carefully defined objects separate concerns so that you can focus on one thing at a time.
    Packages provide the same benefit at a higher level of aggregation.

    Keeping routines short helps reduce your mental workload.
    Writing programs in terms of the problem domain, rather than in terms of low-level implementation details, and
    working at the highest level of abstraction reduce the load on your brain.

    The bottom line is that programmers who compensate for inherent human limitations
    write code that's easier for themselves and others to understand and that has fewer errors."

    -- "Code Complete" 2nd edition by Steve McConnell, перевод: Издательско-торговый дом "Русская Редакция"

..

    📝 "**Главным Техническим Императивом Разработки ПО является управление сложностью.**
    Управлять сложностью будет гораздо легче, если при проектировании вы будете стремиться к простоте.

    Есть два общих способа достижения простоты:
    минимизация объема существенной сложности, с которой приходится иметь дело в любой конкретный момент времени,
    и подавление необязательного роста несущественной сложности.

    **Software's Primary Technical Imperative is managing complexity.**
    This is greatly aided by a design focus on simplicity.

    Simplicity is achieved in two general ways:
    minimizing the amount of essential complexity that anyone's brain has to deal with at any one time,
    and keeping accidental complexity from proliferating needlessly."

    -- "Code Complete" 2nd edition by Steve McConnell, перевод: Издательско-торговый дом "Русская Редакция"

..

    📝 "При выполнении других заданий человек может удерживать в памяти 7±2 дискретных элементов [Miller, 1956].
    Если класс содержит более семи элементов данных-членов, подумайте, не разделить ли его на несколько менее крупных классов [Riel, 1996].

    The number "7±2" has been found to be a number of discrete items a person can remember while performing other tasks [Miller 1956].
    If a class contains more than about seven data members, consider whether the class should be decomposed into multiple smaller classes [Riel 1996].

    [Miller, 1956]
        Miller, G. A. 1956. "The Magical Number Seven, Plus or Minus Two: Some Limits on Our Capacity for Processing Information."
        The Psychological Review 63, no. 2 (2): 81–97.
    [Riel 1996]
        Riel, Arthur J. 1996. Object-Oriented Design Heuristics. Reading, MA: Addison-Wesley."

    -- "Code Complete" 2nd edition by Steve McConnell, перевод: Издательско-торговый дом "Русская Редакция"

По поводу последнего изречения - лучше один раз увидеть на примере метафоры в виде картинки со схожим эффектом:

.. figure:: _media/software-design/12-points.jpg
   :alt: Просто ваши глаза не могут увидеть все 12 точек одновременно. 
         Ninio's extinction illusion. Twelve black dots cannot be seen at once.
         Ninio, J. and Stevens, K. A. (2000) Variations on the Hermann grid: an extinction illusion. Perception, 29, 1209-1217.
         The image source is a post by Akiyoshi Kitaoka https://www.facebook.com/akiyoshi.kitaoka/posts/10207806663219295
   :align: left
   :width: 90%

   Просто ваши глаза не могут увидеть все 12 точек одновременно.
   Ninio's extinction illusion. Twelve black dots cannot be seen at once.
   Ninio, J. and Stevens, K. A. (2000) Variations on the Hermann grid: an extinction illusion. Perception, 29, 1209-1217.
   The image source is "`a post <https://www.facebook.com/akiyoshi.kitaoka/posts/10207806663219295>`__" by Akiyoshi Kitaoka.

Как и в "Законе Миллера", суть картинки сводится к тому, что у человека есть предел способности **воспринимать** информацию, и если количество единиц поступающей информации превышает этот предел (не зависимо от его природы, будь то особенность работы рецепторов сетчатки или предел возможностей краткосрочной памяти), то начинается "жонглирование", т.е. неспособность рассмотреть (в прямом и в переносном смыслах) всю информацию единовременно и изолированно.

Вероятное объяснение этого явления заключается в том, что:

    💬 "Your eye's receptors are stimulated and influenced by the activity of neighboring receptors. In a complex, repetitive grid like this, one receptor can have trouble perceiving the dots accurately because of stimulation occurring in a nearby receptor."

    -- `источник <https://www.brainhq.com/brain-resources/brain-teasers/ninios-extinction-illusion/>`__

**Внимание** - это избирательная направленность **восприятия**.
Периферийное зрение - это способность видеть те предметы, которые выходят за **фокус** основного **внимания**.
Слово "сфокусировать" - означает "сосредоточить", как в прямом (оптическом), так и в переносном (сконцентрироваться) смыслах.
Основной принцип управления сложностью - это её декомпозиция до такого уровня, над которым обеспечивается перевес умственных возможностей человека. Т.е. когда объем рассматриваемой изолированно сложности "вмещается" в **фокус** внимания человека.

См. также ":ref:`emacsway-icebreaker-principle`".

.. _emacsway-kent-beck-constantine's-law:

    📝 "These were elucidated in the mid-70s by Yourdon & Constantine in `Structured Design <https://amzn.to/2GsuXvQ>`__ and haven't changed.
    Their argument goes like this:

    #. We design software to reduce its cost.
    #. The cost of software is ≈ the cost of changing the software.
    #. The cost of changing the software is ≈ the cost of the expensive changes (power laws and all that).
    #. The cost of the expensive changes is generated by cascading changes — if I change this then I have to change that and that, and if I change that then…
    #. Coupling between elements of a design is this propensity for a change to propagate.
    #. So, design ≈ cost ≈ change ≈ big change ≈ coupling. Transitively, software design ≈ managing coupling.

    (This skips loads of interesting stuff, but I'm just trying to set up the argument for why rapid decomposition of a monolith into micro-services is counter-productive.)"

    Managing Coupling

    Note I don't say, "Eliminating coupling."
    Decoupling comes with its own costs, both the cost of the decoupling itself and the future costs of unanticipated changes.
    The more perfectly a design is adapted to one set of changes, the more likely it is to be blind-sided by novel changes. And so we have the classic tradeoff curve:

    .. figure:: _media/software-design/balancing-coupling-decoupling.jpeg
       :alt: Classic tradeoff curve of balancing cost of Coupling vs. cost of Decoupling. The image source is article "Monolith -> Services: Theory & Practice" by Kent Beck https://medium.com/@kentbeck_7670/monolith-services-theory-practice-617e4546a879
       :align: left
       :width: 90%

       Classic tradeoff curve of balancing cost of Coupling vs. cost of Decoupling. The image source is article "`Monolith -> Services: Theory & Practice <https://medium.com/@kentbeck_7670/monolith-services-theory-practice-617e4546a879>`__" by Kent Beck.

    You manage coupling one of two ways:

    1. Eliminate coupling. A client and server with hard-coded read() and write() functions are coupled with respect to protocol changes. Change a write() and you'll have to change the read(). Introduce an interface definition language, though, and you can add to the protocol in one place and have the change propagate automatically to read() and write().
    2. Reduce coupling's scope. If changing one element implies changing ten others, then it's better if those elements are together than if they are scattered all over the system —less to navigate, less to examine, less to test. The number of elements to change is the same, but the cost per change is smaller. (This is also known as the "manure in one pile" principle, or less-aromatically "cohesion".)

    -- "`Monolith -> Services: Theory & Practice <https://medium.com/@kentbeck_7670/monolith-services-theory-practice-617e4546a879>`__" by Kent Beck

Eric Evans дает хорошее определение Constantine's Law нетехническим языком:

    💬 "МОДУЛИ дают возможность посмотреть на модель с разных сторон: во-первых, можно изучить подроб­ ности устройства модуля, не вникая в сложное целое; во-вторых, удобно рассматривать взаимоотношения между модулями, не вдава я сь в детали их внутреннего устройства.

    <...>

    То, что при делении на модули должна соблюдаться низкая внешняя зависимость 
(low coupling) при высокой внутренней связности (high cohesion)- это общие слова. 
    Определения зависимости и связности rрешат уклоном в чисто технические, количест­венные критерии, по которым их якобы можно измерить, подсчитав количество ассо­циаций и взаимодействий. Но это не просто механические характеристики подразде­ления кода на модули, а идейные концепции. Человек не может одновременно удер­живать в уме слишком много предметов (отсюда низкая внешняя зависимость). 
    А плохо связанные между собой фраrменты информации так же трудно понять, как неструктурированную "кашу" из идей (отсюда высокая внутренняя связность).

    MODULES give people two views of the model: They can look at detail within a MODULE without being overwhelmed by the whole, or they can look at relationships between MODULES in views that exclude interior detail.

    <...>

    It is a truism that there should be low coupling between MODULES and high cohesion 
within them.
    Explanations of coupling and cohesion tend to make them sound like technical metrics, to be judged mechanically based on the distributions of associations and interactions. Yet it isn't just code being divided into MODULES, but concepts.
    There is a limit to how many things a person can think about at once (hence low coupling). 
Incoherent fragments of ideas are as hard to understand as an undifferentiated soup of ideas (hence high cohesion)."

    -- "Domain-Driven Design: Tackling Complexity in the Heart of Software" by Eric Evans


Оправдано ли качество?
======================

Martin Fowler
-------------

    📝 "In most contexts higher quality ⇒ expensive. But high internal quality of software allows us to develop features faster and cheaper."

    -- "`Tradable Quality Hypothesis <https://martinfowler.com/bliki/TradableQualityHypothesis.html>`__" by Martin Fowler

.. _emacsway-design-stamina-graph:

.. figure:: _media/software-design/design-stamina-graph.png
   :alt: The pseudo-graph plots delivered functionality (cumulative) versus time for two imaginary stereotypical projects: one with good design and one with no design. The image from "Design Stamina Hypothesis" by Martin Fowler. https://martinfowler.com/bliki/DesignStaminaHypothesis.html
   :align: left
   :width: 90%

   The pseudo-graph plots delivered functionality (cumulative) versus time for two imaginary stereotypical projects: one with good design and one with no design. The image from "`Design Stamina Hypothesis <https://martinfowler.com/bliki/DesignStaminaHypothesis.html>`__" by Martin Fowler.

..

    📝 "... the true value of internal quality - that it's the enabler to speed. The purpose of internal quality is to go faster."

    -- "`Tradable Quality Hypothesis <https://martinfowler.com/bliki/TradableQualityHypothesis.html>`__" by Martin Fowler

..

    📝 "The value of good software design is economic: you can continue to add new functionality quickly even as the code-base grows in size."

    -- "`Design Stamina Hypothesis <https://martinfowler.com/bliki/DesignStaminaHypothesis.html>`__" by Martin Fowler

..

    📝 "We usually perceive that it costs more to get higher quality, but software internal quality actually reduces costs."

    -- "`Is High Quality Software Worth the Cost? <https://martinfowler.com/articles/is-quality-worth-cost.html>`__" by Martin Fowler

..

    📝 "The fundamental role of internal quality is that it lowers the cost of future change.
    But there is some extra effort required to write good software, which does impose some cost in the short term."

    -- "`Is High Quality Software Worth the Cost? <https://martinfowler.com/articles/is-quality-worth-cost.html>`__" by Martin Fowler

..

    📝 "The whole point of good design and clean code is to make you go faster - if it didn't people like Uncle Bob, Kent Beck, and Ward Cunningham wouldn't be spending time talking about it."

    -- "`Technical Debt Quadrant <https://martinfowler.com/bliki/TechnicalDebtQuadrant.html>`__" by Martin Fowler

..

    📝 "Sadly, software developers usually don't do a good job of explaining this situation.
    Countless times I've talked to development teams who say "they (management) won't let us write good quality code because it takes too long".
    Developers often justify attention to quality by justifying through the need for proper professionalism.
    But this moralistic argument implies that this quality comes at a cost - dooming their argument.
    The annoying thing is that the resulting crufty code both makes developers' lives harder, and costs the customer money.
    When thinking about internal quality, I stress that we should only approach it as an economic argument.
    High internal quality reduces the cost of future features, meaning that putting the time into writing good code actually reduces cost.

    This is why the question that heads this article misses the point.
    The "cost" of high internal quality software is negative.
    The usual trade-off between cost and quality, one that we are used to for most decisions in our life, does not make sense with the internal quality of software.
    (It does for external quality, such as a carefully crafted user-experience.)
    Because the relationship between cost and internal quality is an unusual and counter-intuitive relationship, it's usually hard to absorb.
    But understanding it is critical to developing software at maximum efficiency."

    -- "`Is High Quality Software Worth the Cost? <https://martinfowler.com/articles/is-quality-worth-cost.html>`__" by Martin Fowler

..

    📝 "Рефакторинг ускоряет написание программ

    В конечном итоге все сказанное сводится к одному: рефакторинг ускоряет написание программ.

    Создается впечатление внутреннего противоречия.
    Когда я рассказываю о рефакторинге, становится очевидно, что он повышает качество кода.
    Улучшение проекта, повышение удобочитаемости, уменьшение количества ошибок — все это способствует качеству кода.
    Но разве скорость разработки не снижается из-за всего этого?

    Когда я общаюсь с разработчиками программного обеспечения, которые какое-то время работали над системой, я часто слышу, что сначала им удалось быстро продвинуться вперед, но теперь добавление новых функциональных возможностей занимает гораздо больше времени.
    Каждая новая функция требует все больше и больше времени, чтобы понять, как вписать ее в существующую кодовую базу, а после ее добавления часто возникают ошибки, исправление которых занимает еще больше времени.
    Кодовая база начинает выглядеть как серия исправлений, исправляющих предыдущие исправления, и требуются навыки археолога, чтобы выяснить, как все это работает.
    Все это замедляет добавление новых функциональных возможностей до такой степени, что зачастую разработчики хотят начать все заново с чистого листа.

    Визуализировать это положение вещей можно с помощью следующего псевдографика.

    Но некоторые команды сообщают о другом опыте.
    Они утверждают, что могут добавлять новые функциональные возможности быстрее, потому что они могут использовать уже существующий код, опираясь на то, что уже имеется в наличии.

    Разница между этими проектами заключается во внутреннем качестве программного обеспечения.
    Программное обеспечение с хорошим внутренним проектом позволяет легко найти, какие нужно внести изменения, чтобы добавить новую функциональную возможность, и где.
    Хорошая модульность позволяет понять только небольшое подмножество кода, в которое нужно вносить изменения.
    Если код понятен, меньше вероятность внести ошибку, а если это и произойдет, процесс отладки будет намного проще.
    Так кодовая база превращается в платформу для создания новых функциональных возможностей для своей предметной области.

    Я называю этот эффект гипотезой стойкости проекта (`Design Stamina Hypothesis <https://martinfowler.com/bliki/DesignStaminaHypothesis.html>`__):
    создавая хороший внутренний проект, мы повышаем стойкость программного обеспечения, позволяющую двигаться быстрее.
    Я не могу доказать, что это так, поэтому называю это утверждение гипотезой.
    Но так подсказывает мой опыт, а также опыт сотен отличных программистов, с которыми я познакомился за свою карьеру.

    Двадцать лет назад общепринятым было мнение, что для создания хорошего проекта нужно завершить проектирование до начала кодирования, потому что, как только мы написали код, мы можем столкнуться только с ухудшением и упадком.
    Рефакторинг меняет эту картину.
    Теперь мы знаем, что можем улучшить проект существующего кода, так что мы можем формировать и улучшать проект с течением времени, даже когда меняются потребности программы.
    Поскольку очень сложно сделать хороший проект заранее, рефакторинг становится жизненно важным.

    Refactoring Helps Me Program Faster

    In the end, all the earlier points come down to this: Refactoring helps me develop code more quickly.

    This sounds counterintuitive.
    When I talk about refactoring, people can easily see that it improves quality.
    Better internal design, readability, reducing bugs—all theseimprove quality.
    But doesn't the time I spend on refactoring reduce the speed of development?

    When I talk to software developers who have been working on a system for a while, I often hear that they were able to make progress rapidly at first, but now it takes much longer to add new features.
    Every new feature requires more and more time to understand how to fit it into the existing code base, and once it's added, bugs often crop up that take even longer to fix.
    The code base starts looking like a series of patches covering patches, and it takes an exercise in archaeology to figure out how things work.
    This burden slows down adding new features — to the point that developers wish they could start again from a blank slate.

    I can visualize this state of affairs with :ref:`the following pseudograph <emacsway-design-stamina-graph>`.

    But some teams report a different experience.
    They find they can add new features faster because they can leverage the existing things by quickly building on what's already there.

    The difference between these two is the internal quality of the software.
    Software with a good internal design allows me to easily find how and where I need to make changes to add a new feature.
    Good modularity allows me to only have to understand a small subset of the code base to make a change.
    If the code is clear, I'm less likely to introduce a bug, and if I do, the debugging effort is much easier.
    Done well, my code base turns into a platform for building new features for its domain.

    I refer to this effect as the `Design Stamina Hypothesis <https://martinfowler.com/bliki/DesignStaminaHypothesis.html>`__:
    By putting our effort into a good internal design, we increase the stamina of the software effort, allowing us to go faster for longer.
    I can't prove that this is the case, which is why I refer to it as a hypothesis.
    But it explains my experience, together with the experience of hundreds of great programmers that I've got to know over my career.

    Twenty years ago, the conventional wisdom was that to get this kind of good design, it had to be completed before starting to program — because once we wrote the code, we could only face decay.
    Refactoring changes this picture.
    We now know we can improve the design of existing code—so we can form and improve a design over time, even as the needs of the program change.
    Since it is very difficult to do a good design up front, refactoring becomes vital to achieving that virtuous path of rapid functionality."

    -- "Refactoring: Improving the Design of Existing Code" 2nd edition by Martin Fowler, Kent Beck, перевод И.В. Красикова под редакцией С.Н. Тригуб

..

    📝 "In its common usage, evolutionary design is a disaster.
    The design ends up being the aggregation of a bunch of ad-hoc tactical decisions, each of which makes the code harder to alter.
    In many ways you might argue this is no design, certainly it usually leads to a poor design.
    As Kent puts it, **design is there to enable you to keep changing the software easily in the long term.**
    **As design deteriorates, so does your ability to make changes effectively.**
    You have the state of software entropy, over time the design gets worse and worse.
    Not only does this make the software harder to change, it also makes bugs both easier to breed and harder to find and safely kill.
    This is the "code and fix" nightmare, where the bugs become exponentially more expensive to fix as the project goes on."

    -- "`Is Design Dead? <https://martinfowler.com/articles/designDead.html>`__" by Martin Fowler

..

    📝 "If you're a manager or customer how can you tell if the software is well designed?
    It matters to you because poorly designed software will be more expensive to modify in the future."

    -- "`Is Design Dead? <https://martinfowler.com/articles/designDead.html>`__" by Martin Fowler

..

    📝 "From the very earliest days of agile methods, people have asked what role there is for architectural or design thinking.
    A common misconception is that since agile methods drop the notion of a detailed up-front design artifact, that there is no room for architecture in an agile project.
    In my keynote at the first-ever agile conference, I pointed out that design was every bit as important for agile projects, but it manifests itself differently, becoming an evolutionary approach."

    -- "`Agile Software Development <https://martinfowler.com/agile.html>`__" by Martin Fowler


Kent Beck
---------

    📝 "Nothing kills speed more effectively than poor internal quality."

    -- "Planning Extreme Programming" by Kent Beck, Martin Fowler

..

    📝 "... the activity of design is not an option. It must be given serious thought for software development to be effective."

    -- "Extreme Programming Explained" by Kent Beck

..

    📝 "Качество — это еще одна весьма странная переменная.
    Зачастую, настаивая на улучшении качества, мы можете завершить проект быстрее, чем запланировано.
    Или вы можете успеть сделать больше за заданный интервал времени.
    Именно это случилось со мной, когда я приступил к разработке тестов для программного модуля, работа над которым описывалась в главе 2.
    Как только я закончил работу над всеми тестами, я был настолько уверен в своем коде, что смог разработать код модуля существенно быстрее, без каких-либо липших сомнений и размышлений.
    Я смог подчистить мою систему с меньшим количеством усилий, в результате я существенно упростил дальнейшую разработку.
    Мне часто приходится наблюдать, как подобное происходит с целыми командами разработчиков.
    Как только они приступают к тестированию или как только они разрабатывают общие для всех стандарты кодирования, работа начинает идти существенно быстрее.

    Существует весьма странная зависимость между внутренним и внешним качеством.
    Внешнее качество — это качество, измерением которого занимается заказчик.
    Внутреннее качество оценивается программистами.
    Если вы намерены временно пожертвовать внутренним качеством для того, чтобы сократить время разработки, и при этом надеетесь на то, что внешнее качество не пострадает слишком сильно, имейте в виду, что вы стремитесь к достижению краткосрочной цели.
    Возможно, закрыв глаза на качество внутренней отделки, вам удастся сэкономить пару недель или даже месяц, однако с течением времени количество внутренних проблем может увеличиться настолько, что разрабатываемую вами систему будет чрезвычайно сложно сопровождать и развивать;
    кроме того, возможно, вам не удастся достичь приемлемого уровня внешнего качества.

    Quality is another strange variable.
    Often, by insisting on better quality you can get projects done sooner, or you can get more done in a given amount of time.
    This happened to me when I started writing unit tests (as described in Chapter 2, A Development Episode, page 7).
    As soon as I had my tests, I had so much more confidence in my code that I wrote faster, without stress.
    I could clean up my system more easily, which made further development easier.
    I've also seen this happen with teams.
    As soon as they start testing, or as soon as they agree on coding standards, they start going faster.

    There is a strange relationship between internal and external quality.
    External quality is quality as measured by the customer.
    Internal quality is quality as measured by the programmers.
    Temporarily sacrificing internal quality to reduce time to market in hopes that external quality won't suffer too much is a tempting short-term play.
    And you can often get away with making a mess for a matter of weeks or months.
    Eventually, though, internal quality problems will catch up with you and make your software prohibitively expensive to maintain, or unable to reach a competitive level of external quality."

    -- "Extreme Programming Explained" 1st edition by Kent Beck, "Chapter 4. Four Variables :: Interactions Between the Variables", перевод ООО Издательство "Питер"

..

    📝 "Why can't you just listen, write a test case, make it run, listen, write a test case, make it run indefinitely?
    Because we know it doesn't work that way.
    You can do that for a while.
    In a forgiving language you may even be able to do that for a long while.
    Eventually, though, you get stuck.
    The only way to make the next test case run is to break another.
    Or the only way to make the test case run is far more trouble than it is worth.
    Entropy claims another victim.

    The only way to avoid this is to design.
    Designing is creating a structure that organizes the logic in the system.
    Good design organizes the logic so that a change 45 in one part of the system doesn't always require a change in another part of the system.
    Good design ensures that every piece of logic in the system has one and only one home.
    Good design puts the logic near the data it operates allows the extension of the system with changes in only one place."

    -- "Extreme Programming Explained" by Kent Beck


Robert Martin
-------------

    💬 "The only way to make the deadline -- the only way to go fast -- is to keep the code as clean as possible at all times."

    -- "Clean Code: A Handbook of Agile Software Craftsmanship" by Robert C. Martin

..

    📝 "The way to go fast, and to keep the deadlines at bay, is to stay clean.
    Professionals do not succumb to the temptation to create a mess in order to move quickly.
    Professionals realize that "quick and dirty" is an oxymoron.
    Dirty always means slow!"

    -- "Clean Coder" by Robert Martin

..

    📝 "The goal of good software design? That goal is nothing less than my utopian description:

        The goal of software architecture is to minimize the human resources required to build and maintain the required system.

    The measure of design quality is simply the measure of the effort required to meet the needs of the customer.
    If that effort is low, and stays low throughout the lifetime of the system, the design is good.
    If that effort grows with each new release, the design is bad.
    It's as simple as that."

    -- "Clean Architecture: A Craftsman's Guide to Software Structure and Design" by Robert C. Martin

..

    📝 "Напомню, что целью архитектора является минимизация трудозатрат на создание и сопровождение системы.
    Что может помешать достижению этой цели?
    Зависимость — и особенно зависимость от преждевременных решений.

    Recall that the goal of an architect is to minimize the human resources required to build and maintain the required system.
    What it is that saps this kind of peoplepower?
    Coupling—and especially coupling to premature decisions."

    -- "Clean Architecture: A Craftsman's Guide to Software Structure and Design" by Robert C. Martin, перевод ООО Издательство "Питер"


Agile Manifesto
---------------

    📝 "Continuous attention to technical excellence and good design enhances agility."

    -- "`Principles behind the Agile Manifesto <http://agilemanifesto.org/principles.html>`__"


Ralph Johnson
-------------

    📝 "In most successful software projects, the expert developers working on that project have
    a shared understanding of the system design.
    **This shared understanding is called 'architecture.'**
    This understanding includes how the system is divided into components and how the components interact through interfaces.
    These components are usually composed of smaller components, but the architecture only
    includes the components and interfaces that are understood by all the developers."

    -- `Ralph Johnson <https://martinfowler.com/ieeeSoftware/whoNeedsArchitect.pdf>`__


Steve McConnell
---------------

    📝 "The General Principle of Software Quality is that improving quality reduces development costs.

    Understanding this principle depends on understanding a key observation: the best way
    to improve productivity and quality is to reduce the time spent reworking code, whether
    the rework arises from changes in requirements, changes in design, or debugging.
    The industry-average productivity for a software product is about 10 to 50 of lines of
    delivered code per person per day (including all noncoding overhead).
    It takes only a matter of minutes to type in 10 to 50 lines of code, so how is the rest of the day spent?
    Part of the reason for these seemingly low productivity figures is that industry average
    numbers like these factor nonprogrammer time into the lines-of-code-per-day figure.
    Tester time, project manager time, and administrative support time are all included.
    Noncoding activities, such as requirements development and architecture work, are also
    typically factored into those lines-of-code-per-day figures.
    But none of that is what takes up so much time.

    The single biggest activity on most projects is debugging and correcting code that
    doesn't work properly.
    Debugging and associated refactoring and other rework consume
    about 50 percent of the time on a traditional, naive software-development cycle.
    (See Section 3.1, "Importance of Prerequisites," for more details.) Reducing debugging by
    preventing errors improves productivity.
    Therefore, the most obvious method of shortening a development schedule is to improve the quality of the product and decrease
    the amount of time spent debugging and reworking the software.
    This analysis is confirmed by field data.
    In a review of 50 development projects involving over 400 work-years of effort and
    almost 3 million lines of code, a study at NASA's Software
    Engineering Laboratory found that increased quality assurance was
    associated with decreased error rate but did not increase overalldevelopment cost (Card 1987).

    A study at IBM produced similar findings:

        Software projects with the lowest levels of defects had the shortest development
        schedules and the highest development productivity.... software defect removal is
        actually the most expensive and time-consuming form of work for software (Jones 2000).

        -- Jones, Capers. 2000. Software Assessments, Benchmarks, and Best Practices. Reading, MA: Addison-Wesley.

    The same effect holds true at the small end of the scale.
    In a 1985 study, 166 professional programmers wrote programs from the
    same specification.
    The resulting programs averaged 220 lines of
    code and a little under five hours to write.
    The fascinating result was that programmers who took the median time to complete their
    programs produced programs with the greatest number of errors.
    The programmers who took more or less than the median time
    produced programs with significantly fewer errors (DeMarco and Lister 1985).

    The two slowest groups took about five times as long to achieve roughly the same
    defect rate as the fastest group.
    It's not necessarily the case that writing software without
    defects takes more time than writing software with defects.
    As the graph shows, it can take less."

    -- "Code Complete" 2nd edition by Steve McConnell

..

    📝 "Watts Humphrey reports that teams using the Team Software Process
    (TSP) have achieved defect levels of about 0.06 defects per 1000 lines of code.
    TSP focuses on training developers not to create defects in the first place (Weber 2003).
    [Morales, Alexandra Weber. 2003. "The Consummate Coach: Watts Humphrey, Father of Cmm and Author of Winning with Software, Explains How to Get Better at What You Do," SD Show Daily, September 16, 2003.]

    The results of the TSP and cleanroom projects confirm another version of the General
    Principle of Software Quality: it's cheaper to build high-quality software than it is to build and fix low-quality software.
    Productivity for a fully checked-out, 80,000-line cleanroom project was 740 lines of code per work-month.
    The industry average rate for fully checked-out code is closer to 250–300 lines per work-month, including all noncoding overhead (Cusumano et al 2003).
    [Cusumano, Michael , et al. 2003. "Software Development Worldwide: The State of the Practice," IEEE Software, November/ December 2003, 28–34.]
    The cost savings and productivity come from the fact that virtually no time is devoted to debugging on TSP or cleanroom projects.
    No time spent on debugging?
    That is truly a worthy goal!"

    -- "Code Complete" 2nd edition by Steve McConnell

..

    📝 "A six-month study conducted by IBM found that maintenance programmers "most often said that **understanding the original programmer's intent was the most difficult problem**" (Fjelstad and Hamlen 1979).
    [Fjelstad, R. K. , and W. T. Hamlen. 1979. "Applications Program Maintenance Study: Report to our Respondents." Proceedings Guide 48, Philadelphia. Reprinted in Tutorial on Software Maintenance, G. Parikh and N. Zvegintzov eds. Los Alamitos, CA: CS Press, 1983: 13–27.]"

    -- "Code Complete" 2nd edition by Steve McConnell


Сергей Тепляков
---------------

    📝 "Хороший дизайн заключается в простом решении, когда изменения требований ведут к линейным трудозатратам."

    -- "`Принцип YAGNI <http://sergeyteplyakov.blogspot.com/2016/08/yagni.html>`__", Сергей Тепляков


Народное творчество
-------------------

Старый программистский анекдот:

    💬 Идет мужик по лесу. Смотрит, другой мужик лес рубит.
    | \- Привет, что делаешь?
    | \- Не видишь? Лес рублю...
    | \- Так бензопила же лежит рядом. Возьми её - быстрее будет.
    | \- Я не умею.
    | \- Так инструкция же рядом лежит. Возьми, прочти...
    | \- Мне некогда её читать - мне лес рубить надо.


Randy Shoup
-----------

    | \- We don't have time to do it right!
    | \- Do you have time to do it twice?

    -- `Randy Shoup <https://www.infoq.com/presentations/microservices-data-centric>`_, VP Engineering at Stitch Fix in San Francisco

.. figure:: _media/software-design/do-it-right.png
   :alt: Do it right! Иллюстрация из открытых источников неизвестного автора.
   :align: left
   :width: 90%

   Do it right! Иллюстрация из открытых источников неизвестного автора.

.. seealso::

   - ":doc:`../crash-course-in-software-development-economics`"
   - ":ref:`emacsway-icebreaker-principle`"
   - ":ref:`emacsway-adaptation`"
   - ":ref:`emacsway-agile-development`"
   - ":ref:`emacsway-agile-patterns`"
