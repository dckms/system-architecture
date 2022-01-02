:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: Requirements; in Agile
   :name: emacsway-agile-requirements


==================
Agile Requirements
==================

.. sectionauthor:: Ivan Zakrevsky

.. contents:: Содержание


Что такое требование
====================

Следует различать значения терминов needs, requirements и specification.
Вопросы возникают, как правило, там, где их не различают, и под требованиями зачастую понимают System Requirements Specification (SyRS), на который распространяется стандарт "ISO/IEC/IEEE 29148:2011 Systems and software engineering - Life cycle processes - Requirements engineering".

Однако, на сами требования распространяются стандарты SDLC:

- "ISO/IEC/IEEE 12207:2017 Systems and software engineering - Software life cycle processes"
- "ISO/IEC/IEEE 15288:2015 Systems and software engineering - System life cycle processes"

Итак, что такое требование:

    📝 "**requirement** - statement that translates or expresses a need and its associated constraints and conditions.
    [SOURCE: ISO/IEC/IEEE 29148:2011, modified, NOTE has been removed.]"

    -- "ISO/IEC/IEEE 12207:2017 Systems and software engineering - Software life cycle processes"

Это все. Ни больше, ни меньше.


Product Backlog Item и Requirements
===================================

Итак, какая же все-таки связь между Product Backlog Item и Requirements?

    📝 "In the Scrum framework a **product backlog** lists all of the **requirements** for a solution, including both **customer and technical requirements**."

    -- Agile Extension to the BABOK® Guide version 1 (obsolete, на момент написания статьи)

..

    📝 "**Backlog Item** - An item on the backlog which represents one or more **requirements**.

    <...>

    **User Stories** are used to convey a customer **requirement** for the delivery team."

    -- Agile Extension to the BABOK® Guide version 2 (actual, на момент написания статьи)

Актуальная версия Agile-расширения BABoK (на момент написания статьи) - вторая, хотя актуальная версия самого BABoK - третья.

    📝 "The **Product Backlog** is a list of **functional and nonfunctional requirements** that, when turned into functionality, will deliver this vision."

    -- "Agile Project Management with Scrum" by Ken Schwaber

..

    📝 "Agile projects generally maintain **requirements** in the form of user stories in a **product backlog**."

    -- "Software Requirements (Developer Best Practices)" 3rd Edition by Karl Wiegers

..

    📝 "Instead of compiling a large inventory of detailed **requirements** up front, we create placeholders for the **requirements**, called **product backlog items (PBIs)**"

    -- "Essential Scrum: A Practical Guide to the Most Popular Agile Process" by Kenneth Rubin

..

    📝 "XP- originated “**user story**” as the primary currency for expressing application **requirements**."

    📝 "**User stories** are the agile replacement for most of what has been traditionally expressed as software **requirements** statements (or use cases in RUP and UML), and they are the workhorses of agile development."

    -- "Agile Software Requirements: Lean Requirements Practices for Teams, Programs, and the Enterprise" by Dean Leffingwell

И есть у него еще такой `документ <https://scalingsoftwareagility.files.wordpress.com/2007/03/a-lean-and-scalable-requirements-information-model-for-agile-enterprises-pdf.pdf>`__, который прекрасно раскрывает связь между требованиями и PBI.

    📝 "**User Story**: agile software development practice from Extreme Programming to express **requirements** from an end user perspective, emphasizing verbal communication.
    In Scrum, it is often used to express functional items on the Product Backlog."

    -- Официальный сайт Ken Schwaber, `glossary <https://www.scrum.org/resources/professional-scrum-developer-glossary>`__

..

    📝 "The agile approach to **requirements** is based on **user stories**: units of functionality corresponding to interactions of users with the system."

    -- "Agile! The Good, the Hype and the Ugly" by Bertrand Meyer

У Mike Cohn есть прекрасная статья на тему, чем отличается User Story от других способов документирования требований, и начинается она со слов:

    📝 "Extreme programming (XP) introduced the practice of expressing **requirements** in the form of **user stories**"

    -- "`Advantages of User Stories for Requirements <https://www.mountaingoatsoftware.com/articles/advantages-of-user-stories-for-requirements>`__" by Mike Cohn


Почему User Story, а не Requirements
====================================

Kent Beck разъясняет, почему он использовал термин Story вместо Requirements.
Ключевым аргументом здесь выступает семантическое различие - требования переменны, а не константны.
А так же то, что полнота требований недостижима.

    📝 "Software development has been steered wrong by the word "requirement", defined in the dictionary as \"something mandatory or obligatory.\"
    The word carries a connotation of absolutism and permanence, inhibitors to embracing change.
    And the word \"requirement\" is just plain wrong.
    Out of one thousand pages of \"requirements\", if you deploy a system with the right 20% or 10% or even 5%, you will likely realize all of the business benefit envisioned for the whole system.
    So what were the other 80%? Not "requirements"; they weren't really mandatory or obligatory.

    **Early estimation is a key difference between stories and other requirements practices.**
    Estimation gives the business and technical perspectives a chance to interact, which creates value early, when an idea has the most potential.
    When the team knows the cost of features it can split, combine, or extend scope based on what it knows about the features' value."

    -- "Extreme Programming Explained" 2nd edition by Kent Beck

Bertrand Meyer о том, в чем отличия между User Story и Requirements.
Обратите внимание, Bertrand Meyer, как и Kent Beck, так же делает акцент на недостижимость полноты требований, и указывает на семантическое отличие термина Requirements по своему смыслу, хотя по стандарту итеративная разработка освобождается от полноты требований (и даже предназначается для её разрешения).

    📝 "Agile development accepts change.
    In software projects, full requirements cannot be determined at the beginning; needs emerge as the project develops, and evolve as customers and others try intermediate releases.
    Such change is considered a normal part of the development process.

    <...>

    The last principle gives us the second part of the replacement for requirements: use scenarios to define functionality.
    A scenario is a description of a particular interaction of a user with the system, for example (if we are building mobile phone software) a phone conversation from the time the caller dials the number to the time the two parties get disconnected.
    “Scenario” is not a common agile term, but covers variants such as use cases and user stories which differ by their level of granularity (a use case is a complete interaction, a user story an application of a smaller unit of functionality).
    Scenarios are obtained from customers and indicate the fundamental properties of the system’s functionality as seen from the user perspective.
    Collecting scenarios, usually in the form of user stories, is the principal agile technique for requirements; it differs from traditional requirements elicitation in two fundamental ways: 

    - A scenario is just one example; unlike requirements, it cannot lay claim to completeness. A set of scenarios, however large, cannot come even close to achieving this goal, in the same way that no number of tests of a program can replace a specification. 
    - In agile development, requirements are not collected at the beginning of the project but throughout, as development progresses. Note, however, that this difference is not as absolute as the agile literature suggests when it blasts “waterfall approaches”: while the traditional software engineering view presents requirements as a specific lifecycle step, coming early in the process, it does not rule out — except in the imagination of agile authors — a scheme in which the requirements are constantly updated in the rest of the lifecycle.

    <...>

    The agile approach to requirements is based on user stories: units of functionality corresponding to interactions of users with the system.

    <...>

    We note once again the confusion inherent in such agile criticism as Beck’s comment that “Requirements gathering isn’t a phase that produces a static document”, as if having a requirements phase implied that the resulting requirements document will be static.
    The two matters are separate."

    -- "Agile! The Good, the Hype and the Ugly" by Bertrand Meyer


Подведем итог: требование в условиях недостаточной полноты требований, которое может быть изменено по мере снижения уровня неопределенности, традиционно называется User Story или PBI.
В таком случае требования уточняются по мере снижения уровня неопределенности, что является базовым принципом :ref:`итеративной модели <emacsway-iterative-development>` разработки.


.. index:: Literature

Литература про Agile-requirements
=================================

- "`Handbook of RE@Agile According to the IREB Standard Education and Training for IREB Certified Professional for Requirements Engineering Advanced Level RE@Agile <https://www.ireb.org/content/downloads/22-cpre-advanced-level-re-agile-handbook/handbook_cpre_al_re%40agile_en_v1.0.2.pdf>`__"
- "`Agile Practice Guide <https://www.pmi.org/pmbok-guide-standards/practice-guides/agile>`__" by Project Management Institute, 2017
- "Agile Extension to the BABOK® Guide" version 2 (actual, на момент написания статьи)
- "`Agile Software Requirements: Lean Requirements Practices for Teams, Programs, and the Enterprise <https://www.amazon.com/Agile-Software-Requirements-Enterprise-Development/dp/0321635841>`__" by Dean Leffingwell.
- "Software Requirements (Developer Best Practices)" 3rd Edition by Karl Wiegers

См. также:

- "`Agile Modeling :: Requirements-Analysis Models <http://agilemodeling.com/artifacts/#Requirements>`__"
- "`SAFe Requirements Model <https://www.scaledagileframework.com/safe-requirements-model/>`__"

- "`Library of IREB artifacts <https://www.ireb.org/en/downloads/tag:handbook>`__"

TODO:

- https://t.me/emacsway_log/531
- https://t.me/emacsway_log/157
- https://t.me/emacsway_log/158

.. seealso::

   - ":ref:`emacsway-adaptation`"
