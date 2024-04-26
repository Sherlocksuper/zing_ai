//获取version列表
import {request} from "../../utils/request";

export const getVersionListApi = () => {
    return request({
        url: '/version/all',
        method: 'GET',
    })
};

export interface AddVersionParams {
    id?: number;
    version: string;
    introduction: string;
    enable: boolean;
    downloadUrl: string;
}

//添加version
export const addVersionApi = (params: AddVersionParams) => {
    return request({
        url: '/version/add',
        method: 'POST',
        params: params
    })
};

//update version
export const updateVersionApi = (params: AddVersionParams) => {
    return request({
        url: '/version/update',
        method: 'POST',
        params: params
    })
};