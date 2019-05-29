# Trello as a microservice
An OMG service for trello, it is a task management app that gives you a visual overview of what is being worked on and who is working on it.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-trello.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-trello)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-trello/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-trello)


## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Create Board
```sh
$ omg run create_board -a board_name=<BOARD_NAME> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Create list on board
```sh
omg run create_list -a board_id=<BOARD_ID> -a list_name=<LIST_NAME> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Get List
```sh
omg run get_lists -a board_id=<BOARD_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Add card to list
```sh
$ omg run add_card -a name=<CARD_NAME> -a description=<DESCRIPTION> -a list_id=<LIST_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Get cards
```sh
$ omg run get_cards -a board_id=<BOARD_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Move card to list
```sh
$ omg run move_card -a card_id=<CARD_ID> -a list_id=<LIST_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Copy card to list
```sh
$ omg run copy_card -a card_id=<CARD_ID> -a list_id=<LIST_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Subscribe card
```sh
$ omg subscribe receive card -a board_id=<BOARD_ID> -a list_id=<LIST_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Delete Board
```sh
$ omg run delete_board -a board_id=<BOARD_ID> -e API_KEY=<API_KEY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```

## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-trello .
```
### RUN
```
docker run -p 3000:3000 microservice-trello
```
