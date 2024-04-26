// 导入antd组件和CSS样式
import {Button, Table, Space, Input, message, TableProps} from 'antd';
import React, {useEffect, useState} from 'react';

// 定义接口DatabaseBackup，用于备份数据的结构
interface DatabaseBackup {
    id: number;
    name: string;
    size: number; // 以KB为单位
    date: string; // 备份日期
}

const onDownload = (id: number) => {
    console.log(id)
}

const onDelete = (id: number) => {
    console.log(id)
}

// 定义表格的列
const columns: TableProps<DatabaseBackup>['columns'] = [
    {title: 'ID', dataIndex: 'id'},
    {title: 'Name', dataIndex: 'name'},
    {title: 'Size (KB)', dataIndex: 'size'},
    {title: 'Date', dataIndex: 'date'},
    {
        title: 'manage', key: 'action', render: (text, record) => (
            <Space>
                <Button type="primary" onClick={() => {
                    onDownload(record.id)
                }}>Download</Button>
                <Button type="primary" onClick={() => {
                    onDelete(record.id)
                }}>Delete</Button>
            </Space>
        )
    },
];

// 定义组件的props和state
const DatabaseManager: () => void = () => {
        const [backups, setBackups] = useState<DatabaseBackup[]>([]);

        // 模拟获取数据库备份列表的函数
        const fetchBackups = async () => {
            const mockBackups: DatabaseBackup[] = [
                {id: 1, name: 'Backup 1', size: 1024, date: '2024-04-01'},
                {id: 2, name: 'Backup 2', size: 2048, date: '2024-04-02'},
                // 更多备份...
            ];
            setBackups(mockBackups);
        };

        // 备份数据库的函数
        const backupDatabase = async () => {
            try {
                // 这里应该是一个API调用，备份数据库
                // 以下是模拟数据
                const newBackup: DatabaseBackup = {
                    id: backups.length + 1,
                    name: `Backup ${backups.length + 1}`,
                    size: Math.floor(Math.random() * 1024) + 1024,
                    date: new Date().toISOString(),
                };
                setBackups([...backups, newBackup]);
                message.success('Database backup successful');
            } catch (error) {
                message.error('Database backup failed');
            }
        }

        useEffect(() => {
            fetchBackups()
        }, []);

        return (
            <>
                <h1>Database Manager</h1>
                <Space>
                    <Button onClick={backupDatabase}>Backup Database</Button>
                    <Input.Search placeholder="Search backups..."/>
                </Space>
                <Table dataSource={backups} columns={columns} rowKey="id"/>
            </>
        );
    }
export default DatabaseManager;