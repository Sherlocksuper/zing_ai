import './home.css';
import {ChatContext, useChatManager} from "../../chat_context.ts";
import ChatPage from "../chat/Chat.tsx";
import {ChatListComponent} from "../chat/chat_list/ChatList.tsx";

export const Home = () => {
    const chatContext = useChatManager()

    const handleNewChat = () => {
        if (window.confirm("确定要创建新会话吗?"))
            chatContext.onAddChat()
    }

    const handleHistoryChat = () => {
        chatContext.openHistory()
    }

    return (
        <ChatContext.Provider value={chatContext}>
            <div id={'home'}>
                <div id={'home-content'}>
                    <div>
                        <button onClick={handleNewChat} className={"action-button"}>创建新会话</button>
                        <button onClick={handleHistoryChat} className={"action-button"}>历史会话</button>
                    </div>

                    <div id={'user-container'}>
                        <img className={'user-avatar'} src={'src//assets/user.jpg'} alt={'user-avatar'}/>
                    </div>
                </div>
                <ChatPage chat={chatContext.currentChat}/>
                <ChatListComponent />
            </div>
        </ChatContext.Provider>
    );
}
