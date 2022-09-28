import React from 'react';
import clsx from 'clsx';

import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';
import AppBar from '@material-ui/core/AppBar';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';

import { colors } from '../styles/globalStyles'

const useStyles = makeStyles(() => ({
    bigBox: {
        flexGrow: '1',
        height: '100vh',
    },
    appBar: {
        padding: "1rem",
        backgroundColor: colors.general01Color
    },
    title: {
        color: colors.general02Color
    },
    mainBox: {
        flexGrow: '1',
        marginTop: "6rem",
        background: 'white',
    },
    mainContainer: {
        flexGrow: '1',
        display: 'flex',
        flexDirection: 'column',
        padding: '2rem',
    },
}));

// This componenet is meant to be used by all pages to layout the page.
// It will render the App bar which is the top bar as well as a nav bar on the left for navigation
// This was partly referenced from https://github.com/mui-org/material-ui/tree/master/docs/src/pages/getting-started/templates/dashboard
const PageLayout = ({ children, title }) => {
    const classes = useStyles();

    return (
        <Box className={clsx(classes.bigBox)}>
            <AppBar className={clsx(classes.appBar)}>
                <Container maxWidth="lg">
                    <Typography
                        component="h1"
                        variant="h6"
                        color="inherit"
                        noWrap
                        className={clsx(classes.title)}
                    >
                        {title}
                    </Typography>
                </Container>
            </AppBar>
            <Box component="main" className={clsx(classes.mainBox)}>
                <Container maxWidth="lg">
                    {children}
                </Container>
            </Box>
        </Box>
    );
};

export default PageLayout;
