import {Message} from "./format.ts";


interface Choice {
    index: number,
    message?: Message,
    delta?: { content: string },
    finish_reason: string
}

interface Usage {
    prompt_tokens: number,
    completion_tokens: number,
    total_tokens: number
}

export interface AIStreamResponseProps {
    id: string,
    model: string,
    object: string,
    choices: Array<Choice>,
    usage: Usage,
    created: number,
    stream: boolean
}


export const defaultToken = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyLWNlbnRlciIsImV4cCI6MTcyMDc5Mjg4NCwiaWF0IjoxNzEzMDE2ODg0LCJqdGkiOiJjb2Q5MGQxa3FxNG44aWtiZDI1ZyIsInR5cCI6InJlZnJlc2giLCJzdWIiOiJjbzRpbGw2Y3A3ZjJocDE4azgwZyIsInNwYWNlX2lkIjoiY280aWxsNmNwN2YyaHAxOGs4MDAiLCJhYnN0cmFjdF91c2VyX2lkIjoiY280aWxsNmNwN2YyaHAxOGs3dmcifQ.NIpyWNfvcMmkIwB_briD-PCMTaUuKOaHOHuFtcPFyFPu85-o-SpCKLmkeB8HWZx5BzCyASoC9o1OrYiDHa6GnQ"


export const defaultAIStreamResponse: AIStreamResponseProps = {
    id: "",
    model: "",
    object: "",
    choices: [
        {
            index: 0,
            finish_reason: "",
            delta: {
                content: ""
            }
        }
    ],
    usage: {
        prompt_tokens: 0,
        completion_tokens: 0,
        total_tokens: 0
    },
    created: 0,
    stream: true
}
