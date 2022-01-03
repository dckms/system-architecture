:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: Team Topologies; at scale
   :name: emacsway-team-topologies-at-scale

==============================================================
Роль архитектуры в масштабировании команд, DDD и микросервисах
==============================================================

.. sectionauthor:: Ivan Zakrevsky


.. contents:: Содержание


Закон Брукса и роль автономности команд
=======================================


Закон Брукса
------------

Сегодня, наверное, каждый знает Закон Брукса:

    📝 "Если проект не укладывается в сроки, то добавление рабочей силы задержит его ещё больше.

    Adding manpower to a late software project makes it later."

    -- The Brooks's Law

    📝 "Brooks' law is based on the idea that communications overhead is a significant factor on software projects, and that work on a software project is not easily partitioned into isolated, independent tasks. Ten people can pick cotton ten times as fast as one person because the work is almost perfectly partitionable, requiring little communication or coordination. But nine women can't have a baby any faster than one woman can because the work is not partitionable. Brooks argues that work on a software project is more like having a baby than picking cotton. When new staff are brought into a late project, they aren't immediately productive, and they must be trained. The staff who must train them are already productive, but they lose productivity while they're training new staff. Brooks argues that, on balance, more effort is lost to training and additional coordination and communications overhead than is gained when the new staff eventually becomes productive."

    -- "`Brooks' Law Repealed? <https://stevemcconnell.com/articles/brooks-law-repealed/>`__" by Steve McConnell

..

    📝 "Число занятых [специалистов] и число месяцев [в термине человеко-месяц] являются взаимозаменяемыми величинами лишь тогда, когда задачу можно распределить среди ряда работников, которые не имеют между собой взаимосвязи (рис. 2.1).

    Men and months [in term man-month] are interchangeable commodities only when a task can be partitioned among many workers with no communication among them (Fig. 2.1)."

    -- "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., перевод С. Маккавеева

Сравните это с

    📝 "Microservices\' main benefit, in my view, is enabling parallel development by establishing a hard-to-cross boundary between different parts of your system."

    -- "`Don't start with a monolith <https://martinfowler.com/articles/dont-start-monolith.html>`__" by Stefan Tilkov, a co-founder and principal consultant at innoQ


Противоречие между эффективностью и производительностью
-------------------------------------------------------

Frederick Brooks формулирует противоречие. С одной стороны:

    📝 "Выше я доказал, что само число разработчиков, действия которых нужно согласовывать, оказывает влияние на стоимость проекта, поскольку значительная часть издержек вызвана необходимостью общения и устранения отрицательных последствий разобщенности (системная отладка). Это также наводит на мысль, что **желательно разрабатывать системы возможно меньшим числом людей**. Действительно, опыт разработки больших программных систем, как правило, показывает, что подход с позиций грубой силы влечет удорожание, замедленность, неэффективность, а создаваемые в результате системы не являются концептуально целостными. Список, иллюстрирующий это, бесконечен: OS/360, Exec 8, Scop 6600, Multics, TSS, SAGE и другие.

    I have earlier argued that the sheer number of minds to be coordinated affects the cost of the effort, for a major part of the cost is communication and correcting the ill effects of miscommunication (system debugging). This, too, suggests that one wants the system to be built by **as few minds as possible**. Indeed, most experience with large programming systems shows that the brute-force approach is costly, slow, inefficient, and produces systems that are not conceptually integrated. OS/360, Exec 8, Scope 6600, Multics, TSS, SAGE, etc.—the list goes on and on."

    -- "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., перевод С. Маккавеева

Таким же наблюдением делится и Sam Newman:

    📝 "... небольшие команды, работающие с небольшим объемом исходного кода, как правило, показывают более высокую продуктивность.

    ... smaller teams working on smaller codebases tend to be more productive!"

    -- "Building Microservices. Designing Fine-Grained Systems" by Sam Newman

Но, с другой стороны:

    📝 "В этом и состоит **изъян идеи маленькой активной команды: для создания по-настоящему крупных систем ей потребуется слишком много времени**. Посмотрим, как разработка OS/360 осуществлялась бы маленькой активной командой, допустим, из 10 человек. Положим, что они в семь раз более продуктивны средних программистов (что далеко от истины). Допустим, что уменьшение объема общения благодаря малочисленности команды позволило еще в семь раз повысить производительность. Допустим, что на протяжении всего проекта работает одна и та же команда. Таким образом, 5000/(10*7*7)=10, т.е. работу в 5000 человеко-лет они выполнят за 10 лет. Будет ли продукт представлять интерес через 10 лет после начала разработки или устареет благодаря стремительному развитию программных технологий?

    This then is **the problem with the small, sharp team concept: it is too slow for really big systems. Consider the OS/360 job as it might be tackled with a small, sharp team**. Postulate a 10-man team. As a bound, let them be seven times as productive as mediocre programmers in both programming and documentation, because they are sharp. Assume OS/360 was built only by mediocre programmers (which is far from the truth). As a bound, assume that another productivity improvement factor of seven comes from reduced communication on the part of the smaller team. Assume the same team stays on the entire job. Well, 5000/(10 X 7 X 7 ) = 10; they can do the 5000 man-year job in 10 years. Will the product be interesting 10 years after its initial design? Or will it have been made obsolete by the rapidly developing software technology?"

    -- "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., перевод С. Маккавеева

Возникает противоречие:

    📝 "Дилемма представляется жестокой. Для эффективности и концептуальной целостности предпочтительнее, чтобы проектирование и создание системы осуществили **несколько светлых голов**. Однако для больших систем желательно поставить под ружье **значительный контингент**, чтобы продукт мог увидеть свет вовремя. **Как можно примирить эти два желания?**

    The dilemma is a cruel one. For efficiency and conceptual integrity, one prefers **a few good minds** doing design and construction. Yet for large systems one wants a way to bring **considerable manpower to bear**, so that the product can make a timely appearance. **How can these two needs be reconciled?**"

    -- "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., перевод С. Маккавеева

Добавим немного контекста:

    📝 "Вторая ошибка рассуждений заключена в самой единице измерения, используемой при оценивании и планировании: человеко-месяц. Стоимость действительно измеряется как произведения числа занятых на количество затраченных месяцев. Но не достигнутый результат. Поэтому использование человеко-месяца как единицы измерения объема работы является опасным заблуждением.

    Число занятых и число месяцев являются взаимозаменяемыми величинами лишь тогда, когда задачу можно распределить среди ряда работников, которые не имеют между собой взаимосвязи (рис. 2.1). Это верно, когда жнут пшеницу или собирают хлопок, но в системном программировании это далеко не так.

    .. figure:: _media/harlan-mills'-proposal/fig-2.1-perfectly-partitionable-task.png
       :alt: Рис. 2.1 Зависимость времени от числа занятых — полностью разделимая задача. Fig. 2.1 Time versus number of workers—perfectly partitionable task. The image source is "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., "Chapter 2 The Mythical Man-Month", перевод С. Маккавеева.
       :align: left
       :width: 70%

       Рис. 2.1 Зависимость времени от числа занятых — полностью разделимая задача. Fig. 2.1 Time versus number of workers—perfectly partitionable task. The image source is "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., "Chapter 2 The Mythical Man-Month", перевод С. Маккавеева.

    Если задачу нельзя разбить на части, поскольку существуют ограничения на последовательность выполнения этапов, то увеличение затрат не оказывает влияния на график (рис. 2.2). Чтобы родить ребенка требуется девять месяцев независимо от того, сколько женщин привлечено к решению данной задачи. Многие задачи программирования относятся к этому типу, поскольку отладка по своей сути носит последовательный характер.

    .. figure:: _media/harlan-mills'-proposal/fig-2.2-unpartitionable-task.png
       :alt: Рис. 2.2 Зависимость времени от числа занятых — неразделимая задача. Fig. 2.2 Time versus number of workers—unpartitionable task. The image source is "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., "Chapter 2 The Mythical Man-Month", перевод С. Маккавеева.
       :align: left
       :width: 70%

       Рис. 2.2 Зависимость времени от числа занятых — неразделимая задача. Fig. 2.2 Time versus number of workers—unpartitionable task. The image source is "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., "Chapter 2 The Mythical Man-Month", перевод С. Маккавеева.

    Для задач, которые могут быть разбиты на части, но требуют обмена данными между подзадачами, затраты на обмен данными должны быть добавлены к общему объему необходимых работ. Поэтому достижимый наилучший результат оказывается несколько хуже, чем простое соответствие числа занятых и количества месяцев (рис. 2.3).

    .. figure:: _media/harlan-mills'-proposal/fig-2.3-partitionable-task-requiring-communication.png
       :alt: Рис. 2.3 Зависимость времени от числа занятых — разделимая задача, требующая обмена данными. Fig. 2.3 Time versus number of workers—partitionable task requiring communication. The image source is "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., "Chapter 2 The Mythical Man-Month", перевод С. Маккавеева.
       :align: left
       :width: 70%

       Рис. 2.3 Зависимость времени от числа занятых — разделимая задача, требующая обмена данными. Fig. 2.3 Time versus number of workers—partitionable task requiring communication. The image source is "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., "Chapter 2 The Mythical Man-Month", перевод С. Маккавеева.

    Дополнительная нагрузка состоит из двух частей — обучения и обмена данными. Каждого работника нужно обучить технологии, целям проекта, общей стратегии и плану работы. Это обучение нельзя разбить на части, поэтому данная часть затрат изменяется линейно в зависимости от числа занятых.

    **С обменом данными дело обстоит хуже. Если все части задания должны быть отдельно скоординированы между собой, то затраты возрастают как n(n-2)/2. Для трех работников требуется втрое больше попарного общения, чем для двух, для четырех — вшестеро. Если помимо этого возникает необходимость в совещаниях трех, четырех и т.д. работников для совместного решения вопросов, положение становится еще хуже. Дополнительные затраты на обмен данными могут полностью обесценить результат дробления исходной задачи и привести к положению, описываемому рисунком 2.4.**

    .. figure:: _media/harlan-mills'-proposal/fig-2.4-task-with-complex-interrelationships.png
       :alt: Рис. 2.4 Зависимость времени от числа занятых — задача со сложными взаимосвязями. Fig. 2.4 Time versus number of workers—task with complex interrelationships. The image source is "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., "Chapter 2 The Mythical Man-Month", перевод С. Маккавеева.
       :align: left
       :width: 70%

       Рис. 2.4 Зависимость времени от числа занятых — задача со сложными взаимосвязями. Fig. 2.4 Time versus number of workers—task with complex interrelationships. The image source is "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., "Chapter 2 The Mythical Man-Month", перевод С. Маккавеева.

    Поскольку создание программного продукта является по сути системным проектом — практикой сложных взаимосвязей, затраты на обмен данными велики и быстро начинают преобладать над сокращением сроков, достигаемым в результате разбиения задачи на более мелкие подзадачи. В этом случае привлечение дополнительных работников не сокращает, а удлиняет график работ.

    The second fallacious thought mode is expressed in the very unit of effort used in estimating and scheduling: the man-month. Cost does indeed vary as the product of the number of men and the number of months. Progress does not. Hence the man-month as a unit for measuring the size of a job is a dangerous and deceptive myth. It implies that men and months are interchangeable.

    Men and months are interchangeable commodities only when a task can be partitioned among many workers with no communication among them (Fig. 2.1). This is true of reaping wheat or picking cotton; it is not even approximately true of systems programming.

    When a task cannot be partitioned because of sequential constraints, the application of more effort has no effect on the schedule (Fig. 2.2). The bearing of a child takes nine months, no matter how many women are assigned. Many software tasks have this characteristic because of the sequential nature of debugging.

    In tasks that can be partitioned but which require communication among the subtasks, the effort of communication must be added to the amount of work to be done. Therefore the best that can be done is somewhat poorer than an even trade of men for months (Fig. 2.3).

    The added burden of communication is made up of two parts, training and intercommunication. Each worker must be trained in the technology, the goals of the effort, the overall strategy, and the plan of work. This training cannot be partitioned, so this part of the added effort varies linearly with the number of workers.

    **Intercommunication is worse. If each part of the task must be separately coordinated with each other part/ the effort increases as n(n-I)/2. Three workers require three times as much pairwise intercommunication as two; four require six times as much as two. If, moreover, there need to be conferences among three, four, etc., workers to resolve things jointly, matters get worse yet. The added effort of communicating may fully counteract the division of the original task and bring us to the situation of Fig. 2.4.**

    Since software construction is inherently a systems effort—an exercise in complex interrelationships—communication effort is great, and it quickly dominates the decrease in individual task time brought about by partitioning. Adding more men then lengthens, not shortens, the schedule."

    -- "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., перевод С. Маккавеева

Мы остановились на дилемме: с одной стороны, чем меньше численность людей, принимающих проектные решения, тем выше их продуктивность.
С другой стороны, чем больше людей задействовано в разработке, тем скорее продукт сможет выйти на рынок.
Проанализируем о том, как эту дилемму можно решить.

Решение этой дилеммы становится возможным с качественным отделением архитектуры от реализации (с чем отлично справляются сетевые границы Bounded Contexts):

    📝 "**Архитектура и разработка должны быть тщательно разделены.** Как сказал Блау (Blaauw), «архитектура говорит, что должно произойти, а разработка - как сделать, чтобы это произошло». В качестве простого примера он приводит часы, архитектура которых состоит из циферблата, стрелок и заводной головки. Ребенок, освоивший это архитектуру, с одинаковой легкостью может узнать время и по ручным часам, и по часам на колокольне. Исполнение же и его реализация описывают, что происходит внутри: передача усилий и управление точностью каждым из многих механизмов.

    **Architecture must be carefully distinguished from implementation.** As Blaauw has said, "Where architecture tells what happens, implementation tells how it is made to happen." He gives as a simple example a clock, whose architecture consists of the face, the hands, and the winding knob. When a child has learned this architecture, he can tell time as easily from a wristwatch as from a church tower. The implementation, however, and its realization, describe what goes on inside the case—powering by any of many mechanisms and accuracy control by any of many."

    -- "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., перевод С. Маккавеева


.. index:: Team Topologies
   :name: emacsway-harlan-mills'-proposal

Предложение Харлана Миллза
--------------------------

Harlan Mills' Proposal (Предложение Харлана Миллза):

    📝 "Предложение Харлана Миллза дает свежее и творческое решение. Миллз предложил, чтобы на каждом участке работы была **команда разработчиков, организованная наподобие бригады хирургов, а не мясников**. Имеется в виду, что не каждый участник группы будет врезаться в задачу, но резать будет один, а остальные оказывать ему всевозможную поддержку, повышая его производительность и плодотворность.

    При некотором размышлении ясно, что **эта идея приведет к желаемому, если ее удастся осуществить**. Лишь несколько голов занято проектированием и разработкой, и в то же время много работников находится на подхвате. Будет ли такая организация работать? Кто играет роль анестезиологов и операционных сестер в группе программистов, а как осуществляется разделение труда? Чтобы нарисовать картину работы такой команды с включением всех мыслимых видов поддержки, я позволю себе вольное обращение к метафорам.

    A proposal by Harlan Mills offers a fresh and creative solution. Mills proposes that each segment of a large job be tackled by a team, but that **the team be organized like a surgical team rather than a hog-butchering team**. That is, instead of each member cutting away on the problem, one does the cutting and the others give him every support that will enhance his effectiveness and productivity.

    A little thought shows that **this concept meets the desiderata, if it can be made to work**. Few minds are involved in design and construction, yet many hands are brought to bear. Can it work? Who are the anesthesiologists and nurses on a programming team, and how is the work divided? Let me freely mix metaphors to suggest how such a team might work if enlarged to include all conceivable support."

    -- "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., перевод С. Маккавеева

Обратите внимание на слова "эта идея приведет к желаемому, если ее удастся осуществить".
Именно эту задачу успешно решают Bounded Contexts, позволяя совместить большой размер команды и продуктивность, т.е., осуществить масштабирование команды без ущерба производительности.

    📝 "... мы стремимся к тому, чтобы сервисы создавались в результате разбиения системы на такие части, при которых **темпы развития внутри сервисов были намного выше темпов развития между сервисами**.

    ... we aim to ensure our services are decomposed such that the **pace of change inside a service is much higher than the pace of change between services**."

    -- "Building Microservices. Designing Fine-Grained Systems" by Sam Newman

..

    📝 "One case study was particularly interesting. The team had made the wrong choice, using microservices on a system that wasn't complex enough to cover the Microservice Premium. The project got in trouble and needed to be rescued, so lots more people were thrown onto the project. At this point the microservice architecture became helpful, because **the system was able to absorb the rapid influx of developers and the team was able to leverage the larger team numbers much more easily than is typical with a monolith**. As a result the project accelerated to a productivity greater than would have been expected with a monolith, enabling the team to catch up. The result was still a net negative, in that the software cost more staff-hours than it would have done if they had gone with a monolith, but the **microservices architecture did support ramp up**."

    -- "`Microservice Trade-Offs <https://martinfowler.com/articles/microservice-trade-offs.html>`__" by Martin Fowler

..

    📝 "Netflix и Amazon

    Наверное, идея обязательной согласованности организации и архитектуры может быть неплохо проиллюстрирована на примере Amazon и Netflix. В Amazon довольно рано начали понимать преимущества владения командами полным жизненным циклом управляемых ими систем. Там решили, что команды должны всецело распоряжаться теми системами, за которые они отвечают, управляя всем жизненным циклом этих систем. Но в Amazon также знали, что **небольшие команды могут работать быстрее больших**. Это привело к созданию команд, которые можно было бы накормить двумя пиццами. Это стремление к созданию небольших команд, владеющих полным жизненным циклом своих сервисов, и стало основной причиной того, что в Amazon была разработана платформа Amazon Web Services. Для обеспечения самодостаточности своих команд компании понадобилось создать соответствующий инструментарий. Этот пример был взят на вооружение компанией Netflix и с самого начала определил формирование ее структуры вокруг небольших независимых команд, образуемых с прицелом на то, что создаваемые ими сервисы также будут независимы друг от друга. Тем самым обеспечивалась оптимизация скорости изменения архитектуры систем. Фактически в Netflix разработали организационную структуру для желаемой архитектуры создаваемых систем.

    Netflix and Amazon

    Probably the two poster children for the idea that organizations and architecture should be aligned are Amazon and Netflix. Early on, Amazon started to understand the benefits of teams owning the whole lifecycle of the systems they managed. It wanted teams to own and operate the systems they looked after, managing the entire lifecycle. But Amazon also knew that **small teams can work faster than large teams**. This led famously to its two-pizza teams, where no team should be so big that it could not be fed with two pizzas. This driver for small teams owning the whole lifecycle of their services is a major reason why Amazon developed Amazon Web Services. It needed to create the tooling to allow its teams to be self-sufficient. Netflix learned from this example, and ensured that from the beginning it structured itself around small, independent teams, so that the services they created would also be independent from each other. This ensured that the architecture of the system was optimized for speed of change. Effectively, Netflix designed the organizational structure for the system architecture it wanted."

    -- "Building Microservices. Designing Fine-Grained Systems" by Sam Newman

я хотел бы добавить еще одно высказывание от разработчиков популярного микросервисного фреймворка go-kit:

    📝 "Почему микросервисы? Практически вся современная разработка программного обеспечения ориентирована на единственную цель - улучшить время выхода на рынок. Микросервисы представляют собой эволюцию модели сервис-ориентированной архитектуры, которая элегантно устраняет организационные противоречия, **предоставляя вашим инженерам и командам автономию**, необходимую им для непрерывной доставки, итерации и улучшения.

    Why microservices? Almost all of contemporary software engineering is focused on the singular goal of improving time-to-market. Microservices are an evolution of the service-oriented architecture pattern that elegantly eliminate organizational friction, **giving your engineers and teams the autonomy** they need to continuously ship, iterate, and improve."

    -- "`Go kit. A toolkit for microservices. <https://gokit.io/>`__"

На самом деле, если у вас армейская дисциплина среди разработчиков, то нет необходимости укреплять сетевыми границами пределы автономности команд, поскольку границами автономности команды является Bounded Context, который не обязательно должен быть выражен микорсервисом(-ами).

    📝 "In theory, you don’t need microservices for this if you simply have the discipline to follow clear rules and establish clear boundaries within your monolithic application; in practice, I’ve found this to be the case only very rarely.)"

    -- "Don’t start with a monolith" by Stefan Tilkov
    https://martinfowler.com/articles/dont-start-monolith.html

Сетевые границы решают проблему, известную как creeping coupling или abstraction leak. А это позволяет снизить квалификационные требования к разработчикам, тем более, что Microservices First обычно имеет экономическую целесообразность только при задействовании большого количества разработчиков.

    📝 "Обмен данными между самими сервисами ведется через сетевые вызовы, чтобы **упрочить обособленность сервисов** и избежать рисков, сопряженных с тесными связями.

    All communication between the services themselves are via network calls, **to enforce separation between the services** and avoid the perils of tight coupling."

    -- "Building Microservices. Designing Fine-Grained Systems" by Sam Newman


SAFe:

    📝 "The second dimension of the team and technical agility competency is teams of Agile teams. Even with good, local execution, building enterprise-class solutions typically requires more scope and breadth of skills than a single Agile team can provide. Therefore, Agile teams operate in the context of an ART, which is a long-lived team of Agile teams. The ART incrementally develops, delivers, and (where applicable) operates one or more solutions (Figure 6-5)."

    📝 "System Architect/Engineering is an individual or team that defines the overall architecture of the system. They work at a level of abstraction above the teams and components and define Non-Functional Requirements (NFRs), major system elements, subsystems, and interfaces."

    📝 "System Teams typically assist in building and supporting DevOps infrastructure for development, continuous integration, automated testing, and deployment into the staging environment. In larger systems they may do end-to-end testing, which cannot be readily accomplished by individual Agile teams."

    📝 "Shared Services are specialists—for example, data security, information architects, Database Administrators (DBAs)—who are necessary for the success of an ART but cannot be dedicated to a specific train."

    📝 "With the right architecture, elements of the system may be released independently. Figure 8-8 illustrates an autonomous delivery system that was architected to enable system elements to be released independently."

    📝 "Figure 8-8. Architecture impacts the ability to release system elements independently"

    -- "SAFe® 5.0: The World’s Leading Framework for Business Agility" by Richard Knaster, Dean Leffingwell


.. index::
   single: Literature; in Team Topologies at scale
   :name: emacsway-team-topologies-at-scale-literature

Литература
==========

- "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr.
- "Team Topologies: Organizing Business and Technology Teams for Fast Flow" by Matthew Skelton
- "`Architecture Modernization with Strategic Domain-Driven Design. A Guide for Technology Leaders. <https://leanpub.com/arch-modernization-ddd>`__" by Nick Tune


Ссылки по теме
==============

1. "`Architecture Ownership Patterns For Team Topologies. Part 1: A Business Architecture Model <https://medium.com/nick-tune-tech-strategy-blog/team-responsibility-ownership-patterns-part-1-a-business-architecture-model-63597c4e60e1>`__" by Nick Tune
#. "`Architecture Ownership Patterns for Team Topologies. Part 2: Single Team Patterns <https://medium.com/nick-tune-tech-strategy-blog/architecture-ownership-patterns-for-team-topologies-part-2-single-team-patterns-943d31854285>`__" by Nick Tune
#. "`Architecture Ownership Patterns for Team Topologies. Part 3: Multi-Team Patterns <https://medium.com/nick-tune-tech-strategy-blog/architecture-ownership-patterns-for-team-topologies-part-3-multi-team-patterns-eecc146ddb28>`__" by Nick Tune

- "`About Team Topologies and Context Mapping <https://blog.avanscoperta.it/2021/04/22/about-team-topologies-and-context-mapping/>`__" by Alberto Brandolini

- "`The strong and weak forces of architecture <https://martinfowler.com/articles/strong-weak-arch.html>`__" by Evan Bottcher
- "`Compliance in a DevOps Culture. Integrating Compliance Controls and Audit into CI/CD Processes <https://martinfowler.com/articles/devops-compliance.html>`__ by Carl Nygard

- "`How using events helps in a teams' autonomy <https://event-driven.io/en/how_using_events_help_in_teams_autonomy/>`__" by Oskar Dudycz

1. "`Agile Teams <https://www.scaledagileframework.com/agile-teams/>`__"
#. "`Organizing Agile Teams and ARTs: Team Topologies at Scale <https://www.scaledagileframework.com/organizing-agile-teams-and-arts-team-topologies-at-scale/>`__"
#. "`System and Solution Architect/Engineering <https://www.scaledagileframework.com/system-and-solution-architect-engineering/>`__"
#. "`Enterprise Architect <https://www.scaledagileframework.com/enterprise-architect/>`__"
#. "`Architectural Runway <https://www.scaledagileframework.com/architectural-runway/>`__"
#. "`Agile Architecture in SAFe <https://www.scaledagileframework.com/agile-architecture/>`__"

- "`Open Agile Architecture. A Standard of The Open Group <https://pubs.opengroup.org/architecture/o-aa-standard/>`__"

