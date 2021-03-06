:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: Technical Debt; as compound interest
   :name: emacsway-compound-interest

================================
Technical Debt и сложный процент
================================

.. sectionauthor:: Ivan Zakrevsky

.. contents:: Содержание

":ref:`Ложная Коммутативность <emacsway-planning-error-false-commutativity>`" в управленческих решениях - это когда совокупная стоимость реализации нескольких задач зависит от последовательности их выполнения.
Особенно это касается платформенных (технических) задач, которые иногда нужно уметь "продать" бизнесу, если только у вас нет выделенной `платформенной команды <https://www.scaledagileframework.com/agile-teams/>`__.

И, зачастую, представителям бизнеса, находящимся в контексте бизнес-проблем, а не технических проблем, сложно понять, почему первой должна выполняться задача "А", если задача "Б" более востребована бизнесом.
Это называется :ref:`поиском баланса между бизнес-интересами и техническими интересами <emacsway-agile-balancing-business-technical-concerns>`.
Острота этого вопроса в Agile даже привела к тому, что Quality, изначально бывшее 4-ой переменной управления разработкой (наряду с Cost, Time, Scope), очень быстро было :ref:`преобразовано  в константу <emacsway-xp2-balancing-business-technical-concerns>`.

Выявлением причинно-следственных циклов занимается "Системное Мышление", и вам сильно повезло, если ваши представители бизнеса знакомы с этой наукой.
В противном случае, как пишет Craig Larman, проблемы могут возникать даже в таких компаниях, как Microsoft, являющихся "колыбелью архитектуры" (откуда вышли такие авторы, как Steve McConnell), см. "`Systems Thinking <https://less.works/less/principles/systems-thinking.html>`__" by Craig Larman (`на русском <https://less.works/ru/less/principles/systems-thinking.html>`__).

Наверное, наилучшим образом помочь донести представителям бизнеса проблему "Ложной Коммутативности" может "`Сложный Процент <https://quote.rbc.ru/news/training/5e280d059a7947eb63f54970>`__".

Основная проблема термина "Technical Debt" в русском менталитете заключается в том, что зачастую его влияние на стоимость разработки воспринимается представителями бизнеса как константная величина - сколько сегодня позаимствовали, столько завтра и вернем.

На самом же деле, основная суть термина "`Technical Debt <https://martinfowler.com/bliki/TechnicalDebt.html>`__", вложенная в этот термин by Ward Cunningham, заключается в "interest on the debt":

    📝 "Thinking of this as paying interest versus paying of principal can help decide which cruft to tackle."

    -- "`Technical Debt <https://martinfowler.com/bliki/TechnicalDebt.html>`__" by Martin Fowler

..

    📝 "...that would slow us down which is like paying interest on a loan."

    -- "`Debt Metaphor <https://youtu.be/pqeJFYwnkjE?t=90>`__" by Ward Cunningham

А он рассчитывается по правилу "сложного процента".
Вся суть в том, чтобы перенести внимание бизнеса с тела долга на геометрическую прогрессию роста его процента:

    ::

        S = P*(1+ i)^n (в степени n),

    где S – получаемая сумма, P – цена облигации при покупке, i – процент в сотых долях, а n – число выплат.

    -- "`Как рассчитать доходность облигаций <https://bankiros.ru/wiki/term/kak-rasscitat-dohodnost-obligacij>`__" / Редакция Bankiros.ru

По утверждению ряда авторитетных авторов (Kent Beck, Martin Fowler, Robert C. Martin и др.), рост стоимости разработки при накоплении техдолга может обретать характер, близкий именно к геометрической прогрессии (или к экспоненте).
Статистику конкретного проекта смотрите в "Clean Architecture: A Craftsman's Guide to Software Structure and Design" by Robert C. Martin.

Термин "TechnicalDebt" у нас воспринимается несколько по-иному, безобиднее, еще и потому, что он ассоциируется в нашем менталитете с "академической задолженностью" в ВУЗе, где в центре внимания лежит тело долга (paying of principal), а не геометрический характер роста процентов (paying interest).
А так же находит отражение традиция безпроцентного долга в виде взаимовыручки "до зарплаты", который является распространенным явлением в дружеской среде.
Поэтому, термин "TechnicalDebt" не несет своего изначального смысла в нашем менталитете, и не выполняет своей функции, как метафоры, надлежащим образом, что приводит к росту напряженности и недопонимания между представителями бизнеса и техническими специалистами.
Для разработчика техдолг означает экспоненциальный рост стоимости изменения кода, в то время, как для представителя бизнеса, - это просто отложенное выполнение задачи.

    📝 "Asking "is something technical debt" is usually uninteresting. The metaphor's value is comparing paying interest vs principal. We judge debt of $1K differently if we are paying $1/month to service it vs $100/month"

    -- `Martin Fowler <https://twitter.com/martinfowler/status/1517152168775614473?t=BMpST8vXSLBnY36k9o-lUg&s=19>`__

.. seealso::

   - ":ref:`emacsway-architecture-options`"
