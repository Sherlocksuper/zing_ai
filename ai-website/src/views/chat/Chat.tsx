import React, {useContext, useEffect, useState} from 'react';
import './ChatPage.css';
import {getFinalFormByList, MessageType, OriginMessageFormat, Role} from "../../api/request/format.ts";
import ChatBubble from "./chat_bubble/ChatBubble.tsx";
import {autoComplete} from "../../api/request/chats.ts";
import {updateToAIStreamResponse} from "../../api/request/chat.ts";
import {ChatFormat} from "../../api/standard/chatFormat.ts";
import {ChatContext} from "../../chat_context.ts";

// 定义ChatPage组件
//传入一个chat
const ChatPage = ({chat}: { chat: ChatFormat }) => {
    const [inputMessage, setInputMessage] = useState<string>('');
    const [messages, setMessages] = useState<OriginMessageFormat[]>([]);
    const [useSearch, setUseSearch] = useState<boolean>(true);
    const chatContext = useContext(ChatContext)

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setInputMessage(e.target.value);
    }

    useEffect(() => {


        setMessages(chat.messages)
        scrollToBottom();
    }, [chat.messages])

    const scrollToBottom = () => {
        const chatBox = document.querySelector('.messages');
        chatBox?.scrollTo(0, chatBox.scrollHeight);
    }

    const handleSendMessage = async () => {
        if (inputMessage === "" || inputMessage === '') return

        if (chat.messages.length === 0) {
            chat.chatName = inputMessage.slice(0, 5) + "...";
            chatContext.onUpdateChat(chat);
        }

        setInputMessage('');
        const currentMessages = [...messages]
        const newMessage: OriginMessageFormat = {
            message: {
                role: Role.user,
                text: inputMessage
            },
            ref: [],
            use_search: useSearch
        }
        currentMessages.push(newMessage);
        setMessages(currentMessages);
        const finalFormat = getFinalFormByList(currentMessages)

        const aiResponse: OriginMessageFormat = {
            message: {
                role: Role.ai,
                text: ""
            },
            ref: [],
            use_search: false,
        }

        await autoComplete(finalFormat).then(async (res) => {
            const reader = res.body!.getReader();
            while (true) {
                const {done, value} = await reader!.read();
                if (done) break;
                const resultJson = updateToAIStreamResponse(value!);
                try {
                    if (resultJson!.choices[0].delta!.content.length > 0) {
                        if (!resultJson!.choices[0].delta!.content.startsWith("检索")) {
                            aiResponse.message.text += resultJson!.choices[0].delta!.content;

                            setMessages([...currentMessages, aiResponse])
                        } else {
                            aiResponse.ref.push({
                                type: MessageType.web,
                                url: resultJson!.choices[0].delta!.content
                            });
                        }
                    }
                } catch (e) {
                    console.log(e)
                }
                scrollToBottom();
            }
        })

        //更新chat
        currentMessages.push(aiResponse)
        chat.messages = currentMessages;
        chatContext.onUpdateChat(chat);
    }

    return (
        <div className="chat-box">
            <div className="chat-title">
                {chat.chatName}
            </div>
            <div className="messages">
                {messages.map((messageItem: OriginMessageFormat, index) => (
                    <ChatBubble orignMessage={messageItem} key={index}/>
                ))}
            </div>
            <div className="input-area">
                <input
                    type="text"
                    value={inputMessage}
                    onChange={handleInputChange}
                    placeholder="输入消息..."
                    onKeyUp={(e) => {
                        if (e.key === 'Enter') {
                            handleSendMessage();
                        }
                    }}
                />
                <IconButton
                    handleClick={handleSendMessage}
                    condition={inputMessage !== "" && inputMessage !== ''}
                    trueImage="/src/assets/send.svg"
                    falseImage="/src/assets/send_false.svg"
                />

                <IconButton
                    handleClick={() => setUseSearch(!useSearch)}
                    condition={useSearch}
                    trueImage="/src/assets/search.svg"
                    falseImage="/src/assets/search_false.svg"
                />

            </div>
        </div>
    );
}
// @ts-ignore
const IconButton = ({handleClick, condition, trueImage, falseImage}) => (
    <button onClick={handleClick} className="input-area-button">
        {condition ? <img src={trueImage} alt="icon"/> : <img src={falseImage} alt="icon"/>}
    </button>
);
export default ChatPage;

