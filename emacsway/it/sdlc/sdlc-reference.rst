:canonical-base-url: https://dckms.github.io/system-architecture

.. index:: SDLC, Literature

===============================================
Systems Development Life Cycle (SDLC) Reference
===============================================

.. sectionauthor:: Ivan Zakrevsky

Grady Booch выделяет два ключевых критерия, которые определяют успех или неудачу проекта:

    📝 "Traits of Successful Projects 

    A successful software project is one in which the deliverables satisfy and possibly exceed the customer's expectations, the development occurred in a timely and economical fashion, and the result is resilient to change and adaptation. By this measure, we have observed several traits that are common to virtually all of the successful object-oriented systems we have encountered and noticeably absent from the ones that we count as failures: 

    - Existence of a strong architectural vision
    - Application of a well-managed iterative and incremental development lifecycle"

    -- "Object-Oriented Analysis and Design with Applications" 3rd edition by Grady Booch, Robert A. Maksimchuk, Michael W. Engle, Bobbi J. Young Ph.D., Jim Conallen, Kelli A. Houston

Если с "architectural vision" все понятно, то с "well-managed development lifecycle" у многих могут возникать вопросы. Тем не менее, для успешности проекта процессы должны быть грамотно выбраны и качественно отлажены.

- "`Учебник 4CIO <https://book4cio.ru/#page-14>`__" - глава 3.4. Управление разработкой ПО
- "`SEBoK: Life Cycle Models <https://www.sebokwiki.org/wiki/Life_Cycle_Models>`__"
- "`SEBoK: System Life Cycle Process Models: Iterative <https://www.sebokwiki.org/wiki/System_Life_Cycle_Process_Models:_Iterative>`__"
- Неплохой "`SDLC Tutorial <https://www.tutorialspoint.com/sdlc/index.htm>`__"
- "`SDLC - Quick Guide <https://www.tutorialspoint.com/sdlc/sdlc_quick_guide.htm>`__"
- "`ITABoK: Architecture Methodologies and Frameworks <https://itabok.iasaglobal.org/itabok3_0/architecture-methodologies-and-frameworks/>`__"
- "`ITABoK: What is Agility <https://itabok.iasaglobal.org/itabok3_0/digital-outcome-model/agility/>`__"
- "`Systems engineering handbook. A guide for System Life Cycle Processes and activities <https://www.incose.org/products-and-publications/se-handbook>`__" by INCOSE

- "`ISO/IEC/IEEE 12207:2017 Systems and software engineering — Software life cycle processes <https://www.iso.org/standard/63712.html>`__"
- "`ISO/IEC/IEEE 15288:2015 Systems and software engineering — System life cycle processes <https://www.iso.org/standard/63711.html>`__"

- "ISO/IEC/IEEE 29148:2011 Systems and software engineering — Life cycle processes — Requirements engineering"
- "ISO/IEC/IEEE 15289:2019 Systems and software engineering — Content of life-cycle information items (documentation)"

- "ISO/IEC/IEEE 24765:2017 Systems and software engineering — Vocabulary"
- "ISO 9000:2005 Quality management systems — Fundamentals and vocabulary"

- "ISO/IEC 33001:2015 Information technology — Process assessment — Concepts and terminology"

- "ГОСТ Р ИСО/МЭК 12207-2010 Информационная технология. Системная и программная инженерия. Процессы жизненного цикла программных средств."
- "ГОСТ Р 57193-2016 Системная и программная инженерия. Процессы жизненного цикла систем."


- "Object-Oriented Analysis and Design with Applications" 3rd edition by Grady Booch, Robert A. Maksimchuk, Michael W. Engle, Bobbi J. Young Ph.D., Jim Conallen, Kelli A. Houston - "Chapter 6. Process"
- "Software Architecture in Practice" 3d edition by Len Bass, Paul Clements, Rick Kazman - "Chapter 15. Architecture in agile projects"
- "Extreme Programming Explained" 1st edition by Kent Beck (именно первое издание) - кто этой книги не читал, тот ничего в Agile не понимает.

Три превосходные книги Dean Leffingwell:

- "Scaling Software Agility: Best Practices for Large Enterprises" by Dean Leffingwell - о проблемах масштабирования команд.
- "Agile Software Requirements: Lean Requirements Practices for Teams, Programs, and the Enterprise" by Dean Leffingwell - о проблемах интегрирования аналитической и архитектурной работы в Agile.
- "SAFe® 5.0: The World's Leading Framework for Business Agility" by Richard Knaster, Dean Leffingwell - наиболее удачная масштабируемая Agile-модель на сегодня.

..

- "Essential Scrum: A Practical Guide to the Most Popular Agile Process" by Kenneth Rubin - главы 3 и 8 просто необходимы для понимания области применения Scrum.

- https://t.me/emacsway_log/431
- https://t.me/emacsway_log/35
- https://t.me/emacsway_log/151

- "`Краткая история итеративной разработки <https://www.craiglarman.com/wiki/downloads/misc/history-of-iterative-larman-and-basili-ieee-computer.pdf>`__" by Craig Larman
- "`LeSS by Craig Larman <https://less.works/less/framework/introduction>`__" - куча полезной информации, которую можно использовать автономно. (`на русском <https://less.works/ru/less/framework/introduction>`__)

Нужно заметить, что Software является подмножеством System.
Раньше и Software, и System описывались одним стандартом ISO/IEC 12207:2008.
Потом их разделили на ISO/IEC/IEEE 12207:2017 и ISO/IEC/IEEE 15288:2015.

    📝 "This document has a strong relationship with ISO/IEC/IEEE 15288:2015, Systems and Software Engineering System Life Cycle Processes, and is more applicable to software systems.
    To account for situations in which both ISO/IEC/IEEE 15288:2015 and ISO/IEC/IEEE 12207:2017 are applied (e.g., a development of a system containing software, or the development of a software system containing hardware), their process structures are harmonized to be identical.
    The processes of this document directly correspond to processes of ISO/IEC/IEEE 15288 with specialization for software products and services."

    -- "ISO/IEC/IEEE 12207:2017 Systems and software engineering - Software life cycle processes"
