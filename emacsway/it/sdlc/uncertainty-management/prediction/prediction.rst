:canonical-base-url: https://dckms.github.io/system-architecture

.. index:: Prediction
   :name: emacsway-prediction

====================
Что такое Prediction
====================

.. sectionauthor:: Ivan Zakrevsky

..

    📝 A second common style of definition for architecture is that it's "the design decisions that need to be made early in a project", but Ralph complained about this too, saying that it was more like the decisions you wish you could get right early in a project.

    -- "`Software Architecture Guide <https://martinfowler.com/architecture/>`__" by Martin Fowler

К Prediction (Прогнозированию) относится ряд активностей на основе правил логического вывода, предшествующих производству Системного Инкремента и направленных на заблаговременное разрешение неопределенности.
К ним относятся Business/System Requirements Definition and Analysis, Architecture Definition, Design Definition, разработка макетов UX/UI-Design, Estimation и Planning.
Но главным образом к ним относится разрешение неопределенности в problem-space (т.е. требований), что оказывает существенное влияние на выбор SDLC-модели.

В Scrum эти активности традиционно выражаются в событиях, предшествующих Definition Of Ready (DOR):

- Со стороны Product Owner:
    - Reveal needs of stakeholders
    - Creating PBI
- Со стороны Team:
    - PBR
    - Planning
    - Spike

Основная проблема Prediction заключается в том, что характер роста стоимости Прогнозирования, в зависимости от его точности, близок к экспоненциальному.
В то время как характер роста бизнес-выгод от точности Прогнозирования близок к логарифмическому.
Пересечение этих графиков образует предел экономической целесообразности разрешения неопределенности путем Прогнозирования.
Здесь подразумевается, что стоимость самой реализации уже вычтена из бизнес-выгод.

    💬 "WaterFall is based on the empirical observation of 30 years ago (ref: BarryBoehm, Software Engineering Economics, Prentice Hall, 1981.) that the cost of change rises exponentially (base 10) by phases. The conclusion is that you should make the big decisions up front, because changing them is so expensive."

    -- "`Water Fall <https://wiki.c2.com/?WaterFall>`__" at c2.com

..

    💬 "There is a fundamental truth to work breakdown structure estimation - the only way to estimate using a work breakdown structure to really accurately to get the number right is to implement the project.
    Then you'll have the estimate.

    The cost of improving the estimate, the initial work breakdown estimation, the cost of refining that grows exponentially.
    With every next layer you want to take it down.
    And the benefit of doing that does not grow exponentially.
    It grows logarithmically.

    Managers understand there is an extremely high cost involved with refining the estimate."

    -- `"YOW! 2016 Robert C. Martin - Effective Estimation (or: How not to Lie)" at 33:15 <https://youtu.be/eisuQefYw_o?t=1995>`__

Robert C. Martin упомянул о таком способе обработки неопределенности, как реализовать Проект (или Системный Инкремент).
Это уже другой, эмпирический (т.е. опытным путем) способ обработки неопределенности, который называется :ref:`Adaptation <emacsway-adaptation>` и составляет основу итеративной разработки.
Чем сложнее предметная область Проекта, тем раньше наступает предел экономической целесообразности Prediction.
