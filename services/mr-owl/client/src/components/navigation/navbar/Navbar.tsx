import { Box, Text } from '@chakra-ui/react'
import { APP_TITLE, SIDEBAR_WIDTH } from '../../../constants'

const Navbar = () => {
    return (
        <Box height={'100%'}>
            <Box
                width={SIDEBAR_WIDTH}
                height={'100%'}
                borderRight={'1px solid'}
                borderColor={'brand.grey_light'}
                display={'flex'}
                justifyContent={'center'}
                alignItems={'center'}
            >
                <Text fontSize={'32px'} fontWeight={'bold'}>
                    {APP_TITLE}
                </Text>
            </Box>
        </Box>
    )
}

export default Navbar
