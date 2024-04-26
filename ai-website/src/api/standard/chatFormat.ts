import { OriginMessageFormat} from "../request/format.ts";

export interface ChatFormat {
    chatName?: string;
    chatStartAt: number;
    messages: OriginMessageFormat[];
}

export interface ChatListItem {
    chatName: string;
    chatStartAt: number;
}

