package graphiql

import "net/http"

//Handler - Define the handler type
type Handler struct {
	endpoint string
}

//ServeHTTP required to match with the http.Handler
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) { w.Write([]byte(`
<html>

<head>
    <style>
        body {
            height: 100%;
            margin: 0;
            width: 100%;
            overflow: hidden;
        }
        
        #graphiql {
            height: 100vh;
        }
    </style>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.7.8/graphiql.css" />
    <script src="//cdn.jsdelivr.net/es6-promise/4.0.5/es6-promise.auto.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.0.0/fetch.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.4.2/react.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.4.2/react-dom.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.7.8/graphiql.js"></script>
</head>

<body>
    <div id="graphiql"></div>
    <script>
        !function(e) {
            function n(r) {
                if (t[r]) return t[r].exports;
                var o = t[r] = {
                    exports: {},
                    id: r,
                    loaded: !1
                };
                return e[r].call(o.exports, o, o.exports, n), o.loaded = !0, o.exports
            }
            var t = {};
            return n.m = e, n.c = t, n.p = "/", n(0)
        }
	([function(e, n, t) {
            e.exports = t(1)
        }, function(e, n) {
            "use strict";
            function t(e) {
                u.query = e, a()
            }
            function r(e) {
                u.variables = e, a()
            }
            function o(e) {
                u.operationName = e, a()
            }
            function a() {
                var e = "?" + Object.keys(u).filter(function(e) {
                    return Boolean(u[e])
                }).map(function(e) {
                    return encodeURIComponent(e) + "=" + encodeURIComponent(u[e])
                }).join("&");
                history.replaceState(null, null, e)
            }
            function i(e) {
                return fetch('` + h.endpoint + `', {
                    method: "post",
                    headers: {
                        Accept: "application/json",
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify(e),
                    credentials: "include"
                }).then(function(e) {
                    return e.text()
                }).then(function(e) {
                    try {
                        return JSON.parse(e)
                    } catch (n) {
                        return e
                    }
                })
            }
            var c = window.location.search, u = {};
            if (c.substr(1).split("&").forEach(function(e) {
                    var n = e.indexOf("=");
                    n >= 0 && (u[decodeURIComponent(e.slice(0, n))] = decodeURIComponent(e.slice(n + 1)))
                }), u.variables) try {
                u.variables = JSON.stringify(JSON.parse(u.variables), null, 2)
            } catch (e) {}
            ReactDOM.render(React.createElement(GraphiQL, {
                fetcher: i,
                query: u.query,
                defaultQuery: "",
                variables: u.variables,
                operationName: u.operationName,
                onEditQuery: t,
                onEditVariables: r,
                onEditOperationName: o
            }), document.getElementById("graphiql"))
        }]);
    </script>
</body>

</html>

`))}

//New graphiql tool
func New(endpoint string) *Handler {
	return &Handler{
		endpoint: endpoint,
	}
}
