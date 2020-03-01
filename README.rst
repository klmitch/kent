====
Kent
====

.. image:: https://img.shields.io/github/tag/klmitch/kent.svg
    :target: https://github.com/klmitch/kent/tags
.. image:: https://img.shields.io/hexpm/l/plug.svg
    :target: https://github.com/klmitch/kent/blob/master/LICENSE
.. image:: https://travis-ci.org/klmitch/kent.svg?branch=master
    :target: https://travis-ci.org/klmitch/kent
.. image:: https://coveralls.io/repos/github/klmitch/kent/badge.svg?branch=master
    :target: https://coveralls.io/github/klmitch/kent?branch=master
.. image:: https://godoc.org/github.com/klmitch/kent?status.svg
    :target: http://godoc.org/github.com/klmitch/kent
.. image:: https://img.shields.io/github/issues/klmitch/kent.svg
    :target: https://github.com/klmitch/kent/issues
.. image:: https://img.shields.io/github/issues-pr/klmitch/kent.svg
    :target: https://github.com/klmitch/kent/pulls
.. image:: https://goreportcard.com/badge/github.com/klmitch/kent
    :target: https://goreportcard.com/report/github.com/klmitch/kent

This repository contains the go package kent.  The kent package
provides the concept of a ``Reporter``, which allows for the reporting
of errors and warnings in a variety of ways, including outputting to
an ``io.Writer``, to a ``log.Logger``, or just counting the number of
errors and warnings.  A ``Reporter`` is part of a chain of
``Reporter`` instances; the ``Reporter.Unwrap`` method may be used to
retrieve the next ``Reporter`` in the chain.  There is also an ``As``
function provided, which allows retrieving the first of a specified
type of reporter from the chain.

Warnings
========

The kent package provides a ``Warning`` interface.  A ``Warning`` is
an ``error`` with the addition of a ``Warning`` method, which returns
the same string as the ``Error`` method.  The function ``NewWarning``
constructs a new warning, and is analogous to ``errors.New``;
similarly, ``Warningf`` constructs a new warning from a format string,
and is analogous to ``fmt.Errorf``, including the ``%w`` format
conversion to wrap another error or warning.  There is also a
``WarningWrap`` function, which takes a normal ``error`` and wraps it
so that it will appear to be a ``Warning``.  Finally, the utility
function ``IsWarning`` checks whether an ``error`` is a ``Warning`` or
not, utilizing the ``errors.Unwrap`` utility function to explore all
errors in an error chain.

The ``Warning`` concept enables reporting of warnings through the
``Reporter`` support; a ``Warning`` is intended to be an error
condition that is not fatal to whatever procedure is reporting it.
This could, for instance, be used to report a compilation warning, or
the use of a deprecated construct in a configuration file.

Provided Reporters
==================

The kent package provides several useful ``Reporter`` implementations,
all of which are designed for thread safety.  The root of any
``Reporter`` tree should be the *root* reporter, which is returned by
the ``Root`` function; this is similar to the ``context.Context`` that
is returned by ``context.Background``.  Similarly, there is also a
*TODO* reporter, which is returned by the ``TODO`` function; the
concept is analogous to ``context.TODO``.  The root reporter does
nothing when ``Report`` is called, and has no child reporters; it
simply provides an anchor for the root of the ``Reporter`` tree.

The ``TeeReporter``, which can be constructed with a call to
``NewTeeReporter``, constructs a ``Reporter`` implementation that
allows the dynamic addition and removal of other ``Reporter``
instances.  Calls to ``Report`` are simply passed on to all of the
children of the ``TeeReporter`` instance.

The ``CountingReporter``, constructed with a call to
``NewCountingReporter``, constructs a ``Reporter`` implementation that
counts the number of errors and warnings that are passed to its
``Report`` method.  Besides then passing those errors and warnings on
to its child, the ``CountingReporter`` does nothing else with those
errors and warnings.  The number of errors can be retrieved using the
``Errors`` method, and ``Warnings`` returns the number of warnings.

The ``WritingReporter``, constructed with a call to
``NewWritingReporter``, constructs a ``Reporter`` implementation that
emits the error or warning--with a prefix consisting of "ERROR:" or
"WARNING:", as appropriate--to a specified ``io.Writer`` instance.
The error or warning is then passed on to the child reporter.

The ``LoggingReporter``, constructed with a call to
``NewLoggingReporter``, constructs a ``Reporter`` implementation that
emits the error or warning, with the same prefix as for
``WritingReporter``, but sends the message to either the default
``log.Logger`` or to a specified ``log.Logger`` instance.

Mocking Reporters
=================

To facilitate testing of code that utilizes ``Reporter``, the kent
package also provides ``MockReporter``, using the
``github.com/stretchr/testify/mock`` go package.
