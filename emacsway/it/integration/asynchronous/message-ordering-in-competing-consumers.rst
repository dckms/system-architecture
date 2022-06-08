:canonical-base-url: https://dckms.github.io/system-architecture

.. index:: Messaging, Message Broker, Causal Consistency, Eventual Consistency, Integration Patterns, NATS Streaming Server, Actor Model, Asynchronous Communication

======================================================
О гонке сообщений в условиях конкурирующих подписчиков
======================================================

.. sectionauthor:: Ivan Zakrevsky

Одной из непростых  тем в DDD и микросервисной архитектуре является т.н. **проблема "конкурирующих подписчиков"**. Это когда два причинно-зависимых события попадают на конкурирующие узлы обработки событий, и второе событие может "обогнать" первое, например, по причине того, что при обработке первого события возникли сетевые издержки, или запустился сборщик мусора, или по какой-либо причине первое сообщение не было обработано и подтверждено (ack) с первого раза. Возникает гонка сообщений.

.. contents:: Содержание

Например, `NATS использует Round-robin для балансировки подписчиков группы <https://docs.nats.io/nats-concepts/queue>`__, и там эта проблема хорошо проявляется. Партиционирование каналов `появилось <https://bravenewgeek.com/building-a-distributed-log-from-scratch-part-5-sketching-a-new-system/>`__ только в пока еще нестабильном `jetstream <https://github.com/nats-io/jetstream>`__.

    Scaling with queue subscribers

    This is ideal **if you do not rely on message order**.

    -- "`Slow Consumers - NATS Docs <https://docs.nats.io/running-a-nats-service/nats_admin/slow_consumers#handling-slow-consumers>`__"

Обходной путь:

    Create a subject namespace that can scale

    You can distribute work further through the subject namespace, with some forethought in design. This approach is useful if you need to preserve message order. The general idea is to publish to a deep subject namespace, and consume with wildcard subscriptions while giving yourself room to expand and distribute work in the future.

    For a simple example, if you have a service that receives telemetry data from IoT devices located throughout a city, you can publish to a subject namespace like Sensors.North, Sensors.South, Sensors.East and Sensors.West. Initially, you'll subscribe to Sensors.> to process everything in one consumer. As your enterprise grows and data rates exceed what one consumer can handle, you can replace your single consumer with four consuming applications to subscribe to each subject representing a smaller segment of your data. Note that your publishing applications remain untouched.

    -- "`Slow Consumers - NATS Docs <https://docs.nats.io/running-a-nats-service/nats_admin/slow_consumers#handling-slow-consumers>`__"

Еще одина возможная причина нарушения очередности обработки сообщений:

    Note: For a given subscription, messages are dispatched serially, one message at a time. If your application **does not care about processing ordering** and would prefer the messages to be dispatched concurrently, it is the application's responsibility to move them to some internal queue to be picked up by threads/go routines.

    -- "`Asynchronous Subscriptions - NATS Docs <https://docs.nats.io/using-nats/developer/receiving/async>`__"

Кроме того, доставка сообщений может пакетироваться из соображений оптимизации.

Один из примеров, который мне запомнился (с какой-то статьи) - это когда один из пользователей соц.сети удаляет из списка друзей другого пользователя, и тут же шлет оставшимся друзьям письмо, в котором дискредитирует удаленного друга. Возникает два события, первое - на удаление друга, второе - на отправку сообщения списку оставшихся друзей. Причем, второе сообщение находится в причинной зависимости от первого, и должно быть обработано после первого. Возникает гонка событий.

В условиях конкурирующих подписчиков, хронология обработки событий может измениться. И тогда, в момент отправки дискредитирующего письма списку друзей, удаленный пользователь все еще будет присутствовать в списке получателей.

Существует несколько стратегий решения этой проблемы:

1. Нивелировать побочные эффекты (устранить симптомы) от нарушения очередности событий (коммутативность).
2. Исключить причины нарушения очередности событий.
3. Восстановить очередность сообщений.
4. Восстановить очередность обработки сообщений.

Будем рассматривать каждый из вариантов поочередно в отдельных постах.

А пока - список литературы, который хорошо освещает эту проблему:

- "Designing Data-Intensive Applications. The Big Ideas Behind Reliable, Scalable, and Maintainable Systems" by Martin Kleppmann
- "`Lecture notes (PDF) (including exercises) <https://martin.kleppmann.com/2020/11/18/distributed-systems-and-elliptic-curves.html>`__" by Martin Kleppmann (`download <https://www.cl.cam.ac.uk/teaching/2021/ConcDisSys/dist-sys-notes.pdf>`__, `source code <https://github.com/ept/dist-sys>`__, `video <https://www.youtube.com/playlist?list=PLeKd45zvjcDFUEv_ohr_HdUFe97RItdiB>`__)
- "Database Internals: A Deep Dive into How Distributed Data Systems Work" by Alex Petrov
- "Distributed systems: principles and paradigms" 3d edition by Andrew S. Tanenbaum, Maarten Van Steen
- "`Введение в распределенные вычисления <http://books.ifmo.ru/file/pdf/1551.pdf>`__" / Косяков М. С. — СПб: НИУ ИТМО, 2014. — С. 75-82. — 155 с.
- "Building Event-Driven Microservices" by Adam Bellemare, "`Chapter 6. Deterministic Stream Processing <https://www.oreilly.com/library/view/building-event-driven-microservices/9781492057888/ch06.html>`__"
- "`Distributed systems: for fun and profit <http://book.mixu.net/distsys/>`__" (2013). An introduction to distributed systems. (`source code <https://github.com/mixu/distsysbook>`__)
- "Database Reliability Engineering. Designing and Operating Resilient Database Systems." by Laine Campbell and Charity Majors
- "`Event Sourced Building Blocks for Domain-Driven Design with Python <https://leanpub.com/dddwithpython>`__" by John Bywater

Статьи по теме:

- "`Don't Settle for Eventual Consistency. Stronger properties for low-latency geo-replicated storage. <https://queue.acm.org/detail.cfm?id=2610533>`__" (`pdf <https://dl.acm.org/ft_gateway.cfm?id=2610533&ftid=1449165&dwn=1>`__) by Wyatt Lloyd, Facebook; Michael J. Freedman, Princeton University; Michael Kaminsky, Intel Labs; David G. Andersen, Carnegie Mellon University
- "`Bolt-on Causal Consistency <http://www.bailis.org/papers/bolton-sigmod2013.pdf>`__" by Peter Bailis, Ali Ghodsi, Joseph M. Hellerstein†, Ion Stoica, UC Berkeley KTH/Royal Institute of Technology
- "`Detecting Causal Relationships in Distributed Computations:In Search of the Holy Grail <https://disco.ethz.ch/courses/hs08/seminar/papers/mattern4.pdf>`__" by Reinhard Schwarz, Friedemann Mattern
- "`Principles of Eventual Consistency <https://www.microsoft.com/en-us/research/publication/principles-of-eventual-consistency/>`__" (`pdf <https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/final-printversion-10-5-14.pdf>`__) by Sebastian Burckhardt, Microsoft Research
- "`HighLoad++, Михаил Тюленев (MongoDB): Causal consistency: от теории к практике <https://habr.com/ru/company/ua-hosting/blog/487638/>`__"
- "`Version Vector <https://martinfowler.com/articles/patterns-of-distributed-systems/version-vector.html>`__" by Unmesh Joshi
- "`Nobody Needs Reliable Messaging <https://www.infoq.com/articles/no-reliable-messaging/>`__" by Marc de Graauw
- "`Error Handling Patterns for Apache Kafka Applications <https://www.confluent.io/blog/error-handling-patterns-in-kafka/>`__" by Gerardo Villeda

Список литературы по интеграционным паттернам:

- "Enterprise Integration Patterns: Designing, Building, and Deploying Messaging Solutions" by Gregor Hohpe, Bobby Woolf
- "Reactive Messaging Patterns with the Actor Model: Applications and Integration in Scala and Akka" by Vaughn Vernon
- "Camel in Action" 2nd Edition by Claus Ibsen and Jonathan Anstey

Примеры интеграционных паттернов:

- https://github.com/VaughnVernon/ReactiveMessagingPatterns_ActorModel
- https://camel.apache.org/components/latest/eips/enterprise-integration-patterns.html
- https://github.com/camelinaction/camelinaction2
- https://www.enterpriseintegrationpatterns.com/patterns/messaging/

Каталог моделей согласованности:

- https://jepsen.io/consistency

Шпаргалка по EIP-паттернам:

- "`Enterprise Integration Patterns Tutorial Reference Chart <https://www.enterpriseintegrationpatterns.com/download/EIPTutorialReferenceChart.pdf>`__"

Каталоги:

- "`Cloud Design Patterns <https://docs.microsoft.com/en-us/azure/architecture/patterns/>`__"
- "`Cloud Design Patterns. Prescriptive architecture guidance for cloud applications <https://docs.microsoft.com/en-us/previous-versions/msp-n-p/dn568099(v=pandp.10)>`__" by Alex Homer, John Sharp, Larry Brader, Masashi Narumoto, Trent Swanson.

Code Samples:

- http://aka.ms/cloud-design-patterns-sample
- "`Cloud Best Practices <https://docs.microsoft.com/en-us/azure/architecture/best-practices/>`__" by Microsoft Corporation and community


Поддержка коммутативности
=========================

Первая из перечисленных стратегий решения проблемы "конкурирующих подписчиков" - это "**нивелировать побочные эффекты (устранить симптомы) от нарушения очередности событий (коммутативность)**".

Часто бывает так, что два действия подряд над одним и тем же агрегатом приводят к тому, что, в условиях конкурирующих подписчиков, сообщение второго события может обогнать сообщение первого события. Если при этом используется "**Event-Carried State Transfer**" ( https://martinfowler.com/articles/201701-event-driven.html ), то при обработке следующего сообщения (которое было отправлено первым), система будет оперировать уже устаревшими данными.

Как один из вариантов решения проблемы в таком случае, может быть переход на "**Event Notification**". В некоторых случаях прокатывает. Но он ухудшает availability (CAP-Theorem) из-за каскадного синхронного запроса.

В некоторых случаях также прокатывает игнорирование предыдущего события, если последующее событие уже было обработано.


Исключение причин нарушения очередности событий
===============================================

Вторая из перечисленных стратегий решения проблемы "конкурирующих подписчиков" - это "**исключить причины нарушения очередности событий**".

Этому способу решения проблемы посвящена глава "`3.3.5 Competing receivers and message ordering <https://livebook.manning.com/book/microservices-patterns/chapter-3/section-3-3-5?origin=product-toc>`__" книги "Microservices Patterns: With examples in Java" by Chris Richardson

Если mеssaging system не поддерживает партиционирование каналов, то его можно реализовать с помощью паттерна EIP "`Content-Based Router <https://www.enterpriseintegrationpatterns.com/patterns/messaging/ContentBasedRouter.html>`__"

Например, `используя Camel Framework <https://camel.apache.org/components/latest/eips/content-based-router-eip.html>`__.

С помощью партиционирования каналов мы добиваемся того, что все сообщения одного и того же **причинно-зависимого (causal) потока** попадают на один и тот же узел группы подписчиков. Нет конкуренции - нет проблемы. Здесь вводится новый и достаточно обширный термин "**Causal Consistency**", имеющий критически важное значение для всех, кто имеет дело с распределенными системами.

Vaughn Vernon в "Reactive Messaging Patterns with the Actor Model: Applications and Integration in Scala and Akka" (RMPwAM) ссылается на следующие две статьи по этому вопросу:

- https://queue.acm.org/detail.cfm?id=2610533
- http://www.bailis.org/papers/bolton-sigmod2013.pdf

Каталог моделей согласованности:

- https://jepsen.io/consistency

Было бы, наверное, уместно упомянуть в контексте этого обсуждения пару превосходных материалов на тему CAP-theorem и Consistency:

Самое понятное объяснение CAP-Theorem, которое я когда-либо видел:

- "`A plain english introduction to CAP Theorem <http://ksat.me/a-plain-english-introduction-to-cap-theorem>`__" by Kaushik Sathupadi (`перевод на русский <https://habr.com/ru/post/130577/>`__)

Превосходная статья от CTO of Amazon.com Werner Vogels:

- "`Eventually Consistent - Revisited <https://www.allthingsdistributed.com/2008/12/eventually_consistent.html>`__"

Превосходная статья по Causal Consistency (Causal Dependencies) доступным языком:

- "`HighLoad++, Михаил Тюленев (MongoDB): Causal consistency: от теории к практике <https://habr.com/ru/company/ua-hosting/blog/487638/>`__"


Восстановление очередности сообщений
====================================

Третья из перечисленных стратегий решения проблемы "конкурирующих подписчиков" - это "**восстановить очередность сообщений**".

    📝 "Хьюитт был против включения требований о том, что сообщения должны прибывать в том порядке, в котором они отправлены на модель актора. Если желательно упорядочить входящие сообщения, то это можно смоделировать с помощью очереди акторов, которая обеспечивает такую функциональность. Такие очереди акторов упорядочивали бы поступающие сообщения так, чтобы они были получены в порядке FIFO. В общем же случае, если актор X отправляет сообщение M1 актору Y, а затем тот же актор X отправляет другое сообщение M2 к Y, то не существует никаких требований о том, что M1 придёт к Y раньше M2."

    -- Pаздел "`Никаких требований о порядке поступления сообщений <https://ru.wikipedia.org/wiki/%D0%9C%D0%BE%D0%B4%D0%B5%D0%BB%D1%8C_%D0%B0%D0%BA%D1%82%D0%BE%D1%80%D0%BE%D0%B2#%D0%9D%D0%B8%D0%BA%D0%B0%D0%BA%D0%B8%D1%85_%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B9_%D0%BE_%D0%BF%D0%BE%D1%80%D1%8F%D0%B4%D0%BA%D0%B5_%D0%BF%D0%BE%D1%81%D1%82%D1%83%D0%BF%D0%BB%D0%B5%D0%BD%D0%B8%D1%8F_%D1%81%D0%BE%D0%BE%D0%B1%D1%89%D0%B5%D0%BD%D0%B8%D0%B9>`__" статьи "Модель акторов" Википедии

Для решения этой задачи можно использовать EIP Pattern "`Resequencer <https://www.enterpriseintegrationpatterns.com/patterns/messaging/Resequencer.html>`__". Например, `используя Camel Framework <https://camel.apache.org/components/latest/eips/resequence-eip.html>`__.


Восстановление очередности обработки сообщений
==============================================

Четвертая из перечисленных стратегий решения проблемы "конкурирующих подписчиков" - это "**восстановить очередность обработки сообщений**".

Иными словами, можно пойти другим путем, и отказаться от гарантированной очередности доставки сообщений. Но, в таком случае, подписчик сам должен будет решать, может ли он обработать поступившее сообщение, или же причинно-предшествующее сообщение еще пока не было обработано, и тогда он должен оставить поступившее сообщение в очереди. Правда, на выяснение этого требуется потратить ресурсы (где-то нужно фиксировать обработку сообщений и потом удостоверяться, что предшествующее причинное сообщение уже было обработано).

Как красиво заметил Alexey Zimarev, "мир occasionally-connected устройств по определению не упорядочен".

Такой подход применяется в Actor Model:

    📝 "... модель акторов зеркально отражает систему коммутации пакетов, которая не гарантирует, что пакеты будут получены в порядке отправления. Отсутствие гарантий порядка доставки сообщений позволяет системе коммутации пакетов буферизовать пакеты, использовать несколько путей отправки пакетов, повторно пересылать повреждённые пакеты и использовать другие методы оптимизации."

    -- Pаздел "`Никаких требований о порядке поступления сообщений <https://ru.wikipedia.org/wiki/%D0%9C%D0%BE%D0%B4%D0%B5%D0%BB%D1%8C_%D0%B0%D0%BA%D1%82%D0%BE%D1%80%D0%BE%D0%B2#%D0%9D%D0%B8%D0%BA%D0%B0%D0%BA%D0%B8%D1%85_%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B9_%D0%BE_%D0%BF%D0%BE%D1%80%D1%8F%D0%B4%D0%BA%D0%B5_%D0%BF%D0%BE%D1%81%D1%82%D1%83%D0%BF%D0%BB%D0%B5%D0%BD%D0%B8%D1%8F_%D1%81%D0%BE%D0%BE%D0%B1%D1%89%D0%B5%D0%BD%D0%B8%D0%B9>`__" статьи "Модель акторов" Википедии

..

    📝 "Messages in the Actor model are generalizations of packets in Internet computing in that they need not be received in the order sent. Not implementing the order of delivery, allows packet switching to buffer packets, use multiple paths to send packets, resend damaged packets, and to provide other optimizations.

    For example, Actors are allowed to pipeline the processing of messages. What this means is that in the course of processing a message m1, an Actor can designate how to process the next message, and then in fact begin processing another message m2 before it has finished processing m1. Just because an Actor is allowed to pipeline the processing of messages does not mean that it must pipeline the processing. Whether a message is pipelined is an engineering tradeoff."

    -- "`Actor  Model  of  Computation: Scalable  Robust  Information Systems <https://arxiv.org/abs/1008.1459>`__" by Carl Hewitt

Тут нужно сделать короткое отступление. Хотя, как говорилось ранее, "*Хьюитт был против включения требований о том, что сообщения должны прибывать в том порядке, в котором они отправлены на модель актора*", в современных реализациях Actor Model mailbox представлен как FIFO-queue:

    📝 "One of the guarantees of the Actor model is sequential message delivery. That is, by default actor mailboxes are first-in, first-out (FIFO) channels. When a message arrives through the actor's channel, it will be received in the order in which it was sent. Thus, if actor A sends a message to actor B and then actor A sends a second message to actor B, the message that was sent first will be the first message received by actor B."

Однако, вопрос все-равно остается открытым:

    📝 "What if you introduce a third actor, C? Now actor A and actor C both send one or more messages to actor B. There is no guarantee which message actor B will receive first, either the first from actor A or the first from actor C. Nevertheless, the first message from actor A will always be received by actor B before the second message that actor A sends, and the first message from actor C will always
    be received by actor B before the second message that actor C sends...

    What is implied? Actors must be prepared to accept and reject messages based on their current state, which is reflected by the order in which previous messages were received. Sometimes a latent message could be accepted even if it is not perfect timing, but the actor's reaction to the latent message may have to carefully take into account its current state beforehand. This may be dealt with more gracefully by using the actors become() capabilities."

    -- "Reactive Messaging Patterns with the Actor Model: Applications and Integration in Scala and Akka" by Vaughn Vernon, Chapter "5. Messaging Channels :: Point-to-Point Channel"

Кроме того,

    📝 "Because individual messages may follow different routes, some messages are likely to pass through the processing steps sooner than others, **resulting in the messages getting out of order**. However, some subsequent processing steps do require in-sequence processing of messages, for example to maintain referential integrity.

    One common way things get out of sequence is the fact that different messages may take different processing paths. Let's look at a simple example. Let's assume we are dealing with a numbered sequence of messages. If all even numbered messages have to undergo a special transformation whereas all odd numbered messages can be passed right through, then odd numbered messages will appear on the resulting channel while the even ones queue up at the transformation. If the transformation is quite slow, all odd messages may appear on the output channel before a single even message makes it, bringing the sequence completely out of order.

    To avoid getting the messages out of order, we could introduce a loop-back (acknowledgment) mechanism that makes sure that only one message at a time passes through the system. The next message will not be sent until the last one is done processing. This conservative approach will resolve the issue, but has two significant drawbacks. First, it can slow the system significantly. If we have a large number of parallel processing units, we would severely underutilize the processing power. In many instances, the reason for parallel processing is that we need to increase performance, so throttling traffic to one message at a time would complete erase the purpose of the solution. The second issue is that this approach requires us to have control over messages being sent into the processing units. However, often we find ourselves at the receiving end of an out-of-sequence message stream without having control over the message origin."

    -- "Enterprise Integration Patterns: Designing, Building, and Deploying Messaging Solutions" by Gregor Hohpe, Bobby Woolf

Решение?

    📝 "While not discussed in detail here, Message Metadata can be used to achieve causal consistency [`AMC-Causal Consistency <https://queue.acm.org/detail.cfm?id=2610533>`__] among Messages (130) that must be replicated across a network with full ordering preserved [`Bolt-on Causal Consistency <http://www.bailis.org/papers/bolton-sigmod2013.pdf>`__]."

    -- "Reactive Messaging Patterns with the Actor Model: Applications and Integration in Scala and Akka" by Vaughn Vernon, Chapter "10. System Management and Infrastructure :: Message Metadata/History"

..

    📝 "Even so, a technique called causal consistency [`AMC-Causal Consistency <https://queue.acm.org/detail.cfm?id=2610533>`__] can be used to achieve the same."

    -- "Reactive Messaging Patterns with the Actor Model: Applications and Integration in Scala and Akka" by Vaughn Vernon, Chapter "10. System Management and Infrastructure :: Message Journal/Store"

..

    📝 "To see the full power that results from using Domain Events , consider the concept of causal consistency. A business domain provides causal consistency if its operations that are causally related —one operation causes another—are seen by every dependent node of a distributed system in the same order [`Causal <https://queue.acm.org/detail.cfm?id=2610533>`__] . This means that causally related operations must occur in a specific order, and thus one thing cannot happen unless another thing happens before it. Perhaps this means that one Aggregate cannot be created or modified until it is clear that a specific operation occurred to another
    Aggregate."

    -- "Domain-Driven Design Distilled" by Vaughn Vernon

Посмотреть вживую `обеспечение Causal Consistency <https://eventsourcing.readthedocs.io/en/v8.3.0/topics/process.html#causal-dependencies>`__ на уровне подписчика можно в EventSourcing Framework. Реализация `здесь <https://github.com/johnbywater/eventsourcing/blob/fd73c5dbd97c0ae759c59f7bb0700afb12db7532/eventsourcing/application/process.py#L273>`__.

Собственно, Causal является промежуточным уровнем строгости согласованности, чтобы избежать строгую линеаризацию сообщений (которая часто избыточна) из соображений сохранения параллелизма и повышения производительности, но при этом, не допускать параллелизма в потоках причинно-зависимых сообщений (где очередность сообщений, действительно, востребована).

Обычно идентификатором потока (``streamId``) выступает идентификатор агрегата. А идентификатором последовательности события в этом потоке (``position``) обычно `выступает номер версии агрегата <https://github.com/johnbywater/eventsourcing/blob/fd73c5dbd97c0ae759c59f7bb0700afb12db7532/eventsourcing/application/process.py#L82>`__

Другой пример кода, реализующего Causal Store можно посмотреть в главе "6.4.2 Causal Store" статьи "`Principles of Eventual Consistency <https://www.microsoft.com/en-us/research/publication/principles-of-eventual-consistency/>`__" (`pdf <https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/final-printversion-10-5-14.pdf>`__) by Sebastian Burckhardt, Microsoft Research.

Реализация Vector Clock на Golang - `vclock <https://labix.org/vclock>`__.
Статья об этой библиотеке на сайте автора: "`Vector clock support for Go <https://blog.labix.org/2010/12/21/vector-clock-support-for-go>`__" by Gustavo Niemeyer.

    📝 "Note that just **saving the Domain Event in its causal order doesn't guarantee that it will arrive at other distributed nodes in the same order**. Thus, it is also the responsibility of the consuming Bounded Context to recognize proper **causality**. It might be the Domain Event type itself that can indicate causality, or it may be **metadata** associated with the Domain Event, such as a **sequence** or **causal identifier**. The **sequence** or **causal identifier** would **indicate what caused this Domain Event, and if the cause was not yet seen, the consumer must wait to apply the newly arrived event until its cause arrives**. In some cases it is possible to ignore latent Domain Events that have already been superseded by the actions associated with a later one; in this case causality has a dismissible impact [об этом способе уже говорилось ранее, прим. моё]."

    -- "Domain-Driven Design Distilled" by Vaughn Vernon, Chapter "6. Tactical Design with Domain Events:: Designing, Implementing, and Using Domain Events"

..

    📝 "The first option is to use message sessions, a feature of the Azure Service Bus. If you use **message sessions**, this guarantees that messages within a session are delivered in the same order that they were sent.
    The second alternative is to modify the handlers within the application to detect out-of-order messages through the use of sequence numbers or timestamps added to the messages when they are sent. **If the receiving handler detects an out-of-order message, it rejects the message and puts it back onto the queue or topic to be processed later, after it has processed the messages that were sent before the rejected message.**"

    -- "CQRS Journey" by Dominic Betts, Julián Domínguez, Grigori Melnik, Fernando Simonazzi, Mani Subramanian, Chapter "`Journey 6: Versioning Our System :: Message ordering <https://docs.microsoft.com/ru-ru/previous-versions/msp-n-p/jj591565(v=pandp.10)#message-ordering>`__"

..

    📝 "**Actors must be prepared to accept and reject messages based on their current state, which is reflected by the order in which previous messages were received.** Sometimes a latent message could be accepted even if it is not perfect timing, but the actor's reaction to the latent message may have to carefully take into account its current state beforehand. This may be dealt with more gracefully by using the actors become() capabilities."

    -- "Reactive Messaging Patterns with the Actor Model: Applications and Integration in Scala and Akka" by Vaughn Vernon, Chapter "5. Messaging Channels :: Point-to-Point Channel"

Возникает вопрос о том, нужно ли заниматься восстановлением очередности сообщений на уровне Domain Logic, или на уровне Application Logic.
В статье "`Nobody Needs Reliable Messaging <https://www.infoq.com/articles/no-reliable-messaging/>`__" by Marc de Graauw приводятся убедительные аргументы о том, что если это важно для бизнеса, то это должно быть на уровне бизнес-логики (Domain Logic).
Однако, нужно учитывать, что термина "Сообщение" в предметной области вообще не существует (есть только "Событие").
Зато существует термин "время", которое едино для всего в предметной области, в отличии от времени приложения в распределенной системе.

Таким образом, очередность доставки сообщений - это проблема, свойственная не предметной области, а приложению.
Нужно ли решать её на уровне бизнеса?
Ответ зависит от конкретных обстоятельств.

Еще один из способов решения проблемы согласованности - это дублирование данных, сохранение, обработка и передача зависимых данных атомарно.
Этот прием часто используется для обеспечения границ согласованности Aggregate в DDD, для обеспечения автономности микросервисов и Bounded Contexts.

    An implementation consistent with this model would guarantee the invariant relating PO [Purchase Order] and its items, while changes to the price of a part would not have to immediately affect the items that reference it.
    Broader consistency rules could be addressed in other ways.
    For example, the system could present a queue of items with outdated prices to the users each day, so they could update or exempt each one.
    But this is not an invariant that must be enforced at all times.
    By making the dependency of line items on parts looser, we avoid contention and reflect the realities of the business better.
    At the same time, tightening the relationship of the PO and its line items guarantees that an important business rule will be followed.

    -- "Domain-Driven Design" by Eric Evans

Родственные EIP patterns:

- "`Correlation Identifier <https://www.enterpriseintegrationpatterns.com/patterns/messaging/CorrelationIdentifier.html>`__"
- "`Message Sequence <https://www.enterpriseintegrationpatterns.com/patterns/messaging/MessageSequence.html>`__"


Применяется в том числе и в Event Sourcing.

В  метаданных eventstore есть переменные ``$causationid`` and ``$correlationid``.

    📝 "The are both really simple patterns I have never quite understood why they end up so misunderstood.
    Let's say every message has 3 ids. 1 is its id. Another is correlation the last it causation.
    The rules are quite simple. If you are responding to a message, you copy its correlation id as your correlation id, its message id is your causation id.
    This allows you to see an entire conversation (correlation id) or to see what causes what (causation id).
    Cheers,
    Greg Young"

    https://discuss.eventstore.com/t/causation-or-correlation-id/828/4

Примеры:

- `раз <https://github.com/microsoftarchive/cqrs-journey/blob/6ffd9a8c8e865a9f8209552c52fa793fbd496d1f/scripts/CreateDatabaseObjects.sql#L57-L62>`__
- `два <https://github.com/kgrzybek/modular-monolith-with-ddd/blob/4e2d66d16f97b3c863fbecd072dad52338516882/src/Modules/Payments/Infrastructure/AggregateStore/SqlStreamAggregateStore.cs#L44-L45>`__

Шпаргалка по EIP-паттернам:

- "`Enterprise Integration Patterns Tutorial Reference Chart <https://www.enterpriseintegrationpatterns.com/download/EIPTutorialReferenceChart.pdf>`__"

Но даже если подписчик всего один, и сообщения доставляются последовательно, то и тогда очередность обработки сообщений может быть нарушена. Пример из NATS Streaming Server:

    📝 "With the redelivery feature, order can't be guaranteed, since by definition server will resend messages that have not been acknowledged after a period of time. Suppose your consumer receives messages 1, 2 and 3, does not acknowledge 2. Then message 4 is produced, server sends this message to the consumer. The redelivery timer then kicks in and server will resend message 2. The consumer would see messages: 1, 2, 3, 4, 2, 5, etc...

    In conclusion, the server does not offer this guarantee although it tries to redeliver messages first thing on startup. That being said, if the durable is stalled (number of outstanding messages >= MaxInflight), then the redelivery will also be stalled, and new messages will be allowed to be sent. When the consumer resumes acking messages, then it may receive redelivered and new messages interleaved (new messages will be in order though)."

    -- nats-streaming-server, `issue #187 "Order of delivery" <https://github.com/nats-io/nats-streaming-server/issues/187#issuecomment-257024506>`__, comment by Ivan Kozlovic

Кстати, проблема очередности доставки сообщений хорошо описана в главе "Projections and Queries :: Building read models from events :: Subscriptions" книги "`Hands-On Domain-Driven Design with .NET Core: Tackling complexity in the heart of software by putting DDD principles into practice <https://www.amazon.com/Hands-Domain-Driven-Design-NET-ebook/dp/B07C5WSR9B>`__" by Alexey Zimarev. И он добавил несколько `интересных аргументов в чат канала <https://t.me/emacsway_chat/85>`__.
