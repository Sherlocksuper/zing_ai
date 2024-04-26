import React, {useState} from 'react';
import navigation_bar from "../../components/demo/navigation_bar";
import {useNavigate} from "react-router-dom";

interface Credentials {
    username: string;
    password: string;
}

const Login: React.FC = () => {
    const navigate = useNavigate();
    const [credentials, setCredentials] = useState<Credentials>({
        username: '',
        password: '',
    });

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const {name, value} = e.target;
        setCredentials({...credentials, [name]: value});
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (credentials.username === '' || credentials.password === '') {
            alert('请输入账户名和密码');
            return;
        }
        if (credentials.username !== 'Sherlock' || credentials.password !== '123456') {
            alert('账户名或密码错误');
            return;
        }
        navigate('/');
        console.log('登录信息:', credentials);
    };

    return (
        <div style={{display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh'}}>
            <div style={{
                width: '300px',
                textAlign: 'center',
                border: '1px solid #ccc',
                boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
                borderRadius: '8px',
                padding: '20px'
            }}>
                <h2 style={{marginBottom: '20px'}}>用户登录</h2>
                <form onSubmit={handleSubmit}>
                    <div
                        style={{marginBottom: '20px', display: 'flex', flexDirection: 'column', alignItems: 'stretch'}}>
                        <label style={{marginBottom: '5px'}}>账户名:</label>
                        <input
                            type="text"
                            name="username"
                            value={credentials.username}
                            onChange={handleInputChange}
                            style={{padding: '8px', borderRadius: '5px', border: '1px solid #ccc'}}
                        />
                    </div>
                    <div
                        style={{marginBottom: '20px', display: 'flex', flexDirection: 'column', alignItems: 'stretch'}}>
                        <label style={{marginBottom: '5px'}}>密码:</label>
                        <input
                            type="password"
                            name="password"
                            value={credentials.password}
                            onChange={handleInputChange}
                            style={{padding: '8px', borderRadius: '5px', border: '1px solid #ccc'}}
                        />
                    </div>
                    <button type="submit" style={{
                        padding: '8px 16px',
                        borderRadius: '5px',
                        background: 'blue',
                        color: 'white',
                        border: 'none'
                    }}>登录
                    </button>
                </form>
            </div>
        </div>
    );
};

export default Login;
