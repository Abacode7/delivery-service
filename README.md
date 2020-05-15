# Delivery Service

A service to highlight a decent application structure in Go.

The cmd and pkg packages included follow the community standards in Go.

The cmd package contain server or command-lne versions of the application.

The pkg package contain all other project specific files.

Under the pkg package we have the domain package which contains structs
for the various entities in the application. It also includes a
contract (an Interface) for the Storage/Repository layer to follow.

Moving on, we have the storage sub-package which contains various 
storage implementations for the application. Observe that each implementation
has its own folder to ensure package level privacy, and they implement the
interface specified by the models.

Included also is the service package (excluded at times), which contains 
various domain level functions which aren't to be implemented by a single
entity alone. They use the contracts defined by the domain and inject the
storage/repo implementation at runtime for functions that involve storage.
They also define interfaces to be used by handlers/controllers that implement
them.
Observe that for this particular project services are divided according to
the domain (i.e user, courier), for other applications they're divided according to application
context (i.e in a beer reviewing app - the context could be adding, listing and reviewing).
Note that the repository and service interface have to be in sync
with this division.

Finally, the http package specifies the handler for the application. They make
use of the services to function. Note that this package is further divided based
on the architecture you desired to use e.g rest, rpc, e.t.c.
You could also include a top-level handler/controller folder, if you desire
different protocols in your application.
