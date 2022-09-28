import React, { useState } from 'react';
import { useHistory } from "react-router-dom";
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';

import PageLayout from '../components/PageLayout';
import RequestService from '../requestService';
import { getEnvironmentAndCountryFromDomain } from '../helper';
import { colors } from '../styles/globalStyles';

const COUNTRY_ISD = (country) => {
    switch (country) {
        case "SG": return "65"
        case "JP": return "81"
        case "ID": return "62"
        case "TW": return "886"
        case "AU": return "61"
        default: return ""
    }
}

const useStyles = makeStyles(() => ({
    inputField: {
        marginBottom: '0.5rem',
    },
    container: {
        marginBottom: '1rem'
    },
    errorText: {
        color: colors.alert01Color100
    }
}))

const setLoggedInCookie = (history) => {
    // Set a isLoggedIn cookie that expires in 1 hour to track user's logged in state
    // This is needed as the browser doesnt have access to the X-auth token set by the server and is unable to derive the login state from that.
    const expiryDate = new Date()
    expiryDate.setTime(expiryDate.getTime() + (60*60*1000));
    document.cookie = `isLoggedIn=true;expires=${expiryDate.toUTCString()}`
    history.push("/default");
}

// Login component displays 2 login options for user.
// google login button gives all @circles.asia users read access
// freeipa login gives users with `quilt-ui` freeipa group write access
const Login = () => {
    const { host } = window.location
    const { country } = getEnvironmentAndCountryFromDomain(host)

    const history = useHistory();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [phoneNumber, setPhoneNumber] = useState("");
    const [errorMessage, setErrorMessage] = useState("");

    const classes = useStyles();

    const loginWithFreeipa = async (e) => {
        e.preventDefault()
        // Passing username and password up as header as the middleware that is handling auth isn't configured to easily parse the request body
        try {
            const res = await RequestService.post(`${process.env.QUILT_DOMAIN}api/v1/quilt/ui/auth`, {}, {
                headers: {
                "X-PHONE-NUMBER": phoneNumber,
                "X-ISD-CODE": COUNTRY_ISD(country),
                "X-Deviceid": process.env.DEVICE_ID,
                "username": username,
                "password": password,
                },
                withCredentials: true
            })
            if (res.status === 201) {
                setLoggedInCookie(history)
            }
        } catch (error) {
            setErrorMessage("Error logging in. Check that you used a valid phone number.")
        }
    }
    return (<PageLayout title="Login">
        <Grid container direction="row" className={clsx(classes.container)}>
            <Grid item xs={12}>
                <Typography variant="h5">{country} Phone number to use for this session</Typography>
                <TextField 
                    label="phonenumber" 
                    variant="outlined" 
                    value={phoneNumber}
                    className={clsx(classes.inputField)}
                    onChange={(event) => { setPhoneNumber(event.target.value) }} 
                />
            </Grid>
        </Grid>
        <Grid container direction="row" className={clsx(classes.container)}>
            <Grid item xs={12}>
                <Typography variant="h5">Write Access</Typography>
                <Typography variant="body1">Login using your freeipa account if you have quilt write access role to make changes.</Typography>
            </Grid>
            <Grid item xs={3}>
                <TextField 
                    label="username" 
                    variant="outlined" 
                    value={username}
                    className={clsx(classes.inputField)}
                    onChange={(event) => { setUsername(event.target.value) }} 
                />
            </Grid>
            <Grid item xs={3}>
                <TextField 
                    label="password" 
                    variant="outlined" 
                    type="password"
                    value={password} 
                    className={clsx(classes.inputField)}
                    onChange={(event) => { setPassword(event.target.value) }} 
                />
            </Grid>
            <Grid item xs={12}>
                <Button 
                    variant="contained" 
                    color="primary" 
                    onClick={loginWithFreeipa}
                    disabled={phoneNumber === "" || username === "" || password === ""}
                >
                    Login with Freeipa
                </Button>
            </Grid>
            {errorMessage && <Grid item xs={12}>
                <Typography className={clsx(classes.errorText)} variant="body1">Error: {errorMessage}</Typography>
            </Grid>}
        </Grid>
    </PageLayout>)
};

export default Login;
