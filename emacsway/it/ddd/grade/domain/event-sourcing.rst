:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: Event Sourcing; in Golang
   :name: emacsway-golang-event-sourcing

==============
Event Sourcing
==============

.. sectionauthor:: Ivan Zakrevsky

.. contents:: Содержание

Термины StreamId/StreamName, StreamType, StreamPosition не несут информативности в домене, и должны появляться в домене только `на стадии сохранения <https://github.com/VaughnVernon/IDDD_Samples/blob/05d95572f2ad6b85357b216d7d617b27359a360d/iddd_collaboration/src/main/java/com/saasovation/collaboration/port/adapter/persistence/repository/EventStoreCalendarRepository.java#L54>`__.

Иными словами, они не несут информации обработчикам доменных событий, которым нужен первичный идентификатор Агрегата, а не его сериализованное в строку представление.
