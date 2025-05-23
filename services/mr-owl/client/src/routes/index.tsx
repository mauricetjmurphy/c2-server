import { useRoutes } from 'react-router-dom'
import Dashboard from '../pages/dashboard/Dashboard'
import Layout from '../components/layout/MainLayout'
import Home from '../pages/home/Home'
import ScreenLayout from '../components/layout/ScreenLayout'
import User from '../pages/user/User'
import Log from '../pages/log/Log'
import Device from '../pages/device/Device'
import Listener from '../pages/listener/Listener'

const AppRoutes = () => {
    return useRoutes([
        {
            path: '/',
            element: <Layout />,
            children: [
                { index: true, element: <Home /> },
                {
                    path: '/screens',
                    element: <ScreenLayout />,
                    children: [
                        { path: '/screens/dashboard', element: <Dashboard /> },
                        { path: '/screens/devices', element: <Device /> },
                        { path: '/screens/listeners', element: <Listener /> },
                        { path: '/screens/users', element: <User /> },
                        { path: '/screens/logs', element: <Log /> },
                    ],
                },
                { path: '/dash', element: <Dashboard /> },
            ],
        },
    ])
}

export default AppRoutes
