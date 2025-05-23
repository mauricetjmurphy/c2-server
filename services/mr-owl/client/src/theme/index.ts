import { extendTheme } from '@chakra-ui/react'

const theme = extendTheme({
    colors: {
        brand: {
            bg_main: '#26263e',
            bg_sidebar: '#232639',
            bg_sidebar_selected: '#1d1e30',
            bg_navbar: '#2f2f4b',
            bg_widget: '#25263b',

            white: '#e6e8e7',
            purple: '#7067de',
            green: '#4e9693',
            blue: '#6077a0',
            grey_light: '#dfddda80',
        },
        customGreen: {
            50: '#e6f9f880',
            100: '#c0eceb80',
            200: '#97dfdd80',
            300: '#6dd3d080',
            400: '#4ec8c480',
            500: '#4e969380',
            600: '#457b7980',
            700: '#3b626080',
            800: '#314a4880',
            900: '#26333080',
        },
    },
})

export default theme
