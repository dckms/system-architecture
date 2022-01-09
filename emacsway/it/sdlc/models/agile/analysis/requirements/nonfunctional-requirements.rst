:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: Requirements; nonfunctional in Agile
   :name: emacsway-agile-nonfunctional-requirements


================================
Agile nonfunctional Requirements
================================

.. sectionauthor:: Ivan Zakrevsky

Если вы знакомы с ATAM, ADD, QAW, то вы знакомы со сценариями атрибутов качеств ("Software Architecture in Practice" 3d edition by Len Bass, Paul Clements, Rick Kazman, глава 16), и легко поймете, о чем пишет K.Wigers - один из топовых авторов по требованиям и аналитике:

    📝 "**Nonfunctional requirements need to have priority alongside user stories; you can't defer their implementation until a later iteration.**
    It's possible to specify quality attributes in the form of stories: "As a help desk technician, I want the knowledge base to respond to queries **within five seconds** so the customer doesn't get frustrated and hang up."

    <...>

    Developers need to keep nonfunctional requirements in mind as they consider the implications of implementing individual user stories.
    As more functionality is added through a series of iterations, the system's efficiency and hence performance can deteriorate.
    Specify performance goals and begin performance testing **with early iterations**, so you can become aware of concerns early enough to take corrective actions.

    <...>

    As with user stories, it's possible to write acceptance tests for quality attributes.
    This is a way to quantify the quality attributes.
    If a performance goal is stated simply as "The knowledge base must return search results quickly," you can't write tests to define what constitutes "quickly." A better acceptance test would be: "Keyword search of the knowledge base **takes less than 5 seconds**, and preferably **less than 3 seconds**, to return a result.

    <...>

    Part of **accepting an iteration as being complete** is to assess whether the pertinent nonfunctional requirements are satisfied.
    Often there is a range of acceptable performance, with some outcomes more desirable than others.
    As it does for any other software development approach, satisfying quality requirements can distinguish delight from disappointment on agile projects."

    -- Глава 14, подраздел "Handling quality attributes on agile projects", книги "Software Requirements" 3rd Edition, Wiegers K., Beatty J.

Приоритезация и реализация нефункциональных требований совместно (alongside) с функциональными выбивает почву для появления технических задач, оторванных от бизнес-целей.
А поскольку один из топовых авторов по Scrum, K.Rubin, утверждает, что нефункциональными требованиями должен заниматься Product Owner, то это автоматически устраняет необходимость в лоббировании технических задач техническими специалистами.

Нужно заметить, что в SAFe нефункциональными требованиями занимается уже не Product Owner, а System Architect.

    System Architect/Engineering is an individual or team that defines the overall architecture of the system. They work at a level of abstraction above the teams and components and define Non-Functional Requirements (NFRs), major system elements, subsystems, and interfaces.

    -- "SAFe® 5.0: The World's Leading Framework for Business Agility" by Richard Knaster, Dean Leffingwell

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
    Each nonfunctional requirement is a prime target for inclusion in the team's definition of done.
    If the team includes the "Web Browser Support" nonfunctional requirement in the definition of done, the team will have to test any new features added in the sprint with all of the listed browsers.
    If it doesn't work with all of them, the story isn't done.

    I recommend that teams try to include as many of the nonfunctional requirements in their definitions of done as they possibly can.
    Waiting to test nonfunctional requirements until late in the development effort defers getting fast feedback on critical system performance characteristics.

    <...>

    nonfunctional requirement.

    1. A requirement that does not relate to functionality but to attributes such as reliability, efficiency, usability, maintainability, and portability, which product backlog items must possess in order to be fully accepted by the stakeholders.
    2. **Each nonfunctional requirement is a candidate for inclusion in the definition of done**. See also definition of done."

    -- "Essential Scrum : a practical guide to the most popular agile process" by Kenneth S. Rubin.

Топовые авторы по аналитике (K.Wigers) и Scrum (K.Rubin) сошлись в одном мнении. Один дает описание с точки зрения работы с требованиями, а другой - с точки зрения гибких процессов разработки.


    📝 "So, in the information model, we have modeled NFRs as backlog constraints, as illustrated in Figure 17–1.

    From the diagram, we see that a backlog item may be constrained by (zero, one, or more) nonfunctional requirements.
    An example from the case study appears in Figure 17–2.

    Also, nonfunctional requirements apply to zero or more backlog items.
    For example, a nonfunctional requirement such as support 100 concurrent users might apply to zero, one, or many backlog items.

    Once identified, relevant nonfunctional requirements must be captured and communicated to all teams who may be affected by the constraints.
    In agile, with its focus on the backlog, there is no obvious place to model them, so in the Big Picture, we just call them backlog constraints and represent them as shown in Figure 17–3."

    -- "Agile Software Requirements: Lean Requirements Practices for Teams, Programs, and the Enterprise" by Dean Leffingwell

..

    Persisting nonfunctional requirements

    Another difference between user stories and nonfunctional requirements is that they typically need to persist differently in the development life cycle. We've described how user stories are lightweight and generally don't have to be maintained, which is one of the key benefits. We've also shown that the details of a user story are captured in the acceptance test, which persist inside the team's automated or manual regression test environment. That is why we can throw the user story away after implementation—because we have memorialized the important details in our test cases.

    That can work for some NFRs, too, but it gets a bit riskier. For example, if a system must support 1000 concurrent users, we could develop an automated test that simulated that load and build it in the regression test suite. That would be an excellent practice because we could refactor the code at will, and if we accidentally created a performance bottleneck, it would be quickly discovered. In that case, we could forget about the NFR once we have seen it the fi rst time, because the automated test remembers it for us.

    There are other types of NFRs, however, that must be treated quite differently. Here are some examples:

    - Maintain PCI compliance (credit card industry user security standards) in all applications
    - Localize the application in all then-current, supported languages prior to release in any language 
    - No open source without a CFO license review

    We surely can't forget these, and we can't write automated test cases for them, either. 
    So, the teams must have an organized way to save them, find them, and review them when necessary. In practice, we've seen agile teams take a number of approaches to persisting NFRs.

    - Create a separate backlog in the agile project management tool. Most enterprises will adopt agile project management tooling as a central repository for stories and tasks, as well as iteration and release objects that support scheduling, burndown, and feature status reporting. Teams can create a special project/product backlog to hold and maintain the NFRs within the tool. Access privileges must be granted to all team members who are working on the program.

    - Store and manage them in a wiki. This method works well because it provides continuous visibility; is available to all team members; is persistent; fosters communication, comments, and interaction; and doesn't require any special tooling.

    - Maintain a supplementary specification. This label/document was originally developed as an auxiliary document to RUP's use case models and use case specifications and served exactly this role (organizing nonfunctional requirements). Remember, as agilists, we "favor working software over comprehensive documentation," but that doesn't mean we can't create the documentation we need. Even more importantly, we like to do the simplest thing that can possibly work, and when we know something is important, it makes sense to write it down. Table 17–3 later in this chapter provides an example template for a supplemental specif i cation.

    - Build the NFRs into the definition of done, and point to the special backlog, wiki, or supplemental specification that contains the details. In this approach, a team can't be done until the NFRs are satisfied as well. Different definitions of done, requiring different amounts of regression testing, inspection, and so on, can be established for various iteration, potentially shippable increment, and release milestones.

    No matter the approach, it is mandatory that the teams do something to maintain and manage these specif i cations, because they could make the difference between success and failure.

    -- "Agile Software Requirements: Lean Requirements Practices for Teams, Programs, and the Enterprise" by Dean Leffingwell


См. также:

- "`Nonfunctional Requirements <https://www.scaledagileframework.com/nonfunctional-requirements/>`__" at SAFe
- "`An Agile Architectural Epic Kanban System: Part 2 – The Model <https://scalingsoftwareagility.wordpress.com/2010/03/05/an-agile-architectural-epic-kanban-system-part-2-%E2%80%93-the-model/>`__" by Dean Leffingwell
- `A Lean and Scalable Requirements Information Model for the Agile Enterprise <https://scalingsoftwareagility.files.wordpress.com/2007/03/a-lean-and-scalable-requirements-information-model-for-agile-enterprises-pdf.pdf>`__ by Dean Leffingwell with Juha‐Markus Aalto
