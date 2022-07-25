:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: CanExecute pattern; in Golang
   :name: emacsway-golang-can-execute

=============================================
Использовать ли в проекте CanExecute pattern?
=============================================

.. sectionauthor:: Ivan Zakrevsky

CanExecute pattern был описан в статье "`Validation and DDD <https://enterprisecraftsmanship.com/posts/validation-and-ddd/>`__" by Vladimir Khorikov.

Vladimir Khorikov - авторитетный специалист, смело принимающий на себя все риски первопроходца, который существенным образом повлиял на развитие индустрии, а также на мое становление как специалиста, за что я ему очень признателен.

CanExecute pattern, продемонстрированный Владимиром, является, действительно, очень удобным подходом, без которого очень сложно обойтись там, где это вызвано объективной необходимостью:

1. в распределенных процедурах (процессах);
2. там, где существует вероятность установления частично-валидного состояния композитного объекта, т.е. для обеспечения атомарности валидации с целью осуществимости атомарности изменения состояния.

В этом я убедился на собственной практике.

В других случаях, ключевой аргумент использования CanExecute pattern, приводимый в статье, сводится к CQS principle.

Я обсудил с ним этот вопрос, и он согласился с тем, что этот вопрос не совсем однозначный.

В этом подходе меня смущает тот факт, что образуется логическая зависимость - клиент класса действует исходя из предположения о том, что оба метода используют один и тот же инвариант.

В заметке ":ref:`emacsway-cqs-command-status-code`" этот вопрос уже затрагивался, поэтому, повторюсь:

    💬️ "It is important here two deal with two common objections to the side-effect-free style.

    The first has to do with error handling.
    Sometimes a function with side effects is really a procedure, which in addition to doing its job returns a status code indicating how things went.
    But there are better ways to do this; roughly speaking, the proper O-O technique is to **enable the client, after an operation on an object, to perform a query on the status, represented for example by an attribute of the object**, as in

    target.some_operation(...)

    how_did_it_go := target.status

    Note that the technique of returning a status as function result is lame anyway.
    It transforms a procedure into a function by adding the status as a result;
    **but it does not work if the routine was already a function, which already has a result of its own**.
    It is also problematic if you need more than one status indicator.
    In such cases the C approach is either to return a "structure" (the equivalent of an object) with several components, which is getting close to the above scheme, or to use global variables — which raises a whole set of new problems, especially in a large system where many modules can trigger errors."

    -- "Object-Oriented Software Construction" 2nd edition by Bertrand Meyer, chapter "23.1 SIDE EFFECTS IN FUNCTIONS"

Единственная причина, по которой Bertrand Meyer избегал возврата ошибок в то время, заключается в том, что тогда было невозможно возвратить одновременно и результат, и ошибку.
Сегодня таких проблем нет.
В Golang это поддерживается на уровне синтаксиса языка, а в других языках поддерживается объект Result.

Но даже если и разделять, то Bertrand Meyer рекомендует проверять ошибки после попытки, а не перед попыткой.

