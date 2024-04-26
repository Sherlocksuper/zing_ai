import {FinalMessageFormat} from "./format.ts";
import {fetchData} from "../base_request.ts";

export const autoComplete = (data: FinalMessageFormat) => {
    return fetchData({
        data: data,
    })
}