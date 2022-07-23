:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: Encapsulation; in Golang
   :name: emacsway-encapsulation-golang

====================================================
Как сохранить Агрегат в БД не разрушая инкапсуляции?
====================================================

.. sectionauthor:: Ivan Zakrevsky

Инварианты лишены смысла своего существования в условиях дырявой, как решето, инкапсуляции.
Вопрос в том, как сохранить инкапсуляцию Агрегатов в Golang, когда нам требуется его внутреннее состояние для формирования SQL-запроса, или, наоборот, требуется установить состояние Агрегата в результате выполнения SQL-запроса.

.. contents:: Содержание


Memento pattern
===============

Memento оказался близко, но не по назначению. Суть Memento в том, что он не должен раскрывать свое состояние никому, кроме своего создателя:

    1. Preserving encapsulation boundaries. Memento avoids exposing information that only an originator should manage but that must be stored nevertheless outside the originator.
       The pattern shields other objects from potentially complex Originator internals, thereby preserving encapsulation boundaries.

    -- "Design Patterns: Elements of Reusable Object-Oriented Software" by Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides


Walker
======

Walker представляет собою модификацию паттерна Visitor с целью сохранить инкапсуляцию Агрегатов. К числу недостатков паттерна паттерна Visitor относится:

    6. Breaking encapsulation. Visitor's approach assumes that the ConcreteElement interface is powerful enough to let visitors do their job.
    As a result, the pattern often forces you to provide public operations that access an element's internal state, which may compromise its encapsulation.

    -- "Design Patterns: Elements of Reusable Object-Oriented Software" by Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides

Что будет создавать Walker в случае обхода иерархической структуры Агрегата с несколькими вложенными Сущностями?
Вероятно, это будет будет несколько SQL-запросов с параметрами, т.е. некий композитный объект, выраженный некой структурой данных.
Это лишает смысла использование Visitor, если можно сразу возвратить структуру данных, причем, абстрагированную от SQL.

Технически, можно сделать так:

.. code-block:: go

    type Walkable interface {
        Accept(Walker)
    }

    type Walker interface {
        SetField(string)
        WalkWalkable(Walkable)
        WalkUint8(uint8)
        WalkUint64(uint64)
        WalkUint(uint)
        WalkTime(time.Time)
    }


.. code-block:: go

    func (e Endorsement) Accept(walker interfaces.Walker) {
        walker.SetField("recognizerId")
        walker.WalkWalkable(e.recognizerId)
        walker.SetField("recognizerGrade")
        walker.WalkWalkable(e.recognizerGrade)
        // ...
    }

.. code-block:: go

    func (id RecognizerId) Accept(walker interfaces.Walker) {
        walker.WalkUint64(id.Value())
    }


Проблема в том, что на этапе создания SQL-запроса нам пока еще могут быть неизвестны первичные ключи Агрегатов, чтобы их можно было бы проставить в SQL-запросы их Сущностей.

Кроме того, осведомленность о способах образования SQL-запросов размазывается между Walkers и Repositories, что вызывает "Разлет Дроби" (Code Smell) в случае изменения существующего способа построения SQL (например, в случае изменения диалекта БД или в случае внедрения какого-либо QueryBuilder).
Walker начинает быть слишком осведомленным о деталях реализации Repository.

В целях достижения DRY возникает целесообразность возложить на Walker генерирование только части SQL, и освободить его от осведомленности знания таблиц в БД, что будет разрывать обязанность за построение SQL на несколько объектов и подрывать Cohesion.


Valuer & Scanner
================

- `Valuer <https://pkg.go.dev/database/sql/driver#Valuer>`__
- `Scanner <https://pkg.go.dev/database/sql#Scanner>`__

Интерфейс Scanner открывает дверь к изменяемости ValueObject, что противоречит основной его сути.
А так же открывает брешь в инкапсуляции Агрегата.

С другой стороны, Valuer может возвращать только примитивные типы, а значит, он не пригоден для экспорта иерархической структуры состояния Агрегата:


    It is either nil, a type handled by a database driver's NamedValueChecker interface, or an instance of one of these types:

    - int64
    - float64
    - bool
    - []byte
    - string
    - time.Time

    -- `Источник <https://pkg.go.dev/database/sql/driver#Value>`__


Reflection
==========

В документации `отсутствуют <https://pkg.go.dev/reflect#Value.FieldByName>`__ какие-либо упоминания об ограничении доступа к защищенным атрибутам структуры данных посредством рефлекции.

Может быть через рефлексию и заработало бы - я не пробовал.
Но использовать рефлексию в production для таких целей как-то не сильно хочется, в т.ч. и по соображениям производительности.
К тому же этот метод является, по сути, еще одним способом пробить брешь в инкапсуляции.


Exporter
========

Ссылки по теме:

- "`More on getters and setters <https://www.infoworld.com/article/2072302/more-on-getters-and-setters.html>`__" by Allen Holub
- "`Save and load objects without breaking encapsulation <https://stackoverflow.com/questions/24921227/save-and-load-objects-without-breaking-encapsulation>`__" at Stackoverflow

Идею можно посмотреть на примере:

.. code-block:: java
   :caption: `Example by Allen Holub <https://www.infoworld.com/article/2072302/more-on-getters-and-setters.html>`__
   :name: emacsway-code-exporter-example-1
   :linenos:

    import java.util.Locale;

    public class Employee
    {   private Name        name;
        private EmployeeId  id;
        private Money       salary;

        public interface Exporter
        {   void addName    ( String name   );
            void addID      ( String id     );
            void addSalary  ( String salary );
        }

        public interface Importer
        {   String provideName();
            String provideID();
            String provideSalary();
            void   open();
             void   close();
        }

        public Employee( Importer builder )
        {   builder.open();
            this.name   = new Name      ( builder.provideName()     );
             this.id     = new EmployeeId( builder.provideID()       );
             this.salary = new Money     ( builder.provideSalary(),
                                      new Locale("en", "US") );
            builder.close();
        }

        public void export( Exporter builder )
        {   builder.addName  ( name.toString()   );
             builder.addID    ( id.toString()     );
            builder.addSalary( salary.toString() );
        }

        //...
    }

Или на более лаконичном примере:

.. code-block:: java
   :caption: `Example from Stackoverflow <https://stackoverflow.com/questions/24921227/save-and-load-objects-without-breaking-encapsulation>`__
   :name: emacsway-code-exporter-example-2
   :linenos:

    interface PersonImporter {

        public int getAge();

        public String getId();
    }

    interface PersonExporter {

        public void setDetails(String id, int age);

    }

    class Person {

        private int age;
        private String id;

        public Person(PersonImporter importer) {
            age = importer.getAge();
            id = importer.getId();
        }

        public void export(PersonExporter exporter) {
            exporter.setDetails(id, age);
        }

    }

Замечательный вариант, но проблема в том, что он использует интерефейсы, и это получается слишком многословно - требуется декларировать сам тип (структуру), интерфейс, сеттеры.
Вряд ли кто-то будет этим заниматься, когда можно просто обязать Агрегат вернуть простую структуру.

:ref:`Второй <emacsway-code-exporter-example-2>` из приведенных примеров содержит пакетированный сеттер, что делает его несколько менее многословным.
В проекте он реализован на примере Агрегатов ``Recognizer.ExportTo(ex RecognizerExporter)`` и ``Endorsed.ExportTo(ex EndorsedExporter)``.

Внутренние объекты структуры Exporter должны иметь сеттеры.
Реализовать эти объекты можно по разному - как структуру или как примитивный тип.
Но в первом случае у нас даже плоская Сущность будет экспортировать двухуровневую структуру, что вносит неудобство, но при этом структура позволяет использовать параметрический полиморфизм (generics) для повышения ее реиспользования.
А во втором случае, хотя объект Exporter и будет "наследоваться" от примитивного типа, но будет требовать преобразование типов в определенных случаях, что также вносит определенное неудобство.
Второй вариант реализован в проекте на примере типов ``Uint8Exporter``, ``Uint64Exporter`` и др.

Немного смущает смешивание парадигм FP и OOP для ValueObject.
Хотя ValueObject и остается неизменяемым, но сам факт того, что функционально чистый объект вызывает мутирующие методы другого объекта, вызывает состояние легкого смущения.
Возникает вопрос - почему функционально чистый объект не может просто взять и вернуть другой функционально чистый объект?
Если бы Golang поддерживать Generics для методов, тогда могло бы получиться что-то похожее на ``Recognizer.Export[T](exporterFactory function(attr1, attr2, attr3) T) T``.

И несколько многословным получается использование такого способа экспорта состояния в тестовых кейсах: нужно проинициализировать переменную, заэкспортировать в нее состояние, осуществить это дважы и для ожидаемой переменной и для тестируемой переменной, и лишь потом сравнить.
На каждый assert получается по четыре лишние строчки кода (по три, если пакетировать ининциализацию переменных).
В проекте это хорошо видно на сравнении тестовых кейстов ``TestRecognizerExport`` и ``TestRecognizerExportTo``.

Можно было бы сказать, что тестировать нужно по принципам :ref:`черного ящика <emacsway-tdd-black-box>`, т.е. только внешнее поведение.
Совершенно верно, но только нам требуется не только внешнее поведение, но и достоверность сохранения введенной в конструктор Агрегата информации в БД.

Возникает целесообразность облегчить метод экспортирования, придав ему сигнатуру ``Recognizer.Export() RecognizerState`` вместо ``Recognizer.ExportTo(ex RecognizerExporter)``.
Получится что-то типа DTO с тем лишь отличием, что он пересекает не сетевые границы, а границы инкапсуляции Агрегата.
Пока что этот вариант выглядит чуть более привлекательным.

В проекте реализованы оба варианта для наглядности, на примере Агрегатов Recognizer и Endorsed, а также Cущности Endorsement.

