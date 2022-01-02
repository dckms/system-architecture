:canonical-base-url: https://dckms.github.io/system-architecture

.. index:: Prediction, Adaptation
   :name: emacsway-balancing-prediction-adaptation

===============================
Balancing Prediction/Adaptation
===============================

    📝 McConnell writes, "In ten years the pendulum has swung from 'design everything' to 'design nothing.'
    But the alternative to BDUF [Big Design Up Front] isn't no design up front, it's a Little Design Up Front (LDUF) or Enough Design Up Front (ENUF)."
    This is a strawman argument.
    The alternative to designing before implementing is designing after implementing.
    Some design up-front is necessary, but just enough to get the initial implementation.
    Further design takes place once the implementation is in place and the real constraints on the design are obvious.
    Far from "design nothing," the XP strategy is "design always."

    -- "Extreme Programming Explained" 2nd edition by Kent Beck

..

    📝 "The incremental and iterative nature of Agile development can facilitate efficient technical and management processes and practices to reduce the cost associated with change.
    In comparison, projects managed at the waterfall end of the continuum seek to reduce total rework cost by minimizing the number of changes, limiting the number of control points, and baselining detailed specifications which are reviewed and traced throughout the project."

    -- "ISO/IEC/IEEE 12207:2017 Systems and software engineering - Software life cycle processes"


Alberto Brandolini about Prediction/Adaptation
==============================================

.. sectionauthor:: Андрей Ганичев

Андрей Ганичев, contributor of "`Full Modular Monolith application with Domain-Driven Design approach <https://github.com/kgrzybek/modular-monolith-with-ddd>`__", на тему поиска баланса Prediction/Adaptation:

Когда читал книгу Брандолини про "`Introducing EventStorming: An act of Deliberate Collective Learning <https://leanpub.com/introducing_eventstorming>`__" by Alberto Brandolini (та которая недописанная), обратил внимание что и он вскользь проходит по этой теме.

Глава Pretending to solve the problem writing software, раздел Embrace Change:

    📝 "...iterative development is expensive. It is the best approach for developing software in very complex, and lean-demanding domains. However, the initial starting point matters, a lot. A big refactoring will cost a lot more than iterative fine tuning (think splitting a database, vs renaming a variable). So I’ll do everything possible to start iterating from the most reasonable starting point."

    -- "`Introducing EventStorming: An act of Deliberate Collective Learning <https://leanpub.com/introducing_eventstorming>`__" by Alberto Brandolini

..

    📝 "Upfront is a terrible word in the agile jargon. It recalls memories the old times analysis phase in the worst corporate waterfall. Given this infamous legacy, the word has been banned from agile environments like blasphemy. But unfortunately …there’s always something upfront. Even the worst developer thinks before typing the firs line of code."

    -- "`Introducing EventStorming: An act of Deliberate Collective Learning <https://leanpub.com/introducing_eventstorming>`__" by Alberto Brandolini

Источник: https://t.me/emacsway_log/819

TODO:

- https://t.me/emacsway_log/812
- https://t.me/emacsway_log/731
