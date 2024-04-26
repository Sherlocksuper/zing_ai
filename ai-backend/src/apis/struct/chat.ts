//type Chat struct {
// 	gorm.Model
// 	Title         string    `json:"title" gorm:"default:默认标题"`
// 	UserID        uint      `json:"userId"`
// 	SystemMessage string    `json:"systemMessage"`
// 	Messages      []Message `json:"messages"`
// 	Type          ChatType  `json:"type" gorm:"default:Text"`
// }

import {Message} from "./message";

export interface Chat {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
    title: string;
    userId: number;
    systemMessage: string;
    messages: Message[];
    type: string;
    key: number;
}