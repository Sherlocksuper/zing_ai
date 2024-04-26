import React from 'react';
import './ChatBubble.css';
import Markdown from "react-markdown";
import {CodeBlock} from "../../../components/code_block/code_block.tsx";
import {OriginMessageFormat, Role} from "../../../api/request/format.ts"; // 引入CSS样式文件

interface ChatBubbleProps {
    orignMessage: OriginMessageFormat;
}

const ChatBubble: React.FC<ChatBubbleProps> = ({orignMessage}) => {

    const isFromUser = orignMessage.message.role === Role.user
    return (
        <div className={`chat-line  ${isFromUser ? 'user-line' : 'ai-line'}`}>
            <div className={`chat-bubble`}>
                <WebUrlBox referance={orignMessage.ref}/>
                <Markdown
                    components={{
                        // eslint-disable-next-line @typescript-eslint/no-unused-vars
                        pre: ({children}) => {
                            return <CodeBlock children={children}/>;
                        },
                    }}
                >
                    {orignMessage.message.text ?? "hal"}
                </Markdown>

            </div>
            <img
                className="chat-avatar"
                src={isFromUser ? '/src/assets/user.jpg' : '/src/assets/ai.jpg'}
                alt="avatar"/>
        </div>
    );
};

const WebUrlBox: React.FC<{ referance: { type: string, url: string }[] }> = ({referance}) => {

    const [shrinkBox, setShrinkBox] = React.useState(false);

    const changeShrinkBox = () => {
        setShrinkBox(!shrinkBox);
    }

    if (!referance || !referance.length) return null;
    return <div className={'web-url-box'} style={{height: "auto"}}>
        {
            <div onClick={changeShrinkBox} style={{
                cursor: "pointer",
                color: "blue",
                textDecoration: "underline"
            }}>{shrinkBox ? "展开" : "收起"}</div>
        }

        {!shrinkBox && referance.map((ref) => {
            return <div className={"url-item"}>
                {ref.type === "web" && <a href={
                    ref.url.slice(ref.url.indexOf("(") + 1, ref.url.indexOf(")"))
                } target="_blank">{
                    ref.url.slice(3, ref.url.indexOf("("))
                }</a>}
            </div>
        })}
    </div>
}


export default ChatBubble;