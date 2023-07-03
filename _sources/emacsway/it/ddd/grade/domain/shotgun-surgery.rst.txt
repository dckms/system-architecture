:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   Shotgun Surgery


===============
Shotgun Surgery
===============

Может показаться, что используя Raw-SQL мы обретаем классифицированный Code Smell, известный как Shotgun Surgery (Разлет Дроби), ибо добавление одного поля в Сущность требует правки многих файлов.
Есть два способа решить эту проблему (и снизить Coupling), о которых писал Martin Fowler в главе "Metadata Mapping" книги "Patterns of Enterprise Application Architecture": "reflective program" и "code generation", причем, сам он лично предпочитает второй вариант:

    Generated code is more explicit so you can see what's going on in the debugger;
    as a result I usually prefer generation to reflection,
    and I think it's usually easier for less sophisticated developers
    (which I guess makes me unsophisticated).

    -- "Patterns of Enterprise Application Architecture" by Martin Fowler, David Rice, Matthew Foemmel, Edward Hieatt, Robert Mee, Randy Stafford

Подавляющее большинство ORM использует "reflective program", в то время, как в Golang-сообществе традиционно широко применяется в практике "code generation".

В данном проекте, на определенном этапе развития, появится инструмент кодогенерации по образу `sqlc <https://github.com/kyleconroy/sqlc>`__.

