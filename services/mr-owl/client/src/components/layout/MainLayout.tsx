import { Outlet } from 'react-router-dom'
import { Grid, GridItem } from '@chakra-ui/react'
import Navbar from '../navigation/navbar/Navbar'
import Footer from '../navigation/footer/Footer'
import Sidebar from '../navigation/sidebar/Sidebar'
import { FOOTER_HEIGHT, NAVBAR_HEIGHT, SIDEBAR_WIDTH } from '../../constants'

const MainLayout = () => {
    return (
        <Grid
            templateAreas={`
                "header header"
                "nav main"
                "nav footer"`}
            gridTemplateRows={`${NAVBAR_HEIGHT} 1fr ${FOOTER_HEIGHT}`}
            gridTemplateColumns={`${SIDEBAR_WIDTH} 1fr`}
            h="100vh"
            color={'brand.white'}
        >
            <GridItem bg="brand.bg_navbar" area={'header'}>
                <Navbar />
            </GridItem>
            <GridItem bg="brand.bg_sidebar" area={'nav'}>
                <Sidebar />
            </GridItem>
            <GridItem bg="brand.bg_main" area={'main'}>
                <Outlet />
            </GridItem>
            <GridItem bg="brand.bg_main" area={'footer'}>
                <Footer />
            </GridItem>
        </Grid>
    )
}

export default MainLayout
