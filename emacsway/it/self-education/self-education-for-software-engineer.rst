:canonical-base-url: https://dckms.github.io/system-architecture

.. index:: Literature, Self-education

========================================================================
Список литературы для самообучения разработчика программного обеспечения
========================================================================

.. sectionauthor:: Ivan Zakrevsky

Один из частых вопросов, который я наблюдаю регулярно, - это "посоветуйте список литературы в области разработки программного обеспечения".
В этой статье я изложу свое видение самообучения и приведу список тематической литературы, с учетом моего личного опыта.

.. contents:: Содержание


Предисловие
===========

Классическая ошибка новичков - жажда к знаниям, нетерпеливость.
Обычно это приводит к тому, что, в погоне за количеством, они :doc:`надрываются </emacsway/soft-skills/planning-in-psychology>` (объем знаний, который предстоит освоить, действительно, огромный), осознают невыполнимость желаемого, впадают в депрессию, а затем и в состояние `психологической защиты <https://ru.wikipedia.org/wiki/%D0%97%D0%B0%D1%89%D0%B8%D1%82%D0%BD%D1%8B%D0%B9_%D0%BC%D0%B5%D1%85%D0%B0%D0%BD%D0%B8%D0%B7%D0%BC>`__ (мол, "академичность" неуместна на практике), и прекращают развиваться.
Решается эта проблема очень просто - жажда должна быть не к знаниям, а к дисциплине.
А уж дисциплина обязательно приведет к обретению знаний.
Дисциплина - это, своего рода, производная знаний.
Она поддерживает постоянную скорость на пути освоения знаний.
Сперва нужно выработать привычку, а затем привычка будет работать на вас.
Как говорится, сохраняйте порядок, и порядок сохранит вас.

Достаточно читать по 5 страниц в день.
Тут главное - стабильность.
Пусть будет по чуть-чуть, но постоянно.
Дисциплина - мать победы, говорил А.В. Суворов.
Гнаться за количеством не нужно.

    "A little reading goes a long way toward professional advancement. If you read even one
    good programming book every two months, roughly 35 pages a week, you'll soon have
    a firm grasp on the industry and distinguish yourself from nearly everyone around you."

    \- "Code Complete" by Steve McConnell

..

    "We become authorities and experts in the practical and scientific spheres
    by so many separate acts and hours of work.
    If a person keeps faithfully busy each hour of the working day,
    he can count on waking up some morning to find himself one of the competent
    ones of his generation."

    \- William James, cited by Steve McConnell in "Code Complete"

И, желательно, чтобы читаемая книга совпадала с тематикой текущего проекта, чтобы через практику хорошо легла в память.
Я по этой причине часто изменял свой график чтения.
Обычно я читал в параллели 2-3 книги. Одну - планово, другие - по потребностям проекта.

Еще одной ошибкой является неудачный выбор литературы.
Сегодня штампуется много литературы, но далеко не каждая книга достойна внимания.
`Закон Парето <https://ru.wikipedia.org/wiki/%D0%97%D0%B0%D0%BA%D0%BE%D0%BD_%D0%9F%D0%B0%D1%80%D0%B5%D1%82%D0%BE>`__ работает и здесь.

Хорошей вещью для систематизации собственных знаний является написание статей и участие в профессиональных дискуссиях.
Ничто так не систематизирует собственные знания, как попытка объяснить что-то другому человеку.
Вы, конечно, будете периодически ошибаться, но для кристализации знаний это лучше, чем ничего не делать.
К тому же, это хорошо развивает сдержанность в аргументации, что немаловажно.

На первых порах критически важно участвовать в Open Source проектах.
Можно завести свои собственные Open Source проекты.
Можно принимать участие в каких-то существующих проектах с авторитетными комьюнити, которые будут помогать избавляться от ошибок.
В любом случае, не надейтесь на то, что профессиональные проекты предоставят вам достаточную практику для закрепления знаний.
А Open Source проекты - очень даже предоставят.
Я даже считаю, что практика должна предшествовать теории, потому что трудно запомнить какое-то решение, если вам на практике не знакома решаемая проблема.
Потребность в теории должна назреть.
Когда я приступал к теории, то у меня был накоплен уже солидный багаж проблем, решение которых я искал.
Когда я впервые прочитал о мотивации паттерна Bridge, у меня в голове промелькнуло: "так вот, оказывается, как решается та самая проблема".
Когда я читал каталог Code Smells, я частенько вспоминал свой код.
В результате, решения навечно запечатлелись в памяти.

Очень правильно `сказал <https://sergeyteplyakov.blogspot.com/2017/02/reading-books-considered-harmful.html>`__ Сергей Тепляков: "Полноценное обучение – это не теория vs. практика. Это комбинация этих вещей, при этом процент одного и другого зависит от человека и изучаемой темы."

Ну и, главное, не впадать в фанатизм.
Засасывает.
Нужно себя уравновешивать другими интересами, семья, спорт, физкультура, шашлыки, друзья, путешествия...
Непредвзятый и свободный взгляд намного важнее изобилия знаний.
Путешествие должно быть на легке, как говорил Кент Бек.
По сути, знания нужны только для того, чтобы избавиться от всего лишнего.
Архитектура - это, на самом деле, наука об ограничениях (т.е. о том, как не надо делать).


Кандидатский минимум
====================


Учимся обучению
---------------

Это может показаться немного удивительным, но первая книга будет посвящена не техническим знаниям, а вопросам самоорганизации, управления временем, психологии, методикам работы под стрессом, оцеванию задач по разработке программного обеспечения, вопросам коммуникации и поведению в конфликтных ситуациях, и, самое главное, - науке быть правдивым.
Именно правдивость является важнейшим отличительным признаком настоящего профессионала.
И это не так просто, как может показаться на первый взгляд.
Есть разница между кодером и профессионалом.
И эта книга о том, как стать профессионалом.
Без знаний, изложенных в этой книге, вы просто не сможете изыскать время на самообучение, и список остальных книг вам может просто не понадобиться:

- "The Clean Coder" by Robert C. Martin


Изучаем основную используемую технологию
----------------------------------------

Следующая книга должна быть посвящена основной используемой технологии, т.е. синтаксическим возможностям языка программирования.
Для Python-разработчиков хорошим выбором была бы книга:

- "Learning Python" 5th edition by Mark Lutz

Для Golang интересно выглядят книги:

- "Hands-On Software Architecture with Golang. Design and architect highly scalable and robust applications using Go" by Jyotiswarup Raiturkar
- "The Go Programming Language" by Alan A.A. Donovan Google Inc., Brian W. Kernighan Princeton University

Для Erlang:

- "Programming Erlang: Software for a Concurrent World (Pragmatic Programmers)" 2nd edition by Joe Armstrong

Для frontend-разработчиков, работающих с Angular, имеет смысл обратить внимание на книгу:

- "ng-book2. The Complete Book on Angular 6" by Nate Murray, Felipe Coury, Ari Lerner, and Carlos Taborda


Азбука программирования
-----------------------

Подразумевается что вы уже хорошо знаете синтаксис основного языка программирования.
Но, знание букв еще не делает вас поэтом.
Следующие книги являются азбукой программирования.
Я привожу их в таком порядке, в каком я рекомендую их прочтение:

- "Design Patterns: Elements of Reusable Object-Oriented Software" by Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides
- "Patterns of Enterprise Application Architecture" by Martin Fowler, David Rice, Matthew Foemmel, Edward Hieatt, Robert Mee, Randy Stafford
- "Refactoring: Improving the Design of Existing Code" 1st edition by Martin Fowler, Kent Beck, John Brant, William Opdyke, Don Roberts
- "Clean Code: A Handbook of Agile Software Craftsmanship" by Robert C. Martin
- "Code Complete" 2nd edition by Steve McConnell
- "UML Distilled. A Brief Guide to the Standard Object Modeling Language" 3d edition by Martin Fowler
- "`KISS Principles <https://people.apache.org/~fhanik/kiss.html>`__"


Учимся быть эффективным
-----------------------

Знаний предыдущих пяти книг достаточно для того, чтобы вы стали работать в разы эффективней.
Но нужно не только знать, а еще и :doc:`уметь быть эффективным на практике </emacsway/it/tdd/tdd>`.
Никто не раскрывает этот вопрос лучше, чем Kent Beck:

- "Test-Driven Development By Example" by Kent Beck


Учимся делать команду эффективной
---------------------------------

Следующий барьер - умение сделать команду эффективной.
Вы не сможете быть эффективным в изоляции, поскольку ваша эффективность определяется качеством кодовой базы, а она разрабатывается всей командой.
Или вы сделаете команду эффективной, или ваша эффективность так и останется мечтательством.
Опять же, лучший наставник в этих вопросах - Kent Beck:

- "Extreme Programming Explained" 1st edition by Kent Beck

На данном этапе, этой книги достаточно.
Обратите внимание, я советую именно первое издание, так как оно лучше раскрывает смысл и назначение `Agile разработки <https://emacsway.github.io/ru/it/agile/easily-about-agile-way-to-rapid-development/>`__.


Изучаем операционную систему
----------------------------

Вот по операционным системам я мало что могу посоветовать, так как низкоуровневым программированием я практически не занимался.
Но вам обязательно нужно получить представление о том, как работают регистры процессора, память, и как управлять операционной системой.

Я в свое время читал эти книги (к сожалению, сегодня они устарели):

- "The Linux® Kernel Primer: A Top-Down Approach for x86 and PowerPC Architectures" by Claudia Salzberg Rodriguez, Gordon Fischer, Steven Smolski
- "Digital computers and microprocessors" by Aliyev / "Цифровая вычислительная техника и микропроцессоры" М.М.Алиев

А вот этот справочник у меня всегда под рукой:

- "Unix and Linux System Administration Handbook" 5th edition by Evi Nemeth, Garth Snyder, Trent R. Hein, Ben Whaley, Dan Mackin


Изучаем основы алгоритмов и структур данных
-------------------------------------------

Алгоритмы хоть и используются редко в прикладной разработке (если вы только не пишете поисковые системы, системные утилиты, языки программирования и операционные системы, системы маршрутизации, биржевые анализаторы и т.п.), но знать хотя бы базовые основы необходимо.
Существует книга, которая за двести с небольшим страниц может дать эти базовые основы в легкой и популярной форме:

- "Algorithms Unlocked" 3d edition by Thomas H. Cormen

Данная книга не акцентируется на математике, что, с одной стороны, облегчает освоение материала, но, с другой стороны, оставляет невосполненным важный аспект профессиональных знаний.
К счастью, существует книга, которая обеспечивает легкий вход в алгоритмы, включая их математический анализ:

- "Introduction to the Design and Analysis of Algorithms" 3d edition by A.Levitin

При чтении этой книги могут возникать вопросы справочного характера по математике, ответы на которые можно найти в приложении этой книги (Appendix A: Useful Formulas for the Analysis of Algorithms, Appendix B: Short Tutorial on Recurrence Relations), в математических справочниках (например, М.Я. Выгодского, А.А. Гусака) или в справочном разделе по математике "VIII Appendix: Mathematical Background" книги "Introduction to Algorithms" 3d edition by Thomas H. Cormen, Charles E. Leiserson, Ronald L. Rivest, Clifford Stein.

В качестве минималистичного ликбеза по теоретическим основам может неплохо подойти книга:

- "Computer Science Distilled" by Wladston Ferreira Filho

Она содержит минималистичные основы математики (логика, комбинаторика, вероятность), алгоритмы и структуры данных, основы Баз Данных (RDBMS, NoSQL), описание Парадигм Программирования и основы архитектуры железа.


Изучаем математику
------------------

Существует монументальная книга, которую стоит упомянуть отдельно (обратите внимание на фамилии авторов, которые в представлении не нуждаются).
Чтобы не тормозить общий процесс обучения, ее лучше читать в параллельно-фоновом режиме.
К тому же математические знания следует всегда поддерживать в актуальном состоянии, и регулярно освежать их в голове в фоновом режиме.

- "Concrete Mathematics: A Foundation for Computer Science" 2nd edition by Ronald L. Graham, Donald E. Knuth, Oren Patashnik

Эта книга дает прекрасную математическую базу для функционального программирования.
И хорошо заходит в сочетании с "The Art Of Computer Programming" Volume 1 3d edition by Donald Knuth, поскольку у них многие темы пересекаются и раскрываются с разных точек зрения, что дает полноту понимания.
Справочник математических нотаций в конце книги нередко оказывается полезным.

Книги по математике и алгоритмам - сложные книги, и я хотел бы поделиться одним советом, который я услышал еще в студенчестве.
Если что-то непонятно - прочитай три раза:

1. Первый раз просто прочитай, оставив попытки что-то понять, - нужно просто получить обзорность материала.
2. Второй раз прочитай уже пытаясь слегка вникать.
3. И третий раз прочитай уже вникая полностью.


Учимся архитектуре
------------------

Теперь можно приступить и к архитектуре:

- "Clean Architecture: A Craftsman's Guide to Software Structure and Design" by Robert C. Martin


Изучаем распределенные системы
------------------------------

- "NoSQL Distilled. A Brief Guide to the Emerging World of Polyglot Persistence." by Pramod J. Sadalage, Martin Fowler
- "Building Microservices. Designing Fine-Grained Systems" 2nd edition by Sam Newman
- "`A plain english introduction to CAP Theorem <http://ksat.me/a-plain-english-introduction-to-cap-theorem>`__" (`Russian <https://habr.com/ru/post/130577/>`__) by Kaushik Sathupadi
- "`Map Reduce: A really simple introduction <http://ksat.me/map-reduce-a-really-simple-introduction-kloudo>`__" by Kaushik Sathupadi
- "`Eventually Consistent - Revisited <https://www.allthingsdistributed.com/2008/12/eventually_consistent.html>`__" by Werner Vogels
- "`Distributed systems: for fun and profit <http://book.mixu.net/distsys/>`__" (2013). An introduction to distributed systems. (`source code <https://github.com/mixu/distsysbook>`__)
- "`Lecture notes (PDF) (including exercises) <https://martin.kleppmann.com/2020/11/18/distributed-systems-and-elliptic-curves.html>`__" by Martin Kleppmann (`download <https://www.cl.cam.ac.uk/teaching/2021/ConcDisSys/dist-sys-notes.pdf>`__, `source code <https://github.com/ept/dist-sys>`__, `video <https://www.youtube.com/playlist?list=PLeKd45zvjcDFUEv_ohr_HdUFe97RItdiB>`__)
- "`Literature references for "Designing Data-Intensive Applications" <https://github.com/ept/ddia-references>`__" by Martin Kleppmann
- "`Resources and community around CRDT technology - papers, blog posts, code and more. <https://crdt.tech/>`__" by Martin Kleppmann (`source code <https://github.com/ept/crdt-website>`__)


Изучаем распределенные системы. Углубляем навыки.
-------------------------------------------------

Книг по этой теме предстоит прочитать слишком много.
Вряд-ли ваша работа будет ждать, пока вы прочитаете их все.
К счастью, сообщество .NET разработчиков подготовило краткий справочник, который заменит вам прочтение десятка книг:

- "`.NET Microservices: Architecture for Containerized .NET Applications <https://docs.microsoft.com/en-us/dotnet/standard/microservices-architecture/index>`__" edition v2.2.1 (`mirror <https://aka.ms/microservicesebook>`__) by Cesar de la Torre, Bill Wagner, Mike Rousos

К этой книге существует эталонное приложение, которое наглядно демонстрирует практическое применение изложенной в книге информации:

- https://github.com/dotnet-architecture/eShopOnContainers (CQRS, DDD, Microservices)

Еще одно хорошее краткое руководство от Microsoft:

- "`Building microservices on Azure <https://docs.microsoft.com/en-us/azure/architecture/microservices/>`__"

И можно сюда включить еще и книгу:

- "`CQRS Journey <https://docs.microsoft.com/en-US/previous-versions/msp-n-p/jj554200(v=pandp.10)>`__" by Dominic Betts, Julián Domínguez, Grigori Melnik, Fernando Simonazzi, Mani Subramanian

К ней также существует демонстрационное приложение:

- https://github.com/microsoftarchive/cqrs-journey (Event Sourcing, SAGA transactions)


Изучаем DDD
-----------

Начинать я рекомендовал бы с прекрасного краткого руководства:

- "`What Is Domain-Driven Design? <https://www.oreilly.com/library/view/what-is-domain-driven/9781492057802/>`__" by Vladik Khononov

Или с более новой книги этого же автора:

- "Learning Domain-Driven Design: Aligning Software Architecture and Business Strategy" 1st Edition by Vlad Khononov

Затем приступить к классике:

- "Domain-Driven Design" by Eric Evans
- "`Implementing Domain-Driven Design <https://kalele.io/books/>`__" by Vaughn Vernon

Существуют краткие изложения этих двух книг по DDD.

Краткие изложения "Domain-Driven Design" by Eric Evans:

- "`Domain-Driven Design Reference <https://domainlanguage.com/ddd/reference/>`__" by Eric Evans
- "`Domain-Driven Design Quickly <https://www.infoq.com/books/domain-driven-design-quickly/>`__"

Краткое изложение "Implementing Domain-Driven Design" by Vaughn Vernon:

- "`Domain-Driven Design Distilled <https://kalele.io/books/>`__" by V.Vernon


Статьи на частые вопросы по DDD
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

- `Patterns related to Domain Driven Design <https://martinfowler.com/tags/domain%20driven%20design.html>`__ by Martin Fowler


Aggregate & Domain Modeling
~~~~~~~~~~~~~~~~~~~~~~~~~~~

- "`What is domain logic? <https://enterprisecraftsmanship.com/posts/what-is-domain-logic/>`__" by Vladimir Khorikov
- "`Domain services vs Application services <https://enterprisecraftsmanship.com/posts/domain-vs-application-services/>`__" by Vladimir Khorikov
- "`Domain model isolation <https://enterprisecraftsmanship.com/posts/domain-model-isolation/>`__" by Vladimir Khorikov
- "`Email uniqueness as an aggregate invariant <https://enterprisecraftsmanship.com/posts/email-uniqueness-as-aggregate-invariant/>`__" by Vladimir Khorikov
- "`How to know if your Domain model is properly isolated? <https://enterprisecraftsmanship.com/posts/how-to-know-if-your-domain-model-is-properly-isolated/>`__" by Vladimir Khorikov
- "`Domain model purity vs. domain model completeness <https://enterprisecraftsmanship.com/posts/domain-model-purity-completeness/>`__" by Vladimir Khorikov
- "`Domain model purity and lazy loading <https://enterprisecraftsmanship.com/posts/domain-model-purity-lazy-loading/>`__" by Vladimir Khorikov
- "`Domain model purity and the current time <https://enterprisecraftsmanship.com/posts/domain-model-purity-current-time/>`__" by Vladimir Khorikov
- "`Immutable architecture <https://enterprisecraftsmanship.com/posts/immutable-architecture/>`__" by Vladimir Khorikov
- "`Link to an aggregate: reference or Id? <https://enterprisecraftsmanship.com/posts/link-to-an-aggregate-reference-or-id/>`__" by Vladimir Khorikov

- "`How to create fully encapsulated Domain Models <https://udidahan.com/2008/02/29/how-to-create-fully-encapsulated-domain-models/>`__" by Udi Dahan

- "`Effective Aggregate Design <https://dddcommunity.org/library/vernon_2011/>`__" by Vaughn Vernon

- "`Designing a Domain Model to enforce No Duplicate Names <https://github.com/ardalis/DDD-NoDuplicates>`__ by Steve Smith


CQRS & Event Sourcing
~~~~~~~~~~~~~~~~~~~~~

- "`Overselling Event Sourcing <https://zimarev.com/blog/event-sourcing/myth-busting/2020-07-09-overselling-event-sourcing/>`__" by Alexey Zimarev
- "`Event Sourcing and Microservices <https://zimarev.com/blog/event-sourcing/microservices/>`__" by Alexey Zimarev
- "`Projections in Event Sourcing <https://zimarev.com/blog/event-sourcing/projections/>`__" by Alexey Zimarev
- "`Event Sourcing and CQRS <https://zimarev.com/blog/event-sourcing/cqrs/>`__" by Alexey Zimarev
- "`Entities as event streams <https://zimarev.com/blog/event-sourcing/entities-as-streams/>`__" by Alexey Zimarev
- "`Event Sourcing basics <https://zimarev.com/blog/event-sourcing/introduction/>`__" by Alexey Zimarev
- "`What is Event Sourcing? <https://eventstore.com/blog/what-is-event-sourcing/>`__" by Alexey Zimarev
- "`Event Sourcing and CQRS <https://eventstore.com/blog/event-sourcing-and-cqrs/>`__" by Alexey Zimarev

- "`CQRS, Task Based UIs, Event Sourcing agh! <http://codebetter.com/gregyoung/2010/02/16/cqrs-task-based-uis-event-sourcing-agh/>`__" by Greg Young
- "`CQRS Documents <https://cqrs.files.wordpress.com/2010/11/cqrs_documents.pdf>`__" by Greg Young
- "`Versioning in an Event Sourced System <https://leanpub.com/esversioning>`__" by Greg Young ("`читать online <https://leanpub.com/esversioning/read>`__", "`конспект книги <https://github.com/luque/Notes--Versioning-Event-Sourced-System>`__")
- "`Clarified CQRS <http://udidahan.com/2009/12/09/clarified-cqrs/>`__" by Udi Dahan
- "`Busting some CQRS myths <https://lostechies.com/jimmybogard/2012/08/22/busting-some-cqrs-myths/>`__" by Jimmy Bogard


Bounded Context and Microservices
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

- "`Bounded Contexts are NOT Microservices <https://vladikk.com/2018/01/21/bounded-contexts-vs-microservices/>`__" by Vladik Khononov
- "`Tackling Complexity in Microservices <https://vladikk.com/2018/02/28/microservices/>`__" by Vladik Khononov
- "`DDDDD: Bounded Contexts, Microservices, and Everything In Between <https://youtu.be/Z0RgR9xIQE4>`__" by Vladik Khononov

- "`Reactive Microservices <https://kalele.io/reactive-microservices/>`__" by Vaughn Vernon
- "`Microservices and [Micro]services <https://kalele.io/microservices-and-microservices/>`__" by Vaughn Vernon

- "`About Bounded Contexts and Microservices <https://blog.avanscoperta.it/2020/06/11/about-bounded-contexts-and-microservices/>`__" by Alberto Brandolini

- "`Using domain analysis to model microservices <https://docs.microsoft.com/en-us/azure/architecture/microservices/model/domain-analysis>`__"
- "`Identifying microservice boundaries <https://docs.microsoft.com/en-us/azure/architecture/microservices/model/microservice-boundaries>`__"

- "`Domain, Subdomain, Bounded Context, Problem/Solution Space in DDD: Clearly Defined <https://medium.com/nick-tune-tech-strategy-blog/domains-subdomain-problem-solution-space-in-ddd-clearly-defined-e0b49c7b586c>`__" by Nick Tune

- "`Monolith -> Services: Theory & Practice <https://medium.com/@kentbeck_7670/monolith-services-theory-practice-617e4546a879>`__" by Kent Beck

- "`How to break a Monolith into Microservices :: Go Macro First, then Micro <https://martinfowler.com/articles/break-monolith-into-microservices.html#GoMacroFirstThenMicro>`__" by Zhamak Dehghani


Domain Events
~~~~~~~~~~~~~

- ":doc:`/emacsway/it/ddd/tactical-design/domain-model/domain-events/domain-events-in-ddd`"


API-Design
~~~~~~~~~~

- "`Designing APIs for microservices <https://docs.microsoft.com/en-us/azure/architecture/microservices/design/api-design>`__"
- "`Web API design <https://docs.microsoft.com/en-us/azure/architecture/best-practices/api-design>`__"
- "`Web API implementation <https://docs.microsoft.com/en-us/azure/architecture/best-practices/api-implementation>`__"
- "`Microsoft REST API Guidelines <https://github.com/Microsoft/api-guidelines>`__"
- "`Microsoft Graph API <https://docs.microsoft.com/en-us/graph/query-parameters#filter-parameter>`__"
- "`OData protocol <https://docs.oasis-open.org/odata/odata/v4.0/errata03/os/complete/part2-url-conventions/odata-v4.0-errata03-os-part2-url-conventions-complete.html#_Toc453752358>`__"
- "`Google REST API Guidelines <https://google.aip.dev/general>`__"
- "`Microservice API Patterns <https://microservice-api-patterns.org/>`__"
- "`Good Practices for Capability URLs <https://w3ctag.github.io/capability-urls/>`__", W3C Draft
- "`Web API Design - Crafting Interfaces that Developers Love <https://pages.apigee.com/rs/apigee/images/api-design-ebook-2012-03.pdf>`__"
- "`REST vs. GraphQL: A Critical Review <https://goodapi.co/blog/rest-vs-graphql>`__"
- "`5 reasons you shouldn’t be using GraphQL <https://blog.logrocket.com/5-reasons-you-shouldnt-be-using-graphql-61c7846e7ed3/?gi=f67074d77004>`__" (`перевод на Русский <https://medium.com/devschacht/esteban-herrera-5-reasons-you-shouldnt-use-graphql-bae94ab105bc>`__)
- "`OpenAPIs <https://www.openapis.org/>`__"
- "`AsyncAPI <https://www.asyncapi.com/>`__"
- "`Resource Query Language (RQL) <http://www.persvr.org/rql/>`__"
- "`JSON:API <https://jsonapi.org/>`__"
- "`JSONPath specification - XPath for JSON <https://goessner.net/articles/JsonPath/>`__", "`Introduction to JsonPath <https://www.baeldung.com/guide-to-jayway-jsonpath>`__"
- "`Falcor <https://netflix.github.io/falcor/starter/what-is-falcor.html>`__"
- "`Cheat Sheet a.k.a. API Design Heuristics <https://microservice-api-patterns.org/cheatsheet>`__" - шпаргалка по "Microservices API Patterns"
- "`REST API Design - Resource Modeling  <https://www.thoughtworks.com/insights/blog/rest-api-design-resource-modeling>`__" by Prakash Subramaniam, WhoughtWorks
- "`CQRS and REST: the perfect match <https://lostechies.com/jimmybogard/2016/06/01/cqrs-and-rest-the-perfect-match/>`__" by Jimmy Bogard
- "`Entities aren’t resources, resources aren’t representations <https://lostechies.com/jimmybogard/2016/05/12/entities-arent-resources-resources-arent-representations/>`__" by Jimmy Bogard



Event Storming
~~~~~~~~~~~~~~


By Alberto Brandolini (`twitter <https://twitter.com/ziobrando>`__):

- "Domain-Driven Design: The First 15 Years", chapter "Discovering Bounded Contexts with EventStorming" by Alberto Brandolini
- "`Introducing Event Storming <http://ziobrando.blogspot.com/2013/11/introducing-event-storming.html>`__" by Alberto Brandolini
- "`Remote EventStorming <https://blog.avanscoperta.it/2020/03/26/remote-eventstorming/>`__" by Alberto Brandolini
- "`EventStorming in COVID-19 times <https://blog.avanscoperta.it/2020/03/26/eventstorming-in-covid-19-times/>`__" by Alberto Brandolini
- "`Leanpub: Introducing EventStorming <https://leanpub.com/introducing_eventstorming>`__" by Alberto Brandolini
- `EventStorming.com <https://www.eventstorming.com/>`__


Others:

- "Domain-Driven Design Distilled" by Vaughn Vernon, chapter "Chapter 7 Acceleration and Management Tools :: Event Storming"
- "`What is Domain-Driven Design? <https://www.oreilly.com/library/view/what-is-domain-driven/9781492057802/>`__" by Vladik Khononov, chapter "Chapter 8: Event Storming"
- "`EventStorming Glossary & Cheat sheet <https://ddd-crew.github.io/eventstorming-glossary-cheat-sheet/>`__" by Nick Tune
- "Open Agile Architecture", chapter "`19. Event Storming <https://pubs.opengroup.org/architecture/o-aa-standard/>`__"
- "`Event Storming на практических кейсах <http://agilemindset.ru/event-storming-%D0%BD%D0%B0-%D0%BF%D1%80%D0%B0%D0%BA%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D1%85-%D0%BA%D0%B5%D0%B9%D1%81%D0%B0%D1%85/>`__", Сергей Баранов (`видео <https://www.youtube.com/watch?v=kJjuTuviZ-E>`__)
- "`Reactive in practice, Unit 1: Event storming the stock trader domain <https://developer.ibm.com/tutorials/reactive-in-practice-1/>`__" by Kevin Webber, Dana Harrington
- "`Event storming at ibm.com <https://www.ibm.com/cloud/architecture/architecture/practices/event-storming-methodology-architecture/>`__"
- "`Event driven solution implementation methodology <https://ibm-cloud-architecture.github.io/refarch-eda/methodology/event-storming/>`__"
- "`Awesome EventStorming <https://github.com/mariuszgil/awesome-eventstorming>`__"


Tools:

- `EventStorming для PlantUML <https://github.com/tmorin/plantuml-libs/blob/master/dist/eventstorming/README.md>`__
- `miro.com <https://miro.com/>`__, см. `Event Storming template <https://miro.com/miroverse/category/ideation-and-brainstorming/event-storming>`__
- EventStorming для ArchiMate:
    - см. "`Figure 13: Event Storming Model <https://nicea.nic.in/download-files.php?nid=247>`__"
    - `Model used by Jean-Baptiste Sarrodie for presentation "Enterprise Architecture Modelling with ArchiMate in an Agile at Scale Programme" <https://community.opengroup.org/archimate-user-community/home/-/issues/8>`__


Modelling
~~~~~~~~~

- "`Getting started with DDD. Definitions of DDD and fundamental concepts to reduce the learning curve and confusion. <https://github.com/ddd-crew/welcome-to-ddd>`__" by Nick Tune & DDD-Crew
- "`Domain-Driven Design Starter Modelling Process. If you're new to DDD and not sure where to start, this process will guide you step-by-step. <https://github.com/ddd-crew/ddd-starter-modelling-process>`__" by Nick Tune & DDD-Crew
- "`Legacy Architecture Modernisation With Strategic Domain-Driven Design <https://medium.com/nick-tune-tech-strategy-blog/legacy-architecture-modernisation-with-strategic-domain-driven-design-3e7c05bb383f>`__" by Nick Tune


Собственно, этих знаний достаточно для того, чтобы стать зрелым специалистом.
Своего рода - кандидатский минимум.
Далее - порядок чтения может быть произвольным.
Читать весь список необязательно.


Дополнительная литература (на выбор)
====================================


SDLS
----


Single-Team Agile
^^^^^^^^^^^^^^^^^

- "Extreme Programming Explained" 2nd edition by Kent Beck
- "Planning Extreme Programming" by Kent Beck, Martin Fowler
- "More Effective Agile: A Roadmap for Software Leaders" by Steve McConnell
- "Clean Agile: Back to Basics" by Robert C. Martin
- "Agile! The Good, the Hype and the Ugly" by Bertrand Meyer
- "Scrum and XP from the Trenches: How We Do Scrum" 2nd edition by Henrik Kniberg
- "Essential Scrum: A Practical Guide to the Most Popular Agile Process" by Kenneth Rubin
- "Succeeding with Agile: Software Development Using Scrum" by Mike Cohn
- "User Stories Applied: For Agile Software Development" by Mike Cohn


Scaled Agile
^^^^^^^^^^^^

- "Scaling Software Agility: Best Practices for Large Enterprises" by Dean Leffingwell
- "Agile Software Requirements: Lean Requirements Practices for Teams, Programs, and the Enterprise" by Dean Leffingwell
- "SAFe® 5.0: The World’s Leading Framework for Business Agility" by Richard Knaster, Dean Leffingwell
- "Large-Scale Scrum: More with LeSS" by Craig Larman
- "`LeSS <https://less.works/less/framework/introduction>`__" (`перевод на Русский <https://less.works/ru/less/framework/introduction>`__)
- "`Agile Practice Guide <https://www.pmi.org/pmbok-guide-standards/practice-guides/agile>`__" by Project Management Institute


Стандарты
^^^^^^^^^

- "ISO/IEC/IEEE 12207:2017 Systems and software engineering — Software life cycle processes"
- "ISO/IEC/IEEE 15288:2015 Systems and software engineering — System life cycle processes"
- "ISO/IEC/IEEE 29148:2011 Systems and software engineering — Life cycle processes — Requirements engineering"
- "ISO/IEC/IEEE 15289:2019 Systems and software engineering — Content of life-cycle information items (documentation)"

- "ISO/IEC/IEEE 24765:2017 Systems and software engineering — Vocabulary"
- "ISO 9000:2005 Quality management systems — Fundamentals and vocabulary"

- "ISO/IEC 33001:2015 Information technology — Process assessment — Concepts and terminology"

- "ГОСТ Р ИСО/МЭК 12207-2010 Информационная технология. Системная и программная инженерия. Процессы жизненного цикла программных средств."
- "ГОСТ Р 57193-2016 Системная и программная инженерия. Процессы жизненного цикла систем."


Менеджмент
----------

- "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr.
- "`Systems Thinking <https://less.works/less/principles/systems-thinking.html>`__" by Craig Larman (`перевод на Русский <https://less.works/ru/less/principles/systems-thinking.html>`__)
- "Art of Project Management" by Scott Berkun
- "Менеджмент: Учебник для вузов." 3-е изд. Глухов В. В.
- "Оргуправленческое мышление. Идеология, методология, технология" / Щедровицкий Георгий Петрович
- "`Managing Digital Concepts and Practices <https://publications.opengroup.org/g183>`__"


Развитие личностных профессиональных качеств
--------------------------------------------

- "The Pragmatic Programmer: From Journeyman to Master" 1st edition by David Thomas, Andrew Hunt
- "The Pragmatic Programmer: your journey to mastery, 20th Anniversary Edition" 2nd edition by David Thomas, Andrew Hunt
- "A Mind for Numbers: How to Excel at Math and Science" by Barbara Ann Oakley
- "Systems Thinking. Quality Software Management. New York: Dorset House." by Gerald M. Weinberg, 1992,  ISBN: 0932633226
- "An Introduction to General Systems Thinking" by Gerald M. Weinberg
- "Becoming a Technical Leader" by Gerald M. Weinberg
- "Harvard Business Review on Decision Making" by Harvard Business School Press
- "The Software Architect Elevator: Redefining the Architect's Role in the Digital Enterprise 1st Edition" by Gregor Hohpe
- "Fundamentals of Software Architecture: An Engineering Approach" 1st edition by Mark Richards, Neal Ford
- "Software Architecture: The Hard Parts: Modern Trade-Off Analyses for Distributed Architectures" 1st Edition by Neal Ford, Mark Richards, Pramod Sadalage, Zhamak Dehghani
- "The Book: 37 Things One Architect Knows About IT Transformation" by Gregor Hohpe
- "Eat or Be Eaten!: Jungle Warfare for the Corporate Master Politician" by Phil Porter
- "Presentation patterns: techniques for crafting better presentations" by Neal Ford, Matthew McCullough, Nathaniel Schutta
- "Technology Strategy Patterns: Architecture as Strategy" 1st edition by Eben Hewitt
- "Thinking in Systems: A Primer" by Donella H. Meadows, Diana Wright
- "Social psychology" 13th edition by David G. Myers. Перевод: "Социальная психология" / Майерс Д. Пер. с англ. З. Замчук; Зав. ред. кол. Л. Винокуров. — 7-е изд. — СПб.: Питер, 2006.
- "Never split the difference: negotiating as if life depended on it" by Chris Voss. Перевод: "Договориться не проблема. Как добиваться своего без конфликтов и ненужных уступок." / Крис Восс
- "Искусство спора. О теории и практике спора." / Поварнин С.И.
- "Эристика, или Искусство побеждать в спорах" / Шопенгауэр Артур. English: "The Art of Being Right: 38 Ways to Win an Argument" by Arthur Schopenhauer
- "Как читать книги" / Поварнин С.И.
- "`Искусство спора (обучающие материалы) <https://ruxpert.ru/%D0%98%D1%81%D0%BA%D1%83%D1%81%D1%81%D1%82%D0%B2%D0%BE_%D1%81%D0%BF%D0%BE%D1%80%D0%B0_(%D0%BE%D0%B1%D1%83%D1%87%D0%B0%D1%8E%D1%89%D0%B8%D0%B5_%D0%BC%D0%B0%D1%82%D0%B5%D1%80%D0%B8%D0%B0%D0%BB%D1%8B)>`__"
- "`Книги по риторике <https://m.vk.com/wall-56611080_127534>`__"

Простой и доходчивый видеокурс по SoftSkills:

- "`Soft Skills Pro <https://youtube.com/channel/UCSN7G8syJUaRiXrw1l0qk_g>`__"


Базы данных
-----------

- "Mastering PostgreSQL In Application Development" by Dimitri Fontaine
- "The Art of PostgreSQL" 2nd edition by Dimitri Fontaine - is the new title of "Mastering PostgreSQL in Application Development"
- "SQL Antipatterns. Avoiding the Pitfalls of Database Programming." by Bill Karwin
- "Refactoring Databases. Evolutionary Database Design" by Scott J Ambler and Pramod J. Sadalage
- "An Introduction to Database Systems" by C.J. Date
- "PostgreSQL 10 High Performance" by Ibrar Ahmed, Gregory Smith, Enrico Pirozzi

PostgresPro представил `три книги <https://postgrespro.ru/education/books>`__ для трех разных уровней подготовленности читателей, от совершенно неосведомленного человека до разработчика баз данных.
Книги дают комплексные знания в лаконичной форме.
Все книги доступны для скачивания в свободном доступе:

1. "`Postgres: первое знакомство <https://postgrespro.ru/education/books/introbook>`__" / Л.П. Вениаминович, Р.Е. Валерьевич, Л.И. Викторович
2. "`PostgreSQL. Основы языка SQL: учеб. пособие <https://postgrespro.ru/education/books/sqlprimer>`__"  / Е.П. Моргунов; под ред. Е.В. Рогова, П.В. Лузанова.
3. "`Основы технологий баз данных: учеб. пособие <https://postgrespro.ru/education/books/dbtech>`__" / Б. А. Новиков, Е. А. Горшкова, Н. Г. Графеева; под ред. Е. В. Рогова.

Так же доступны `учебные материалы курсов <https://postgrespro.ru/education/courses>`__: слайды, видео, руководства. Скачать можно все материалы каждого курса одним архивом.

`Видеозаписи курсов <https://postgrespro.ru/education/where>`__.

Превосходная подборка статей с фундаментальной информацией простым языком о внутреннем устройстве PostgreSQL, от разработчиков PostgresPro:

- `MVCC-1. Изоляция <https://m.habr.com/ru/company/postgrespro/blog/442804/>`__
- `WAL в PostgreSQL: 1. Буферный кеш <https://m.habr.com/ru/company/postgrespro/blog/458186/>`__

Шпаргалка по выбору типа хранилища данных:

- "`Understand data store models <https://docs.microsoft.com/en-us/azure/architecture/guide/technology-choices/data-store-comparison>`__"
- "`Select an Azure data store for your application <https://docs.microsoft.com/en-us/azure/architecture/guide/technology-choices/data-store-decision-tree>`__"

Jepsen's analysis over two dozen databases, coordination services, and queues—and we’ve found replica divergence, data loss, stale reads, read skew, lock conflicts, and much more:

- "`Analyses <https://jepsen.io/analyses>`__"
- "`Everything Tagged "Jepsen" <https://aphyr.com/tags/jepsen>`__"

Рейтинг хранилищ данных:

- "`DB-Engines Ranking <https://db-engines.com/en/ranking>`__"


Изучаем распределенные системы. Третий заход.
---------------------------------------------

- "Enterprise Integration Patterns: Designing, Building, and Deploying Messaging Solutions" by Gregor Hohpe, Bobby Woolf
- "Service Design Patterns: Fundamental Design Solutions for SOAP/WSDL and RESTful Web Services" by Robert Daigneau
- "Microsoft .NET: Architecting Applications for the Enterprise" 2nd edition by Dino Esposito, Andrea Saltarello
- "`Cloud Design Patterns <https://docs.microsoft.com/en-us/azure/architecture/patterns/>`__"
- "`Cloud Design Patterns. Prescriptive architecture guidance for cloud applications <https://docs.microsoft.com/en-us/previous-versions/msp-n-p/dn568099(v=pandp.10)>`__" by Alex Homer, John Sharp, Larry Brader, Masashi Narumoto, Trent Swanson. (`Code Samples <http://aka.ms/cloud-design-patterns-sample>`__)
- "`Build Microservices on Azure <https://docs.microsoft.com/en-us/azure/architecture/microservices>`__" by Microsoft Corporation and community
- "`Cloud Best Practices <https://docs.microsoft.com/en-us/azure/architecture/best-practices/>`__" by Microsoft Corporation and community
- "`Performance Antipatterns <https://docs.microsoft.com/en-us/azure/architecture/antipatterns>`__" by Microsoft Corporation and community
- "`Azure Application Architecture Guide <https://docs.microsoft.com/en-us/azure/architecture/guide/>`__" by Microsoft Corporation and community
- "`Azure Data Architecture Guide <https://docs.microsoft.com/en-us/azure/architecture/data-guide/>`__" by Microsoft Corporation and community
- "Release It! Design and Deploy Production-Ready Software" 2nd edition by Michael Nygard
- "`Microservices Patterns: With examples in Java <https://www.manning.com/books/microservice-patterns>`__" 1st edition by Chris Richardson (`more info <https://microservices.io/book>`__)
- "Monolith to Microservices Evolutionary Patterns to Transform Your Monolith" by Sam Newman
- "Microservices AntiPatterns and Pitfalls" by Mark Richards
- "Microservices vs. Service-Oriented Architecture" by Mark Richards
- "`Site Reliability Engineering: How Google runs production systems <https://landing.google.com/sre/books/>`__" edited by Betsy Beyer, Chris Jones, Jennifer Petoff & Niall Richard Murphy
- "`The Site Reliability Workbook: Practical Ways to Implement SRE. <https://landing.google.com/sre/books/>`__" by Betsy Beyer, Niall Richard Murphy, David K. Rensin, Kent Kawahara & Stephen Thorne
- "`Building Secure & Reliable Systems: Best Practices for Designing, Implementing and Maintaining Systems. <https://landing.google.com/sre/books/>`__" by Heather Adkins, Betsy Beyer, Paul Blankinship, Ana Oprea, Piotr Lewandowski, Adam Stubblefield
- "Database Reliability Engineering. Designing and Operating Resilient Database Systems." by Laine Campbell and Charity Majors
- "Designing Data-Intensive Applications. The Big Ideas Behind Reliable, Scalable, and Maintainable Systems" by Martin Kleppmann
- "Database Internals: A Deep Dive into How Distributed Data Systems Work" by Alex Petrov
- "`Distributed systems: principles and paradigms <https://www.distributed-systems.net/index.php/books/ds3/>`__" 3d edition by Andrew S. Tanenbaum, Maarten Van Steen
- "`Введение в распределенные вычисления <http://books.ifmo.ru/file/pdf/1551.pdf>`__" / Косяков М. С. — СПб: НИУ ИТМО, 2014. — С. 75-82. — 155 с.
- "Service-Oriented Architecture Analysis and Design for Services and Microservices" by Thomas Erl
- "Workflow patterns: the definitive guide" by Aalst, Wil van der, Russell, Nick, Ter Hofstede, Arthur
- "Real-Life BPMN (4th edition): Includes an introduction to DMN" by Jakob Freund, Bernd Rücker
- "Practical Process Automation" by Bernd Ruecker


API-Design
----------

- "REST in Practice: Hypermedia and Systems Architecture" by Savas Parastatidis, Jim Webber, Ian Robinson
- "RESTful Web APIs: Services for a Changing World" by Leonard Richardson, Sam Ruby, Mike Amundsen
- "Web API Design Crafting Interfaces that Developers Love" by Brian Mulloy
- "REST API Design Rulebook" by Mark Massé
- "Principles of Web API Design: Delivering Value with APIs and Microservices" by James Higginbotham


Streaming Processing
--------------------

- "Streaming Data: Understanding the real-time pipeline" 1st edition by Andrew Psaltis
- "Big Data: Principles and best practices of scalable realtime data systems" 1st edition by Nathan Marz, James Warren
- "Kafka Streams in Action: Real-time apps and microservices with the Kafka Streams API" 1st edition by Bill Bejeck 
- "The Enterprise Big Data Lake: Delivering the Promise of Big Data and Data Science" 1st edition by Alex Gorelik


Углубляем DDD
-------------

- "Reactive Messaging Patterns with the Actor Model: Applications and Integration in Scala and Akka" by Vaughn Vernon
- "Patterns, Principles, and Practices of Domain-Driven Design" by Scott Millett, Nick Tune
- "Hands-On Domain-Driven Design with .NET Core: Tackling complexity in the heart of software by putting DDD principles into practice" by Alexey Zimarev
- "Balancing Coupling in Software Design: Successful Software Architecture in General and Distributed Systems" by Vladislav Khononov
- "`The Addison-Wesley Signature Series: Vaughn Vernon <https://www.informit.com/imprint/series_detail.aspx?ser=7937178>`__"
- "`Event Sourced Building Blocks for Domain-Driven Design with Python <https://leanpub.com/dddwithpython>`__" by John Bywater


Изучаем проектирование
----------------------

- "Agile Software Development. Principles, Patterns, and Practices." by Robert C. Martin, James W. Newkirk, Robert S. Koss
- "Analysis Patterns. Reusable Object Models" by Martin Fowler
- "Implementation Patterns" by Kent Beck
- "Smalltalk Best Practice Patterns" by Kent Beck
- "`Development of Further Patterns of Enterprise Application Architecture <https://martinfowler.com/eaaDev/>`__" by Martin Fowler
- "Domain Specific Languages" by Martin Fowler (with Rebecca Parsons)
- "Pattern Hatching: Design Patterns Applied" by John Vlissides
- "`Microsoft Application Architecture Guide <https://docs.microsoft.com/en-us/previous-versions/msp-n-p/ff650706(v=pandp.10)?redirectedfrom=MSDN>`__" 2nd edition (Patterns & Practices) by Microsoft Corporation (J.D. Meier, David Hill, Alex Homer, Jason Taylor, Prashant Bansode, Lonnie Wall, Rob Boucher Jr., Akshay Bogawat)
- "Applying UML and Patterns: An Introduction to Object-Oriented Analysis and Design and Iterative Development" by Craig Larman
- "Object-Oriented Software Construction" 2nd edition by Bertrand Meyer
- "Working Effectively with Legacy Code" by Michael C. Feathers
- "Refactoring To Patterns" by Joshua Kerievsky
- "Structure and Interpretation of Computer Programs" (aka SICP) 2nd edition (MIT Electrical Engineering and Computer Science) by Harold Abelson, Gerald Jay Sussman, Julie Sussman
- "Object Oriented Software Engineering: A Use Case Driven Approach" by Ivar Jacobson
- "Object-Oriented Analysis and Design with Applications" 3rd edition by Grady Booch, Robert A. Maksimchuk, Michael W. Engle, Bobbi J. Young Ph.D., Jim Conallen, Kelli A. Houston


POSA
----

- "Pattern-Oriented Software Architecture: A System of Patterns, Volume 1" by Frank Buschmann, Regine Meunier, Hans Rohnert, Peter Sommerlad, Michael Stal
- "Pattern-Oriented Software Architecture: Patterns for Concurrent and Networked Objects, Volume 2" by Douglas C. Schmidt, Michael Stal, Hans Rohnert, Frank Buschmann
- "Pattern-Oriented Software Architecture: Patterns for Resource Management, Volume 3" by Michael Kircher, Prashant Jain
- "Pattern-Oriented Software Architecture: A Pattern Language for Distributed Computing, Volume 4" by Frank Buschmann, Kevin Henney, Douglas C. Schmidt
- "Pattern-Oriented Software Architecture: On Patterns and Pattern Languages, Volume 5" by Frank Buschmann, Kevin Henney, Douglas C. Schmidt


Алгоритмы. Второй заход.
------------------------

- "Introduction to Algorithms" 3d edition by Thomas H. Cormen, Charles E. Leiserson, Ronald L. Rivest, Clifford Stein
- "Algorithms and Data Structures" (Oberon version: August 2004) by N.Wirth

Donald E. Knuth:

- "The Art of Computer Programming, Volume 1: Fundamental Algorithms" 3d edition by Donald Knuth
- "The Art of Computer Programming, Volume 1, Fascicle 1: MMIX; A RISC Computer for the New Millennium" 1st edition by Donald Knuth
- "The Art of Computer Programming, Volume 2, Seminumerical Algorithms" 3rd edition by Donald E. Knuth
- "The Art of Computer Programming, Volume 3, Sorting and Searching" 2nd edition by Donald E. Knuth
- "The Art of Computer Programming, Volume 4, Fascicle 0: Introduction to Combinatorial Algorithms and Boolean Functions" 1st edition by Donald E. Knuth
- "The Art of Computer Programming, Volume 4, Fascicle 1: Bitwise Tricks & Techniques; Binary Decision Diagrams" 1st edition by Donald E. Knuth
- "The Art of Computer Programming, Volume 4, Fascicle 2: Generating All Tuples and Permutations" 1st edition by Donald E. Knuth
- "The Art of Computer Programming, Volume 4, Fascicle 3: Generating All Combinations and Partitions Paperback" 1st edition by Donald E. Knuth
- "Art of Computer Programming, Volume 4, Fascicle 4: Generating All Trees; History of Combinatorial Generation 1st edition by Donald E. Knuth
- "The Art of Computer Programming" Volume 4, Fascicle 5: Mathematical Preliminaries Redux; Introduction to Backtracking; Dancing Links" 1st edition by Donald E. Knuth
- "The Art of Computer Programming, Volume 4, Fascicle 6: Satisfiability" 1st edition by Donald E. Knuth
- "The Art of Computer Programming, Volume 4A, Combinatorial Algorithms, Part 1" 1st edition by Donald E. Knuth

Хорошая подборка книг по алгоритмам: http://e-maxx.ru/bookz/


Тестирование
------------

- "xUnit Test Patterns. Refactoring Test Code." by Gerard Meszaros
- "Unit Testing Principles, Practices, and Patterns: Effective testing styles, patterns, and reliable automation for unit testing, mocking, and integration testing with examples in C#" 1st Edition by Vladimir Khorikov
- "Growing Object-Oriented Software, Guided by Tests" by Steve Freeman, Nat Pryce
- "Agile Testing: A Practical Guide for Testers and Agile Teams" by Lisa Crispin, Janet Gregory
- "More Agile Testing: Learning Journeys for the Whole Team" by Lisa Crispin, Janet Gregory
- "ATDD by Example: A Practical Guide to Acceptance Test-Driven Development" by Markus Gärtner
- "Continuous Delivery: Reliable Software Releases through Build, Test, and Deployment Automation" by Jez Humble, David Farley
- "Continuous Integration: Improving Software Quality and Reducing Risk" by Paul M. Duvall, Steve Matyas, Andrew Glover

- "`ISTQB® Related Books <https://www.istqb.org/references/books/istqb-related-books.html>`__"
- "`Referenced Books in ISTQB® Syllabi <https://www.istqb.org/references/books/referenced-books-in-istqb-syllabi.html>`__"


Компиляторы
-----------

- "Compiler Construction" by N.Wirth
- "Compilers: Principles, Techniques, and Tools" 2nd edition by Alfred V. Aho, Monica S. Lam, Ravi Sethi, Jeffrey D. Ullman


Архитектура
-----------

- "Software Architecture in Practice" 4th edition by Len Bass, Paul Clements, Rick Kazman
- "Documenting Software Architectures: Views and Beyond" 2nd edition by Paul Clements, Felix Bachmann, Len Bass, David Garlan, James Ivers, Reed Little, Paulo Merson, Robert Nord, Judith Stafford
- "Software Systems Architecture: Working With Stakeholders Using Viewpoints and Perspectives" 2nd edition by Nick Rozanski, Eóin Woods
- "Designing Software Architectures: A Practical Approach (SEI Series in Software Engineering)" 1st edition by Humberto Cervantes, Rick Kazman
- "Fundamentals of Software Architecture: An Engineering Approach" 1st edition by Mark Richards, Neal Ford
- "Introduction to Solution Architecture Paperback" by Alan McSweeney
- "Systems Analysis and Design" 7th edition by Alan Dennis, Barbara Haley Wixom, Roberta M. Roth
- "The Design of Design: Essays from a Computer Scientist" by Frederick P. Brooks
- "Living Documentation: Continuous Knowledge Sharing by Design" by Cyrille Martraire
- "Just Enough Software Architecture: A Risk-Driven Approach" by George H. Fairbanks
- "The Book: 37 Things One Architect Knows About IT Transformation" by Gregor Hohpe
- "The Software Architect Elevator: Redefining the Architect's Role in the Digital Enterprise 1st Edition" by Gregor Hohpe
- "Cloud Strategy: A Decision-based Approach to Successful Cloud Migration" by Gregor Hohpe, Michele Danieli, Jean-Francois Landreau, Tahir Hashmi
- "Architecting for Scale" 2nd Edition by Lee Atchison
- "Software Engineering: A Practitioner's Approach" 9th edition by Roger S. Pressman, Bruce Maxim
- "Presentation patterns: techniques for crafting better presentations" by Neal Ford, Matthew McCullough, Nathaniel Schutta
- "Team Topologies: Organizing Business and Technology Teams for Fast Flow" by Matthew Skelton, Manuel Pais
- "Technology Strategy Patterns: Architecture as Strategy" 1st edition by Eben Hewitt

Архитектура в Agile:

- "Building Evolutionary Architectures: Support Constant Change" 1st Edition by Neal Ford, Rebecca Parsons, Patrick Kua
- "Agile Software Architecture: Aligning Agile Processes and Software Architectures" by Muhammad Ali Babar, Alan W. Brown, Kai Koskimies, Ivan Mistrík
- "Continuous Architecture: Sustainable Architecture in an Agile and Cloud-Centric World" by Murat Erder, Pierre Pureur


Стандарты:

- "`Open Agile Architecture: A Standard of The Open Group <https://pubs.opengroup.org/architecture/o-aa-standard/>`__"
- "`ISO/IEC/IEEE 42010:2011(en) Systems and software engineering — Architecture description <https://www.iso.org/standard/50508.html>`__"
- "`ISO/IEC/IEEE 42020:2019 Software, systems and enterprise — Architecture processes <https://www.iso.org/standard/68982.html>`__"
- "`ISO/IEC/IEEE 42030:2019 Software, systems and enterprise — Architecture evaluation framework <https://www.iso.org/standard/73436.html>`__"
- `ГОСТ Р 57100-2016 Системная и программная инженерия. Описание архитектуры <https://allgosts.ru/35/080/gost_r_57100-2016>`__

Рейтинг инструментов для упраления требованиями/архитектурой/SDLC/etc. от Gartner по категориям: "`Reviews Organized by Markets <https://www.gartner.com/reviews/markets>`__".


Аналитика
---------

- "Software Requirements (Developer Best Practices)" 3rd Edition by Karl Wiegers
- "INCOSE Guide for Writing Requirements" by INCOSE
- "`Systems engineering handbook. A guide for System Life Cycle Processes and activities. <https://www.incose.org/products-and-publications/se-handbook>`__" by INCOSE
- "Managing Software Requirements: A Unified Approach" 1st edition by Dean Leffingwell, Don Widrig
- "Managing Software Requirements (paperback): A Use Case Approach" 2d Edition by Dean Leffingwell, Don Widrig
- "Requirements Engineering Fundamentals: A Study Guide for the Certified Professional for Requirements Engineering Exam - Foundation Level - IREB compliant" 2nd Edition by Klaus Pohl, Chris Rupp

- "`A Guide to the Business Analysis Body of Knowledge (BABOK®) <https://www.iiba.org/career-resources/a-business-analysis-professionals-foundation-for-success/babok/>`__"
- "`Guide to the Systems Engineering Body of Knowledge (SEBoK) <https://www.sebokwiki.org/wiki/Download_SEBoK_PDF>`__"

- "`Library of IREB artifacts <https://www.ireb.org/en/downloads/tag:handbook>`__"
- "`Handbook for the CPRE Foundation Level according to the IREB Standard Education and Training for Certified Professional for Requirements Engineering (CPRE) Foundation Level <https://www.ireb.org/content/downloads/5-cpre-foundation-level-handbook/cpre_foundationlevel_handbook_en_v1.0.pdf>`__" Version 1.0.0
- "`Handbook of Advanced Level Elicitation according to the IREB Standard Education and Training for IREB Certified Professional for Requirements Engineering Advanced Level Advanced Level Elicitation <https://www.ireb.org/content/downloads/13-cpre-advanced-level-elicitation-handbook/advanced_level_elicitation_handbook_en.pdf>`__" Version 1.0.3
- "`Requirements Management according to the IREB Standard Education and Training for the IREB Certified Professional for Requirements Engineering Qualification Advanced Level Requirements Management <https://www.ireb.org/content/downloads/16-cpre-advanced-level-requirements-management-handbook/ireb-cpre-handbook-for-requirements-management-advanced-level-en-v1.1.pdf>`__" Version 1.1.0
- "`Handbook of Requirements Modeling According to the IREB Standard Education and Training for IREB Certified Professional for Requirements Engineering Advanced Level Requirements Modeling <https://www.ireb.org/content/downloads/19-handbook-cpre-advanced-level-requirements-modeling/ireb_cpre_handbook_requirements-modeling_advanced-level-v1.3.pdf>`__" Version 1.3
- "`A Glossary of Requirements Engineering Terminology <https://www.ireb.org/content/downloads/1-cpre-glossary/ireb_cpre_glossary_17.pdf>`__" Version 1.7
- "`A  Glossary  of Requirements Engineering Terminology Caution: This glossary is aligned to the CPRE Foundation Level syllabus 3.0 only! <https://www.ireb.org/content/downloads/2-cpre-glossary-2-0/ireb_cpre_glossary_en_2.0.pdf>`__" Version 2.0.0


Аналитика в Agile:

- "Agile Software Requirements: Lean Requirements Practices for Teams, Programs, and the Enterprise" by Dean Leffingwell
- Whitepaper "`A Lean and Scalable Requirements Information Model for the Agile Enterprise <https://scalingsoftwareagility.files.wordpress.com/2007/03/a-lean-and-scalable-requirements-information-model-for-agile-enterprises-pdf.pdf>`__" by Dean Leffingwell with Juha‐Markus Aalto
- "`An Agile Architectural Epic Kanban System: Part 2 – The Model <https://scalingsoftwareagility.wordpress.com/2010/03/05/an-agile-architectural-epic-kanban-system-part-2-%E2%80%93-the-model/>`__" by Dean Leffingwell

- "`Agile Extension to the BABOK® Guide <https://www.iiba.org/career-resources/business-analysis-resources/iiba-bookstore/>`__"

- "`Handbook of RE@Agile According to the IREB Standard Education and Training for IREB Certified Professional for Requirements Engineering Advanced Level RE@Agile <https://www.ireb.org/content/downloads/22-cpre-advanced-level-re-agile-handbook/handbook_cpre_al_re%40agile_en_v1.0.2.pdf>`__" Version 1.0.2
- "`IREB Certified Professional for Requirements Engineering RE@Agile Glossary <https://www.ireb.org/content/downloads/34-re-agile-glossary/ireb_cpre_re%40agile_glossary_v1.0.5.pdf>`__"


Другие подборки литературы по аналитике:

- `Литература по аналитике на сайте Systems.Education <https://systems.education/books>`__
- `Литература по аналитике на сайте Volere Requirements Resources <https://www.volere.org/resources/books/>`__.

Смотрите также список инструментов для управления требованиями:

- `Tools <https://www.volere.org/tools/>`__ on Volere Requirements Resources
- `Requirements Tools <https://www.volere.org/requirements-tools/>`__ on Volere Requirements Resources


Изучаем оценивание задач
------------------------

- "Software Estimation: Demystifying the Black Art (Developer Best Practices)" by Steve McConnell (я встречал в интернете `краткий конспект <http://igorshevchenko.ru/blog/entries/software-estimation>`__)
- "Agile Estimating and Planning" by Mike Cohn


Функциональное программирование
-------------------------------

- `"Software architecture: object-oriented vs functional <http://se.ethz.ch/~meyer/publications/functional/meyer_functional_oo.pdf>`__" by Bertrand Meyer
- "`Category Theory for Programmers <https://bartoszmilewski.com/2014/10/28/category-theory-for-programmers-the-preface/>`__" by Bartosz Milewski (`unofficial PDF and LaTeX source <https://github.com/hmemcpy/milewski-ctfp-pdf>`__)
- "`Domain Modeling Made Functional. Tackle Software Complexity with Domain-Driven Design and F# <https://fsharpforfunandprofit.com/books/>`__" by Scott Wlaschin
- "`F# for Fun and Profit <https://fsharpforfunandprofit.com/>`__" by Scott Wlaschin
- "Functional Programming for the Object-Oriented Programmer" by Brian Marick
- "Functional Thinking" by Neal Ford
- "`Haskell <https://en.wikibooks.org/wiki/Haskell>`__"
- "`The Science of Functional Programming. A Tutorial, with Examples in Scala. <https://github.com/winitzki/sofp/blob/master/sofp-src/sofp.pdf>`__" by Sergei Winitzki, Ph.D.
- "Microservices with Clojure. Develop event-driven, scalable, and reactive microservices with real-time monitoring" by Anuj Kumar

Для Golang-разработчиков:

- "Learning Functional Programming in Go: Change the way you approach your applications using functional programming in Go" by Lex Sheehan


Справочники
-----------

- "Computing Handbook. Computer Science and Software Engineering." 3d edition by Allen Tucker, Teofilo Gonzalez, Jorge Diaz-Herrera


Справочная информация
=====================


Body of Knowledge
-----------------

- "`Guide to the Systems Engineering Body of Knowledge (SEBoK) <https://www.sebokwiki.org/wiki/Download_SEBoK_PDF>`__"
- "`The Information Technology Architecture Body of Knowledge (ITABoK) <https://itabok.iasaglobal.org/>`__"
- "`The Enterprise Architecture Body of Knowledge (EABoK) <https://www.mitre.org/publications/technical-papers/guide-to-the-evolving-enterprise-architecture-body-of-knowledge>`__"
- "`MITRE Systems Engineering Guide <https://www.mitre.org/publications/technical-papers/the-mitre-systems-engineering-guide>`__
- "`A Guide to the Business Architecture Body of Knowledge(R) (BIZBOK(R) Guide) <https://www.businessarchitectureguild.org/page/BIZBOK>`__"
- "`A Guide to the Business Analysis Body of Knowledge (BABOK®) <https://www.iiba.org/career-resources/a-business-analysis-professionals-foundation-for-success/babok/>`__"
- "`Agile Extension to the BABOK® Guide <https://www.iiba.org/career-resources/business-analysis-resources/iiba-bookstore/>`__"
- "`DAMA-DMBOK: Data Management Body of Knowledge <https://www.dama.org/content/what-data-management>`__" 2nd edition by DAMA International
- "`The Project Management Body of Knowledge (PMBoK) <https://www.pmi.org/pmbok-guide-standards/foundational/pmbok>`__" by Project Management Institute (PMI)
- "`Agile Practice Guide <https://www.pmi.org/pmbok-guide-standards/practice-guides/agile>`__" by Project Management Institute (PMI), 2017
- "`Учебник 4CIO. Настольная книга ИТ-Директора <https://book4cio.ru/>`__"
- "`Учебник 4CDTO. Настольная книга руководителя цифровой трансформации <https://4cio.ru/pages/570>`__"

- "`Systems engineering handbook. A guide for System Life Cycle Processes and activities. <https://www.incose.org/products-and-publications/se-handbook>`__" by INCOSE
- "`The global skills and competency framework for a digital world <https://sfia-online.org/en>`__" by SFIA Foundation
- "`List of Bodies of Knowledge <https://sfia-online.org/en/tools-and-resources/bodies-of-knowledge/list-of-bodies-of-knowledge>`__" by SFIA Foundation


ГОСТы
-----

- "`База ГОСТов allgosts.ru - 35. ИНФОРМАЦИОННЫЕ ТЕХНОЛОГИИ. МАШИНЫ КОНТОРСКИЕ <https://allgosts.ru/35/>`__"
- "`StandartGOST.ru - бесплатные ГОСТы и магазин документов. Информационные технологии. Машины конторские <https://standartgost.ru/0/753-informatsionnye_tehnologii_mashiny_kontorskie>`__"


Online-каталоги
---------------

- `Catalog of Refactorings <http://www.refactoring.com/catalog/>`__
- `Code Smell <http://c2.com/cgi/wiki?CodeSmell>`__
- `Anti Patterns Catalog <http://c2.com/cgi/wiki?AntiPatternsCatalog>`__
- `Catalog of Patterns of Enterprise Application Architecture <https://martinfowler.com/eaaCatalog/>`__
- `List of DSL Patterns <https://www.martinfowler.com/dslCatalog/>`__
- `Enterprise Integration Patterns <http://www.enterpriseintegrationpatterns.com/>`__ (`шпаргалка по EIP <https://www.enterpriseintegrationpatterns.com/download/EIPTutorialReferenceChart.pdf>`__)
- `DDD and Messaging Architectures <https://verraes.net/2019/05/ddd-msg-arch/>`__ - an overview of different series on patterns in distributed systems by Mathias Verraes.
- `Service Design Patterns <http://servicedesignpatterns.com/>`__
- `SOAPatterns.org <https://patterns.arcitura.com/soa-patterns>`__
- `CloudPatterns.org <https://patterns.arcitura.com/cloud-computing-patterns>`__
- `BigDataPatterns.org <https://patterns.arcitura.com/big-data-patterns>`__
- `Cloud Design Patterns | Microsoft Docs <https://docs.microsoft.com/en-us/azure/architecture/patterns/>`__
- `Workflow Patterns <http://workflowpatterns.com/patterns/>`__
- `Microservices Patterns <https://microservices.io/patterns/>`__
- `Microservices Patterns (book) <https://www.manning.com/books/microservice-patterns>`__
- `Microservices Patterns from Sam Newman <https://samnewman.io/patterns/>`__
- `About DDD on the site of Ward Cunningham <http://ddd.fed.wiki.org/>`__
- `Refactoring Databases <http://www.databaserefactoring.com/>`__
- `XUnit Test Patterns <http://xunitpatterns.com/>`__
- `Refactoring Databases <https://databaserefactoring.com/>`__
- `Catalog of Database Refactorings <http://www.agiledata.org/essays/databaseRefactoringCatalog.html>`__
- `Extreme Programming Rules <http://www.extremeprogramming.org/rules.html>`__
- `Consistency Models - a clickable map <https://jepsen.io/consistency>`__
- `Subway Map to Agile Practices - a clickable map <https://www.agilealliance.org/agile101/subway-map-to-agile-practices/>`__
- `The Arcitura Education Patterns, Mechanisms and Metrics Master Catalog <https://patterns.arcitura.com/>`__
- `Microservice API Patterns <https://microservice-api-patterns.org/>`__
- `OpenAPIs <https://www.openapis.org/>`__
- `AsyncAPI <https://www.asyncapi.com/>`__
- `Architecture Playbook <https://nocomplexity.com/documents/arplaybook/>`__ (`source <https://github.com/nocomplexity/ArchitecturePlaybook>`__)
- `Software Systems Architecture <https://www.viewpoints-and-perspectives.info/>`__ - This web site contains a selection of supporting material for the book ("Software Systems Architecture: Working With Stakeholders Using Viewpoints and Perspectives" 2nd edition by Nick Rozanski, Eóin Woods), including sample chapters, references and white papers.


Code Smell catalogs
-------------------

- Chapter 17: "Smells and Heuristics" of the book "Clean Code: A Handbook of Agile Software Craftsmanship" by Robert C. Martin
- Chapter 3. "Bad Smells in Code" of the book "Refactoring: Improving the Design of Existing Code" by Martin Fowler, Kent Beck, John Brant, William Opdyke, Don Roberts
- `Code Smell <http://c2.com/cgi/wiki?CodeSmell>`__ catalog on the site of Ward Cunningham
- "Refactoring To Patterns" by Joshua Kerievsky


Другие подборки литературы
--------------------------

- `Awesome lists <https://github.com/sindresorhus/awesome>`__
- `Awesome Domain-Driven Design <https://github.com/heynickc/awesome-ddd>`__
- `Domain Driven Design in Python, Ruby and other dynamic languages resources <https://github.com/valignatev/ddd-dynamic>`__
- `Awesome Microservices <https://github.com/mfornos/awesome-microservices>`__
- `Solution Architecture links, articles, books, video lessons, etc. <https://github.com/unlight/solution-architecture>`__
- `Awesome Algorithms <https://github.com/tayllan/awesome-algorithms>`__
- `Awesome Algorithms Education <https://github.com/gaerae/awesome-algorithms-education>`__
- `The System Design Primer <https://github.com/donnemartin/system-design-primer>`__ - Learn how to design large-scale systems. Prep for the system design interview.
- `List of awesome university courses for learning Computer Science <https://github.com/prakhar1989/awesome-courses>`__
- `MAXimal :: bookz - электронные версии различных книг по алгоритмам <http://e-maxx.ru/bookz/>`__
- `Programming and design learning resources by Kamil Grzybek <http://www.kamilgrzybek.com/programming-and-design-resources/>`__
- `Список книг от Сергея Теплякова <https://sergeyteplyakov.blogspot.com/2013/08/blog-post.html>`__
- `Список книг от Grady Booch <https://handbookofsoftwarearchitecture.com/books/>`__
- `Книги по направлению Архитектура и проектирование ПО от эксперта luxoft <https://www.luxoft-training.ru/about/experts/answers/302/30945/>`__
- `The Architect’s Path (Part 1 - Model) <https://architectelevator.com/architecture/architect-path/>`__ by Gregor Hohpe
- `The Architect’s Path (Part 2 - Implementation) <https://architectelevator.com/architecture/architect-bookshelf/>`__ by Gregor Hohpe
- `Software Architecture Book References <https://www.developertoarchitect.com/books.html>`__ by Mark Richards
- `Как прокачаться в проектировании программного обеспечения — список книг <https://apolomodov.medium.com/software-design-books-743be52e4c71>`__ by Alexander Polomodov
- `Искусство спора (обучающие материалы) <https://ruxpert.ru/%D0%98%D1%81%D0%BA%D1%83%D1%81%D1%81%D1%82%D0%B2%D0%BE_%D1%81%D0%BF%D0%BE%D1%80%D0%B0_(%D0%BE%D0%B1%D1%83%D1%87%D0%B0%D1%8E%D1%89%D0%B8%D0%B5_%D0%BC%D0%B0%D1%82%D0%B5%D1%80%D0%B8%D0%B0%D0%BB%D1%8B)>`__
- `Книги по риторике <https://m.vk.com/wall-56611080_127534>`__


Почтовые рассылки и сообщества
------------------------------

- `Domain Driven Design Community <http://dddcommunity.org/>`__
- `Domain Driven Design Weekly <http://dddweekly.com/>`__
- `Microservice Weekly <https://microserviceweekly.com/>`__


.. _reference-applications-ru:

Эталонные демонстрационные приложения
=====================================

- `eShopOnContainers <https://github.com/dotnet-architecture/eShopOnContainers>`__ (CQRS, DDD, Microservices)
- `Microsoft patterns & pratices CQRS Journey sample application <https://github.com/microsoftarchive/cqrs-journey>`__ (CQRS, DDD, Event Sourcing, SAGA transactions)

..

    "A perfect example of this [you can see] if you go look at the CQRS and Event Sourcing by Microsoft Patterns and Practices, which is heavily focused on doing this inside of Azure using their toolkits."

    \- Greg Young, "`A Decade of DDD, CQRS, Event Sourcing <https://youtu.be/LDW0QWie21s?t=1092>`__" at 18:15

- `Full Modular Monolith application with Domain-Driven Design approach <https://github.com/kgrzybek/modular-monolith-with-ddd>`__ by Kamil Grzybek
- `Sample .NET Core REST API CQRS implementation with raw SQL and DDD using Clean Architecture <https://github.com/kgrzybek/sample-dotnet-core-cqrs-api>`__ by Kamil Grzybek
- `Refactoring from anemic to rich Domain Model example <https://github.com/kgrzybek/refactoring-from-anemic-to-rich-domain-model-example>`__ by Kamil Grzybek
- `Sample Bounded Contexts for C#.NET from the book "Implementing Domain-Driven Design" <https://github.com/VaughnVernon/IDDD_Samples_NET>`__ by Vaughn Vernon
- `Sample Bounded Contexts from the book "Implementing Domain-Driven Design" <https://github.com/VaughnVernon/IDDD_Samples>`__ by Vaughn Vernon
- `xoom-examples <https://github.com/vlingo/xoom-examples>`__ - the VLINGO XOOM examples demonstrating features and functionality available in the reactive components.
- Implementation of samples from the book "Domain-Driven Design" by Eric Evans in `Java <https://github.com/citerus/dddsample-core>`__, `C# <https://github.com/SzymonPobiega/DDDSample.Net>`__, `Ruby <https://github.com/paulrayner/ddd_sample_app_ruby>`__, `Golang <https://github.com/marcusolsson/goddd>`__ (`yet another Golang <https://github.com/go-kit/kit/tree/master/examples/shipping>`__). See also `the article <https://www.citerus.se/go-ddd>`__.
- `Goa <https://goa.design/>`__ provides a holistic approach for developing remote APIs and microservices in Go.
- `Simple CQRS example <https://github.com/gregoryyoung/m-r>`__ by Greg Young (приложение так же реализует Event Sourcing)
- `Greg Young's Simple CQRS in F# <https://github.com/thinkbeforecoding/m-r>`__ by Jérémie Chassaing
- `Complete serverless application to show how to apply DDD, Clean Architecture, and CQRS by practical refactoring of a Go project <https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example>`__ (`more info <https://threedots.tech/post/serverless-cloud-run-firebase-modern-go-application/>`__) by Robert Laszczak
- `Clean Monolith Shop <https://github.com/ThreeDotsLabs/monolith-microservice-shop>`__ by Robert Laszczak - Source code for `Why using Microservices or Monolith can be just a detail? <https://threedots.tech/post/microservices-or-monolith-its-detail/>`__ article
- `go-iddd - showcase project for implementing DDD in Go <https://github.com/AntonStoeckl/go-iddd>`__ by Anton Stöckl (see more info `here <https://medium.com/@TonyBologni/implementing-domain-driven-design-and-hexagonal-architecture-with-go-1-292938c0a4d4>`__ and `here <https://medium.com/@TonyBologni/implementing-domain-driven-design-and-hexagonal-architecture-with-go-2-efd432505554>`__).
- `Real-time Map <https://github.com/asynkron/realtimemap-go>`__ displays real-time positions of public transport vehicles in Helsinki. It's a showcase for `Proto.Actor <https://proto.actor/>`__ - an ultra-fast distributed actors solution for Go, C#, and Java/Kotlin. See also `realtimemap-dotnet <https://github.com/asynkron/realtimemap-dotnet>`__ implementation in .NET.
- `Demo taxi system, using eventsourcing library <https://github.com/johnbywater/es-example-taxi-demo>`__ by John Bywater
- `Example "bank accounts" application using the Python eventsourcing library <https://github.com/johnbywater/es-example-bank-accounts>`__ by John Bywater
- `Example "cargo shipping" application using the Python eventsourcing library <https://github.com/johnbywater/es-example-cargo-shipping>`__ by John Bywater
- `Examples of using eventsourcing library <https://github.com/johnbywater/eventsourcing/tree/main/eventsourcing/examples>`__ by John Bywater
- `FTGO example application. Example code for the book Microservice patterns <https://github.com/microservices-patterns/ftgo-application>`__ by Chris Richardson
- `Eventuate Tram Customers and Orders <https://github.com/eventuate-tram/eventuate-tram-examples-customers-and-orders/>`__ by Chris Richardson
- `Eventuate Tram Customers and Orders - .NET version <https://github.com/eventuate-examples/eventuate-tram-core-dotnet-examples-customers-and-orders>`__ by Chris Richardson
- `eventuate-examples <https://github.com/eventuate-examples>`__ by Chris Richardson
- `Sample code for the book Principles, Practices and Patterns of Domain-Driven Design <https://github.com/elbandit/PPPDDD>`__ by Scott Millett, Nick Tune
- `Hands-On Domain-Driven Design with .NET Core, published by Packt <https://github.com/PacktPublishing/Hands-On-Domain-Driven-Design-with-.NET-Core>`__ by Alexey Zimarev
- "`dotnet-sample <https://github.com/Eventuous/dotnet-sample>`__" - Sample application using Eventuous .NET by Alexey Zimarev
- `Extended code samples related to the book "Domain Modeling Made Functional" <https://github.com/swlaschin/DomainModelingMadeFunctional>`__ by Scott Wlaschin
- `Railway-Oriented-Programming-Example <https://github.com/swlaschin/Railway-Oriented-Programming-Example>`__ by Scott Wlaschin
- `Order Taking Service <https://github.com/andorp/order-taking>`__ - Idris version of Domain Modeling Made Functional Book.
- `DDD with Actors <https://github.com/VaughnVernon/DDDwithActors>`__ by Vaughn Vernon
- `The examples for the book "Reactive Messaging Patterns with the Actor Model" <https://github.com/VaughnVernon/ReactiveMessagingPatterns_ActorModel>`__ by Vaughn Vernon
- `A Stock Trader system to demonstrate reactive systems development <https://github.com/VaughnVernon/reactive-stock-trader>`__ (`source <https://github.com/RedElastic/reactive-stock-trader>`__ by RedElastic)
- `ContosoUniversityDotNetCore-Pages <https://github.com/jbogard/ContosoUniversityDotNetCore-Pages>`__ by Jimmy Bogard

- `RedElastic: reactive-stock-trader <https://github.com/RedElastic/reactive-stock-trader>`__ - A reference architecture for stock trading to demonstrate the concepts of reactive systems development. Based on the original Stock Trader by IBM and implemented with Lagom by Lightbend. "`Reactive in practice: A complete guide to event-driven systems development in Java. <https://developer.ibm.com/series/reactive-in-practice/>`__"

- `IBM Stock Trader <https://github.com/IBMStockTrader>`__ - Org containing a repository per microservice in the IBM Stock Trader cloud-native sample application. "`Introduction to the IBM Stock Trader sample. <https://developer.ibm.com/blogs/introducing-stocktrader/>`__"

- `Refactoring from Anemic Domain Model Towards a Rich One <https://github.com/vkhorikov/AnemicDomainModel>`__ by Vladimir Khorikov
- `DDD in Practice <https://github.com/vkhorikov/DddInAction>`__ by Vladimir Khorikov
- `DDD and EF Core <https://github.com/vkhorikov/DddAndEFCore>`__ by Vladimir Khorikov
- `CQRS in Practice <https://github.com/vkhorikov/CqrsInPractice>`__ by Vladimir Khorikov
- `Applying Functional Principles in C# <https://github.com/vkhorikov/FuntionalPrinciplesCsharp>`__ by Vladimir Khorikov
- `Specification Pattern in C# <https://github.com/vkhorikov/SpecPattern>`__ by Vladimir Khorikov
- `Specification pattern implementation in C# <https://github.com/vkhorikov/SpecificationPattern>`__ by Vladimir Khorikov
- `Validation in DDD course <https://github.com/vkhorikov/ValidationInDDD>`__ by Vladimir Khorikov

Варианты реализации OO/Functional Aggregates на примере Reference Applications by Chris Richardson:

- `Traditional OO mutable Domain Objects <https://github.com/cer/event-sourcing-examples/tree/master/java-spring>`__
- `Functional Scala witn immutable Domain Objects <https://github.com/cer/event-sourcing-using-scala-typeclasses>`__
- `Hybrid OO/Functional Scala with immutable Domain Objects <https://github.com/cer/event-sourcing-examples/tree/master/scala-spring>`__

Others:

- `DDD Sample Projects <https://github.com/heynickc/awesome-ddd#sample-projects>`__

..
    - "Rapid Development: Taming Wild Software Schedules" by Steve McConnell
    - "The Definitive Guide to MongoDB" by David Hows, Peter Membrey, Eelco Plugge, Tim Hawkins
    - "High Performance MySQL" by Baron Schwartz, Peter Zaitsev, and Vadim Tkachenko
    - "PostgreSQL: Up and Running" by Regina Obe and Leo Hsu

