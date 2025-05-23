import { ReactNode } from 'react'
import { BrowserRouter as Router } from 'react-router-dom'
import { ChakraProvider } from '@chakra-ui/react'
import theme from '../theme'

interface AppProviderProps {
    children: ReactNode
}

const AppProvider = ({ children }: AppProviderProps) => {
    return (
        <Router>
            <ChakraProvider theme={theme}>{children}</ChakraProvider>
        </Router>
    )
}

export default AppProvider
