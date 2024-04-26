import './App.css';
import NavigationBar from "../../components/demo/navigation_bar";
import {Outlet} from "react-router-dom";


function App() {
    return (
        <div style={{
            display: 'flex',
            height: '100vh',
            width: '100vw',
        }}>
            <NavigationBar/>
            <div style={{
                display: 'flex',
                flexDirection: 'column',
                width: '100%',
                height: '100vh',
                overflow: 'auto'
            }}>
                <Outlet/>
            </div>
        </div>
    );
}


export default App;
