///TODO 这里是发送消息的格式


//{
//     // 模型名称随意填写，如果不希望输出检索过程模型名称请包含silent_search
//     "model": "kimi",
//     "messages": [
//         {
//             "role": "user",
//             "content": [
//                 {
//                     "type": "file",
//                     "file_url": {
//                         "url": "https://mj101-1317487292.cos.ap-shanghai.myqcloud.com/ai/test.pdf"
//                     }
//                 },
//                 {
//                     "type": "text",
//                     "text": "文档里说了什么？"
//                 }
//             }
//         },

//     ],
//     // 建议关闭联网搜索，防止干扰解读结果
//     "use_search": false
// }

const defaultModel = "kimi";
const defaultStream = true;


export enum Role {
    user = "user",
    ai = "assistant",
    system = "system",
}

export enum MessageType {
    text = "text",
    file = "file",
    image = "image",
    web = "web"
}

export interface Message {
    role: Role,
    content: {
        type: string,
        file_url?: {
            url: string
        },
        image_url?: {
            url: string
        },
        web_url?: {
            url: string
        },
        text?: string
    }[]
}

export interface FinalMessageFormat {
    model: string,
    messages: Message[],
    use_search: boolean,
    stream: boolean
}

// 这里的message只有role和text，
// ref是一个数组，里面有type和url,根据type把url的内容放到image_url或者file_url
export interface OriginMessageFormat {
    message: {
        role: Role,
        text: string
    },
    ref: {
        type: string,
        url: string
    }[],
    use_search: boolean,
}

// 这里只需要format user的消息即可
export const getRequestFormat =
    ({text, ref, use_search}: { text: string, ref: { type: string, url: string }[], use_search: boolean }
    ): OriginMessageFormat => {
        return {
            message: {
                role: Role.user,
                text: text
            },
            ref: ref,
            use_search: use_search
        }
    }


export const getFinalFormByOrigin = (origin: OriginMessageFormat): FinalMessageFormat => {
    const messages: Message[] = [
        {
            role: origin.message.role,
            content: [
                {
                    type: MessageType.text,
                    text: origin.message.text
                }
            ]
        }
    ]
    origin.ref.forEach((item) => {
        const message = messages[0];
        if (item.type === MessageType.file) {
            message.content.push({
                type: MessageType.file,
                file_url: {
                    url: item.url
                }
            })
        } else if (item.type === MessageType.image) {
            message.content.push({
                type: MessageType.image,
                image_url: {
                    url: item.url
                }
            })
        }
    })
    return {
        model: defaultModel,
        messages: messages,
        use_search: origin.use_search,
        stream: defaultStream
    }
}

export const getFinalFormByList = (list: OriginMessageFormat[]): FinalMessageFormat => {
    const messages: Message[] = []
    list.forEach((origin) => {
        const message: Message = {
            role: origin.message.role,
            content: [
                {
                    type: MessageType.text,
                    text: origin.message.text
                }
            ]
        }
        origin.ref.forEach((item) => {
            if (item.type === MessageType.file) {
                message.content.push({
                    type: MessageType.file,
                    file_url: {
                        url: item.url
                    }
                })
            } else if (item.type === MessageType.image) {
                message.content.push({
                    type: MessageType.image,
                    image_url: {
                        url: item.url
                    }
                })
            }
        })
        messages.push(message)
    })
    return {
        model: defaultModel,
        messages: messages,
        use_search: list[list.length - 1].use_search,
        stream: defaultStream
    }
}