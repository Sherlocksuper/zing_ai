import {createBrowserRouter} from "react-router-dom";
// @ts-ignore
import {Home} from "../views/home/home.tsx";
// @ts-ignore
import ChatPage from "../views/chat/Chat.tsx";

export const routes = createBrowserRouter([
    {
        path: "/",
        element: <Home/>,

    },

    {
        path: "/about",
        element: <div>
            <h1>About</h1>
            <p>Welcome to the about page!</p>
        </div>
    }
])