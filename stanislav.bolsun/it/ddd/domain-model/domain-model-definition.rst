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

