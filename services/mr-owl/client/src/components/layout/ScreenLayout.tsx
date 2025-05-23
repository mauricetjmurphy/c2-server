import { Box } from '@chakra-ui/react'
import { Outlet } from 'react-router-dom'

const ScreenLayout = () => {
    return (
        <Box p={'32px'}>
            <Outlet />
        </Box>
    )
}

export default ScreenLayout
