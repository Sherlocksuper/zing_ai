import {defaultSetting} from "../constants";

export interface Setting {
    model?: string,
    use_search?: boolean,
    stream?: boolean
}

const initSetting = () => {
    if (!getSetting()) {
        setSetting(defaultSetting);
    }
}

const setSetting = (data: Setting) => {
    localStorage.setItem('setting', JSON.stringify(data));
}

const getSetting = (): Setting => {
    const setting = localStorage.getItem('setting');
    return setting ? JSON.parse(setting) : null;
}

//更新设置
const updateSetting = (data: Setting) => {
    const setting = getSetting();
    if (setting) {
        setSetting({
            ...setting,
            ...data
        })
    }
}

//更新模型
const updateModel = (model: string) => {
    updateSetting({model})
}

//更新搜索
const updateSearch = (use_search: boolean) => {
    updateSetting({use_search})
}

//更新流式
const updateStream = (stream: boolean) => {
    updateSetting({stream})
}


const settingUtils = {
    initSetting,
    setSetting,
    getSetting,
    updateSetting,
    updateModel,
    updateSearch,
    updateStream
}

export default settingUtils;

