import type {TableProps} from 'antd';
import {Button, Space, Table} from 'antd';
import {banUser, getUserListApi, unbanUser} from "../../../apis/request/user";
import React, {useEffect} from "react";
import {AccountStatus, User} from "../../../apis/struct/user";
import Search from "antd/es/input/Search";

const columns: TableProps<User>['columns'] = [
    {
        title: 'ID',
        dataIndex: 'id',
    },
    {
        title: 'Name',
        dataIndex: 'name',
    },
    {
        title: 'Role',
        dataIndex: 'role',
    },
    {
        title: 'Email',
        dataIndex: 'email',
    },
    {
        title: 'Created At',
        dataIndex: 'createdAt',
    },
    {
        title: 'Last Login Time',
        dataIndex: 'lastLoginTime',
    },
    {
        title: 'Status',
        dataIndex: 'accountStatus',
    },
    {
        title: 'Total Chats',
        dataIndex: 'chatNum',
    },
    {
        title: 'Action',
        key: 'action',
    },
];
const getUserList = async () => {
    return await getUserListApi({});
}

const App: React.FC = () => {
    const [data, setData] = React.useState<User[]>([]);

    useEffect(() => {
        getUserList().then(res => {
            setData(res.data);
            console.log(res.data)
            res.data.forEach((item: User) => {
                item.key = item.id;

                // @ts-ignore
                item.createdAt = new Date(item.createdAt).toLocaleString();
            });
        }).catch(err => {
            console.error(err);
        })
    }, []);

    const handleBan = (id: number) => {
        banUser(id).then(() => {
            const newData = data.map(item => {
                if (item.id === id) item.accountStatus = AccountStatus.BAN
                return item;
            });
            setData(newData);
        }).catch(err => {
            console.error(err);
        })
    }

    const handleUnban = (id: number) => {
        console.log('unban')
        unbanUser(id).then(() => {
            console.log('unban success')
            const newData = data.map(item => {
                if (item.id === id) item.accountStatus = AccountStatus.NORMAL
                return item;
            });
            setData(newData);
        }).catch(err => {
            console.error(err);
        })
    }

    const handleClick = (record: User) => () => {
        if (record.role === 'Admin') {
            window.alert('Admin用户不允许封禁')
            return;
        }
        if (record.accountStatus === AccountStatus.BAN) {
            console.log('unban')
            handleUnban(record.id)
        } else if (record.accountStatus === AccountStatus.NORMAL) {
            console.log('ban')
            handleBan(record.id)
        }
    }

    const onSearch = (value: string) => {
        getUserListApi({name: value}).then(res => {
            setData(res.data);
            res.data.forEach((item: User) => {
                item.key = item.id;
            });
        }).catch(err => {
            console.error(err);
        })
    }

    columns[columns.length - 1].render = (_value, record, index) => (
        <Space>
            <Button type="primary" onClick={handleClick(record)}>
                {record.accountStatus === AccountStatus.BAN ? '解封' : '封禁'}
            </Button>
        </Space>
    );

    return (
        <>
            <Search placeholder="input search text" onSearch={onSearch} enterButton size={"large"}
                    style={{
                        width: "300px"
                    }}/>


            <Table columns={columns} dataSource={data} key={data.length} style={{
                marginTop: "20px"
            }} size={"middle"} pagination={{pageSize: 10}}/>
        </>
    )
}

export default App;