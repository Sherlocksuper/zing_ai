import React from "react";
import chatLocalUtils from "./utils/chatUtils.ts";
import {ChatFormat} from "./api/standard/chatFormat.ts";

//创建一个context
export const ChatContext = React.createContext<ReturnType<typeof useChatManager>>({} as any);

export const useChatManager = () => {
    const defaultChat: ChatFormat = {
        chatName: "默认标题",
        chatStartAt: 0,
        messages: []
    }

    const [openHis, setOpenHis] = React.useState<boolean>(false)
    const [currentChat, setCurrentChat] = React.useState<ChatFormat>({...defaultChat});

    //当添加chat时
    const onAddChat = () => {
        const newChat: ChatFormat = {
            chatName: "新会话",
            chatStartAt: Date.now(),
            messages: []
        }
        setCurrentChat(newChat);
        chatLocalUtils.addChat(newChat)
    }

    //删除chat,不能删除当前的chat
    const onDeleteChat = (chatStartAt: string) => {
        if (chatStartAt === currentChat.chatStartAt.toString()) {
            alert("不能删除当前的chat")
            return
        }
        chatLocalUtils.removeChat(Number(chatStartAt))
    }

    //更换chat
    const onSwitchChat = (chatStartAt: string) => {
        const chat = chatLocalUtils.getChat(Number(chatStartAt))
        console.log("切换到chat" + chatStartAt + ",现在具有messages:" + chat!.messages.length)
        setCurrentChat(chat!)
        closeHistory()
    }

    //发送消息以后更新当前chat
    const onUpdateChat = (chat: ChatFormat) => {
        console.log("更新了" + chat.chatStartAt, ",现在具有messages:" + chat.messages.length)
        chatLocalUtils.updateChat(chat)
    }

    const openHistory = () => {
        setOpenHis(true)
    }

    const closeHistory = () => {
        setOpenHis(false)
    }


    return {
        openHis,
        currentChat,
        onAddChat,
        onDeleteChat,
        onSwitchChat,
        onUpdateChat,
        openHistory,
        closeHistory,
    }
}