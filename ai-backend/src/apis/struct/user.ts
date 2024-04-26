import {Chat} from "./chat";

export interface User {
    id: number;
    key?: number;
    name: string;
    role: 'Admin' | 'User';
    email?: string;
    createdAt?: string;
    lastLoginAt?: string;
    totalChats?: number;
    chats?: Chat[];
    chatNum?: number;
    accountStatus?: AccountStatus.BAN | AccountStatus.NORMAL;
}

export enum AccountStatus {
    NORMAL = 'normal',
    BAN = 'ban'
}
