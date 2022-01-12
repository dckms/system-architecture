:canonical-base-url: https://dckms.github.io/system-architecture

.. index::
   single: Technical Debt; as architecture options
   :name: emacsway-architecture-options

=============================
Architecture: Selling Options
=============================

.. sectionauthor:: Ivan Zakrevsky

.. contents:: Содержание

Gregor Hohpe увидел другую возможность донести представителям бизнеса стоимость архитектурных решений, используя терминологию фондового рынка, и разъясняет это на примере фондовых опционов.


    📝 "In the financial world, an option is well-known as the right, but not the obligation, to buy or sell a financial instrument at a future point in time (or over a future time span for American-style options).
    An option is therefore a way to defer a decision: instead of deciding to buy or sell a stock today, you have the right to make that decision in the future, at a known price.

    Any person involved in the financial industry knows that options aren’t free: there’s a whole market for buying and selling options and other derivatives.
    If there was any doubt, it was cleared out by Fischer Black and Myron Scholes, who managed to compute the value of an option with their famous Black-Scholes Formula (see `Wikipedia <http://en.wikipedia.org/wiki/Black%E2%80%93Scholes_model>`__).
    One important parameter in establishing the value of the option is the strike price, i.e. the price at which the stock can be purchased in the future.
    The lower this strike price, the higher the value of the option.

    <...>

    Just as with financial options, it’s important that the right to exercise the option in the future is tied to a known price.
    In the world of IT architecture this means that a future change or addition to the system can be made at the same or similar cost as doing it today.
    Following Black-Scholes, options whose strike price is higher than the stock’s current price still have value.
    So it’s OK if exercising the (architecture) option in the future has a slightly higher price than today.
    The value of the option originates from being able to defer the decision until you have more information while fixing the price."


    -- "`Architecture: Selling Options. How do you explain the value of architecture to business stakeholders? Deferring to the Nobel-prize winning economists Black and Scholes can work surprisingly well. <https://architectelevator.com/architecture/architecture-options/>`__" by Gregor Hohpe

.. seealso::

   - ":ref:`emacsway-compound-interest`"
