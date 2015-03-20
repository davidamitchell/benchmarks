accounting_service
==================
# Thoughts
* Made up of many smaller micro services which can be pipelined together.
* These smaller micro services will provide functionality simular to what is seen in __middleware__
  * persistance
  * event capture
  * validation
  * translation
  * lookups
  * caching
  * etc...
* These micro services (MS) will be able to communicate via __tcp__, __ipc__, __http__, __https__, or simple unix pipes.
* Each MS will only except one type of message and output one type of message.
* Each MS must be designed in a way which will allow it to work on the same server or different servers
* They must be designed in a way to allow them to be running 1..n number of instances
* The number of instances one MS is running must be transparent to the other MS


# Questions
* How/when to provide a response to an external initiator of a message (i.e. a user of the externally exposed api)

# Purpose
The purpose of this package is to provide the framework (and beginnings of an implemeations) for a system which would allow for plug-able microservices (MS).
> A note on the term microservice (MS): These could also be thought of as building blocks of the larger system/service (accounting service).  Perhaps __component deamons__ might be a better description.

The thought is that a __domain topography__ could be _configured_ (see section on configuration) by piecing together these microservices to achieve the desired goal.  

# Overview

The process of any message which comes into an accounting system would likely go through a set of actions.  Firstly the message must be captured, then any authentication is done.  Afterwards there will likely be some validation of some sort followed by a translation.  At this point there is likely to be some lookups and business logic applied.  These lookups may be to external services and could be quite time intensive, so these might be cached.

The process of any message to an accounting system asking for information would likely go through simular set of actions.  One starts with message capture then authentication and validation of the message.  At this point there may be a small bit of transactions followed by a lookup.  Then a further translation and perhaps a callback to an external hook.

# Outline
Two types of messages dealt with in the context of the larger service:
* Incoming messages _sending of information_ `<-`
* Outgoing messages _asking for information_ `->`

Below I will use `->` for MS which are used in __outgoing__ messages, `<-` for __incoming__ and `<->` for those which are able to be used in both

### Event capture
These should be simple endpoints which can deal with capture of a message.
### Authentication
Additionally, very simple MS here to preform authentication.  It could possibly require a lookup/cache service.
### Validation
Any means of validation.  Also may require a lookup/cache service.
### Translation
Simple transaltion of the payload from one form to another, again, lookup/cache service most likely required.
### Lookups/caching
Services for dealing with lookups each would require caching to allow for rapid responses.  These lookups could be external or internal.
### Business logic
Kind of the krux of the whole system.  Think of this as decision gates which would call off to other translation services.  Really the only MS which should be able to make decisions.
### Persistance
Deals with the persistance of the messages from the event capture.
### Notifications
Deals with the sending of information to other systems. Think webhooks.


# Configuration

Ideally would like to have a GUI for building a __domain topography__.  But a config file would also suffice.

---------------------------------------

# Proof of concept
Looking for a simple problem which would have this type of solution. Which would use only two of these microservices.  Would have to define at least three so we can produce two __domain topographies__.  Perhaps 4 total.  Two event captures, and two translations.
