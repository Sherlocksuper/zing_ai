import {defaultToken, Role} from "../src/api/request/chat_base_setting";
import {autoCompleteChat} from "../src/api/request/chat";
import {checkToken} from "../src/api/request/token";


test('getAutoComplete', async () => {
    const message = {
        role: Role.user,
        content: "测试"
    }
    const data = await autoCompleteChat(message);
    expect(data).not.toBeNull();
});

test('checkToken', async () => {
    const data = await checkToken({token:defaultToken});
    console.log(JSON.stringify(data));
});