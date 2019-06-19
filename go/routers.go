/*
 * 挣闲钱
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"GetDelivery",
		strings.ToUpper("Get"),
		"/task/delivery/{taskId}",
		GetDelivery,
	},

	Route{
		"GetQuestionare",
		strings.ToUpper("Get"),
		"/task/questionare/{taskId}",
		GetQuestionare,
	},

	Route{
		"QAcceptPage",
		strings.ToUpper("Get"),
		"/task/qAcceptPage",
		QAcceptPage,
	},

	Route{
		"QPublishPage",
		strings.ToUpper("Get"),
		"/task/qPublishPage",
		QPublishPage,
	},

	Route{
		"QueryPageD",
		strings.ToUpper("Get"),
		"/task/queryPageD",
		QueryPageD,
	},

	Route{
		"QueryPageQ",
		strings.ToUpper("Get"),
		"/task/queryPageQ",
		QueryPageQ,
	},

	Route{
		"QueryTitle",
		strings.ToUpper("Get"),
		"/task/queryTitle",
		QueryTitle,
	},

	Route{
		"AcceptTask",
		strings.ToUpper("Post"),
		"/user/accept",
		AcceptTask,
	},

	Route{
		"FinishAccept",
		strings.ToUpper("Put"),
		"/user/finishAccept",
		FinishAccept,
	},

	Route{
		"FinishPublish",
		strings.ToUpper("Put"),
		"/user/finishPublish",
		FinishPublish,
	},

	Route{
		"GetUser",
		strings.ToUpper("Get"),
		"/user/{userId}",
		GetUser,
	},

	Route{
		"PublishDTask",
		strings.ToUpper("Post"),
		"/user/publishDelivery",
		PublishDTask,
	},

	Route{
		"PublishQTask",
		strings.ToUpper("Post"),
		"/user/publishQuery",
		PublishQTask,
	},

	Route{
		"SignIn",
		strings.ToUpper("Post"),
		"/user/signin",
		SignIn,
	},

	Route{
		"SignUp",
		strings.ToUpper("Post"),
		"/user/signup",
		SignUp,
	},
	Route{
		"SignOut",
		strings.ToUpper("DELETE"),
		"/user/signout",
		SignOut,
	},
	Route{
		"FillQuery",
		strings.ToUpper("Post"),
		"/user/fillQuestionare",
		FillQuery,
	},
}
