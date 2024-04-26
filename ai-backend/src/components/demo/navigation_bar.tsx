import React, {useState} from 'react';
import {
    ContainerOutlined,
    DesktopOutlined,
    PieChartOutlined,
} from '@ant-design/icons';
import type {MenuProps} from 'antd';
import {Menu} from 'antd';
import {useNavigate} from "react-router-dom";

type MenuItem = Required<MenuProps>['items'][number];

function getItem(
    label: React.ReactNode,
    key: React.Key,
    icon?: React.ReactNode,
    children?: MenuItem[],
    type?: 'group',
): MenuItem {
    return {
        key,
        icon,
        children,
        label,
        type,
    };
}

const items: MenuItem[] = [
    //数据库


    //Version管理
    getItem('Version管理', '0', <PieChartOutlined/>, [
        getItem('Version列表', 'versions', <ContainerOutlined/>),
    ]),
    //用户管理
    getItem('用户管理', '1', <PieChartOutlined/>, [
        getItem('用户列表', 'users', <ContainerOutlined/>),
    ]),
    getItem('数据库', '3', <PieChartOutlined/>, [
        getItem('SQL', 'mysql', <ContainerOutlined/>),
    ]),

];

const App: React.FC = () => {
    const [collapsed, setCollapsed] = useState(false);
    const navigate = useNavigate();

    return (
        <div style={{width: 256, height: "100vh"}}>
            <Menu
                defaultSelectedKeys={['1']}
                defaultOpenKeys={['sub1']}
                mode="inline"
                theme="dark"
                inlineCollapsed={collapsed}
                items={items}
                style={{height: '100%'}}
                onClick={(e) => {
                    navigate(`/${e.key}`)
                }}
            />
        </div>
    );
};

export default App;