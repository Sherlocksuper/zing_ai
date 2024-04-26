import "./code_block.css"
import {ReactNode} from "react";

interface CodeBlockProps {
    children: ReactNode;
}


export const CodeBlock = ({children}: CodeBlockProps) => {
    return (
        <>
            <div className={'container'}>
                <button className={'copy'} onClick={
                    () => {
                        // @ts-ignore
                        window.navigator.clipboard.writeText(children['props']['children'])
                    }
                }>复制
                </button>
                <pre className={'code'}>
                      {children}
                </pre>
            </div>
        </>
    );

}