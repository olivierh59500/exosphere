Feature: help command

  As a developer not familiar with how to call "exo clean" correctly
  I want to be given helpful information how to use this application
  So that I can do the things I intended without having to look this up somewhere else.


  Scenario: the user enters 'exo help'
    When running "exo help" in the terminal
    Then I see:
      """
      Usage:
        exo [command]

      Available Commands:
        add         Adds a new service to the current application
        clean       Removes dangling Docker images and volumes
        configure   Configures secrets for an Exosphere application deployed to the cloud
        create      Creates a new Exosphere application
        deploy      Deploys Exosphere application to the cloud
        run         Runs an Exosphere application
        template    Manages remote service templates
        test        Runs tests for the application
        version     Displays the version

      Use "exo [command] --help" for more information about a command
      """


  Scenario: the user enters 'exo add help'
    When running "exo add help" in the terminal
    Then I see:
      """
      Adds a new service to the current application

      Usage:
        exo add
      """


  Scenario: the user enters 'exo help clean'
    When running "exo clean help" in the terminal
    Then I see:
      """
      Removes dangling Docker images and volumes

      Usage:
        exo clean
      """


  Scenario: the user enters 'exo create help'
    When running "exo create help" in the terminal
    Then I see:
      """
      Creates a new Exosphere application

      Usage:
        exo create
      """

  Scenario: the user enters 'exo help deploy'
    When running "exo help deploy" in the terminal
    Then I see:
      """
      Deploys Exosphere application to the cloud

      Usage:
        exo deploy
      """

  Scenario: the user enters 'exo template help'
    When running "exo template help" in the terminal
    Then I see:
      """
      Manages remote service templates

      Usage:
        exo template [command]

      Available Commands:
        add         Adds a remote service template to .exosphere
        fetch       Fetches updates for an existing service template in .exosphere
        remove      Removes an existing service template from .exosphere
        test        Tests service templates
      """


  Scenario: the user enters 'exo template add help'
    When running "exo template add help" in the terminal
    Then I see:
      """
      Adds a remote service template to .exosphere

      Usage:
        exo template add <name> <url> [<commit-ish>]
      """


  Scenario: the user enters 'exo template fetch help'
    When running "exo template fetch help" in the terminal
    Then I see:
      """
      Fetches updates for an existing service template in .exosphere

      Usage:
        exo template fetch <name>
      """


  Scenario: the user enters 'exo template remove help'
    When running "exo template remove help" in the terminal
    Then I see:
      """
      Removes an existing service template from .exosphere

      Usage:
        exo template remove <name>
      """

  Scenario: the user enters 'exo template test help'
    When running "exo template test help" in the terminal
    Then I see:
      """
      Tests service templates by adding to an exopshere application and running tests.
      This command must be run in the directory of an exosphere template.

      Usage:
        exo template test
      """

  Scenario: the user enters 'exo run help'
    When running "exo run help" in the terminal
    Then I see:
      """
      Runs an Exosphere application

      Usage:
        exo run
      """
