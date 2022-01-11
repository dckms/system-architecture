:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: Requirements; in balancing business and technical concerns in Agile
   :name: emacsway-agile-balancing-business-technical-concerns

===============================================
Достижение баланса Бизнес/Технических интересов
===============================================

.. sectionauthor:: Ivan Zakrevsky

.. contents:: Содержание

Существует довольно распространенная проблема поиска баланса между краткосрочными бизнес-интересами и долгосрочными техническими интересами проекта.


Поиск баланса
=============

В первой версии XP управление разработкой осуществлялось посредством четырех переменных:

- затраты (cost)
- время (time)
- качество (quality)
- объем работ (scope)

Iron Triangle изначально был квадратом, и только во второй версии XP он превратился в треугольник.
Но обо всем по порядку.

    📝 "Одним из основополагающих правил нашей стратегии является то, что технари концентрируются на решении технических проблем, а бизнесмены — на решении бизнес-проблем.
    Проект должен управляться бизнес-решениями, однако для принятия бизнес-решений должна использоваться информация о затратах и риске, предоставляемая техническими специалистами.

    Эта информация является результатом технических решений.
    **Существуют два широко распространенных неправильных режима взаимоотношений между бизнесом и разработчиками: когда либо бизнес, либо разработчики получают слишком большую власть над проектом, проект начинает страдать.**

    One key to our strategy is to keep technical people focused on technical problems and business people focused on business problems.
    The project must be driven by business decisions, but the business decisions must be informed by technical decisions about cost and risk.

    There are two common failure modes in the relationship between Business and Development.
    **If either Business or Development gains too much power, the project suffers.**"

    -- "Extreme Programming Explained" 1st edition by Kent Beck, "Chapter 14. Splitting Business and Technical Responsibility", перевод ООО Издательство "Питер"


.. index::
   single: Requirements; in predominance of technical concerns in Agile
   :name: emacsway-agile-technical-concerns-predominance

Преобладание технических интересов
----------------------------------

Начнем с преобладания технических интересов.

.. index::
   single: The Second-System Effect; in Agile
   :name: emacsway-second-system-effect

Frederick Brooks в своем бестселлере "Мифический человеко-месяц" говорит об "`Эффекте второй системы <https://ru.m.wikipedia.org/wiki/%D0%AD%D1%84%D1%84%D0%B5%D0%BA%D1%82_%D0%B2%D1%82%D0%BE%D1%80%D0%BE%D0%B9_%D1%81%D0%B8%D1%81%D1%82%D0%B5%D0%BC%D1%8B>`__".
Я приведу небольшой фрагмент:

    📝 "Если ответственность за спецификацию функций отделить от ответственности за быстрое создание недорогого продукта, то **чем сдержать изобретательский энтузиазм архитектора**?

    If one separates responsibility for functional specification from responsibility for building a fast, cheap product, **what discipline bounds the architect's inventive enthusiasm**?"

    -- "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., перевод ООО Издательство "Питер"

Похожую проблему описывает и Kent Beck в "Extreme Programming" 1st edition:

    📝 "Когда разработчикам предоставляется чрезмерная свобода, они начинают использовать все те новые технологии и процессы, для которых у них никогда не хватает времени, если "эти белые воротнички" постоянно подгоняют их.
    Когда разработчикам предоставляется свобода, они устанавливают и начинают использовать новые инструменты разработки, новые языки программирования, новые технологии.
    При этом инструменты, языки и технологии выбираются исходя из того, что они очень интересны и суперсовременны.
    Все только что появившееся на рынке связано с риском.
    (Если мы не попробуем это сейчас, то когда же еще?)

    Таким образом, в результате предоставления разработчикам слишком широких полномочий, они прикладывают слишком много усилий и генерируют слишком много риска, при этом **они обеспечивают слишком незначительную отдачу**.

    When Development is in charge, they put in place all the process and technology that they never had time for when "those suits" were pushing them around.
    They install new tools, new languages, new technologies.
    And the tools, languages, and technologies are chosen because they are interesting and cutting edge.
    Cutting edge implies risk.
    (If we haven't learned that by now, when will we?)

    So, the net result of the "Development in Charge" scenario is too much effort and way, way **too much risk for too little return**."

    -- "Extreme Programming Explained" 1st edition by Kent Beck, перевод ООО Издательство "Питер"

..

    📝 "Есть масса примеров, подсказанных другими искусствами и ремеслами, которые
    подводят к мнению, что дисциплина идет на пользу.
    Действительно, афоризм художника гласит, что "форма освобождает".
    Самые ужасны строения — это те, бюджет которых был слишком велик для поставленных целей.
    Творческую активность Баха едва ли могла подавлять еженедельная необходимость изготавливать кантату определенного вида.
    Я уверен, что архитектура компьютера Stretch стала бы лучше, если бы на нее наложили более жесткие ограничения; так, ограничения, наложенные бюджетом на System/360 Model 30, по моему мнению, принесли лишь пользу архитектуре Model 75.

    Аналогично, я считаю, что получение архитектуры извне усиливает, а не подавляет творческую активность группы исполнителей.
    Они сразу сосредоточиваются на той части задачи, которой никто не занимался, и в результате изобретательность бьет ключом.
    В не ограничиваемой группе большая часть обдумывания и обсуждения посвящена архитектурным решениям в ущерб реализации. [5]

    There are many examples from other arts and crafts that lead one to believe that discipline is good for art.
    Indeed, an artist's aphorism asserts, "Form is liberating."
    The worst buildings are those whose budget was too great for the purposes to be served.
    Bach's creative output hardly seems to have been squelched by the necessity of producing a limited-form cantata each week.
    I am sure that the Stretch computer would have had a better architecture had it been more tightly constrained; the constraints imposed by the System/360 Model 30's budget were in my opinion entirely beneficial for the Model 75's architecture.

    Similarly, I observe that the external provision of an architecture enhances, not cramps, the creative style of an implementing group.
    They focus at once on the part of the problem no one has addressed, and inventions begin to flow.
    In an unconstrained implementing group, most thought and debate goes into architectural decisions, and implementation proper gets short shrift. [5]"

    5. Englebart, D., and W. English, "A research center for augmenting human intellect," AFIPS Conference Proceedings, Fall Joint Computer Conference, San Francisco (Dec. 9-11, 1968), pp. 395-410.

    -- "The Mythical Man-Month Essays on Software Engineering Anniversary Edition" by Frederick P. Brooks, Jr., перевод ООО Издательство "Питер"

К этой же категории относится и т.н. Resume-Driven Development, когда разработчики безобоснованно переусложняют проект сложными технологиями ради достижения в резюме.


.. index::
   single: Requirements; in predominance of business concerns in Agile
   :name: emacsway-agile-business-concerns-predominance

Преобладание бизнес-интересов
-----------------------------

Но есть и обратная проблема - когда технические специалисты ущемлены в своих полномочиях в пользу представителей бизнеса, проект неизменно загнивает, а :ref:`экономика разработки деградирует с зависимостью, приближенной к экспоненциальной <emacsway-agile-development-difficulties>`.

Решение этой проблемы хорошо описано в главе "Chapter 14. Splitting Business and Technical Responsibility" книги "Extreme Programming Explained" 1st edition by Kent Beck.
Решение слишком объемное, чтобы его сюда поместить.

    📝 "Когда бизнесмены получают слишком много полномочий, они начинают диктовать разработчиком значения для всех четырех переменных.
    "Вот то, что ты должен сделать.
    Это должно быть сделано тогда-то и тогда-то.
    Нет, тебе не дадут ни одной дополнительной рабочей станции.
    И для тебя будет лучше, если ты сделаешь эту работу с наивысшим возможным качеством, иначе у тебя будут проблемы.
    Ты меня хорошо понял? Скотина ленивая!"

    В такой ситуации бизнес предписывает слишком многое.
    Некоторые элементы в списке требований абсолютно обязательны, но некоторые — нет.
    И если у разработчиков не будет никаких полномочий, они не смогут возразить.
    Они не смогут принудить бизнес выбрать правильный вариант.
    И тогда разработчики, понурив голову, идут работать над невыполнимой задачей, которую перед ними поставили.

    Как правило, наименее важные требования являются причиной наибольшего риска.
    Похоже, это является следствием их природы.
    Они меньше всего обдумываются, меньше всего анализируются и меньше всего осмысливаются, поэтому вероятность того, что именно они изменятся в процессе разработки, выше всего.
    Очень часто такие требования оказываются также наиболее рискованными с технической точки зрения.

    **В результате, если бизнес получает слишком большие полномочия, проект требует слишком много усилий и генерирует слишком много риска, при этом он** :ref:`обеспечивает слишком незначительную отдачу <emacsway-agile-development-difficulties>`.

    If Business has the power, they feel fit to dictate all four variables to Development.
    "Here is what you will do.
    Here is when it will be done.
    No, you can't have any new workstations.
    And it better be of the highest quality or you're in trouble, buster."

    In this scenario, Business always specifies too much.
    Some of the items on the list of requirements are absolutely essential.
    But some are not.
    And if Development doesn't have any power, they can't object; they can't force Business to choose which is which.
    So Development dutifully goes to work, heads down, on the impossible task they have been given.

    It seems to be in the nature of the less important requirements that they entail the greatest risk.
    They are typically the poorest understood, so there is great risk that the requirements will change all during development.
    Somehow, they also tend to be technically riskier.

    **The result of the "Business in Charge" scenario, then, is that the project takes on too much effort and way, way too much risk for :ref:too little return** <emacsway-agile-development-difficulties>`."

    -- "Extreme Programming Explained" 1st edition by Kent Beck, "Chapter 14. Splitting Business and Technical Responsibility", перевод ООО Издательство "Питер"

..

    📝 "Закон Вайнберга-Брукса: От действий менеджеров, основанных на неправильных моделях системы, пострадало больше проектов, чем от всех остальных причин вместе взятых.

    Weinberg-Brooks' Law: More software projects have gone awry from management's taking action based on incorrect system models than for all other causes combined."

    -- "`Systems Thinking <https://less.works/ru/less/principles/systems-thinking.html>`__ by Craig Larman

..

    📝 "Scrum is by far the most widely used agile framework in the world, but we've also found that 58% of Scrum implementations fail."

    -- "`Better Scrum with Essence <https://www.scruminc.com/better-scrum-with-essence/>`__" Jeff Sutherland



Решение
-------


.. index::
   single: Requirements; in balancing business and technical concerns in XP
   :name: emacsway-xp-balancing-business-technical-concerns

Extreme Programming
^^^^^^^^^^^^^^^^^^^

Первая версия XP
""""""""""""""""

    📝 "Что делать?

    Решение состоит в том, чтобы определенным образом **разделить полномочия и ответственность между бизнесом и разработчиками**.
    **Бизнесмены должны принимать решения в своей области компетенции, а программисты должны принимать решения в своей области компетенции.**
    Решения, принятые одной стороной, должны стать базой для решений, принимаемых другой стороной.
    Ни одна сторона не должна в одностороннем порядке решать абсолютно все.

    What to Do?

    The solution is to somehow **split the responsibility and power between Business and Development**.
    **Business people should make the decisions for which they are suited.**
    **Programmers should make the decisions for which they are suited.**
    Each party's decisions should inform the other's.
    Neither party should be able to unilaterally decide anything."

    -- "Extreme Programming Explained" 1st edition by Kent Beck, "Chapter 14. Splitting Business and Technical Responsibility", перевод ООО Издательство "Питер"

..

    📝 "В данной главе я расскажу вам о модели разработки программного обеспечения, которая представляет собой систему контролируемых переменных.
    В рамках данной модели разработка программного обеспечения определяется с использованием следующих четырех переменных:

    - затраты (cost);
    - время (time);
    - качество (quality);
    - объем работ (scope).

    В данном случае игра в разработку программного обеспечения выглядит следующим образом: **внешние силы (заказчики, менеджеры) должны определить значения для любых трех переменных из указанного набора, при этом команда разработчиков должна выбрать результирующее значение для оставшейся переменной**.

    Некоторые менеджеры и заказчики полагают, что они обладают правом с успехом установить значение для всех четырех переменных.
    "Вы обязаны реализовать все, что указано в техническом задании к первому числу следующего месяца, работая в текущем составе, то есть без увеличения численности, при этом качество должно стоять на первом месте и не уступать нашим обычным стандартам".
    Когда происходит подобное, :ref:`качество, как правило, летит ко всем чертям <emacsway-agile-development-difficulties>` (и это, к сожалению, как раз и является общераспространенным стандартом), потому что никто не в состоянии хорошо делать свою работу под слишком большим давлением.
    Помимо качества, время, как правило, также выходит из-под контроля.
    Таким образом, вы производите некачественное программное обеспечение, не успевая при этом сдать работу к сроку.

    Чтобы решить проблему, необходимо сделать все четыре переменные наблюдаемыми.
    Если все — программисты, заказчики и менеджеры — смогут наблюдать за поведением всех четырех переменных, будет легче сознательно выбрать, какие из четырех переменных следует контролировать.
    Если результирующее значение четвертой переменной окажется неприемлемым, можно будет либо изменить входные значения, либо выбрать для контроля другие три переменные.

    Here is a model of software development from the perspective of a system of control variables.
    In this model, there are four variables in software development:

    - Cost
    - Time
    - Quality
    - Scope

    The way the software development game is played in this model is that **external forces (customers, managers) get to pick the values of any three of the variables**.
    **The development team gets to pick the resultant value of the fourth variable.**

    Some managers and customers believe they can pick the value of all four variables.
    "You are going to get all these requirements done by the first of next month with exactly this team.
    And quality is job one here, so it will be up to our usual standards."
    When this happens, :ref:`quality always goes out the window <emacsway-agile-development-difficulties>` (this is generally up to the usual standards, though), since nobody does good work under too much stress.
    Also likely to go out of control is time.
    You get crappy software late.

    The solution is to make the four variables visible.
    If everyone—programmers, customers, and managers—can see all four variables, they can consciously choose which variables to control.
    If they don't like the result implied for the fourth variable, they can change the inputs, or they can pick a different three variables to control."

    -- "Extreme Programming Explained" 1st edition by Kent Beck, "Chapter 4. Four Variables", перевод ООО Издательство "Питер"


Вторая версия XP
""""""""""""""""

Сам же Kent Beck и преобразовал позже квадрат (Quality, Cost, Time, Scope) в треугольник (Cost, Time, Scope), путем преобразования качества (Quality) из переменной в константу.

Если в первой версии XP он боролся за то, чтобы качество (или хотя бы любую одну из 4-х переменных контроля разработки) контролировали технические специалисты, то во второй версии он и вовсе преобразовал качество в константу.

Вот что он пишет во втором издании:

    📝 "Quality

    Sacrificing quality is not effective as a means of control.
    **Quality is not a control variable.**
    Projects don't go faster by accepting lower quality.
    They don't go slower by demanding higher quality.
    Pushing quality higher often results in faster delivery; while lowering quality standards often results in later, less predictable delivery.

    One of my biggest surprises since the first edition of Extreme Programming Explained was released has been just how far teams have been able to push quality as measured in defects, design quality, and the experience of development.
    Each increase in quality leads to improvements in other desirable project properties, like productivity and effectiveness, as well.
    There is no apparent limit to the benefits of quality, only limits in our ability to understand how to achieve higher quality.

    Quality isn't a purely economic factor.
    People need to do work they are proud of.
    I remember talking to the manager of a mediocre team.
    He went home on the weekends and made fancy ironwork as a blacksmith.
    He met his need for quality; he just met it outside of work.

    If you can't control projects by controlling quality, how can you control them?
    Time and cost are most often fixed.
    XP chooses scope as the primary means of planning, tracking, and steering projects.
    Since scope is never known precisely in advance, it makes a good lever.
    The weekly and quarterly cycles provide explicit points for tracking and choosing scope.

    A concern for quality is no excuse for inaction.
    If you don't know a clean way to do a job that has to be done, do it the best way you can.
    If you know a clean way but it would take too long, do the job as well as you have time for now.
    Resolve to finish doing it the clean way later.
    This often occurs during architectural evolution, where you have to live with two architectures solving the same problem while you transition from one to the other.
    Then the transition itself becomes a demonstration of quality: making a big change efficiently in small, safe steps."

    -- "Extreme Programming Explained" 2nd edition by Kent Beck


.. index::
   single: Requirements; in balancing business and technical concerns in Scrum
   :name: emacsway-scrum-balancing-business-technical-concerns

Scrum
^^^^^


Scrum Guide™
""""""""""""

А как обстоят дела в Scrum? Официальный Scrum Guide™ не допускает технического долга вообще, как и XP второй версии:

    📝 "During the Sprint: Quality does not decrease;"

    -- "`The 2020 Scrum Guide™ <https://scrumguides.org/scrum-guide.html>`__"

А все отклонения продукта должны устраняться как можно скорее:

    📝 "If any aspects of a process deviate outside acceptable limits or if the resulting **product is unacceptable**, the process being applied or the **materials being produced must be adjusted**.
    The adjustment must be made **as soon as possible to minimize further deviation**".

    -- "`The 2020 Scrum Guide™ <https://scrumguides.org/scrum-guide.html>`__"

Сам Ken Schwaber под "прозрачностью" понимает полное отсутствие техдолга:

    📝 "Transparency means the software is ready.
    It can either be immediately deployed or built upon without regression.
    **It has no technical debt.**"

    -- "`Can Software Developers Meet the Need? <https://kenschwaber.wordpress.com/2014/04/09/can-software-developers-meet-the-needs/>`__ by Ken Schwaber

А баланс бизнес и технических интересов обеспечивается тем, что решения PO инспектируемы:

    📝 "For Product Owners to succeed, the entireorganization must respect their decisions.
    These decisions are visible in the content and ordering of the Product Backlog,
    and through the **inspectable** Increment at the Sprint Review."

    -- "`The 2020 Scrum Guide™ <https://scrumguides.org/scrum-guide.html>`__"

А инспектирует их сбалансированный круг внутренних (команда) и внешних стейкхолдеров:

    📝 "Scrum Definition: The **Scrum Team and its stakeholders inspect** the results and adjust for the next Sprint.
    <...>
    Sprint Review: During the event, the **Scrum Team and stakeholders review** what was accomplished in the Sprint
    and what has changed in their environment."

    -- "`The 2020 Scrum Guide™ <https://scrumguides.org/scrum-guide.html>`__"

Это работает для маленьких команд.
В больших коллективах лучше работают практики для работы со стейкхолдерами типа QAW, Mini-QAW, etc.


Суть от первоисточника
""""""""""""""""""""""

Bertrand Meyer был прав - лучший способ понять суть вещей - это обратиться к первоисточнику.
Jeffrey Sutherland о том, как и зачем он ввел роль Product Owner:

    📝 "When I started the first Scrum team in 1993, I didn't have a Product Owner.
    I was part of the leadership team and had a bunch of other responsibilities besides figuring out exactly what the team should do in each Sprint.
    I carried out management and marketing duties, dealt with customers, and plotted strategy.
    But in that first Sprint I figured I could handle the Backlog.
    I just needed to make sure I had enough "stories" and features for the team to work on during the next Sprint.
    The problem was, after the second Sprint we introduced the Daily Stand-up meeting.
    Velocity went up 400 percent in the next Sprint, and the team finished in a week what we thought would take us a month.
    There was no more Backlog for them to work on! I thought I'd have a month to create more "stories." A great problem to have, admittedly, but one that had to be addressed.
    So I thought about this role of Product Owner and what qualities someone would need to execute it properly.

    My inspiration for the role came from Toyota's Chief Engineer.
    A Chief Engineer at Toyota is responsible for a whole product line, such as the Corolla or the Camry.
    To do this, they have to draw on the talents of groups specializing in body engineering, or chassis, or electrical, or whatever.
    The Chief Engineer has to draw from all those groups to create a cross-functional team capable of creating a car.
    Outside of Toyota everyone thinks of these legendary Chief Engineers (or Shusas, as they were originally called) as all-powerful leaders of the "Toyota Way." And in a way they are.
    But what they don't have is authority.
    No one reports to them—rather, they report to their own groups.
    People can tell Chief Engineers that they're wrong, so they have to make sure they're right.
    They don't give anyone performance appraisals or promotions or raises.
    But they do decide on the vision of the car, and how the car will be made—by persuasion, not coercion.

    It's this idea that I wanted to embody within Scrum.
    John Shook of the Lean Enterprise Institute once began his description of the Chief Engineer role by quoting the US Marine Corps leadership manual:

    "An individual's responsibility for leadership is not dependent on authority.… the deep-rooted assumption that authority should equal responsibility is the root of much organizational evil.
    I believe misunderstanding around this issue is rampant, problematic, and runs so deep in our consciousness that we don't even realize it." [Shook, John. "The Remarkable Chief Engineer." Lean Enterprise Institute, February 3, 2009]

    Reflecting on my time at West Point and in Vietnam, I found myself agreeing that leadership has nothing to do with authority.
    Rather, it has to do with—among other things—knowledge and being a servant-leader.
    The Chief Engineer can't simply say something has to be done a particular way.
    He has to persuade, cajole, and demonstrate that his way is the right way, the best way.
    It usually takes someone with thirty years of experience to fill the role.
    I wanted that in Scrum, but I'm also well aware that very few people have that level of skill and experience.
    So I split the role in two, giving the Scrum Master the how and the Product Owner the what.

    Even in those early days of Scrum I knew that I needed someone who was deeply connected to the customer.
    The Product Owner needed to be able to deliver feedback to the team from the customer each and every Sprint.
    They needed to spend half their time talking to the people buying the product (getting their thoughts on the latest incremental release and how it delivered value) and half their time with the team creating the Backlog (showing them what the customers valued and what they didn't)."

    -- "Scrum: The Art of Doing Twice the Work in Half the Time" by Jeffrey Sutherland

Отдельно следует выделить два критически важных момента, непонимание которых является корнем проблем большинства Scrum-проектов:

    📝 "So I split the role in two, giving the Scrum Master the **how** and the Product Owner the **what**.

    <...>

    The Scrum Master and the team are responsible for **how fast they're going and how much faster they can get**.
    The Product Owner is accountable for **translating the team's productivity into value**."

    -- "Scrum: The Art of Doing Twice the Work in Half the Time" by Jeffrey Sutherland

P.S.: От себя добавлю, что я пробовал различные Agile-методологии, и XP - это единственное, что у нас работало эффективно, из числа non-scaled методик.

.. todo::

   here

      - PO is a stakeholder
      - https://t.me/emacsway_log/531
      - https://t.me/emacsway_log/480
      - https://t.me/emacsway_log/488

   TechDebt

      - https://t.me/emacsway_log/125
      - https://t.me/emacsway_log/130
      - https://t.me/emacsway_log/785
      - https://t.me/emacsway_log/662
      - https://t.me/emacsway_log/393

   Refactoring

      - https://t.me/emacsway_log/131
      - https://t.me/emacsway_log/132

   YAGNI

      - https://t.me/emacsway_log/135
      - https://t.me/emacsway_log/136

   common planning mistakes (in reguirements)

      - https://t.me/emacsway_log/458
      - https://t.me/emacsway_log/462

