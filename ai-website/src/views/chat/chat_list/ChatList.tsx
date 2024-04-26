import chatLocalUtils from "../../../utils/chatUtils.ts";
import {useContext} from "react";
import {ChatContext} from "../../../chat_context.ts";
import "./ChatList.css"


export const ChatListComponent = () => {
    const chatList = chatLocalUtils.getChatList()

    const chatContext = useContext(ChatContext)

    return (
        <div className={'chat-list-container'} style={{display: chatContext.openHis ? 'block' : 'none'}}>
            <div className={'container_bar'}>
                <h3 className={'chat-list-title'}>会话列表</h3>
                <button className={'close-history'} onClick={chatContext.closeHistory}>关闭</button>
            </div>

            <div className={'chat-list'}>
                {chatList.chatsStartTimes.map((chatListItem) => (
                    <ChatItem key={chatListItem.chatStartAt} chatStartAt={chatListItem.chatStartAt}
                              chatName={chatListItem.chatName}/>
                ))}
            </div>
        </div>
    );
}

const ChatItem = ({chatStartAt, chatName}: { chatStartAt: number, chatName: string }) => {
    const chatContext = useContext(ChatContext)

    return (
        <div className={'chat-item'} onClick={() => {
            chatContext.onSwitchChat(String(chatStartAt))
        }}>
            <span className={'chat-name'}>{chatName}</span>
            <span className={'chat-time'}>{new Date(Number(chatStartAt)).toLocaleString()}</span>
        </div>
    );
}
