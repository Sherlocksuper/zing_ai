export interface Message {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
    chatId: number;
    role: string;
    content: string;
    key: number;
}