===================================
Repository and Anticorruption Layer
===================================

..

    💬 In Implementing #DDDesign why use Anticorruption Layer rather than a Repository to integrate with another Bounded Context?

    Typically the Repository pattern is used as an Adapter in front of a database.
    In general you could think of a Repository as an AcL over a database.

    But then again, an AcL is at least both of these: (1) an Adapter, and (2) a Translator.
    Typically a Repository doesn't translate from database data into a model's data.
    Usually database data is at most reshaped when hydrating models. In other words, you think of an AcL more as a translator between two different model languages, where one language might not be carefully defined.
    **A Repository isn't translating between languages, just adapting database data shapes to/from domain model data shapes.**

    With AcL, make sure you have very good reason to duplicate data across context boundaries.
    In general replicating/duplicating data across boundaries is a bad thing.

    Every chapter of Implementing #DDDesign, and Chapter 13 in this case, demonstrates different ways of handling similar challenges. The Agile PM Context has a very high autonomy quality attribute.
    The Collaboration Context is designed earlier before SaaSOvation had a lot of experience with DDD.
    Decisions are never static and certainly imperfect.
    Continuous improvement matters.

    If you are going to deliver software that is useful and provides valuable customer outcomes, you must make decisions now and later.

    As with all software decisions, pattern choice depends. This thread should help make choices clearer.

    -- `Vaughn Vernon <https://twitter.com/VaughnVernon/status/1506090113582841859?s=20>`__

Другим отличительным признаком Repository от Adapter (`Ports & Adapter <https://alistair.cockburn.us/hexagonal-architecture/>`__) является их цель: Adapter реализует Port, который явно обозначает границу компонента, в то время как Repository имеет цель эту границу скрыть, и подделать удаленный источник локальным:

    💬 Mediates between the domain and data mapping layers using **a collection-like interface** for accessing domain objects.

    💬 A Repository mediates between the domain and data mapping layers, acting like an in-memory domain object collection.
    Client objects construct query specifications declaratively and submit them to Repository for satisfaction.
    **Objects can be added to and removed from the Repository, as they can from a simple collection of objects**, and the mapping code encapsulated by the Repository will carry out the appropriate operations behind the scenes.

    -- `Repository homepage <https://martinfowler.com/eaaCatalog/repository.html>`__
