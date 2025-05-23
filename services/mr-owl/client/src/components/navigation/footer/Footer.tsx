import { Box, Text } from '@chakra-ui/react'

const Footer = () => {
    return (
        <Box
            height={'100%'}
            display={'flex'}
            justifyContent={'center'}
            alignItems={'center'}
        >
            <Text fontSize={'13px'}>Brought to you by Mr. Owl</Text>
        </Box>
    )
}

export default Footer
