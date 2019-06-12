# _Trello_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-trello.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-trello)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-trello/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-trello)

Trello microservice allows to create board, list, cards and also to subscribe the entire board or list, it has a variety of work and personal uses and overview of what is being worked on and what is the current status of card.

## Usage in [Storyscript](https://storyscript.io/)

### Create Board
```coffee
>>> trello createBoard boardName:'boardName'

```
### Get Board
```coffee
>>> trello getBoard boardId:'boardId'

```
### Create list on board
```coffee
>>> trello createList boardId:'boardId' listName:'listName'

```
### Get List
```coffee
>>> trello getLists boardId:'boardId'

```
### Add card to list
```coffee
>>> trello addCard name:'name' description:'description' listId:'listId'

```
### Get cards
```coffee
>>> trello getCards boardId:'boardId'

```
### Move card to list
```coffee
>>> trello moveCard cardId:'cardId' listId:'listId'

```
### Copy card to list
```coffee
>>> trello copyCard cardId:'cardId' listId:'listId'

```
### Delete Board
```coffee
>>> trello deleteBoard boardId:'boardId'

```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)
##### Create Board
```shell
$ omg run createBoard -a boardName=<BOARD_NAME> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Get Board
```shell
$ omg run getBoard -a boardId=<BOARD_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Create list on board
```shell
$ omg run createList -a boardId=<BOARD_ID> -a listName=<LIST_NAME> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Get List
```shell
$ omg run getLists -a boardId=<BOARD_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Add card to list
```shell
$ omg run addCard -a name=<CARD_NAME> -a description=<DESCRIPTION> -a listId=<LIST_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Get cards
```shell
$ omg run getCards -a boardId=<BOARD_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Move card to list
```shell
$ omg run moveCard -a cardId=<CARD_ID> -a listId=<LIST_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Copy card to list
```shell
$ omg run copyCard -a cardId=<CARD_ID> -a listId=<LIST_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Delete Board
```shell
$ omg run deleteBoard -a boardId=<BOARD_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Subscribe card
```shell
$ omg subscribe receive card -a boardId=<BOARD_ID> -a listId=<LIST_ID> -a existing=<BOOLEAN> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://choosealicense.com/licenses/mit/).

