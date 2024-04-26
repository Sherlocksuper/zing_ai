import {CHAT_URL} from "./url.ts";
import {tokenUtils} from "../utils/tokenUtils.ts";

const headerInit: HeadersInit = {
    'Authorization': `Bearer ${tokenUtils.getToken()}`,
    'Content-Type': 'application/json',
}

interface RequestProps {
    data?: any
}

export const fetchData = ({data}: RequestProps) => {
    return fetch(CHAT_URL, {
        method: 'POST',
        headers: headerInit,
        body: JSON.stringify(data),
    })
}