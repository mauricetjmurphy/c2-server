import { ReactElement } from 'react'
import { Box, Button, Text } from '@chakra-ui/react'
import useStore from '../../../store'
import { useNavigate } from 'react-router-dom'
import { MenuState } from '../../../types/menu'

import { IconLayoutDashboard } from '@tabler/icons-react'
import { IconDeviceImacMinus } from '@tabler/icons-react'
import { IconEar } from '@tabler/icons-react'
import { IconUsersGroup } from '@tabler/icons-react'
import { IconLogs } from '@tabler/icons-react'
import { MENU_ICON_SIZE } from '../../../constants'

interface Tab {
    id: number
    name: string
    path: string
    icon: ReactElement
}

const tabs: Tab[] = [
    {
        id: 1,
        name: 'Dashboard',
        path: '/screens/dashboard',
        icon: <IconLayoutDashboard stroke="1px" size={MENU_ICON_SIZE} />,
    },
    {
        id: 2,
        name: 'Devices',
        path: '/screens/devices',
        icon: <IconDeviceImacMinus stroke="1px" size={MENU_ICON_SIZE} />,
    },
    {
        id: 3,
        name: 'Listeners',
        path: '/screens/listeners',
        icon: <IconEar stroke="1px" size={MENU_ICON_SIZE} />,
    },
    {
        id: 4,
        name: 'Users',
        path: '/screens/users',
        icon: <IconUsersGroup stroke="1px" size={MENU_ICON_SIZE} />,
    },
    {
        id: 5,
        name: 'Logs',
        path: '/screens/logs',
        icon: <IconLogs stroke="1px" size={MENU_ICON_SIZE} />,
    },
]

const Sidebar = () => {
    const navigate = useNavigate()

    const selectedId = useStore((state: MenuState) => state.selectedMenuItem)
    const setSelectedId = useStore(
        (state: MenuState) => state.setSelectedMenuItem
    )

    const tabSelected = (tab: Tab) => {
        setSelectedId(tab.id)
        navigate(tab.path)
    }

    return (
        <Box>
            {tabs.map((tab: Tab) => (
                <Button
                    key={tab.id}
                    colorScheme="brand.sidebar"
                    height={'120px'}
                    display={'flex'}
                    flexDirection={'column'}
                    justifyContent={'center'}
                    alignItems={'center'}
                    width={'100%'}
                    borderRadius={'0px'}
                    bg={
                        selectedId === tab.id
                            ? 'brand.bg_sidebar_selected'
                            : 'none'
                    }
                    borderLeft={selectedId === tab.id ? '3px solid' : 'none'}
                    borderColor={
                        selectedId === tab.id ? 'brand.purple' : 'transparent'
                    }
                    onClick={() => tabSelected(tab)}
                >
                    <Box color={'brand.white'}>
                        <Box
                            color={
                                selectedId === tab.id
                                    ? 'brand.purple'
                                    : 'brand.white'
                            }
                            pb={3}
                            display={'flex'}
                            justifyContent={'center'}
                        >
                            {tab.icon}
                        </Box>
                        <Text
                            color={
                                selectedId === tab.id
                                    ? 'brand.purple'
                                    : 'brand.white'
                            }
                            fontWeight={300}
                            fontSize={'14px'}
                        >
                            {tab.name}
                        </Text>
                    </Box>
                </Button>
            ))}
        </Box>
    )
}

export default Sidebar
