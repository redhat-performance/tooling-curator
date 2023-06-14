import { createTheme } from '@mui/material/styles';
import { red } from '@mui/material/colors';

// A custom theme for this app
const theme = createTheme({
    palette: {
        primary: {
            main: '#556cd6',
        },
        secondary: {
            main: '#19857b',
        },
        warning: {
            main: '#FF6666',
        },
        error: {
            main: red.A400,
        },
    },
});

export default theme;