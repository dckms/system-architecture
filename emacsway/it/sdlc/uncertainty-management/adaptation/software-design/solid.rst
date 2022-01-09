:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: SOLID; in Agile
   :name: emacsway-agile-solid


=================================
Role of SOLID principles in Agile
=================================

.. sectionauthor:: Ivan Zakrevsky

.. contents:: Содержание


.. index::
   single: SOLID; in misunderstanding
   :name: emacsway-agile-solid-misunderstanding


Ошибочная трактовка SRP
=======================

SRP is not do just one thing
----------------------------

Существуют две распространенные ошибки применения принципа SRP:

1. SRP якобы подразумевает делать только одну вещь.
2. SRP якобы применяется к компонентам, например, к микросервисам.

Давайте немного исследуем этот вопрос.

В своей книге "Clean Architecture: A Craftsman's Guide to Software Structure and Design", Robert C. Martin сожалеет, что выбрал такое название (SRP):

    📝 "Of all the SOLID principles, the Single Responsibility Principle (SRP) might be **the least well understood**.
    That's likely because it has a particularly **inappropriate name**. 
    It is too easy for programmers **to hear the name and then assume that it means that every module should do just one thing**.

    Make no mistake, there is a principle like that.
    A function should do one, and only one, thing.
    We use that principle when we are refactoring large functions into smaller functions; we use it at the lowest levels.
    **But it is not one of the SOLID principles—it is not the SRP**.

    Historically, the SRP has been described this way:

        *A module should have one, and only one, reason to change.*

    Software systems are changed to satisfy users and stakeholders; those users and stakeholders are the "reason to change" that the principle is talking about.
    Indeed, we can rephrase the principle to say this:

        *A module should be responsible to one, and only one, user or stakeholder.*

    Unfortunately, the words "user" and "stakeholder" aren't really the right words to use here.
    There will likely be more than one user or stakeholder who wants the system changed in the same way.
    Instead, we're really referring to a group—one or more people who require that change.
    We'll refer to that group as an actor.

    Thus the final version of the SRP is:

        *A module should be responsible to one, and only one, actor.*

    Now, what do we mean by the word "module"? The simplest definition is just a source file.
    Most of the time that definition works fine.
    Some languages and development environments, though, don't use source files to contain their code.
    In those cases a module is just a cohesive set of functions and data structures.

    That word "cohesive" implies the SRP.
    Cohesion is the force that binds together the code responsible to a single actor.

    Perhaps the best way to understand this principle is by looking at the symptoms of violating it..."

    -- "Clean Architecture: A Craftsman's Guide to Software Structure and Design" by Robert C. Martin


SRP is about Cohesion
---------------------

А в своей книге 2002 года, которая вышла в свет через год после подписания Agile Manifesto (подписание которого, собственно, он же и организовал), Robert C. Martin выводит определение SRP из понятия "сфокусированности" (cohesion) класса:

    📝 "SRP: The Single-Responsibility Principle

    **This principle was described in the work of Tom DeMarco [1] and Meilir Page-Jones [2].**
    **They called it cohesion.**
    They defined cohesion as the functional relatedness of the elements of a module.
    In this chapter we'll shift that meaning a bit and relate cohesion to the forces that cause a module, or a class, to change.

    A class should have only one reason to change.

    1. [DeMarco79], p. 310.
    2. [Page-Jones88], Chapter 6, p. 82.

    1. DeMarco, Tom. Structured Analysis and System Specification. Yourdon Press Computing Series. Englewood Cliff, NJ: 1979.
    2. Page-Jones, Meilir. The Practical Guide to Structured Systems Design, 2d ed. Englewood Cliff, NJ: Yourdon Press Computing Series, 1988."

    -- "Agile Software Development. Principles, Patterns, and Practices." by Robert C. Martin, James W. Newkirk, Robert S. Koss

Многое встает на свое место, если принимать во внимание Cohesion, т.е. использовать изначальный принцип "`Low Coupling & High Cohesion <http://wiki.c2.com/?CouplingAndCohesion>`__".

Часто можно слышать, что применение принципов SOLID ведет к появлению нечитаемого кода.
Очень хорошо подобную проблему выразил Eric Evans (правда, по другим причинам):

    📝 "Если требования архитектурной среды к распределению обязанностей таковы, что элементы, реализующие концептуальные объекты, оказываются физически разделенными, то код больше не выражает модель.

    Нельзя разделять до бесконечности, у человеческого ума есть свои пределы, до которых он еще способен соединять разделенное; если среда выходит за эти пределы, разработчики предметной области теряют способность расчленять модель на осмысленные фрагменты.

    If the framework's partitioning conventions pull apart the elements implementing the conceptual objects, the code no longer reveals the model.

    There is only so much partitioning a mind can stitch back together, and if the framework uses it all up, the domain developers lose their ability to chunk the model into meaningful pieces."

    -- "Domain-Driven Design: Tackling Complexity in the Heart of Software" by Eric Evans, перевод В.Л. Бродового

О том, что использованием принципов SOLID можно переусложнить программу, пишет и весьма авторитетный в области программной разработки Сергей Тепляков:

- "`Принцип YAGNI <http://sergeyteplyakov.blogspot.com/2016/08/yagni.html>`__"
- "`Критика книги Боба Мартина "Принципы, паттерны и методики гибкой разработки на языке C#" <http://sergeyteplyakov.blogspot.com/2013/12/about-agile-principles-patterns-and.html>`__"
- "`Идеальная архитектура <http://sergeyteplyakov.blogspot.com/2011/11/blog-post_23.html>`__"
- "`Шпаргалка по SOLID принципам <http://sergeyteplyakov.blogspot.com/2014/10/solid.html>`__"
- "`О принципах проектирования <http://sergeyteplyakov.blogspot.com/2014/10/about-design-principles.html>`__"
- "`О дизайне <http://sergeyteplyakov.blogspot.com/2012/07/blog-post.html>`__"
- "`О повторном использовании кода <http://sergeyteplyakov.blogspot.com/2012/04/blog-post_19.html>`__"

Лично мне на практике не доводилось наблюдать сложности от использования принципов SOLID, разве что только в проектах с использованием Redux.
Кстати, у Udi Dahan есть прекрасная статья "`Clarified CQRS <http://udidahan.com/2009/12/09/clarified-cqrs/>`__" о том, как грамотно разделять бизнес-логику и логику приложения в CQRS-приложении (а `Redux реализаует принципы CQRS <https://redux.js.org/understanding/thinking-in-redux/motivation>`__), чтобы предотвратить фрагментирование бизнес-логики.

Я обнаружил еще одну причину столь широкого недопонимания этого принципа.
В переводе книги "Clean Code" термин "Single" переводится как "Единый".
А в книге "Clean Architecture" - как "Единственный".

Эти термины похожи, но не идентичны.
Так, например, "Единое гражданство" означает то, что административно-территориальные единицы государства не могут вводить свое собственное гражданство.
Но при этом, граждане могут иметь двойное гражданство.
А вот "Единственное гражданство" уже подразумевает запрет на двойное гражданство.

Таким образом, термин "Единый" подразумевает "Сфокусированный" на конкретной задаче, т.е. нефрагментированный.
Иными словами, речь идет о "High Cohesion", что восходит к Constantine's Law - "Low Coupling & High Cohesion", о чем прямо говорит Robert C. Martin по приведенным выше ссылкам.

    📝 "**That word "cohesive" implies the SRP.**
    Cohesion is the force that binds together the code responsible to a single actor."

    -- "Clean Architecture: A Craftsman's Guide to Software Structure and Design" by Robert C. Martin

Так, например, метод рефакторинга "`Inline Class <https://refactoring.com/catalog/inlineClass.html>`__" не противоречит SRP, хотя класс и отбирает обязанность у другого класса в случае, когда её недостаточно для самостоятельного существования.

С другой стороны, если фрагментировать класс, понижая его Cohesion, то это будет противоречить принципу SRP, хотя мы и получим классы с дистиллированными обязанностями без примесей.

Качественный Software Design должен облегчать понимание кода, а не затруднять.

К сожалению, сложности перевода встречаются нередко.
Так, например, до сих пор нет единого мнения о том, как правильно переводить термины "Coupling" и "Cohesion", и различные источники дают прямо противоположный перевод.


SRP and "Conway's law"
----------------------

А здесь Robert C. Martin выводит понимание SRP исходя из "Conway's law":

    📝 "SRP: The Single Responsibility Principle

    An active **corollary to Conway's law**: The best structure for a software system is heavily influenced by the social structure of the organization that uses it so that each software module has one, and only one, reason to change."

    -- "Clean Architecture: A Craftsman's Guide to Software Structure and Design" by Robert C. Martin


Common Closure Principle (CCP)
------------------------------

К компонентам применяется похожий, но другой, принцип, который называется "Common Closure Principle (CCP)":

    📝 "THE COMMON CLOSURE PRINCIPLE

    Gather into components those classes that change for the same reasons and at the same times.
    Separate into different components those classes that change at different times and for different reasons.

    This is the Single Responsibility Principle restated for components.
    Just as the SRP says that a class should not contain multiples reasons to change, so the Common Closure Principle (CCP) says that a component should not have multiple reasons to change.

    For most applications, maintainability is more important than reusability.
    If the code in an application must change, you would rather that all of the changes occur in one component, rather than being distributed across many components. [1]
    If changes are confined to a single component, then we need to redeploy only the one changed component.
    Other components that don't depend on the changed component do not need to be revalidated or redeployed.

    The CCP prompts us to gather together in one place all the classes that are likely to change for the same reasons.
    If two classes are so tightly bound, either physically or conceptually, that they always change together, then they belong in the same component.
    This minimizes the workload related to releasing, revalidating, and redeploying the software.

    This principle is closely associated with the Open Closed Principle (OCP).
    Indeed, it is "closure" in the OCP sense of the word that the CCP addresses.
    The OCP states that classes should be closed for modification but open for extension.
    Because 100% closure is not attainable, closure must be strategic.
    We design our classes such that they are closed to the most common kinds of changes that we expect or have experienced.

    The CCP amplifies this lesson by gathering together into the same component those classes that are closed to the same types of changes.
    Thus, when a change in requirements comes along, that change has a good chance of being restricted to a minimal number of components."

    1. See the section on "The Kitty Problem" in Chapter 27, "Services: Great and Small."

    -- "Clean Architecture: A Craftsman's Guide to Software Structure and Design" by Robert C. Martin

Нужно учитывать, что под компонентом Robert C. Martin понимает единицу развертывания (в других источниках этот термин может иметь другое значение):

    📝 "Components are the units of deployment.
    They are the smallest entities that can be deployed as part of a system.
    In Java, they are jar files.
    In Ruby, they are gem files.
    In .Net, they are DLLs.
    In compiled languages, they are aggregations of binary files.
    In interpreted languages, they are aggregations of source files.
    In all languages, they are the granule of deployment."

    -- "Clean Architecture: A Craftsman's Guide to Software Structure and Design" by Robert C. Martin


SRP and complexity
------------------

Еще существует распространенное мнение, что SOLID уменьшает сложность программы.

    📝 сложность
        1. Составленность из нескольких частей; многообразность по составу входящих частей и связей между ними.
        2. Трудность, запутанность. Противоположное понятие — простота.

    -- "`Словарь практического психолога <https://psychology.academic.ru/2331/%D1%81%D0%BB%D0%BE%D0%B6%D0%BD%D0%BE%D1%81%D1%82%D1%8C>`__". — М.: АСТ, Харвест. С. Ю. Головин. 1998.

..

    📝 сложный
        1. Состоящий из нескольких частей, элементов. от т. перен. Характеризующийся многими переплетающимися явлениями, признаками, отношениями.
        2. перен. Представляющий трудность для понимания, разрешения, осуществления; трудный.
        3. перен. Обладающий противоречивыми качествами, свойствами, особенностями.

    -- "`Толковый словарь Ефремовой <https://dic.academic.ru/dic.nsf/efremova/246102/%D1%81%D0%BB%D0%BE%D0%B6%D0%BD%D1%8B%D0%B9>`__". Т. Ф. Ефремова. 2000.

..

    📝 Complex

    Com"plex (kŏm"plĕks), a. [L. complexus, p. p. of complecti to entwine around, comprise; com- + plectere to twist, akin to plicare to fold. See Plait, n.]

    1. Composed of two or more parts; composite; not simple; as, a complex being; a complex idea.

        Ideas thus made up of several simple ones put together, I call complex; such as beauty, gratitude, a man, an army, the universe.
        Locke.

    2. Involving many parts; complicated; intricate.

        When the actual motions of the heavens are calculated in the best possible way, the process is difficult and complex.
        Whewell.

    Complex fraction. See Fraction. -- Complex number (Math.), in the theory of numbers, an expression of the form a + b√-1, when a and b are ordinary integers.

    Syn. -- See Intricate.

    Com"plex, n. [L. complexus] Assemblage of related things; collection; complication.

        This parable of the wedding supper comprehends in it the whole complex of all the blessings and privileges exhibited by the gospel.
        South.

    Complex of lines (Geom.), all the possible straight lines in space being considered, the entire system of lines which satisfy a single relation constitute a complex; as, all the lines which meet a given curve make up a complex. The lines which satisfy two relations constitute a congruency of lines; as, the entire system of lines, each one of which meets two given surfaces, is a congruency.

    -- `webster's 1913 <http://www.websters1913.com/words/Complex>`__. Connoisseur's reference to American English - a dictionary for writers and wordsmiths

Если рассматривать термин в значении "простота понимания", то правильное использование принципов SOLID, наоборот, облегчает понимание.

Если же рассматривать термин в значении "многообразность по составу входящих частей", то совокупная сложность программы (в общей сложности) не уменьшается, а наоборот возрастает.
Задача архитектурных принципов сводится не к тому, чтобы уменьшить сложность, а к тому, чтобы управлять сложностью.
Это позволяет формировать структуру программы таким образом, чтобы отдельные её части можно было рассматривать изолированно, сохраняя рассматриваемый уровень сложности в пределах `возможностей краткосрочной памяти человека <https://ru.m.wikipedia.org/wiki/%D0%9C%D0%B0%D0%B3%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%BE%D0%B5_%D1%87%D0%B8%D1%81%D0%BB%D0%BE_%D1%81%D0%B5%D0%BC%D1%8C_%D0%BF%D0%BB%D1%8E%D1%81-%D0%BC%D0%B8%D0%BD%D1%83%D1%81_%D0%B4%D0%B2%D0%B0>`__.

Отсюда вывод - применение любого паттерна или принципа, вносящего в систему несущественную сложность (accidental complexity), должно себя окупать, т.е. позволять управлять еще большим уровнем сложности.
Тогда применение принципов и паттернов, хотя и будет (математически) усложнять программу, но будет упрощать понимание программы, формируя такие уровни абстракции, которые человеческий мозг сможет рассматривать изолированно.

Методики управления сложностью позволяют предотвратить Уроборос.

Но есть еще одно значение этого термина:

    📝 "Structural Complexity looks at the system elements and relationships.
    In particular, structural complexity looks at how many different ways system elements can be combined.
    Thus, it is related to the potential for the system to adapt to external needs."

    -- "`Guide to the Systems Engineering Body of Knowledge (SEBoK) <https://www.sebokwiki.org/wiki/Complexity>`__"

Если рассматривать термин в этом значении, то SOLID увеличивает сложность, но это не имеет негативного влияния на понимание устройства системы.


Ad hominem
----------

Еще одно распространенное мнение, которое часто можно услышать, заключается в том, что Robert C. Martin - оторванный от практики теоретик, придумывающий в своем иллюзорном мирке всякие нежизнеспособные принципы вроде SOLID, которые на практике только ухудшают код.

Принципы SOLID действительно, имеют под собой теоретическое обоснование, только эта теория не связана с Robert C. Martin.
А вот, например, Bertrand Meyer, действительно, является серьезным научным теоретиком, и его авторство Robert C. Martin не скрывал в своей оригинальной статье "`The Open-Closed Principle <https://web.archive.org/web/20060822033314/http://www.objectmentor.com/resources/articles/ocp.pdf>`__".

Итак, вывод первый - если кто и является теоретиком, то это не Robert C. Martin. Он-то как раз практик.

В архитекторских кругах отношение к Robert C. Martin можно назвать, мягко говоря, неоднозначным.
Зато к Gregor Hohpe отношение - почти единодушно уважительное.

Но, странное дело, первая книга в `списке рекомендованной литературы Gregor Hohpe <https://architectelevator.com/architecture/architect-bookshelf/>`__ - это именно книга "Clean Code" by Robert C. Martin.

Мое же мнение сводится к тому, что говорить о влиянии внутреннего качества кода на характер роста стоимости изменения кода - нужно.
Именно этим и занимается Robert C. Martin.
И это `важнее манеры донесения информации <https://en.m.wikipedia.org/wiki/Ad_hominem>`__.
Потому что это - один из наиболее чувствительных вопросов индустрии:

- https://t.me/emacsway_log/458
- https://t.me/emacsway_log/462


SOLID and Agile
===============

Принципы SOLID впервые появились в статье "Design Principles and Design Patterns" 2000 года:

- "`Источник 1 <https://sites.google.com/site/unclebobconsultingllc/getting-a-solid-start>`__"
- "`Источник 2 <https://web.archive.org/web/20150906155800/http://www.objectmentor.com/resources/articles/Principles_and_Patterns.pdf>`__"
- "`Источник 3 <https://fi.ort.edu.uy/innovaportal/file/2032/1/design_principles.pdf>`__"

Вышла эта статья за год до того, как тот же Robert C. Martin организовал встречу 17-ти, на которой был принят Agile-Manifesto.
Как между собой связаны два этих события?

Все просто.
Agile - это адаптивная методика, которая имеет экономическую целесообразность только в том случае, если :ref:`стоимость адаптации ниже стоимости заблаговременного проектирования (BDUF) <emacsway-agile-development>`.

А стоимость адаптации, благодаря которой итеративная разработка вообще обретает смысл, определяется характером роста стоимости изменения кода.
А это уже :ref:`задача архитектурная <emacsway-agile-software-design>`, и это объясняет, почему на подписании Agile-manifesto присутствовало столько людей из мира архитектуры.
Кстати, на этой встрече предполагалось присутствие и Grady Booch, но, не вышло.

И это так же объясняет, почему первая книга, которую выпустил организатор встречи Agile-Manifesto после его подписания, была посвящена не столько процессам, сколько принципам конструирования (гибкого) кода, обладающего низкой стоимостью изменения.
Это лишний раз подчеркивает :ref:`важность технической составляющей в Agile (гибкой) разработке <emacsway-agile-development-difficulties>`.

Итак, следующий важный вывод: если бы принципы конструирования гибкого кода, включая SOLID, не имели бы практического улучшения экономических показателей разработки, тогда Robert C. Martin с единомышленниками никогда не смог бы доказать бизнесу, что Agile обладает экономическим превосходством перед BDUF, и рынок просто его проигнорировал бы.


Баланс краткосрочных и долгосрочных интересов
=============================================

`Quality Attributes <https://en.m.wikipedia.org/wiki/List_of_system_quality_attributes>`__ противоречивы между собой, и удовлетворить их все не представляется возможным.
Поэтому, приложение не может быть лучше или хуже - оно может соответствовать или не соответствовать требуемым атрибутам качества.

Принципы SOLID направлены на удовлетворение атрибута качества `Modifiability <https://resources.sei.cmu.edu/library/asset-view.cfm?assetid=8299>`__ (см. "Software Architecture in Practice" 3d edition by Len Bass, Paul Clements, Rick Kazman) в долгосрочной перспективе.
Т.е. они призваны обеспечить пологий характер роста стоимости изменения кода, максимально приближенный к горизонтальной асимптоте.
Напомню, принципы SOLID были опубликованы в контексте Agile разработки, где это требование является критически необходимым для достижения экономического превосходства Agile-разработки перед BDUF.

Чтобы находить баланс наименьшей стоимости разработки как в долгосрочной, так и в краткосрочной перспективе, нужно сочетать принципы SOLID с принципом YAGNI (который отвечает за снижение стоимости в краткосрочной перспективе), о чем писал Сегей Тепляков в статьях:

- "`Принцип YAGNI <http://sergeyteplyakov.blogspot.com/2016/08/yagni.html>`__"
- "`О повторном использовании кода <http://sergeyteplyakov.blogspot.com/2012/04/blog-post_19.html>`__"

Сам Robert Martin дает такое определение качеству дизайна:

    📝 "The measure of design quality is simply the measure of the effort required to meet the needs of the customer.
    If that effort is low, and stays low throughout the lifetime of the system, the design is good.
    If that effort grows with each new release, the design is bad.
    It's as simple as that."

    -- "Clean Architecture: A Craftsman's Guide to Software Structure and Design" by Robert C. Martin

Ключевым здесь является "stays low throughout the lifetime of the system" (т.е. в долгосрочной перспективе), поскольку существует `Design Payoff Line <https://martinfowler.com/bliki/DesignPayoffLine.html>`__.

Не должно быть принципов ради принципов, когда за деревьями леса не видно.
Если принципы применяются, а стоимость разработки возрастает, значит, применяются либо не те принципы, либо не так.
Как говорил Craig Larman:

    📝 "в продуктовой разработке нет такого понятия как "лучшие практики" - есть только практики, применение которых целесообразно в конкретном контексте.
    Практики ситуационны, и беспечное объявление их "лучшими" отрывает их от мотивации и контекста.
    Они превращаются в ритуалы, и навязывание так называемых "best practices" убивает культуру обучения, задавания вопросов, вовлечения и непрерывных улучшений.
    Зачем людям искать чего-то лучшего, если все уже придумано за них?"

    -- "`Знакомство с LeSS <https://less.works/ru/less/framework/introduction>`__"

    📝 "There are no such things as best practices in product development.
    There are only practices that are adequate within a certain context.
    Practices are situational; blithely claiming they are "best" disconnects them from motivation and context.
    They become rituals. And pushing so-called best practices kills a culture of learning, questioning, engagement, and continuous improvement.
    Why would people challenge best?"

    -- "`Introduction to LeSS <https://less.works/ru/less/framework/introduction>`__"

Иными словами, нужно осознавать, для достижения какого именно требования применяется тот или иной принцип, следить за фидбэком от его применения, и анализировать успешность достижения этого требования. Без этого применение принципов может легко превратиться в `Карго-Культ <http://sergeyteplyakov.blogspot.com/2013/09/blog-post_24.html>`__.


SOLID и первоисточники
======================

Остается еще одно распространенное мнение - Robert C. Martin исказил оригинальный смысл первоисточников.
Где-то в чем-то он может и ошибся, да он и сам об этом говорил.
Но он не присваивал себе чужие идеи, и всегда открыто отсылал к первоисточникам, таким образом, привлекая к ним внимание.

OCP
---

Вот, например, многие из ваших коллег узнали бы об OCP из оригинала в изложении Bertrand Meyer?
Даже Martin Fowler говорил, что:

    📝 "the second  edition [of "Object Oriented Software Construction"] is good but you'll need several months in a gym before you  can lift it."

    -- M.Fowler, `Command Query Separation <https://martinfowler.com/bliki/CommandQuerySeparation.html>`__

При этом, Robert C. Martin, действительно возлагал на старые принципы новые задачи, исходя из исторического контекста того времени.
Очень хорошо этот вопрос рассматривается в статье "`OCP vs YAGNI <https://enterprisecraftsmanship.com/posts/ocp-vs-yagni/>`__" by Vladimir Khorikov.

    📝 "There are two interpretations of the Open/Closed Principle:

        1. The original Bertrand Meyer's one is about backward compatibility. You need to close the API of your module/library/service if it's meant for external use. Not implementation but exactly the API part of it. And only when it's used by external teams.
        2. The Bob Martin's one is about avoiding ripple effects: you need to be able to extend the software behavior with modifying little or no original code. This is achieved by putting extension points to your code base."

    -- "`OCP vs YAGNI <https://enterprisecraftsmanship.com/posts/ocp-vs-yagni/>`__" by Vladimir Khorikov

SRP
---

В качестве источника этого принципа, Robert C. Martin в своей книге "Agile Software Development. Principles, Patterns, and Practices." указывает:

    1. DeMarco, Tom. Structured Analysis and System Specification. Yourdon Press Computing Series. Englewood Cliff, NJ: 1979.
    2. Page-Jones, Meilir. The Practical Guide to Structured Systems Design, 2d ed. Englewood Cliff, NJ: Yourdon Press Computing Series, 1988.

    -- "Agile Software Development. Principles, Patterns, and Practices." by Robert C. Martin

А здесь он дает более развернутую историю:

    📝 "The 1970s and 1980s were a fertile time for principles of software architecture.
    Structured Programming and Design were all the rage.
    During that time the notions of Coupling and Cohesion were introduced by Larry Constantine, and amplified by Tom DeMarco, Meilir Page-Jones and many others."

    -- "`The Single Responsibility Principle <https://blog.cleancoder.com/uncle-bob/2014/05/08/SingleReponsibilityPrinciple.html>`__" by Robert C. Martin

Т.е. он отсылает к `Constantine's Law <http://wiki.c2.com/?CouplingAndCohesion>`__.

Этот принцип не имеет отношения к OOP, хотя и активно используется в OOP, в частности в книгах:

- "Code Complete" by Steve McConnell
- "Applying UML and Patterns: An Introduction to Object-Oriented Analysis and Design and Iterative Development" by Craig Larman, где этот принцип известен под аббревиатурой GRASP

И этот принцип является одним из самых фундаментальных в разработке ПО.
Он применяется в структурном программировании, в OOP, в DDD при моделировании агрегатов, в микросервисах при поиске границ микросервисов и т.п.

А Kent Beck назвал его в одной из своих недавних статей

    📝 "the basic forces acting on software design.
    These were elucidated in the mid-70s by Yourdon & Constantine in `Structured Design <https://amzn.to/2GsuXvQ>`__ and haven't changed."

    -- "`Monolith -> Services: Theory & Practice <https://medium.com/@kentbeck_7670/monolith-services-theory-practice-617e4546a879>`__" by Kent Beck

Каким бы образом это ни сделал Robert C. Martin, но он достиг поставленной цели - привлек массовое внимание к архитектурным принципам, имеющим действительно важное значение в Agile разработке.


No Silver Bullet
----------------

Но правда и в том, что `серебряной пули нет <https://en.m.wikipedia.org/wiki/No_Silver_Bullet>`__.
Нет магических пяти букв, которые волшебным образом сделают код экономически высокоэффективным, особенно, если про них узнали не из первоисточника, а из Википедии.
Зайдите на github, и вы увидите огромное многообразие взаимно-противоречивых демонстрационных приложений (reference applications), и каждое из них претендует на роль самой эталонной реализации Clean Architecture (это уже не совсем SOLID, но из этой же области).
Кто в лес, кто по дрова.
И только работы парней, обладающих `широким архитектурным кругозором и литературным бэкграундом <http://www.kamilgrzybek.com/programming-and-design-resources/>`__, таких, как `Kamil Grzybek <https://github.com/kgrzybek>`__, могут претендовать на роль эталонной реализации:

- `Full Modular Monolith application with Domain-Driven Design approach <https://github.com/kgrzybek/modular-monolith-with-ddd>`__ by Kamil Grzybek
- `Sample .NET Core REST API CQRS implementation with raw SQL and DDD using Clean Architecture <https://github.com/kgrzybek/sample-dotnet-core-cqrs-api>`__ by Kamil Grzybek

В таком случае, SOLID уже не мешает создавать экономически высокоэффективные приложения (`раз <http://www.kamilgrzybek.com/design/clean-domain-model-attributes/>`__ и `два <http://www.kamilgrzybek.com/design/grasp-explained/>`__).

Еще одно, наверное, самое популярное reference application - `eShopOnContainers <https://github.com/dotnet-architecture/eShopOnContainers>`__, было разработано при значительном участии известного авторитета в области архитектуры Cesar De La Torre, который в своем `блог-посте <https://devblogs.microsoft.com/cesardelatorre/free-ebookguide-on-net-microservices-architecture-for-containerized-net-applications/>`__ пишет:

    "Using SOLID principles and Dependency Injection"

Речь идет об `этом месте гайда <https://docs.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/microservice-application-layer-web-api-design>`__.

В качестве итога можно привести слова самого Robert C. Martin:

    📝 "**Following the rules on the paint can won't teach you how to paint.**
    This is an important point.
    Principles will not turn a bad programmer into a good programmer.
    Principles have to be applied with judgement.
    If they are applied by rote it is just as bad as if they are not applied at all.
    Having said that, if you want to paint well, I suggest you learn the rules on the paint can.
    You may not agree with them all.
    You may not always apply the ones you do agree with.
    But you'd better know them.
    Knowledge of the principles and patterns gives you the justification to decide when and where to apply them.
    If you don't know them, your decisions are much more arbitrary."

    -- "`Getting a SOLID start. <https://sites.google.com/site/unclebobconsultingllc/getting-a-solid-start>`__" by Robert C. Martin

См. также:

- "`Getting a SOLID start <https://sites.google.com/site/unclebobconsultingllc/getting-a-solid-start>`__"
- "`The Single Responsibility Principle <https://blog.cleancoder.com/uncle-bob/2014/05/08/SingleReponsibilityPrinciple.html>`__"
- "`The Open Closed Principle <https://blog.cleancoder.com/uncle-bob/2014/05/12/TheOpenClosedPrinciple.html>`__"
- "`Solid Relevance <http://blog.cleancoder.com/uncle-bob/2020/10/18/Solid-Relevance.html>`__"

.. seealso::

   - ":doc:`../crash-course-in-software-development-economics`"


.. todo::

   - https://t.me/emacsway_log/458
   - https://t.me/emacsway_log/462
   - https://t.me/emacsway_log/393
