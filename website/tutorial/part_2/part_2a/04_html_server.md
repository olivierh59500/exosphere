<table>
  <tr>
    <td><a href="03_microservices.md"><b>&lt;&lt;</b> microservices</a></td>
    <th>The Web Server Service</th>
    <td><a href="05_communication.md">communication <b>&gt;&gt;</b></a></td>
  </tr>
</table>


# The HTML Server Service

Let's build the first service for our application:
the HTML server!
If we were building our Todo app as a traditional monolith,
this would be the only code base
and perform all of the application's functionality:
receiving requests,
storing todo items, user accounts, and API tokens in the database,
configuring and using the search engine,
tracking user sessions and permissions,
rendering HTML for the browser, JSON for the REST API,
sending out emails,
etc.

In Exosphere's microservice world,
each code base has only one responsibility.
The html server's job is to interact with the user via an HTML UI.
Most of the things mentioned above are not a direct part of this responsibility,
and are therefore implemented outside of the html server,
as separate services.
Because of this much narrower set of responsibilities,
the html server is a lot smaller and simpler
than it would be in a traditional monolithic application.

Since our html server is so simple,
we'll build it using [ExpressJS](http://expressjs.com).
Exosphere provides a template for building ExpressJS html servers.
Let's use it:

<a class="tr_runConsoleCommand">
```
$ cd todo-app
exo add service --templateName=htmlserver-express-es6
```

Again, the generator asks for the information it needs interactively.
Please enter:

<table>
  <tr>
    <th>prompt</th>
    <th>text you enter</th>
  </tr>
  <tr>
    <td>Name of the service to add</td>
    <td>html-server</td>
  </tr>
  <tr>
    <td>Author</td>
    <td>test-author</td>
  </tr>
  <tr>
    <td>Description</td>
    <td>serves the HTML UI of the Todo app</td>
  </tr>
  <tr>
    <td>Name of the data model</td>
    <td></td>
  </tr>
</table>

</a>

Now we see the service registered in our application configuration file:

<a class="tr_verifyWorkspaceFileContent">
__todo-app/application.yml__

```yml
name: todo-app
description: An example Exosphere application
version: 0.0.1

services:
  public:
    html-server:
      location: ./html-server
```

</a>

Here is the current architecture of our application:

<table>
  <tr>
    <td width="280">
      <img alt="architecture for step 2" src="04_architecture.png" width="258">
    </td>
    <td>
      <ol>
        <li>
          The user browses to our homepage.
          In order to show that page, her web browser requests the HTML for it.
        </li>
        <li>
          This request goes to our <i>html server service</i>.
          It replies with the HTML for the page.
        </li>
      </ol>
    </td>
  </tr>
</table>



## The html service folder

The html service is located in a subdirectory of the application,
in <a class="tr_verifyWorkspaceContainsDirectory">`todo-app/html-server/`</a>.
This makes sense because it is an integral part of our application,
and doesn't make sense outside of it.

Most of the files in this folder
are just a normal [ExpressJS](http://expressjs.com) application,
plus some extra tools like linters.
Let's check out how the service looks like internally.


### The service configuration file

This file tells the Exosphere framework everything it needs to know about this service.

<a class="tr_verifyWorkspaceFileContent">
__todo-app/html-server/service.yml__

```yml
type: html-server
description: serves the HTML UI of the Todo app
author: test-author

# defines the commands to make the service runnable:
# install its dependencies, compile it, etc.
setup: yarn install

# defines how to boot up the service
startup:

  # the command to boot up the service
  command: node app

  # the string to look for in the terminal output
  # to determine when the service is fully started
  online-text: HTML server is running

# the messages that this service will send and receive
messages:
  sends:
  receives:

# other services this service needs to run,
# e.g. databases
dependencies:

docker:
  publish:
```
</a>

The other files in this directory are just a normal
[ExpressJS](http://expressjs.com)
application.


## Setting up the service

With all files in place,
the Exosphere CLI has all the information to set up our application.
Let's check that the overall configuration is correct,
and have Exosphere set up the service for us:

<a class="tr_runConsoleCommand">
```
$ cd todo-app
$ exo setup
```
</a>

We see how it uses Node's package management system (NPM)
to download and install
the external [ExpressJS](http://expressjs.com) and [Pug](http://jade-lang.com/) (formerly Jade) modules for us,
so that the service is ready to run.
The output should look something like:

<a class="tr_verifyRunConsoleCommandOutput">
```
Setting up todo-app 0.0.1
html-server  starting setup
html-server  setup finished
html-server  preparing Docker image
html-server  Docker setup finished
  exo-setup  setup complete
```
</a>


## Booting up the application

To test that everything works, let's check that the application boots up:

<a class="tr_startConsoleCommand">
```
$ cd todo-app
$ exo run
```
</a>

It prints this output:

<a class="tr_waitForOutput">
```
Running todo-app 0.0.1

exocom  Exosphere Development Communications server
exocom  Ctrl-C to stop
html-server  ExoRelay for 'html-server' online
html-server  HTML server online
html-server  HTML server is running
exo-run  'html-server' is running
exo-run  all services online
```
</a>

The Exosphere framework itself is written as a bunch of loosely coupled services.
We see a number of them in action here:
* __exorun__ is the command that runs Exosphere applications.
  It starts the other services.
* __html-server__ is our html server service.
  We can see that exorun starts it,
  and recognizes right after the output `Todo app running at port 3000`
  that our html server is online.
* The Exosphere runtime also starts a service called __exocom__.
  This is the messaging system
  for communication between services.
  More about it later.

Finally, exorun tells us that the application is now fully started
and ready to be used.
Open a browser and navigate to [http://localhost:3000](http://localhost:3000).
We got a running microservice-based web site!

Takeaway:
> The web server in a microservice application is much simpler than in a monolith,
> because it only focuses on interacting with the user via HTML.

Next, let's look at how services communicate with each other!

<table>
  <tr>
    <td><a href="05_communication.md"><b>&gt;&gt;</b></a></td>
  </tr>
</table>
