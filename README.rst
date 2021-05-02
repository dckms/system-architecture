=================
Как пользоваться?
=================

.. contents:: Содержание


Почему?
=======

Объем информации в современном мире стал достаточно большим и возникла потребность этим объемом как-то управлять.

Опыт показал, что высокое качество объемного материала легче достигают **коллективы авторов**, как, например, "`.NET Microservices: Architecture for Containerized .NET Applications <https://docs.microsoft.com/en-us/dotnet/architecture/microservices/>`__" (`более 4.5 тыс. форков <https://github.com/dotnet/docs>`__), "`Azure Architecture Center <https://docs.microsoft.com/en-us/azure/architecture/>`__ (`более 1.1 тыс. форков <https://github.com/MicrosoftDocs/architecture-center>`__).
Можно заметить, что некоторые пользователи ведут автономные версии форков, периодически вливая в них главный бранч.
Т.е., пользователи стремятся создавать **собственные пространства знаний**, которые **обогащаются** из других источников и **дистиллируются** от всего лишнего.

Это наводит на мысль - а что, если совместить персональную систему управления знаниями с коллективной системой?
Вы сами решаете, какими именно своими заметками вы хотите поделиться, а какие вы хотите добавить в свою базу от других авторов.

Основная идея данной системы заключается в том, что **субъектом субъектом знаний в такой моделе является именно пользователь**, в то время, как в большинстве других систем управления знаниями на основе wiki-движков, пользователь является, скорее, объектом системы, что порождает ряд вопросов о редакционной политике, достижении консенсуса и разрешении противоречий, выработке норм поведения, управлении правами, защите от диверсий путем уничтожения содержимого или постановки информационных помех.

Распределенный подход к хранению знаний избавляет от этих проблем - вы сами решаете, чем поделиться, и от кого и что добавить в свою базу знаний.
Ваша база знаний не обязана следовать чьим-то чужим точкам зрения. Все как в реальной жизни.

Другой фактор, который поспособствовал появлению этой системы, заключается в том, что сегодня человек живет в условиях стесненного времени.
Каждому из нас есть чем поделиться и обогатить коллективные знания, но не у каждого есть время прорабатывать по этой теме статьи.
Именно поэтому короткие заметки стали сегодня так популярны.
В конечном итоге, люди тратят на Telegram (и другие мессенджеры) намного больше времени, чем требуется на написание статей.
Вопрос только в том, что в Telegram они это делают короткими интервалами времени.
Однако, в Telegram их информация быстро теряется в безорганизованной свалке информации, и, несмотря на сам факт своего существования, эта информация быстро становится бесполезной в условиях отсутствия навигации.

Возникло  противоречие: там, где можно структурировать, - там не пишут, а там, где пишут, - структурировать нельзя.
Можно ли это противоречие разрешить?

Цель данного проекта заключается в разрешении этого противоречия, путем обеспечения навигации в распределенной коллективной информации коротких сообщений (заметок).
Благодаря RSS-каналу, новые сообщения можно легко отражать в Telegram-channel или в другие мессенджеры, например, в Mattermost, посредством ботов и плагинов к мессенджерам.
Таким образом достигается и цель уведомления о новых сообщениях, и сохраняется навигация по сообщениям.

Система представляет собой минималистичный набор принципов и соглашений, реализованный на Open Source системе документирования `Sphinx-doc <https://www.sphinx-doc.org/>`__ и использущий reStructuredText и Markdown форматы разметок.
Sphinx-doc предоставляет и тегирование (директива "`index <https://www.sphinx-doc.org/en/master/usage/restructuredtext/directives.html#index-generating-markup>`__"), и перекрестные ссылки, и Table of Content (ToC, директива "`toctree <https://www.sphinx-doc.org/en/master/usage/restructuredtext/directives.html#table-of-contents>`__"), и неограниченную иерархию файлов, и перекрестные связи между иерархиями файлов и иерархиями ToC, и клиентский полнотекстовый поиск (средствами JS браузера), и `TODO <https://www.sphinx-doc.org/en/master/usage/extensions/todo.html>`__-листы, и RSS (в виде стороннего расширения), и подсветку синтаксиса языков программирования, и расширяемость с большим количеством готовых к использованию расширений.


Порядок использования
=====================

Система работает следующим образом:

#. Создайте форк репозитария.
#. Перейдите в приватный бранч "private".
#. Свои приватные заметки ведите в пространстве имен "private" (``/private``, ``_media/private``, ``_html_extra/private``).
#. Создайте свой публичный бранч, например, "ivan.ivanov". Приватные директории сразу же внесите в файл ".gitignore".
#. Создайте пространство имен для своих публичных заметок, которыми вы хотите поделиться, например, "ivan.ivanov" (``/ivan.ivanov``, ``_media/ivan.ivanov``, ``_html_extra/ivan.ivanov``). Таким образом вы облегчите читателям навигацию по вашим заметкам и сохраните очевидность авторства за собой. Это необходимо еще и потому, что древовидное устройство файловой системы сложно унифицировать для всех авторов - у каждого автора есть свое видение на классификацию его материала. Благодаря гибкости директивы "`toctree <https://www.sphinx-doc.org/en/master/usage/restructuredtext/directives.html#table-of-contents>`__", вы легко можете включать в дерево своего содержания поддеревья или страницы других авторов.
#. Тегируйте свой материал с помощью директивы "`index <https://www.sphinx-doc.org/en/master/usage/restructuredtext/directives.html#index-generating-markup>`__"
#. Ненужные вам заметки других авторов вы можете удалить в своем приватном бранче. А нужные - добавить, как целиком, так и выборочно, используя `cherry-pick <https://git-scm.com/docs/git-cherry-pick>`__.
#. Используя `UUID4 <https://www.uuidgenerator.net/version4>`__, создавайте `перекрестные ссылки <https://www.sphinx-doc.org/en/master/usage/restructuredtext/roles.html#ref-role>`__ между связанными заметками, следуя лучшим практикам `Zettelkasten <https://zettelkasten.de/posts/overview/>`__.
#. Ведите `TODO <https://www.sphinx-doc.org/en/master/usage/extensions/todo.html>`__.
#. Создайте Pull Request из своего именного публичного бранча ("ivan.ivanov") в trunk-branch. Может быть множество trunk-бранчей, и, в качестве одного из них, можете использовать `этот <https://github.com/dckms/dckms-trunk>`__. Trunk-branch можно сравнить с шиной событий в Event Sourcing системе.
#. Стройте свою распределенную коллективную базу знаний.

Можно добавить, что GitHub планирует добавить `поддержку cherry-pick в свой web-интерфейс <https://github.com/isaacs/github/issues/629>`__, а в `Desktop-client она уже реализована <https://github.blog/2021-03-30-github-desktop-now-supports-cherry-picking/>`__.
А вот GitLab уже реализовал `поддержку cherry-pick в web-интерфейсе <https://docs.gitlab.com/ee/user/project/merge_requests/cherry_pick_changes.html>`__.


Как собрать html?
=================

#. Если не установлен Python, то `установите его <https://docs.python.org/3/installing/index.html>`__.
#. Установите зависимости. Для этого, из корневой директории проекта выполните команду: ``pip install -r requirements.freeze.txt``
#. Отредактируйте файл conf.py, подробности смотрите в `документации <https://www.sphinx-doc.org/en/master/usage/configuration.html>`__.
#. Произведите сборку: ``make html``
#. Подробнее `здесь <https://www.sphinx-doc.org/en/master/usage/quickstart.html>`__.


О Zettelkasten
==============

- `Zettelkasten <https://zettelkasten.de/posts/overview/>`__
- `The Introduction to the Zettelkasten Method <https://zettelkasten.de/introduction/>`__
- `Как я веду Zettelkasten в Notion уже год: стартовый набор и полезные трюки <https://habr.com/ru/post/509756/>`__
- `Zettelkasten: как один немецкий учёный стал невероятно продуктивным <https://habr.com/ru/post/508672/>`__


Философия
=========

Основные принципы системы:

- минимизация рисков и внешних зависимостей (от конкретного типа текстового редактора, вендора)
- минимализм
- неограниченная расширяемость
- автономность
- субъектность пользователя и полный контроль над информацией
- распределенность и коллективность
- свободное обогащение и дистилляция информации


Близкие по духу системы
=======================


Obsidian
--------

    In our age when cloud services can shut down, get bought, or change privacy policy any day, the last thing you want is proprietary formats and data lock-in.

    With Obsidian, your data sits in a local folder. Never leave your life's work held hostage in the cloud again.

    Plain text Markdown also gives you the unparalleled interoperability to use any kind of sync, encryption, or data processing that works with plain text files.

    https://obsidian.md/


Neuron Zettelkasten
-------------------

    Neuron was designed with these criteria in mind:

    - Future-proof: store notes locally1 as plain-text (Markdown) files
    - Not tied2 to a single text editor
    - Statically generated web site, for browsing and publishing on the web
    - Remain as simple to use as possible, whilst being feature-rich via Plugins

    https://neuron.zettel.page/philosophy


Markdown
========

Markdown - популярный язык разметки.
Приводимые в начале этой страницы архитектурные руководства Microsoft написаны на Markdown.

Вы легко можете использовать Markdown, благодаря расширению `MyST-Parser <https://myst-parser.readthedocs.io/en/latest/>`__ (`порядок установки <https://www.sphinx-doc.org/en/master/usage/markdown.html>`__).
Расширение позволяет использовать в Markdown все директивы и роли Sphinx-doc, и является мостом Docutils к `markdown-it-py <https://github.com/executablebooks/markdown-it-py>`__, который поддерживает синтаксис `CommonMark <https://commonmark.org/>`__.

Как вариант, возможна и обычная статическая конвертация Markdown в reStructuredText:

- `m2r <https://github.com/miyakogi/m2r>`__ - Markdown to reStructuredText converter 
- `mdToRst <https://github.com/kata198/mdToRst>`__ - tool and library to convert markdown [md] to restructed text [rst] (md to rst).


Как работать на мобильных устройствах
=====================================


Как работать на Android
-----------------------

- Markor - популярный Markdown-редактор на Android: `GitHub <https://github.com/gsantner/markor>`__, `F-Droid <https://f-droid.org/packages/net.gsantner.markor>`__, `Google Play <https://play.google.com/store/apps/details?id=net.gsantner.markor>`__.
- `Termux <https://termux.com/>`__ - a unix-like environment for Android, for git and python3.
- `GitJournal <https://gitjournal.io/>`__ - mobile first Markdown Notes integrated with Git: `GitHub <https://github.com/GitJournal/GitJournal>`__, `Google Play <https://play.google.com/store/apps/details?id=io.gitjournal.gitjournal&pcampaignid=website>`__.


Как работать на iPhone
----------------------

- `GitJournal <https://gitjournal.io/>`__ - mobile first Markdown Notes integrated with Git: `GitHub <https://github.com/GitJournal/GitJournal>`__, `App Store <https://apps.apple.com/app/gitjournal/id1466519634>`__.
- `Working Copy <https://apps.apple.com/ca/app/working-copy-git-client/id896694807>`__ - a Git client.
- `1Writer <https://1writerapp.com/>`__ - powerful, beautiful Markdown editor for iOS.
- `iA Writer <https://ia.net/writer>`__ - the simple, award-winning design of iA Writer delivers the essential writing experience.
- `Editorial <https://www.omz-software.com/editorial/>`__ is a plain text editor for iOS with great Markdown support and powerful automation tools.
- `Editorial-obsidian <https://tekacs.github.io/editorial-obsidian/>`__ - Editorial scripts for Obsidian (unofficial): `GitHub <https://github.com/tekacs/editorial-obsidian>`__.
- `iTextEditors <https://brettterpstra.com/ios-text-editors/>`__ - the iOS Text Editor roundup.


Аналоги
=======

`Zettelkasten <https://github.com/roalyr/zettelkasten>`__ - a template for a Zettelkasten based on markdown files.

Neuron Zettelkasten может представлять интерес для тех, кто предпочитает минимизацию внешних зависимостей, минимализм и неограниченность:

- https://neuron.zettel.page/philosophy
- https://neuron.zettel.page/tutorial
- https://srid.github.io/neuron-template/README
- https://github.com/srid/neuron
- https://github.com/srid/neuron-template

Парень дает `сравнение Neuron Zettelkasten и Sphinx-doc <https://lobste.rs/s/kydg6q/neuron_0_4_zettelkasten_note_management#c_me2hhh>`__.


Интеграция с другими системами
==============================

Интеграция с другими системами, сервисами и приложениями возможна в пределах пересекающегося подмножества поддерживаемого Markdown-синтаксиса.


Интеграция с Obsidian
---------------------

Идея Obsidian так же построена на локальных Markown-файлах, но с GUI-клиентом (недавно появился и `мобильный клиент <https://help.obsidian.md/Obsidian/Mobile+app+beta>`__).
Теоретически это означает, что вы можете шарить файлы между двумя системами.
На практике я не пробовал это сделать (если попробуете - расскажите, пожалуйста, как получилось).

Зато сообщество Obsidian `дает много дельных советов <https://forum.obsidian.md/t/how-do-i-work-with-obsidian-on-mobile/471>`__, как работать с Markdown-файлами на мобильных устройствах.


Интеграция с Notion
-------------------

Notion позволяет экспортировать содержимое в Markdown-файлы.
Теоретически это означает, что вы можете шарить файлы между двумя системами.
На практике я не пробовал это сделать (если попробуете - расскажите, пожалуйста, как получилось).
Массового импорта в Notion я не встречал, но есть варианты, например `Notion.so Markdown Importer <https://github.com/Cobertos/md2notion/>`__.


Интеграция с Evernote
---------------------

Существуют решения для экспорта заметок из Evernote:

- `evernote2md <https://github.com/wormi4ok/evernote2md>`__ - convert Evernote .enex files to Markdown.
-  `ever2simple <https://github.com/claytron/ever2simple>`__ - migrate from evernote to simplenote with markdown formatting.
-  `ever2text <https://github.com/nicholaskuechler/ever2text>`__ - convert Evernote exports to text files.


RSS
---

Существует несколько коробочных решений RSS-feed для Sphinx:

- https://github.com/sphinx-contrib/yasfb
- https://github.com/sphinx-contrib/feed
- https://github.com/lsaffre/sphinxfeed
- https://github.com/prometheusresearch/sphinxcontrib-newsfeed

Смотрите так же https://github.com/sphinx-doc/sphinx/issues/2


Интегация RSS с Mattermost
~~~~~~~~~~~~~~~~~~~~~~~~~~

- https://integrations.mattermost.com/rssfeed-plugin/
- https://github.com/wbernest/mattermost-plugin-rssfeed


Интегация RSS с Telegram
~~~~~~~~~~~~~~~~~~~~~~~~

- https://github.com/BoKKeR/RSS-to-Telegram-Bot
- https://thefeedreaderbot.com/ ( https://telegram.me/TheFeedReaderBot )
- https://www.integromat.com/en/integrations/rss/telegram
- https://core.telegram.org/bots/faq


Послесловие
===========

Проект в состоянии развития. Стабильность пока не гарантируется.

Технически, в отдаленной перспективе можно было бы приспособить под принципы и соглашения системы одно из Open Source приложений для заметок, но у меня такая цель на данный момент не стоит. В таком приложении можно было бы выбирать источники подписок, автоматизировать и облегчить просмотр и принятие коммитов в свою базу знаний, например, если коммит содержит новую заметку, связанную с одной из уже принятых ранее заметок, или является её обновлением, тогда принимать коммит автоматически.