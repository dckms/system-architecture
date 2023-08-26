:canonical-base-url: https://dckms.github.io/system-architecture

.. index:: Domain Model
   :name: domain-model-definition


=======================
Domain Model Definition
=======================

.. sectionauthor:: Stanislav Bolsun

Один из самых популярных и сложных вопросов - "Что такое доменная модель?", и в этой статье будет попытка объяснить максимально
простым языком что это такое и с чем она связана.

.. contents:: Содержание

Доменная модель
---------------

Чтобы разобраться что такое модель, необходимо так же понять и другие термины, идущие вплотную, ограниченный контекст и единый язык.

Начнем с главной задачи модели - это борьба со сложностью в системе. Доменная модель является упрощенной интерпретацией реальности, используемая для решения бизнес задач или целей.

Каждая модель имеет свой контекст применимости, без контекста применимости мы не сможем создать модель, так как не знаем какую проблему решаем (то есть какие свойства и поведение нужны для конкретной проблемы). Из этого следует, что ограниченный контекст является границей применимости нашей модели.

.. figure:: _media/model_of_earth_processes.png
   :alt: Model of Earth processes
   :align: center
   :width: 100%
На изображении выше, мы видим очень упрощенную модель процессов Земли, эта модель служит одной цели и моделирует объект не полностью, а только в необходимом минимуме для решения поставленной задачи.

На это и делают акцент Эванс, Вернон и Зимарев в определениях модели:

    📝 "every model represents some aspect of reality or an idea that is of interest.
    A model is a simplification. It is an interpretation of reality that abstracts the aspects relevant to solving
    the problem at hand and ignores extraneous detail..."

    -- "DDD" by Eric Evans



    📝 "So, models represent some artifact of the real world, but with a narrow purpose.
    How much space the building will occupy and how high the whole complex will be, for example,
    are often just enough for a rough model, during the first review stage of the building project.
    Models do not intend to replicate real life. Instead, they represent some particular aspects of real life at a certain level of detail,
    depending on the purpose of the model...

    Going back to Chapter 1,  Why Domain-Driven Design?, if the business domain and the particular problems we have to
    solve are in our problem space, the domain model is purely in our solution space.
    We will be modeling our solution, and those models will be our domain models."

    -- "Hands-On Domain-Driven Design with .NET Core: Tackling complexity in the heart of software by putting DDD principles into practice" by Alexey Zimarev



    📝 "What’s a Domain Model?
    It’s a software model of the very specific business domain you are working in. Often it’s implemented as an object model,
    where those objects have both data and behavior with literal and accurate business meaning.
    Creating a unique, carefully crafted domain model at the heart of a core, strategic application or subsystem is essential to
    practicing DDD. With DDD your domain models will tend to be smallish, very focused.
    Using DDD, you never try to model the whole business enterprise with a single, large domain model. Phew, that’s good!""

    -- "IDDD" by Vaughn Vernon


Важное уточнение: Модель - это абстракция, которая формирует реализацию, но не является реализацией,
хотя реализация и может выражать эту модель. Модель - это solution space (см следующую цитату).

    📝 "A domain model is not a particular diagram; it is the idea that the diagram is intended to convey.
    It is not just the knowledge in a domain expert's head;
    it is a rigorously organized and selective abstraction of that knowledge.
    A diagram can represent and communicate a model, as can carefully written code, as can an English sentence...

    The model and the heart of the design shape each other. It is the intimate link between the model and the implementation that makes the model relevant and ensures that the analysis that went into it applies to the final product, a running program.
    This binding of model and implementation also helps during maintenance and continuing development, because the code can be interpreted based on understanding the model. (See Chapter 3.)"

    -- "DDD" by Eric Evans



А что если смоделировать "единую" модель системы?
-------------------------------------------------

Если контекст применимости не ясен или не проработан до должного уровня, то у нас есть два возможных пути: либо модель не создавать вообще, либо создавать модель на все случаи жизни,
но тогда придется полностью воспроизвести объект моделирования (что не позволит эффективно решать бизнес задачи).

    📝 "Because the term domain model includes the word domain, we might get the idea that we should create a single,
    cohesive, all-inclusive model of an organization’s entire business domain—you know, like an enterprise model. However,
    when using DDD, that is not our goal. DDD places emphasis on just the opposite. The whole Domain of the organization is composed of Subdomains.
    Using DDD, models are developed in Bounded Contexts. In fact, developing a Domain Model is actually one way that we focus on only one specific area of the whole business domain.
    Any attempt to define the business of even a moderately complex organization in a single, all-encompassing model will be at best extremely difficult and will usually fail.
    As is made clear in this chapter, vigorously separating distinct areas of the whole business domain will help us succeed.

    So, if a domain model shouldn’t be all-inclusive of what the organization does and how it does it, what should it be, exactly?

    Almost every software Domain has multiple Subdomains. It really doesn’t matter whether the organization is huge and extremely complex or consists of just a few people and the software they use.
    There are different functions that make any business successful, so it’s advantageous to think about each of those business functions separately."

    -- "IDDD" by Vaughn Vernon



Возьмем известный пример Эрика Эванса про различные виды карт Земли. Если мы будем пытаться полностью воспроизвести Землю, то есть, все возможные виды карт в одной модели карты,
то с такой картой мы не сможем решать вообще никакую проблему из-за сложности в целом и огромного количества нюансов/атрибутов, не связанных с нашей задачей.
Именно поэтому модель является упрощенной интерпретации реальности и работает внутри определенного контекста применимости (и для решения строго определенных задач).

В качестве иллюстрации сложности "единой" модели, рассмотрим примеры из доклада Эрика Эванса (Eric Evans — Tackling Complexity in the Heart of Software, Domain-Driven Design Europe 2016 - Brussels, January 26-29, 2016)

1. Карта морского ориентирования (цилиндрическая проекция Меркатора)

.. figure:: _media/mercator_projection.png
   :alt: Mercator projection
   :align: center
   :width: 100%

Такие карты используют относительное искажение размеров объектов относительно друг друга, но помогают направлять компас в сторону нужной конечной точки (направление на карте полностью совпадет со стрелкой компаса для ориентирования морских судов).
На этой карте Африка и Гренландия выглядят равными по площади, но в действительности, Африка в 14 раз больше Гренландии, то есть у карты есть четкое предназначение, задача для которой она нужна, и только для нее - навигация судов (то есть, для школьных уроков географии она не подходит из-за искажения той информации, которая важна для урока).

2. Картографическая проекция земного шара на поверхность многогранника (проекция Димаксион (Фуллера))

.. figure:: _media/fuller_projection.png
   :alt: Fuller projection
   :align: center
   :width: 100%

Фуллер утверждал, что его проекция имеет ряд преимуществ над другими проекциями земного шара. Она имеет меньшие искажения относительных размеров объектов, особенно в сравнении с проекцией Меркатора,
то есть, модель карты может служить более точным инструментом определения относительных размеров объектов земли.



Важное уточнение: модель по Тарасенко
-------------------------------------

