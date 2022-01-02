:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: Requirements; nonfunctional in Agile
   :name: emacsway-agile-nonfunctional-requirements


================================
Agile nonfunctional Requirements
================================

.. sectionauthor:: Ivan Zakrevsky

Если вы знакомы с ATAM, ADD, QAW, то вы знакомы со сценариями атрибутов качеств ("Software Architecture in Practice" 3d edition by Len Bass, Paul Clements, Rick Kazman, глава 16), и легко поймете, о чем пишет K.Wigers - один из топовых авторов по требованиям и аналитике:

    📝 "**Nonfunctional requirements need to have priority alongside user stories; you can’t defer their implementation until a later iteration.**
    It’s possible to specify quality attributes in the form of stories: "As a help desk technician, I want the knowledge base to respond to queries **within five seconds** so the customer doesn’t get frustrated and hang up."

    <...>

    Developers need to keep nonfunctional requirements in mind as they consider the implications of implementing individual user stories.
    As more functionality is added through a series of iterations, the system’s efficiency and hence performance can deteriorate.
    Specify performance goals and begin performance testing **with early iterations**, so you can become aware of concerns early enough to take corrective actions.

    <...>

    As with user stories, it’s possible to write acceptance tests for quality attributes.
    This is a way to quantify the quality attributes.
    If a performance goal is stated simply as “The knowledge base must return search results quickly,” you can’t write tests to define what constitutes “quickly.” A better acceptance test would be: "Keyword search of the knowledge base **takes less than 5 seconds**, and preferably **less than 3 seconds**, to return a result.

    <...>

    Part of **accepting an iteration as being complete** is to assess whether the pertinent nonfunctional requirements are satisfied.
    Often there is a range of acceptable performance, with some outcomes more desirable than others.
    As it does for any other software development approach, satisfying quality requirements can distinguish delight from disappointment on agile projects."

    -- Глава 14, подраздел "Handling quality attributes on agile projects", книги "Software Requirements" 3rd Edition, Wiegers K., Beatty J.

Приоритезация и реализация нефункциональных требований совместно (alongside) с функциональными выбивает почву для появления технических задач, оторванных от бизнес-целей.
А поскольку один из топовых авторов по Scrum, K.Rubin, утверждает, что нефункциональными требованиями должен заниматься Product Owner, то это автоматически устраняет необходимость в лоббировании технических задач техническими специалистами.

Нужно заметить, что в SAFe нефункциональными требованиями занимается уже не Product Owner, а System Architect.

    System Architect/Engineering is an individual or team that defines the overall architecture of the system. They work at a level of abstraction above the teams and components and define Non-Functional Requirements (NFRs), major system elements, subsystems, and interfaces.

    -- "SAFe® 5.0: The World’s Leading Framework for Business Agility" by Richard Knaster, Dean Leffingwell

Вообще говоря, в Scrum этим не обязательно должен заниматься Product Owner.
По офф.гайду команда должна сотрудничать со стейкхолдерами, а Product Owner может делегировать свои полномочия команде, например, аналитику или архитектору.
Тут важно не кто, а как.

Был упомянут K.Rubin, посмотрим, что именно он говорит:

    📝 "The **product owner is responsible** for defining the acceptance criteria for each product backlog item.
    These are the conditions under which the product owner would be satisfied that the functional and **nonfunctional requirements** have been met.
    The product owner may also write acceptance tests corresponding to the acceptance criteria, or he could enlist the assistance of subject matter experts (SMEs) or development team members.
    In either case, the product owner should ensure that these acceptance criteria (and frequently specific acceptance tests) are created before an item is considered at a sprint-planning meeting.
    Without them, the team would have an incomplete understanding of the item and would not be ready to include it in a sprint.
    For this reason, many Scrum teams include the existence of clear acceptance criteria as an item on their definition-of-ready checklist (see Chapter 6).

    <...>

    The team must also decide when to test all of the browsers.
    Each nonfunctional requirement is a prime target for inclusion in the team’s definition of done.
    If the team includes the “Web Browser Support” nonfunctional requirement in the definition of done, the team will have to test any new features added in the sprint with all of the listed browsers.
    If it doesn’t work with all of them, the story isn’t done.

    I recommend that teams try to include as many of the nonfunctional requirements in their definitions of done as they possibly can.
    Waiting to test nonfunctional requirements until late in the development effort defers getting fast feedback on critical system performance characteristics.

    <...>

    nonfunctional requirement.
    1. A requirement that does not relate to functionality but to attributes such as reliability, efficiency, usability, maintainability, and portability, which product backlog items must possess in order to be fully accepted by the stakeholders.
    2. **Each nonfunctional requirement is a candidate for inclusion in the definition of done**. See also definition of done."

    -- "Essential Scrum : a practical guide to the most popular agile process" by Kenneth S. Rubin.

Топовые авторы по аналитике (K.Wigers) и Scrum (K.Rubin) сошлись в одном мнении. Один дает описание с точки зрения работы с требованиями, а другой - с точки зрения гибких процессов разработки.

См. также:

- "`Nonfunctional Requirements <https://www.scaledagileframework.com/nonfunctional-requirements/>`__" at SAFe

TODO:

- Dean Leffingwell тоже дает описание работы с нефункциональными требованиями.
