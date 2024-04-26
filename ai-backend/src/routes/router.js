import {createBrowserRouter} from "react-router-dom";
import App from "../views/demo/App";
import ErrorPage from "./err-page";
import Versions from "../views/version/info";
import Users from "../views/user/list/list";
import Login from "../views/login/login";
import DatabaseManager from "../views/data/sql";

export const router = createBrowserRouter([
    {
        path: "/",
        element: <App/>,
        errorElement: <ErrorPage/>,
        children: [
            {
                index: true,
                path: "/versions",
                element: <Versions/>,
            },
            {
                path: "/users",
                element: <Users/>,
            },
            {
                path: "/mysql",
                element: <DatabaseManager/>
            }
        ],
    },
    {
        path: "/login",
        element: <Login/>,
    },
    {
        path: "*",
        element: <ErrorPage/>,
    }
]);