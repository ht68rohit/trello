# _Trello_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-trello.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-trello)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-trello/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-trello)

Trello microservice allows to create board, list, cards and also to subscribe the entire board or list, it has a variety of work and personal uses and overview of what is being worked on and what is the current status of card.

## Usage in [Storyscript](https://storyscript.io/)

```coffee
_insert example usage here_
```

Curious to [learn more](https://docs.storyscript.io/)?

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)
##### Create Board
```shell
_$ omg run createBoard -a boardName=<BOARD_NAME> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>_
```
##### Get Board
```shell
_$ omg run getBoard -a boardId=<BOARD_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>_
```
##### Create list on board
```shell
_$ omg run createList -a boardId=<BOARD_ID> -a listName=<LIST_NAME> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>_
```
##### Get List
```shell
_$ omg run getLists -a boardId=<BOARD_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>_
```
##### Add card to list
```shell
_$ omg run addCard -a name=<CARD_NAME> -a description=<DESCRIPTION> -a listId=<LIST_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>_
```
##### Get cards
```shell
_$ omg run getCards -a boardId=<BOARD_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>_
```
##### Move card to list
```shell
_$ omg run moveCard -a cardId=<CARD_ID> -a listId=<LIST_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>_
```
##### Copy card to list
```shell
_$ omg run copyCard -a cardId=<CARD_ID> -a listId=<LIST_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>_
```
##### Delete Board
```shell
_$ omg run deleteBoard -a boardId=<BOARD_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>_
```
##### Subscribe card
```shell
_$ omg subscribe receive card -a boardId=<BOARD_ID> -a listId=<LIST_ID> -a existing=<BOOLEAN> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>_
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://choosealicense.com/licenses/mit/).

