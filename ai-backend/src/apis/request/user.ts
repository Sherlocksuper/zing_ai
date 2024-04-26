import {request} from '../../utils/request'
import {GET_USER_DETAIL, GET_USER_LIST} from "../url/user";
import {AccountStatus} from "../struct/user";


//获取用户列表
export const getUserListApi = (params: any) => {
    return request({
        url: GET_USER_LIST,
        method: 'GET',
        params: params
    })
}

//获取用户详情
export const getUserDetailApi = (params: any) => {
    return request({
        url: GET_USER_DETAIL,
        method: 'get',
        params
    })
}


interface UpdateUserParams {
    id: number;
    name?: string;
    email?: string;
    password?: string;
    accountStatus?: AccountStatus.NORMAL | AccountStatus.BAN;
}

//update user
export const updateUser = (params: UpdateUserParams) => {
    return request({
        url: '/user/update',
        method: 'post',
        params: params
    })
}


//封禁用户
export const banUser = (userId: number) => {
    const params: UpdateUserParams = {
        id: userId,
        accountStatus: AccountStatus.BAN
    }
    return updateUser(params)
}

//解封用户
export const unbanUser = (userId: number) => {
    const params: UpdateUserParams = {
        id: userId,
        accountStatus: AccountStatus.NORMAL
    }
    return updateUser(params)
}




