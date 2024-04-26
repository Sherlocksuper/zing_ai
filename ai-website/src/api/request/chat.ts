import {AIStreamResponseProps, defaultAIStreamResponse} from "./format_save.ts";


//AutoComplete
//对话补全
// 需要参数：1.历史messages
// 2.新的messages
// 历史messages可以存到localStorage中
// 返回值：AIStreamResponseProps

// 把返回的数据转换成AIStreamResponseProps
export const updateToAIStreamResponse = (value: Uint8Array): AIStreamResponseProps => {
    const decoder = new TextDecoder();
    const text: string = decoder.decode(value);
    try {
        return JSON.parse(text.slice(text.indexOf("{"), text.lastIndexOf("}") + 1));
    } catch (e) {
        const textArr = text.split("\n").filter((item) => item !== "").map((item) => item.replace("data:", ""));
        const result = JSON.parse(JSON.stringify(defaultAIStreamResponse));
        textArr.forEach((item) => {
            if (isValidJSON(item)) {
                result.choices[0].delta!.content += JSON.parse(item).choices[0].delta.content;
            }
        });
        return result;
    }
}

function isValidJSON(text: string) {
    try {
        JSON.parse(text);
        return true;
    } catch {
        return false;
    }
}