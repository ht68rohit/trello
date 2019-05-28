package route

import (
    "github.com/gorilla/mux"
    trello "github.com/heaptracetechnology/microservice-trello/trello"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "GetCards",
        "POST",
        "/getcards",
        trello.GetCards,
    },
    Route{
        "AddCard",
        "POST",
        "/addcard",
        trello.AddCard,
    },
    Route{
        "MoveCard",
        "POST",
        "/movecard",
        trello.MoveCard,
    },
    Route{
        "SubscribeCard",
        "POST",
        "/subscribe",
        trello.SubscribeCard,
    },
    Route{
        "CopyCard",
        "POST",
        "/copycard",
        trello.CopyCard,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
