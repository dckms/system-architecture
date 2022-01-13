:canonical-base-url: https://dckms.github.io/system-architecture

.. index:: YAGNI
   :name: emacsway-yagni


=====
YAGNI
=====

.. sectionauthor:: Ivan Zakrevsky

.. contents:: Содержание

..

    📝 "Если стоимость сегодняшнего решения высока, вероятность того, что оно окажется правильным, низка, вероятность того, что завтра вы найдете лучший способ решить проблему, высока, а стоимость внесения изменений в дизайн завтра низка, то мы можем прийти к выводу, что если сегодня мы можем обойтись без решения, значит, мы ни в коем случае не должны принимать это решение сегодня.
    Именно такой подход используется в рамках ХР.
    "Количество сложностей ровно на один день и не более того".

    Однако некоторые факторы могут стереть наши выводы в порошок.
    Если затраты, которые возникнут в случае, если мы будем принимать решение завтра, существенно больше сегодняшних, значит, мы должны принять решение сегодня в надежде на то, что завтра мы окажемся правы.
    Если инерция дизайна достаточно низка (над проектом работают очень-очень умные люди), значит, у дизайна, формируемого по мере разработки, остается все меньше и меньше преимуществ.
    Если вы действительно очень хороший провидец, значит, вы можете спроектировать все без исключения с самого начала, а затем приступать к реализации готового завершенного плана.
    Однако для всех остальных обычных людей я не вижу иной альтернативы, кроме той, в рамках которой предлагается проектировать сегодня только то, что требует проектирования именно сегодня, и откладывать на завтра то, что можно спроектировать завтра.

    If the cost of today's decision is high, and the probability of its being right is low, and the probability of knowing a better way tomorrow is high, and the cost of putting in the design tomorrow is low, then we can conclude that we should never make a design decision today if we don't need it today.
    In fact, that is what XP concludes.
    "Sufficient to the day are the troubles thereof."

    Now, several factors can make the above evaluation null and void.
    If the cost of making the change tomorrow is very much higher, then we should make the decision today on the off chance that we are right.
    If the inertia of the design is low enough (for example, you have really, really smart people), then the benefits of just-in-time design are less.
    If you are a really, really good guesser, then you could go ahead and design everything today.
    For the rest of us, however, I don't see any alternative to the conclusion that today's design should be done today and tomorrow's design should be done tomorrow."

    -- "Extreme Programming Explained" 1st edition by Kent Beck, "Chapter 17. Design Strategy", перевод ООО Издательство "Питер"


.. seealso::

   - ":doc:`/emacsway/it/sdlc/uncertainty-management/adaptation/crash-course-in-software-development-economics`"
