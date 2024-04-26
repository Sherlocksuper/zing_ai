import React, {useEffect, useState} from 'react';
import {Button, Form, Input, Modal, Radio, Space, Table} from 'antd';
import type {TableProps} from 'antd';
import {Version} from "../../apis/struct/version";
import {addVersionApi, AddVersionParams, getVersionListApi, updateVersionApi} from "../../apis/request/version";


const getVersion = async () => {
    return await getVersionListApi()
}
const addVersion = async (params: AddVersionParams) => {
    return await addVersionApi(params)
}

const columns: TableProps<Version>['columns'] = [
    {
        title: 'ID',
        dataIndex: 'id',
    },
    {
        title: 'Version',
        dataIndex: 'version',
    },
    {
        title: 'Introduction',
        dataIndex: 'introduction',
    },
    {
        title: 'Created At',
        dataIndex: 'createdAt',
    },
    {
        title: 'Download Url',
        dataIndex: 'downloadUrl',
    },
    {
        title: 'Enable',
        dataIndex: 'enable',
        render: (text, record) => (
            <Radio checked={record.enable} disabled/>
        )
    },
    {
        title: 'Action',
        key: 'action',
    },
];

const App: React.FC = () => {

    const [versionData, setVersionData] = React.useState<Version[]>([] as Version[])
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [form] = Form.useForm();

    const changeEnable = (record: Version, enable: boolean) => {
        updateVersionApi({
            id: record.id,
            version: record.version,
            introduction: record.introduction,
            enable: enable,
            downloadUrl: record.downloadUrl
        }).then(r => {
            console.log(r)
        })
        //更新本地
        const newData = versionData.map(item => {
            if (item.id === record.id) {
                item.enable = enable;
            }
            return item;
        })
        setVersionData(newData)
    }

    useEffect(() => {
        getVersion().then(res => {
            setVersionData(res.data)
            res.data.forEach((item: Version) => {
                item.key = item.id;
                item.createdAt = new Date(item.createdAt).toLocaleString();
            });
        }).catch(err => {
            console.error(err);
        })
    }, []);

    columns[columns.length - 1].render = (text, record) => (
        <Space size="middle">
            <Button type="primary" onClick={() => changeEnable(record, false)}>Disable</Button>
            <Button type="primary" onClick={() => changeEnable(record, true)}>Enable</Button>
        </Space>
    )
    const showModal = () => {
        setIsModalOpen(true);
    };
    const hiddenModal = () => {
        setIsModalOpen(false);
    };
    const handleAdd = () => {
        addVersion(form.getFieldsValue() as AddVersionParams).then(res => {
            console.log(res)
        }).catch(err => {
            //弹窗
            console.error(err)
        })
        hiddenModal()
    }

    return (
        <div style={{
            display: "flex",
            flexDirection: "column",
            margin: "10px",
        }}>
            <Button type={"primary"} style={{width: '200px'}} onClick={showModal}>Add</Button>
            <Table columns={columns} dataSource={versionData} style={{marginTop: "20px"}}/>;
            <Modal title="Basic Modal" open={isModalOpen} onOk={handleAdd} onCancel={hiddenModal}>
                <Form form={form} layout="vertical">
                    <Form.Item label="Version" name={"version"}>
                        <Input/>
                    </Form.Item>
                    <Form.Item label="Introduction" name={"introduction"}>
                        <Input.TextArea/>
                    </Form.Item>
                    <Form.Item label="Enable" name={"enable"}>
                        <Radio.Group>
                            <Radio value={true}>True</Radio>
                            <Radio value={false}>False</Radio>
                        </Radio.Group>
                    </Form.Item>
                    <Form.Item label="Download Url" name={"downloadUrl"}>
                        <Input/>
                    </Form.Item>
                </Form>
            </Modal>
        </div>
    )
}


export default App;